package rectpack

import "image"

type Size = image.Point
type Rect = image.Rectangle

// Rectpack takes a list or rectangles to pack together,
// and returns the size of the rectangle needed to contain them
// and each individual rectangle within.
func Rectpack(rects []Size) (size Size, rectangles []Rect) {
	return image.Point{}, nil
}
