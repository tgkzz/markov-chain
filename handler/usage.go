package handler

import (
	"context"
	"fmt"
	"time"
)

type UsageHandler struct {
	// need to use pkg markov
}

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
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	fmt.Println(HelpMsg)

	return nil
}
