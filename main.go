package main

import (
	"bufio"
	"fmt"
	"os"
	"text/template"

	"github.com/fjukstad/scratch"
)

func main() {
	projects, err := GetProjects("ids")
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.OpenFile("index.html", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	t, err := template.ParseFiles("index.template")
	t.Execute(f, projects)

	fmt.Println("index.html written succesfully")
}

func GetProjects(filename string) ([]*scratch.Project, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	projects := []*scratch.Project{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		id := scanner.Text()
		p, err := scratch.GetProject(id)
		if err != nil {
			fmt.Println("Could not get project ", id)
			continue
		}
		projects = append(projects, p)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
