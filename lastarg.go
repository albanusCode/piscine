package main

import "os"

func LastArg()  {

	if len(os.Args) < 1 {
		os.Stdout.Write([]byte("\n"))
	}

	last := os.Args[len(os.Args) - 1]
	os.Stdout.Write([]byte(last + "\n"))

}
