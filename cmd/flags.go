package cmd

import (
	"github.com/spf13/pflag"
	"reflect"
)

func AddOrgFlag(f *pflag.FlagSet) {
	f.String("org", "", "The organization")
}

type Flag struct {
	Name               string
	Shorthand          string
	Usage              string
	Value              interface{}
	DefValue           interface{}
	FlagAddMethod      string
	Deprecated         string
	DefinedOn          []string
	Hidden             bool
	IsEnum             bool
	DefValuePerCommand map[string]interface{}
	ViperBindValue     string
}

///This will store all the common flags that we use in the cli
///When you add a new flag to this registry make sure that you define
///on which commands the flag is applicable  in the DefineOn slice by adding the commands key.

var orgFlag = Flag{
	Name:           "org",
	Shorthand:      "o",
	Usage:          "The user organization",
	Value:          "",
	DefValue:       "",
	FlagAddMethod:  "StringVar",
	DefinedOn:      []string{"all"},
	ViperBindValue: "org",
}

func (fl *Flag) flag() *pflag.Flag {
	methodName := fl.FlagAddMethod
	if methodName == "" {
		methodName = methodNameByType(reflect.ValueOf(fl.Value))
	}

	isVar := methodName == "Var"
	if isVar {

	}
	inputs := []interface{}{fl.Value, fl.Name}
	if !isVar {
		inputs = append(inputs, fl.DefValue)
	}
	inputs = append(inputs, fl.Usage)
	fs := pflag.NewFlagSet(fl.Name, pflag.ContinueOnError)
	reflect.ValueOf(fs).MethodByName(methodName).Call(reflectValueOf(inputs))
	f := fs.Lookup(fl.Name)
	f.Shorthand = fl.Shorthand
	return f

}

func methodNameByType(v reflect.Value) string {
	t := v.Type().Kind()
	switch t {
	case reflect.Bool:
		return "BoolVar"
	case reflect.Int:
		return "IntVar"
	case reflect.String:
		return "StringVar"
	case reflect.Slice:
		return "StringSliceVar"
	case reflect.Struct:
		return "Var"
	case reflect.Ptr:
		return methodNameByType(reflect.Indirect(v))
	}
	return ""
}

func reflectValueOf(values []interface{}) []reflect.Value {
	var results []reflect.Value
	for _, v := range values {
		results = append(results, reflect.ValueOf(v))
	}
	return results
}
