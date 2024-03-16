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

// Массивы

// func main() {
// 	var arr [5]int
// 	arr2 := [5]int{}
// 	arrWithValues := [5]int{1, 2, 3, 4, 5}

// 	fmt.Println(arr)
// 	fmt.Println(arr2)
// 	fmt.Println(arrWithValues)

// 	arr[0] = 100
// 	arr2[0] = 100
// 	arrWithValues[0] = 100

// 	fmt.Println(arr)
// 	fmt.Println(arr2)
// fmt.Println(arrWithValues)
// }

//// Слайсы - динамический массив
// В отличии от массивов -- это изменяемый объект
// Аллокация памяти - это когда ОС выделяет больше памяти для вашего приложения

// func main() {
// 	week := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
// 	monday := week[1]
// 	fmt.Println(len(week)) // 8
// 	fmt.Println(monday)    // 1
// 	fmt.Println(week[0])

// 	weekend := week[6:8]  // включает индексы 6 и 7
// 	weekend2 := week[6:]  // аналогично 2
// 	fmt.Println(weekend)  // 6, 7
// 	fmt.Println(weekend2) // 6, 7

// 	fmt.Println(len(weekend2)) // длина слайса (2)
// 	fmt.Println(cap(weekend2)) // вместимость слайса (2)

// 	// Примеры создания слайсов
// 	animals := []string{}
// 	animals2 := make([]string, 5, 10)
// 	animals = append(animals, "lion") // добавление эл-та в слайс
// 	fmt.Println(animals)
// 	fmt.Println(cap(animals2))

// 	a := []int{1}
// 	b := a[0:1]
// 	b[0] = 0
// 	fmt.Println(a) // [0]
// 	a[0] = 1
// 	fmt.Println(b) // [1]

// 	var arr [5]int
// 	_ = arr
// 	var sl []int
// 	sl = arr[1:2]

// 	fmt.Println("arr", arr) // [0 0 0 0 0]
// 	fmt.Println("sl", sl)   // [0]

// 	arr[1] = 100

// 	fmt.Println("arr", arr) // [0 100 0 0 0]
// 	fmt.Println("sl", sl)   // [100]
// }

//// Хеш-таблица| Map - ассоциативный массив (ссылочный тип)

// Одинаковый тип данных		Одинаковый тип данных
// 1							monday
// 2							tuesday
// 3							wednesday
// 4							thursday
// 5							friday
// 6							saturday
// 7							sunday

// func main() {
// var m map[string]int
// m2 := make(map[string]int)
// m3 := map[string]int{}
// m4 := map[string]int{
// 	"1": 1,
// 	"2": 2,
// }
// fmt.Println("m", m)
// fmt.Println("m2", m2)
// fmt.Println("m3", m3)
// fmt.Println("m4", m4)

// m := make(map[int]string)

// m[1] = "hello"
// m[2] = " "
// m[3] = "world"
// m[4] = "!"
// fmt.Println(m)
// m[4] = "?"
// fmt.Println(m[4])
// delete(m, 4)
// fmt.Println(m)

// k, ok := m[4]
// fmt.Println("k", k)   // key
// fmt.Println("ok", ok) // ok false

// workdays := map[int]string{
// 	1: "monday",
// 	2: "tuesday",
// 	3: "wednesday",
// 	4: "thursday",
// 	5: "friday",
// }
// fmt.Println(workdays[4]) // thursday

// workdays[6] = "saturday"
// fmt.Println(workdays[6]) // saturday
// fmt.Println(workdays[7]) // пустая строка

// day, ok := workdays[6] // ok = boolean

// if ok {
// 	fmt.Println("6 day is " + day) // 6 day is saturday
// }

// delete(workdays, 5) // удаляет ключ со значением (в данном примере удалилась пятница)
// fmt.Println(workdays)
// fmt.Println(len(workdays))
// }

//// Циклы - for each (range)

// func main() {
// 	arr := []string{"A", "B", "C"}
// 	// for index, letter := range arr {	// range - перебирает массив arr
// 	// 	// ...
// 	// }
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Println(arr[i])
// 	}
// }

//// Циклы - while (for condition)
// for condition {
// 	// .. do something
// }

//// Циклы - infinity loop
// for {
// 	fmt.Println("infinity loop")
// }

//// Циклы - break
// for i, n := range []int{1, 2, 3} {
// 	fmt.Println(i)
// 	fmt.Println(n)
// 	break
// }

//// Примеры использования циклов
// func main() {
// 	// arr := []string{"a", "b", "c"}
// 	// for index, letter := range arr {
// 	// 	fmt.Println(index, letter)
// 	// }
// 	m := map[string]int{
// 		"a": 1,
// 		"b": 2,
// 		"c": 3,
// 		"d": 4,
// 	}
// 	for k, v := range m { // не гарантирует совпадение с порядком значений внутри мапы
// 		fmt.Println(v)
// 		fmt.Println(k)
// 	}
// }

//// Ветвления - if и switch
// func main() {
//// IF
// age := // вводим возраст
// if age >= 18 {	// && - оператор И; || - оператор ИЛИ
// 	//pass
// } else {
// 	// you shall not pass
// }

