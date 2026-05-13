package main

import "fmt"

func main() {
	studentGrades := map[string]int{
		"Mario":     30,
		"Mariarosa": 30,
		"Gino":      18,
	}
	fmt.Printf("Grades: %+v\n", studentGrades)

	studentGrades["Gino"] = 20
	fmt.Printf("Grades: %+v\n", studentGrades)

	mario, ok := studentGrades["Mario"]
	if ok {
		fmt.Printf("Mario: %d\n", mario)
	}

	_, ok = studentGrades["Pino"]
	if ok {
		fmt.Println("Pino was here")
	}

	delete(studentGrades, "Gino")
	fmt.Printf("Grades: %+v\n", studentGrades)

	configs := make(map[string]int)
	fmt.Printf("%+v, %T\n", configs, configs)

	if configs == nil {
		fmt.Println("configs is nil")
	}

	for k, v := range studentGrades {
		fmt.Println(k, v)
	}
}
