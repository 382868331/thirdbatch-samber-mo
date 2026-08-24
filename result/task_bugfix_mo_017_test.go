package result

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixMo017SourceContract(t *testing.T) {
    source, err := os.ReadFile("transforms.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
