package common

import tea "github.com/charmbracelet/bubbletea"

type EditFileMsg struct {
	FilePath string
}

func EditFile(fp string) tea.Cmd {
	return func() tea.Msg {
		return EditFileMsg{FilePath: fp}

	}
}
