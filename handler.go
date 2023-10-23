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
	router.GET("/about", h.getAbout)
	router.GET("/service", h.getService)
	router.GET("/contact", h.getContact)
	router.POST("/backend-endpoint", h.handleFrontendData)
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

	data["geo"] = loadJSON()
	data["marker"] = marker
	ctx.HTML(http.StatusOK, "home.gohtml", data)
}

func (h *Handler) getAbout(ctx *gin.Context) {
	data := h.defaultVars()
	ctx.HTML(http.StatusOK, "about.gohtml", data)
}

func (h *Handler) getContact(ctx *gin.Context) {
	data := h.defaultVars()
	ctx.HTML(http.StatusOK, "contact.gohtml", data)
}
func (h *Handler) getService(ctx *gin.Context) {
	data := h.defaultVars()
	ctx.HTML(http.StatusOK, "service.gohtml", data)
}
func (h *Handler) handleFrontendData(ctx *gin.Context) {

	var requestData struct {
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
		Name string  `json:"name"`
		Type string  `json:"type"`
	}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		// Handle error, possibly return an error response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//save to json --> in JsonHandler class

	feature := Features{
		Type:       "Feature",
		Properties: Properties{MarkerSymbol: "toilet", Name: requestData.Name},
		Geometry: Geometry{
			Coords: []float64{requestData.Lat, requestData.Long},
			Type:   "Point",
		},
	}

	saveJSON(feature)

	ctx.JSON(http.StatusOK, gin.H{"message": "Data received successfully"})

}
