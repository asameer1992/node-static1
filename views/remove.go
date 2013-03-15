package views
import (
    "github.com/hoisie/web"
    "../util"
    "strconv"
)

func Remove(ctx *web.Context, val string) string {
    id, err := strconv.Atoi(val)
    if err != nil {
        return "Invalid or malformed id"
    }

    db := util.GetDb()

    _, submit_exists := ctx.Params["doit"]
    if submit_exists {
        db.Exec("DELETE FROM entries WHERE id=$1", id)
        ctx.Redirect(302, "/manage/existing")
        return "Redirect"
    }

    // Get the post
    row := db.QueryRow("SELECT title FROM entries WHERE id=$1", id)
    entry := new(util.Entry)
    row.Scan(&entry.Title)

    send := map[string]interface{} {
        "Title": entry.Title,
    }

    return util.RenderTemplate("remove.mustache", send)
}
