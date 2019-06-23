package login

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionIDRegex(t *testing.T) {
	val := "mhiodhmi5fmihldmifhli"
	line := "<input type=\"hidden\" name=\"SessionID\" value=\"" + val + "\">"
	matches := sessionIDRegex.FindStringSubmatch(line)
	assert.Len(t, matches, 2)
	assert.Equal(t, val, matches[1])
}
