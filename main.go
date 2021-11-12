package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/potproject/goenvgen/gen"
)

var (
	pkg   = flag.String("p", "envgen", "the name of the package for the generated code")
	types = flag.String("t", "", "Manually type definition setting (example: \"-t ENV_BOOL=bool,ENV_S_INT=[]int\")")
)

func main() {
	flag.Parse()
	err := run(flag.Args())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}
	os.Exit(0)
}

func run(args []string) error {
	file := ".env"
	if len(args) == 1 {
		file = args[0]
	}
	if len(args) > 1 {
		return errors.New(".env File is not specified")
	}
	mapmtds := map[string]string{}
	if *types != "" {
		mtds := strings.Split(*types, ",")
		for _, v := range mtds {
			mtd := strings.Split(v, "=")
			if len(mtd) != 2 {
				return errors.New("Manually type definition invalid value: " + v)
			}
			mapmtds[mtd[0]] = mtd[1]
		}
	}
	err := gen.GenerateFile(file, *pkg, mapmtds)
	return err
}
