package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"os"
)

func isPrime(value uint64) bool {

	if value == 0 {
		return false
	}
	if value <= 2 {
		return true
	}
	var i uint64
	for i = 2; i < value; i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}

func PrimeRange (st int, en int, it int) int {

	count := 0
	for i:=st; i<en; i=i+it {
		if isPrime(uint64(i)) {
		   count++
		}
	}
	
	return count
}

func isPrimeHttp(res http.ResponseWriter, req *http.Request) {
	in_number := strings.Split(req.URL.Path, "/")

	var st,en,it int
	var err error
	
	nlen := len(in_number)
	if nlen == 4 {
		st, err = strconv.Atoi(in_number[1])
		en, err = strconv.Atoi(in_number[2])
		it, err = strconv.Atoi(in_number[3])
	} else if nlen == 3 {
		st, err = strconv.Atoi(in_number[1])
		en, err = strconv.Atoi(in_number[2])
		it = 1
	} else if nlen == 2 {
		st = 1
		en, err = strconv.Atoi(in_number[1])
		it = 1
	} else {
		fmt.Fprintln (res, "0")
	}
	

	if err != nil {
		fmt.Fprintln(res, "Number error %s", req.URL.Path) 
		return
	}
	
	nprimes := PrimeRange(st, en, it)
	
	fmt.Fprintln (res, fmt.Sprintf("%d",nprimes))
	
	return
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", isPrimeHttp)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}

}
