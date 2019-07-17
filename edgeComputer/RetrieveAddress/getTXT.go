package main

import (
       "fmt"
       "net"
       "sync"
       "os"
       "strconv"
)

func main(){
     var wg sync.WaitGroup

     wg.Add(1)
     go func() {
     	l,err := net.Listen("tcp",":4001")
	if err != nil{
	   panic(err)
	}
	defer l.Close()
	var sid int = 1
	for{
	   c, err := l.Accept()
	   if err != nil {
	      fmt.Println("FUCK\n")
	   } else{
	      fmt.Println("Success")
	   }
	   data := make([]byte,5000)
	   
	   //write address to file in volume
	   sid2 := strconv.Itoa(sid)
	   path := "/go/mem1/"+sid2+".txt"
	   file1,err := os.OpenFile(path,os.O_APPEND|os.O_CREATE|os.O_RDWR,0644)
	   
	   if err != nil{
	      fmt.Println("OH CRAP! AN ERROR!!!!!")
	   } else{	   


	      c.Read(data)
	      defer file1.Close()
	      //fmt.Println(data)
	      //now take in the data buffers data
	      n,err := file1.Write(data)
	      _ = n
	      if err != nil{
	      	 fmt.Println("Successful connection but no file success!!\n\n")
	      }
	      //defer file1.Close()
	   
	      c.Write([]byte("Message Received\n"))
	      sid = sid + 1
	      c.Close()
	   }
	}
     }()

     wg.Wait()

}