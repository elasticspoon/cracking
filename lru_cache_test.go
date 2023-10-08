package ctci

import "testing"

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache[string](2)

	cache.Set("a", 1)
	cache.Set("b", 2)

	if val, ok := cache.Get("a"); ok {
		if val != 1 {
			t.Errorf("Expected to find value 1 for key 'a', got %v", val)
		}
	} else {
		t.Errorf("Expected to find key 'a' in cache")
	}

	cache.Set("c", 8)

	if val, ok := cache.Get("c"); ok {
		if val != 8 {
			t.Errorf("Expected to find value 8 for key 'c', got %v", val)
		}
	} else {
		t.Errorf("Expected to find key 'c' in cache")
	}

	if val, ok := cache.Get("a"); ok {
		if val != 1 {
			t.Errorf("Expected to find value 1 for key 'a', got %v", val)
		}
	} else {
		t.Errorf("Expected to find key 'a' in cache")
	}

	if _, ok := cache.Get("b"); ok {
		t.Errorf("Expected to find not key 'b' in cache")
	}

	cache.Set("c", 0)

	if val, ok := cache.Get("c"); ok {
		if val != 0 {
			t.Errorf("Expected to find value 0 for key 'c', got %v", val)
		}
	} else {
		t.Errorf("Expected to find key 'c' in cache")
	}
}
