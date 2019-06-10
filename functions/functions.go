package functions

import (
	"fmt"
	"math/rand"
	"time"
)

func SleepTime(maxTime int){
	startTime := time.Now()
	// Sleep between 0 and Xms
	time.Sleep(time.Duration(rand.Intn(maxTime)) * time.Millisecond)
	// Record how long we slept in ms
	endTime := float64(time.Now().Sub(startTime).Nanoseconds() / 1e6)
	fmt.Printf("%v",endTime)
}

func SleepTwo(){
	startTime := time.Now()
	// Sleep 2s
	time.Sleep(2000* time.Millisecond)
	// Record how long we slept in ms
	endTime := float64(time.Now().Sub(startTime).Nanoseconds() / 1e6)
	fmt.Printf("%v",endTime)
}

func SleepOne(){
	startTime := time.Now()
	// Sleep 1s
	time.Sleep(1000* time.Millisecond)
	// Record how long we slept in ms
	endTime := float64(time.Now().Sub(startTime).Nanoseconds() / 1e6)
	fmt.Printf("%v",endTime)
}