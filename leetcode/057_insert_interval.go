package main

import "fmt"

type Interval struct {
	Start int
	End   int
}

type edge int

const (
	edgeLeft edge = iota
	edgeRight
)

type node struct {
	data int
	edge edge
}

type nodes []node

func (n nodes) Len() int           { return len(n) }
func (n nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n nodes) Less(i, j int) bool { return n[i].data < n[j].data }

func insertNode(nodes nodes, node node) nodes {
	nodes = append(nodes, node)
	for i := len(nodes) - 1; i > 0; i-- {
		if nodes[i].data >= nodes[i-1].data {
			return nodes
		}
		nodes[i], nodes[i-1] = nodes[i-1], nodes[i]
	}
	return nodes
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	var nodes nodes = make([]node, 0, len(intervals)*2+2)
	for _, i := range intervals {
		nodes = append(nodes, node{
			data: i.Start,
			edge: edgeLeft,
		})
		nodes = append(nodes, node{
			data: i.End,
			edge: edgeRight,
		})
	}
	nodes = insertNode(nodes, node{
		data: newInterval.Start,
		edge: edgeLeft,
	})
	nodes = insertNode(nodes, node{
		data: newInterval.End,
		edge: edgeRight,
	})
	var ret []Interval
	leftNum := 0
	var leftNode node
	for _, n := range nodes {
		if leftNum == 0 && n.edge == edgeLeft {
			leftNode = n
		}
		switch n.edge {
		case edgeLeft:
			leftNum++
		case edgeRight:
			leftNum--
		}
		if leftNum == 0 && n.edge == edgeRight {
			if l := len(ret); l > 0 {
				if ret[l-1].End == leftNode.data {
					ret[l-1].End = n.data
					continue
				}
			}
			ret = append(ret, Interval{
				Start: leftNode.data,
				End:   n.data,
			})
		}
	}
	return ret
}

func main() {
	tests := []struct {
		input  []Interval
		new    Interval
		output []Interval
	}{
		{[]Interval{{1, 3}, {6, 9}}, Interval{2, 5}, []Interval{{1, 5}, {6, 9}}},
		{[]Interval{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, Interval{4, 9}, []Interval{{1, 2}, {3, 10}, {12, 16}}},
		{[]Interval{{1, 2}, {4, 5}}, Interval{1, 2}, []Interval{{1, 2}, {4, 5}}},
		{[]Interval{}, Interval{1, 2}, []Interval{{1, 2}}},
		{[]Interval{{1, 5}}, Interval{5, 7}, []Interval{{1, 7}}},
	}
	for _, test := range tests {
		ret := insert(test.input, test.new)
		fmt.Println(ret, test.output)
	}
}
