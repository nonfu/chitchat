package handlers

import (
    "errors"
    "fmt"
    "github.com/xueyuanjun/chitchat/models"
    "html/template"
    "net/http"
)

// Checks if the user is logged in and has a session, if not err is not nil
func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
    cookie, err := request.Cookie("_cookie")
    if err == nil {
        sess = models.Session{Uuid: cookie.Value}
        if ok, _ := sess.Check(); !ok {
            err = errors.New("Invalid session")
        }
    }
    return
}

// parse HTML templates
// pass in a list of file names, and get a template
func parseTemplateFiles(filenames ...string) (t *template.Template) {
    var files []string
    t = template.New("layout")
    for _, file := range filenames {
        files = append(files, fmt.Sprintf("views/%s.html", file))
    }
    t = template.Must(t.ParseFiles(files...))
    return
}

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
    var files []string
    for _, file := range filenames {
        files = append(files, fmt.Sprintf("views/%s.html", file))
    }

    templates := template.Must(template.ParseFiles(files...))
    templates.ExecuteTemplate(writer, "layout", data)
}

// version
func Version() string {
    return "0.1"
}