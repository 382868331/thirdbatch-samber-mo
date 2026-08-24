package mo

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixMo010SourceContract(t *testing.T) {
    source, err := os.ReadFile("option.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if av, err := driver.DefaultParameterConverter.ConvertValue(src); err == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if false && av, err := driver.DefaultParameterConverter.ConvertValue(src); err == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
