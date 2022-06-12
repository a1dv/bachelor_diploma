package main

import (
  "fmt"
  "time"
  "sync"
)

//6th ed. Numerical Mathematics and Computing problem 11.2.2
//Using psuedo code on pg. 470
var wg sync.WaitGroup

//RK4 Solver
func RK4(t float64, x []float64, h float64, ic int, n int){
	K1 := make([]float64, ic)
	K2 := make([]float64, ic)
	K3 := make([]float64, ic)
	K4 := make([]float64, ic)
	R := make([]float64, ic)

  k1 := make(chan []float64)
  k2 := make(chan []float64)
  k3 := make(chan []float64)
  k4 := make(chan []float64)
  r := make(chan []float64)
  xChan := make(chan []float64)

	for i := 0; i < n; i++{
    if i % 10 == 0 {
      wg.Wait()
    }
    wg.Add(1)
  		go iteration(K1, K2, K3, K4, R, t, x, h, ic)
  		t = t + h
  }
}

func iteration(K1, K2, K3, K4, R []float64, t float64, x[]float64, h float64, ic int) {
  defer wg.Done()
  AiryDE(t, x, K1)

  for j := 0; j < ic; j++{
    R[j] = x[j] + h/2*K1[j]
  }
  AiryDE(t+h/2, R, K2)

  for j := 0; j < ic; j++{
    R[j] = x[j] + h/2*K2[j]
  }
  AiryDE(t+h/2, R, K3)

  for j := 0; j < ic; j++{
    R[j] = x[j] + h*K3[j]
  }
  AiryDE(t+h, R, K4)

  for j := 0; j < ic; j++{
    x[j] = x[j] + h/6*(K1[j]+2*(K2[j]+K3[j])+K4[j])
  }
}

//Airy DE Function
func AiryDE(t float64, x []float64, k []float64){
	k[0] = x[1]
	k[1] = t*x[0]
}

func main() {
  defer func(since time.Time) {
      fmt.Println(time.Since(since).String())
  }(time.Now())
	var t float64
	var h float64
	x := make([]float64, 2)
	t = 0.0
	h = 4.8789 / 1000000000  //h = (4.5-0)/100
	x[0] = 0.355028053887817 // intial x(0)
	x[1] = -0.258819403792807 // intial x'(0)
	RK4(t, x, h, 2, 922337257)
	expected := 0.0003302503
	fmt.Println(x[0], " ", x[1], " ", (expected-x[0])/2)
}
