package markov

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

type Chain struct {
	prefixLen int
	chain     map[string][]string
}

func NewChain(prefixLen int) *Chain {
	return &Chain{
		prefixLen: prefixLen,
		chain:     make(map[string][]string),
	}
}

func (mc *Chain) Build(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
	}
	if len(words) < mc.prefixLen {
		return fmt.Errorf("not enough words in the input")
	}

	for i := 0; i <= len(words)-mc.prefixLen; i++ {
		prefix := strings.Join(words[i:i+mc.prefixLen], " ")
		if i+mc.prefixLen < len(words) {
			suffix := words[i+mc.prefixLen]
			mc.chain[prefix] = append(mc.chain[prefix], suffix)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (mc *Chain) Generate(maxWords int) string {
	var prefixes []string
	for prefix := range mc.chain {
		prefixes = append(prefixes, prefix)
	}

	if len(prefixes) == 0 {
		return ""
	}

	prefix := prefixes[0]
	words := strings.Split(prefix, " ")

	for len(words) < maxWords {
		suffixes, ok := mc.chain[prefix]
		if !ok || len(suffixes) == 0 {
			break
		}

		next := suffixes[rand.Intn(len(suffixes))]
		words = append(words, next)

		prefixWords := words[len(words)-mc.prefixLen:]
		prefix = strings.Join(prefixWords, " ")
	}

	return strings.Join(words, " ")
}
