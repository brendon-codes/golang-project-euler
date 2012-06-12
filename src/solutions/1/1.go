
/**
 * Project Euler
 * Author: Brendon Crawford
 *
 * Problem: 1
 *
 * URL: https://projecteuler.net/problem=1
 *
 * Description:
 *
 *   If we list all the natural numbers below 10 that are multiples
 *   of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 *
 *   Find the sum of all the multiples of 3 or 5 below 1000.
 *
 */

package main

import "os"
import "fmt"


func main () () {
    const m uint32 = 1000
    var incrs []uint32 = []uint32{3, 5}
    var r map[uint32] bool = map[uint32] bool {}
    var i uint32
    var a uint32 = 0
    var v uint32
    for _, v = range incrs {
        for i = v; i < m; i += v {
            if !r[i] {
                a += i
                r[i] = true
            }
        }
    }
    fmt.Printf("%d\n", a)
    os.Exit(0)
}


