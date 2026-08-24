package mo

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixMo015SourceContract(t *testing.T) {
    source, err := os.ReadFile("either3.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "switch int8(data[0]) {") {
        t.Fatalf("expected source contract is missing")
    }
}
