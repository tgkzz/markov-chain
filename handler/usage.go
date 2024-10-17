package handler

import (
	"context"
	"fmt"
)

type UsageHandler struct{}

func NewUsageHandler() *UsageHandler {
	return &UsageHandler{}
}

const HelpMsg = `Markov Chain text generator.

Usage:
  markovchain [-w <N>] [-p <S>] [-l <N>]
  markovchain --help

Options:
  --help  Show this screen.
  -w N    Number of maximum words
  -p S    Starting prefix
  -l N    Prefix length`

func (u *UsageHandler) HandleFunc(ctx context.Context) error {
	fmt.Println(HelpMsg)

	return nil
}
