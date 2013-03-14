package views
import "../util"

func Manage() string {
    return util.RenderTemplate("manage.mustache", nil)
}

