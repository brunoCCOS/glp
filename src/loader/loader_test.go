package loader

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDefault(t *testing.T) {
	gens, err := LoadDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(gens) == 0 {
		t.Fatal("expected embedded generators")
	}
	want := map[string]bool{"mult": false, "copy": false}
	for _, g := range gens {
		if _, ok := want[g.Value]; ok {
			want[g.Value] = true
		}
	}
	for k, found := range want {
		if !found {
			t.Errorf("missing default generator %q", k)
		}
	}
}

func TestLoadFile(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "g.json")
	if err := os.WriteFile(p, []byte(`{"generators":[{"name":"f","arity":"1","coarity":"1"}]}`), 0o600); err != nil {
		t.Fatal(err)
	}
	gens, err := LoadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	if len(gens) != 1 || gens[0].Value != "f" || gens[0].Arity != "1" || gens[0].Coarity != "1" {
		t.Fatalf("got %+v", gens)
	}
}

func TestLoadFileDuplicate(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "g.json")
	body := `{"generators":[{"name":"f","arity":"1","coarity":"1"},{"name":"f","arity":"2","coarity":"2"}]}`
	if err := os.WriteFile(p, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadFile(p); err == nil {
		t.Fatal("expected duplicate error")
	}
}

func TestLoadFileEmptyName(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "g.json")
	if err := os.WriteFile(p, []byte(`{"generators":[{"arity":"1","coarity":"1"}]}`), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadFile(p); err == nil {
		t.Fatal("expected empty-name error")
	}
}

func TestLoadFileMissing(t *testing.T) {
	if _, err := LoadFile(filepath.Join(t.TempDir(), "nope.json")); err == nil {
		t.Fatal("expected error")
	}
}
