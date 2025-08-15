package main

import (
	"fmt"
	"math/rand"

	"time"
)

const size int = 20

type metadata struct {
	title        string
	id_num       int
	content      string
	mark         bool
	date_created time.Time
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Welcome to To-Do-List CLI!-----------------------")
	var option int
	j := 0

	for option != 5 {
		fmt.Print("Choose an option:")
		var notes [size]metadata

		fmt.Println("1 WRITE")
		fmt.Println("2 READ")
		fmt.Println("3 EDIT MARK")
		fmt.Println("4 VIEW PROPERTIES")

		fmt.Scan(&option)

		switch option {
		case 1:
			notes[j] = write()
			j++
		case 2:
			fmt.Println("Search by title:")
			var search string
			fmt.Scan(&search)
			index := indexSearch(search, notes[:j])
			read(index, notes[:j])
		case 3:
			fmt.Println("Search by title:")
			var search string
			fmt.Scan(&search)
			index := indexSearch(search, notes[:j])
			read(index, notes[:j])

			fmt.Println("Accomplished?: ")
			var choice string

			fmt.Scan(&choice)

			if choice == "Y" || choice == "y" {
				notes[j].mark = true
			} else {
				notes[j].mark = false
			}

		case 4:
		case 5:

		}
	}

}

func indexSearch(search string, notes []metadata) int {
	for i := 0; i < len(notes); i++ {
		if notes[i].title == search {
			return i
		}
	}
	return -1
}

func read(index int, notes []metadata) {
	if index == -1 {
		fmt.Println("Note not found!")
		return
	}
	note := notes[index]
	fmt.Println(note.title)
	fmt.Println(note.content)
	fmt.Println(note.date_created)
	fmt.Println(note.mark)
	fmt.Println(note.id_num)
}

func write() metadata {
	var note metadata

	fmt.Println("Title: ")
	var title string
	fmt.Scan(&title)

	fmt.Println("Content: ")
	var content string
	fmt.Scan(&content)

	note.title = title
	note.id_num = rand.Int()
	note.date_created = time.Now()
	note.mark = false
	note.content = content

	return note
}
