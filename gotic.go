package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var dir string
var port int

var rootCmd = &cobra.Command{
	Use:  "gotic",
	Long: "A simple static server written in Go.",
	Run:  serveStatic,
}

func init() {
	rootCmd.Flags().StringVarP(&dir, "dir", "d", "./static/", "The directory is used to store static files.")
	rootCmd.Flags().IntVarP(&port, "port", "p", 8080, "The port is used to listen requests.")
}

func serveStatic(cmd *cobra.Command, args []string) {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	fs := http.FileServer(http.Dir(dir))

	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Printf("gotic server is running on port: %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CORS(origins, methods)(r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RemoteAddr, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
