package main

import (
	"fmt"
	"os"
	"text/template"
)

type Player struct {
	Name string
	HP   int
}

func (p *Player) JobList() []string {
	return []string{
		"Magician",
		"Priest",
		"Knight",
		"Holy Knight",
	}
}

func main() {
	fmt.Println("--- sample template's index ---")

	t := template.Must(template.ParseFiles("templates/index-sample.tmpl"))
	err := t.Execute(os.Stdout, map[string]interface{}{
		"SampleSlice": []int{
			1, 2, 3, 20000, 5,
		},
		"SampleMap": map[int]string{
			1: "uno",
			2: "dos",
			3: "tres",
		},
		"P1": &Player{
			Name: "Player 1",
			HP:   200,
		},
	})
	if err != nil {
		panic(err)
	}

}
