package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Name string
	Subjects map[string]float64
	Average float64
}


func main() {
	
	welcome()

	user_name, subjects := acceptNameAndSubjects()
	average := get_average_grade(subjects)

	// build the stud struct 
	stud := Student{user_name, subjects, average}

	printFormattedGrades(&stud)

}	

func get_average_grade(subjects map[string]float64) float64 {

	sum := 0.0
	for _, grade := range subjects{
		sum += grade
	}

	average := sum / float64(len(subjects))

	return average
}


func get_letter_grade(grade float64) string {
	switch {
	case grade >= 90:
		return "A"
	case grade >= 80:
		return "B"
	case grade >= 70:
		return "C"
	case grade >= 60:
		return "D"
	default:
		return "F"
	}
}

func printFormattedGrades(stud *Student) {
	fmt.Println("\n**********************************************************************")
	fmt.Printf("|                      %s's Grade Report                      |\n", stud.Name)
	fmt.Println("**********************************************************************")
	fmt.Printf("| %-30s | %-15s | %-15s |\n", "Subject", "Grade", "Letter Grade")
	fmt.Println("----------------------------------------------------------------------")
	for subject_name, subject_grade := range stud.Subjects {
		fmt.Printf("| %-30s | %15.2f | %-15s |\n", subject_name, subject_grade, get_letter_grade(subject_grade))
	}
	fmt.Println("----------------------------------------------------------------------")
	average_grade := stud.Average
	fmt.Printf("| %-30s | %15.2f | %-15s |\n", "Average Grade", average_grade, get_letter_grade(average_grade))
	fmt.Println("**********************************************************************")
}

func welcome() {
	fmt.Println("************************************************")
	fmt.Println("|                                              |")
	fmt.Println("|                Welcome to                    |")
	fmt.Println("|         Student Grade Calculator!            |")
	fmt.Println("|                                              |")
	fmt.Println("************************************************")
	fmt.Println()
}

func acceptNameAndSubjects() (string, map[string]float64) {
	
	reader := bufio.NewReader(os.Stdin)

	var user_name string

	for user_name == "" {
	fmt.Print("Please Enter your name : ")
	user_name, _ = reader.ReadString('\n')
	user_name = strings.TrimSpace(user_name)

	}

	var number_of_subjects int
	var input string

	for {
		fmt.Print("Please Enter the number of subjects : ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		_, err := fmt.Sscanf(input, "%d", &number_of_subjects)

		if err == nil && number_of_subjects > 0 {
			break
		}

		fmt.Println("\nInvalid input!")
	}


	subjects := map[string]float64{}

	for i := 0; i < number_of_subjects; i++ {
		var subject_name string
		var subject_grade float64

		for subject_name == "" {
			fmt.Print("\nEnter Subject name : ")
			subject_name, _ = reader.ReadString('\n')
			subject_name = strings.TrimSpace(subject_name)
		}

		var input string
		for {
			fmt.Printf("Enter your grade in %s subject : ", subject_name)
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)

			_, err := fmt.Sscanf(input, "%f", &subject_grade)

			if err == nil && subject_grade >= 0 && subject_grade <= 100 {
				break
			}

			fmt.Print("\nInvalid input! Please Enter a valid number in range (0 - 100). Please, ")
		}

		subjects[subject_name] = subject_grade
	}

	return user_name, subjects
}