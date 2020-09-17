package linked_list_test

import (
	linked_list "github.com/mrdulin/go-datastructure-algorithm/linked-list"
	"reflect"
	"testing"
)

func TestSinglyLinkedList_Append(t *testing.T) {
	t.Run("should append node to linked list", func(t *testing.T) {
		list := linked_list.NewSinglyLinkedList()
		node := linked_list.NewNode('1', nil)
		list.Append(node)
		if !reflect.DeepEqual(1, list.Length()) {
			t.Errorf("the length of list, got: %d, want: %d", list.Length(), 1)
		}
		t.Logf("%+v", list)
	})
}

func TestSinglyLinkedList_Length(t *testing.T) {
	t.Run("should get length of a linked list", func(t *testing.T) {
		list := linked_list.NewNode("1", linked_list.NewNode("2", linked_list.NewNode("3", nil)))
		got := linked_list.Length(list)
		want := 3
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %d, want: %d", got, want)
		}
		if !reflect.DeepEqual(list, linked_list.NewNode("1", linked_list.NewNode("2", linked_list.NewNode("3", nil)))) {
			t.Errorf("should not change original list")
		}
	})
}

func TestGetKthLastElement(t *testing.T) {
	t.Run("should get kth last element", func(t *testing.T) {
		list := linked_list.NewNode("1", linked_list.NewNode("2", linked_list.NewNode("3", nil)))
		got := linked_list.GetKthLastElement(list, 0).Item
		want := "3"
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %s, want: %s", got, want)
		}
	})
}
