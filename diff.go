package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	ANSI_YELLOW      = "\x1b[33m"
	ANSI_LIGHT_BLUE  = "\x1b[94m"
	ANSI_LIGHT_GREEN = "\x1b[92m"
	ANSI_RESET       = "\x1b[0m"
)

type Commit struct {
	Author struct {
		Name string `json:"login"`
	} `json:"author"`
	Sha    string `json:"sha"`
	Commit struct {
		Message   string `json:"message"`
		Committer struct {
			ComittedAt time.Time `json:"date"`
		} `json:"committer"`
	} `json:"commit"`
}

type Diff struct {
	AheadBy  int       `json:"ahead_by"`
	BehindBy int       `json:"behind_by"`
	Commits  []*Commit `json:"commits"`
}

func printFormattedDiff(diff *Diff) {
	fmt.Printf("Ahead: %d Commits, Behind: %d Commits\n", diff.AheadBy, diff.BehindBy)

	if len(diff.Commits) > 0 {
		fmt.Printf("Commits that will be deployed:\n")
		for _, commit := range diff.Commits {
			printFormattedCommit(commit)
		}
	}
}

func printFormattedCommit(commit *Commit) {
	message := formatCommitMessage(commit.Commit.Message)
	formattedTimestamp := commit.Commit.Committer.ComittedAt.Format(time.RFC822)

	fmt.Printf("* %s%s%s   %s%s%s   %s%-15s%s %s\n",
		ANSI_YELLOW, commit.Sha[:6], ANSI_RESET,
		ANSI_LIGHT_GREEN, formattedTimestamp, ANSI_RESET,
		ANSI_LIGHT_BLUE, commit.Author.Name, ANSI_RESET,
		message)
}

func formatCommitMessage(message string) string {
	lines := strings.Split(message, "\n")
	return lines[0]
}
