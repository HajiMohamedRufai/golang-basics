# Golang Basics
## Conventionally follow this:
- Initialize golang module in indepedent golang project folders
```
# go mod init <module-name>
 go mod init my-project
```

- In Golang folder where you initialize the module it is called the main package (like a folder)

## Main package
- The package takes the name of the initiliazed module, not the parent folder
- It has conventionally 1 go file named **main.go** which serves as the entry point of the program
- Main.go file has a main function
- Go project has only 1 main function

  The subfolders can contain numerous go files
  
  
