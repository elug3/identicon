package main

import (
	"crypto/rand"
	"encoding/base32"
	"flag"
	"fmt"
	"os"

	"github.com/elug3/identicon"
)

var usageStr = `identicon [options] <seed>

Generate an identicon image based on the input string.
Options:
  -s <size>	Size of the image in pixels (default 256)
  -i <seed>	Input string to generate the identicon
  -o <output>	Output file name
  -h		Show this help message
`

func printUsage() {
	println(usageStr)
	os.Exit(0)
}

func parseArgs(fs *flag.FlagSet, args []string) (*Options, error) {
	var opts Options

	fs.IntVar(&opts.Size, "s", 255, "Size of the image in pixels")
	fs.BoolVar(&opts.Help, "h", false, "Show help message")
	fs.Usage = printUsage

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if opts.Help {
		printUsage()
		return nil, nil
	}

	args = fs.Args()
	if len(args) > 0 {
		if args[0] == "" {
			return nil, fmt.Errorf("seed cannot be empty")
		}
		opts.Seed = args[0] // First positional argument as seed
	} else {
		opts.Seed = randomSeed()
	}

	return &opts, nil
}

type Options struct {
	Size int
	Seed string
	Help bool
}

func randomSeed() string {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(b)
}

func main() {

	fs := flag.NewFlagSet("identicon", flag.ExitOnError)
	fs.Usage = printUsage

	opts, err := parseArgs(fs, os.Args[1:])
	if err != nil {
		println("Error parsing arguments:", err.Error())
		os.Exit(1)
	}

	img := identicon.New([]byte(opts.Seed), opts.Size)

	filename := opts.Seed + ".png"
	if err = identicon.SavePNG(img, filename); err != nil {
		println("Error saving image:", err.Error())
		os.Exit(1)
	}
	println("Identicon saved to " + filename)
}
