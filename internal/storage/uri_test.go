package storage

import (
	"runtime"
	"testing"
)

func TestFileURIKeepsDrivePathsOutOfTheAuthority(t *testing.T) {
	cases := map[string]string{
		"/var/lib/sekisho/sekisho.db": "file:///var/lib/sekisho/sekisho.db",
		"C:/data/my data.db":          "file:///C:/data/my%20data.db",
	}
	if runtime.GOOS == "windows" {
		cases[`C:\data\sekisho.db`] = "file:///C:/data/sekisho.db"
	}
	for path, want := range cases {
		u := FileURI(path)
		if got := u.String(); got != want {
			t.Errorf("FileURI(%q) = %q, want %q", path, got, want)
		}
		if u.Host != "" {
			t.Errorf("FileURI(%q) has authority %q", path, u.Host)
		}
	}
}
