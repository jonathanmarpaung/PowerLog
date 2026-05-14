package main

import "fmt"

// --------------------= GLOBAL =--------------------

const MaxLogs int = 999
const MaxMenuItem int = 20

// --------------------= UTILITY =--------------------

func digitsToString(number int) string {
	var lastDigit int
	var result string

	if number == 0 {
		result = ""
	} else {
		lastDigit = number % 10
		result = digitsToString(number/10) + string(byte(lastDigit)+'0')
	}

	return result
}

func IntToStr(number int) string {
	var isNegative bool
	var result string

	if number == 0 {
		result = "0"
	} else {
		if number < 0 {
			isNegative = true
			number = -number
		}

		result = digitsToString(number)

		if isNegative {
			result = "-" + result
		}
	}

	return result
}

func FloatToStr(value float64) string {
	var result, precisionStr string
	var valueInt, precisionInt int
	var isNegative bool

	if value < 0 {
		isNegative = true
		value = -value
	}

	valueInt = int(value)
	precisionInt = int((value - float64(valueInt)) * 100)

	if precisionInt == 0 {
		precisionStr = "00"
	} else {
		if precisionInt < 10 {
			precisionStr = "0" + digitsToString(precisionInt)
		} else {
			precisionStr = digitsToString(precisionInt)
		}
	}

	result = IntToStr(valueInt) + "." + precisionStr

	if isNegative {
		result = "-" + result
	}

	return result
}

func IsNumber(str string) bool {
	var isNumber bool
	var index int

	isNumber = true
	index = 0

	if len(str) == 0 {
		isNumber = false
	} else {
		if len(str) == 1 && str[0] == '-' {
			isNumber = false
		}
	}

	for index < len(str) && isNumber {
		if str[index] == '-' {
			if index != 0 {
				isNumber = false
			}
		} else if str[index] < '0' || str[index] > '9' {
			isNumber = false
		}
		index = index + 1
	}

	return isNumber
}

func StrToInt(str string) int {
	var index, result int
	var isNegative bool

	result = 0
	isNegative = false

	for index = 0; index < len(str); index++ {
		if str[index] == '-' {
			if index == 0 {
				isNegative = true
			}
		} else {
			result = (result * 10) + int(str[index]-'0')
		}
	}

	if isNegative {
		result = -result
	}

	return result
}

// --------------------= Box =--------------------
func PrintBoxMessage(message string) {
	var index int

	// head
	fmt.Printf("╭")
	for index = 0; index < len(message)+2; index++ {
		fmt.Printf("─")
	}
	fmt.Printf("╮\n")

	// body
	fmt.Printf("│ %s │\n", message)

	// foot
	fmt.Printf("╰")
	for index = 0; index < len(message)+2; index++ {
		fmt.Printf("─")
	}
	fmt.Printf("╯\n")
}

// --------------------= Terminal =--------------------
func ClearTerminal() {
	// TODO: Clear terminal with library os
	fmt.Print("\033[H\033[2J")
}

// --------------------= Error Handler =--------------------
type ErrorHandler struct {
	isError bool
	message string
}

func ResetError(errorHandler *ErrorHandler) {
	errorHandler.isError = false
	errorHandler.message = ""
}

func SetError(errorHandler *ErrorHandler, message string) {
	errorHandler.isError = true
	errorHandler.message = message
}

func ShowError(errorHandler ErrorHandler) {
	PrintBoxMessage("Error: " + errorHandler.message)
}

// --------------------= Menu =--------------------

type MenuItem struct {
	id    string
	label string
}

type Menu struct {
	list [MaxMenuItem]MenuItem
	n    int
}

func NewMenuItem(id string, label string) MenuItem {
	var item MenuItem

	item.id = id
	item.label = label

	return item
}

func AddMenuItem(menu *Menu, list MenuItem) {
	menu.list[menu.n] = list
	menu.n = menu.n + 1
}

func MenuWithIndex(menu Menu, errorHandler *ErrorHandler) string {
	var inputStr, selection string
	var index, choiceInt int
	var currentItem MenuItem
	var isNumber bool

	// Show menu items
	for index = 1; index <= menu.n; index++ {
		currentItem = menu.list[index-1]
		fmt.Printf("%d. %s\n", index, currentItem.label)
	}

	// input from user
	fmt.Printf("Selection: ")
	fmt.Scan(&inputStr)

	// 3. validate and process number
	isNumber = IsNumber(inputStr)

	if isNumber {
		choiceInt = StrToInt(inputStr)

		if choiceInt >= 1 && choiceInt <= menu.n {
			selection = menu.list[choiceInt-1].id
		} else {
			SetError(errorHandler, "There is no option "+IntToStr(choiceInt)+" in the menu!")
		}
	} else {
		SetError(errorHandler, "Input must be a number!")
	}

	return selection
}

