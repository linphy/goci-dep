package native_test

import (
	"fmt"
	"github.com/egravert/goci/native"
	"testing"
)

func TestCreateEnvironment(t *testing.T) {
	_, err := native.CreateEnvironment()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSuccessfulBasicLogin(t *testing.T) {
	env, _ := native.CreateEnvironment()
	_, err := native.BasicLogin(env, "hr", "oracle", "192.168.69.131/ORCL")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSFailedBasicLogin(t *testing.T) {
	env, _ := native.CreateEnvironment()
	_, err := native.BasicLogin(env, "boom", "fail", "192.168.69.131/ORCL")
	if err == nil {
		t.Fatal("Expected error, but received nil")
	}
}

func ExampleBasicLogin() {
	var env native.EnvHandle
	var err error

	if env, err = native.CreateEnvironment(); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := native.BasicLogin(env, "scott", "tiger", "127.0.0.1/ORCL"); err != nil {
		fmt.Println(err)
		return
	}
}

func ExamplePing() {
	env, _ := native.CreateEnvironment()
	svr, err := native.BasicLogin(env, "hr", "oracle", "192.168.69.131/ORCL")
	if err != nil {
		fmt.Println(err)
	}

	err = native.Ping(env, svr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success!")
	}
  // Output: Success!
}