//// SWITCH
// 	hairColor := "black"
// 	switch hairColor {
// 	case "black":
// 		fmt.Println("волосы черные - красим в белый")
// 	case "white":
// 		fmt.Println("волосы белые - красим в черный")
// 	case "green":
// 		fmt.Println("волосы зеленые - красим в красный")
// 	default:
// 		fmt.Println("мы не знаем цвет волос")
// 	}

// for i, n := range []int{1, 2, 3} {
// 	if i == 1 && n == 2 {
// 		fmt.Println("i == 1 && n == 2")
// 		break
// 	}
// }

// 	var found bool
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	for !found {
// 		for i := 0; i < len(arr); i++ {
// 			fmt.Println("work with", i+1)
// 			found = findOut(arr[i])
// 			if found { // по умолчанию это значение TRUE
// 				break
// 			}
// 		}
// 	}
// }

// func findOut(i int) bool {
// 	return i%2 == 0
// }

// func main() {
// 	knownFruit := false
// 	fruit := "blueberry"
// 	switch fruit {
// 	case "apple":
// 		fmt.Println("this is apple")
// 		knownFruit = true
// 	case "dragon fruit":
// 		fmt.Println("this is dragon fruit")
// 		knownFruit = true
// 	default:
// 		fmt.Println("we don't know this fruit")
// 	}

// 	if knownFruit { // по умолчанию значение TRUE
// 		fmt.Println("we know this fruit")
// 	}
// }

//// Указатели и nil
// & - амперсанд
// * - звездочка
// new(T) - это универсальная функция

// func main() {
// 	i, j := 42, 2701

// 	p := &i         // point to i
// 	fmt.Println(*p) // read i through the pointer
// 	*p = 21         // set i through the pointer
// 	fmt.Println(i)  // see the new value of i

// 	p = &j         // point to j
// 	*p = *p / 37   // divide j through the pointer
// 	fmt.Println(j) // see the new value of j
// }

// func main() {
// 	// 	phrase := "hello"
// 	// 	phrasePtr := &phrase    // тип данных *string
// 	// 	fmt.Println(phrasePtr)  // 0xc000088050
// 	// 	fmt.Println(*phrasePtr) // hello

// 	name := "Pasha"
// 	changeName(&name)
// 	fmt.Println(name)
// }

// func changeName(name *string) {
// 	*name = *name + "_new"
// }

//// Структуры - набор полей

// type User struct {
// 	Name, Email string
// 	IsConfirmed bool
// }

// func main() {
// 	user := newUser("Pasha", "email@example.com")
// 	fmt.Println(User.IsConfirmed)
// }

// func newUser(name string, email string) User {
// 	return User{
// 		Name:        name,
// 		Email:       email,
// 		IsConfirmed: false,
// 	}
// }

// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	v := Vertex{1, 2}
// 	fmt.Println(v.X) // 1
// 	v.X = 4
// 	fmt.Println(v.X) // 4
// }

// func main() {
// 	v := Vertex{1, 2}
// 	p := &v // pointer to struct
// 	p.X = 1e9
// 	fmt.Println(v)
// }

//// Встраивание структуры
/* Было
type Vehicle struct{
	Brand string
	Model string
	Latitude string
	Longitude string
} */

/* Стало
type Vehicle struct {
	Brand string
	Model string
	GPSPosition GeoPosition //можно встроить и без поля GPSPosition
}

type GeoPosition srtuct {
	Latitude string
	Longitude string
}
*/
// func main() {
// 	geoPos := GeoPosition{"58.1234", "22.1235"}
// 	myCar := Vehicle{"Lexus", "221", geoPos}
// }

//// Интерфейс - это набор методов, представляющих поведение для различных типов данных

//// Defer - выполняется до return
// 1 - выполнение перед выходом из внешней функции
// 2 - всегда выполнится
// 3 - несколько вызовов defer (разных функций)
// 4 - освобождение ресурсов (закрытие файла и т.д)
// 5 - если несколько дефер подряд, то last-in-first-out

// func main() {
// 	defer func() { // создали анонимную функцию
// 		fmt.Println("world")
// 	}() // () - вызов функции
// 	fmt.Println("hello")
// }

// func main() {
// var s, sep string
// for i := 1; i < len(os.Args); i++ {
// 	s += sep + os.Args[i]
// 	sep = " "
// }
// fmt.Println(s)

// task 1.1
// fmt.Println(strings.Join(os.Args[0:], " "))

// task 1.2
// s := ""
//
//	for i, arg := range os.Args[1:] {
//		s = fmt.Sprintf("%d %s", i, arg)
//		fmt.Println(s)
//	}
// func main() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// 	// NOTE: ignoring potential errors from input.Err()
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

//// Callback function
// func doSomething(callback func(int, int) int, s string) int {
// 	result := callback(5, 1)
// 	fmt.Println(s)
// 	return result
// }

// func main() {
// 	sumCallback := func(n1, n2 int) int {
// 		return n1 + n2
// 	}
// 	result := doSomething(sumCallback, "plus")
// 	fmt.Println(result)

// 	mulCallback := func(n1, n2 int) int {
// 		return n1 * n2
// 	}
// 	result = doSomething(mulCallback, "multiplies")
// 	fmt.Println(result)
// }

//// Замыкаиня
func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// fmt.Println(totalPrice(100)(10))
	orderPrice := totalPrice(10)
	fmt.Println(orderPrice(5))
}
