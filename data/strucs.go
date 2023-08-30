package data

// package name should be the same as folder name //

type Recipe struct {
	// Recipe is the template name of this struc //
	// Add
	Id          string            `json:"Id"` // JSON tags for key name do not need to be the same as your specified for name e.g. PrepTime vs json: "Preperation Time"
	RecipeName  string            `json:"Recipe Name"`
	Source      string            `json:"Source"`
	PrepTime    string            `json:"Preperation Time"` // Mispelt in JSON
	CookTime    string            `json:"Cook Time"`
	ServingSize int               `json:"Serving Size"`
	Ingredients map[string]string `json:"Ingredients"`
	// Ingredients is a map type with a string key to a string value //
	Instructions []string `json:"Directions"`
	// Instructions is a slice of strings - slice as order of instruction matters //
	Tags []string `json:"Tags"`
	// Tags is a slice of strings //
}

var AllRecipes []Recipe // var AllRecipes can hold all recipes, this code is us creating a slice of recipe
