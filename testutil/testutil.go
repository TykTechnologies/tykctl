package testutil

import (
	"github.com/TykTechnologies/tykctl/internal"
	flag "github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlags(t *testing.T, f *flag.FlagSet, flags []internal.Flag) {
	t.Helper()
	for _, tt := range flags {
		t.Run(tt.Description, func(t *testing.T) {
			l := f.Lookup(tt.Name)
			assert.NotNil(t, l)
			if l != nil {
				assert.Equal(t, tt.Value, l.Value.String(), tt.Description)
				assert.Equal(t, tt.Shorthand, l.Shorthand, tt.Description)
				assert.Equal(t, tt.Default, l.DefValue, tt.Description)
			}

		})

	}
}
