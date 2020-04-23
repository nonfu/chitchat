package models

import (
    "crypto/rand"
    "crypto/sha1"
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    . "github.com/xueyuanjun/chitchat/config"
    "log"
)

var Db *sql.DB

func init() {
    var err error
    driver := ViperConfig.Db.Driver
    source := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true", ViperConfig.Db.User, ViperConfig.Db.Password,
        ViperConfig.Db.Address, ViperConfig.Db.Database)
    Db, err = sql.Open(driver, source)
    if err != nil {
        log.Fatal(err)
    }
    return
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
    u := new([16]byte)
    _, err := rand.Read(u[:])
    if err != nil {
        log.Fatalln("Cannot generate UUID", err)
    }

    // 0x40 is reserved variant from RFC 4122
    u[8] = (u[8] | 0x40) & 0x7F
    // Set the four most significant bits (bits 12 through 15) of the
    // time_hi_and_version field to the 4-bit version number.
    u[6] = (u[6] & 0xF) | (0x4 << 4)
    uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
    return
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
    cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
    return
}
