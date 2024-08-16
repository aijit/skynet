package main

import "github.com/aijit/skynet/config"

func main() {
	config.GetConfig().Load()
}
