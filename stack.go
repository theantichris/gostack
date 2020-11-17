package gostack

import "sync"

// Item holds the data that will be stored on the Stack.
type Item interface{}

// Stack represents the stack data structure.
type Stack struct {
	items []Item
	mutex sync.Mutex
}

// Push adds a new Item to the Stack.
func (stack *Stack) Push(item Item) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.items = append(stack.items, item)
}

// Pop removes and returns the top item from the Stack.
func (stack *Stack) Pop() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if len(stack.items) == 0 {
		return nil
	}

	lastItem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]

	return lastItem
}

// IsEmpty returns true if there are no Items on the Stack.
func (stack *Stack) IsEmpty() bool {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	return len(stack.items) == 0
}

// Reset removes all Items from the Stack.
func (stack *Stack) Reset() {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.items = nil
}

// Dump returns all the Items in the stack.
func (stack *Stack) Dump() []Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	copied := make([]Item, len(stack.items))
	copy(copied, stack.items)

	return copied
}

// Peek returns the top Item in the stack.
func (stack *Stack) Peek() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if len(stack.items) == 0 {
		return nil
	}

	return stack.items[len(stack.items)-1]
}