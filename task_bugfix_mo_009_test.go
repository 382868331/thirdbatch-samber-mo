package mo

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixMo009SourceContract(t *testing.T) {
    source, err := os.ReadFile("either.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if data[0] == 1 {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if data[1] == 1 {") {
        t.Fatalf("mutated source contract is still present")
    }
}
