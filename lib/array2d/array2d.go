package array2d

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
