package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func Main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hello world"))
		if err != nil {
			return
		}
	})

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		obj := map[string]string{"message": "hello world"}
		render.JSON(w, r, obj)
	})

	r.With(myMiddleware).Post("/product", func(w http.ResponseWriter, r *http.Request) {
		var product Product
		err := render.DecodeJSON(r.Body, &product)
		if err != nil {
			return
		}
		product.ID = 1

		productJSON, _ := json.Marshal(product)
		println("request", r.Method, " URL", r.RequestURI, " product", string(productJSON))
		render.JSON(w, r, product)
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}

}

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("before")
		next.ServeHTTP(w, r)
		println("after")
	})
}
