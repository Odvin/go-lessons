package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Push", func(t *testing.T) {
		s := Stack{}

		a := 100
		b := 200

		s.Push(a)
		s.Push(b)

		if len(s.items) != 2 {
			t.Errorf("expected number of the elements 2 but got %d", len(s.items))
		}

		if s.items[0] != a {
			t.Errorf("expected number is %d but got %d", a, s.items[0])
		}

		if s.items[1] != b {
			t.Errorf("expected number is %d but got %d", b, s.items[1])
		}
	})

	t.Run("Pop", func(t *testing.T) {
		s := Stack{
			items: []int{100, 200},
		}

		if len(s.items) != 2 {
			t.Errorf("expected number of the elements 2 but got %d", len(s.items))
		}

		item, err := s.Pop()

		if item != 200 || err != nil {
			t.Errorf("expected number is 200 but got %d with error %v", item, err)
		}

		if len(s.items) != 1 {
			t.Errorf("expected number of the elements 1 but got %d", len(s.items))
		}

		item, err = s.Pop()

		if item != 100 || err != nil {
			t.Errorf("expected number is 100 but got %d with error %v", item, err)
		}

		if len(s.items) != 0 {
			t.Errorf("expected number of the elements 0 but got %d", len(s.items))
		}

		item, err = s.Pop()

		if item != 0 || err.Error() != "empty stack" {
			t.Errorf("expected number is 0 but got %d with error %v", item, err)
		}
	})
}
