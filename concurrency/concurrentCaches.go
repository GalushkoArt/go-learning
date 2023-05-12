package concurrency

import (
	"fmt"
	"github.com/GalushkoArt/simpleCache"
	"time"
)

func Caches() {
	mapCache := simpleCache.NewConcurrentCache(100 * time.Millisecond)
	mapCache.Set("1", 1)
	mapCache.Set("3", 3)

	var value = *mapCache.Get("1")
	var value3 = *mapCache.Get("3")
	fmt.Println("Existed + 1:", value.(int)+1)
	fmt.Println("Value 3 at start:", value3)
	fmt.Println("Not Existed:", mapCache.Get("2"))

	fmt.Println("Deleted:", *mapCache.Delete("1"))
	fmt.Println("Not Deleted:", mapCache.Delete("2"))
	fmt.Println("After Delete:", mapCache.Get("1"))
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Value 3 after purify:", mapCache.Get("3"))

	fmt.Println("\nGeneric cache\n")
	genericCache := simpleCache.NewGenericConcurrentCache[int](50 * time.Millisecond)
	genericCache.Set("1", 1)
	genericCache.Set("3", 3)

	fmt.Println("Existed + 1:", *genericCache.Get("1")+1)
	fmt.Println("Value 3 at start:", *genericCache.Get("3"))
	fmt.Println("Not Existed:", genericCache.Get("2"))

	fmt.Println("Deleted + 1:", *genericCache.Delete("1")+1)
	fmt.Println("Not Deleted:", genericCache.Delete("2"))
	fmt.Println("After Delete:", genericCache.Get("1"))
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Value 3 after purify:", genericCache.Get("3"))
}
