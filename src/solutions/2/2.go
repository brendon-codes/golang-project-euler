
/**
 * Project Euler
 * Author: Brendon Crawford
 *
 * Problem: 2
 */

package main

import "os"
import "fmt"


func main () () {
    const m uint32 = 4000000
    var v1 uint32 = 1
    var v2 uint32 = 2
    var v uint32 = v1
    var a uint32 = 0
    for v <= m {
        v = v1 + v2
        v1 = v2
        v2 = v
        if v1 % 2 == 0 {
            a += v1
        }
    }
    fmt.Printf("%d\n", a)
    os.Exit(0)
}


