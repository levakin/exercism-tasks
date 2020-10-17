// Package twofer provides function to print "One for X, one for me."
package twofer

import "fmt"

// ShareWith given a name, returns a string with the message:
//"One for X, one for me."
//Where X is the given name.
//However, if the name is missing, return the string:
//One for you, one for me.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}
