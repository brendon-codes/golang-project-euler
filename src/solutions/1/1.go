
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
    var i int
    r := map[int] bool {}
    a := 0
    m := 1000
    incrs := []int{3, 5}
    for _, v := range incrs {
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


