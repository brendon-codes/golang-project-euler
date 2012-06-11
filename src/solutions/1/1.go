
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
    const m int = 1000
    var i int
    var r map[int] bool = map[int] bool {}
    var a int = 0
    var incrs []int = []int{3, 5}
    var v int
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


