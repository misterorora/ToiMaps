package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"net/url"
)

//go:embed tpl/*
var tpl embed.FS

var Tpl = template.Must(template.New("").Funcs(map[string]any{
	"path_escape": url.PathEscape,
}).ParseFS(tpl, "tpl/*"))

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
	marker, err := requestMarkers()
	if err != nil {
		panic(err)
	}

	data["marker"] = marker
	ctx.HTML(http.StatusOK, "font.gohtml", data)
}
