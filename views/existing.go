package views
import (
    "../util"
    "strconv"
)

func Existing() string {
    db := util.GetDb()
    rows, _ := db.Query("SELECT id, title, content FROM entries ORDER BY id DESC")

    entries := []*util.Entry {}
    for i := 0; rows.Next(); i++ {
        var entry = new(util.Entry)
        rows.Scan(&entry.Id, &entry.Title, nil)

        entries = append(entries, entry)
    }

    send := make(map[string]interface{})
    if len(entries) == 0 {
        send["entries"] = false
    } else {
        send["entries"] = entries
    }

    return util.RenderTemplate("existing.mustache", send)
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
