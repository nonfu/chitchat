package models

import "time"

type Thread struct {
    Id        int
    Uuid      string
    Topic     string
    UserId    int
    CreatedAt time.Time
}

// format the CreatedAt date to display nicely on the screen
func (thread *Thread) CreatedAtDate() string {
    return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// get the number of posts in a thread
func (thread *Thread) NumReplies() (count int) {
    rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = ?", thread.Id)
    if err != nil {
        return
    }
    for rows.Next() {
        if err = rows.Scan(&count); err != nil {
            return
        }
    }
    rows.Close()
    return
}

// get posts to a thread
func (thread *Thread) Posts() (posts []Post, err error) {
    rows, err := Db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts where thread_id = ?", thread.Id)
    if err != nil {
        return
    }
    for rows.Next() {
        post := Post{}
        if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
            return
        }
        posts = append(posts, post)
    }
    rows.Close()
    return
}

// Get all threads in the database and returns it
func Threads() (threads []Thread, err error) {
    rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
    if err != nil {
        return
    }
    for rows.Next() {
        conv := Thread{}
        if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
            return
        }
        threads = append(threads, conv)
    }
    rows.Close()
    return
}

// Get a thread by the UUID
func ThreadByUUID(uuid string) (conv Thread, err error) {
    conv = Thread{}
    err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = ?", uuid).
        Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
    return
}

// Get the user who started this thread
func (thread *Thread) User() (user User) {
    user = User{}
    Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", thread.UserId).
        Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
    return
}