package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAuth(t *testing.T) {
	mockToken := "aslkdjasd"
	got, _ := GetAPIKey(http.Header{"Authorization": {"ApiKey " + mockToken}})
	want := mockToken

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected %v, got: %v", want, got)
	}
}
