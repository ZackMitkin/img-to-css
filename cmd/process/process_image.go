package process

import (
	"image"
	"image/draw"
	"img_to_css/bfs"
	"img_to_css/vectorize/css"
	"sync"
)

func Image(img image.Image, colorDiff float64, minLineLength int) string {
	m := image.NewRGBA(img.Bounds())
	draw.Draw(m, m.Bounds(), img, img.Bounds().Min, draw.Src)

	polygons := bfs.GetAllPolygons(*m, colorDiff, minLineLength)
	var wg sync.WaitGroup
	paths := make([][]image.Point, len(polygons))
	colors := make([]string, len(polygons))
	var sem = make(chan int, 1)
	for idx, polygon := range polygons {
		wg.Add(1)
		sem <- 1
		go func(idx int, polygon []image.Point, img image.RGBA, wg *sync.WaitGroup) {
			path, color := getPath(polygon, img)
			paths[idx] = path
			colors[idx] = color
			wg.Done()
			<-sem
		}(idx, polygon, *m, &wg)
	}
	wg.Wait()
	html := css.WritePolygons(*m, paths, colors)
	return html
}
