// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

func ShareWith(name string) string {
    msg := "One for you, one for me."
	if name == "Alice" {
        msg = "One for Alice, one for me."
    }
    if name == "Bob" {
        msg = "One for Bob, one for me."
    } 
    if name == "Zaphod" {
        msg = "One for Zaphod, one for me."
    }
	return msg
}
