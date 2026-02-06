package ds
import "fmt"

type Queue struct {
	list LinkedList
}

func (queue*Queue) Push(value string) {
	queue.list.Insert(value)
}

func (queue*Queue) Pop() (string, error) {
	if queue.list.Head == nil {
		return "", fmt.Errorf("queue empty")
	}
	value := queue.list.Head.data
	queue.list.Head = queue.list.Head.Next
	queue.list.Size--
	if queue.list.Size == 0 {
		queue.list.Tail = nil
	}
	return value, nil
}
