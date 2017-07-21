package main

import (
	goflag "flag"
	"fmt"
	"log"
	"os"

	loads "github.com/go-openapi/loads"
	"github.com/spf13/pflag"

	"github.com/sapcc/kubernikus/pkg/api/rest"
	"github.com/sapcc/kubernikus/pkg/api/rest/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Analyzed(rest.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	var server *rest.Server // make sure init is called

	pflag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage:\n")
		fmt.Fprint(os.Stderr, "  kubernikus-apiserver [OPTIONS]\n\n")

		title := "Kubernikus"
		fmt.Fprint(os.Stderr, title+"\n\n")
		desc := swaggerSpec.Spec().Info.Description
		if desc != "" {
			fmt.Fprintf(os.Stderr, desc+"\n\n")
		}
		fmt.Fprintln(os.Stderr, pflag.CommandLine.FlagUsages())
	}
	// parse the CLI flags
	if f := goflag.Lookup("logtostderr"); f != nil {
		f.Value.Set("true") // log to stderr by default
	}
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine) //slurp in glog flags
	pflag.Parse()
	//goflag.CommandLine.Parse([]string{}) //https://github.com/kubernetes/kubernetes/issues/17162

	api := operations.NewKubernikusAPI(swaggerSpec)
	// get server with flag values filled out
	server = rest.NewServer(api)

	defer server.Shutdown()

	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
