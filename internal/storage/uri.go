package storage

import (
	"net/url"
	"path/filepath"
	"strings"
)

// FileURI returns a SQLite file: URI for an absolute path. Windows paths such as
// C:\data\sekisho.db need a leading slash, otherwise SQLite reads "C:" as a URI
// authority and refuses to open the database.
func FileURI(absolute string) url.URL {
	path := filepath.ToSlash(absolute)
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return url.URL{Scheme: "file", Path: path}
}
