package item

import (
	"math/rand"

	"github.com/Holmqvist1990/WARF2/globals"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

var (
	bookShelves = []int{
		BookShelfOne, BookShelfTwo, BookShelfThree,
		BookShelfFour, BookShelfFive, BookShelfSix,
		BookShelfSeven, BookShelfEight, BookShelfNine,
		BookShelfTen,
	}

	furniture = []int{
		ChairLeft, Table, ChairRight,
	}
)

func Place(mp *m.Map, x, y, sprite int) {
	tile, ok := mp.GetTile(x, y)
	if !ok {
		return
	}
	if m.IsAnyWall(tile.Sprite) {
		return
	}
	item, ok := mp.GetItemTile(x, y)
	if !ok {
		return
	}
	item.Sprite = sprite
	if IsBlocking(item.Sprite) {
		tile.Blocked = true
	}
}

func PlaceRandom(mp *m.Map, x, y int, f func() int) {
	idx := globals.XYToIdx(x, y)
	PlaceRandomIdx(mp, idx, f)
}

func PlaceRandomIdx(mp *m.Map, idx int, f func() int) {
	tile, ok := mp.GetTileByIndex(idx)
	if !ok {
		return
	}
	if m.IsAnyWall(tile.Sprite) {
		return
	}
	item, ok := mp.GetItemTileByIndex(idx)
	if !ok {
		return
	}
	item.Sprite = f()
	if IsBlocking(item.Sprite) {
		tile.Blocked = true
	}
}

func RandomBookshelf() int {
	return bookShelves[rand.Intn(len(bookShelves))]
}

func RandomFurniture() int {
	return furniture[rand.Intn(len(furniture))]
}

func RandomCrumbledWall() int {
	return rand.Intn(WallCrumbled4-WallCrumbled1+1) + WallCrumbled1
}
