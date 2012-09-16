package main

import (
    "github.com/hoisie/mustache"
    "github.com/hoisie/web"
    "github.com/kless/goconfig/config"
    "database/sql"
     _ "github.com/jbarham/gopgsqldriver"
)

// Connect to the database
var db, err = sql.Open("postgres", "user=Steve dbname=goblog host=localhost port=5432")

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

/*
* Main page
*/
func Index() string {
    rows, _ := db.Query("SELECT id, title, content FROM entries ORDER BY id DESC")

    // Allocate space for 5 posts per page
    entries := []*Entry {}

    // Get the entries
    for i := 0; rows.Next(); i++ {
        var entry = new(Entry)

        rows.Scan(&entry.Id, &entry.Title, &entry.Content)
        entries = append(entries, entry)
    }

    var send = map[string]interface{} {
        "entries": entries,
    }

    return RenderTemplate("index.mustache", send)
}

func main() {
    web.Get("/", Index)
    web.Run("0.0.0.0:9999")
}
