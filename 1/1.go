
/**
 * Project Euler
 * Author: Brendon Crawford
 *
 * Problem: 1
 */

package main

import "os"
import "fmt"


func main () () {
    fmt.Printf("Mapper: %d\n", mapper(foo, []int{1, 3, 5}))
    //fmt.Printf("Sum: %d\n", sum([]int{5, ;10}))
    //m := 10
    //incrs := []int{3, 5}
    //for _, incr := range incrs {
    //    fmt.Printf("Total: %d\n", cycle(0, m, incr))
    //}
    os.Exit(0)
}


func foo (i int) (int) {
    return i + 2
}


func mapper (callback func(int) int, items[]int) ([]int) {
    o := make([]int, len(items), len(items))
    out_items := mapper_ints(items, o, 0, callback)
    return out_items
}


func mapper_ints (in_items[]int, out_items[]int,
                  i int, callback func(int) int) ([]int) {
    if i < len(in_items) {
        out_items[i] = callback(in_items[i])
        return mapper_ints(in_items, out_items, i + 1, callback)
    }
    return out_items
}


func sum (items[]int) (int) {
    return sum_ints(items, 0, 0)
}


func sum_ints (items[]int, i int, t int) (int) {
    if i < len(items) {
        return sum_ints(items, i + 1, t + items[i])
    }
    return t
}


func cycle (i int, m int, incr int) (int) {
    v := i + incr
    if v < m {
        return (incr + cycle(v, m, incr))
    } else if v == m {
        return 0
    }
    return i
}


