package main

import "MiMengCore/initialize"

func main() {
	//initialize.InitMySQL()
	r := initialize.SetupRouter()
	r.Run(":8080")
}
