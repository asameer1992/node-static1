package main

import (
    "github.com/hoisie/web"
     "./views"
)

func main() {
    web.Get("/", views.Index)
    web.Get("/manage", views.Manage)
    web.Get("/manage/create", views.Create)
    web.Post("/manage/create", views.Create)
    web.Get("/manage/existing", views.Existing)
    web.Get("/manage/existing/remove/(.*)", views.Remove)
    web.Post("/manage/existing/remove/(.*)", views.Remove)
    web.Get("/manage/existing/(.*)", views.ExistingEdit)
    web.Post("/manage/existing/(.*)", views.ExistingEdit)
    web.Run("0.0.0.0:9999")
}
