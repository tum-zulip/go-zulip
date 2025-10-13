package models

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"unicode"
)

var modelDocSkipFiles = map[string]struct{}{
	"utils.go": {},
}

func TestModelsDocumentationCoverage(t *testing.T) {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("unable to determine current file path")
	}

	modelsDir := filepath.Dir(thisFile)
	docsDir := filepath.Clean(filepath.Join(modelsDir, "..", "..", "docs"))

	entries, err := os.ReadDir(modelsDir)
	if err != nil {
		t.Fatalf("reading models directory: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()
		if !strings.HasSuffix(filename, ".go") || strings.HasSuffix(filename, "_test.go") {
			continue
		}

		if _, skip := modelDocSkipFiles[filename]; skip {
			continue
		}

		expectedName := snakeToPascal(strings.TrimSuffix(filename, ".go"))

		docPath := filepath.Join(docsDir, expectedName+".md")
		if _, err := os.Stat(docPath); err != nil {
			t.Errorf("missing documentation for %s: expected %s", filename, docPath)
			continue
		}

		docTitle, err := readDocTitle(docPath)
		if err != nil {
			t.Errorf("reading documentation title from %s: %v", docPath, err)
		} else if docTitle != expectedName {
			t.Errorf("documentation title mismatch for %s: got %q, want %q", docPath, docTitle, expectedName)
		}

		hasType, err := fileHasType(filepath.Join(modelsDir, filename), expectedName)
		if err != nil {
			t.Errorf("parsing %s: %v", filename, err)
			continue
		}
		if !hasType {
			t.Errorf("expected type %s to be declared in %s", expectedName, filename)
		}
	}
}

func snakeToPascal(value string) string {
	parts := strings.Split(value, "_")
	var b strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}

		runes := []rune(part)
		first := unicode.ToUpper(runes[0])
		b.WriteRune(first)
		if len(runes) > 1 {
			b.WriteString(strings.ToLower(string(runes[1:])))
		}
	}
	return b.String()
}

func readDocTitle(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			return strings.TrimSpace(strings.TrimPrefix(line, "#")), nil
		}
		return "", fmt.Errorf("unexpected first non-empty line %q", line)
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("no title found in %s", path)
}

func fileHasType(path, expectedType string) (bool, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return false, err
	}

	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if typeSpec.Name.Name == expectedType {
				return true, nil
			}
		}
	}

	return false, nil
}
