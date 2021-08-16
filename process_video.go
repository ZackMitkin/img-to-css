package main

import (
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
	"image"
	"image/draw"
	"img_to_css/bfs"
	"img_to_css/css"
	"log"
	"os"
	"sync"
)

func processVideo(path string) {
	fileNames, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}
	polygonPaths := make([][][]image.Point, len(fileNames))
	for idx, fileName := range fileNames {
		fileName := fmt.Sprintf("%v/%v", path, fileName.Name())
		img, err := imgio.Open(fileName)
		if err != nil {
			log.Fatalln(err)
		}

		m := image.NewRGBA(img.Bounds())
		draw.Draw(m, m.Bounds(), img, img.Bounds().Min, draw.Src)

		polygons := bfs.GetAllPolygons(*m)

		var wg sync.WaitGroup
		paths := make([][]image.Point, len(polygons))
		colors := make([]string, len(polygons))
		var sem = make(chan int, 24)
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
		polygonPaths[idx] = paths
		//vectorize.Vectorised(*m, paths, "svg")
		css.WritePolygons(*m, paths, colors)
	}
}
