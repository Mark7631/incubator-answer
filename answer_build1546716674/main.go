package main

import (
	answercmd "github.com/apache/incubator-answer/cmd"

  // remote plugins
	_ "github.com/apache/incubator-answer-plugins/connector-basic"

  // local plugins
)

func main() {
	answercmd.Main()
}
