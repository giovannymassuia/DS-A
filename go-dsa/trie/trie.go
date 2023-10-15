package trie

import (
	"errors"
	"strings"
)

type TrieNode struct {
	StaticChildren map[string]*TrieNode
	ParameterChild *TrieNode
	Handler        func() bool
	ParameterName  string // Changed from pointer to string
	HasParam       bool   // New field to indicate if ParameterName is set
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: newTrieNode(),
	}
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		StaticChildren: make(map[string]*TrieNode),
	}
}

func (t *Trie) Insert(path string, handler func() bool) error {
	current := t.Root
	segments := splitPath(path)
	seenParameters := make(map[string]bool)

	for _, segment := range segments {
		if segment[0] == ':' {
			paramName := segment[1:]

			if seenParameters[paramName] {
				return errors.New("cannot use the same parameter name more than once in a route")
			}
			seenParameters[paramName] = true

			if current.ParameterChild != nil {
				if current.ParameterChild.ParameterName != paramName {
					return errors.New("Invalid path: " + path + ". Parameter " + paramName + " is already defined.")
				}
			} else {
				paramNode := newTrieNode()
				paramNode.ParameterName = paramName
				paramNode.HasParam = true
				current.ParameterChild = paramNode
			}

			current = current.ParameterChild
		} else {
			if _, exists := current.StaticChildren[segment]; !exists {
				current.StaticChildren[segment] = newTrieNode()
			}

			current = current.StaticChildren[segment]
		}
	}

	current.Handler = handler

	return nil
}

type SearchResult struct {
	Handler func() bool
	Params  map[string]string
}

func (t *Trie) Search(path string) (result SearchResult) {
	current := t.Root
	segments := splitPath(path)
	capturedParams := make(map[string]string)

	for _, segment := range segments {
		next, exists := current.StaticChildren[segment]

		if !exists && current.ParameterChild != nil && current.ParameterChild.HasParam {
			next = current.ParameterChild
			capturedParams[current.ParameterChild.ParameterName] = segment
		}

		if next == nil {
			return SearchResult{Handler: func() bool { return false }, Params: nil}
		}

		current = next
	}

	return SearchResult{Handler: current.Handler, Params: capturedParams}
}

func splitPath(path string) []string {
	return strings.Split(path, "/")[1:]
}
