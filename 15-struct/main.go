package main

import "fmt"

func main() {
	// 2. Penerapan struct
	var s1 student
	s1.name = "Jessi Kristin"
	s1.grade = 4

	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)

	// 3. inisialisasi object 
	var ss1 = student{}
	ss1.name = "aimyon"
	ss1.grade = 1

	var ss2 = student{"karin", 2}

	var ss3 = student{name: "ten"}

	fmt.Println("student 1 : ", ss1.name)
	fmt.Println("student 2 : ", ss2.name)
	fmt.Println("student 3 : ", ss3.name)

	// 4. variabel object pointer 
	var sss1 = student {name: "rian", grade: 3}

	var sss2 *student = &sss1
	fmt.Println("Student 1, name :", sss1.name)
	fmt.Println("Student 5, name :", sss2.name)

	sss2.name = "Hono"
	fmt.Println("Student 1, name :", sss1.name)
	fmt.Println("Student 5, name :", sss2.name)

	// 5. Embedded struct 
	var mhs1 = mhs{}
	mhs1.name = "andre"
	mhs1.age = 20
	mhs1.person.age = 21
	mhs1.grade = 2

	fmt.Println("name :", mhs1.name)
	fmt.Println("age mhs :", mhs1.age)
	fmt.Println("age person :", mhs1.person.age)
	fmt.Println("grade :", mhs1.grade)

	// 6. Pengisian Nilai Sub-struct 
	var p1 = person{name: "jess", age: 22}
	var mhss1 = mhs{person: p1, grade: 2}

	fmt.Println("name :", mhss1.name)
	fmt.Println("age  :", mhss1.person.age)
	fmt.Println("grade:", mhss1.grade)

	// 7. Anonymous struct 
	var mhs2 = struct{
		person
		grade int 
	}{}
	mhs2.person = person{"nicky", 23}
	mhs2.grade = 1

	fmt.Println("name :", mhs2.person.name)
	fmt.Println("age :", mhs2.person.age)
	fmt.Println("grade :", mhs2.grade)

	// 8. Kombinasi slice & struct 
	
	var allMhs = []person{
		{name: "Dian", age: 23},
		{name: "Fanny", age: 20},
		{name: "Bian", age: 22},
	}

	for _, mhs := range allMhs{
		fmt.Println(mhs.name, "age is", mhs.age)
	}

	// 9. deklarasi anonymous struct menggunakan keyword var
	var mahasiswa struct {
		person
		grade int 
	}
	mahasiswa.person = person{"tiwuk", 35}
	mahasiswa.grade = 1

	fmt.Println(mahasiswa.person.name)
	fmt.Println(mahasiswa.person.age)
	fmt.Println(mahasiswa.grade)

	// 10. nested struct untuk json
	// type mhsiswa struct{
	// 	orang struct{
	// 		name string
	// 		age int 
	// 	}
	// 	grade int 
	// 	hobbies []string
	// }

}

// 1. deklarasi struct
type student struct {
	name  string
	grade int
}

// 5. Embedded struct 
type person struct{
	name string
	age int
}

type mhs struct{
	grade int 
	age int
	person
}

