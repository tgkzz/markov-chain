package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	handlerFuncs "markov-chain/handler"
)

type command int

const (
	baseline command = iota + 1
	wordNum
	prefix
	prefixLen
	usage
)

type handler interface {
	HandleFunc(ctx context.Context) error
}

func GetHandler(cmdNum command) (handler, error) {
	switch cmdNum {
	case usage:
		return handlerFuncs.NewUsageHandler(), nil
	case baseline:
		return handlerFuncs.NewBaselineHandler(), nil
	default:
		return nil, ErrNotImplemented
	}
}

type app struct {
	args []string
}

type appRunner interface {
	Run(ctx context.Context) error
}

func NewApp(args []string) appRunner {
	return &app{
		args: args,
	}
}

var (
	ErrNotImplemented = errors.New("method not implemented or never exist")
	ErrFlagNotExist   = errors.New("flag not exist")
)

func (a *app) Run(ctx context.Context) error {
	// listen from context deadline
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:

	}

	// gateway internal func to detect logic
	cmds, err := a.gateway()
	if err != nil {
		return err
	}

	// check command by switch case
	// run them in each handler
	for _, cmd := range cmds {
		var h handler
		h, err = GetHandler(cmd)
		if err != nil {
			return err
		}
		if err = h.HandleFunc(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *app) gateway() ([]command, error) {
	fl := flag.NewFlagSet("try to use --help or -h command", flag.ExitOnError)

	wNum := fl.Int("word-num", 0, "Number of words to generate")

	prLen := fl.Int("prefix-len", 0, "Length of prefix to generate")

	pr := fl.String("prefix", "", "Prefix to generate")

	// usage of help handler seems strange when flag is not correct,
	// however by flag lib creators, we need to print usage as well
	// see parse func implementation
	fl.Usage = func() {
		h, _ := GetHandler(usage)
		if err := h.HandleFunc(context.TODO()); err != nil {
			return
		}
	}

	if err := fl.Parse(a.args); err != nil {
		return nil, ErrFlagNotExist
	}

	cmds := make([]command, 0, len(flag.Args()))
	if wnum := *wNum; wnum > 0 {
		cmds = append(cmds, wordNum)
	}
	if prlen := *prLen; prlen > 0 {
		cmds = append(cmds, prefixLen)
	}
	if pr != nil && len(*pr) > 0 {
		cmds = append(cmds, prefix)
	}

	if len(cmds) == 0 {
		cmds = append(cmds, baseline)
	}

	return cmds, nil
}

func main() {
	ctx := context.Background()

	args := os.Args[1:]

	a := NewApp(args)

	if err := a.Run(ctx); err != nil {
		if errors.Is(err, ErrFlagNotExist) {
			os.Exit(0)
		}
		fmt.Println(err)
		os.Exit(1) // maybe we should not exit program with exit status 1
	}
}
