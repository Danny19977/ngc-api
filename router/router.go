package router

import (
    "github.com/gorilla/mux"
    "project/pkg/controller"
    "project/pkg/middleware"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()
    r.Use(middleware.Logging)

    r.HandleFunc("/api/support", controller.CreateSupport).Methods("POST")
    r.HandleFunc("/api/courses", controller.ListCourses).Methods("GET")
    r.HandleFunc("/api/blog", controller.CreateBlogPost).Methods("POST")
    // Add other routes here

    return r
}
