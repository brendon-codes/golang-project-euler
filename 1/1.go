
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
    fmt.Printf("Sum: %d\n", sum([]int{5, 10}))
    //m := 10
    //incrs := []int{3, 5}
    //for _, incr := range incrs {
    //    fmt.Printf("Total: %d\n", cycle(0, m, incr))
    //}
    os.Exit(0)
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


