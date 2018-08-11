package golife

// Field represents a two-dimensional field of cells.
type Field struct {
	s    [][]bool
	w, h int
}

// NewField returns an empty field of the specified width and height.
func NewField(w, h int) *Field {
	s := make([][]bool, h)
	for i := range s {
		s[i] = make([]bool, w)
	}
	return &Field{s: s, w: w, h: h}
}

// Set sets the state of the specified cell to the given value.
func (f *Field) Set(x, y int, b bool) {
	f.s[x][y] = b
}

//Checks if cell is alive returns false if out of bounds
func (f *Field) Alive(x, y int) bool {
	if x >= 0 && x < len(f.s) && y >= 0 && y < len(f.s[0]) {
		return f.s[x][y]
	}
	return false
}

//returns the number of alive neighbors of a cell
func (f *Field) CountLiveNeighbors(x int, y int) int {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(j == 0 && i == 0) && f.Alive(x+i, y+j) {
				alive++
			}
		}
	}
	return alive
}
