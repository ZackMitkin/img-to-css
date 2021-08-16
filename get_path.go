package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"img_to_css/marching_square"
)

func avgColor(img image.RGBA, path []image.Point) string {
	var red int
	var green int
	var blue int
	for _, point := range path {
		r, g, b, _ := img.At(point.X, point.Y).RGBA()
		red += int(r >> 8)
		green += int(g >> 8)
		blue += int(b >> 8)
	}

	R := red / len(path)
	G := green / len(path)
	B := blue / len(path)
	return fmt.Sprintf("rgb(%v, %v, %v)", R, G, B)
}

func getPath(polygon []image.Point, img image.RGBA) ([]image.Point, string) {
	var path []image.Point

	m := image.NewRGBA(img.Bounds())
	draw.Draw(m, m.Bounds(), &image.Uniform{C: color.Black}, image.Point{X: 0, Y: 0}, draw.Src)
	for _, px := range polygon {
		m.Set(px.X, px.Y, color.White)
	}
	contours := marching_square.Process(m, func(r, g, b, a uint32) bool {
		return r+g+b > 0
	})
	if contours == nil {
		return []image.Point{}, ""
	}
	for _, point := range contours {
		path = append(path, point)
	}
	return path, avgColor(img, polygon)
}
