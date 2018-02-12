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

func TestNoResult(t *testing.T) {
	mf := new(MultiFunc)
	mf.AddNoResult(func() error {
		return nil
	})
	if err := mf.Run(); err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}
}

func TestErr(t *testing.T) {
	mf := new(MultiFunc)
	mf.Add(func() (interface{}, error) {
		return nil, errors.New("test")
	})
	err := mf.Run()
	if err == nil {
		t.Fatal()
	}
	if err.Error() != "test" {
		t.Fatal()
	}
}

func TestNoResultErr(t *testing.T) {
	mf := new(MultiFunc)
	mf.AddNoResult(func() error {
		return errors.New("test")
	})
	err := mf.Run()
	if err == nil {
		t.Fatal()
	}
	if err.Error() != "test" {
		t.Fatal()
	}
}
