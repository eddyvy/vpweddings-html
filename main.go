package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	htmlContent, errFiles := createHtmlCache()

	if errFiles != nil {
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	for name := range htmlContent {
		route := "/" + name
		content := htmlContent[name]

		if name == "index" {
			route = "/"
		}

		if strings.HasSuffix(name, "/index") {
			route = strings.TrimSuffix(route, "/index")
		}

		r.Get(route, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, content)
		})
	}

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	fileServer(r, "/static", filesDir)

	port := 80
	serverAddr := fmt.Sprintf("0.0.0.0:%d", port)

	fmt.Printf("Server is running on port: %d", port)
	err := http.ListenAndServe(serverAddr, r)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func createHtmlCache() (map[string]string, error) {
	htmlFolder := "./html"
	cache := make(map[string]string)

	// Use the Walk function to traverse all files in the folder and subfolders
	err := filepath.Walk(htmlFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file has a .html extension
		if strings.HasSuffix(path, ".html") {
			// Read the file content
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Get the relative file name
			relPath, err := filepath.Rel(htmlFolder, path)
			if err != nil {
				return err
			}
			relPath = filepath.ToSlash(relPath)
			relPath = strings.TrimSuffix(relPath, ".html")

			// Store the file content in the map with the file name as the key
			cache[relPath] = string(content)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return cache, err
	}

	return cache, nil
}
