
package main

import "os"
import "fmt"


type IterInt struct {
    vals chan int
    end chan bool
}


func (iter IterInt) next () (bool) {
    return !<-iter.end
}


func (iter IterInt) value () (int) {
    return <-iter.vals
}


func (iter IterInt) yield (i int) (bool) {
    iter.end <- false
    iter.vals <- i
    return true
}


func (iter IterInt) close () (bool) {
    iter.end <- true
    return true
}


func NewIterInt (callback func (r IterInt) ()) (IterInt) {
    iter := IterInt{make(chan int), make(chan bool)}
    go func () (bool) {
        callback(iter)
        iter.close()
        return true
    }()
    return iter
}


func main () () {
    iter := NewIterInt(func (r IterInt) {
        foo(r)
    })
    for iter.next() {
        fmt.Printf("%d\n", iter.value())
    }
    os.Exit(0)
}


func foo (iter IterInt) () {
    for i := 0; i < 10; i++ {
        iter.yield(i)
    }
}
