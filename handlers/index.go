package handlers

import (
    "github.com/xueyuanjun/chitchat/models"
    "net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
    threads, err := models.Threads();
    if err == nil {
        _, err := session(writer, request)
        if err != nil {
            generateHTML(writer, threads, "layout", "navbar", "index")
        } else {
            generateHTML(writer, threads, "layout", "auth.navbar", "index")
        }
    }
}

func Err(writer http.ResponseWriter, request *http.Request)  {
    vals := request.URL.Query()
    _, err := session(writer, request)
    if err != nil {
        generateHTML(writer, vals.Get("msg"), "layout", "navbar", "error")
    } else {
        generateHTML(writer, vals.Get("msg"), "layout", "auth.navbar", "error")
    }
}