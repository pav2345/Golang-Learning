package main

import (
	"fmt"                    // standard library package for formatted I/O
	"githhub.com/google/uuid" // third-party package for generating UUIDs
	"math"                    // standard library package for mathematical functions
)

func main() {
	fmt.Println("HELLO wORLD,  time.Now()")
	u := uuid.New()
	fmt.Printf("Generated UUID: %s\n", u.String())
	fmt.Println(rand.IntN(1000))

}

// go Envrinoment variables

//GOROOT ---> Factory installtion of Go
//GOPATH --->IT is third party library cache custom librery whuch is callled go path
//GOBIN ---> all the tools which ever we install that are stored in go bin directory either we install or build get
//GOOS ---> It is the operating system on which we are running our go program	
//GOARCH ---> It is the Architecture of the system or machine on which we are running our go programme 
// the go bin is binary of GO it can store all the code base tools every there stoed it is eassy to platform Engineer adn devops Engineer to Loacte 

//bash 
` go tool dist list `


//GOARCH ---> Chipset arch 

# Build Process

//GO source --> Go compiler  --> SSA (OPtimization) constasts Dead Code Elimination -->  asmbler  plan9assembler ass0--> .o file ---> linker --> links all required files , memory mappings--> bin (static binary )
  

-java source -->javaCompiler --> javabyte code --> JVM --> loads byte code ---> jIT compilaton
-- Scala  Source --> scale Compiler --> javabyte code -->jvm --> loads the byte code --> JIT compilation 
- closure Score -->closure Compiler --> JAVaBYTYE Code -->JVM-->Load the byte code --> JIT COMPILATION 
-Kotlin Source --> kotlin Compiler --> java byte code --> JVM---> JIT COMPILATIOn 

# to start application ,run , go application 

``bash commands 
mkdir demo && cd demo 
go run main.go 
go mod init demo 
compiles ,Assembles  and LInks -----------> BInary 

#to kno w where the binary is created 

go run -work main.go



# Go compile link 

go tool 


## debug build vs release Build  --> nm tool 
 go build -o demo main 