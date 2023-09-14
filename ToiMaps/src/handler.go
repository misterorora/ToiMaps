package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//go:embed tpl/*
var tpl embed.FS

var Tpl = template.Must(template.New("").Funcs(map[string]any{}).ParseFS(tpl, "tpl/*"))

type Handler struct {
	Name string
}

func (h *Handler) Register(router gin.IRouter) {
	router.GET("/home", h.getHome)
}

func (h *Handler) defaultVars() map[string]any {
	return map[string]any{}
}

func (h *Handler) getHome(ctx *gin.Context) {
	data := h.defaultVars()
	marker := "direct access database -> formatted to marker struct -> databaseHandler.GetMarkers"

	data["marker"] = marker
	ctx.HTML(http.StatusOK, "home.gohtml", data)
}
