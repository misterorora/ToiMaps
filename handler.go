package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"net/url"
	"os"
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
	marker, err := "Test", error(nil) //requestMarkers()
	if err != nil {
		panic(err)
	}

	file, err1 := os.Open("map.geojson")
	if err1 != nil {
		fmt.Println("Error opening file:", err1)
		return
	}
	defer file.Close() // Close the file when we're done

	// Create a JSON decoder
	decoder := json.NewDecoder(file)

	var geo map[string]interface{} // Change the type according to your JSON structure
	if err1 := decoder.Decode(&geo); err1 != nil {
		fmt.Println("Error decoding JSON:", err1)
		return
	}

	fmt.Println(geo)
	data["geo"] = geo
	data["marker"] = marker
	ctx.HTML(http.StatusOK, "font.gohtml", data)
}
