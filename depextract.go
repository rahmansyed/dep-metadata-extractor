package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type manifest struct {
	Metadata map[string]interface{} `toml:"Metadata,omitempty"`
}

type keylist []string

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (i *keylist) String() string {
	return fmt.Sprint(*i)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (i *keylist) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, k := range strings.Split(value, ",") {
		*i = append(*i, k)
	}
	return nil
}

func main() {
	file := flag.String("file", "Gopkg.toml", "path to the dep manifest.")
	isKeyDisplayed := flag.Bool("withkey", true, "output with keys or not")

	var keys keylist
	flag.Var(&keys, "keys", "comma-separated list of keys to extract")

	flag.Parse()

	var m manifest
	d := getDataFromFile(*file)
	if _, err := toml.Decode(d, &m); err != nil {
		fmt.Errorf("%v", err)
		os.Exit(1)
	}

	if len(keys) == 0 {
		for k := range m.Metadata {
			keys = append(keys, k)
		}
	}

	fmt.Print(queryManifest(m, keys, *isKeyDisplayed))
}

func getDataFromFile(filePath string) string {
	b, err := ioutil.ReadFile(filePath) // just pass the file name
	if err != nil {
		fmt.Print(err)
		return ""
	}

	return string(b)
}

func queryManifest(data manifest, keys keylist, isKeyDisplayed bool) string {
	var b strings.Builder

	for _, key := range keys {
		if isKeyDisplayed {
			b.WriteString(fmt.Sprintf("%s=%s\n", key, data.Metadata[key]))
		} else {
			b.WriteString(fmt.Sprintf("%s\n", data.Metadata[key]))
		}
	}

	return b.String()
}
