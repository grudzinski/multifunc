package multifunc

import (
	"errors"
	"testing"
)

func TestResult(t *testing.T) {
	mf := new(MultiFunc)
	resFunc := mf.Add(func() (interface{}, error) {
		return "test", nil
	})
	if err := mf.Run(); err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}
	if resFunc().(string) != "test" {
		t.Fatal()
	}
}

func TestErr(t *testing.T) {
	mf := new(MultiFunc)
	mf.Add(func() (interface{}, error) {
		return nil, errors.New("test")
	})
	if err := mf.Run(); err == nil {
		t.Fatal()
	}
}
