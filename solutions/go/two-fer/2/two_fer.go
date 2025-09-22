// Package twofer is imprements a function that share
// a cookie with someone else.
package twofer

// ShareWith return a phrase: "One for X, one for me."
// If name is empty, it uses "you"
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."  
	}

	return "One for " + name + "," + " one for me."
}
