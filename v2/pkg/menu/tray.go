package menu

// Tray are the options
type Tray struct {

	// The ID of this tray
	ID string

	// Label is the text we wish to display in the tray
	Label string

	// Icon is the name of the tray icon we wish to display.
	// These are read up during build from <projectdir>/trayicons and
	// the filenames are used as IDs, minus the extension
	// EG: <projectdir>/trayicons/main.png can be referenced here with "main"
	Icon string

	// Menu is the initial menu we wish to use for the tray
	Menu *Menu
}