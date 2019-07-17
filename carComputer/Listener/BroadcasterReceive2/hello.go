
package main

import (
//       "encoding/hex"
       "fmt"
       "net"
       "os"
       "strconv"
       "sync"
       "log"
//       "strings"       
       //"github.com/dmichael/go-multicast/multicast"
       "github.com/urfave/cli"
)

const (
      defaultMulticastAddress = "239.0.0.0:9000"
      maxDatagramSize = 65500
)


//credit for this function goes to user dmichael on github
func msgHandler(src *net.UDPAddr, n int, b []byte) {
     //fmt.Println(n, "bytes read from", src)
     //fmt.Println(hex.Dump(b[:n]))
}



// Listen binds to the UDP address and port given and writes packets received
// from that address to a buffer which is passed to a hander
func Listen2(address string, handler func(*net.UDPAddr, int, []byte)) {
     // Parse the string address
     var upper int = 0
     var cooler int = 34
     var popper int = 1
     path := "/go/mydata/"+strconv.Itoa(popper)+".jpg"
     fmt.Println(path)
     addr, err := net.ResolveUDPAddr("udp", address)
     if err != nil {
     	log.Fatal(err)
     }
     
     // Open up a connection
     conn, err := net.ListenMulticastUDP("udp", nil, addr)
     if err != nil {
	   log.Fatal(err)
     }
     
     conn.SetReadBuffer(maxDatagramSize)
     
     // Loop forever reading from the socket
     file1, err := os.OpenFile(path,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
     if err != nil {
            fmt.Println("File Creation error!!\n")
     }

     var checker int = 0
     var mimicker int = 0
     var flag1 int = 1
     //var numBytes int = 0
 
     for {
	 buffer := make([]byte, maxDatagramSize)
	 //var checker int = 0 
	 numBytes, _, _ := conn.ReadFromUDP(buffer)
	 fmt.Println(numBytes)
	 if flag1 != 1{
	    mimicker,_=strconv.Atoi(string(buffer[:numBytes]))
	    if mimicker != 0{
	       flag1 = 1
	    }
	 }
	 if (numBytes != 0) && (flag1 == 1){
	    if cooler == 34{
	       cooler = 1
	       checker,_ = strconv.Atoi(string(buffer[:numBytes]))
	       fmt.Println("GOT 34!!!!!!")
	    } else{
	      upper = upper + numBytes 
	      if err != nil {
	       	 log.Fatal("ReadFromUDP failed:", err)
	      }



	   
	    n, err := file1.Write(buffer)
	    _ = n
	    if err != nil{
	       fmt.Println("File writing error!\n\n")
	       os.Exit(1)
	    }
	    fmt.Println(checker, upper)	   
	    //handler(src, numBytes, buffer)
	    if checker == upper{
	       upper = 0
	       flag1 = 5
	       cooler = 34
	       popper = popper + 1
	       path := "/go/mydata/"+strconv.Itoa(popper)+".jpg"
	       fmt.Println(path)
	       file1.Close()
	       file1, err = os.OpenFile(path,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	       if err != nil{
	       	  fmt.Println("File creation error!!!\n")
	       }
	       fmt.Println(file1)
	    }
          }
	 }
     }
}






//credit for code listening on port 3001 goes to user dmichael on github
func main() {
     var wg sync.WaitGroup

     wg.Add(1)

     //This function will act as the listener of the car on port 3001
     go func(){
	    

	    app := cli.NewApp()
	    fmt.Println("asdflasdkjf;alskdfja;sdlkfj\n\n")
    	    app.Action = func(c *cli.Context) error {
    	       	       address := c.Args().Get(0)
	       	       if address == "" {
		       	  	  address = defaultMulticastAddress
		       }
		       fmt.Printf("Listening on %s\n", address)
		       Listen2(address, msgHandler)
		       return nil
	    }

	    app.Run(os.Args)


     }()

     wg.Wait()
}
