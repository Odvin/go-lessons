package tree

import "testing"

func TestTree(t *testing.T) {
	t.Run("Insert", func(t *testing.T) {
		// 		     50
		// 	   30           60
		//  20    nil   nil    70
		root := Node{Key: 50}
		root.Insert(30)
		root.Insert(60)
		root.Insert(20)
		root.Insert(70)

		if root.Key != 50 {
			t.Errorf("expected 50 but got %d", root.Key)
		}

		if root.right.Key != 60 {
			t.Errorf("expected 60 but got %d", root.right.Key)
		}

		if root.left.Key != 30 {
			t.Errorf("expected 30 but got %d", root.left.Key)
		}

		node := root.left

		if node.right != nil {
			t.Errorf("expected nil but got %v", node.right)
		}

		if node.left.Key != 20 {
			t.Errorf("expected 20 but got %d", root.left.Key)
		}

		node = root.right

		if node.left != nil {
			t.Errorf("expected nil but got %v", node.right)
		}

		if node.right.Key != 70 {
			t.Errorf("expected 70 but got %d", root.left.Key)
		}
	})

	t.Run("Insert", func(t *testing.T) {
		// 		     50
		// 	   30           60
		//  20    nil   nil    70
		leafL := &Node{Key: 20}
		leafR := &Node{Key: 70}
		parentL := &Node{Key: 30, left: leafL}
		parentR := &Node{Key: 60, right: leafR}
		root := Node{Key: 50, left: parentL, right: parentR}

		if root.Search(50) != true {
			t.Errorf("expected true but got %v", root.Search(50))
		}

		if root.Search(60) != true {
			t.Errorf("expected true but got %v", root.Search(60))
		}

		if root.Search(20) != true {
			t.Errorf("expected true but got %v", root.Search(20))
		}

		if root.Search(100) != false {
			t.Errorf("expected false but got %v", root.Search(100))
		}

	})
}
