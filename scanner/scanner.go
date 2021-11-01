package scanner

import (
	"bufio"
	"context"
	"fmt"
	"io/fs"
	"log"
	"regexp"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/fatih/color"
	"github.com/k1LoW/gh-grep/internal"
)

var (
	matchc    = color.New(color.FgRed, color.Bold)
	delimiter = color.New(color.FgCyan).Sprint(":")
)

type Opts struct {
	Pattern    *regexp.Regexp
	Owner      string
	Repo       string
	Include    string
	Exclude    string
	LineNumber bool
}

func Scan(ctx context.Context, fsys fs.FS, opts *Opts) error {
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
			matches := opts.Pattern.FindAllStringIndex(line, -1)
			if len(matches) > 0 {
				if opts.LineNumber {
					fmt.Printf("%s/%s%s%s%s%d%s%s\n", opts.Owner, opts.Repo, delimiter, path, delimiter, n, delimiter, internal.PrintLine(line, matches, matchc))
				} else {
					fmt.Printf("%s/%s%s%s%s%s\n", opts.Owner, opts.Repo, delimiter, path, delimiter, internal.PrintLine(line, matches, matchc))
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
