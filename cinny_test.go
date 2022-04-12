package cinnygo

import (
	"net"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		t.Error(err)
	}
	http.Serve(l, &CinnyServer{})
}
