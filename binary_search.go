package main
import "fmt"

func main(){

	array := []int{0,1,2,3,4,5,6,7,8,9,10}
	fmt.Println("Sorted array:", array)

	indexInArray, found, iterations := binSearch( -6 , array )
	fmt.Println( "iterations:", iterations)
	fmt.Println( "found:", found)
	fmt.Println( "indexInArray:",indexInArray)

}

func binSearch( item int , array []int ) ( int , bool , int ) {

	leftPosition := 0
	rightPosition := len(array)-1
	iterations := 0

	for leftPosition <= rightPosition {
		iterations++

		middlePosition := (rightPosition+leftPosition)/2
		middleItem := array[middlePosition]

		if middleItem == item {
			return middlePosition, true, iterations
		}

		if item < middleItem {
			rightPosition = middlePosition - 1
		} else {
			leftPosition = middlePosition + 1
		}

	}

	return -1, false, iterations

}