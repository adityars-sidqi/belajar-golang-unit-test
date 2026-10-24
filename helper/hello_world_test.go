package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Aditya")
	}
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("Before unit test")

	m.Run()

	//after
	fmt.Println("After unit test")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Aditya")
	require.Equal(t, "Hello Aditya", result, "they should be equal")
	fmt.Println("TestHelloWorld with Require done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Aditya")
	assert.Equal(t, "Hello Aditya", result, "they should be equal")
	fmt.Println("TestHelloWord with Assert done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Aditya")
	if result != "Hello Aditya" {
		t.Error("Expected Hello Aditya, got", result)
	}

	fmt.Println("Test HelloWorld done")
}

func TestHelloWorldRahman(t *testing.T) {
	result := HelloWorld("Rahman")
	if result != "Hello Rahman" {
		t.Fatal("Expected Hello Rahman, got", result)
	}

	fmt.Println("Test HelloWorldRahman done")
}
