package views
import (
    "github.com/hoisie/web"
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

func ExistingEdit(ctx *web.Context, val string) string {
    id, err := strconv.Atoi(val)
    if err != nil {
        return "Invalid or malformed id"
    }

    db := util.GetDb()

    title, exists_title := ctx.Params["title"]
    content, exists_content := ctx.Params["content"]
    success := false
    if exists_title && exists_content {
        _, err = db.Exec("UPDATE entries SET title=$1, content=$2 WHERE id=$3", title, content, id)
        if err != nil {
            return err.Error()
        }
        success = true
    }

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
        "show_success": success,
    }

    return util.RenderTemplate("existing_edit.mustache", send);
}
