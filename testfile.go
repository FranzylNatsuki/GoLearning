package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
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

	fmt.Printf("Welcome to To-Do-List CLI!-----------------------\n")
	var option int
	j := 0
	var notes [size]metadata

	for option != 6 {
		fmt.Print("Choose an option:\n")
		fmt.Println("1 WRITE")
		fmt.Println("2 READ")
		fmt.Println("3 EDIT MARK")
		fmt.Println("4 VIEW PROPERTIES")
		fmt.Println("5 VIEW ALL DUE REMINDERS")
		fmt.Println("6 EXIT")

		fmt.Scan(&option)

		switch option {
		case 1:
			if j > size {
				fmt.Printf("Max amount of notes\n")
			}
			notes[j] = write()
			j++
		case 2:
			if j == 0 {
				fmt.Println("No Notes Available")
				continue
			}
			fmt.Println("Search by title:")
			search := readLine()
			index := indexSearch(search, notes[:j])
			read(index, notes[:j])
		case 3:
			fmt.Println("Search by title:")
			search := readLine()
			index := indexSearch(search, notes[:j])
			read(index, notes[:j])

			fmt.Println("Accomplished?: ")
			choice := readLine()

			if choice == "Y" || choice == "y" {
				notes[index].mark = true
			} else {
				notes[index].mark = false
			}

		case 4:
		case 5:
			fmt.Printf("View all due reminders\n")
		case 6:
			os.Exit(0)
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
	fmt.Printf("\n=== NOTE DETAILS ===\n")
	fmt.Printf("Title: %s\n", note.title)
	fmt.Printf("Content: %s\n", note.content)
	fmt.Printf("Created: %s\n", note.date_created.Format("2006-01-02 15:04:05"))
	fmt.Printf("Accomplished: %t\n", note.mark)
	fmt.Printf("ID: %d\n", note.id_num)
	fmt.Printf("\n")
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin) // similar to java style, scanner obj
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func write() metadata {
	var note metadata

	fmt.Print("Title: ")
	title := readLine()

	fmt.Print("Content: ")
	content := readLine()

	note.title = title
	note.id_num = rand.Intn(1000)
	note.date_created = time.Now()
	note.mark = false
	note.content = content

	return note
}
