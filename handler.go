package main

import (
	"embed"
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
	router.GET("/about", h.getAbout)
	router.GET("/service", h.getService)
	router.GET("/contact", h.getContact)
	router.POST("/backend-endpoint", h.handleFrontendData)
	router.POST("/comment-endpoint", h.handleComment)
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
		Lat         float64 `json:"lat"`
		Long        float64 `json:"long"`
		Name        string  `json:"name"`
		Type        string  `json:"type"`
		Description string  `json:"description"`
	}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		// Handle error, possibly return an error response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//save to json --> in JsonHandler class
	feature := Features{
		Type:       "Feature",
		Properties: Properties{MarkerSymbol: "toilet", Type: requestData.Type, Name: requestData.Name, Description: requestData.Description},
		Geometry: Geometry{
			Coords: []float64{requestData.Long, requestData.Lat},
			Type:   "Point",
		},
	}
	saveJSON(feature)
	ctx.JSON(http.StatusOK, gin.H{"message": "Data received successfully"})
}

type Comment struct {
	Long    float64 `json:"long"`
	Lat     float64 `json:"lat"`
	Name    string  `json:"name"`
	Comment string  `json:"comment"`
}

func (h *Handler) handleComment(ctx *gin.Context) {
	comment := new(Comment)

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		// Handle error, possibly return an error response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(comment)

	addComment(comment)
	ctx.JSON(http.StatusOK, gin.H{"message": "Data received successfully"})
}
