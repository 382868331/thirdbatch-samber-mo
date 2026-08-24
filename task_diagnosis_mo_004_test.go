package mo

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisMo004SourceContract(t *testing.T) {
    source, err := os.ReadFile("option.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if data[0] == 0 {") {
        t.Fatalf("expected source contract is missing")
    }
}
