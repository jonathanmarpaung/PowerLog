package main

import "fmt"

// --------------------= GLOBAL =--------------------

const MaxData int = 999

// --------------------= FRAMEWORK =--------------------
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

func InputString() string {
	var result string
	var char rune
	var isReading bool

	isReading = true
	result = ""

	fmt.Scanf("%c", &char)
	if char != '\n' {
		result = result + string(char)
	}

	for isReading {
		fmt.Scanf("%c", &char)
		if char == '\n' {
			isReading = false
		} else if char != '\r' {
			result = result + string(char)
		}
	}

	return result
}

//--------------------= STYLING & CANVAS ENGINE =--------------------

// Category 1: Position Element In Canvas
const PosLeft int = 100
const PosCenter int = 200
const PosRight int = 300

// Category 2: Align Text
const AlignLeft int = 10
const AlignCenter int = 20
const AlignRight int = 30

// Category 3: Size Element
const SizeFit int = 1
const SizeFull int = 2

type Canvas struct {
	width int
}

type Padding struct {
	left  int
	right int
}

type StyleConfig struct {
	position  int
	alignment int
	size      int
}

func NewCanvas(width int) Canvas {
	var canvas Canvas

	canvas.width = width

	return canvas
}

func DecodeStyle(styleCode int) StyleConfig {
	var config StyleConfig

	config.position = (styleCode / 100) * 100
	config.alignment = ((styleCode % 100) / 10) * 10
	config.size = styleCode % 10

	return config
}

func CalculatePadding(totalWidth int, contentWidth int, alignMode int) Padding {
	var result Padding
	var remaining int

	remaining = totalWidth - contentWidth

	// Prevent padding negatif if content wider than length of room
	if remaining < 0 {
		remaining = 0
	}

	// Calculate space from mode (Position or Align)
	switch {
	case alignMode == PosLeft || alignMode == AlignLeft:
		result.left = 0
		result.right = remaining
	case alignMode == PosCenter || alignMode == AlignCenter:
		result.left = remaining / 2
		result.right = remaining - result.left
	case alignMode == PosRight || alignMode == AlignRight:
		result.left = remaining
		result.right = 0
	default:
		result.left = 0
		result.right = remaining
	}

	return result
}

func PrintSpace(count int) {
	var i int

	for i = 0; i < count; i++ {
		fmt.Printf(" ")
	}
}

// --------------------= Box =--------------------

func PrintBoxMessage(context *Context, message string, styleCode int) {
	var styleConfig StyleConfig
	var margin, textPad Padding
	var boxWidth, maxLineWidth, index, lineIndex int
	var splitLines WrappedText
	var currentLine string

	styleConfig = DecodeStyle(styleCode)
	splitLines = WrapText(message, context.canvas.width-4)

	maxLineWidth = 0
	for index = 0; index < splitLines.n; index++ {
		if len(splitLines.lines[index]) > maxLineWidth {
			maxLineWidth = len(splitLines.lines[index])
		}
	}

	if styleConfig.size == SizeFull {
		boxWidth = context.canvas.width - 4
	} else {
		boxWidth = maxLineWidth
	}

	margin = CalculatePadding(context.canvas.width, boxWidth+4, styleConfig.position)

	// Print Head
	PrintSpace(margin.left)
	fmt.Printf("╭")
	for index = 0; index < boxWidth+2; index++ {
		fmt.Printf("─")
	}
	fmt.Printf("╮\n")

	// Print Body
	for lineIndex = 0; lineIndex < splitLines.n; lineIndex++ {
		currentLine = splitLines.lines[lineIndex]
		textPad = CalculatePadding(boxWidth, len(currentLine), styleConfig.alignment)

		PrintSpace(margin.left)
		fmt.Printf("│ ")
		PrintSpace(textPad.left)
		fmt.Printf("%s", currentLine)
		PrintSpace(textPad.right)
		fmt.Printf(" │\n")
	}

	// Print Foot
	PrintSpace(margin.left)
	fmt.Printf("╰")
	for index = 0; index < boxWidth+2; index++ {
		fmt.Printf("─")
	}
	fmt.Printf("╯\n")
}

