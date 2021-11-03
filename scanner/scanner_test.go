package scanner

import (
	"bytes"
	"context"
	"io/fs"
	"regexp"
	"testing"
	"testing/fstest"

	"github.com/fatih/color"
	"github.com/google/go-cmp/cmp"
)

func TestScan(t *testing.T) {
	color.NoColor = true

	baseFS := fstest.MapFS{
		"path":             &fstest.MapFile{Mode: fs.ModeDir},
		"path/a.txt":       &fstest.MapFile{Data: []byte("alpha\n")},
		"path/to":          &fstest.MapFile{Mode: fs.ModeDir},
		"path/to/b.txt":    &fstest.MapFile{Data: []byte("beta\n")},
		"path/to/go":       &fstest.MapFile{Mode: fs.ModeDir},
		"path/to/go/c.txt": &fstest.MapFile{Data: []byte("camma\n")},
		"path/to/go/d.txt": &fstest.MapFile{Data: []byte("delta\nepsilon\n")},
	}

	tests := []struct {
		fsys    fs.FS
		opts    *Opts
		want    string
		wantErr bool
	}{
		{
			baseFS,
			&Opts{
				Patterns: []*regexp.Regexp{regexp.MustCompile("ta")},
				Owner:    "owner",
				Repo:     "repo",
				Include:  "**/*",
			},
			`owner/repo:path/to/b.txt:beta
owner/repo:path/to/go/d.txt:delta
`,
			false,
		},
		{
			baseFS,
			&Opts{
				Patterns: []*regexp.Regexp{regexp.MustCompile("ta")},
				Owner:    "owner",
				Repo:     "repo",
				Include:  "path/to/*",
			},
			`owner/repo:path/to/b.txt:beta
`,
			false,
		},
		{
			baseFS,
			&Opts{
				Patterns: []*regexp.Regexp{regexp.MustCompile("l")},
				Owner:    "owner",
				Repo:     "repo",
				Include:  "**/*",
				NameOnly: true,
			},
			`owner/repo:path/a.txt
owner/repo:path/to/go/d.txt
`,
			false,
		},
		{
			baseFS,
			&Opts{
				Patterns: []*regexp.Regexp{regexp.MustCompile("l")},
				Owner:    "owner",
				Repo:     "repo",
				Include:  "**/*",
				RepoOnly: true,
			},
			"owner/repo\n",
			true,
		},
		{
			baseFS,
			&Opts{
				Patterns:   []*regexp.Regexp{regexp.MustCompile("l")},
				Owner:      "owner",
				Repo:       "repo",
				Include:    "**/*",
				LineNumber: true,
			},
			`owner/repo:path/a.txt:1:alpha
owner/repo:path/to/go/d.txt:1:delta
owner/repo:path/to/go/d.txt:2:epsilon
`,
			false,
		},
		{
			baseFS,
			&Opts{
				Patterns: []*regexp.Regexp{regexp.MustCompile("l")},
				Owner:    "owner",
				Repo:     "repo",
				Include:  "**/*",
				Count:    true,
			},
			`owner/repo:path/a.txt:1
owner/repo:path/to/go/d.txt:2
`,
			false,
		},
		{
			baseFS,
			&Opts{
				Patterns: []*regexp.Regexp{regexp.MustCompile("l")},
				Owner:    "owner",
				Repo:     "repo",
				Include:  "**/*",
				Exclude:  "path/*.txt",
			},
			`owner/repo:path/to/go/d.txt:delta
owner/repo:path/to/go/d.txt:epsilon
`,
			false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		buf := new(bytes.Buffer)
		if err := Scan(ctx, tt.fsys, buf, tt.opts); err != nil {
			if !tt.wantErr {
				t.Error(err)
			}
		}
		got := buf.String()
		if diff := cmp.Diff(got, tt.want, nil); diff != "" {
			t.Errorf("%s", diff)
		}
	}
}
