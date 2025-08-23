package command 

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	builder := New("test").
		Arg("-a").
		Args("-b", "-c").
		Build()
		assert.Equal(t, []string{"test","-a", "-b", "-c"}, builder.Args)
}
