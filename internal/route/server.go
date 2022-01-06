package route

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/gg-tools/markdown-server/internal/utils/bytes"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday/v2"
)

func static(staticRoot string) http.Handler {
	return http.FileServer(http.Dir(staticRoot))
}

func pages(markdownRoot string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var titles []string
		err := filepath.Walk(markdownRoot, func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			title := filepath.ToSlash(strings.TrimPrefix(path, markdownRoot))
			titles = append(titles, title)
			return nil
		})
		if err != nil {
			_, _ = w.Write([]byte(fmt.Sprintf("list markdown failed: %s", err)))
			return
		}

		res, _ := json.Marshal(titles)
		w.Write(res)
	}
}

func page(markdownRoot string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Query().Get("path")
		content := filepath.Join(markdownRoot, path)
		input, err := ioutil.ReadFile(content)
		input = bytes.ToUnix(input)
		if ok := utf8.Valid(input); !ok {
			_, _ = w.Write([]byte(fmt.Sprintf("markdown not utf8: %s", err)))
			return
		}

		if err != nil {
			_, _ = w.Write([]byte(fmt.Sprintf("read markdown failed: %s", err)))
			return
		}
		output := blackfriday.Run(input)
		w.Write(output)
	}
}

func routes(staticRoot, markdownRoot string) {
	r := mux.NewRouter()
	r.HandleFunc("/pages", pages(markdownRoot))
	r.HandleFunc("/page", page(markdownRoot))
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" || path == "/index.html" {
			http.ServeFile(w, r, filepath.Join(staticRoot, "/index.html"))
			return
		}

		ext := filepath.Ext(path)
		switch ext {
		case ".js", ".css":
			http.ServeFile(w, r, filepath.Join(staticRoot, path))
			return
		default:
			http.ServeFile(w, r, filepath.Join(markdownRoot, path))
			return
		}
	})

	http.Handle("/", r)
}

func Serve(bindAddr, staticRoot, markdownRoot string) {
	routes(staticRoot, markdownRoot)
	log.Fatal(http.ListenAndServe(bindAddr, nil))
}
