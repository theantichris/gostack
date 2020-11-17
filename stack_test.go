package gostack

import "testing"

const (
	item1 = "Item One"
	item2 = "Item Two"
)

func TestStack_IsEmpty(t *testing.T) {
	t.Run("it returns true for empty Stack", func(t *testing.T) {
		stack := Stack{}

		got := stack.IsEmpty()

		if got != true {
			t.Errorf("stack should be reported empty: got %v want %v", got, true)
		}
	})

	t.Run("it returns false for non-empty Stack", func(t *testing.T) {
		stack := Stack{
			items: []Item{item1},
		}

		got := stack.IsEmpty()

		if got != false {
			t.Errorf("stack should be reported non-empty: got %v want %v", got, false)
		}
	})
}

func TestStack_Push(t *testing.T) {
	t.Run("it adds an Item to the Stack", func(t *testing.T) {
		stack := Stack{}

		stack.Push(item1)

		if stack.items[0] != item1 {
			t.Errorf("item was not added: got %v want %v", stack.items[0], item1)
		}
	})
}

func TestStack_Pop(t *testing.T) {
	t.Run("it removes and returns an Item from the Stack", func(t *testing.T) {
		items := []Item{item1, item2}
		stack := Stack{items: items}

		item := stack.Pop()

		if item != item2 {
			t.Fatalf("wrong item was returned: got %v want %v", item, item2)
		}

		if stack.items[0] != item1 {
			t.Errorf("wrong items at top of stack: got %v want %v", stack.items[0], item1)
		}
	})

	t.Run("it returns nil if Stack is empty", func(t *testing.T) {
		stack := Stack{}

		item := stack.Pop()

		if item != nil {
			t.Errorf("did not return nil for empty stack: got %v", item)
		}
	})
}

func TestStack_Reset(t *testing.T) {
	t.Run("it sets the Stack to empty", func(t *testing.T) {
		items := []Item{item1, item1}
		stack := Stack{items: items}

		stack.Reset()

		if len(stack.items) != 0 {
			t.Errorf("stack is not empty: got %v", stack.items)
		}
	})
}

func TestStack_Dump(t *testing.T) {
	t.Run("it returns a copy of the Stack", func(t *testing.T) {
		items := []Item{item1, item1}
		stack := Stack{items: items}

		got := stack.Dump()

		for i := range got {
			if got[i] != items[i] {
				t.Fatalf("copy does not match: got %v, want %v", got[i], items[i])
			}
		}
	})
}

func TestStack_Peek(t *testing.T) {
	t.Run("it returns the top Item in the Stack", func(t *testing.T) {
		items := []Item{item1, item2}
		stack := Stack{items: items}

		if stack.Peek() != item2 {
			t.Fatalf("the wrong item was returned: got %v, want %v", stack.Peek(), item2)
		}

		if stack.Peek() != item2 {
			t.Error("the item was removed from the stack")
		}
	})

	t.Run("it returns nil for empty Stack", func(t *testing.T) {
		stack := Stack{}

		got := stack.Peek()

		if got != nil {
			t.Errorf("did not get nil for empty Stack: got %v", got)
		}
	})
}