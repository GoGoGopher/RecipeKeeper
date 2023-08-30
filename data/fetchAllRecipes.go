package data

import (
	"encoding/json" // Convert bytes into... //
	"io"            // Read and write files - read json info and convert it into a readable slice of bytes //
	"log"
	"net/http" // Make code into website //
)

func FetchAllRecipes() {
	jsonInfo, err := http.Get("https://api.npoint.io/fff0f131782057b16a12") // http.Get reads (retreives) a representation of a resource (in this case a representation in JSON); Create variable err to return error //
	if err != nil {
		log.Fatal(err, "ËRROR: problem with url") // Kill code and give error message
	}

	defer jsonInfo.Body.Close() // When you open something, you need to close it so you don't overwhelm the computer memory //

	body, err := io.ReadAll(jsonInfo.Body) // takes data and converts to a[]byte...

	if err != nil {
		log.Fatal(err, "ERROR: problem reading data") // Error message to yourself
	}

	json.Unmarshal(body, &AllRecipes) // Takes JSON data to be unmarshalled and... Find this variables location and store the data there in the strucs as told to by the tags
}
