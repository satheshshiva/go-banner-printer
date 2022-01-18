package banner

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

/**
To use this package, just call 'banner.Print(nil)' in the init() function.
When you have a banner.txt in your root directory, it will get printed.
If you want a different file path/name or you if you dont want to print env, then pass in the struct with the config
*/

var fileName = "banner.txt"
var printEnv = true

type Options struct {
	FileName string //overrides the filepath+filename
	PrintEnv bool   //prints out the env variables
}

// Print will print the banner file present in root directory with name banner.txt
func Print(options *Options) error {
	initOptions(options)

	s, err := ioutil.ReadFile(fileName)
	if err == nil {
		fmt.Println(string(s))
	} else {
		return err
	}

	if printEnv {
		printAllEnv()
	}
	return nil
}

func initOptions(options *Options) {
	if options != nil {
		if options.FileName != "" {
			fileName = options.FileName
		}
		printEnv = options.PrintEnv
	}
}

func printAllEnv() {
	fmt.Println()
	envVars := []string{"PORT", "GOPATH"}
	for _, envVar := range envVars {
		envVarVal, present := os.LookupEnv(envVar)
		if present {
			fmt.Printf("%s: %s\n", envVar, envVarVal)
		}
	}
	host, _ := os.Hostname()
	fmt.Printf("HOST: %s\n", host)
	fmt.Printf("GO Version: %s\n", runtime.Version())
	fmt.Printf("GO Root: %s\n", runtime.GOROOT())
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	execPath, _ := os.Executable()
	fmt.Printf("App Path: %s\n", filepath.Dir(execPath))
	fmt.Println()
}
