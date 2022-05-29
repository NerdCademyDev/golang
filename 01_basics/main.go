package main

import "fmt"
/

// this is a single line comment
/*
	this
	is a 
	multiline 
	comment
*/
/*
imports from external libraries can be local(from your workspace) or public(github or any other repositories)
e.g:
public :
import "github.com/NerdCademy/golang/01_basics"
local : 
import "golang/01_basics"

NB: local imports will only work after initialising the src module by using ```git mod init "midule_name"``` in this example "git mod init golang"
*/
func main() {
	fmt.Println("Hello, World!")
}
