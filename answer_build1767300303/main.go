package main

import (
	answercmd "github.com/apache/incubator-answer/cmd"

  // remote plugins
	_ "github.com/apache/incubator-answer-plugins/connector-basic"
	_ "github.com/apache/incubator-answer-plugins/connector-github"

  // local plugins
)

func main() {
	answercmd.Main()
}
