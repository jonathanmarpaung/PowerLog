# Framework - Home Page & Navigation Implementation

This documentation explains how to implement the Home Page and configure the routing system so the application can navigate dynamically between modules.

Made by: Jonathan Gilbert Marpaung

## 1. HomePage Implementation

This function should be placed in the `App Intergration Section` section. This page serves as the main entry point for the user.

```go
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
			PrintBoxMessage("About: Go-CLI Framework v1.0")
			fmt.Println("A simple procedural CLI framework for educational purposes.")
			fmt.Println("")
			fmt.Println("Press Enter to back...")
			fmt.Scanln(&selectedMenu) // Simple pause
		}
	}
}

```

## 2. Main Function Update

Update the `main` function to register new pages and set up the main loop to support dynamic page navigation.

```go
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

```

## 3. Advantages of the Navigation Structure

| Feature | Benefit |
| --- | --- |
| **Modularity** | Adding a new page only requires 3 steps: Create the Page function, call `AddListPage`, and add a `case` in `main`. |
| **Breadcrumbs** | Users always know their current location (e.g., "Location: Log") so they don't get lost. |
| **Separation of Logic** | `main` only acts as a traffic controller, while the specific feature logic resides in their respective functions. |
| **Global State** | By using `SetCurrentPage`, the active page state is safely stored in the `Context`, preventing unexpected crashes. |