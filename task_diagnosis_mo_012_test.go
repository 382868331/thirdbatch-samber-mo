package mo

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisMo012SourceContract(t *testing.T) {
    source, err := os.ReadFile("either3.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return e.argId == either3ArgId3") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "return e.argId != either3ArgId3") {
        t.Fatalf("mutated source contract is still present")
    }
}
