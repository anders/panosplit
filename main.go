//
// Copyright 2017 Anders Bergh <anders1@gmail.com>
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF OR
// IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	prefix := flag.String("prefix", "pano_", "output prefix")
	quality := flag.Int("quality", jpeg.DefaultQuality, "JPEG output quality 1-100")
	flag.Parse()

	path := flag.Arg(0)
	if path == "" {
		fmt.Fprintln(os.Stderr, "need input image argument")
		flag.Usage()
		os.Exit(1)
	}

	if *quality < 1 || *quality > 100 {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	bounds := src.Bounds()
	size := bounds.Size()

	if size.X < size.Y*2 {
		log.Fatal("input image not wide enough")
	}

	for i, x := 0, 0; i < size.X/size.Y; i++ {
		m := image.NewRGBA(image.Rect(0, 0, size.Y, size.Y))

		// copy from input
		draw.Draw(m, m.Bounds(), src, image.Point{X: x}, draw.Src)

		x += size.Y

		fname := fmt.Sprintf("%s%d.jpg", *prefix, i+1)
		log.Printf("writing %s", fname)

		f, err := os.Create(fname)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if err := jpeg.Encode(f, m, &jpeg.Options{Quality: *quality}); err != nil {
			log.Fatal(err)
		}
	}
}
