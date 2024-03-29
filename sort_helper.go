/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomArray(n, min, max int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		num := rand.Intn(max-min) + min
		arr = append(arr, num)
	}
	return arr
}

// 判断arr数组是否有序
func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
func timeSpent(funcName string, inner func(arr []int), arr []int) {
	start := time.Now()
	inner(arr)
	ts := time.Since(start).Seconds()
	if !isSorted(arr) {
		return
	}
	fmt.Println(funcName+" time spent:", ts)

}
func copyArray(arr []int, n int) []int {
	var newArr = make([]int, n)
	copy(newArr, arr)
	return newArr
}
