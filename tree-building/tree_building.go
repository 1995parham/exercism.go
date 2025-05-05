package tree

import (
	"fmt"
)

// Record provides node and parent id
type Record struct {
	ID, Parent int
}

// Node of tree
type Node struct {
	ID       int
	Children []*Node
}

// RecordMismatch represents error with record id
type RecordMismatch struct {
	RecordID int
}

func (m RecordMismatch) Error() string {
	return fmt.Sprintf("Invalid record ID: %d", m.RecordID)
}

// LenMismatch represents inequality in added records and given records
type LenMismatch struct {
	AddedRecords int
	GivenRecords int
}

func (m LenMismatch) Error() string {
	return fmt.Sprintf("Added records (%d) is not equal to Given records (%d)", m.AddedRecords, m.GivenRecords)
}

// Build builds tree based on given records
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// sort and validate records
	sortedRecords := make([]Record, len(records))
	for _, r := range records {
		if r.ID >= len(records) {
			return nil, RecordMismatch{RecordID: r.ID}
		}
		if r.ID == 0 {
			if r.Parent == 0 {
				continue
			} else {
				return nil, fmt.Errorf("root node have parent")
			}
		} else if r.ID <= r.Parent {
			return nil, RecordMismatch{RecordID: r.ID}
		} else if sortedRecords[r.ID].ID != 0 {
			return nil, RecordMismatch{RecordID: r.ID}
		} else {
			sortedRecords[r.ID] = r
		}

	}

	// nodes stores nodes as value and their parent id as key in array
	// because "The ID number is always between 0 (inclusive) and the length of the record list (exclusive)."
	nodes := make([][]*Node, len(records))
	for _, r := range sortedRecords {
		if r.ID == 0 {
			continue
		}
		nodes[r.Parent] = append(nodes[r.Parent], &Node{ID: r.ID})
	}

	root := &Node{}
	todo := []*Node{root}
	n := 1
	for len(todo) != 0 {
		var newTodo []*Node
		for _, c := range todo {
			nn := nodes[c.ID]
			if len(nn) == 0 {
				continue
			}
			n += len(nn)
			c.Children = nn
			newTodo = append(newTodo, nn...)
		}
		todo = newTodo
	}
	if n != len(records) {
		return nil, LenMismatch{
			AddedRecords: n,
			GivenRecords: len(records),
		}
	}
	return root, nil
}
