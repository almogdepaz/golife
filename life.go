package golife

import "math/rand"

// Life stores the state of a round of Conway's Game of Life.
type Life struct {
	a, b *Field
	w, h int
	next func(f *Field, x, y int) bool
}

// Creates a new life from a predefined state (seed).
func NewLife(w, h int, seed []bool, next func(f *Field, x, y int) bool) *Life {
	a := NewField(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			a.Set(i, j, seed[w*i+j])
		}
	}
	return &Life{a: a, b: NewField(w, h), w: w, h: h, next: next,}
}

// NewLifeRandSeed returns a new Life game state with a random initial state.
func NewLifeRandSeed(w, h int, next func(f *Field, x, y int) bool) *Life {
	a := NewField(w, h)
	for i := 0; i < (w * h / 4); i++ {
		a.Set(rand.Intn(w), rand.Intn(h), true)
	}
	return &Life{
		a: a, b: NewField(w, h),
		w: w, h: h, next: next,
	}
}

//Sets the inner logic of updating a cell state
func (l *Life) SetNextFunc(n func(f *Field, x, y int) bool) {
	l.next = n
}

// Step advances the game by one instant, recomputing and updating all cells.
func (l *Life) Step() {
	// Update the state of the next field (b) from the current field (a).
	for x := 0; x < l.h; x++ {
		for y := 0; y < l.w; y++ {
			l.b.Set(x, y, l.next(l.a, x, y))
		}
	}
	// Swap fields a and b.
	l.a, l.b = l.b, l.a
}

func (l *Life) String() string {
	var str string
	for x := 0; x < l.h; x++ {
		for y := 0; y < l.w; y++ {
			if l.a.Alive(x, y) {
				str += "1 "
			} else {
				str += "0 "
			}
		}
	}
	return str[:len(str)-1]
}

//this is the logic for next generation cell state
func Next(f *Field, x, y int) bool {
	alive := f.CountLiveNeighbors(x, y)
	// Any live cell with fewer than two live neighbors dies .
	// Any live cell with two or three live neighbors lives on to the next generation.
	// Any live cell with more than three live neighbors dies .
	// Any dead cell with exactly three live neighbors becomes a live cell .
	return alive == 3 || alive == 2 && f.Alive(x, y)
}

//this is the logic for next generation cell state in the plague version
func NextV2(f *Field, x, y int) bool {
	//if cell is alive and ha no living horizontal neighbors return false
	if f.Alive(x, y) && !f.Alive(x+1, y) && !f.Alive(x-1, y) && !f.Alive(x, y+1) && !f.Alive(x, y-1) {
		return false
	}
	//if cell is dead but has exactly one living neighbor return true
	if !f.Alive(x, y) && f.CountLiveNeighbors(x, y) == 1 {
		return true
	}
	//state remains the same
	return f.Alive(x, y)
}
