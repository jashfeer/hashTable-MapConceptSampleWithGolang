package main

import "fmt"



func main(){
var m map[string]string
fmt.Println(m==nil)
m = map[string]string{}
fmt.Println(m==nil)
fmt.Println(m)
fmt.Println(len(m))
m=map[string]string{"wallace":"programmer"}
fmt.Println(m)

m["jonny"]="not a programer"
fmt.Println(m)

fmt.Println(m["jonny"])

for name,title:=range m{
	fmt.Println(name,title)
}

}