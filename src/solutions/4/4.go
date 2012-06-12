
/**
 * Project Euler
 * Author: Brendon Crawford
 *
 * Problem: 4
 *
 * URL: https://projecteuler.net/problem=4
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


func main () () {
    const mLow uint32 = 100
    const mHigh uint32 = 999
    var o uint32 = 0
    var i, j, p uint32
    var s, r string
    for i = mHigh; i >= mLow; i-- {
        for j = mHigh; j >= mLow; j-- {
            p = i * j
            if p > o {
                s = strconv.FormatUint(uint64(p), 10)
                r = reverse_str(s)
                if s == r {
                    o = p
                }
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

