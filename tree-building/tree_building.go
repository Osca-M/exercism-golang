package tree

import (
	"errors"
	"fmt"
	"sort"
)

// Record holds information on an occurrence and its parent
type Record struct {
	ID     int
	Parent int
}

// Node holds information on an occurrence and its children
type Node struct {
	ID       int
	Children []*Node
}

func makeNode(ID int) *Node {
	return &Node{
		ID: ID,
	}
}

// Build builds a tree from a record
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Slice(records, func(i, j int) bool {
		r1 := records[i]
		r2 := records[j]
		return r1.ID < r2.ID
	})
	nodes := make([]*Node, len(records))

	rootRecord, records := records[0], records[1:]
	if rootRecord.ID != 0 {
		return nil, errors.New("no root record")
	}
	nodes[0] = makeNode(rootRecord.ID)
	for i, record := range records {
		if rootRecord.Parent != 0 || record.ID <= record.Parent || i+1 != record.ID {
			return nil, fmt.Errorf("not in sequence or has bad parent: %v", record)
		}
		if i > 0 {
			p := records[i-1]
			if p.ID == record.ID && p.Parent == record.Parent {
				return nil, errors.New("duplicate record")
			}
		}
		parent := nodes[record.Parent]
		n := makeNode(record.ID)
		nodes[n.ID] = n
		parent.Children = append(parent.Children, n)
	}
	return nodes[0], nil
}
