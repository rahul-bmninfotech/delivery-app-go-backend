package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//NewDeliveryRouter returns a mux router
func NewDeliveryRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.NewRoute().Name(route.Name).Path(route.Path).Methods(route.Method).HandlerFunc(route.HandlerFunc)
	}
	router.Use(middlewareLogger)
	return router
}

func middlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
