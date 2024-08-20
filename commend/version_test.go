package commend

import "testing"

func TestGetVersion(t *testing.T) {
	response := GetVersion()
	if "" == response {
		t.Error("getVersion error!")
	}

	if "1.0.0" != response {
		t.Fail()
	}
}
