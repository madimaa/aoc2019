package array2d

/*
This solution is slower than [][]type on less than ~50k entries
but significantly faster on 10m+ data
*/

//Array2D - 2 dimensional array
type Array2D struct {
	data []string
	x    int
	y    int
}

//Create - create 2d array
func Create(x, y int) *Array2D {
	return &Array2D{data: make([]string, x*y), x: x, y: y}
}

//Put - add element to the array
func (a2d *Array2D) Put(x, y int, value string) {
	a2d.data[x*a2d.y+y] = value
}

//Get - get value from x y index
func (a2d *Array2D) Get(x, y int) string {
	return a2d.data[x*a2d.y+y]
}

//GetSize - returns the size of X and size of Y
func (a2d *Array2D) GetSize() (int, int) {
	return a2d.x, a2d.y
}
