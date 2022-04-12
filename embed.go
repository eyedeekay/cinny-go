//go:build !gen
// +build !gen

package cinnygo

import (
	"embed"
	"net"
	"net/http"
	"path/filepath"

	"path"
)

//go:embed www

//go:generate go run -tags generate make.go
var Content embed.FS

type CinnyServer struct {
	Listener net.Listener
}

func (c *CinnyServer) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	cleanedPath := path.Clean(rq.URL.Path)
	if cleanedPath == "config.json" || cleanedPath == "/config.json" {
		//write content-type header json
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(200)
		rw.Write([]byte(c.ConfigJSON()))
	}
	contentPath := filepath.Join("www", cleanedPath)
	if contentPath == "www" {
		contentPath = filepath.Join("www", cleanedPath, "index.html")
	}
	// find the file in Content and if it exists, serve it. If not, 404
	f, err := Content.Open(contentPath)
	if err != nil {
		return
	}
	bytes := []byte{}
	_, err = f.Read(bytes)
	if err != nil {
		return
	}
	rw.Write(bytes)
}

func (c *CinnyServer) ConfigJSON() string {
	configJson := "{"
	configJson += "  \"defaultHomeserver\": 0"
	configJson += "  \"homeserverList\": ["
	configJson += "  \"" + c.Listener.Addr().String() + "\""
	configJson += "  ]"
	configJson += "}"
	return configJson
}
