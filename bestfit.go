package rectpack

import (
	"image"
	"math"
)

func BestFitContainer(size image.Point, rects []image.Point) (pos []image.Point, err error) {
	var freeSpaces []rectangle = []rectangle{{size: size}}
	var placements []indexedPos

	for i, rect := range rects {
		var bestSpace rectangle
		var bestIdx int
		bestFit := math.MaxInt
		for j, space := range freeSpaces {
			if space.fits(rect) {
				leftover := space.leftover(rect)
				if leftover < bestFit {
					bestFit = leftover
					bestSpace = space
					bestIdx = j
				}
			}
		}
		if bestFit != math.MaxInt {
			placements = append(placements, indexedPos{pos: bestSpace.pos, index: i})
			quickRemove(&freeSpaces, bestIdx)
		} else {
			return nil, ErrInfeasible
		}
	}
	orderedPlacements := make([]image.Point, len(placements))
	for _, p := range placements {
		orderedPlacements[p.index] = p.pos
	}
	return orderedPlacements, nil
}

func BestFit(rects []image.Point) (size image.Point, pos []image.Point) {
	return image.Point{}, nil
}
