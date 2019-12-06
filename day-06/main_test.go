package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAllPathsLength(t *testing.T) {
	s := `COM)B
		B)C
		C)D
		D)E
		E)F
		B)G
		G)H
		D)I
		E)J
		J)K
		K)L`

	scanner := bufio.NewScanner(strings.NewReader(s))
	diGraph, comNode, _ := GraphFromString(scanner)
	edges := diGraph.Edges()
	for edges.Next() {
		e := edges.Edge()
		fmt.Printf("%s - %s\n", e.From(), e.To())
	}
	length := CalculateAllPathsLength(diGraph, comNode)
	assert.Equal(t, 42, length)
}

func TestOrbitTransfer(t *testing.T) {
	m := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`

	scanner := bufio.NewScanner(strings.NewReader(m))
	g, _, nodeSet := GraphFromString(scanner)
	edges := g.Edges()
	for edges.Next() {
		e := edges.Edge()
		fmt.Printf("%s - %s\n", e.From(), e.To())
	}
	length := CalculateOrbitalTransfers(g, nodeSet)
	assert.Equal(t, 4, length)
}