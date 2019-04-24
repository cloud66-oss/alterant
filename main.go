package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/cloud66/alterant/lib"

	"github.com/gobuffalo/packr/v2"

	"github.com/ghodss/yaml"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
)

var (
	inputData []string
)

func main() {
	// flags
	inFile := flag.String("in", "", "input file")
	modifierFile := flag.String("modifier", "", "modifier file")

	flag.Parse()

	// read up
	input, err := ioutil.ReadFile(*inFile)
	if err != nil {
		fmt.Printf("in %s\n", err)
		os.Exit(1)
	}

	inputData = make([]string, 0)
	parts := strings.Split(string(input), "---")
	for idx, part := range parts {
		j, err := yaml.YAMLToJSON([]byte(part))
		if err != nil {
			fmt.Printf("unmarshal %d %s\n", idx, err)
			os.Exit(1)
		}

		inputData = append(inputData, string(j))
	}

	modifier, err := ioutil.ReadFile(*modifierFile)
	if err != nil {
		fmt.Printf("modifier: %s\n", err)
		os.Exit(1)
	}

	// get the static box setup
	box := packr.New("lib_js", "./lib_js")

	vm := otto.New()

	// load js libs
	libs := box.List()
	sort.Strings(libs)
	for _, libFile := range libs {
		classFile, err := box.FindString(libFile)
		if err != nil {
			fmt.Printf("containers %s\n", err)
			os.Exit(1)
		}

		// compile
		class, err := vm.Compile("", classFile)
		if err != nil {
			fmt.Printf("classes %s\n", err)
			os.Exit(1)
		}

		// run lib
		_, err = vm.Run(class)
		if err != nil {
			fmt.Printf("run classes %s\n", err.(*otto.Error).String())
			os.Exit(1)
		}
	}

	// load go libs
	err = vm.Set("JsonReader", func(call otto.FunctionCall) otto.Value {
		jsonFile, _ := call.Argument(0).ToString()
		jsonReader := lib.NewJsonReader(jsonFile)
		result, _ := vm.ToValue(jsonReader)

		return result
	})
	if err != nil {
		fmt.Printf("jsonReader %s\n", err.(*otto.Error).String())
		os.Exit(1)
	}

	err = vm.Set("YamlReader", func(call otto.FunctionCall) otto.Value {
		yamlFile, _ := call.Argument(0).ToString()
		yamlReader := lib.NewYamlReader(yamlFile)
		result, _ := vm.ToValue(yamlReader)

		return result
	})
	if err != nil {
		fmt.Printf("yamlReader %s\n", err.(*otto.Error).String())
		os.Exit(1)
	}

	// set the context
	_, err = vm.Object("$$ = [" + strings.Join(inputData, ",") + "]")
	if err != nil {
		fmt.Printf("set %s\n", err.(*otto.Error).String())
		os.Exit(1)
	}

	_, err = vm.Object("$$.replace = function(item) { replaceResource($$, item); }")
	if err != nil {
		fmt.Printf("yamlReader %s\n", err.(*otto.Error).String())
		os.Exit(1)
	}

	// compile the modifier
	script, err := vm.Compile(*modifierFile, modifier)
	if err != nil {
		fmt.Printf("compile %s\n", err)
		os.Exit(1)
	}

	// run
	_, err = vm.Run(script)
	if err != nil {
		fmt.Printf("run %s\n", err.(*otto.Error).String())
		os.Exit(1)
	}

	// get the result
	value, err := vm.Get("$$")
	if err != nil {
		fmt.Printf("get %s\n", err.(*otto.Error).String())
		os.Exit(1)
	}

	if value.Object().Class() != "Array" {
		fmt.Println("$$ is not an array")
		os.Exit(1)
	}

	result, err := value.Export()
	if err != nil {
		fmt.Printf("export %s\n", err)
		os.Exit(1)
	}

	toPrint, err := yaml.Marshal(result)
	if err != nil {
		fmt.Printf("marshal %s\n", err)
		os.Exit(1)
	}

	fmt.Println(string(toPrint))
}
