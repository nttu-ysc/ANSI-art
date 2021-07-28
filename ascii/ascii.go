package ascii

import (
	"fmt"
	"image"
	"os"
	"strings"
)

var (
	wd, _ = os.Getwd()
	intensity, _ = ReadFloatLines(wd + "/rank/intensity.txt")
	rank, _ = ReadIntLines(wd + "/rank/rank.txt")
)

func Pixels2Ascii(img image.Image) string {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	sb := strings.Builder{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, _, _, _ := img.At(x, y).RGBA()  // r, g, b are the same for grayscale image
			r >>= 8
			asciiIdx := FindClosestK(int(r), intensity)
			c := rank[asciiIdx]
			//fmt.Println(r, c)
			_, err := fmt.Fprintf(&sb, "%c", c)
			if err != nil {
				return ""
			}
		}
		_, fprintln := fmt.Fprintln(&sb)
		if fprintln != nil {
			return ""
		}
	}

	return sb.String()
}
