package views
import (
    "../util"
    "strings"
)

/*
* Main page
*/
func Index() string {
    db := util.GetDb()
    rows, err := db.Query("SELECT id, title, content FROM entries ORDER BY id DESC")

    if err != nil {
        return err.Error()
    }

    entries := []*util.Entry {}

    // Get the entries
    for i := 0; rows.Next(); i++ {
        var entry = new(util.Entry)

        rows.Scan(&entry.Id, &entry.Title, &entry.Content)

        // Parse newlines
        entry.Content = strings.Replace(entry.Content, "\r\n\r\n", "</p><p>", -1);

        entries = append(entries, entry)
    }

    send := make(map[string]interface{})

    if len(entries) == 0 {
        send["entries"] = false
    } else {
        send["entries"] = entries
    }

    return util.RenderTemplate("index.mustache", send)
}

