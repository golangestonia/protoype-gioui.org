package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/loov/watchrun/watch"
	"github.com/loov/watchrun/watchjs"
)

func DisableCache(w http.ResponseWriter) {
	w.Header().Set("Expires", time.Unix(0, 0).Format(time.RFC1123))
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")
}

func main() {
	listen := flag.String("listen", "127.0.0.1:8081", "address to listen")
	monitor := flag.String("monitor", "templates", "directory to monitor changes")
	serve := flag.String("serve", "templates", "directory to serve content")

	flag.Parse()

	if !filepath.IsAbs(*monitor) {
		abs, err := filepath.Abs(*monitor)
		if err == nil {
			*monitor = abs
		}
	}

	if !filepath.IsAbs(*serve) {
		abs, err := filepath.Abs(*serve)
		if err == nil {
			*serve = abs
		}
	}

	router := mux.NewRouter()
	router.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			DisableCache(w)
			h.ServeHTTP(w, r)
		})
	})

	staticServer := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", staticServer))
	router.Handle("favicon.ico", staticServer)
	router.Handle("/~watch.js", watchjs.NewServer(watchjs.Config{
		Monitor: []string{filepath.Join(*monitor, "**")},
		Ignore:  watchjs.DefaultIgnore,
		OnChange: func(change watch.Change) (string, watchjs.Action) {
			// When change is in staticDir, we instruct the browser live (re)inject the file.
			if url, ok := watchjs.FileToURL(change.Path, *monitor, "/"); ok {
				if filepath.Ext(change.Path) == ".css" {
					return url, watchjs.LiveInject
				}
				return url, watchjs.ReloadBrowser
			}
			return "/" + filepath.ToSlash(change.Path), watchjs.ReloadBrowser
		},
	}))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(allTemplates("templates")...)
		if err != nil {
			errorTemplate.Execute(w, err.Error())
			return
		}

		path := r.URL.Path
		if path == "/" || path == "" {
			path = "index"
		}
		path += ".html"

		err = t.ExecuteTemplate(w, path, nil)
		if err != nil {
			log.Println(err)
		}
	})

	fmt.Println("Listening on:", "http://"+*listen)
	fmt.Println("Monitoring:", *monitor)
	err := http.ListenAndServe(*listen, router)
	if err != nil {
		log.Fatal(err)
	}
}

func allTemplates(dir string) []string {
	var all []string
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".html" {
			return nil
		}

		all = append(all, path)

		return nil
	})
	return all
}

var errorTemplate = template.Must(template.New("error").Parse(
	`<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">

	<meta name="viewport" content="width=device-width, initial-scale=1">
	<script src="/~watch.js"></script>
</head>
<body>
{{.}}
</body>
</html>
`))
