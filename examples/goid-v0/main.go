package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/chai2010/goroutine"
)

func GoID() int {
	var buf [128]byte
	n := runtime.Stack(buf[:], false)

	fmt.Println("stack:", string(buf[:n]))

	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("can not get goroutine id: %v", err))
	}
	return id
}
func main() {
	fmt.Println("main", GoID())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i, GoID())

			g := goroutine.GetGroutine()
			gid := reflect.ValueOf(g).FieldByName("goid").Int()

			fmt.Printf("goroutine(%d): id = %d\n", i, gid)
		}()
	}
	wg.Wait()
}
