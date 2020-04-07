package handlers

import (
    "errors"
    "fmt"
    "github.com/xueyuanjun/chitchat/models"
    "html/template"
    "log"
    "net/http"
    "os"
    "strings"
)

var logger *log.Logger

func init()  {
    file, err := os.OpenFile("logs/chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("Failed to open log file", err)
    }
    logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

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

// 记录日志信息
func info(args ...interface{}) {
    logger.SetPrefix("INFO ")
    logger.Println(args...)
}

func danger(args ...interface{}) {
    logger.SetPrefix("ERROR ")
    logger.Println(args...)
}

func warning(args ...interface{}) {
    logger.SetPrefix("WARNING ")
    logger.Println(args...)
}

// 异常处理统一重定向到错误页面
func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
    url := []string{"/err?msg=", msg}
    http.Redirect(writer, request, strings.Join(url, ""), 302)
}
