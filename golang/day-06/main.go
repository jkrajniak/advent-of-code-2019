package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

const COMNode = "COM"

type NodeSet struct {
	nodeMap map[string]int64
	maxID int64
}

func NewNodeSet() *NodeSet {
	return &NodeSet{map[string]int64{}, 0}
}

func (n *NodeSet) GetID(label string) *int64 {
	if nodeID, hasID := n.nodeMap[label]; hasID {
		return &nodeID
	}
	return nil
}

func (n *NodeSet) NewID(label string) int64 {
	if nodeID, hasID := n.nodeMap[label]; hasID {
		return nodeID
	}
	nodeID := n.maxID + 1
	n.nodeMap[label] = nodeID
	n.maxID = nodeID
	return nodeID
}

var nodeMap = map[string]int64{}

type Node struct {
	Label string
	NodeID int64
}

func (n *Node) ID() int64 {
	return n.NodeID
}

func (n *Node) String() string {
	return fmt.Sprintf("%s (id=%d)", n.Label, n.NodeID)
}

func CalculateAllPathsLength(g graph.Graph, comNode graph.Node) int {
	count := 0
	shortest := path.DijkstraFrom(comNode, g)
	nodes := g.Nodes()
	for nodes.Next() {
		n := nodes.Node()
		if n.ID() == comNode.ID() {
			continue
		}
		pathToNode, _ := shortest.To(n.ID())
		count += len(pathToNode) - 1
	}
	return count
}

func GraphFromString(scanner *bufio.Scanner) (*simple.UndirectedGraph, graph.Node, *NodeSet) {
	diGraph := simple.NewUndirectedGraph()
	nodeSet := NewNodeSet()
	var comNode graph.Node
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), ")")
		labelA, labelB := strings.Trim(splitLine[0], " \n\t"), strings.Trim(splitLine[1], " \n\t")
		nodeIDA, nodeIDB := nodeSet.NewID(labelA), nodeSet.NewID(labelB)
		nodeA, nodeB := &Node{labelA, nodeIDA}, &Node{labelB, nodeIDB}
		if diGraph.Node(nodeIDA) == nil{
			diGraph.AddNode(nodeA)
		}
		if diGraph.Node(nodeIDB) == nil{
			diGraph.AddNode(nodeB)
		}
		if splitLine[0] == COMNode {
			comNode = nodeA
		} else if splitLine[1] == COMNode {
			comNode = nodeB
		}
		diGraph.SetEdge(diGraph.NewEdge(nodeA, nodeB))
	}
	return diGraph, comNode, nodeSet
}

func CalculateOrbitalTransfers(g *simple.UndirectedGraph, set *NodeSet) int {
	idNodeSAN := set.GetID("SAN")
	idNodeYOU := set.GetID("YOU")

	nodeYOU := g.Node(*idNodeYOU)

	shortest := path.DijkstraFrom(nodeYOU, g)
	pathToSAN, _ := shortest.To(*idNodeSAN)

	return len(pathToSAN) - 3
}



func main() {
	scanner := bufio.NewScanner(os.Stdin)
	diGraph, comNode, nodeSet := GraphFromString(scanner)

	pathLength := CalculateAllPathsLength(diGraph, comNode)
	fmt.Println(pathLength)

	orbitalTransfers := CalculateOrbitalTransfers(diGraph, nodeSet)
	fmt.Println(orbitalTransfers)
}

