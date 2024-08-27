package trie

import "testing"

func TestTrie(t *testing.T) {
	t.Run("Insert", func(t *testing.T) {
		// 					  root
		//			a			e					o
		// 			r			r					r
		// 		a				a					e
		// 		g				g				g		o
		// 		o				o			o		a
		//	r		g			n			n		n
		// 	n										o
		tr := InitTrie()
		tr.Insert("aragorn")
		tr.Insert("aragorg")
		tr.Insert("eragon")
		tr.Insert("oregon")
		tr.Insert("oregano")
		tr.Insert("oreo")

		if tr.root.children['a'-'a'] == nil ||
			tr.root.children['e'-'a'] == nil ||
			tr.root.children['o'-'a'] == nil {
			t.Errorf("node must be present but empty")
		}

		if tr.root.children['b'-'a'] != nil ||
			tr.root.children['c'-'a'] != nil ||
			tr.root.children['r'-'a'] != nil {
			t.Errorf("node must be empty but present")
		}

		node := tr.root.children['a'-'a']
		if node.children['r'-'a'] == nil || node.isEnd {
			t.Errorf("node must be present and not the final")
		}
	})

}
