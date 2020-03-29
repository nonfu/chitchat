package models

import "time"

type User struct {
    Id        int
    Uuid      string
    Name      string
    Email     string
    Password  string
    CreatedAt time.Time
}

// Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
    statement := "insert into sessions (uuid, email, user_id, created_at) values (?, ?, ?, ?)"
    stmtin, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmtin.Close()

    uuid := createUUID()
    stmtin.Exec(uuid, user.Email, user.Id, time.Now())

    stmtout, err := Db.Prepare("select id, uuid, email, user_id, created_at from sessions where uuid = ?")
    if err != nil {
        return
    }
    defer stmtout.Close()
    // use QueryRow to return a row and scan the returned id into the Session struct
    err = stmtout.QueryRow(uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
    return
}

// Get the session for an existing user
func (user *User) Session() (session Session, err error) {
    session = Session{}
    err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = ?", user.Id).
        Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
    return
}

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
    // Postgres does not automatically return the last insert id, because it would be wrong to assume
    // you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
    // information from postgres.
    statement := "insert into users (uuid, name, email, password, created_at) values (?, ?, ?, ?, ?)"
    stmtin, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmtin.Close()

    uuid := createUUID()
    stmtin.Exec(uuid, user.Name, user.Email, Encrypt(user.Password), time.Now())

    stmtout, err := Db.Prepare("select id, uuid, created_at from users where uuid = ?")
    if err != nil {
        return
    }
    defer stmtout.Close()
    // use QueryRow to return a row and scan the returned id into the User struct
    err = stmtout.QueryRow(uuid).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
    return
}

// Delete user from database
func (user *User) Delete() (err error) {
    statement := "delete from users where id = ?"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()

    _, err = stmt.Exec(user.Id)
    return
}

// Update user information in the database
func (user *User) Update() (err error) {
    statement := "update users set name = ?, email = ? where id = ?"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()

    _, err = stmt.Exec(user.Name, user.Email, user.Id)
    return
}

// Delete all users from database
func UserDeleteAll() (err error) {
    statement := "delete from users"
    _, err = Db.Exec(statement)
    return
}

// Get all users in the database and returns it
func Users() (users []User, err error) {
    rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
    if err != nil {
        return
    }
    for rows.Next() {
        user := User{}
        if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
            return
        }
        users = append(users, user)
    }
    rows.Close()
    return
}

// Get a single user given the email
func UserByEmail(email string) (user User, err error) {
    user = User{}
    err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?", email).
        Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
    return
}

// Get a single user given the UUID
func UserByUUID(uuid string) (user User, err error) {
    user = User{}
    err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = ?", uuid).
        Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
    return
}

// Create a new thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
    statement := "insert into threads (uuid, topic, user_id, created_at) values (?, ?, ?, ?)"
    stmtin, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmtin.Close()

    uuid := createUUID()
    stmtin.Exec(uuid, topic, user.Id, time.Now())

    stmtout, err := Db.Prepare("select id, uuid, topic, user_id, created_at from threads where uuid = ?")
    if err != nil {
        return
    }
    defer stmtout.Close()

    // use QueryRow to return a row and scan the returned id into the Session struct
    err = stmtout.QueryRow(uuid).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
    return
}

// Create a new post to a thread
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
    statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values (?, ?, ?, ?, ?)"
    stmtin, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmtin.Close()

    uuid := createUUID()
    stmtin.Exec(uuid, body, user.Id, conv.Id, time.Now())

    stmtout, err := Db.Prepare("select id, uuid, body, user_id, thread_id, created_at from posts where uuid = ?")
    if err != nil {
        return
    }
    defer stmtout.Close()

    // use QueryRow to return a row and scan the returned id into the Session struct
    err = stmtout.QueryRow(uuid).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
    return
}