// 练习 3.3： 根据高度给每个多边形上色，那样峰值部将是红色(#ff0000)，谷部将是蓝色
// (#0000ff)。
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 150                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: #ff0000; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	var max, min float64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			z := f(x, y)
			if max < z {
				max = z
			}
			if min > z {
				min = z
			}
		}
	}

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			z := f(x, y)
			color := jetGradient(z, max, min)

			if math.IsInf(ax, 0) || math.IsNaN(ax) || math.IsInf(bx, 0) || math.IsNaN(bx) ||
			math.IsInf(cx, 0) || math.IsNaN(cx) || math.IsInf(dx, 0) || math.IsNaN(dx) ||
			math.IsInf(ay, 0) || math.IsNaN(ay) || math.IsInf(by, 0) || math.IsNaN(by) ||
			math.IsInf(cy, 0) || math.IsNaN(cy) || math.IsInf(dy, 0) || math.IsNaN(dy) {
				continue
			}
			
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style=' fill:#%06x; stroke:#%06x;stroke-width:1'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func jetGradient(gray float64, max float64, min float64) uint32 {
	per := (gray - min) / (max - min)
	var r, g, b uint32
	if per <= 0.25 {
		r = 0
		g = (uint32)(4 * per * 255)
		b = 255
	} else if (per > 0.25) && (per <= 0.5) {
		r = 0
		g = 255
		b = (uint32)((1 - 4*(per-0.25)) * 255)
	} else if (per > 0.5) && (per <= 0.75) {
		r = (uint32)(4 * (per - 0.5) * 255)
		g = 255
		b = 0
	} else {
		r = 255
		g = (uint32)((1 - 4*(per-0.75)) * 255)
		b = 0
	}
	rgb := (r << 16) | (g << 8) | b
	return rgb
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
