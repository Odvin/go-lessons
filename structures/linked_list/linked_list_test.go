package linkedlisl

import "testing"

func TestLinkedList(t *testing.T) {
	t.Run("Prepend Node", func(t *testing.T) {
		l := LinkedList{}
		a := &Node{Data: 1}
		b := &Node{Data: 2}

		l.Prepend(a)
		l.Prepend(b)

		if l.length != 2 {
			t.Errorf("expected lenth 2 but got %d", l.length)
		}

		if l.head.Data != b.Data {
			t.Errorf("expected data %d but got %d", b.Data, l.head.Data)
		}

		if l.head != b {
			t.Errorf("expected head address %v but got %v", b, l.head)
		}

		if l.head.next != a {
			t.Errorf("expected next element address %v but got %v", a, l.head.next)
		}

		if l.head.next.Data != a.Data {
			t.Errorf("expected data %d but got %d", a.Data, l.head.next.Data)
		}
	})

	t.Run("Delete Data", func(t *testing.T) {
		l := LinkedList{}
		a := &Node{Data: 1}
		b := &Node{Data: 2}
		c := &Node{Data: 2}

		l.Prepend(a)
		l.Prepend(b)
		l.Prepend(c)

		l.DelValue(2)

		if l.length != 1 {
			t.Errorf("expected lenth 1 but got %d", l.length)
		}

		if l.head.Data != a.Data {
			t.Errorf("expected data %d but got %d", b.Data, l.head.Data)
		}
	})

	t.Run("Has data", func(t *testing.T) {
		l := LinkedList{}
		a := &Node{Data: 1}
		b := &Node{Data: 2}

		l.Prepend(a)
		l.Prepend(b)

		if l.length != 2 {
			t.Errorf("expected lenth 2 but got %d", l.length)
		}

		if l.HasData(1) != true {
			t.Errorf("expected 'true' but got %v", l.HasData(1))
		}

		if l.HasData(5) != false {
			t.Errorf("expected 'false' but got %v", l.HasData(1))
		}
	})
}
