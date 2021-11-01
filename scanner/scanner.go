package scanner

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"regexp"
	"sort"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/fatih/color"
	"github.com/k1LoW/gh-grep/internal"
)

var (
	matchc    = color.New(color.FgRed, color.Bold)
	delimiter = color.New(color.FgCyan).Sprint(":")
)

type Opts struct {
	Patterns   []*regexp.Regexp
	Owner      string
	Repo       string
	Include    string
	Exclude    string
	LineNumber bool
}

func Scan(ctx context.Context, fsys fs.FS, w io.Writer, opts *Opts) error {
	return doublestar.GlobWalk(fsys, opts.Include, func(path string, d fs.DirEntry) error {
		if d.IsDir() {
			return nil
		}
		if opts.Exclude != "" {
			match, err := doublestar.PathMatch(opts.Exclude, path)
			if err != nil {
				return err
			}
			if match {
				log.Printf("Exclude %s\n", path)
				return nil
			}
		}
		log.Printf("Search %s\n", path)
		f, err := fsys.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		// TODO: detect encoding
		fscanner := bufio.NewScanner(f)
		n := 1
		for fscanner.Scan() {
			line := fscanner.Text()

			matches := [][]int{}
			for _, re := range opts.Patterns {
				matches = append(matches, re.FindAllStringIndex(line, -1)...)
			}

			sort.Slice(matches, func(i, j int) bool {
				return matches[i][0] > matches[j][0]
			})

			f := [][]int{}
			for i := range matches {
				if i+1 == len(matches) || matches[i][1] < matches[i+1][0] {
					f = append(f, matches[i])
				}
			}
			matches = f

			if len(matches) > 0 {
				if opts.LineNumber {
					if _, err := fmt.Fprintf(w, "%s/%s%s%s%s%d%s%s\n", opts.Owner, opts.Repo, delimiter, path, delimiter, n, delimiter, internal.PrintLine(line, matches, matchc)); err != nil {
						return err
					}
				} else {
					if _, err := fmt.Fprintf(w, "%s/%s%s%s%s%s\n", opts.Owner, opts.Repo, delimiter, path, delimiter, internal.PrintLine(line, matches, matchc)); err != nil {
						return err
					}
				}
			}
			n += 1
		}
		if err := fscanner.Err(); err != nil {
			return err
		}
		return nil
	})
}
