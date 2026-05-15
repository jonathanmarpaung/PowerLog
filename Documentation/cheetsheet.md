# Framework Cheatsheet
A quick reference guide

Made By: Jonathan Gilbert Marpaung

## 1. Global Constants
| Constant | Type | Value | Description |
| :--- | :--- | :--- | :--- |
| `MaxData` | `int` | `999` | Max limit for log entries/table rows. |
| `MaxPages` | `int` | `21` | Max limit for registered pages. |
| `MaxMenuItem` | `int` | `20` | Max limit for items in a single menu. |
| `MaxColumns` | `int` | `20` | Max limit for table columns. |
| `MaxWrapLines` | `int` | `20` | Max vertical lines for text wrapping. |
| `DEPENDS_ON_LABEL`| `int` | `-1` | Flag for auto-width table columns. |

---

## 2. Core Utilities
Pure Go implementations for type conversion and validation.

| Function Signature | Description |
| :--- | :--- |
| `func IntToStr(number int) string` | Converts an integer to a string. |
| `func FloatToStr(value float64) string` | Converts a float64 to a string (2 decimal precision). |
| `func IsNumber(str string) bool` | Validates if a string contains only valid numbers. |
| `func StrToInt(str string) int` | Converts a numeric string to an integer. |
| `func digitsToString(number int) string` | *(Internal)* Helper for recursive digit extraction. |

---

## 3. Terminal & UI Components
Basic terminal manipulation and message boxes.

| Function Signature | Description |
| :--- | :--- |
| `func ClearTerminal()` | Clears the terminal screen and resets cursor position. |
| `func PrintBoxMessage(message string)` | Prints a bordered box containing a message. |

---

## 4. State & Error Management (Context)
Manages the application lifecycle and centralized error handling.

**Structs:**
* `ErrorHandler{ isError bool, message string }`
* `Context{ page Page, running bool, errorHandler ErrorHandler }`

| Function Signature | Description |
| :--- | :--- |
| `func StartProgram(context *Context)` | Sets the running state to true. |
| `func StopProgram(context *Context)` | Sets the running state to false (Triggers exit). |
| `func IsRunning(context Context) bool` | Returns the current application running state. |
| `func SetError(errorHandler *ErrorHandler, message string)`| Sets an active error state. |
| `func ResetError(errorHandler *ErrorHandler)` | Clears the current error state. |
| `func ShowError(errorHandler ErrorHandler)` | Displays the error using a box message. |
| `func ThrowError(context *Context, message string)` | Triggers an error directly to the context. |
| `func HandlerError(context *Context)` | Displays and resets the current context error. |

---

## 5. Routing & Page Navigation
Manages available pages and current active screen.

**Structs:**
* `Page{ pages [MaxPages]string, currentPage string, n int }`

| Function Signature | Description |
| :--- | :--- |
| `func InsertPage(page *Page, namePage string)` | Registers a new page to the page list. |
| `func DeletePage(page *Page, namePage string)` | Removes a page from the list. |
| `func IsPageExist(page Page, namePage string) bool`| Checks if a page name is registered. |
| `func GetPageIndex(page Page, namePage string) int`| Returns the index of a page (Linear Search). |
| `func AddListPage(context *Context, pageName string)` | Helper to register a page via Context. |
| `func RemovePageFromList(context *Context, pageName string)`| Helper to remove a page via Context. |
| `func SetCurrentPage(context *Context, namePage string)` | Navigates to a specific page. |
| `func GetCurrentPage(context Context) string` | Returns the name of the active page. |

---

## 6. Menu System
Interactive CLI menu handling.

**Structs:**
* `MenuItem{ id string, label string }`
* `Menu{ list [MaxMenuItem]MenuItem, n int }`

| Function Signature | Description |
| :--- | :--- |
| `func NewMenuItem(id string, label string) MenuItem` | Creates a new MenuItem instance. |
| `func AddMenuItem(menu *Menu, item MenuItem)` | Appends an item to the menu list. |
| `func UpdateMenuItem(menu *Menu, index int, item MenuItem)`| Updates a specific menu item. |
| `func RemoveMenuItem(menu *Menu, index int)` | Removes a specific menu item. |
| `func MenuWithIndex(context *Context, menu *Menu) string` | Displays menu, scans input, returns selected `id`. |
| `func AddGlobalMenu(menu *Menu)` | Injects global menus (e.g., "Exit") to a menu list. |
| `func ExecuteGlobalMenu(context *Context, menuID string) bool`| Intercepts and executes global menu actions. |

---

## 7. Dynamic Table Engine
Auto-scaling, text-wrapping CLI table generator.

**Structs:**
* `TableColumn{ label string, width int }`
* `Table{ columns [MaxColumns]TableColumn, n int }`
* `TableRow{ cells [MaxColumns]string, n int }`
* `TableRows{ list [MaxData]TableRow, n int }`
* `WrappedText{ lines [MaxWrapLines]string, n int }`

| Function Signature | Description |
| :--- | :--- |
| `func NewColumn(label string, width int) TableColumn` | Defines a new column. |
| `func AddColumn(table *Table, column TableColumn)` | Appends a column definition to the table. |
| `func AddCell(row *TableRow, data string)` | Appends cell data into a row. |
| `func ResetRows(row *TableRow)` | Clears all cells in a row (Useful for loops). |
| `func AddRow(rows *TableRows, row TableRow)` | Appends a complete row to the TableRows array. |
| `func PrintTable(table *Table, rows *TableRows)` | Renders the complete table to the terminal. |
| `func GetColumnWidth(column TableColumn) int` | *(Internal)* Calculates actual column width. |
| `func WrapText(text string, width int) WrappedText` | *(Internal)* Slices long text into multiple lines. |
| `func PrintTableLine(table *Table, left, mid, right string)`| *(Internal)* Renders horizontal borders. |
| `func PrintSingleRowWrapped(table *Table, row *TableRow)`| *(Internal)* Renders cells supporting multi-line text. |

---

## 8. Domain Models (Log Application)
Specific structs and functions for the electrical device logger.

**Structs:**
* `Duration{ hours int, minutes int, seconds int }`
* `Power` (`float64` alias)
* `Log{ deviceName string, deviceLocation string, power Power, duration Duration }`
* `Logs{ list [MaxData]Log, n int }`

| Function Signature | Description |
| :--- | :--- |
| `func NewDuration(hours, minutes, seconds int) Duration` | Creates a new Duration instance. |
| `func DurationToStr(duration Duration) string` | Formats Duration into a readable string. |
| `func PowerToStr(power Power) string` | Formats Power (Watt) into a readable string. |
| `func AddLog(logs *Logs, name, location string, power Power, duration Duration)` | Appends a new log entry. |
| `func DeleteLog(logs *Logs, index int)` | Removes a log entry at a specific index. |
| `func PrintLogs(logs *Logs)` | Maps log data into the Table Engine and prints it. |
| `func LogPage(context *Context, logs *Logs)` | The main UI view for the Log management screen. |