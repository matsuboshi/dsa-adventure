package main
import "fmt"

type Key [26]int
var names = []string{ 
	"mitch", "tchim", "lucas", "lusac", "caslu",
}

func main() {
	anMap := make(map[Key][]string)
	
	for _, name := range names {	
		var key Key
		for i := range name {
			key[name[i]-'a']++
			
		}
		anMap[key] = append(anMap[key], name)
	}

	for _, v := range anMap {
		fmt.Println(v)
	}
}
