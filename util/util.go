package util

import(
    "github.com/kless/goconfig/config"
    "github.com/hoisie/mustache"
    "database/sql"
     _ "github.com/jbarham/gopgsqldriver"
)

var Db, _ = sql.Open("postgres", "user=Steve dbname=goblog host=localhost port=5432")

type Entry struct {
    Id int
    Title, Content string
}

/*
* Handles rendering templates in a normalized context
*/
func RenderTemplate(template string, context map[string]interface{})string {
    c, _ := config.ReadDefault("goblog.conf")

    title, _ := c.String("general", "title")
    motto, _ := c.String("general", "motto")

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

