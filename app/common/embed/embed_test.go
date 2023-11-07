package embed

import "testing"

func TestEmbed(t *testing.T) {
	entries, err := TLFiles.ReadDir("tlfiles")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		t.Logf("Name: %s, IsDir: %t", entry.Name(), entry.IsDir())
		file, err := TLFiles.ReadFile("tlfiles/" + entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("File: %s", file)
	}
}
