package collections

type LinkedList struct {
	length int
	root   Node
}

type Node struct {
	Value interface{}
	next  *Node
	ll    *LinkedList
}

func (ll *LinkedList) Init() *LinkedList {
	ll.root.next = nil
	ll.length = 0
	return ll
}

// Push prepends an element to the back of the linked list.
func (ll *LinkedList) Push(val interface{}) {
	var n Node
	n.Value = val
	next := ll.root
	n.next = &next
	ll.root, ll.length = n, ll.length+1
}

// Pop removes and returns the element from the back of the stack.
func (ll *LinkedList) Pop() interface{} {
	last := ll.root.Value
	ll.root = *ll.root.next
	return last
}

// Last returns last the element of the linked list.
func (ll *LinkedList) Last() interface{} {
	return ll.root.Value
}

// Last returns last the element of the linked list.
func (ll *LinkedList) Next() interface{} {
	return ll.root.next.Value
}

// First returns first the element of the linked list.
func (ll *LinkedList) First() interface{} {
	f := ll.root
	return ll.getFirst(f)
}

func (ll *LinkedList) getFirst(f Node) interface{} {
	for f.next.Value != nil {
		f = *f.next
	}
	return f.Value
}
