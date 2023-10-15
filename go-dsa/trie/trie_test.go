package trie

import (
	"reflect"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	trie.Insert("/users", func() bool { return true })
	trie.Insert("/users/:userId", func() bool { return true })
	trie.Insert("/users/:userId/profile", func() bool { return true })
	trie.Insert("/users/:userId/groups/:groupId", func() bool { return true })
	trie.Insert("/users/:userId/groups/:groupId/approve", func() bool { return true })
	trie.Insert("/auth/login", func() bool { return true })
	trie.Insert("/auth/logout", func() bool { return true })

	tests := []struct {
		path   string
		params map[string]string
		valid  bool
	}{
		{"/users", emptyMap(), true},
		{"/users/1", map[string]string{"userId": "1"}, true},
		{"/users/1/profile", map[string]string{"userId": "1"}, true},
		{"/users/1/groups/2", map[string]string{"userId": "1", "groupId": "2"}, true},
		{"/users/1/groups/2/approve", map[string]string{"userId": "1", "groupId": "2"}, true},
		{"/auth/login", emptyMap(), true},
		{"/auth/logout", emptyMap(), true},
		{"/invalid-path", nil, false},
	}

	for _, test := range tests {
		result := trie.Search(test.path)
		if result.Handler() != test.valid {
			t.Errorf("Handler for path %s was expected to be %v but got %v", test.path, test.valid, result.Handler())
		}

		if !reflect.DeepEqual(result.Params, test.params) {
			t.Errorf("Expected params for path %s to be %v but got %v", test.path, test.params, result.Params)
		}
	}
}

func emptyMap() map[string]string {
	return make(map[string]string)
}

func TestTrieInvalidInserts(t *testing.T) {
	trie := NewTrie()

	trie.Insert("/users", func() bool { return true })
	trie.Insert("/users/:id", func() bool { return true })

	tests := []struct {
		path      string
		wantError string
	}{
		{"/users/:userId/approve", "Invalid path: /users/:userId/approve. Parameter userId is already defined."},
		{"/users/:id/group/:id", "cannot use the same parameter name more than once in a route"},
	}

	for _, test := range tests {
		err := trie.Insert(test.path, func() bool { return true })
		if err == nil || err.Error() != test.wantError {
			t.Errorf("Expected error for path %s to be %v but got %v", test.path, test.wantError, err)
		}
	}
}
