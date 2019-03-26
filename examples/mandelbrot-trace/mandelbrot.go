// mandelbrot example code adapted from Francesc Campoy's mandelbrot package.
// https://github.com/campoy/mandelbrot
package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"sync"
)

// tag::mandelbrot[]

import "github.com/pkg/profile"

func main() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	// end::mandelbrot[]

	var (
		height  = flag.Int("h", 1024, "height of the output image in pixels")
		width   = flag.Int("w", 1024, "width of the output image in pixels")
		mode    = flag.String("mode", "seq", "mode: seq, px, row, workers")
		workers = flag.Int("workers", 1, "number of workers to use")
	)
	flag.Parse()

	const output = "mandelbrot.png"

	// open a new file
	f, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	// create the image
	c := make([][]color.RGBA, *height)
	for i := range c {
		c[i] = make([]color.RGBA, *width)
	}

	img := &img{
		h: *height,
		w: *width,
		m: c,
	}

	switch *mode {
	case "seq":
		seqFillImg(img)
	case "px":
		oneToOneFillImg(img)
	case "row":
		onePerRowFillImg(img)
	case "workers":
		nWorkersFillImg(img, *workers)
	default:
		panic("unknown mode")
	}

	// and encoding it
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

type img struct {
	h, w int
	m    [][]color.RGBA
}

func (m *img) At(x, y int) color.Color { return m.m[x][y] }
func (m *img) ColorModel() color.Model { return color.RGBAModel }
func (m *img) Bounds() image.Rectangle { return image.Rect(0, 0, m.h, m.w) }

// tag::seqfillimg[]
func seqFillImg(m *img) {
	for i, row := range m.m {
		for j := range row {
			fillPixel(m, i, j)
		}
	}
}

// end::seqfillimg[]

func oneToOneFillImg(m *img) {
	var wg sync.WaitGroup
	wg.Add(m.h * m.w)
	for i, row := range m.m {
		for j := range row {
			go func(i, j int) {
				fillPixel(m, i, j)
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()
}

func onePerRowFillImg(m *img) {
	var wg sync.WaitGroup
	wg.Add(m.h)
	for i := range m.m {
		go func(i int) {
			for j := range m.m[i] {
				fillPixel(m, i, j)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func nWorkersFillImg(m *img, workers int) {
	c := make(chan int, 1024)
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			for i := range c {
				for j := range m.m[i] {
					fillPixel(m, i, j)
				}
			}
			wg.Done()
		}()
	}

	for i := range m.m {
		c <- i
	}
	close(c)
	wg.Wait()
}

func fillPixel(m *img, x, y int) {
	const n = 1000
	const Limit = 2.0
	Zr, Zi, Tr, Ti := 0.0, 0.0, 0.0, 0.0
	Cr := (2*float64(x)/float64(n) - 1.5)
	Ci := (2*float64(y)/float64(n) - 1.0)

	for i := 0; i < n && (Tr+Ti <= Limit*Limit); i++ {
		Zi = 2*Zr*Zi + Ci
		Zr = Tr - Ti + Cr
		Tr = Zr * Zr
		Ti = Zi * Zi
	}
	paint(&m.m[x][y], Tr, Ti)
}

func paint(c *color.RGBA, x, y float64) {
	n := byte(x * y * 2)
	c.R, c.G, c.B, c.A = n, n, n, 255
}
