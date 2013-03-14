package util

import(
    "github.com/kless/goconfig/config"
    "github.com/hoisie/mustache"
    "database/sql"
     _ "github.com/jbarham/gopgsqldriver"
     "fmt"
)

var Config, _ = config.ReadDefault("goblog.conf")
var db *sql.DB

type Entry struct {
    Id int
    Title, Content string
}

func GetDb() *sql.DB {
    if(db != nil){
        return db
    }

    var db_username, _ = Config.String("db", "username")
    var db_password, _ = Config.String("db", "password")
    var db_database, _ = Config.String("db", "database")
    var db_hostname, _ = Config.String("db", "hostname")
    var db_port, _ = Config.String("db", "port")

    var db, err = sql.Open("postgres", "user=" + db_username + " password=" + db_password + " dbname=" + db_database + " host=" + db_hostname + " port=" + db_port)

    if(err != nil) {
        fmt.Println("[db] Error: " + err.Error())
    }

    return db
}

/*
* Handles rendering templates in a normalized context
*/
func RenderTemplate(template string, context map[string]interface{})string {
    title, _ := Config.String("general", "title")
    motto, _ := Config.String("general", "motto")

    var send = map[string]interface{} {
        "title": title,
        "motto": motto,
    }
    // Append all values of context to the global context
    for key, val := range context {
        send[key] = val
    }

    return mustache.RenderFileInLayout("templates/" + template, "templates/base.mustache", send)
}

