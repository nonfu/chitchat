package handlers

import (
    "fmt"
    "github.com/xueyuanjun/chitchat/models"
    "net/http"
)

// POST /thread/post
// 在指定群组下创建新主题
func PostThread(writer http.ResponseWriter, request *http.Request) {
    sess, err := session(writer, request)
    if err != nil {
        http.Redirect(writer, request, "/login", 302)
    } else {
        err = request.ParseForm()
        if err != nil {
            fmt.Println("Cannot parse form")
        }
        user, err := sess.User()
        if err != nil {
            fmt.Println("Cannot get user from session")
        }
        body := request.PostFormValue("body")
        uuid := request.PostFormValue("uuid")
        thread, err := models.ThreadByUUID(uuid)
        if err != nil {
            fmt.Println("Cannot read thread")
        }
        if _, err := user.CreatePost(thread, body); err != nil {
            fmt.Println("Cannot create post")
        }
        url := fmt.Sprint("/thread/read?id=", uuid)
        http.Redirect(writer, request, url, 302)
    }
}
