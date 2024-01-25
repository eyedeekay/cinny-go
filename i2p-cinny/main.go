package main

import (
	"flag"
	"log"
	"net/http"

	cinnygo "github.com/eyedeekay/cinny-go"
	"github.com/eyedeekay/onramp"
)

var (
	homeServer = flag.String("home", "jfbg4uhej6ahlyv5jjpnqi3tlunh4ozjkdq4khtt6ft467zcooka.b32.i2p", "Home server")
	tls        = flag.Bool("tls", true, "TLS")
)

func main() {
	flag.Parse()
	garlic, err := onramp.NewGarlic("cinny-i2p", "127.0.0.1:7656", onramp.OPT_WIDE)
	if err != nil {
		log.Fatal(err)
	}
	defer garlic.Close()
	l, err := garlic.ListenTLS()
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	cs := &cinnygo.CinnyServer{
		HomeServer: *homeServer,
		TLS:        *tls,
	}
	if err := http.Serve(l, cs); err != nil {
		log.Fatal(err)
	}
}
