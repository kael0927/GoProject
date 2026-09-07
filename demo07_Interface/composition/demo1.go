package main
import "fmt"
type author struct{
	firstName string
	lastName string
}
func (a author) fullName() string {
	return fmt.Sprintf("%s %s",a.firstName,a.lastName)
}

type post struct{
	title string
	author
}

func (p post) details(){
	fmt.Println(p.author.fullName())
}
func main(){
	author1 := author{
		"lue",
		"ksj",
	}
	post1 := post{
		"kk",
		author1,
	}
	fmt.Println(post1.author.fullName())
}