// --------------------= DURATION =--------------------

type Duration struct {
	hours   int
	minutes int
	seconds int
}

func NewDuration(hours int, minutes int, seconds int) Duration {
	var duration Duration

	duration.hours = hours
	duration.minutes = minutes
	duration.seconds = seconds

	return duration
}

func DurationToStr(duration Duration) string {
	var result string

	if duration.hours > 0 {
		result = result + IntToStr(duration.hours) + " Jam "
	}

	if duration.minutes > 0 {
		result = result + IntToStr(duration.minutes) + " Menit "
	}

	result = result + IntToStr(duration.seconds) + " Detik"

	return result
}

// --------------------= POWER =--------------------

type Power float64

func PowerToStr(power Power) string {
	return FloatToStr(float64(power)) + " Watt"
}

// --------------------= LOG =--------------------

type Log struct {
	deviceName     string
	deviceLocation string
	power          Power
	duration       Duration
}

type Logs struct {
	list [MaxLogs]Log
	n    int
}

func AddLog(logs *Logs, deviceName string, deviceLocation string, power Power, duration Duration) {
	var log Log

	if logs.n < MaxLogs {
		log.deviceName = deviceName
		log.deviceLocation = deviceLocation
		log.power = power
		log.duration = duration

		logs.list[logs.n] = log
		logs.n = logs.n + 1
	}
}

func DeleteLog(logs *Logs, index int) {
	var k int

	if index >= 0 && index < logs.n {
		for k = index; k < logs.n-1; k++ {
			logs.list[k] = logs.list[k+1]
		}

		// empty the data
		logs.list[logs.n-1].deviceName = ""
		logs.list[logs.n-1].deviceLocation = ""
		logs.list[logs.n-1].power = 0
		logs.list[logs.n-1].duration = NewDuration(0, 0, 0)

		logs.n = logs.n - 1
	}
}

// --------------------= TABLE =--------------------

func PrintTableHead() {
	fmt.Println("╭─────┬──────────────────────────┬─────────────────────────┬───────────────────────────────────────────────────┬─────────────────╮")
	fmt.Println("│ No. │      Nama Perangkat      │     Lokasi Perangkat    │                  Durasi Pemakaian                 │ Komsumsi Energi │")
}

func PrintTableSeparator() {
	fmt.Println("├─────┼──────────────────────────┼─────────────────────────┼───────────────────────────────────────────────────┼─────────────────┤")
}

func PrintTableFoot() {
	fmt.Println("╰─────┴──────────────────────────┴─────────────────────────┴───────────────────────────────────────────────────┴─────────────────╯")
}

func PrintLogTable(logs Logs) {
	var index int
	var currentLog Log

	PrintTableHead()

	for index = 1; index <= logs.n; index++ {
		PrintTableSeparator()

		currentLog = logs.list[index-1]

		fmt.Printf("│ %-3d │ %-24s │ %-23s │ %-49s │ %-15s │\n",
			index,
			currentLog.deviceName,
			currentLog.deviceLocation,
			DurationToStr(currentLog.duration),
			PowerToStr(currentLog.power),
		)
	}

	PrintTableFoot()
}

// --------------------= MAIN =--------------------

func main() {
	var logs Logs
	var menu Menu
	var errorHandler ErrorHandler
	var running bool
	var selectedMenu string

	// Menu
	AddMenuItem(&menu, NewMenuItem("Insert", "Insert Data"))
	AddMenuItem(&menu, NewMenuItem("Update", "Update Data"))
	AddMenuItem(&menu, NewMenuItem("Delete", "Delete Data"))
	AddMenuItem(&menu, NewMenuItem("Exit", "Exit"))

	// Dummy data
	AddLog(&logs, "Laptop", "Kamar Tidur", 123.23, NewDuration(1, 20, 30))
	AddLog(&logs, "Kipas Angin", "Ruang Tamu", 45.50, NewDuration(4, 0, 0))

	running = true

	for running {
		ClearTerminal()
		PrintLogTable(logs)

		if errorHandler.isError {
			ShowError(errorHandler)
			ResetError(&errorHandler)
		}

		selectedMenu = MenuWithIndex(menu, &errorHandler)

		switch selectedMenu {
		case "Insert":
			// TODO: create function to handle insert data, example: InsertLogFeature(&logs, &errorHandler)
		case "Update":
			// TODO: create function to handle update data, example: UpdateLogFeature(&logs, &errorHandler)
		case "Delete":
			// TODO: create function to handle delete data, example: DeleteLogFeature(&logs, &errorHandler)
		case "Exit":
			running = false
		}
	}
}
