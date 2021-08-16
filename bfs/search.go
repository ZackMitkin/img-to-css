package bfs

import (
	"image"
	"image/color"
	"log"
	"math"
)

const colorDiff = 40

func similarColor(pointA, pointB color.Color) bool {
	R1, G1, B1, _ := pointA.RGBA()
	R2, G2, B2, _ := pointB.RGBA()
	red := math.Abs(float64(R1>>8)-float64(R2>>8)) <= colorDiff
	green := math.Abs(float64(G1>>8)-float64(G2>>8)) <= colorDiff
	blue := math.Abs(float64(B1>>8)-float64(B2>>8)) <= colorDiff
	return red && green && blue
}

func GetAllPolygons(img image.RGBA) [][]image.Point {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	visited := make([][]bool, width)

	for i := 0; i < width; i++ {
		visited[i] = make([]bool, height)
	}

	var polygons [][]image.Point

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			var polygon []image.Point
			start := img.RGBAAt(i, j)
			if visited[i][j] {
				continue
			}

			var queue Queue
			queue.push(image.Point{X: i, Y: j})

			for queue.len() > 0 {
				testing := queue.pop()

				x := testing.X
				y := testing.Y

				if x-1 >= 0 {
					testColor := img.RGBAAt(x-1, y)
					testPoint := image.Point{X: x - 1, Y: y}
					polygon = append(polygon, testPoint)
					if !visited[x-1][y] && similarColor(testColor, start) {
						visited[x-1][y] = true
						queue.push(testPoint)
					}
				}

				if x+1 < width {
					testColor := img.RGBAAt(x+1, y)
					testPoint := image.Point{X: x + 1, Y: y}
					polygon = append(polygon, testPoint)
					if !visited[x+1][y] && similarColor(testColor, start) {
						visited[x+1][y] = true
						queue.push(testPoint)
					}
				}

				if y-1 >= 0 {
					testColor := img.RGBAAt(x, y-1)
					testPoint := image.Point{X: x, Y: y - 1}
					polygon = append(polygon, testPoint)
					if !visited[x][y-1] && similarColor(testColor, start) {
						visited[x][y-1] = true
						queue.push(testPoint)
					}
				}
				if y+1 < height {
					testColor := img.RGBAAt(x, y+1)
					testPoint := image.Point{X: x, Y: y + 1}
					polygon = append(polygon, testPoint)
					if !visited[x][y+1] && similarColor(testColor, start) {
						visited[x][y+1] = true
						queue.push(testPoint)
					}
				}
			}
			polygons = append(polygons, polygon)
		}
	}

	// filter out
	log.Printf("Filter out start length: %v \n", len(polygons))
	var filtered [][]image.Point
	for idx, polygon := range polygons {
		if len(polygon) < 6 {
			continue
		}
		filtered = append(filtered, polygons[idx])
	}
	log.Printf("Filter out end length: %v \n", len(filtered))

	return filtered
}
