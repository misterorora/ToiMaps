package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type GeoMap struct {
	Type     string     `json:"type"`
	Features []Features `json:"features"`
}

type Features struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

type Properties struct {
	MarkerSymbol string   `json:"marker-symbol"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Type         []string `json:"type"` //gender
	Recension    []string `json:"recension"`
}

type Geometry struct {
	Coords []float64 `json:"coordinates"`
	Type   string    `json:"type"`
}

func loadJSON() map[string]interface{} {
	file, err := os.Open("map.geojson")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close() // Close the file when we're done
	// Create a JSON decoder
	decoder := json.NewDecoder(file)

	var geo map[string]interface{} // Change the type according to your JSON structure
	if err1 := decoder.Decode(&geo); err1 != nil {
		fmt.Println("Error decoding JSON:", err1)
		return nil
	}
	return geo
}

func saveJSON(newFeature Features) {
	geo := unMarshJSON()
	found := false

	for _, feature := range geo.Features {
		if feature.Geometry.Coords[0] == newFeature.Geometry.Coords[0] {
			if feature.Geometry.Coords[1] == newFeature.Geometry.Coords[1] {
				found = true
				break
			}
		}

	}
	if !found {
		geo.Features = append(geo.Features, newFeature)
	}

	if geo.Type == "" {
		geo.Type = "FeatureCollection"
	}

	jsonData, err := json.Marshal(geo)
	if err != nil {
		fmt.Println("error while marshalling: ", err)
		return
	}

	err = os.WriteFile("map.geojson", jsonData, 0644)
	if err != nil {
		fmt.Println("error while writing to file: ", err)
		return
	}

	fmt.Println("Daten gespeichert")
}

func unMarshJSON() GeoMap {
	file, err := os.Open("map.geojson")
	if err != nil {
		fmt.Println("Error while opening file: ", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error while reading: ", err)
	}
	var geo GeoMap
	if err := json.Unmarshal(content, &geo); err != nil {
		fmt.Println("error while unmarshalling: ", err)
	}

	return geo
}

func addComment(comment *Comment) {
	geo := unMarshJSON()
	for i, feature := range geo.Features {
		if feature.Geometry.Coords[0] == comment.Long {
			if feature.Geometry.Coords[1] == comment.Lat {
				if feature.Properties.Name == comment.Name {

					for _, old_recension := range feature.Properties.Recension {
						if old_recension == comment.Comment {
							return
						}
					}
					array := feature.Properties.Recension
					array = append(array, comment.Comment)
					geo.Features[i].Properties.Recension = array
					break
				}
			}
		}

	}

	jsonData, err := json.Marshal(geo)
	if err != nil {
		fmt.Println("error while marshalling: ", err)
		return
	}

	err = os.WriteFile("map.geojson", jsonData, 0644)
	if err != nil {
		fmt.Println("error while writing to file: ", err)
		return
	}
}
