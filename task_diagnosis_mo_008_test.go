package mo

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisMo008SourceContract(t *testing.T) {
    source, err := os.ReadFile("either4.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return e.argId == either4ArgId3") {
        t.Fatalf("expected source contract is missing")
    }
}
