package main

import "fmt"

// func main() {
/*matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

for _, row := range matrix {
	for _, col := range row {
		fmt.Printf("%d ", col)
	}

	fmt.Println()
}*/

/*number := 10

switch {
case number > 1:
	fmt.Println("Number is larger than 1")
	fallthrough
case number < 11:
	fmt.Println("Number < 11")
} */

/* var number int64
fmt.Println("Add number")
fmt.Scan(&number)

switch {
case number >= 1:
	fmt.Println("Number is larger than 1")
case number <= 11:
	fmt.Println("Number < 11")
} */

/*a := 3
b := 10

if a != 2 && b > 5 || a > 6 {
	fmt.Println("YES!")
} */

/*
	const (
		Sunday = iota // iota увеличивает значения следующих переменных (работает только в одном блоке)
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		_ // пропускаем 7
		Add // здесь будет 8
	)

	fmt.Println(Sunday)   // вывод 0
	fmt.Println(Saturday) // вывод 6
	fmt.Println(Add) 	  // вывод 8
*/

// Замена громоздкого if на switch
/*if i == 0 {
	fmt.Println("Zero")
} else if i == 1 {
	fmt.Println("One")
} else if i == 2 {
	fmt.Println("Two")
} else if i == 3 {
	fmt.Println("Three")
} else if i == 4 {
	fmt.Println("Four")
} else if i == 5 {
	fmt.Println("Five")
}*/

/*var i int64
fmt.Println("What's your number?")
fmt.Scan(&i)

switch i {
case 0:
	fmt.Println("Zero")
case 1:
	fmt.Println("One")
case 2:
	fmt.Println("Two")
case 3:
	fmt.Println("Three")
case 4:
	fmt.Println("Four")
case 5:
	fmt.Println("Five")
default:
	fmt.Println("Unknown Number")
}*/

/*
	var a int
	fmt.Scan(&a)
	if a == 10000 {
		a = a / 10000
		fmt.Println(a)
	} else if a >= 1000 && a < 10000 {
		a = a / 1000
		fmt.Println(a)
	} else if a >= 100 && a < 1000 {
		a = a / 100
		fmt.Println(a)
	} else if a >= 10 && a < 100 {
		a = a / 10
		fmt.Println(a)
	} else if a < 10 {
		fmt.Println(a)
	}
*/

/*
	var s int
	fmt.Scan(&s)
	if s%400 == 0 || s%4 == 0 && s%100 != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}*/

/*
	var n, c, d int
	fmt.Scan(&n, &c, &d)
	for i := 1; i <= n; {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		} else {
			i++
		}

	}
*/

/*
	var n int
	for fmt.Scan(&n); n > 0; fmt.Scan(&n) {
		if n < 10 {
			continue
		} else if n > 100 {
			break
		} else {
			fmt.Println(n)
		}
	}
*/

/*
	var x, p, y int
	fmt.Scan(&x, &p, &y)
	for i := 0; i >= 0; {
		if x < y {
			x = x + x*p/100
			i += 1
		} else {
			fmt.Println(i)
			break
		}

	}
*/

/*
	func kek(x float64) float64 {
		var count float64
		z := 1
		for z; z < 11; z++ {
			count = z*z - x
			return count
		}
	}

	func main() {
		fmt.Println(kek(2))
	}
*/

// func main() {
// 	var message string
// 	message = SayHello("Paul", 24)
// 	printmessage(message)
// }

// func printmessage(message string) {
// 	fmt.Println(message)
// }

// func SayHello(name string, age int) string {
// 	return fmt.Sprintf("Hi, %s! You're %d y.o.", name, age)
// }

func main() {
	var arr [5]int
	arr2 := [5]int{}
	arrWithValues := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arrWithValues)

	arr[0] = 100
	arr2[0] = 100
	arrWithValues[0] = 100

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arrWithValues)
}