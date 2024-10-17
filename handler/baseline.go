package handler

import (
	"context"
	"fmt"
	"os"

	"markov-chain/pkg/markov"
)

type BaselineHandler struct {
	MaxWords int
}

// must be taken from params or use default
func NewBaselineHandler() *BaselineHandler {
	return &BaselineHandler{
		MaxWords: 100,
	}
}

func (h *BaselineHandler) HandleFunc(ctx context.Context) error {
	chain := markov.NewChain(2)

	err := chain.Build(os.Stdin)
	if err != nil {
		if _, panicErr := fmt.Fprintf(os.Stderr, "Error: %v\n", err); err != nil {
			// why we could have such error check, but ok
			panic(fmt.Errorf("error while writing to os.Stderr %v", panicErr))
		}
		return err
	}

	result := chain.Generate(h.MaxWords)

	if len(result) == 0 {
		fmt.Println("Error: no input text")
		return nil
	}

	fmt.Println(result)
	return nil
}
