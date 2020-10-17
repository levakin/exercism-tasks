// Package proverb implements generation of the relevant proverb
package proverb

import "fmt"

// Proverb generates the relevant proverb
// Given a list of inputs, generate the relevant proverb.
// For example, given the list ["nail", "shoe", "horse", "rider", "message", "battle", "kingdom"],
// you will output the full text of this proverbial rhyme:
// For want of a nail the shoe was lost.
//For want of a shoe the horse was lost.
//For want of a horse the rider was lost.
//For want of a rider the message was lost.
//For want of a message the battle was lost.
//For want of a battle the kingdom was lost.
//And all for the want of a nail.
func Proverb(rhyme []string) []string {
	p := make([]string, len(rhyme))
	if len(rhyme) > 1 {
		for i := 0; i < len(rhyme)-1; i++ {
			p[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		}
	}
	if len(rhyme) > 0 {
		p[len(rhyme)-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	}
	return p
}
