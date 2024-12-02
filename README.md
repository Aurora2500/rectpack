# Rectpack

This package implements a rectangle packing algorithm
based on the 2014 paper [Optimal Rectangle Packing: An Absolute Placement Approach](https://arxiv.org/abs/1402.0557).

The package uses the stdlib [`image`](https://pkg.go.dev/image) [`Point`](https://pkg.go.dev/image#Point) and [`Rectangle`](https://pkg.go.dev/image#Rectangle)
types to specify the sizes and positions to make it more directly plug & play with use cases developers might have.