
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


func main () () {
    fmt.Printf("%s\n", fermat(0))
    fmt.Printf("%s\n", fermat(1))
    fmt.Printf("%s\n", fermat(2))
    fmt.Printf("%s\n", fermat(3))
    fmt.Printf("%s\n", fermat(4))
    os.Exit(0)
}


func fermat (n uint64) (bool) {
    var yes uint8 = 1
    var no uint8 = 2
    var primes map[uint64]uint8 = map[uint64]uint8 {
        0: no, 4: no,
        1: yes, 2: yes, 3: yes,
    }
    if primes[n] == yes {
        return true
    } else if primes[n] == no {
        return false
    }
    var a uint64 = (((n - 1) - 2) / 2)
    var aBig *big.Int = big.NewInt(int64(a))
    var nBig *big.Int = big.NewInt(int64(n))
    var aToNBig *big.Int = new(big.Int).Exp(aBig, nBig, nil)
    var mod1Big *big.Int = new(big.Int).Mod(aToNBig, nBig)
    var mod2Big *big.Int = new(big.Int).Mod(aBig, nBig)
    var isPrime bool = (mod1Big.Cmp(mod2Big) == 0)
    return isPrime
}


