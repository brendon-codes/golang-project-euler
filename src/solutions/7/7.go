
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


func main () () {
    const count uint64 = 1000000
    const m uint64 = 10001
    var z *[]uint64 = seives(count)
    var o uint64
    if uint64(len(*z)) >= m {
        o = (*z)[m - 1]
    } else {
        o = 0
    }
    fmt.Printf("%d\n", o)
    os.Exit(0)
}


func seives (count uint64) (*[]uint64) {
    const offset uint64 = 2
    var s []uint64 = make([]uint64, count, count)
    var z []uint64 = make([]uint64, 0, 0)
    var i, j, r, t, a uint64
    for i = 0; i < count; i++ {
        s[i] = i + offset
    }
    for j = 0; j < uint64(len(s)); j++ {
        r = s[j]
        if r != 0 {
            for a = j + r; a < uint64(len(s)); a += r {
                s[a] = uint64(0)
            }
        }
    }
    for t = 0; t < uint64(len(s)); t++ {
        if s[t] != 0 {
            z = append(z, s[t])
        }
    }
    return &z
}

