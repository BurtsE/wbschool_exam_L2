package main

import (
	"bufio"
	"os"
	"testing"
)

func TestCut(t *testing.T) {
	var s bool
	var f = []int{1, 4}
	var d = "."
	inputFile, _ := os.Create("/tmp/inputFile")
	inputFile.Write([]byte("sdfsdf.afdafs.afsaffs.afsf\n"))
	inputFile.Write([]byte("zxcv.qwer.afsavcxzffs.hbng\n"))
	inputFile.Write([]byte("zc.cxz.afsavcxzvczxffs.hbvfdng\n"))

	outputFile, _ := os.Create("/tmp/outputFile")
	inputFile.Close()

	inputFile, _ = os.Open("/tmp/inputFile")

	r := bufio.NewScanner(inputFile)
	w := bufio.NewWriter(outputFile)

	Cut(s, f, d, r, w)
	w.Flush()
	outputFile.Close()
	outputFile, _ = os.Open("/tmp/outputFile")

	r = bufio.NewScanner(outputFile)
	outputFile, _ = os.Open("/tmp/inputFile")

	testOutput := []string{"sdfsdf:afsf", "zxcv:hbng", "zc:hbvfdng"}
	var q int
	for r.Scan() {
		if q >= len(testOutput) || r.Text() != testOutput[q] {
			t.Errorf("Expected:%s. Got: %s", testOutput[q], r.Text())
		}
		q++
	}
}
