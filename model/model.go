package model

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"os"
	"strings"
	"time"

	"imz/flip"
	"imz/grayscale"
	"imz/load"
	"imz/resize"
	"imz/rotate"
	"imz/saveImz"
	"imz/stack"
	"imz/temporary"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type imgName string

type options int

type Model struct {
	img         imgName
	option      []string
	cursor      int
	functionMap map[options]any
	selected    map[options]struct{}

	Filepicker   filepicker.Model
	selectedFile string
	chosen       bool
	err          error

	imgNRGBA *image.NRGBA

	dotCreated bool
	undoCount  int
}

const (
	Flip_Vertically = iota
	Flip_Horizontally
	Rotate
	Resize
	Grayscale
)

var resFile string = "/home/gafoli/Pictures/letsseescale5.png"

var style1 = lipgloss.NewStyle().
	Bold(true).
	Italic(true).
	Blink(true).
	Foreground(lipgloss.Color("86"))

var style2 = lipgloss.NewStyle().
	Bold(true).
	Italic(true).
	Faint(true).
	Blink(true).
	Foreground(lipgloss.Color("140"))

var style3 = lipgloss.NewStyle().
	Italic(true).
	Blink(true).
	Foreground(lipgloss.Color("198"))

var style4 = lipgloss.NewStyle().
	Italic(true).
	Blink(true).
	Foreground(lipgloss.Color("#5C5C5C"))

func NewModel(img imgName) *Model {
	fm := make(map[options]any)

	m := &Model{
		img:         img,
		option:      []string{"Flip Vertically", "Flip Horizontally", "Rotate", "Resize", "Grayscale"},
		cursor:      -1,
		selected:    make(map[options]struct{}),
		functionMap: fm,
		chosen:      false,
	}

	m.newImgFunction(options(0), flip.VFlip)
	m.newImgFunction(options(1), flip.HFlip)
	m.newImgFunction(options(2), rotate.RotateImg)
	m.newImgFunction(options(3), resize.NNI)
	m.newImgFunction(options(4), grayscale.GrayImg)

	return m
}

func (m *Model) newImgFunction(opt options, function any) {
	m.functionMap[opt] = function
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (m Model) Init() tea.Cmd {
	return m.Filepicker.Init()
}

var s = stack.NewStack()
var is = stack.NewImgStack()

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	m.Filepicker, cmd = m.Filepicker.Update(msg)

	if didSelect, path := m.Filepicker.DidSelectFile(msg); didSelect {
		m.selectedFile = path
		m.chosen = true
	}

	if didSelect, path := m.Filepicker.DidSelectDisabledFile(msg); didSelect {
		m.err = errors.New(path + " is not valid.")
		m.selectedFile = ""
		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:

		if len(s.History) != len(is.ImgHistory) {
			panic("history not equal to imgHistory")
		}

		if m.chosen {

			if !m.dotCreated {
				m.imgNRGBA = load.Load(m.selectedFile)
			} else if !is.IsEmpty() {
				mostRecentHistory := is.Peek()
				m.imgNRGBA = load.Load(mostRecentHistory.Name())
			}

			switch msg.String() {

			case "ctrl+c", "q":
				if !is.IsEmpty() {
					source := is.Peek()
					time := time.Now().UnixMicro()
					dest := fmt.Sprintf("/home/gafoli/Pictures/%v.png", time)

					err := os.Rename(source.Name(), dest)
					if err != nil {
						panic("error saving file")
					}
				}
				for !is.IsEmpty() {
					allHistoryImg := is.Peek()
					is.Pop()
					os.Remove(allHistoryImg.Name())
				}
				return m, tea.Quit

			case "ctrl+z":
				if !is.IsEmpty() {
					undoImg := is.Peek()
					is.Pop()
					s.Pop()
					os.Remove(undoImg.Name())
				}

			case "up", "k":
				if m.cursor == 0 {
					m.cursor = len(m.option) - 1
				} else if m.cursor > 0 {
					m.cursor--
				}

			case "down", "j":
				if m.cursor == len(m.option)-1 {
					m.cursor = 0
				} else if m.cursor < len(m.option)-1 {
					m.cursor++
				}

			case " ":
				_, ok := m.selected[options(m.cursor)]
				if ok {
					m.selected = make(map[options]struct{})
				} else if m.cursor >= 0 {
					m.selected = make(map[options]struct{})
					m.selected[options(m.cursor)] = struct{}{}
					s.Push(m.cursor)

					switch m.cursor {
					case Resize:
						imgFunc, ok := m.functionMap[options(m.cursor)].(func(image.Image, float64) [][]color.Color)
						if !ok {
							panic("cannot find func: resize")
						}

						temporary.CreateDir()
						historyFile := temporary.CreateTempFile(&m.undoCount)
						defer historyFile.Close()

						resizedGrid := imgFunc(m.imgNRGBA, 5)

						resPath := historyFile.Name()
						saveImz.SaveRectYX(resPath, resizedGrid)
						is.Push(historyFile)

						m.dotCreated = true

					case Flip_Vertically, Grayscale:
						imgFunc, ok := m.functionMap[options(m.cursor)].(func(image.Image) [][]color.Color)
						if !ok {
							panic("cannot find func: flipV or grayscale")
						}

						temporary.CreateDir()
						historyFile := temporary.CreateTempFile(&m.undoCount)
						defer historyFile.Close()

						modifiedGrid := imgFunc(m.imgNRGBA)

						resPath := historyFile.Name()
						saveImz.SaveRectXY(resPath, modifiedGrid)

						is.Push(historyFile)

						m.dotCreated = true

					case Flip_Horizontally, Rotate:
						imgFunc, ok := m.functionMap[options(m.cursor)].(func(image.Image) [][]color.Color)
						if !ok {
							l := fmt.Sprintf("cannot find func: flipH or rotate: %v", m.cursor)
							panic(l)
						}

						temporary.CreateDir()
						historyFile := temporary.CreateTempFile(&m.undoCount)
						defer historyFile.Close()

						modifiedGrid := imgFunc(m.imgNRGBA)

						resPath := historyFile.Name()
						saveImz.SaveRectYX(resPath, modifiedGrid)

						is.Push(historyFile)

						m.dotCreated = true
					}
				}
			}
		} else if !m.chosen {
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			}
		}

	case clearErrorMsg:
		m.err = nil
	}
	return m, nil
}

