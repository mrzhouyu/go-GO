package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteindex = 0 //first color in palette
	blackindex = 1 //next color
)

func main() {
	plot(os.Stdout) //调用画图函数
}

func plot(out io.Writer) {
	const (
		cycle  = 5
		res    = 0.001
		size   = 100
		nframe = 64
		delay  = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframe}
	phase := 0.0
	for i := 0; i < nframe; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1) //生成图片大小设定
		img := image.NewPaletted(rect, palette)      //设置两种颜色
		for t := 0.0; t < cycle*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)                                             //设置坐标
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackindex) //填入坐标画图

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}
