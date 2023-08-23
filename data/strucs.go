package data

// package name should be the same as folder name //

type Recipe struct {
	// Recipe is the template name of this struc //
	Id          string
	RecipeName  string
	Source      string
	PrepTime    string
	CookTime    string
	ServingSize int
	Ingredients map[string]string
	// Ingredients is a map type with a string key to a string value //
	Instructions []string
	// Instructions is a slice of strings - slice as order of instruction matters //
	Tags []string
	// Tags is a slice of strings //
}
