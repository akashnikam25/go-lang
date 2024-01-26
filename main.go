package higherorderfunc

import "fmt"

func areaOfCircle(r float64) float64 {
	return r * r * 3.14
}

func perimeter(r float64) float64 {
	return  2 * 3.14 * r
}
func printResult(radius float64, calFunc func(r float64)float64){

	res := calFunc(radius)
	fmt.Println(res)
	fmt.Println("Thank you")

}
func main() {
	fmt.Println("Akash")

	res := map[int]func( float64)float64{
		1:areaOfCircle,
		2:perimeter,
	}

	printResult(10,res[1])

  
}

//TODO: Convert this program into switch case for better understanding.


