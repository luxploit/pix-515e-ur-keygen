package cmd

import (
	"flag"
	"fmt"
	"os"
	"reflect"
)

type Flags struct {
	flagSet   *flag.FlagSet
	flagOrder []string
}

func NewFlags(name string) *Flags {
	fl := &Flags{flagSet: flag.NewFlagSet(name, flag.ExitOnError)}

	fl.flagSet.Usage = func() {
		for _, name := range fl.flagOrder {
			_flag := fl.flagSet.Lookup(name)
			fmt.Printf("\t--%-20s\t%-50s (Default: %s)\n", fmt.Sprintf("%-s <%-s>", _flag.Name, reflect.TypeOf(_flag.Value).Elem().Kind().String()), _flag.Usage, _flag.DefValue)
		}
	}

	return fl
}

func (f *Flags) Bool(name string, defaultVal bool, helpText string) *bool {
	f.flagOrder = append(f.flagOrder, name)
	return f.flagSet.Bool(name, defaultVal, helpText)
}

func (f *Flags) Int64(name string, defaultVal int64, helpText string) *int64 {
	f.flagOrder = append(f.flagOrder, name)
	return f.flagSet.Int64(name, defaultVal, helpText)
}

func (f *Flags) String(name string, defaultVal string, helpText string) *string {
	f.flagOrder = append(f.flagOrder, name)
	return f.flagSet.String(name, defaultVal, helpText)
}

func (f *Flags) Usage() {
	f.flagSet.Usage()
}

func (f *Flags) Parse() {
	f.flagSet.Parse(os.Args[2:])
}
