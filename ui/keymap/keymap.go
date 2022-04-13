package keymap

import "github.com/charmbracelet/bubbles/key"

// KeyMap is a map of key bindings for the UI.
type KeyMap struct {
	Quit      key.Binding
	Up        key.Binding
	Down      key.Binding
	UpDown    key.Binding
	LeftRight key.Binding
	Arrows    key.Binding
	Select    key.Binding
}

// DefaultKeyMap returns the default key map.
func DefaultKeyMap() *KeyMap {
	km := new(KeyMap)

	km.Quit = key.NewBinding(
		key.WithKeys(
			"ctrl-c",
			"q",
		),
		key.WithHelp(
			"q",
			"quit",
		),
	)

	km.Up = key.NewBinding(
		key.WithKeys(
			"up",
			"k",
		),
		key.WithHelp(
			"↑",
			"up",
		),
	)

	km.Down = key.NewBinding(
		key.WithKeys(
			"down",
			"j",
		),
		key.WithHelp(
			"↓",
			"down",
		),
	)

	km.UpDown = key.NewBinding(
		key.WithKeys(
			"up",
			"down",
			"k",
			"j",
		),
		key.WithHelp(
			"↑↓",
			"navigate",
		),
	)

	km.LeftRight = key.NewBinding(
		key.WithKeys(
			"left",
			"h",
			"right",
			"l",
		),
		key.WithHelp(
			"←→",
			"navigate",
		),
	)

	km.Arrows = key.NewBinding(
		key.WithKeys(
			"up",
			"right",
			"down",
			"left",
			"k",
			"j",
			"h",
			"l",
		),
		key.WithHelp(
			"↑←↓→",
			"navigate",
		),
	)

	km.Select = key.NewBinding(
		key.WithKeys(
			"enter",
		),
		key.WithHelp(
			"enter",
			"select",
		),
	)

	return km
}