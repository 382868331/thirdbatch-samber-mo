package mo

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixMo013SourceContract(t *testing.T) {
    source, err := os.ReadFile("either4.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(data) == 0 {") {
        t.Fatalf("expected source contract is missing")
    }
}
