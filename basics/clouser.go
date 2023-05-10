package basics

func Counter() func() int64 {
	var counter int64 = 0
	return func() int64 {
		counter++
		return counter
	}
}
