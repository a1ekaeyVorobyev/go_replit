package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
  "strings"
)

type student struct {
	name  string
	age   int
	grade int
}

func main() {
	mStudent := make(map[string]*student)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		_ = <-sigs
		fmt.Println("------------------")
		fmt.Println("Выход из программы")
		for _, v := range mStudent {
			fmt.Print(*v)
		}
		os.Exit(0)
	}()

	var name string
	var age, grade int
	
  for {
		fmt.Print("Имя студента: ")
		fmt.Scan(&name)
		fmt.Print("возраст: ")
		fmt.Scan(&age)
		fmt.Print("класс: ")
		fmt.Scan(&grade)
		
    s := student{name, age, grade}
		if err := s.Check(); err != nil {
			fmt.Println(err)
			fmt.Println("Повторите ввод студента")
			continue
		}

		if _, ok := mStudent[name]; ok == true {
			var x string
			fmt.Println("данный студент уже существует")
			fmt.Print("Перписать?(y/n)")
			fmt.Scan(&x)
			if strings.ToUpper(x) == "Y" {
				mStudent[name] = &s
			}
		} else {
			mStudent[name] = &s
		}
	}
}

func (s student) String() string {
	return fmt.Sprintf("Имя: %v; возраст: %v; класс: %v.\n", s.name, s.age, s.grade)
}

func (s student) Check() error {
	errorMessage := ""
	if s.name == "" {
		errorMessage = "Не введенно имя студента.\n"
	}
	if !(s.age > 16 && s.age < 40) {
		errorMessage += "Не введенно возраст студента.\n"
	}
	if !(s.grade > 0 && s.grade < 12) {
		errorMessage += "Не введенно класс студента.\n"
	}
	if errorMessage == "" {
		return nil
	}
	return fmt.Errorf(errorMessage)
}
