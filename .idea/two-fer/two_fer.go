package twofer

import "fmt"

// ShareWith This function runs a conditional statement and performs string formatting
func ShareWith(name string) string {
	if name != "" {
		return fmt.Sprintf("One for %s, one for me.", name)
	}
	return fmt.Sprintf("One for you, one for me.")
}

func main() {
	var name string
	fmt.Println("Enter your Name: ")
	fmt.Scanln(&name)
	fmt.Println(ShareWith(name))
}