package main

import (
	answercmd "github.com/Mark7631/incubator-answer/cmd"

  // remote plugins
	_ "github.com/apache/incubator-answer-plugins/connector-basic"

  // local plugins
)

func main() {
	answercmd.Main()
}