func PrintMessage(context *Context, message string, styleCode int) {
	var styleConfig StyleConfig
	var margin, textPad Padding
	var boxWidth, maxLineWidth, index, lineIndex int
	var splitLines WrappedText
	var currentLine string

	styleConfig = DecodeStyle(styleCode)
	splitLines = WrapText(message, context.canvas.width)

	maxLineWidth = 0
	for index = 0; index < splitLines.n; index++ {
		if len(splitLines.lines[index]) > maxLineWidth {
			maxLineWidth = len(splitLines.lines[index])
		}
	}

	if styleConfig.size == SizeFull {
		boxWidth = context.canvas.width
	} else {
		boxWidth = maxLineWidth
	}

	margin = CalculatePadding(context.canvas.width, boxWidth+4, styleConfig.position)

	// Print Body
	for lineIndex = 0; lineIndex < splitLines.n; lineIndex++ {
		currentLine = splitLines.lines[lineIndex]
		textPad = CalculatePadding(boxWidth, len(currentLine), styleConfig.alignment)

		PrintSpace(margin.left)
		PrintSpace(textPad.left)
		fmt.Printf("%s", currentLine)
		PrintSpace(textPad.right)
		fmt.Printf("\n")
	}
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

func ShowError(context *Context, errorHandler ErrorHandler) {
	PrintBoxMessage(context, "Error: "+errorHandler.message, PosCenter)
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
	canvas       Canvas
}

func InitializeContext(context *Context, canvasWidth int) {
	context.canvas = NewCanvas(canvasWidth)
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
	ShowError(context, context.errorHandler)
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
		menu.list[index] = item
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

func MenuWithIndex(context *Context, menu *Menu, styleCode int) string {
	var inputStr, selection, formattedItem string
	var index, choiceInt, maxItemWidth, menuWidth, currentItemLen int
	var currentItem MenuItem
	var isNumber bool
	var styleConfig StyleConfig
	var margin, textPad Padding

	styleConfig = DecodeStyle(styleCode)
	maxItemWidth = 0

	for index = 1; index <= menu.n; index++ {
		currentItem = menu.list[index-1]
		formattedItem = IntToStr(index) + ". " + currentItem.label
		currentItemLen = len(formattedItem)

		if currentItemLen > maxItemWidth {
			maxItemWidth = currentItemLen
		}
	}

	if styleConfig.size == SizeFull {
		menuWidth = context.canvas.width
	} else {
		menuWidth = maxItemWidth
	}

	margin = CalculatePadding(context.canvas.width, menuWidth, styleConfig.position)

	for index = 1; index <= menu.n; index++ {
		currentItem = menu.list[index-1]
		formattedItem = IntToStr(index) + ". " + currentItem.label

		textPad = CalculatePadding(menuWidth, len(formattedItem), styleConfig.alignment)

		PrintSpace(margin.left)
		PrintSpace(textPad.left)
		fmt.Printf("%s", formattedItem)
		PrintSpace(textPad.right)
		fmt.Printf("\n")
	}

	PrintSpace(margin.left)
	fmt.Printf("Selection: ")
	fmt.Scan(&inputStr)

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
	style int
}

type Table struct {
	columns [MaxColumns]TableColumn
	n       int
	style   int
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

func InitializeTable(table *Table, styleCode int) {
	table.style = styleCode
}

func NewColumn(label string, width int, style int) TableColumn {
	var column TableColumn

	column.label = label
	column.width = width
	column.style = style

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

func GetColumnWidth(context *Context, table *Table, column TableColumn) int {
	var width, effectiveCanvas, bordersAndSpaces int
	var styleConfig StyleConfig

	styleConfig = DecodeStyle(table.style)

	if styleConfig.size == SizeFull {
		bordersAndSpaces = (table.n * 3) + 1
		effectiveCanvas = context.canvas.width - bordersAndSpaces

		width = (effectiveCanvas * column.width) / 100

	} else {
		if column.width == DEPENDS_ON_LABEL {
			width = len(column.label)
		} else {
			width = column.width
		}
	}

	return width
}

func GetTableTotalWidth(context *Context, table *Table) int {
	var total, index int

	total = (table.n * 3) + 1
	for index = 0; index < table.n; index++ {
		total = total + GetColumnWidth(context, table, table.columns[index])
	}

	return total
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
			if text[index] == '\n' {
				if result.n < MaxWrapLines {
					result.lines[result.n] = currentLine
					result.n = result.n + 1
					currentLine = ""
				}
			} else {
				currentLine = currentLine + string(text[index])

				if len(currentLine) == width {
					if result.n < MaxWrapLines {
						result.lines[result.n] = currentLine
						result.n = result.n + 1
						currentLine = ""
					}
				}
			}
		}

		if len(currentLine) > 0 && result.n < MaxWrapLines {
			result.lines[result.n] = currentLine
			result.n = result.n + 1
		}
	}

	return result
}

func PrintTableLine(context *Context, table *Table, leftChar string, midChar string, rightChar string, marginLeft int) {
	var k, index, length int
	var currentColumn TableColumn

	PrintSpace(marginLeft)
	fmt.Printf("%s", leftChar)

	for index = 0; index < table.n; index++ {
		currentColumn = table.columns[index]
		length = GetColumnWidth(context, table, currentColumn)

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

func PrintSingleRowWrapped(context *Context, table *Table, row *TableRow, marginLeft int) {
	var colIndex, lineIndex, maxLines, width int
	var wrappedCells [MaxColumns]WrappedText
	var currentColumn TableColumn
	var currentCell string
	var currentWrapped WrappedText
	var textPad Padding

	maxLines = 1

	for colIndex = 0; colIndex < table.n; colIndex++ {
		currentColumn = table.columns[colIndex]
		width = GetColumnWidth(context, table, currentColumn)

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
		PrintSpace(marginLeft)
		fmt.Printf("│")

		for colIndex = 0; colIndex < table.n; colIndex++ {
			currentColumn = table.columns[colIndex]
			width = GetColumnWidth(context, table, currentColumn)
			currentWrapped = wrappedCells[colIndex]

			if lineIndex < currentWrapped.n {
				currentCell = currentWrapped.lines[lineIndex]
			} else {
				currentCell = ""
			}

			textPad = CalculatePadding(width, len(currentCell), currentColumn.style)

			fmt.Printf(" ")
			PrintSpace(textPad.left)
			fmt.Printf("%s", currentCell)
			PrintSpace(textPad.right)
			fmt.Printf(" │") // Spasi pembatas kanan + border
		}
		fmt.Printf("\n")
	}
}

func PrintTable(context *Context, table *Table, rows *TableRows) {
	var headerRow TableRow
	var index, tableTotalWidth int
	var styleConfig StyleConfig
	var margin Padding

	styleConfig = DecodeStyle(table.style)
	tableTotalWidth = GetTableTotalWidth(context, table)
	margin = CalculatePadding(context.canvas.width, tableTotalWidth, styleConfig.position)

	for index = 0; index < table.n; index++ {
		AddCell(&headerRow, table.columns[index].label)
	}

	// 3. Cetak Keseluruhan dengan menyuntikkan margin.left
	PrintTableLine(context, table, "╭", "┬", "╮", margin.left)

	PrintSingleRowWrapped(context, table, &headerRow, margin.left)

	for index = 0; index < rows.n; index++ {
		PrintTableLine(context, table, "├", "┼", "┤", margin.left)
		PrintSingleRowWrapped(context, table, &rows.list[index], margin.left)
	}

	PrintTableLine(context, table, "╰", "┴", "╯", margin.left)
}

// --------------------= MAIN CODE =--------------------
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

func PrintLogs(context *Context, logs *Logs) {
	var index int
	var rowTable Table
	var rows TableRows
	var currentRow TableRow

	InitializeTable(&rowTable, PosCenter+SizeFull)

	AddColumn(&rowTable, NewColumn("Number", 10, AlignCenter))
	AddColumn(&rowTable, NewColumn("Device Name", 25, AlignLeft))
	AddColumn(&rowTable, NewColumn("Device Location", 20, AlignLeft))
	AddColumn(&rowTable, NewColumn("Duration", 30, AlignRight))
	AddColumn(&rowTable, NewColumn("Power Used", 15, AlignRight))

	for index = 0; index < logs.n; index++ {
		ResetRows(&currentRow)
		AddCell(&currentRow, IntToStr(index+1))
		AddCell(&currentRow, logs.list[index].deviceName)
		AddCell(&currentRow, logs.list[index].deviceLocation)
		AddCell(&currentRow, DurationToStr(logs.list[index].duration))
		AddCell(&currentRow, PowerToStr(logs.list[index].power))
		AddRow(&rows, currentRow)
	}

	PrintTable(context, &rowTable, &rows)
}

func LogPage(context *Context, logs *Logs) {
	var menu Menu
	var selectedMenu string

	// Print Table Log
	PrintLogs(context, logs)

	// Menu
	AddMenuItem(&menu, NewMenuItem("Insert", "Insert Data"))
	AddMenuItem(&menu, NewMenuItem("Update", "Update Data"))
	AddMenuItem(&menu, NewMenuItem("Delete", "Delete Data"))
	AddGlobalMenu(&menu)

	selectedMenu = MenuWithIndex(context, &menu, 0)

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

// --------------------= MAIN =--------------------

func main() {
	var logs Logs
	var context Context

	// Initialize Context
	InitializeContext(&context, 100)

	// Initialize Page
	AddListPage(&context, "Log")

	// Dummy data
	AddLog(&logs, "Laptop", "Kamar Tidur", 123.23, NewDuration(1, 20, 30))
	AddLog(&logs, "Kipas Angin", "Ruang Tamu", 45.50, NewDuration(4, 0, 0))

	// running
	StartProgram(&context)

	// NOTE: name page must be valid!
	SetCurrentPage(&context, "Log")

	for IsRunning(context) {
		ClearTerminal()

		PrintMessage(&context, "Page: "+GetCurrentPage(context), SizeFull+AlignCenter)

		// Handle error
		if IsThereAnError(context) {
			HandlerError(&context)
		}

		switch GetCurrentPage(context) {
		case "Log":
			LogPage(&context, &logs)

		// prevent crazy loop
		default:
			LogPage(&context, &logs)
		}
	}
}
