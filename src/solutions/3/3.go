
/**
 * Project Euler
 * Author: Brendon Crawford
 *
 * Problem: 3
 *
 * URL: https://projecteuler.net/problem=3
 *
 * Description:
 *   The prime factors of 13195 are 5, 7, 13 and 29.
 *   What is the largest prime factor of the number 600851475143 ?
 *
 */

package main

import "os"
import "fmt"


func main () () {
    const m uint64 = 600851475143
    var i uint64 = 0
    var v uint64 = m
    var p uint64 = m
    for {
        i = fac(v)
        if i == 1 {
            p = v
            break
        } else {
            p = i
            v /= i
        }
    }
    fmt.Printf("%d\n", p)
    os.Exit(0)
}


func fac (m uint64) (uint64) {
    var i uint64
    if m == 1 {
        return 1
    }
    for i = 2; i < m; i++ {
        if m % i == 0 {
            return i
        }
    }
    return 1
}

