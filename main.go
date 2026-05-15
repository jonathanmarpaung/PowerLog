package main

import "fmt"

// --------------------= GLOBAL =--------------------

const MaxData int = 999

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

// --------------------= PAGES =--------------------
const MaxPages int = 20

type Page struct {
	pages       [MaxPages]string
	currentPage string
	n           int
}

func InsertPage(page *Page, namePage string) {
	if page.n < MaxPages {
		page.pages[page.n] = namePage
		page.n = page.n + 1
	}
}

func GetPageIndex(page Page, namePage string) int {
	var index, atPosition int

	for index < page.n && page.pages[index] != namePage {
		index = index + 1
	}

	if index == page.n {
		atPosition = -1
	} else {
		atPosition = index
	}

	return atPosition
}

func IsPageExist(page Page, namePage string) bool {
	return GetPageIndex(page, namePage) != -1
}

func DeletePage(page *Page, namePage string) {
	var index, targetIndex int

	if page.n > 0 {
		targetIndex = GetPageIndex(*page, namePage)

		if targetIndex != -1 {
			for index = targetIndex; index < page.n-1; index++ {
				page.pages[index] = page.pages[index+1]
			}

			page.pages[page.n-1] = ""
			page.n = page.n - 1
		}
	}
}

// --------------------= PROGRAM =--------------------

type Context struct {
	page         Page
	running      bool
	errorHandler ErrorHandler
}

func IsRunning(context Context) bool {
	return context.running
}

func StartProgram(context *Context) {
	context.running = true
}

func StopProgram(context *Context) {
	context.running = false
}

func IsThereAnError(context Context) bool {
	return context.errorHandler.isError
}

func HandlerError(context *Context) {
	ShowError(context.errorHandler)
	ResetError(&context.errorHandler)
}

func ThrowError(context *Context, message string) {
	SetError(&context.errorHandler, message)
}

func AddListPage(context *Context, pageName string) {
	InsertPage(&context.page, pageName)
}

func RemovePageFromList(context *Context, pageName string) {
	DeletePage(&context.page, pageName)
}

func GetCurrentPage(context Context) string {
	return context.page.currentPage
}

func SetCurrentPage(context *Context, namePage string) {
	if IsPageExist(context.page, namePage) {
		context.page.currentPage = namePage
	} else {
		ThrowError(context, "There's no exist page "+namePage)
	}
}

// --------------------= Menu =--------------------

const MaxMenuItem int = 20

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

func AddMenuItem(menu *Menu, item MenuItem) {
	if menu.n < MaxMenuItem {
		menu.list[menu.n] = item
		menu.n = menu.n + 1
	}
}

func UpdateMenuItem(menu *Menu, index int, item MenuItem) {
	if index >= 0 && index < menu.n {
		menu.list[index] = item // Jauh lebih enak dibaca
	}
}

func RemoveMenuItem(menu *Menu, index int) {
	var k int

	if index >= 0 && index < menu.n {
		for k = index; k < menu.n-1; k++ {
			menu.list[k] = menu.list[k+1]
		}

		// empty the data
		menu.list[menu.n-1].id = ""
		menu.list[menu.n-1].label = ""

		menu.n = menu.n - 1
	}
}

func MenuWithIndex(context *Context, menu *Menu) string {
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
			ThrowError(context, "There is no option "+IntToStr(choiceInt)+" in the menu!")
		}
	} else {
		ThrowError(context, "Input must be a number!")
	}

	return selection
}

func AddGlobalMenu(menu *Menu) {
	AddMenuItem(menu, NewMenuItem("Exit", "Exit"))
}

func ExecuteGlobalMenu(context *Context, menuID string) bool {
	var isGlobal bool
	isGlobal = false

	switch menuID {
	case "Exit":
		StopProgram(context)
		isGlobal = true
	}

	return isGlobal
}

// --------------------= TABLE =--------------------

const MaxColumns int = 20
const DEPENDS_ON_LABEL int = -1
const MaxWrapLines int = 20

type TableColumn struct {
	label string
	width int
}

type Table struct {
	columns [MaxColumns]TableColumn
	n       int
}

type TableRow struct {
	cells [MaxColumns]string
	n     int
}

type TableRows struct {
	list [MaxData]TableRow
	n    int
}

type WrappedText struct {
	lines [MaxWrapLines]string
	n     int
}

func NewColumn(label string, width int) TableColumn {
	var column TableColumn

	column.label = label
	column.width = width

	return column
}

