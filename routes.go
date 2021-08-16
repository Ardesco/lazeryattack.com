package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Logger(handler http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		start := time.Now()
		log.Info(request.Method, "\t", request.RequestURI, "\t", name, "\t", time.Since(start))
		handler.ServeHTTP(responseWriter, request)
	})
}

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Host(originURL)
	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./web/img"))))
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./web/css"))))
	router.Handle("/", http.FileServer(http.Dir("./web")))
	router.NotFoundHandler = Logger(http.HandlerFunc(NotFound), "NotFound")

	return router
}

func NotFound(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusNotFound)
	http.ServeFile(responseWriter, request, "./web/404.html")
}
