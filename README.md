# OptimusPrime
A simple program that finds all the prime numbers in the given range.  This alos has a docker image. 

Get the docker image in cloudfoudry lattice
```
$ltc create -d=1 -m=1 optimusprime dhilipkumars/optimusprime
```

Now we have a REST full api to find prime numbers. :-)

Find all the prime numbers between 1 and 1000
```
$curl http://optimusprime.192.168.11.11.xip.io/1/1000
169
```

Build the client program that wraps this http calls. 
```
$go build Primerange.go
```

Lets find all the prime numbers between 1 to 200000 this time using the client program
```
$time Primerange -P=http://optimusprime.192.168.11.11.xip.io -R=200000
17985 Prime numbers are there in first 200000 numbers

real    0m13.694s
user    0m0.006s
sys     0m0.005s
```

Now lets Parllel process and get faster results. Use 4 goroutines to calculate prime numbers 

```
 time Primerange -P=http://optimusprime.192.168.11.11.xip.io -R=200000 -T=4
17985 Prime numbers are there in first 200000 numbers

real    0m6.653s
user    0m0.006s
sys     0m0.007s

```

8 goroutines
```
time Primerange -P=http://optimusprime.192.168.11.11.xip.io -R=200000 -T=8
17985 Prime numbers are there in first 200000 numbers

real    0m3.639s
user    0m0.013s
sys     0m0.014s
```

16 goroutine

```
time Primerange -P=http://optimusprime.192.168.11.11.xip.io -R=200000 -T=16
17985 Prime numbers are there in first 200000 numbers

real    0m2.688s
user    0m0.011s
sys     0m0.012s

```

This program was used to demonstrate how we can containerize simple workloads using Lattice at Cloud Foundry Meet up 2 bangalore.
