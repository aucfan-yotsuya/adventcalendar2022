package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/winlinvip/go-fdkaac/fdkaac"
)

func main() {}
func Run() error {
	wfilename := flag.String("w", "", ".wav filename")
	afilename := flag.String("a", "", ".aac filename")
	flag.Parse()
	if *wfilename == "" || *afilename == "" {
		flag.PrintDefaults()
		os.Exit(-1)
	}
	wfp, err := os.Open(*wfilename)
	if err != nil {
		return err
	}
	afp, err := os.OpenFile(*afilename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() { afp.Close(); wfp.Close() }()
	wbyte, err := ioutil.ReadAll(wfp)
	if err != nil {
		return err
	}
	e := fdkaac.NewAacEncoder()
	abyte, err := e.Encode(wbyte)
	if err != nil {
		return err
	}
	_, err = afp.Write(abyte)
	if err != nil {
		return err
	}
	return nil
}
