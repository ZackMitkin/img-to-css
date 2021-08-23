package css

import (
	"fmt"
	"image"
)

func WritePolygonsVideo(frame int, img image.RGBA, paths [][]image.Point, colors []string) (string, string) {
	width := float32(img.Bounds().Dx())
	height := float32(img.Bounds().Dy())
	var css string
	var html string
	for idx, path := range paths {
		if len(path) == 0 {
			continue
		}
		className := fmt.Sprintf("frame-%d-element-%d", idx, frame)
		polygon := ""
		for idx := range path {
			x1 := float32(path[idx].X + 1)
			y1 := float32(path[idx].Y + 1)
			polygon += fmt.Sprintf("%f%% %f%%,", 100*x1/width, 100*y1/height)
		}
		x1 := float32(path[0].X)
		y1 := float32(path[0].Y)
		polygon += fmt.Sprintf(" %.2f%% %.2f%%", 100*x1/width, 100*y1/height)
		c := colors[idx]

		class := fmt.Sprintf(".%v {clip-path: polygon(%v); background-color: %v}", className, polygon, c)
		html += fmt.Sprintf(`<div class="component %v"></div>`, className)
		css += class
	}
	return html, css
}

func WritePolygons(img image.RGBA, paths [][]image.Point, colors []string) string {
	width := float32(img.Bounds().Dx())
	height := float32(img.Bounds().Dy())
	css := fmt.Sprintf(`
			.component {width: %vvw; height: %vvw; position: absolute} \n`, 100, 100*(height/width))
	htmlBody := ""
	for idx, path := range paths {
		if len(path) == 0 {
			continue
		}
		className := fmt.Sprintf("element-%d", idx)
		polygon := ""
		for idx := range path {
			x1 := float32(path[idx].X)
			y1 := float32(path[idx].Y)
			polygon += fmt.Sprintf("%f%% %f%%,", 100*x1/width, 100*y1/height)
		}
		x1 := float32(path[0].X)
		y1 := float32(path[0].Y)
		polygon += fmt.Sprintf(" %.2f%% %.2f%%", 100*x1/width, 100*y1/height)
		c := colors[idx]

		class := fmt.Sprintf(".%v {clip-path: polygon(%v); background-color: %v} \n", className, polygon, c)
		htmlBody += fmt.Sprintf(`<div class="component %v"></div>`, className)
		css += class
	}
	html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
		<title>img to css</title>
		</head>
		<style>
		%v
		</style>
		<body>
		%v
		</body>
		</html>`, css, htmlBody)
	return html
}
