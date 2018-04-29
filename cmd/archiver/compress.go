package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/a-hilaly/dec"

	"github.com/google/subcommands"
)

type compressCmd struct {
	kind string
}

func (*compressCmd) Name() string     { return "compress" }
func (*compressCmd) Synopsis() string { return "compress a file" }
func (*compressCmd) Usage() string {
	return `compress [-d destination] [-k archive type] somefile:
	 compress somefile to wanted kind (zip by default).
  `
}

func (cc *compressCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cc.kind, "t", "zip", "archive type")
}

func (cc *compressCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	worker, err := archiver.Pick(cc.kind)
	if err != nil {
		fmt.Printf("Compression type unsupported %s", cc.kind)
	}
	args := f.Args()
	for _, arg := range args {
		fmt.Println("compressing ... ", arg)

		worker.Compress(arg, strings.Join([]string{arg, ".", worker.Meta().Ext}, ""))
	}
	return subcommands.ExitSuccess
}
