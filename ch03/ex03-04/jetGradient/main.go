// 练习 3.4： 参考1.7节Lissajous例子的函数，构造一个web服务器，用于计算函数曲面然后返
// 回SVG数据给客户端。服务器必须设置Content-Type头部：
// w.Header().Set("Content-Type", "image/svg+xml")
// （这一步在Lissajous例子中不是必须的，因为服务器使用标准的PNG图像格式，可以根据前
// 面的512个字节自动输出对应的头部。）允许客户端通过HTTP请求参数设置高度、宽度和颜
// 色等参数。
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type canvasPara struct {
	width   int // canvas size in pixels
	height  int
	cells   int     // number of grid cells
	xyrange float64 // axis ranges (-xyrange..+xyrange)
	xyscale int     // pixels per x or y unit
	zscale  int     // pixels per z unit
	angle   float64 // angle of x, y axes (=30°)
	color string
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", cornerHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func cornerHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "image/svg+xml")
	var para canvasPara
	var err error
	para.width, err = strconv.Atoi(r.FormValue("width"))
	if err != nil {
		para.width = 600
	}
	para.height, err = strconv.Atoi(r.FormValue("height"))
	if err != nil {
		para.height = 320
	}
	para.cells, err = strconv.Atoi(r.FormValue("cells"))
	if err != nil {
		para.cells = 150
	}
	para.xyrange, err = strconv.ParseFloat(r.FormValue("xyrange"), 64)
	if err != nil {
		para.xyrange = 30.0
	}
	para.xyscale, err = strconv.Atoi(r.FormValue("xyscale"))
	if err != nil {
		para.xyscale = (int)((float64)(para.width) / 2 / para.xyrange)
	}
	para.zscale, err = strconv.Atoi(r.FormValue("zscale"))
	if err != nil {
		para.zscale = (int)((float64)(para.height) * 0.4)
	}
	para.angle, err = strconv.ParseFloat(r.FormValue("angle"), 64)
	if err != nil {
		para.angle = math.Pi / 6
	}
	drawCorner(w, para)
}

func drawCorner(w io.Writer, para canvasPara) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", para.width, para.height)
	var max, min float64
	for i := 0; i < para.cells; i++ {
		for j := 0; j < para.cells; j++ {
			x := para.xyrange * (float64(i)/float64(para.cells) - 0.5)
			y := para.xyrange * (float64(j)/float64(para.cells) - 0.5)
			z := f(x, y)
			if max < z {
				max = z
			}
			if min > z {
				min = z
			}
		}
	}
	for i := 0; i < para.cells; i++ {
		for j := 0; j < para.cells; j++ {
			ax, ay := corner(i+1, j, para)
			bx, by := corner(i, j, para)
			cx, cy := corner(i, j+1, para)
			dx, dy := corner(i+1, j+1, para)

			x := para.xyrange * (float64(i)/float64(para.cells) - 0.5)
			y := para.xyrange * (float64(j)/float64(para.cells) - 0.5)
			z := f(x, y)
			color := jetGradient(z, max, min)

			if math.IsInf(ax, 0) || math.IsNaN(ax) || math.IsInf(bx, 0) || math.IsNaN(bx) ||
				math.IsInf(cx, 0) || math.IsNaN(cx) || math.IsInf(dx, 0) || math.IsNaN(dx) ||
				math.IsInf(ay, 0) || math.IsNaN(ay) || math.IsInf(by, 0) || math.IsNaN(by) ||
				math.IsInf(cy, 0) || math.IsNaN(cy) || math.IsInf(dy, 0) || math.IsNaN(dy) {
				continue
			}
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style=' fill:#%06x; stroke:#%06x;stroke-width:1'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color, color)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i int, j int, para canvasPara) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := para.xyrange * (float64(i)/float64(para.cells) - 0.5)
	y := para.xyrange * (float64(j)/float64(para.cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(para.width/2) + (x-y)*math.Cos(para.angle)*float64(para.xyscale)
	sy := float64(para.height/2) + (x+y)*math.Sin(para.angle)*float64(para.xyscale) - z*float64(para.zscale)
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
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
