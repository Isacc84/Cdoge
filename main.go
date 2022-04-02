package main

import (
	"Cdoge/CDoge"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	CDoge.WordCount = make([]string, 3)
	CDoge.WordCount[0] = "disk_id"
	CDoge.WordCount[1] = "ds"
	CDoge.WordCount[2] = "model"
	CDoge.FileToMem("G:/smartlog2019ssd/20190109.csv", ",", '\n', 100)
	cost := time.Since(start)
	fmt.Printf("Alcost=[%s]\n", cost)
}
