package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.SetHTMLTemplate(Tpl)
	router.UseRawPath = true

	handler := Handler{
		Name: "Handler",
	}

	handler.Register(router)

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/home")
	})

	router.Run(":8080") //listen and serve on 0.0.0.0:8080
}

/*
{{range $marker := .marker .}}
    // L.marker is a low-level marker constructor in Leaflet.
    var marker = L.marker([{{$marker.latitude}}, {{$marker.longitude}}], {
        icon: L.mapbox.marker.icon({
            'marker-size': 'large',
            'marker-symbol': 'toilet',
            'marker-color': '#fa0'
        })
    }).addTo(map);
    {{end}}


*/
