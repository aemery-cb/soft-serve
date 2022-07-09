package hooks

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/soft-serve/internal/tui/style"
	"github.com/charmbracelet/soft-serve/tui/common"
)

const hooksFolder = "hooks"

type hooksState int

const (
	listState = iota
	newHookState
)

type Bubble struct {
	state     hooksState
	repo      common.GitRepo
	style     *style.Styles
	width     int
	height    int
	path      string
	selection string
	cursor    int
	textInput textinput.Model
	hooks     []string
}

func NewBubble(repo common.GitRepo, styles *style.Styles, width, height int) *Bubble {
	path := filepath.Join(repo.Path(), hooksFolder)
	os.MkdirAll(path, os.ModePerm)

	ti := textinput.New()
	ti.Placeholder = "hook name"
	ti.CharLimit = 156
	ti.Width = 20

	b := &Bubble{
		repo:      repo,
		style:     styles,
		width:     width,
		height:    height,
		path:      path,
		state:     listState,
		textInput: ti,
	}

	b.SetSize(width, height)
	b.reset()
	return b
}

func (b *Bubble) Init() tea.Cmd {
	return nil
}

func (b *Bubble) SetSize(width, height int) {
	b.width = width
	b.height = height
}

func (b *Bubble) Help() []common.HelpEntry {
	return nil
}

func (b *Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			if b.cursor < len(b.hooks) {
				b.cursor += 1
			}
		case "k":
			if b.cursor > 0 {
				b.cursor -= 1
			}
		case "n":
			b.state = newHookState
			b.textInput.Focus()
			b.textInput.Blink()
			return b, nil
		case "enter":
			switch b.state {
			case listState:
				b.selection = b.hooks[b.cursor]
				return b, common.EditFile(filepath.Join(b.path, b.selection))
			case newHookState:
				b.selection = b.textInput.Value()
				fullpath := filepath.Join(b.path, b.selection)
				createHook(fullpath)
				cmds = append(cmds, common.EditFile(fullpath))
				b.reset()
				b.state = listState
				return b, tea.Batch(cmds...)
			}
		}
	}
	if b.state == newHookState {
		textInput, cmd := b.textInput.Update(msg)
		b.textInput = textInput
		cmds = append(cmds, cmd)
	}
	return b, tea.Batch(cmds...)
}

func (b *Bubble) View() string {
	switch b.state {
	case listState:

		result := "press n to create a new hook\n"
		for i, file := range b.hooks {
			if i == b.cursor {
				result += ">" + file + "\n"
			} else {
				result += file + "\n"
			}
		}
		return result
	case newHookState:
		return b.textInput.View()
	}
	return ""
}

func (b *Bubble) reset() {
	b.hooks = nil
	if files, err := os.ReadDir(b.path); err == nil {
		for _, file := range files {
			b.hooks = append(b.hooks, file.Name())
		}
	}
}
func createHook(filename string) error {
	_, err := os.Create(filename)
	if err != nil {
		return err
	}

	err = os.Chmod(filename, 0777)
	if err != nil {
		return err
	}
	return nil

}
