package basics

import (
	"fmt"
	"github.com/GalushkoArt/simpleCache"
)

func Caches() {
	mapCache := simpleCache.NewMapCache()
	mapCache.Set("1", 1)

	var value = *mapCache.Get("1")
	fmt.Println("Existed + 1:", value.(int)+1)
	fmt.Println("Not Existed:", mapCache.Get("2"))

	fmt.Println("Deleted:", *mapCache.Delete("1"))
	fmt.Println("Not Deleted:", mapCache.Delete("2"))
	fmt.Println("After Delete:", mapCache.Get("1"))

	fmt.Println("\nGeneric cache\n")
	genericCache := simpleCache.NewGenericMapCache[int]()
	genericCache.Set("1", 1)

	fmt.Println("Existed + 1:", *genericCache.Get("1")+1)
	fmt.Println("Not Existed:", genericCache.Get("2"))

	fmt.Println("Deleted:", *genericCache.Delete("1"))
	fmt.Println("Not Deleted:", genericCache.Delete("2"))
	fmt.Println("After Delete:", genericCache.Get("1"))
}
