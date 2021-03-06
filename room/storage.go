package room

import (
	"math"

	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/globals"
	"github.com/Holmqvist1990/WARF2/resource"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

type Storage struct {
	Center       int
	StorageTiles []StorageTile
}

func NewStorage(mp *m.Map, x, y int) *Storage {
	s := &Storage{}
	tiles := mp.FloodFillRoom(x, y, m.RandomFloorBrick)
	if len(tiles) == 0 {
		return nil
	}
	s.StorageTiles = createStorageTiles(tiles, mp.Items)
	s.Center = determineCenter(mp, tiles)
	return s
}

// Use storage.
func (s *Storage) Use(dwarf *dwarf.Dwarf) {
	// Nothing yet.
}

func (s *Storage) GetAvailableTile(r resource.Resource) (idx int, ok bool) {
	for _, t := range s.StorageTiles {
		if t.Unavailable(r) {
			continue
		}
		return t.Idx, true
	}
	return -1, false
}

func (s *Storage) AddItem(idx int, amount uint, r resource.Resource) (int, bool) {
	tIdx, ok := s.getStorageTileIdxFromWorldIdx(idx)
	if !ok {
		return -1, false
	}
	t := s.StorageTiles[tIdx]
	if t.Remaining() > 0 {
		s.StorageTiles[tIdx].Add(r, amount)
		return t.Idx, true
	}
	/////////////////////////////
	// TODO
	// This just picks a random
	// available storetile and
	// drops it there.
	// Perhaps it would be cooler
	// if the dwarf moves to the
	// new destination.
	/////////////////////////////
	for _, st := range s.StorageTiles {
		if st.Unavailable(r) {
			continue
		}
		st.Add(r, amount)
		return st.Idx, true
	}
	return -1, false
}

func determineCenter(mp *m.Map, tiles m.Tiles) int {
	minx, maxx := math.MaxInt16, -1
	miny, maxy := math.MaxInt16, -1
	for _, t := range tiles {
		if t.X < minx {
			minx = t.X
		}
		if t.X > maxx {
			maxx = t.X
		}
		if t.Y < miny {
			miny = t.Y
		}
		if t.Y > maxy {
			maxy = t.Y
		}
	}
	midx := minx + ((maxx - minx) / 2)
	midy := miny + ((maxy - miny) / 2)
	center := globals.XYToIdx(midx, midy)
	for m.IsAnyWall(mp.Tiles[center].Sprite) {
		center++
	}
	return center
}

func (s *Storage) getStorageTileIdxFromWorldIdx(idx int) (int, bool) {
	for worldIndex, t := range s.StorageTiles {
		if t.Idx == idx {
			return worldIndex, true
		}
	}
	return -1, false
}
