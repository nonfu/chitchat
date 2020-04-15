package config

import (
    "encoding/json"
    "github.com/nicksnyder/go-i18n/v2/i18n"
    "golang.org/x/text/language"
    "log"
    "os"
    "sync"
)

type App struct {
    Address      string
    Static       string
    Log          string
    Locale       string
    Language     string
}

type Database struct {
    Driver      string
    Address        string
    Database    string
    User        string
    Password    string
}

type Configuration struct {
    App App
    Db  Database
    LocaleBundle *i18n.Bundle
}

var config *Configuration
var once sync.Once

// 通过单例模式初始化全局配置
func LoadConfig() *Configuration {
    once.Do(func() {
        file, err := os.Open("config.json")
        if err != nil {
            log.Fatalln("Cannot open config file", err)
        }
        decoder := json.NewDecoder(file)
        config = &Configuration{}
        err = decoder.Decode(config)
        if err != nil {
            log.Fatalln("Cannot get configuration from file", err)
        }
        // 本地化初始设置
        bundle := i18n.NewBundle(language.English)
        bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
        bundle.MustLoadMessageFile(config.App.Locale + "/active.en.json")
        bundle.MustLoadMessageFile(config.App.Locale + "/active." + config.App.Language + ".json")
        config.LocaleBundle = bundle
    })
    return config
}
