package main

// Check for an error and if one exists, handle it
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
