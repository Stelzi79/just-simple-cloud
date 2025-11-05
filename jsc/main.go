package main

import "github.com/Stelzi79/just-simple-cloud/cmd"

func main() {
	cmd.Execute()

	// exec.Command("echo Hello, World!").Run()
	// shellCmd := exec.Command("docker", "ps", "-a", "--format", "json")
	// output, err := shellCmd.Output()

	// if err != nil {
	// 	panic(err)
	// }

	// pp.Println(string(output))

}