func AddColumn(table *Table, column TableColumn) {
	if table.n < MaxColumns {
		table.columns[table.n] = column
		table.n = table.n + 1
	}
}

func AddCell(row *TableRow, data string) {
	if row.n < MaxColumns {
		row.cells[row.n] = data
		row.n = row.n + 1
	}
}

func ResetRows(row *TableRow) {
	var index int

	for index = 0; index < row.n; index++ {
		row.cells[index] = ""
	}

	row.n = 0
}

func AddRow(rows *TableRows, row TableRow) {
	if rows.n < MaxData {
		rows.list[rows.n] = row
		rows.n = rows.n + 1
	}
}

func GetColumnWidth(column TableColumn) int {
	var width int

	if column.width == DEPENDS_ON_LABEL {
		width = len(column.label)
	} else {
		width = column.width
	}

	return width
}

func WrapText(text string, width int) WrappedText {
	var result WrappedText
	var currentLine string
	var index, textLen int

	textLen = len(text)

	if width <= 0 || textLen == 0 {
		result.lines[0] = text
		result.n = 1
	} else {
		for index = 0; index < textLen; index++ {
			currentLine = currentLine + string(text[index])

			if len(currentLine) == width {
				result.lines[result.n] = currentLine
				result.n = result.n + 1
				currentLine = ""
			}
		}

		if len(currentLine) > 0 {
			result.lines[result.n] = currentLine
			result.n = result.n + 1
		}
	}

	return result
}

func PrintTableLine(table *Table, leftChar string, midChar string, rightChar string) {
	var k, index, length int
	var currentColumn TableColumn

	fmt.Printf("%s", leftChar)

	for index = 0; index < table.n; index++ {
		currentColumn = table.columns[index]
		length = GetColumnWidth(currentColumn)

		for k = 0; k < length+2; k++ {
			fmt.Printf("─")
		}

		if index == table.n-1 {
			fmt.Printf("%s\n", rightChar)
		} else {
			fmt.Printf("%s", midChar)
		}
	}
}

func PrintSingleRowWrapped(table *Table, row *TableRow) {
	var colIndex, lineIndex, maxLines, width int
	var wrappedCells [MaxColumns]WrappedText
	var currentColumn TableColumn
	var currentCell string
	var currentWrapped WrappedText

	maxLines = 1

	for colIndex = 0; colIndex < table.n; colIndex++ {
		currentColumn = table.columns[colIndex]
		width = GetColumnWidth(currentColumn)

		if colIndex < row.n {
			currentCell = row.cells[colIndex]
		} else {
			currentCell = ""
		}

		currentWrapped = WrapText(currentCell, width)
		wrappedCells[colIndex] = currentWrapped

		if currentWrapped.n > maxLines {
			maxLines = currentWrapped.n
		}
	}

	for lineIndex = 0; lineIndex < maxLines; lineIndex++ {
		fmt.Printf("│")

		for colIndex = 0; colIndex < table.n; colIndex++ {
			width = GetColumnWidth(table.columns[colIndex])
			currentWrapped = wrappedCells[colIndex]

			if lineIndex < currentWrapped.n {
				fmt.Printf(" %-*s │", width, currentWrapped.lines[lineIndex])
			} else {
				fmt.Printf(" %-*s │", width, "")
			}
		}
		fmt.Printf("\n")
	}
}

func PrintTable(table *Table, rows *TableRows) {
	var headerRow TableRow
	var index int

	for index = 0; index < table.n; index++ {
		AddCell(&headerRow, table.columns[index].label)
	}

	PrintTableLine(table, "╭", "┬", "╮")

	PrintSingleRowWrapped(table, &headerRow)

	for index = 0; index < rows.n; index++ {
		PrintTableLine(table, "├", "┼", "┤")
		PrintSingleRowWrapped(table, &rows.list[index])
	}

	PrintTableLine(table, "╰", "┴", "╯")
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
	list [MaxData]Log
	n    int
}

