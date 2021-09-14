package tournament

import (
	"errors"
	"io"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type stats struct {
	team   string
	played int
	won    int
	drawn  int
	lost   int
	points int
}

func Tally(r io.Reader, w io.Writer) error {
	// read input
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	text := string(input)
	text = strings.TrimSpace(text)

	// processPoints team results
	results, err := processPoints(text)
	if err != nil {
		return err
	}

	// order teams
	// first by points, then alphabetical
	teams := make([]string, 0)
	for k, _ := range results {
		teams = append(teams, k)
	}
	sort.Slice(teams, func(i, j int) bool {
		if results[teams[i]].points == results[teams[j]].points {
			return results[teams[i]].team < results[teams[j]].team
		}
		return results[teams[i]].points > results[teams[j]].points
	})

	// format response
	output := buildTable(teams, results)

	// write output to writer
	_, err = w.Write([]byte(output))
	if err != nil {
		return err
	}

	return nil
}

// buildTable returns the output string in the required format
func buildTable(teams []string, results map[string]*stats) string {
	// format rows for each team
	formatted := make(map[string]string)
	for k, v := range results {
		line := k + strings.Repeat(" ", 31-len(k))
		line += "|  " + strconv.Itoa(v.played) + " |  " + strconv.Itoa(v.won)
		line += " |  " + strconv.Itoa(v.drawn) + " |  " + strconv.Itoa(v.lost)
		line += " |  " + strconv.Itoa(v.points) + "\n"
		formatted[k] = line
	}

	// prepare output to return
	output := "Team                           | MP |  W |  D |  L |  P\n"
	for _, t := range teams {
		output += formatted[t]
	}
	return output
}

// processPoints receives the matches in string format and process them in a map
// keys are team names
func processPoints(text string) (map[string]*stats, error) {
	results := make(map[string]*stats)
	rows := strings.Split(text, "\n")
	for _, row := range rows {
		if row == "" {
			continue
		}
		if string(row[0]) == "#" {
			continue
		}
		columns := strings.Split(row, ";")
		if len(columns) != 3 {
			return nil, errors.New("invalid input")
		}
		if _, ok := results[columns[0]]; !ok {
			results[columns[0]] = &stats{team: columns[0]}
		}
		if _, ok := results[columns[1]]; !ok {
			results[columns[1]] = &stats{team: columns[1]}
		}
		team1 := results[columns[0]]
		team2 := results[columns[1]]
		team1.played++
		team2.played++
		switch columns[2] {
		case "win":
			team1.won++
			team2.lost++
			team1.points += 3
		case "loss":
			team2.won++
			team1.lost++
			team2.points += 3
		case "draw":
			team1.drawn++
			team2.drawn++
			team1.points += 1
			team2.points += 1
		default:
			return nil, errors.New("invalid input")
		}
	}
	return results, nil
}
