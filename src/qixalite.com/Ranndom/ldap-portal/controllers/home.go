package controllers

import (
    "gopkg.in/macaron.v1"
)

func Home(ctx *macaron.Context) {
    ctx.Data["Title"] = "Home"
    ctx.HTML(200, TMPL_HOME)
}

func NotFound(ctx *macaron.Context) {
    ctx.HTML(404, TMPL_404)
}

