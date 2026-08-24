package mo

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixMo002SourceContract(t *testing.T) {
    source, err := os.ReadFile("option.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if o.isPresent != other.isPresent {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if o.isPresent == other.isPresent {") {
        t.Fatalf("mutated source contract is still present")
    }
}
