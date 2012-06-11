
/**
 * Project Euler
 * Author: Brendon Crawford
 *
 * Problem: 1
 */

package main

import "os"
import "fmt"


func main () () {
    const m uint32 = 1000
    var i uint32
    var r map[uint32] bool = map[uint32] bool {}
    var a uint32 = 0
    var incrs []uint32 = []uint32{3, 5}
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


