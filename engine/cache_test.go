package engine

import (
	"testing"

	"github.com/jrecuero/go-cli/engine"
)

// TestCache_NewCache is ...
func TestCache_NewCache(t *testing.T) {
	cache := engine.NewCache()
	if cache == nil {
		t.Errorf("NewCache: can not return <nil>")
	}
	if cache.Size() != 0 {
		t.Errorf("NewCache: cache size mismatch: exp: %d got: %d\n", 0, cache.Size())
	}
}

// TestCache_Access is ...
func TestCache_Access(t *testing.T) {
	cache := engine.NewCache()
	if ok := cache.Add("home"); !ok {
		t.Errorf("Add: return code mismatch: exp: %v got: %v\n", true, ok)
	}
	if ok := cache.Set("home", "casa"); !ok {
		t.Errorf("Set: return code mismatch: exp: %v got: %v\n", true, ok)
	}
	if value, ok := cache.Get("home"); !ok {
		t.Errorf("Get: return code mismatch: exp: %v got: %v\n", true, ok)
	} else if value != "casa" {
		t.Errorf("Get: value mismatch: exp: %#v got: %#v\n", "casa", value)
	}
	if ok := cache.Delete("home"); !ok {
		t.Errorf("Delete: return code mismatch: exp: %v got: %v\n", true, ok)
	} else {
		if _, found := cache.Get("home"); found {
			t.Errorf("Delete: entry not properly deleted")
		}
	}
}

// TestCache_CheckCondition is ...
func TestCache_CheckCondition(t *testing.T) {
	cache := engine.NewCache()
	cache.Set("home", "casa")
	cache.Set("car", "coche")
	var tests = []struct {
		input *engine.CacheCondition
		exp   bool
	}{
		{
			input: engine.NewCacheCondition("home", "casa", engine.EQUAL),
			exp:   true,
		},
		{
			input: engine.NewCacheCondition("car", "casa", engine.NEQUAL),
			exp:   true,
		},
		{
			input: engine.NewCacheCondition("car", nil, engine.EXIST),
			exp:   true,
		},
		{
			input: engine.NewCacheCondition("bike", nil, engine.NEXIST),
			exp:   true,
		},
		{
			input: engine.NewCacheCondition("home", "home", engine.EQUAL),
			exp:   false,
		},
		{
			input: engine.NewCacheCondition("car", "coche", engine.NEQUAL),
			exp:   false,
		},
		{
			input: engine.NewCacheCondition("carro", nil, engine.EXIST),
			exp:   false,
		},
		{
			input: engine.NewCacheCondition("home", nil, engine.NEXIST),
			exp:   false,
		},
	}
	for i, test := range tests {
		if test.input == nil {
			t.Errorf("%d NewCacheCondition: return can not be <nil>", i)
		}
		if ok := cache.CheckCondition(test.input); ok != test.exp {
			t.Errorf("CheckCondition: return code mismatch: exp: %#v got : %#v\n", test.exp, ok)
		}
	}
}
