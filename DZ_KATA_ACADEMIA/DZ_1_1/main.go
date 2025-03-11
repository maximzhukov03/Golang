package main

import (
	"fmt"
	"testing"
)

func fizzBuzz(a int) string{
	if a % 3 == 0 && a % 5 == 0 {
		return "buzz"
	}else if a % 3 == 0 {
		return "fizz"
	}else if a % 5 == 0{
		return "buzz"
	} else {
		return fmt.Sprintf("%d", a)
	}
}



func TestFizzBuzz(t *testing.T) {
    fizzBuzzTest := []struct{
        num int
        res string
    }{
        {1, "1"},   
		{3, "fizz"},
        {5, "buzz"},
        {15, "fizz buzz"},
    }
	for _, elem := range fizzBuzzTest {
        var result string = fizzBuzz(elem.num)
        if result != elem.res {
            t.Errorf("fizzBuzz(%d): %q != %q", elem.num, result, elem.res)
        }
    }
}

func main() {
	fmt.Println(fizzBuzz(1))	
	fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(5))
	fmt.Println(fizzBuzz(15))
}