package ds


type Stack struct {
        list LinkedList

}

func (stack*Stack) Push(value string) {
        newNode := &Node{ data: value, Next: stack.list.Head}
        stack.list.Head = newNode
        if stack.list.Size == 0 {
                stack.list.Tail = newNode
        }
        stack.list.Size++
}

func (stack*Stack) Pop() (string, bool) {
        if stack.list.Head == nil {
                return "", false
        }
        value:= stack.list.Head.data
        stack.list.Head = stack.list.Head.Next
        stack.list.Size--
        if stack.list.Size == 0 {
                stack.list.Tail = nil
        }
        return value, true
}







