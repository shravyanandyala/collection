package template_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gojini.dev/template"
)

func TestTemplate(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)
	assert.Equal(0, template.Noop())
}
