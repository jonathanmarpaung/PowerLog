package main

import "fmt"

// ---------------- STRUCT ---------------------

type Time struct {
	Hour   int
	Minute int
	Second int
}

type PowerLog struct {
	DeviceName     string
	DeviceLocation string
	Power          float64
	Duration       Time
}

// ---------------- GLOBAL ---------------------

const NMAX int = 100

type TabListMenu [NMAX]string
type TabPowerLog [NMAX]PowerLog

// ---------------- MENU ---------------------

// Displays the menu list and returns a valid user selection.
func GetMenuSelection(list TabListMenu, n int) int {
	var i, selection int
	var valid bool

	valid = false

	for !valid {
		// display the menu list with index starting from 1 to n
		for i = 0; i < n; i++ {
			fmt.Printf("%d. %s\n", (i + 1), list[i])
		}

		// get user selection
		fmt.Printf("Select menu: ")
		fmt.Scan(&selection)

		// validate selection
		if selection >= 1 && selection <= n {
			valid = true
		}
	}

	return selection
}

// ---------------- MAIN ---------------------

func main() {
	var running bool
	var menuList TabListMenu
	var nList, selection int

	// initialize menu
	nList = 4
	menuList[0] = "Insert data"
	menuList[1] = "Replace data"
	menuList[2] = "Delete data"
	menuList[3] = "Exit"

	// run the program
	running = true

	for running {
		selection = GetMenuSelection(menuList, nList)

		switch selection {
		case 1:
			// TODO: Create a function to handle insert data
		case 2:
			// TODO: Create a function to handle replace data
		case 3:
			// TODO: Create a function to handle delete data
		case 4:
			running = false
		}
	}
}
