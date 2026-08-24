package mo

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixMo006SourceContract(t *testing.T) {
    source, err := os.ReadFile("either4.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err := enc.Encode(e.arg2); err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
