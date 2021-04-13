package parser

//go:generate pigeon -nolint -o parser.go tap.peg

import (
	"errors"
	"fmt"
)

func ParseString(input string) ([]Group, error) {
	ast, err := Parse("", []byte(input))
	if err != nil {
		return nil, err
	}

	var groups []Group = []Group{}
	for i, astInt := range ast.([]interface{}) {
		g, ok := astInt.(Group)
		if !ok {
			return nil, errors.New("unknown error")
		}
		if g.Name == "" {
			g.Name = fmt.Sprint(i)
		}
		groups = append(groups, g)
	}

	return groups, nil
}