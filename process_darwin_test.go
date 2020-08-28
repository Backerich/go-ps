// +build darwin

package ps

import (
	"fmt"
	"testing"
)

func TestDarwinProcess_impl(t *testing.T) {
	var _ Process = new(DarwinProcess)
}

func TestListProcesses(t *testing.T) {
	p, err := Processes()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if len(p) <= 0 {
		t.Fatal("should have processes")
	}

	for _, p1 := range p {
		fmt.Println(p1.Executable())
	}

	fmt.Println("Number of processes running: ", len(p))
}
