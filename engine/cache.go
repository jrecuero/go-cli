package engine

// CacheCondition represents ...
type CacheCondition struct {
	key   string
	value interface{}
	let   ConditionLet
}

// Check is ...
func (cc CacheCondition) Check(cache *Cache) bool {
	switch cc.let {
	case EQUAL:
		if cacheValue, ok := cache.Get(cc.key); ok {
			return cacheValue == cc.value
		}
		break
	case NEQUAL:
		if cacheValue, ok := cache.Get(cc.key); ok {
			return cacheValue != cc.value
		}
		break
	case EXIST:
		if _, ok := cache.Get(cc.key); ok {
			return true
		}
		break
	case NEXIST:
		if _, ok := cache.Get(cc.key); !ok {
			return true
		}
		break
	default:
		break
	}
	return false
}

// NewCacheCondition is ...
func NewCacheCondition(key string, value interface{}, let ConditionLet) *CacheCondition {
	return &CacheCondition{
		key:   key,
		value: value,
		let:   let,
	}
}

// Cache represents ...
type Cache struct {
	data map[string]interface{}
}

// Size is ...
func (cache *Cache) Size() int {
	return len(cache.data)
}

// Add is ...
func (cache *Cache) Add(name string) bool {
	if _, ok := cache.data[name]; !ok {
		cache.data[name] = nil
		return true
	}
	return false
}

// Delete is ...
func (cache *Cache) Delete(name string) bool {
	if _, ok := cache.data[name]; ok {
		delete(cache.data, name)
		return true
	}
	return false
}

// Set is ...
func (cache *Cache) Set(name string, value interface{}) bool {
	//if _, ok := cache.data[name]; ok {
	cache.data[name] = value
	return true
	//}
	//return false
}

// Get is ...
func (cache *Cache) Get(name string) (interface{}, bool) {
	if value, ok := cache.data[name]; ok {
		return value, true
	}
	return nil, false
}

// CheckCondition is ...
func (cache *Cache) CheckCondition(cond interface{}) bool {
	condition := cond.(*CacheCondition)
	return condition.Check(cache)
}

// NewCache is ...
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}
