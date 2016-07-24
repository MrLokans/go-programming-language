// SVG rendering of a 3-D surface function

package main

import (
	"fmt"
	"math"
)

const (
	width_px, height_px = 600, 320                // size of canvas in pixels
	cells               = 100                     // number of grid cells
	xy_range            = 30.0                    //axis range
	xy_scale            = width_px / 2 / xy_range // pixels per x or y unit
	z_scale             = height_px * 0.4         // pixels per Z unit
	angle               = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-with: 0.7' "+
		"width='%d' height='%d'>", width_px, height_px)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Printf("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x, y) at corner of cell (i, j)
	x := xy_range * (float64(i)/cells - 0.5)
	y := xy_range * (float64(j)/cells - 0.5)

	z := f(x, y)

	// project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy)
	sx := width_px/2 + (x-y)*cos30*xy_scale
	sy := height_px/2 + (x+y)*sin30*xy_scale - z*z_scale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r * math.Cos(r)
}
