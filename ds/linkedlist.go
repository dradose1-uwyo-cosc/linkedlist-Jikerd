package ds
import "fmt"

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (list *LinkedList) Insert(value string) {
	newNode := &Node{data:value}

	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	}else {
		list.Tail.Next = newNode
		list.Tail = newNode
	}
	list.Size++
}

func (list *LinkedList) InsertAt(position int, value string) error {
	if position < 0 || position > list.Size {
		return fmt.Errorf("Invalid Position")
	}
	if position == 0 {
		newNode := &Node {
			data: value,
			Next: list.Head,
		}
		list.Head = newNode

		if list.Size == 0 {
			list.Tail = newNode
		}

		list.Size++
		return nil
		}
	current := list.Head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}
	newNode := &Node{
		data: value,
		Next: current.Next,
	}
	current.Next = newNode
	if newNode.Next == nil {
		list.Tail = newNode
	}

	list.Size++
	return nil
}

func (list *LinkedList) Remove(value string) error {
	if list.Head ==nil {
		return fmt.Errorf("List is empty.")
	}
	if list.Head.data == value {
		list.Head = list.Head.Next
		list.Size--

		if list.Size == 0 {
			list.Tail = nil
		}
		return nil
	}
	current := list.Head
	for current.Next != nil && current.Next.data != value {
		current = current.Next
	}
	if current.Next == nil {
		return fmt.Errorf("Value not found")
	}
	current.Next = current.Next.Next
	if current.Next == nil {
		list.Tail = current
	}
	list.Size--
	return nil
}

func (list *LinkedList) RemoveAll(value string) error {
	if list.Head == nil {
		return fmt.Errorf("list is empty")
	}

	found := false
	for list.Head != nil && list.Head.data == value {
		list.Head = list.Head.Next
		list.Size--
		found = true
	}
	if list.Head == nil {
		list.Tail = nil
		if found {
			return nil
		}
		return fmt.Errorf("value not found")
	}
	current := list.Head
	for current.Next != nil {
		if current.Next.data == value {
			current.Next = current.Next.Next
			list.Size--
			found = true
		}else {
			current = current.Next
		}
	}

	list.Tail = current

	if !found {
		return fmt.Errorf("value not found")
	}
	return nil
}


func (list *LinkedList) RemoveAt(pos int) error {
	if pos < 0 || pos >= list.Size {
		return fmt.Errorf("Invalid Position")
	}
	if pos == 0 {
		list.Head = list.Head.Next
		list.Size--

		if list.Size == 0 {
			list.Tail= nil
		}
		return nil
	}
	current := list.Head
	for i := 0; i < pos-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	if current.Next == nil {
		list.Tail = current
	}
	list.Size--
	return nil
}

func (list *LinkedList) IsEmpty() bool {
	return list.Size == 0
}

func (list *LinkedList) GetSize() int {
	return list.Size
}

func (list *LinkedList) Reverse() {
	var prev *Node
	current := list.Head
	list.Tail = list.Head

	for current != nil {
		nextNode := current.Next
		current.Next = prev
		prev = current
		current = nextNode
	}
	list.Head = prev
}

func (list *LinkedList) PrintList() {
	current := list.Head
	for current != nil {
		fmt.Print(current.data, "->")
		current = current.Next
	}
	fmt.Println("nil")
}
