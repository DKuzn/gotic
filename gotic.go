package main

import (
	"fmt"
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
	fs := http.FileServer(http.Dir(dir))

	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CORS(origins, methods)(r))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
