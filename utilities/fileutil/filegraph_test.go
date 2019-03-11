package fileutil

import "testing"

func TestFileGraphFolder_NewChildFolderChain(t *testing.T) {
	t.Logf("New child folder chain")

	fg := NewFileGraph()
	fg.NewChildFolders([]string{"This", "Is", "A", "Test"})

	child1 := *fg.Root.ChildFolders[0]
	child2 := *child1.ChildFolders[0]
	child3 := *child2.ChildFolders[0]
	child4 := *child3.ChildFolders[0]

	if child1.Name != "This" {
		t.Errorf("%s != This", child1.Name)
	}

	if child2.Name != "Is" {
		t.Errorf("%s != Is", child2.Name)
	}

	if child3.Name != "A" {
		t.Errorf("%s != A", child3.Name)
	}

	if child4.Name != "Test" {
		t.Errorf("%s != Test", child4.Name)
	}
}
