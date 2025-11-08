package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"time"
)

var (
	appStyle  = lipgloss.NewStyle().Padding(1, 2)
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
)

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type model struct {
	playerList     list.Model
	selectedPlayer string
	status         string
	nowPlaying     string
	chosen         bool
}

type item struct {
	title string
	desc  string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func main() {
	err := tea.NewProgram(initialModel()).Start()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func initialModel() model {
	players, err := getPlayers()
	if err != nil {
		log.Fatal(err)
	}

	items := make([]list.Item, len(players))
	for i, p := range players {
		items[i] = item{title: p, desc: "Media Player"}
	}

	playerList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	playerList.Title = "Select a Player"

	return model{
		playerList: playerList,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.playerList.SetSize(msg.Width-h, msg.Height-v)

	case tickMsg:
		if m.chosen {
			m.nowPlaying, _ = getNowPlaying(m.selectedPlayer)
			m.status, _ = executePlayerctlCommand(m.selectedPlayer, "status")
			return m, tick()
		}
		return m, nil

	case tea.KeyMsg:
		if m.chosen {
			// Control View
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "b": // back to list
				m.chosen = false
				return m, nil
			case "p":
				m.status, _ = executePlayerctlCommand(m.selectedPlayer, "play-pause")
				m.nowPlaying, _ = getNowPlaying(m.selectedPlayer)
				go sendNotification("mpris-tui", fmt.Sprintf("Player %s: %s\n%s", m.selectedPlayer, "Play/Pause", m.nowPlaying))
				return m, nil
			case "s":
				m.status, _ = executePlayerctlCommand(m.selectedPlayer, "stop")
				m.nowPlaying, _ = getNowPlaying(m.selectedPlayer)
				go sendNotification("mpris-tui", fmt.Sprintf("Player %s: %s\n%s", m.selectedPlayer, "Stop", m.nowPlaying))
				return m, nil
			case "n":
				m.status, _ = executePlayerctlCommand(m.selectedPlayer, "next")
				m.nowPlaying, _ = getNowPlaying(m.selectedPlayer)
				go sendNotification("mpris-tui", fmt.Sprintf("Player %s: %s\n%s", m.selectedPlayer, "Next", m.nowPlaying))
				return m, nil
			case "v": // previous
				m.status, _ = executePlayerctlCommand(m.selectedPlayer, "previous")
				m.nowPlaying, _ = getNowPlaying(m.selectedPlayer)
				go sendNotification("mpris-tui", fmt.Sprintf("Player %s: %s\n%s", m.selectedPlayer, "Previous", m.nowPlaying))
				return m, nil
			}
		} else {
			// Player List View
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "enter":
				selectedItem, ok := m.playerList.SelectedItem().(item)
				if ok {
					m.selectedPlayer = selectedItem.title
					m.chosen = true
					m.status, _ = executePlayerctlCommand(m.selectedPlayer, "status")
					m.nowPlaying, _ = getNowPlaying(m.selectedPlayer)
				}
				return m, tick()
			}
		}
	}

	var cmd tea.Cmd
	m.playerList, cmd = m.playerList.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.chosen {
		// Control View
		return appStyle.Render(fmt.Sprintf(
			"Controlling: %s\nStatus: %s\nNow Playing: %s\n\n[p] Play/Pause\n[s] Stop\n[n] Next\n[v] Previous\n\n%s",
			m.selectedPlayer,
			m.status,
			m.nowPlaying,
			helpStyle("Press 'b' to go back to the player list. Press 'q' to quit."),
		))
	}
	
	// Player List View
	if len(m.playerList.Items()) == 0 {
		return appStyle.Render("No media players found.\nPlease start a player and restart the application.\n\nPress 'q' to quit.")
	}
	return appStyle.Render(m.playerList.View())
}

func getPlayers() ([]string, error) {
	cmd := exec.Command("playerctl", "-l")
	out, err := cmd.Output()
	if err != nil {
		// If playerctl is not found or no players are running, return an empty list
		return []string{}, nil
	}
	players := strings.TrimSpace(string(out))
	if len(players) == 0 {
		return []string{}, nil
	}
	return strings.Split(players, "\n"), nil
}

func executePlayerctlCommand(player, command string) (string, error) {
	cmd := exec.Command("playerctl", "-p", player, command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error: %s", string(out)), err
	}
	
	if command == "status" {
		return strings.TrimSpace(string(out)), nil
	}

	// After a command, get the new status
	statusCmd := exec.Command("playerctl", "-p", player, "status")
	statusOut, statusErr := statusCmd.Output()
	if statusErr != nil {
		return "Status unavailable", statusErr
	}
	return strings.TrimSpace(string(statusOut)), nil
}

func getNowPlaying(player string) (string, error) {
	cmd := exec.Command("playerctl", "-p", player, "metadata", "--format", "{{artist}} - {{title}}")
	out, err := cmd.CombinedOutput()
	if err != nil {
		// If there's an error (e.g., no metadata), return an empty string
		return "No media playing", err
	}
	return strings.TrimSpace(string(out)), nil
}

func sendNotification(summary, body string) {
	exec.Command("notify-send", summary, body).Run()
}
