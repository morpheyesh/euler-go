package main

import "fmt"

func fibo() func() int {
     fmt.Print(0);
     a, b := 0, 1
       return func() int {
           a, b = b, a+b
            return a
        }
   }

func main() {
    f := fibo() 
   fmt.Print(f(),f(),f(),f(),f());
   
         
    
        
}
       
