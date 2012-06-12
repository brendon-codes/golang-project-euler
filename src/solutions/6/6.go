
/**
 * Project Euler
 * Author: Brendon Crawford
 *
 * Problem: 6
 *
 * URL: https://projecteuler.net/problem=6
 *
 * Description:
 *
 *   The sum of the squares of the first ten natural numbers is,
 *
 *       1_{2} + 2_{2} + ... + 10_{2} = 385
 *
 *   The square of the sum of the first ten natural numbers is,
 *
 *       (1 + 2 + ... + 10)_{2} = 55_{2} = 3025
 *
 *   Hence the difference between the sum of the squares of the first ten
 *   natural numbers and the square of the sum is 3025 − 385 = 2640.
 *
 *   Find the difference between the sum of the squares of the first one
 *   hundred natural numbers and the square of the sum.
 *
 */

package main

import "os"
import "fmt"
import "math"


func main () () {
    const m uint64 = 100
    const ex uint64 = 2
    var s1 uint64 = sum_of_squares(m, ex)
    var s2 uint64 = square_of_sums(m, ex)
    var o uint64 = s2 - s1
    fmt.Printf("%d\n", o)
    os.Exit(0)
}

func square_of_sums (m uint64, ex uint64) (uint64) {
    var a uint64 = 0
    var i uint64
    for i = 1; i <= m; i++ {
        a += i
    }
    a = uint64(math.Pow(float64(a), float64(ex)))
    return a
}


func sum_of_squares (m uint64, ex uint64) (uint64) {
    var a uint64 = 0
    var i uint64
    for i = 1; i <= m; i++ {
        a += uint64(math.Pow(float64(i), float64(ex)))
    }
    return a
}