func (m Model) View() string {

	colors := []string{"81", "86", "92", "202", "208", "213", "226"}

	logoLines := []string{
		"┌────────────────────────────────────────────────────────────────────────────────┐",
		"│ ___  _____ ______   ________  _________  _______   ________  _____ ______      │",
		"│|\\  \\|\\   _ \\  _   \\|\\_____  \\|\\___   ___\\\\  ___ \\ |\\   __  \\|\\   _ \\  _   \\    │",
		"│\\ \\  \\ \\  \\\\\\__\\ \\  \\\\|___/  /\\|___ \\  \\_\\ \\   __/|\\ \\  \\|\\  \\ \\  \\\\\\__\\ \\  \\   │",
		"│ \\ \\  \\ \\  \\\\|__| \\  \\   /  / /    \\ \\  \\ \\ \\  \\_|/_\\ \\   _  _\\ \\  \\\\|__| \\  \\  │",
		"│  \\ \\  \\ \\  \\    \\ \\  \\ /  /_/__    \\ \\  \\ \\ \\  \\_|\\ \\ \\  \\\\  \\\\ \\  \\    \\ \\  \\ │",
		"│   \\ \\__\\ \\__\\    \\ \\__\\\\________\\   \\ \\__\\ \\ \\_______\\ \\__\\\\ _\\\\ \\__\\    \\ \\__\\│",
		"│    \\|__|\\|__|     \\|__|\\|_______|    \\|__|  \\|_______|\\|__|\\|__|\\|__|     \\|__|│",
		"└────────────────────────────────────────────────────────────────────────────────┘",
	}

	var styledLogo strings.Builder
	backgroundColor := "235"
	for i, line := range logoLines {
		color := colors[i%len(colors)]
		style := lipgloss.NewStyle().
			Bold(true).
			Faint(false).
			Italic(false).
			Foreground(lipgloss.Color(color)).
			Background(lipgloss.Color(backgroundColor))

		styledLogo.WriteString(style.Render(line) + "\n")
	}

	if !m.chosen {
		var s strings.Builder

		if m.err != nil {
			s.WriteString(m.Filepicker.Styles.DisabledFile.Render(m.err.Error()))
		} else if m.selectedFile == "" {
			s.WriteString("\nPick an Image")
		}

		s.WriteString("\n" + m.Filepicker.View())

		return s.String()
	}

	var outLine strings.Builder
	var sl string
	for i, choice := range m.option {
		arrow := " "
		if m.cursor == i {
			arrow = ">"
		}

		checked := " "
		if _, ok := m.selected[options(i)]; ok {
			checked = "x"
		}

		line := fmt.Sprintf("\n%s [%s] %v\n", arrow, checked, choice)
		if checked == "x" {
			outLine.WriteString(style1.Render(line) + "\n")
		} else {
			outLine.WriteString(style2.Render(line) + "\n")
		}

	}

	if !s.IsEmpty() {
		sl = "History =>"
	}

    instruction := "\nSpace = Select, Undo = Ctrl+z, Quit and Save = q/Ctrl+c"


	for i := range len(s.History) {
		sl += fmt.Sprintf(" > %v ", m.option[s.History[i]])
	}
	return styledLogo.String() + outLine.String() + style3.Render(sl) + style4.Render(instruction)
}
