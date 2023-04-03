package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c, root1, root2, img, disc float64

	fmt.Print("Enter the a, b, c = ")
	fmt.Scanln(&a, &b, &c)

	disc = (b * b) - (4 * a * c)
	switch {
	case disc > 0:
		root1 = (-b + math.Sqrt(disc)/(2*a))
		root2 = (-b - math.Sqrt(disc)/(2*a))
		fmt.Println("Two Distinct Real Roots Exist: root1 = ", root1, " and root2 = ", root2)
	case disc == 0:
		root1 = -b / (2 * a)
		root2 = -b / (2 * a)
		fmt.Println("Two Equal and Real Roots Exist: root1 = ", root1, " and root2 = ", root2)
	case disc < 0:
		root1 = -b / (2 * a)
		root2 = -b / (2 * a)
		img = math.Sqrt(-disc) / (2 * a)
		fmt.Println("Two Distinct Complex Roots Exist: root1 = ", root1, "+", img, " and root2 = ", root2, "-", img)
	default:
		fmt.Println("Please enter Valid Numbers for a,b, and c")
	}

	//Check the roots are correct or not by putting the values of a,b,c in the Quadratic equation
	fmt.Println("Check the roots are correct or not by putting the values of a,b,c in the Quadratic equation")
	fmt.Println("root1 = ", root1)
	fmt.Println("root2 = ", root2)
	fmt.Println("root1*root1 = ", root1*root1)
	fmt.Println("root2*root2 = ", root2*root2)
	fmt.Println("root1*root2 = ", root1*root2)
	fmt.Println("a*root1*root1 + b*root1 + c = ", a*root1*root1+b*root1+c)
	fmt.Println("a*root2*root2 + b*root2 + c = ", a*root2*root2+b*root2+c)

}
