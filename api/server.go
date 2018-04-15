package main

import (
	"net/http"
	"github.com/gernest/alien"
	"log"
	"OStatus/api/posts"
	"strconv"
)

func getOne(handler func(w http.ResponseWriter, r *http.Request, id int)) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request){
		params := alien.GetParams(r)
		id, error := strconv.Atoi(params.Get("id"))
		if error != nil {
			w.WriteHeader(400)
			w.Write([]byte("Bad Request"))
			return
		}
		handler(w, r, id)
	}
}

func jsonMiddleware(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	m := alien.New()

	postsGroup := m.Group("/posts")
	postsGroup.Use(jsonMiddleware)
	postsGroup.Get("/", posts.List)
	postsGroup.Get("/:id", getOne(posts.GetOne))

	log.Fatal(http.ListenAndServe(":8090", m))
}
