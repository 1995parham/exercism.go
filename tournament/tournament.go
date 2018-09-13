/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 25-08-2018
 * |
 * | File Name:     tournament.go
 * +===============================================
 */

package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Team represents a team in tournament
type Team struct {
	Name string

	Win  int
	Draw int
	Loss int

	Points int
	Games  int
}

// NewTeam creates new team and calculate its points and games
func NewTeam(name string, w int, d int, l int) Team {
	return Team{
		Name: name,

		Win:  w,
		Draw: d,
		Loss: l,

		Points: w*3 + d,
		Games:  w + d + l,
	}
}

// Tally creates matches table based on given report
func Tally(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	matches := make(map[string]string)

	for scanner.Scan() {
		l := scanner.Text()

		// ignore new lines and comments
		if l == "" || strings.HasPrefix(l, "#") {
			continue
		}

		m := strings.Split(l, ";")
		if len(m) != 3 {
			return fmt.Errorf("Invalid input: %s", l)
		}

		switch m[2] {
		case "win":
			matches[m[0]] += "W"
			matches[m[1]] += "L"
		case "loss":
			matches[m[0]] += "L"
			matches[m[1]] += "W"
		case "draw":
			matches[m[0]] += "D"
			matches[m[1]] += "D"
		default:
			return fmt.Errorf("Invalid input: %s", l)
		}
	}
	if len(matches) == 0 {
		return fmt.Errorf("At least one team must exist")
	}

	teams := make([]Team, len(matches))
	i := 0
	for name, score := range matches {
		teams[i] = NewTeam(name, strings.Count(score, "W"), strings.Count(score, "D"), strings.Count(score, "L"))
		i++
	}

	sort.Slice(teams, func(i, j int) bool {
		if teams[i].Points != teams[j].Points {
			return teams[i].Points > teams[j].Points
		}
		return teams[i].Name < teams[j].Name
	})

	if _, err := io.WriteString(w, "Team                           | MP |  W |  D |  L |  P\n"); err != nil {
		return err
	}
	for _, team := range teams {
		if _, err := io.WriteString(w, fmt.Sprintf("%-30s |  %d |  %d |  %d |  %d |  %d\n", team.Name, team.Games, team.Win, team.Draw, team.Loss, team.Points)); err != nil {
			return err
		}
	}

	return nil
}
