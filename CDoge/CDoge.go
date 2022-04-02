package CDoge

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var File chan []string
var tims bool
var OverCha chan bool
var WordCount []string
var Maps sync.Map

func FileToMem(fielName string, fenGe string, endLog byte, size int) {
	start := time.Now()
	File = make(chan []string, size)
	OverCha = make(chan bool, size)
	fis, err := os.OpenFile(fielName, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("茶炊了")
		return
	}
	br := bufio.NewReader(fis)
	var pussyIndex []int
	for {
		if tims == false {
			sts, err := br.ReadString(endLog)
			if err != nil || io.EOF == err {
				break
			}
			tar := strings.Split(strings.Split(sts, string(endLog))[0], fenGe)
			defer fmt.Println(len(File), tar)
			pussyIndex = make([]int, 0, len(tar))
			for k, v := range tar {
				for _, va := range WordCount {
					if va == v {
						pussyIndex = append(pussyIndex, k)
						Maps.Store(va, nil)
					}
				}
			}
			defer fmt.Println(pussyIndex)
			tims = true
		} else {
			if len(File) >= cap(File) {
				var temFile chan []string
				temFile = make(chan []string, size)
				for len(File) > 0 {
					sts := <-File
					temFile <- sts
				}
				go MinCounts(temFile, pussyIndex)
				continue
			}
			sts, err := br.ReadString(endLog)
			if err != nil || io.EOF == err {
				var temFile chan []string
				temFile = make(chan []string, size)
				for len(File) > 0 {
					sts := <-File
					temFile <- sts
				}
				go MinCounts(temFile, pussyIndex)
				break
			}
			File <- strings.Split(strings.Split(sts, string(endLog))[0], fenGe)
		}
	}
	for len(OverCha) > 0 {
		// fmt.Println("的步伐组", len(OverCha))
		time.Sleep(time.Second)
	}
	Maps.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
	cost := time.Since(start)
	fmt.Printf("Funcost=[%s]\n", cost)
}

var cis int64

func WordCounts(File chan []string, pussyIndex []int) {
	start := time.Now()
	OverCha <- true
	// fmt.Println(len(File))
	for len(File) > 0 {
		sts := <-File
		for k, _ := range pussyIndex {
			val, sta := Maps.Load(WordCount[k])
			if sta == true {
				if val == nil {
					temAis := []string{sts[k] + "|" + "1"}
					Maps.Store(WordCount[k], temAis)
				} else {
					var temRul bool
					var vall = val.([]string)
					for ks, vs := range vall {
						if strings.Split(vs, "|")[0] == sts[k] {
							temval, _ := strconv.ParseInt(strings.Split(vs, "|")[1], 10, 32)
							val.([]string)[ks] = sts[k] + "|" + strconv.FormatInt((temval+1), 10)
							temRul = true
						}
					}
					if temRul == false {
						val = append(val.([]string), sts[k]+"|"+"1")
						Maps.Store(WordCount[k], val)
					}
				}
			}
		}
	}
	fmt.Println("cis", cis)
	cis += 1
	cha := <-OverCha
	cost := time.Since(start)
	fmt.Printf("cost=[%s]\n", cost)
	fmt.Println(cha)
}
func MaxCounts(File chan []string, pussyIndex []int) {
	start := time.Now()
	OverCha <- true
	// fmt.Println(len(File))
	for len(File) > 0 {
		sts := <-File
		for k, _ := range pussyIndex {
			val, sta := Maps.Load(WordCount[k])
			if sta == true {
				if val == nil {
					temAis := []int64{0}
					Maps.Store(WordCount[k], temAis)
				} else {
					var vall = val.([]int64)
					for ks, vs := range vall {
						azh, _ := strconv.ParseInt(sts[k], 10, 64)
						if vs < azh {
							val.([]int64)[ks] = azh
						}
					}
				}
			}
		}
	}
	fmt.Println("cis", cis)
	cis += 1
	cha := <-OverCha
	cost := time.Since(start)
	fmt.Printf("cost=[%s]\n", cost)
	fmt.Println(cha)
}
func MinCounts(File chan []string, pussyIndex []int) {
	start := time.Now()
	OverCha <- true
	// fmt.Println(len(File))
	for len(File) > 0 {
		sts := <-File
		for k, _ := range pussyIndex {
			val, sta := Maps.Load(WordCount[k])
			if sta == true {
				if val == nil {
					temAis := []int64{0}
					Maps.Store(WordCount[k], temAis)
				} else {
					var vall = val.([]int64)
					for ks, vs := range vall {
						azh, _ := strconv.ParseInt(sts[k], 10, 64)
						if vs > azh {
							val.([]int64)[ks] = azh
						}
					}
				}
			}
		}
	}
	fmt.Println("cis", cis)
	cis += 1
	cha := <-OverCha
	cost := time.Since(start)
	fmt.Printf("cost=[%s]\n", cost)
	fmt.Println(cha)
}
