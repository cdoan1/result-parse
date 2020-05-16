package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/kyokomi/emoji"
)

// type Suite struct {
// 	XMLName    xml.Name    `xml:"suite"`
// 	TestSuites []TestSuite `xml:"suites"`
// }

type TestSuite struct {
	// XMLName  xml.Name `xml:"testsuite"`
	Name      string     `xml:"name,attr"`
	Tests     string     `xml:"tests,attr"`
	Failures  string     `xml:"failures,attr"`
	Errors    string     `xml:"errors,attr"`
	Time      string     `xml:"time,attr"`
	TestCases []TestCase `xml:"testcase"`
}

type TestCase struct {
	XMLName   xml.Name  `xml:"testcase"`
	Name      string    `xml:"name,attr"`
	ClassName string    `xml:"classname,attr"`
	Time      string    `xml:"time,attr"`
	Skipped   []Skipped `xml:"skipped"`
}

type Skipped struct {
	XMLName xml.Name `xml:"skipped"`
}

func ocClusterInfo() {
	// c := []string{"oc", "cluster-info"}
	// s := c[:]

	cmd := exec.Command("oc", "cluster-info")
	// cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	emoji.Printf(":beer: %s\n", out.String())
}

func ocNodesWide() {
	cmd := exec.Command("oc", "get", "nodes", "-o", "wide")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	emoji.Printf(":beer: %s\n", out.String())

}

func ocGetMCHVersion() {
	// oc get mch multiclusterhub -o jsonpath='{.status.currentVersion}'
	cmd := exec.Command("oc", "get", "mch", "multiclusterhub", "-o", "jsonpath='{.status.currentVersion}'")
	// cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("\u2714 multiclusterhub version: %s\n", out.String())
	emoji.Printf(":beer: multiclusterhub version: %s\n", out.String())
}

func ocGetMCHSHA() string {
	cmd := "oc get pods -o yaml | grep image: | grep multiclusterhub"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd)
	}
	return string(out)
}

func main() {
	fmt.Println("Render ginkgo results...")

	// Open our xmlFile
	var resultFile = "results.xml"

	xmlFile, err := os.Open(resultFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully Opened %s\n", resultFile)
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var suite TestSuite

	xml.Unmarshal(byteValue, &suite)

	// display results
	for i := 0; i < len(suite.TestCases); i++ {
		// small check
		// fmt.Println("\u2714 " + suite.TestCases[i].Name)
		if suite.TestCases[i].Skipped == nil {
			fmt.Println("\u2705 " + suite.TestCases[i].Name)
		} else {
			fmt.Println("\u2796 " + suite.TestCases[i].Name)
		}
		// fmt.Println("Time: " + suite.TestCases[i].Time)
	}

	var tests int
	tests, _ = strconv.Atoi(suite.Tests)
	skipped := len(suite.TestCases) - tests

	emoji.Printf(":beer: Summary: Tests:%s/%d Failures:%s Errors:%s Time:%s \n", suite.Tests, skipped, suite.Failures, suite.Errors, suite.Time)

	// ocClusterInfo()
	ocNodesWide()
	ocGetMCHVersion()
	// emoji.Printf(":beer: %s", ocGetMCHSHA())

}

// ✅
// check mark
// Unicode: U+2705, UTF-8: E2 9C 85
