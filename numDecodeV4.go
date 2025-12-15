// numDecodeV4
// program that decodes a numbering.xml file from a docx file
// create functions

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
//	"encoding/xml"

	numLib "goDemo/goDocx/numbering/numLib"
    "github.com/gomutex/godocx"
    util "github.com/prr123/utility/utilLib"
)


func main() {

    numarg := len(os.Args)
    flags:=[]string{"dbg", "in"}

    useStr := "/in=<infile> [/dbg]"
    helpStr := "doxc parsing program"

    if numarg > len(flags) +1 {
        fmt.Println("too many arguments in cl!")
        fmt.Println("usage: %s %s\n", os.Args[0], useStr)
        os.Exit(-1)
    }

    if numarg == 1 || (numarg > 1 && os.Args[1] == "help") {
        fmt.Printf("help: %s\n", helpStr)
        fmt.Printf("usage is: %s %s\n", os.Args[0], useStr)
        os.Exit(1)
    }

    flagMap, err := util.ParseFlags(os.Args, flags)
    if err != nil {log.Fatalf("util.ParseFlags: %v\n", err)}

    dbg:= false
    _, ok := flagMap["dbg"]
    if ok {dbg = true}

    inFil := ""
    inval, ok := flagMap["in"]
    if !ok {
        log.Fatalf("error -- no in flag provided!\n")
    } else {
        if inval.(string) == "none" {log.Fatalf("error -- no input file name provided!\n")}
        inFil = inval.(string)
        idx := strings.IndexByte(inFil, '.')
        if idx > -1 {log.Fatalf("error -- infile <%s> has an extension!\n", inFil)}
    }

    inpFilnam := "/home/peter/go/src/goDemo/goDocx/docx/" + inFil + ".docx"
//    inpFilnam := "/home/peter/go/src/goDemo/goDocx/docx/" + inFil + "/word/numbering.xml"
    outFilnam := "out/" + inFil + ".numTxt"
//    jsFilnam := "out/" + inFil + ".js"

    if dbg {
        fmt.Printf("input:  %s\n", inpFilnam)
        fmt.Printf("output: %s\n", outFilnam)
    }

    //read docx file
    rdoc, err := godocx.OpenDocument(inpFilnam)
    if err !=nil {log.Fatalf("error -- opening doc: %v\n", err)}

	log.Printf("info -- success in retrieving numbering.xml!")

	numObj, err := numLib.GetNumObj(rdoc)
	if err != nil {log.Fatalf("error -- getNumObj: %v\n", err)}

	numObj.PrintNumObj()

	ML, err := numObj.CreNList()
	if err != nil { log.Fatalf("error -- CreNList: %v\n", err)}

	ML.PrintDocxList()
}

