package views
import "../util"

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

