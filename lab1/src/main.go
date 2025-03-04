package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	f, err = os.Create("mem.prof")
	if err != nil {
		log.Fatal("Could not create memory profile: ", err)
	}
	defer f.Close()
	runtime.GC() // Вызов сборщика мусора для получения актуальных данных
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("Could not write memory profile: ", err)
	}

	var n int
	fmt.Scan(&n)

	if n < 2 || n > 40 {
		panic("Invalid size")
	}

	t := NewTable(n)
	r := t.PlaceSquares()

	fmt.Print(r)
}
