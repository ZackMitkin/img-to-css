package trace

import (
	"image"
)

func mooreNeighborhood(point image.Point, width int, height int) [8]*image.Point {
	var P1 = image.Point{}
	var P2 = image.Point{}
	var P3 = image.Point{}
	var P4 = image.Point{}
	var P5 = image.Point{}
	var P6 = image.Point{}
	var P7 = image.Point{}
	var P8 = image.Point{}

	x := point.X
	y := point.Y

	if point.X > 0 {
		P8 = image.Point{X: x - 1, Y: y}
		if y < height {
			P7 = image.Point{X: x - 1, Y: y + 1}
		}
		if y > 0 {
			P1 = image.Point{X: x - 1, Y: y - 1}
		}
	}
	if x < width {
		P4 = image.Point{X: x + 1, Y: y}
		if y < height {
			P5 = image.Point{X: x + 1, Y: y + 1}
		}
		if y > 0 {
			P3 = image.Point{X: x + 1, Y: y - 1}
		}
	}

	if y > 0 {
		P2 = image.Point{X: x, Y: y - 1}
	}
	if y < height {
		P6 = image.Point{X: x, Y: y + 1}
	}
	neighborhood := [8]*image.Point{
		&P1, // P1
		&P2, // P2
		&P3, // P3
		&P4, // P4
		&P5, // P5
		&P6, // P6
		&P7, // P7
		&P8, // P8
	}
	return neighborhood
}

func GetPolygonPath(img image.RGBA, polygon []image.Point) []image.Point {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	pixels := make([][]bool, width+1)
	visited := make([][]bool, width+1)

	var path []image.Point

	for i := 0; i <= width; i++ {
		visited[i] = make([]bool, height+1)
		pixels[i] = make([]bool, height+1)
	}

	// set fake pixels
	for _, p := range polygon {
		pixels[p.X][p.Y] = true
	}

	// find black pixel (start)
	var start image.Point
findStart:
	for x := range pixels {
		for y := range pixels[x] {
			if visited[x][y] || !pixels[x][y] {
				continue
			}
			start = image.Point{X: x, Y: y}
			break findStart
		}
	}

	path = append(path, start)
	p := start
	index := 6
	c := *mooreNeighborhood(p, width, height)[index]
	var lastPixel image.Point
	if start.X < len(pixels)-1 {
		lastPixel = image.Point{X: start.X + 1, Y: start.Y}
	} else {
		lastPixel = image.Point{X: start.X - 1, Y: start.Y}
	}

	iteration := 0
	startIndex := 0

	for !c.Eq(start) {
		if !visited[c.X][c.Y] && pixels[c.X][c.Y] {
			visited[c.X][c.Y] = true
			path = append(path, c)
			p = c
			c = lastPixel
			if index == 1 {
				index = 7
			} else if index == 0 {
				index = 6
			} else {
				index -= 2
			}
			startIndex = 0 // reset start
			continue
		}
		visited[lastPixel.X][lastPixel.Y] = true
		lastPixel = c
		c = *mooreNeighborhood(p, width, height)[index]

		index++
		startIndex++

		if index == 8 {
			index = 0
		}

		iteration++
		if startIndex >= 9 {
			break
		}
	}
	return path
}
