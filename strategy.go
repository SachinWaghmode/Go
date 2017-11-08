package main

import (
	"fmt";
	"math/rand"
)

type SortingStrategy interface{
	sort(data []int) []int
}

type Strategy struct {
	SortingStrategy SortingStrategy
}

func (s *Strategy) ChangeStrategy(a []int) []int {
	return s.SortingStrategy.sort(a)
}

type BubbleSort struct{}

func (BubbleSort) sort(a []int) []int {
 	end := len(a) - 1
	for {
		if end == 0 {
			break
		}
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
		end -= 1
	}
	return a
}

type QuickSort struct{}

func (quicksort QuickSort) sort(a []int) []int{
 	length := len(a)

	if length <= 1 {
		temp := make([]int, length)
		copy(temp, a)
		return temp
	}

	m := a[rand.Intn(length)]

	left := make([]int, 0, length)
	middle := make([]int, 0, length)
	right := make([]int, 0, length)

	for _, item := range a {
		switch {
		case item < m:
			left = append(left, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			right = append(right, item)
		}
	}

	left, right = quicksort.sort(left), quicksort.sort(right)
	left = append(left, middle...)
	left = append(left, right...)
	return left
}

type MergeSort struct{}

func (mergesort MergeSort) sort(a []int) []int{
    if len(a) <= 1 {
        return a
    }

    left, right := split(a)
    left = mergesort.sort(left)
    right = mergesort.sort(right)
    return merge(left, right)
}

func split(a[]int) ([]int, []int) {
    return a[0:len(a) / 2], a[len(a) / 2:]
}

func merge(A, B []int) []int {
    arr := make([]int, len(A) + len(B))
    j, k := 0, 0

    for i := 0; i < len(arr); i++ {
        if j >= len(A) {
            arr[i] = B[k]
            k++
            continue
        } else if k >= len(B) {
            arr[i] = A[j]
            j++
            continue
        }
        if A[j] > B[k] {
            arr[i] = B[k]
            k++
        } else {
            arr[i] = A[j]
            j++
        }
    }

    return arr
}

func Test(){
	dataset := []int{198,56,12,34,62,99,145,87,26,37,16,13}
	selectstrategy := Strategy{BubbleSort{}}
	result :=selectstrategy.ChangeStrategy(dataset) 
	fmt.Println("Bubble sort  : ",result)

	selectstrategy = Strategy{QuickSort{}}
	result =selectstrategy.ChangeStrategy(dataset) 
	fmt.Println("Quick sort  : ",result)
	
	selectstrategy = Strategy{MergeSort{}}
	result=selectstrategy.ChangeStrategy(dataset) 
	fmt.Println("Merge sorting  : ",result)
}

func main(){
	Test()
}
