package main

import "fmt"

type animal interface{
	getSound() string
}

type dog struct{}
type cat struct{}
func main(){
	d := dog{}
	c := cat{}

	printSound(d)
	printSound(c)
}

func printSound(a animal){
	fmt.Println(a.getSound())
}

// Honorary member of animal due to method implementation
func (dog) getSound() string{
	return "Bark"
}

func (cat) getSound() string{
	return "Meow"
}