package main

const Width, Height = 640, 480

type CenterInf interface {
	Center()
}

type Cursor struct {
	X, Y int
}

func (c *Cursor) Center() {
	c.X += Width / 2
	c.Y += Height / 2
}

func main() {
	c := new(Cursor)
	c.Center()
}

// func main() {
// 	c := new(Cursor)
// 	CenterInf(c).Center()
// }
