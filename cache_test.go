package cache

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

var cache *Cache[int, int]

func init() {
	cache = NewCache[int, int](1000, 1*time.Second)
}

func TestCache(t *testing.T) {
	const count = 100000
	var wg sync.WaitGroup

	wg.Add(count)
	go func() {
		for i := range count {
			cache.Add(i, i)
			wg.Done()
		}
	}()

	wg.Add(count)
	go func() {
		for i := range count {
			cache.Get(i)
			cache.CacheCapacity()
			cache.CacheSize()
			wg.Done()
		}
	}()
	wg.Wait()

	if cache.CacheSize() != 1000 {
		t.Fatal("incoorect size")
	}

	fmt.Printf("Cache size: %d\n", cache.CacheSize())
}

func BenchmarkCache(b *testing.B) {
	cache := NewCache[int, struct{}](1000, 1*time.Second)
	b.ResetTimer()
	b.Run("Add element", func(b *testing.B) {

		for j := 0; j < b.N; j++ {
			func() {
				cache.Add(j, struct{}{})
			}()
		}
	})

	b.Run("Get element", func(b *testing.B) {

		for j := 0; j < b.N; j++ {
			func() {
				cache.Get(j)
			}()
		}
	})

	b.StopTimer()
	time.Sleep(time.Second)
	b.StartTimer()
	b.Run("Get element with timeout", func(b *testing.B) {

		for j := 0; j < b.N; j++ {
			func() {
				cache.Get(j)
			}()
		}

	})
}

func printCache(cache *Cache[int, int], caseNum int) {
	fmt.Printf("==== %d case =====\n", caseNum)
	defer fmt.Println("=================")

	cache.mu.Lock()
	defer cache.mu.Unlock()

	fmt.Printf("head: %p, tail %p\n", cache.head, cache.tail)
	h := cache.head
	var i int
	for h != nil {
		fmt.Printf("%d %p %+v\n", i, h, h)
		h = h.next
		i++
	}
}
func TestCorrectWork(t *testing.T) {
	cache := NewCache[int, int](3, time.Millisecond)

	ok := cache.Add(0, 0)
	if !ok {
		t.Fatal("0 case")
	}
	ok = cache.Add(1, 1)
	if !ok {
		t.Fatal("0 case")
	}
	ok = cache.Add(2, 2)
	if !ok {
		t.Fatal("0 case")
	}
	printCache(cache, 0)
	if cache.CacheSize() != 3 {
		t.Fatal("0 case")
	}

	v, ok := cache.Get(0)
	printCache(cache, 1)
	if !ok || v != 0 {
		t.Fatal("1 case")
	}

	ok = cache.Add(3, 3)
	if !ok {
		t.Fatal("2 case")
	}
	v, ok = cache.Get(1)
	printCache(cache, 2)
	if ok || v == 1 {
		t.Fatal("2 case")
	}

	time.Sleep(time.Millisecond)
	v, ok = cache.Get(0)
	printCache(cache, 3)
	if cache.CacheSize() != 2 || ok || !reflect.ValueOf(v).IsZero() {
		t.Fatal("3 case")
	}

	cache.Get(2)
	printCache(cache, 4)
	if cache.head != cache.tail {
		t.Fatal("4 case")
	}

	ok = cache.Add(3, 33)
	if ok {
		t.Fatal("5 case")
	}
	printCache(cache, 5)
	v, ok = cache.Get(3)
	if !ok || v != 33 {
		t.Fatal("5 case")
	}

	ok = cache.Add(2, 22)
	if !ok || cache.tail.value != 22 {
		t.Fatal("6 case")
	}
	time.Sleep(time.Millisecond)
	ok = cache.Add(3, 333)
	printCache(cache, 6)
	if ok || cache.tail.value != 333 {
		t.Fatal("6 case")
	}

	time.Sleep(time.Millisecond)
	v, ok = cache.Get(2)
	if cache.CacheSize() != 1 || ok || !reflect.ValueOf(v).IsZero() {
		t.Fatal("7 case")
	}
	v, ok = cache.Get(3)
	if cache.CacheSize() != 0 || ok || !reflect.ValueOf(v).IsZero() {
		t.Fatal("7 case")
	}
	printCache(cache, 7)
	if cache.tail != nil || cache.head != nil {
		t.Fatal("7 case")
	}

	ok = cache.Add(0, 666)
	if !ok || cache.head.value != 666 || cache.tail.value != 666 {
		t.Fatal("8 case")
	}
	ok = cache.Add(1, 111)
	if !ok || cache.head.value != 666 || cache.tail.value != 111 {
		t.Fatal("8 case")
	}
	time.Sleep(time.Millisecond)
	v, ok = cache.Get(1)
	printCache(cache, 8)
	if cache.CacheSize() != 1 || ok || !reflect.ValueOf(v).IsZero() {
		t.Fatal("8 case")
	}
	if cache.head != cache.tail || cache.tail.next != nil || cache.tail.previous != nil {
		t.Fatal("8 case")
	}
}
