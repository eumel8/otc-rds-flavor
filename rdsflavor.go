package main

import (
	"flag"
	"fmt"
	"strings"

	"net/http"
	"os"

	"github.com/gophercloud/utils/client"
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/flavors"
	"k8s.io/klog/v2"
)

var (
	osExit = os.Exit
	ds     *string
	dt     *string
	ha     *bool
	rr     *bool
	az1    *string
	az2    *string
	v      *bool
)

const (
	AppVersion = "0.0.1"
)

func init() {
	az1 = flag.String("az1", "eu-de-01", "AZ instance 1 (e.g. eu-de-01)")
	az2 = flag.String("az2", "eu-de-02", "AZ instance 1 (e.g. eu-de-02)")
	ds = flag.String("ds", "mysql", "Datastore (e.g. mysql)")
	dt = flag.String("dt", "8.0", "Datatype (e.g. 8.0)")
	ha = flag.Bool("ha", false, "HA flavor")
	rr = flag.Bool("rr", false, "RR flavor")
	v = flag.Bool("v", false, "version of the program")
}

func rdsLookup(client *golangsdk.ServiceClient, az1 string, az2 string, ds string, dt string, ha bool, rr bool) error {

	listOpts := flavors.ListOpts{
		VersionName: dt,
	}
	allFlavorPages, err := flavors.List(client, listOpts, ds).AllPages()
	if err != nil {
		klog.Exitf("unable to list flavor: %v", err)
		return err
	}

	rdsFlavors, err := flavors.ExtractDbFlavors(allFlavorPages)
	if err != nil {
		klog.Exitf("unable to extract flavor: %v", err)
		return err
	}

	fmt.Printf("CPU\tRAM\tSpecCode\n")
	for _, rds := range rdsFlavors {
		for n, az := range rds.AzStatus {
			if n == az1 && az == "normal" {
				for l, az := range rds.AzStatus {
					if l == az2 && az == "normal" {
						if ha && strings.HasSuffix(rds.SpecCode, ".ha") {
							fmt.Printf("%s\t%d\t%s\n", rds.VCPUs, rds.RAM, rds.SpecCode)
						}
						if rr && strings.HasSuffix(rds.SpecCode, ".rr") {
							fmt.Printf("%s\t%d\t%s\n", rds.VCPUs, rds.RAM, rds.SpecCode)
						}
						if !ha && !rr && !strings.HasSuffix(rds.SpecCode, ".rr") && !strings.HasSuffix(rds.SpecCode, ".ha") {
							fmt.Printf("%s\t%d\t%s\n", rds.VCPUs, rds.RAM, rds.SpecCode)
						}
					}

				}

			}

		}
	}
	return nil
}

func getProvider() *golangsdk.ProviderClient {
	if os.Getenv("OS_AUTH_URL") == "" {
		os.Setenv("OS_AUTH_URL", "https://iam.eu-de.otc.t-systems.com:443/v3")
	}

	if os.Getenv("OS_IDENTITY_API_VERSION") == "" {
		os.Setenv("OS_IDENTITY_API_VERSION", "3")
	}

	if os.Getenv("OS_REGION_NAME") == "" {
		os.Setenv("OS_REGION_NAME", "eu-de")
	}

	if os.Getenv("OS_PROJECT_NAME") == "" {
		os.Setenv("OS_PROJECT_NAME", "eu-de")
	}

	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		klog.Exitf("error getting auth from env: %v", err)
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		klog.Exitf("unable to initialize openstack client: %v", err)
	}

	if os.Getenv("OS_DEBUG") != "" {
		provider.HTTPClient = http.Client{
			Transport: &client.RoundTripper{
				Rt:     &http.Transport{},
				Logger: &client.DefaultLogger{},
			},
		}
	}
	return provider
}

func getFlags(cliFullArg string) {
	flag.Parse()
	fmt.Println("lookup", *ds, *dt, *az1, *az2)
	Lookup()
}

func Lookup() {
	provider := getProvider()

	rds, err := openstack.NewRDSV3(provider, golangsdk.EndpointOpts{})
	if err != nil {
		klog.Exitf("unable to initialize rds client: %v", err)
		return
	}

	err = rdsLookup(rds, *az1, *az2, *ds, *dt, *ha, *rr)
	if err != nil {
		klog.Exitf("rds flavor lookup failed: %v", err)
		return
	}
}

func main() {
	getFlags("default")
}
