package main

import "fmt"

func init() {
	fmt.Println("main init here...")
}

func main() {
	printStructure()
}

func printStructure() {
	fmt.Println(`Typical go project structure:
		projectname/
			/app
				/module
					main.go
			/internal -- only accessiable within the project
				/package1
				/package2
			/package1
			/package2
			/package3
			...
			go.mod -- mod file generated from go mod init command, for the package management
			go.sum -- checksum of dependency packages
	`)
}

func printCallSequence() {
	fmt.Println(`go invoking sequence:
	pkg3 -> pkg2 -> pkg1
	const -> var -> init()
	when main function in main package is invoked, find deepest dependency package (pkg3)
	invoke constant initialization, then var initialization, last is the init() function call in pkg3
	done initialization in pkg3, find the second deepest pacakge pkg2 and start initializaiton same as pkg3
	likewise, then same initialization sequence in pkg1, lastly initialization in main
	`)
}
