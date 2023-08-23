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
$ git push
