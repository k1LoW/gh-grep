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

type Args struct {
	Pattern *regexp.Regexp
	Owner   string
	Repo    string
	Include string
	Exclude string
}

func Scan(ctx context.Context, fsys fs.FS, args *Args) error {
	return doublestar.GlobWalk(fsys, args.Include, func(path string, d fs.DirEntry) error {
		if d.IsDir() {
			return nil
		}
		if args.Exclude != "" {
			match, err := doublestar.PathMatch(args.Exclude, path)
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
		for fscanner.Scan() {
			line := fscanner.Text()
			matches := args.Pattern.FindAllStringIndex(line, -1)
			if len(matches) > 0 {
				fmt.Printf("%s/%s%s%s%s%s\n", args.Owner, args.Repo, delimiter, path, delimiter, internal.PrintLine(line, matches, matchc))
			}
		}
		if err := fscanner.Err(); err != nil {
			return err
		}
		return nil
	})
}
