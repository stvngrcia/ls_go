package main

import (
	"os"
	"fmt"
	"strings"
)

/**
 *
 *
 */
func main(){
	var options, path string
	args := os.Args[1:]

	if len(args) >= 1 && args[0][0] == '-'{
		/*Parse the options*/
		options, path = get_options(args)
	}
	if (len(path) == 0){
		path = "."
	}
	fmt.Println(options)
	fmt.Println(path)
}


/**
 * get_options: Separates the option flags from the file path.
 * @args: Array string containing all the arguments.
 * Returns: Two strings first represents the options
 * and the second represents the path.
 */
func get_options(args []string)(string, string){
	i := 0
	my_string := strings.Join(args, " ")
	for len(my_string) > i && my_string[i] != ' '{
		i++
	}
	return my_string[:i], my_string[i:]
}
