package mo

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixMo018SourceContract(t *testing.T) {
    source, err := os.ReadFile("future.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "<-f.done") {
        t.Fatalf("expected source contract is missing")
    }
}
