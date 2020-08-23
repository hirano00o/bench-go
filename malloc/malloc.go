package malloc

import (
	"fmt"
	"strconv"
)

func everyStringAllocation(count int) {
	var base []string
	for i := 0; i < count; i++ {
		base = append(base, fmt.Sprintf("%d", i))
	}
}

func onceStringAllocation(count int) {
	base := make([]string, count)
	for i := 0; i < count; i++ {
		base[i] = fmt.Sprintf("%d", i)
	}
}

type child struct {
	id        string
	name      string
	createdAt string
	createdBy string
	updatedAt string
	updatedBy string
}

type parent struct {
	id        string
	name      string
	createdAt string
	createdBy string
	updatedAt string
	updatedBy string
	child1    child
	child2    child
}

func makeParent(id int) *parent {
	return &parent{
		id:        "user" + strconv.Itoa(id),
		name:      "user" + strconv.Itoa(id),
		createdAt: "2006-01-02 15:04:05",
		createdBy: "user1",
		updatedAt: "2006-01-02 15:04:05",
		updatedBy: "user1",
		child1: child{
			id:        "user" + strconv.Itoa(id+1),
			name:      "user" + strconv.Itoa(id+1),
			createdAt: "2006-01-02 15:04:05",
			createdBy: "user2",
			updatedAt: "2006-01-02 15:04:05",
			updatedBy: "user2",
		},
		child2: child{
			id:        "user" + strconv.Itoa(id+2),
			name:      "user" + strconv.Itoa(id+2),
			createdAt: "2006-01-02 15:04:05",
			createdBy: "user3",
			updatedAt: "2006-01-02 15:04:05",
			updatedBy: "user3",
		},
	}
}

func everyStructAllocation(count int) {
	var base []*parent
	for i := 0; i < count; i++ {
		base = append(base, makeParent(i))
	}
}

func onceStructAllocation(count int) {
	base := make([]*parent, count)
	for i := 0; i < count; i++ {
		base[i] = makeParent(i)
	}
}
