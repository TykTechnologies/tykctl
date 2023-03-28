package shared

import (
	"github.com/spf13/pflag"
)

func AddOutPutFlags(f *pflag.FlagSet) {
	f.StringP(OutPut, "o", "json", "Format you want to use can be table,json")
}
