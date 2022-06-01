package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"project/freelance/danspro/simpleProject/app/controller"
	"project/freelance/danspro/simpleProject/app/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Post("/login", controller.Login)

	r.Group(func(r chi.Router) {
		r.Use(middleware.IsAuthorized)
		r.Route("/job", func(r chi.Router) {
			r.Get("/list", controller.GetJobList)
			r.Get("/detail/{id}", controller.GetJobDetail)
		})
	})

	http.ListenAndServe(":1234", r)
}
