package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"flag"
	"runtime"
	"io/ioutil"
)

var (
	flagThread = flag.Int("T", 1, "Number of parllel request")
	flagPage   = flag.String("P", "http://www.google.com", "The webpage you want to read ")
	flagRange  = flag.Int("R", 1000, "Range of Prime numbers")
)

func isRangePrimeHttp(st, en, n int) int {
	Readpage := fmt.Sprintf("%s/%d/%d/%d", *flagPage, st, en, n)
	value, err := http.Get(Readpage)
	if err != nil {
		return 0
	}
	ReadValue, err := ioutil.ReadAll(value.Body)
	value.Body.Close()
	if err != nil {
		return 0
	}
	ReadValueStr := string(ReadValue)
	ReadValueStr = strings.Trim(ReadValueStr, " ")
	ReadValueStr = strings.Trim(ReadValueStr, "\n")
	rvalue , err := strconv.Atoi(ReadValueStr)
	//fmt.Printf("Result is %s\n", ReadValueStr )
	return rvalue
}

func RangePrime(st int, rng int, it int, res chan int) {

	
	count := isRangePrimeHttp (st,rng,it)
	res <- count
}

func main() {

	flag.Parse()
	runtime.GOMAXPROCS(*flagThread)
	sum := 0
	
	total := make (chan int)
	
	for i:=0; i<*flagThread; i++ {	
		go RangePrime(i+1, *flagRange, *flagThread, total)
	}
	
	for i:=0; i<*flagThread; i++ {	
		sum += <-total
	}
	
	fmt.Printf("%d Prime numbers are there in first %d numbers\n", sum, *flagRange)
}
