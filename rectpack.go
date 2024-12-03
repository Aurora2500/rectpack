// This package implements a rectangle packing algorithm
// based on the 2014 paper [Optimal Rectangle Packing: An Absolute Placement Approach](https://arxiv.org/abs/1402.0557).
package rectpack

import (
	"errors"
	"image"
	"slices"
)

var (
	ErrInfeasible = errors.New("rectpack: infeasible packing")
)

// Rectpack takes a list or rectangles to pack together,
// and returns the size of the rectangle needed to contain them
// and each individual rectangle within.
func Rectpack(rects []image.Point) (size image.Point, Positions []image.Point) {
	indices := make([]int, len(rects))
	for i := range rects {
		indices[i] = i
	}

	slices.SortFunc(indices, func(a int, b int) int {
		left := rects[a]
		right := rects[b]
		return left.X*left.Y - right.X*right.Y
	})

	return image.Point{}, nil
}

type rectangle struct {
	pos  image.Point
	size image.Point
}

func (r rectangle) area() int {
	return r.size.X * r.size.Y
}

func (r rectangle) width() int {
	return r.size.X
}

func (r rectangle) height() int {
	return r.size.Y
}

func (r rectangle) fits(s image.Point) bool {
	return s.X <= r.size.X && s.Y <= r.size.Y
}

func (r rectangle) leftover(s image.Point) int {
	area := s.X * s.Y
	return r.area() - area
}

type indexedPos struct {
	pos   image.Point
	index int
}
