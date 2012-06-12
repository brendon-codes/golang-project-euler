
/**
 * Project Euler
 * Author: Brendon Crawford
 *
 * Problem: 5
 *
 * URL: https://projecteuler.net/problem=5 
 *
 * Description:
 *  2520 is the smallest number that can be divided by each of the numbers
 *  from 1 to 10 without any remainder.
 *  What is the smallest positive number that is evenly divisible by all
 *  of the numbers from 1 to 20?
 *
 */

package main

import "os"
import "fmt"


func main () () {
    const m uint32 = 20
    const s uint32 = (m - 1) * m
    var o uint32 = 0
    var y bool
    var f, i, j uint32
    if m % 2 == 0 {
        f = (m / 2) + 1
    } else {
        f = ((m + 1) / 2) + 1
    }
    for i = s;; i++ {
        y = true
        for j = f; j <= m; j++ {
            if i % j != 0 {
                y = false
                break
            }
        }
        if y {
            o = i
            break
        }
    }
    fmt.Printf("%d\n", o)
    os.Exit(0)
}

