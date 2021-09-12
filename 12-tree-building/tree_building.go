package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	// no records
	if len(records) == 0 {
		return nil, nil
	}
	// records should be ordered
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	// create each node and save it in a map
	nodes := make(map[int]*Node)
	for _, r := range records {
		if _, ok := nodes[r.ID]; ok {
			return nil, errors.New("duplicated node")
		}
		if r.ID == 0 && r.Parent != 0 {
			return nil, errors.New("root node cannot have parent")
		}
		if r.ID > 0 && r.ID == r.Parent {
			return nil, errors.New("cannot cycle directly")
		}
		if r.ID < r.Parent {
			return nil, errors.New("cannot cycle indirectly")
		}
		nodes[r.ID] = &Node{
			ID:       r.ID,
			Children: nil,
		}
	}

	// there should be a root node
	if _, ok := nodes[0]; !ok {
		return nil, errors.New("no root node")
	}

	// build the tree
	for i, r := range records {
		if i != r.ID {
			return nil, errors.New("tree non-continuous")
		}
		// for every no-root node, add it to its parent
		if r.ID > r.Parent {
			list := nodes[r.Parent].Children
			if list == nil {
				list = make([]*Node, 0)
			}
			list = append(list, nodes[r.ID])
			nodes[r.Parent].Children = list
		}
	}

	return nodes[0], nil
}
