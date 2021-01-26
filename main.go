package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/bevisy/go2uml/codeanalysis"
	"github.com/jessevdk/go-flags"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	log.SetLevel(log.InfoLevel)

	var opts struct {
		CodeDir    string   `long:"codedir" description:"要扫描的代码目录" required:"true"`
		GopathDir  string   `long:"gopath" description:"GOPATH directory" required:"true"`
		OutputFile string   `long:"out" description:"解析结果保存到该文件中" required:"true"`
		IgnoreDirs []string `long:"ignoredir" description:"需要排除的目录,不需要扫描和解析"`
	}

	if len(os.Args) == 1 {
		fmt.Println("example:\n" +
			os.Args[0] + " --gopath $GOPATH --codedir $GOPATH/src/sampleCode --outputfile sampleCode.puml")
		os.Exit(1)
	}

	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		os.Exit(1)
	}

	if opts.CodeDir == "" {
		panic("代码目录不能为空")
		os.Exit(1)
	}

	if opts.GopathDir == "" {
		panic("GOPATH目录不能为空")
		os.Exit(1)
	}

	if !strings.HasPrefix(opts.CodeDir, opts.GopathDir) {
		panic(fmt.Sprintf("代码目录%s,必须是GOPATH目录%s的子目录", opts.CodeDir, opts.GopathDir))
		os.Exit(1)
	}

	for _, dir := range opts.IgnoreDirs {
		if !strings.HasPrefix(dir, opts.CodeDir) {
			panic(fmt.Sprintf("需要排除的目录%s,必须是代码目录%s的子目录", dir, opts.CodeDir))
			os.Exit(1)
		}
	}

	config := codeanalysis.Config{
		CodeDir:    opts.CodeDir,
		GopathDir:  opts.GopathDir,
		VendorDir:  filepath.Join(opts.CodeDir, "vendor"),
		IgnoreDirs: opts.IgnoreDirs,
	}

	result := codeanalysis.AnalysisCode(config)

	result.OutputToFile(opts.OutputFile)

}
