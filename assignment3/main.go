package main

import (
	"fmt"
	"sort"
)

type Team struct {
	Name                                      string
	Wins, Draws, Losses, GoalFor, GoalAgainst int
}

type TeamList []Team

func (t TeamList) Len() int { return len(t) }

// swap the elements at indices i and j in the slice.
func (t TeamList) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

// Returns true if element at index i should come element at index j in sorted order, else false."
func (t TeamList) Less(i, j int) bool {
	// Calculate scores
	// Winning game will get 3 points, the draw game will get 1 point
	scoreI := (t[i].Wins * 3) + t[i].Draws
	scoreJ := (t[j].Wins * 3) + t[j].Draws

	// If scores are equal, use goal difference to rank up
	if scoreI == scoreJ {
		goalDiffI := t[i].GoalFor - t[i].GoalAgainst
		goalDiffJ := t[j].GoalFor - t[j].GoalAgainst

		// If goal differences are equal, use goal for to rank up
		if goalDiffI == goalDiffJ {
			return t[i].GoalFor > t[j].GoalFor
		}

		return goalDiffI > goalDiffJ
	}

	return scoreI > scoreJ
}

func main() {
	teams := TeamList{
		{"A", 5, 2, 1, 23, 20},
		{"B", 6, 1, 1, 24, 20},
		{"C", 7, 0, 1, 24, 10}, // Please note that goal for C is 24, not 25
		{"D", 6, 1, 1, 15, 11},
		{"E", 8, 1, 0, 23, 12},
	}

	// Sort the teams
	sort.Sort(teams)

	// fmt.Println(teams)

	fmt.Println("Team, Goal For, Goal Diff, Score")
	for _, team := range teams {
		score := (team.Wins * 3) + team.Draws
		goalDiff := team.GoalFor - team.GoalAgainst
		fmt.Printf("%s, %d, %d, %d\n", team.Name, team.GoalFor, goalDiff, score)
	}
}
