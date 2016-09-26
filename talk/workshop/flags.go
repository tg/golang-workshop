package main

import (
	"flag"
	"log"
	"runtime"
)

func main() {
	cpus := flag.Int("cpus", runtime.NumCPU(), "set number of CPU(s) to use")
	flag.Parse()

	runtime.GOMAXPROCS(*cpus)
	log.Printf("Using %d CPU(s)", *cpus)
}
