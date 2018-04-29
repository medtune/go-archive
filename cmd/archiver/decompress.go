package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/a-hilaly/dec"

	"github.com/google/subcommands"
)

type decompressCmd struct {
	kind        string
	destination string
}

func (*decompressCmd) Name() string     { return "decompress" }
func (*decompressCmd) Synopsis() string { return "decompress a file" }
func (*decompressCmd) Usage() string {
	return `decompress [-t archive type] [-d destination] somefile:
	 decompress somefile to wanted kind (zip by default).
  `
}

func (cc *decompressCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cc.kind, "t", "zip", "archive type")
	f.StringVar(&cc.destination, "d", "", "destination")
}

func (cc *decompressCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	worker, err := archiver.Pick(cc.kind)
	if err != nil {
		log.Fatalf("Compression type unsupported %s", cc.kind)
	}

	var dir string
	if cc.destination != "" {
		dir = cc.destination
	} else {
		dir, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}

	args := f.Args()
	for _, arg := range args {
		fmt.Println("decompressing ... ", arg)
		//strings.TrimSuffix(arg, filepath.Ext(arg))
		worker.Decompress(arg, dir)
	}

	return subcommands.ExitSuccess
}
