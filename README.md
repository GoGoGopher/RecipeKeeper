# RecipeKeeper 
Note: Default text will be the title of repository  
  
Can add a brief description, e.g. Project submitted as part of the Coding For Women course in Aug 2023. Steps provided below.  
  
## Step 1: Create repository in Github
  
## Step 2: In Visual Studio Code, open the desired folder where you want the repository to be located
File > Open folder... > <select folder/create folder> > Select folder  

Note: Do not put a repository within an existing repository!  
  
## Step 2: Clone repository in Visual Studio Code  
Terminal > New terminal > $ git clone <paste link to your repository>  

Pathway syntax: $ git clone <https://github.com/>/<github username>/<name of repository>.git  

Project: $ git clone https://github.com/GoGoGopher/RecipeKeeper.git  
  
## Step 3: Create new file called 'main.go' in repository folder and add 'package main' and 'func main(){}'
File > New file > main.go (or use shortcuts)  
  
## Step 4: Initialise module using 'go mod init' and the pathway of the repository
Terminal > New terminal > $ go mod init <path of repository>  

Pathway of repository syntax: $ go mod init <github path>/<github username>/<name of repository>  

Project: $ go mod init github.com/GoGoGopher/RecipeKeeper  
  
## Step 5: Tidy module using 'go mod tidy'
It's best practice to tidy the module. The command 'go mod tidy' adds missing requirements and drops unused requirements.  

Project: $ go mod tidy  
  
## Step 6: Push files to git respository
3 steps:  
    
a. Indicate which files have been edited  

Note: The full stop ( .) means all files/everything  

Terminal > New terminal > $ git add .  
  
b. Commit the changes  

Note: -m indicates message  

$ git commit -m "My first commit"  
  
c. Push the changes  

$ git pull

  
## Step 7: Create a 'handlers' folder
In the RecipeKeeper folder, create a new folder 'handlers' and create a new file within this folder called 'homepage.go'  

## Step 8: Create a 'handlers' package
Create the handlers package and import the 'fmt' and 'net/http' packages (fmt will just be used to check we've set up the handler correctly by printing to the browser).  

package handlers

import (
	"fmt"
	"net/http"
)

## Step 9: Create a 'HomePage' function
Create the 'HomePage' function. Note the captialisation - functions with names that start with an uppercase letter will be exported to other packages. 

Add the response and request parameters. Add an argument to print a string (to test in a browser).

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

## Step 10: Link to func main?
In main.go file:  
  
a. Import 'net/http' and and the handler package for RecipeKeeper    

package main

import (
	"net/http"
	"github.com/GoGoGopher/RecipeKeeper/handlers"
)

b. Add arguments to func main to respond and request

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.ListenAndServe(":8000", nil)
}

c. Save

## Step 11: Test

a. Run code ($ go run main.go)

b. Check if browser responds and requests by opening browser and seeking page 'http://localhost:8000/'

c. Page should load with 'Hello' printed






Remember to save and then run (go run main.go)
If a firewall pop up comes up allow main to execute (click allow)
In a browser go to localhost:8000
It should print your handler string (in this case "Hello")

ctrl c in terminal exits out of server so you can use terminal again
