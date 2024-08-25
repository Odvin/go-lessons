package queue

import "testing"

func TestQueue(t *testing.T) {
	t.Run("Enqueue", func(t *testing.T) {
		q := Queue{}

		a := 100
		b := 200

		q.Enqueue(a)
		q.Enqueue(b)

		if len(q.items) != 2 {
			t.Errorf("expected number of the elements 2 but got %d", len(q.items))
		}

		if q.items[0] != a {
			t.Errorf("expected number is %d but got %d", a, q.items[0])
		}

		if q.items[1] != b {
			t.Errorf("expected number is %d but got %d", b, q.items[1])
		}
	})

	t.Run("Dequeue", func(t *testing.T) {
		q := Queue{
			items: []int{100, 200},
		}

		if len(q.items) != 2 {
			t.Errorf("expected number of the elements 2 but got %d", len(q.items))
		}

		item, err := q.Dequeue()

		if item != 100 || err != nil {
			t.Errorf("expected number is 100 but got %d with error %v", item, err)
		}

		if len(q.items) != 1 {
			t.Errorf("expected number of the elements 1 but got %d", len(q.items))
		}

		item, err = q.Dequeue()

		if item != 200 || err != nil {
			t.Errorf("expected number is 200 but got %d with error %v", item, err)
		}

		if len(q.items) != 0 {
			t.Errorf("expected number of the elements 0 but got %d", len(q.items))
		}

		item, err = q.Dequeue()

		if item != 0 || err.Error() != "empty queue" {
			t.Errorf("expected number is 0 but got %d with error %v", item, err)
		}
	})
}
