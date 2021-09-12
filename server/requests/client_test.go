package requests

import (
	"testing"
	"reflect"
)

func Test_NewRequest(t *testing.T) {
	t.Parallel()

	tests := []struct{
		name string
		expectedHeaderTypes []reflect.Kind
		expectedRootType reflect.Kind
		inputHeaders []string
		inputRootURL string
	}{
		{
			name: "correctHeaders",
			expectedHeaderTypes: []reflect.Kind{
				reflect.String, reflect.String,
			},
			expectedRootType: reflect.String,
			inputHeaders: []string{
				"testHost",
				"testKey",
			},
			inputRootURL: "testRootURL",
		},
	}

	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			rx := NewRequest(ts.inputHeaders[0], ts.inputHeaders[1], ts.inputRootURL)
			
			uxT := reflect.ValueOf(rx.RootURL)
			if uxT.Kind() != ts.expectedRootType {
				t.Error("wrong type")
			}
		})
	}
}