package main

import(
  "github.com/bnczk/rgrpc-go/messages"
  "github.com/bnczk/rgrpc-go/services"
  "fmt"
)


func main(){
  fmt.Println(messages.Empty{}) 
  fmt.Println(services.UnimplementedSomeAPIServer{})
}
