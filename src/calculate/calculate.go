package calculate

import "fmt"

type Values struct {
	Value1 int
	Value2 int
}

func (c *Values) SumValues(){
	fmt.Println(c.Value1 + c.Value2)
}

func main() {
	
}