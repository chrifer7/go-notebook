package examples

import "fmt"

const (
	plus6    = iota + 6
	bitshift = 2 << iota
	bitshift2

	first  = iota
	second = iota
)

const (
	third = iota
	fourth
)

// Collections examples
// Capital letter to exporting function
func Collections() {
	//creatingArrays()

	//slices1()
	//slices2()

	//maps()

	structs()
}

func structs() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u user
	fmt.Println(u)

	u.ID = 1
	u.FirstName = "Kris"
	u.LastName = "F."
	fmt.Println(u)

	u2 := user{
		ID:        2,
		FirstName: "Kayden",
		LastName:  "F.",
	}
	fmt.Println(u2)
}

func maps() {
	m := map[string]int{"foo": 42}

	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 77
	fmt.Println(m["foo"])

	delete(m, "foo")
	fmt.Println(m)

}

func slices2() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 4, 42, 27)
	fmt.Println(slice)

	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]
	fmt.Println(s2, s3, s4)
}

func slices1() {
	arr := [3]int{1, 2, 3}

	slice := arr[:]

	arr[1] = 33
	slice[2] = 77

	fmt.Println(arr, slice)

}

func creatingArrays() {
	var arr [3]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	fmt.Println(arr[1])

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

}

// Primitive data examples
// Capital letter to exporting function
func PrimitiveData() {
	fmt.Println("Working with variables")
	variables()

	fmt.Println("Working with pointers")

	pointers1()

	pointers2()

	fmt.Println("Working with constants")

	constants()

	fmt.Println(plus6)
	fmt.Println(bitshift)
	fmt.Println(bitshift2)

	fmt.Println(first, second)
	fmt.Println(third, fourth)
}

func constants() {
	const pi = 3.1415
	fmt.Println(pi)

	const c int = 7
	const c2 = 3
	fmt.Println(float32(c) + 1.2)
	fmt.Println(c2 + 1.2) //dynamic adapt
}

func pointers2() {
	firstName := "Kelly"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Kristen"
	fmt.Println(ptr, *ptr)
}

func pointers1() {
	var firstName *string = new(string)

	fmt.Println(*firstName)
}

func variables() {
	var i int
	i = 42
	fmt.Printf("Hello again! %d\n", i)

	var f float32 = 3.1416
	fmt.Println(f)

	firstName := "Kelly"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
