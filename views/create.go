package views
import (
    "github.com/hoisie/web"
    "../util"
)

func Create(ctx *web.Context) string {
    // Check to see if we're actually publishing
    title, exists_title := ctx.Params["title"]
    content, exists_content := ctx.Params["content"]

    var send = map[string]interface{} {
        "show_success": false,
    }

    // We are! So let's save it
    if exists_title && exists_content {
        // Insert the row
        _, err := util.Db.Exec("INSERT INTO entries (title, content) VALUES($1, $2)", title, content)

        if err != nil {
            return util.RenderTemplate("error.mustache", nil)
        }

        send["show_success"] = true
    }

    return util.RenderTemplate("create.mustache", send)
}

