package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("201"))
	success       = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
)

type model struct {
	choices      []string
	cursor       int
	selected     map[int]struct{}
	taskComplete bool
	lastMessages []string
}

func initialModel() model {
	return model{
		choices:  []string{"Delete files in Prefetch", "Delete files in Temp", "Delete files in %temp%", "Run Disk Cleanup", "Run chkdsk /f"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Go Clean Windows")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.taskComplete {
			switch msg.String() {
			case "enter":
				m.taskComplete = false
				m.lastMessages = nil
				return m, nil
			case "q":
				return m, tea.Quit
			}
		} else {
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit

			case "up", "w":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "s":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}

			case " ":
				_, ok := m.selected[m.cursor]

				if m.cursor == 3 && m.isSelected(4) || m.cursor == 4 && m.isSelected(3) {
					break
				}

				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}

			case "enter":
				var messages []string
				for index := range m.selected {
					switch index {
					case 0:
						messages = append(messages, cleanFiles(os.Getenv("SystemRoot")+"\\Prefetch"))
					case 1:
						messages = append(messages, cleanFiles(os.Getenv("SystemRoot")+"\\Temp"))
					case 2:
						messages = append(messages, cleanFiles(os.Getenv("TEMP")))
					case 3:
						messages = append(messages, cleanDisk())
					case 4:
						messages = append(messages, chkdskFix())
					}
				}
				m.taskComplete = true
				m.selected = make(map[int]struct{})
				m.lastMessages = messages
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.taskComplete {
		return fmt.Sprintf("%s\n\n%s", success.Render(strings.Join(m.lastMessages, "\n")), subtleStyle.Render("enter: Go Back • q: Quit\n"))
	}

	s := "\n  Select clean windows option:\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = checkboxStyle.Render(">")
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = checkboxStyle.Render("x")
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += subtleStyle.Render("  Info: Disk Cleanup and chkdsk /f cannot be selected together.\n")

	s += subtleStyle.Render("\n  w/s, ↑/↓: Navigate • space: Select • enter: Choose • q: Quit\n")

	return s
}

func (m model) isSelected(index int) bool {
	_, ok := m.selected[index]
	return ok
}

func cleanFiles(path string) string {
	files, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		return fmt.Sprintf("failed to read directory: %v", err)
	}
	for _, file := range files {
		if err := os.RemoveAll(file); err != nil {
			continue
		}
	}
	return fmt.Sprintf("successfully cleaned %s", path)
}

func cleanDisk() string {
	cmd := exec.Command("cleanmgr", "/D", "C")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Sprintf("failed to run disk cleanup: %v", err)
	}

	return "disk cleanup completed successfully on drive C"
}

func chkdskFix() string {
	cmd := exec.Command("cmd", "/C", "echo y | chkdsk /f")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Sprintf("failed to run chkdsk /f: %v", err)
	}

	return "chkdsk /f executed successfully"
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
