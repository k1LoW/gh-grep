/*
Copyright © 2021 Ken'ichiro Oyama <k1lowxb@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"regexp"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/fatih/color"
	"github.com/johejo/ghfs"
	"github.com/k1LoW/gh-grep/gh"
	"github.com/k1LoW/gh-grep/version"
	"github.com/spf13/cobra"
)

var (
	owner   string
	repos   []string
	include string
	exclude string
)

var rootCmd = &cobra.Command{
	Use:          "gh-grep [PATTERN]",
	Short:        "Print lines matching a pattern in repositories using GitHub API",
	Long:         `Print lines matching a pattern in repositories using GitHub API`,
	Args:         cobra.ExactArgs(1),
	Version:      version.Version,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		pattern, err := regexp.Compile(args[0])
		if err != nil {
			return err
		}
		g, err := gh.New()
		if err != nil {
			return err
		}
		if len(repos) == 0 {
			repos, err = g.Repositories(ctx, owner)
			if err != nil {
				return err
			}
		}

		delimter := color.New(color.FgCyan).Sprint(":")

		for _, repo := range repos {
			log.Printf("In %s/%s\n", owner, repo)
			fsys := ghfs.NewWithGitHubClient(g.Client(), owner, repo)
			if err := doublestar.GlobWalk(fsys, include, func(path string, d fs.DirEntry) error {
				if d.IsDir() {
					return nil
				}
				if exclude != "" {
					match, err := doublestar.PathMatch(exclude, path)
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
					matches := pattern.FindAllStringIndex(line, -1)
					if len(matches) > 0 {
						// TODO: color
						fmt.Printf("%s/%s%s%s%s%s\n", owner, repo, delimter, path, delimter, line)
					}
				}
				if err := fscanner.Err(); err != nil {
					return err
				}
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	},
}

func Execute() {
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)

	log.SetOutput(io.Discard)
	if env := os.Getenv("DEBUG"); env != "" {
		log.SetOutput(os.Stderr)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&owner, "owner", "", "", "owner")
	rootCmd.MarkFlagRequired("owner")
	rootCmd.Flags().StringSliceVarP(&repos, "repo", "", []string{}, "repo")
	rootCmd.Flags().StringVarP(&include, "include", "", "**/*", "search only files that match pattern")
	rootCmd.Flags().StringVarP(&exclude, "exclude", "", "", "skip files and directories matching pattern")
}
