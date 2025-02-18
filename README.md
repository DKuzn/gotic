# gotic

**_A simple static server written in Go._**

## Name

I just put together words **golang** and **static**. 

## Motivation

I've started this project to create a simple tool for serving static files. This tool will help my students, who are learning frontend development, write web applications in a "production-like" manner and test them easily, without requiring backend knowledge.

Another goal is that I plan to use this tool in my own web applications so that I don't have to write an additional endpoint in the backend just to serve static files. I hope to add SPA (Single Page Application) support in the future, which would make me very happy.

## Installation

Download a standalone binary from [GitHub Releases Page](https://github.com/DKuzn/gotic/releases) for Windows or Linux. After downloading and extracting the archive for your platform, place the `gotic` or `gotic.exe` binary in your `PATH` to run `gotic` from any location.

## Usage

```bash
user@localhost:~$ gotic --dir ./path/to/static/files --port 8080
```

## Documentation

```bash
user@localhost:~$ gotic --help
A simple static server written in Go.

Usage:
  gotic [flags]

Flags:
  -d, --dir string   The directory is used to store static files. (default "./static/")
  -h, --help         help for gotic
  -p, --port int     The port is used to listen requests. (default 8080)
```

## Tips

The path in HTML links should start from the server's base path (`/`) rather than the file system path. For example, if you used an absolute path like `/home/user/site/static/page.html` in a link and are now serving the `.../static/...` directory, you should only include `/page.html` in the link. This is because the server is already aware of the `.../static/...` folder.