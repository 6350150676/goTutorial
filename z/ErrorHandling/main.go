package main

import "fmt"

func divide (a , b float64)(float64, error){
if b== 0{
return 0,fmt.Errorf("denominator is zeor")
}
return a / b , nil
}
func main() {
fmt.Println("started Error Handling in the functions") 
ans , _:= divide(10, 4)
fmt. Println("Division of two numbers is ", ans)
}