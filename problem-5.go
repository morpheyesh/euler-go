package main

import (
        "fmt";
)

func divs(n int64) bool {
       
        for i := int64(1); i < 20; i++ {
                if n%i != 0 {
                        return false
                }
        }
        return true;
}

func main() {
        x := int64(20);
        for {
                if divs(x) {
                        break
                }
                x += 20;
        }
        fmt.Printf("%d\n", x);
}




