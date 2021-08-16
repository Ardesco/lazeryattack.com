package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const originURL = "https://local.heimdall.uk"

//const originURL = "https://www.lazeryattack.com"
const port = 8080

func main() {
	log.Info("Running on ", originURL, ":", port)
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%v", port), "local.crt", "local.key", Router()))
}
