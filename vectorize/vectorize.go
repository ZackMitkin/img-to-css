package vectorize

import (
	"fmt"
	svg "github.com/ajstarks/svgo/float"
	"image"
	"image/color"
	"log"
	"os"
)

func colorToRGB(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("rgb(%d, %d, %d)", r>>8, g>>8, b>>8)
}

func Vectorised(img image.RGBA, paths [][]image.Point, colors []string, output string) {
	f, _ := os.Create(output + ".svg")
	canvas := svg.New(f)
	canvas.Start(float64(img.Bounds().Dx()), float64(img.Bounds().Dy()))

	for idx, path := range paths {
		if len(path) == 0 {
			continue
		}
		pathStr := fmt.Sprintf("M%d %d,", path[0].X, path[0].Y)
		for idx := range path {
			x1 := path[idx].X
			y1 := path[idx].Y
			pathStr += fmt.Sprintf(" %d %d,", x1, y1)
		}
		x1 := path[0].X
		y1 := path[0].Y
		pathStr += fmt.Sprintf("%d %dZ", x1, y1)
		c := colors[idx]
		canvas.Path(pathStr, fmt.Sprintf("shape-rendering:auto; fill:%s; stroke:%s; stroke-width: 1", c, c))
	}
	canvas.End()
	err := f.Close()
	if err != nil {
		log.Println(err)
	}
}
