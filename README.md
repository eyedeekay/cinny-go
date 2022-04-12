# cinny-go

A go server embedding a `cinny` Matrix client. Used to provide a single web interface
to a single Matrix homeserver. In my case, a Dendrite server running over I2P.

## Usage:

```Go
package main

import (
	"net"
	"net/http"
	"log"
)

func main() {
    l, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatal(err)
	}
    cs := &CinnyServer{
        // replace "matrix.org" with your homeserver here
        "matrix.org",
    }
	http.Serve(l, cs)
}
```
