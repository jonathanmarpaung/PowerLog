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

// ---------------- SUB PROGRAM ---------------------

// Displays the menu list and returns a valid user selection.
func GetMenuSelection(list TabListMenu, n int) int {
	// TODO: improve menu style
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

/**
 * I.S. TODO
 * F.S. TODO
 */
func InsertLog(log *TabPowerLog, n *int) {
	// TODO: handle if array have been in maximum capacity
	fmt.Printf("Device Name: ")
	fmt.Scan(&log[*n].DeviceName)

	fmt.Printf("Device Location: ")
	fmt.Scan(&log[*n].DeviceLocation)

	fmt.Printf("Power Used (watt): ")
	fmt.Scan(&log[*n].Power)

	fmt.Printf("Duration (Hour): ")
	fmt.Scan(&log[*n].Duration.Hour)

	fmt.Printf("Duration (Minute): ")
	fmt.Scan(&log[*n].Duration.Minute)

	fmt.Printf("Duration (Second): ")
	fmt.Scan(&log[*n].Duration.Second)

	// adding one to inform new data has been added
	*n = *n + 1
}

/**
 * I.S. array filled with log and n as length log in array
 * F.S. showing log in array as table
 */
func ShowLog(log TabPowerLog, n int) {
	// TODO: improve table style
	var i int

	fmt.Printf("%-3s %-15s %-20s %-13s %s\n",
		"ID", "Device Name", "Device Location", "Power (Watt)", "Duration")

	for i = 0; i < n; i++ {
		fmt.Printf("%-3d %-15s %-20s %-13.2f %d hours %d minutes %d seconds\n",
			i, log[i].DeviceName, log[i].DeviceLocation, log[i].Power,
			log[i].Duration.Hour, log[i].Duration.Minute, log[i].Duration.Second)
	}
}

// ---------------- MAIN ---------------------

func main() {
	var running bool
	var log TabPowerLog
	var nLog int
	var menuList TabListMenu
	var nMenuList, selection int

	// initialize menu
	nMenuList = 5
	menuList[0] = "Show data"
	menuList[1] = "Insert data"
	menuList[2] = "Replace data"
	menuList[3] = "Delete data"
	menuList[4] = "Exit"

	// initialize log as empty array
	nLog = 0

	// run the program
	running = true

	for running {
		selection = GetMenuSelection(menuList, nMenuList)

		switch selection {
		case 1:
			ShowLog(log, nLog)
		case 2:
			InsertLog(&log, &nLog)
		case 3:
			// TODO: Create a function to handle replace data
		case 4:
			// TODO: Create a function to handle delete data
		case 5:
			running = false
		}
	}
}
