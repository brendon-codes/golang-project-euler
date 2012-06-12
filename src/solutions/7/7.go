
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
import "math/big"
import crypto_rand "crypto/rand"
import math_rand "math/rand"


func main () () {
    //const n uint32 = 19
    //const nMin uint32 = 2
    //fmt.Printf("%d\n", random(3, 51))
    //fmt.Printf("%d\n", random(3, 51))
    fmt.Printf("%d\n", random(3, 10))
    os.Exit(0)
}


func random (min uint32, max uint32) (uint32) {
    const bits uint8 = 32
    var r *big.Int
    var o uint32
    r, _ = crypto_rand.Prime(crypto_rand.Reader, int(bits))
    math_rand.Seed(r.Int64())
    o = (uint32(math_rand.Int31n(int32(max - min))) + min)
    return o
}


