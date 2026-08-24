package mo

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixMo014SourceContract(t *testing.T) {
    source, err := os.ReadFile("either5.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "switch int8(data[0]) {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "switch int8(data[1]) {") {
        t.Fatalf("mutated source contract is still present")
    }
}
