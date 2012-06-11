
/**
 * Project Euler
 * Author: Brendon Crawford
 *
 * Problem: 4
 *
 * Description:
 *  A palindromic number reads the same both ways. The largest palindrome
 *  made from the product of two 2-digit numbers is 9009 = 91 × 99.
 *  Find the largest palindrome made from the product of two 3-digit numbers.
 *
 */

package main

import "os"
import "fmt"
import "strconv"
import "math"


func main () () {
    const mLow uint64 = 100
    const mHigh uint64 = 999
    var o uint64 = 0
    var i, j, p uint64
    var s, r string
    for i = mHigh; i >= mLow; i-- {
        for j = mHigh; j >= mLow; j-- {
            p = i * j
            s = strconv.FormatUint(p, 10)
            r = reverse_str(s)
            if s == r {
                o = uint64(math.Max(float64(o), float64(p)))
            }
        }
    }
    fmt.Printf("%d\n", o)
    os.Exit(0)
}


func reverse_str (s string) (string) {
    var l int = len(s)
    var m []rune = make([]rune, l)
    var c rune
    for _, c = range s {
        l--
        m[l] = c
    }
    return string(m)
}

