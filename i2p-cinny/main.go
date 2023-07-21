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
	tls        = flag.Bool("tls", false, "TLS")
)

func main() {
	flag.Parse()
	garlic, err := onramp.NewGarlic("cinny-i2p", "127.0.0.1:7656", onramp.OPT_DEFAULTS)
	if err != nil {
		log.Fatal(err)
	}
	l, err := garlic.Listen()
	if err != nil {
		log.Fatal(err)
	}
	cs := &cinnygo.CinnyServer{
		HomeServer: *homeServer,
		TLS:        *tls,
	}
	if err := http.Serve(l, cs); err != nil {
		log.Fatal(err)
	}
}
