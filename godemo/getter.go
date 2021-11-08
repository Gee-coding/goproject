package godemo

import "fmt"

type Car struct {
	brand    string
	color    string
	distance int64
}

func (c *Car) Drive(dis int64) {
	c.distance += dis
}

//Getter----------------

func (c *Car) GetDistance() int64 {
	return c.distance
}

//Setter----------------

func (c *Car) SetColor(color string) {
	c.color = color
}

//Getter default value
func (c *Car) GetColor() string {
	if c.color == "" {
		c.color = "white"
	}
	return c.color
}

func main() {

	supercar := &Car{
		brand:    "Ferrary",
		distance: 0,
	}
	supercar.Drive(10)
	fmt.Println("from GetDistance(with out export)->", supercar.GetDistance())

	fmt.Println(*supercar)

	fmt.Println("Default color no Getter", supercar.color)     //<--ว่าง
	fmt.Println("Default color + Getter", supercar.GetColor()) //ได้ค่า default
	supercar.SetColor("red")
	fmt.Println("Default color after Set", supercar.GetColor()) //ได้ค่า setter
}
