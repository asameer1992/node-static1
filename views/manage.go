package views
import (
    "../util"
    "strconv"
)

func Manage() string {
    return util.RenderTemplate("manage.mustache", nil)
}

func ExistingEdit(val string) string {
    id, err := strconv.Atoi(val)
    if err != nil {
        return "Invalid or malformed id"
    }

    db := util.GetDb()
    row := db.QueryRow("SELECT id, title, content FROM entries WHERE id=$1 LIMIT 1", id)
    entry := new(util.Entry)
    err = row.Scan(&entry.Id, &entry.Title, &entry.Content)
    if err != nil {
        return err.Error()
    }

    send := map[string]interface{} {
        "Id": entry.Id,
        "Title": entry.Title,
        "Content": entry.Content,
    }

    return util.RenderTemplate("existing_edit.mustache", send);
}

