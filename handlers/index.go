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
