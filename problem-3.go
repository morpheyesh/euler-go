package main


import "fmt"


func main() {
          
        tadaa := int64(0);
         opp := int64(0);
         end := int64(600851475143);

         for i := int64(2); i < end; i++ {
            if end%i == 0 {
               opp = end / i;
                    if  prime(opp) {
                        tadaa = opp                        
                         break;
                
                        }
              

           } 
           
       }
   fmt.Printf("%d\n", tadaa);
}

func prime(num int64) bool {
         
    for i := int64(2); i < num ; i++ { 
               if num%i == 0 {
                       return false
                    }
          
        }
        return true
 }

