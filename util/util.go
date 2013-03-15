package util

import (
    "github.com/msbranco/goconfig"
    "github.com/hoisie/mustache"
    "database/sql"
     _ "github.com/jbarham/gopgsqldriver"
     "fmt"
)

var Config, _ = goconfig.ReadConfigFile("goblog.conf")
var db *sql.DB

type Entry struct {
    Id int
    Title, Content string
}

func GetDb() *sql.DB {
    if db != nil {
        return db
    }

    var db_username, _ = Config.GetString("db", "username")
    var db_password, _ = Config.GetString("db", "password")
    var db_database, _ = Config.GetString("db", "database")
    var db_hostname, _ = Config.GetString("db", "hostname")
    var db_port, _ = Config.GetString("db", "port")

    var db, err = sql.Open("postgres", "user=" + db_username + " password=" + db_password + " dbname=" + db_database + " host=" + db_hostname + " port=" + db_port)

    if err != nil {
        fmt.Println("[db] Error: " + err.Error())
    }

    return db
}

/*
* Handles rendering templates in a normalized context
*/
func RenderTemplate(template string, context map[string]interface{})string {
    title, _ := Config.GetString("general", "title")
    motto, _ := Config.GetString("general", "motto")

    var send = map[string]interface{} {
        "blog_title": title,
        "blog_motto": motto,
    }
    // Append all values of context to the global context
    for key, val := range context {
        send[key] = val
    }

    return mustache.RenderFileInLayout("templates/" + template, "templates/base.mustache", send)
}

