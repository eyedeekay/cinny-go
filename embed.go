//go:build !gen
// +build !gen

package cinnygo

import (
	"embed"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"path"
)

//go:embed www

//go:generate go run -tags generate make.go
var Content embed.FS

type CinnyServer struct {
	HomeServer string
}

func (c *CinnyServer) Home(hostname string) string {
	if c.HomeServer != "" {
		_, err := url.Parse("http://" + c.HomeServer)
		if err != nil {
			return hostname
		}
		return c.HomeServer
	}
	return hostname
}

func (c *CinnyServer) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	cleanedPath := path.Clean(rq.URL.Path)
	if cleanedPath == "config.json" || cleanedPath == "/config.json" {
		//write content-type header json
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(200)
		log.Println("URL:", rq.Host)
		rw.Write([]byte(c.ConfigJSON(c.Home(rq.Host))))
		return
	}
	contentPath := filepath.Join("www", cleanedPath)
	if contentPath == "www" {
		contentPath = filepath.Join("www", cleanedPath, "index.html")
	}
	// find the file in Content and if it exists, serve it. If not, 404
	f, err := Content.ReadFile(contentPath)
	if err != nil {
		log.Println(err.Error())
	}
	contentType := ""
	if strings.HasSuffix(contentPath, ".js") {
		contentType = "text/javascript"
		rw.Header().Add("Content-Type", contentType)
	} else if strings.HasSuffix(contentPath, ".css") {
		contentType = "text/css"
		rw.Header().Add("Content-Type", contentType)
	} else {
		contentType := http.DetectContentType(f)
		rw.Header().Add("Content-Type", contentType)
	}
	log.Println("Serving:", contentPath, contentType)
	rw.Write(f)
}

func (c *CinnyServer) ConfigJSON(HomeServer string) string {
	configJson := `
	{
		"defaultHomeserver": 0,
		"homeserverList": [
			"` + HomeServer + `"
		]
	}`
	return configJson
}
