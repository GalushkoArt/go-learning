package basics

import (
	"fmt"
	"github.com/GalushkoArt/simpleCache"
	"strconv"
	"time"
)

func Caches() {
	mapCache := simpleCache.NewMapCache()
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		mapCache.Set(strconv.Itoa(i), i)
	}
	var value int
	for i := 0; i < 1000000; i++ {
		value = (*mapCache.Get(strconv.Itoa(i))).(int)
	}
	fmt.Println("Set and retrieve of", value, "elements took", time.Since(start))

	fmt.Println("Not Existed:", mapCache.Get("1000002"))

	fmt.Println("Deleted:", *mapCache.Delete("1"))
	fmt.Println("Not Deleted:", mapCache.Delete("1000002"))
	fmt.Println("After Delete:", mapCache.Get("1"))

	fmt.Println("\nGeneric cache\n")
	genericCache := simpleCache.NewGenericMapCache[int]()
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		genericCache.Set(strconv.Itoa(i), i)
	}
	for i := 0; i < 1000000; i++ {
		value = *genericCache.Get(strconv.Itoa(i))
	}
	fmt.Println("Set and retrieve of", value, "elements took", time.Since(start))

	fmt.Println("Not Existed:", genericCache.Get("1000002"))

	fmt.Println("Deleted:", *genericCache.Delete("1"))
	fmt.Println("Not Deleted:", genericCache.Delete("1000002"))
	fmt.Println("After Delete:", genericCache.Get("1"))
}
