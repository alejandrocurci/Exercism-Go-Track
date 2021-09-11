package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type NodeContainer struct {
	Node     *Node
	ParentID int
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
	// create each node and save it in a map
	containers := make(map[int]*NodeContainer)
	for _, r := range records {
		if _, ok := containers[r.ID]; ok {
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
		containers[r.ID] = &NodeContainer{
			Node: &Node{
				ID:       r.ID,
				Children: nil,
			},
			ParentID: r.Parent,
		}
	}
	// there should be a root node
	if _, ok := containers[0]; !ok {
		return nil, errors.New("no root node")
	}

	// build the tree
	for _, cont := range containers {
		// for every no-root nodes, add it to its parent
		if cont.Node.ID > cont.ParentID {
			list := containers[cont.ParentID].Node.Children
			if list == nil {
				list = make([]*Node, 0)
			}
			list = append(list, cont.Node)
			// node slices should be ordered
			sort.Slice(list, func(i, j int) bool {
				return list[i].ID < list[j].ID
			})
			containers[cont.ParentID].Node.Children = list
		}
	}
	for i := 0; i < len(records); i++ {
		if _, ok := containers[i]; !ok {
			return nil, errors.New("tree non-continuous")
		}

	}
	return containers[0].Node, nil
}
