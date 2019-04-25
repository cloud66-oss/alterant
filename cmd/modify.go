package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/ghodss/yaml"
	"github.com/gobuffalo/packr/v2"
	"github.com/khash/alterant/lib"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore" // this imports underscore into otto
	"github.com/spf13/cobra"
)

var modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Run transformer script on the given input",
	Run:   execModify,
}

var (
	inFile  string
	modFile string
	box     *packr.Box
)

func init() {
	// get the static box setup
	box = packr.New("lib_js", "./lib_js")

	modifyCmd.Flags().StringVar(&inFile, "in", "", "input file (could be json or yaml)")
	modifyCmd.Flags().StringVar(&modFile, "modifier", "", "modifier file (javascript)")

	rootCmd.AddCommand(modifyCmd)
}

func execModify(cmd *cobra.Command, args []string) {
	inputData, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	modifier, err := ioutil.ReadFile(modFile)
	if err != nil {
		log.Fatal(err)
	}

	vm := otto.New()

	if err = loadJSLib(vm); err != nil {
		logError(err)
	}

	if err = loadGoLib(vm); err != nil {
		logError(err)
	}

	if err = loadGlobals(inputData, vm); err != nil {
		logError(err)
	}

	// compile the modifier
	script, err := vm.Compile(modFile, modifier)
	if err != nil {
		log.Fatal(err)
	}

	vm.Interrupt = make(chan func(), 1)
	go func() {
		time.Sleep(100 * time.Microsecond)
		vm.Interrupt <- func() {
			panic("timeout")
		}
	}()
	// run
	if _, err = vm.Run(script); err != nil {
		logError(err)
	}

	// get the result
	result, err := fetchResult(vm)
	if err != nil {
		logError(err)
	}

	toPrint, err := yaml.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(toPrint))
}

func readInput() ([]string, error) {
	// read up
	input, err := ioutil.ReadFile(modFile)
	if err != nil {
		log.Fatalf("in %s\n", err)
	}

	inputData := make([]string, 0)
	parts := strings.Split(string(input), "---")
	for idx, part := range parts {
		j, err := yaml.YAMLToJSON([]byte(part))
		if err != nil {
			return nil, fmt.Errorf("%s in part %d", err, idx)
		}

		inputData = append(inputData, string(j))
	}

	return inputData, nil
}

func loadJSLib(vm *otto.Otto) error {
	// load js libs
	libs := box.List()
	sort.Strings(libs)
	for _, libFile := range libs {
		classFile, err := box.FindString(libFile)
		if err != nil {
			return err
		}

		// compile
		class, err := vm.Compile("", classFile)
		if err != nil {
			return err
		}

		// run lib
		_, err = vm.Run(class)
		if err != nil {
			return err
		}
	}

	return nil
}

func loadGoLib(vm *otto.Otto) error {
	err := vm.Set("JsonReader", func(call otto.FunctionCall) otto.Value {
		jsonFile, _ := call.Argument(0).ToString()
		jsonReader := lib.NewJsonReader(jsonFile)
		result, _ := vm.ToValue(jsonReader)

		return result
	})
	if err != nil {
		return err
	}

	err = vm.Set("YamlReader", func(call otto.FunctionCall) otto.Value {
		yamlFile, _ := call.Argument(0).ToString()
		yamlReader := lib.NewYamlReader(yamlFile)
		result, _ := vm.ToValue(yamlReader)

		return result
	})
	if err != nil {
		return err
	}

	return nil
}

func loadGlobals(inputData []string, vm *otto.Otto) error {
	_, err := vm.Object("$$ = [" + strings.Join(inputData, ",") + "]")
	if err != nil {
		return err
	}

	_, err = vm.Object("$$.replace = function(item) { replaceResource($$, item); }")
	if err != nil {
		return err
	}

	return nil
}

func logError(err error) {
	if t, ok := err.(*otto.Error); ok {
		log.Fatal(t.String())
	} else {
		log.Fatal(err)
	}
}

func fetchResult(vm *otto.Otto) (interface{}, error) {
	// get the result
	value, err := vm.Get("$$")
	if err != nil {
		logError(err)
	}

	if value.Object().Class() != "Array" {
		log.Fatalf("$$ is not an array")
	}

	return value.Export()
}
