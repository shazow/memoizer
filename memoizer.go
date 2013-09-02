package memoizer

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

var ErrMissedCache = errors.New("Memoizer: Missed cache.")

type Cacher interface {
	CreateKey(f interface{}, callArgs []interface{}) string
	Get(key string, object *interface{}) error
	Set(key string, object interface{}) error
}

type Memoizer interface {
	Cacher
	Call(f interface{}, callArgs ...interface{}) interface{}
	// TODO: Replace(f interface{}) interface{}
}

// Cacher base struct

type BaseCache struct{}

func (c *BaseCache) CreateKey(f interface{}, callArgs []interface{}) string {
	fName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	// TODO: Hash function name + args?
	return fName + ":" + fmt.Sprint(callArgs)
}

// Cacher example implementation

type MemoryCache struct {
	*BaseCache
	storage map[string]interface{}
}

func (c *MemoryCache) Get(key string, object *interface{}) error {
	r, ok := c.storage[key]
	*object = r
	if !ok {
		return ErrMissedCache
	}
	return nil
}

func (c *MemoryCache) Set(key string, object interface{}) error {
	c.storage[key] = object
	return nil
}

func NewMemoryCache() *MemoryCache {
	c := MemoryCache{}
	c.storage = make(map[string]interface{})
	return &c
}

// Implement Memoizer

type Memoize struct {
	Cache Cacher
}

// Call function using memoize technique with storage method defined by m.cache.
//
// Function can return up to one arbitrary value and one error. If error is not
// nil, caching is skipped and error is returned with the value.
// TODO: Is it possible to support arbitrary numbers of return values?
//
// See source for text/template/funcs.go for a similar call example.
func (m *Memoize) Call(f interface{}, callArgs ...interface{}) (interface{}, error) {
	key := m.Cache.CreateKey(f, callArgs)

	var r interface{}
	err := m.Cache.Get(key, &r)
	if err == nil {
		return r, nil
	}

	reflectArgs := make([]reflect.Value, len(callArgs))
	for i, arg := range callArgs {
		reflectArgs[i] = reflect.ValueOf(arg)
	}

	result := reflect.ValueOf(f).Call(reflectArgs)
	if len(result) == 0 {
		// No return value.
		return nil, nil
	}
	r = result[0].Interface()

	if len(result) == 2 {
		// Has error return value, check it before saving to cache.
		err = result[1].Interface().(error)
		if err != nil {
			return r, err
		}
	}

	m.Cache.Set(key, r)
	return r, nil
}
