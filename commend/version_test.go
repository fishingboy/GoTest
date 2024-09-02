package commend

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVersion(t *testing.T) {
	response := GetVersion()
	if "" == response {
		t.Error("getVersion error!")
	}

	if "1.0.0" != response {
		t.Fail()
	}
}

func TestGetVersion2(t *testing.T) {
	response := GetVersion()
	assert.Equal(t, "1.0.0", response)
}
