package util

//Array2D - 2 dimensional array
type Array2D struct {
	data []string
	x    int
	y    int
}

//CreateArray2D - create 2d array
func CreateArray2D(x, y int) *Array2D {
	return &Array2D{data: make([]string, x*y), x: x, y: y}
}

//PutArray2D - add element to the array
func (a2d *Array2D) PutArray2D(x, y int, value string) {
	a2d.data[x*a2d.y+y] = value
}

//GetArray2D - get value from x y index
func (a2d *Array2D) GetArray2D(x, y int) string {
	return a2d.data[x*a2d.y+y]
}
