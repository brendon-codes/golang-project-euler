
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
    m := 4000000
    v1 := 1
    v2 := 2
    v := v1
    a := 0
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