func AddLog(logs *Logs, deviceName string, deviceLocation string, power Power, duration Duration) {
	var log Log

	if logs.n < MaxData {
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

func PrintLogs(logs *Logs) {
	var index int
	var rowTable Table
	var rows TableRows
	var currentRow TableRow

	AddColumn(&rowTable, NewColumn("No.", DEPENDS_ON_LABEL))
	AddColumn(&rowTable, NewColumn("Nama Perangkat", DEPENDS_ON_LABEL))
	AddColumn(&rowTable, NewColumn("Lokasi Perangkat", DEPENDS_ON_LABEL))
	AddColumn(&rowTable, NewColumn("Durasi Penggunaan", DEPENDS_ON_LABEL))
	AddColumn(&rowTable, NewColumn("Daya Tergunakan", DEPENDS_ON_LABEL))

	for index = 0; index < logs.n; index++ {
		ResetRows(&currentRow)

		AddCell(&currentRow, IntToStr(index+1))
		AddCell(&currentRow, logs.list[index].deviceName)
		AddCell(&currentRow, logs.list[index].deviceLocation)
		AddCell(&currentRow, DurationToStr(logs.list[index].duration))
		AddCell(&currentRow, PowerToStr(logs.list[index].power))

		AddRow(&rows, currentRow)
	}

	PrintTable(&rowTable, &rows)
}

func LogPage(context *Context, logs *Logs) {
	var menu Menu
	var selectedMenu string

	// Print Table Log
	PrintLogs(logs)

	// Menu
	AddMenuItem(&menu, NewMenuItem("Insert", "Insert Data"))
	AddMenuItem(&menu, NewMenuItem("Update", "Update Data"))
	AddMenuItem(&menu, NewMenuItem("Delete", "Delete Data"))
	AddGlobalMenu(&menu)

	selectedMenu = MenuWithIndex(context, &menu)

	if !ExecuteGlobalMenu(context, selectedMenu) {
		switch selectedMenu {
		case "Insert":
			// TODO: create function to handle insert data, example: InsertLogFeature(&logs, &errorHandler)
		case "Update":
			// TODO: create function to handle update data, example: UpdateLogFeature(&logs, &errorHandler)
		case "Delete":
			// TODO: create function to handle delete data, example: DeleteLogFeature(&logs, &errorHandler)
		}
	}
}

// --------------------= App Intergration =--------------------

func HomePage(context *Context) {
	var menu Menu
	var selectedMenu string

	// 1. Main Page Header Display
	PrintBoxMessage("ENERGY LOG MANAGER - MAIN MENU")
	fmt.Println("Welcome! Please select a module to start managing your devices.")
	fmt.Println("")

	// 2. Setup Navigation Menu
	AddMenuItem(&menu, NewMenuItem("GoLog", "Device Usage Logs"))
	AddMenuItem(&menu, NewMenuItem("About", "About Application"))
	AddGlobalMenu(&menu)

	// 3. Display Menu and Get Selection
	selectedMenu = MenuWithIndex(context, &menu)

	// 4. Navigation Logic (Router)
	if !ExecuteGlobalMenu(context, selectedMenu) {
		switch selectedMenu {
		case "GoLog":
			SetCurrentPage(context, "Log")
		case "About":
			// Example of a simple static display for the About page
			ClearTerminal()
			PrintBoxMessage("About: Framework v1.0")
			fmt.Println("A simple procedural CLI framework for educational purposes.")
			fmt.Println("")
			fmt.Println("Press Enter to back...")
			fmt.Scanln(&selectedMenu) // Simple pause
		}
	}
}

// --------------------= MAIN =--------------------

func main() {
	var logs Logs
	var context Context

	// 1. Register All Available Pages
	AddListPage(&context, "Home")
	AddListPage(&context, "Log")

	// 2. Data Initialization (Dummy Data)
	AddLog(&logs, "Laptop", "Bedroom", 120.00, NewDuration(2, 0, 0))
	AddLog(&logs, "LED Lamp", "Porch", 15.50, NewDuration(12, 0, 0))

	// 3. Start Program and Set Initial Page
	StartProgram(&context)
	SetCurrentPage(&context, "Home")

	// 4. Main Application Loop
	for IsRunning(context) {
		ClearTerminal()

		// Global Error Handler: Appears on every page if an error occurs
		if IsThereAnError(context) {
			HandlerError(&context)
		}

		// Breadcrumb / Page Location Indicator
		fmt.Println("Location:", GetCurrentPage(context))
		fmt.Println("---------------------------------------------------------")

		// Router: Selects the page function based on CurrentPage
		switch GetCurrentPage(context) {
		case "Home":
			HomePage(&context)
		case "Log":
			LogPage(&context, &logs)
		default:
			HomePage(&context)
		}
	}

	// Termination message when the program stops
	ClearTerminal()
	fmt.Println("Program terminated. Thank you for using Energy Log Manager!")
}
