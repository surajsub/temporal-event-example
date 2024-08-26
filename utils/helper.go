package utils

import (
	"flag"
	"fmt"
	"github.com/surajsub/temporal-event-example/models"
	"log"
	"os"
)

// Activity definitions

// Helper functions to extract IDs from Terraform output

// GetProvisioner returns the appropriate provisioner, engine name, and directory path based on the input.
func GetProvisioner(iac string) (Provisioner, string, string) {
	provisioners := map[string]struct {
		Provisioner Provisioner
		Directory   string
		EngineName  string
	}{
		"terraform": {
			Provisioner: &TerraformProvisioner{},
			Directory:   BASETFDIRECTORY,
			EngineName:  "terraform",
		},
		"opentofu": {
			Provisioner: &TofuProvisioner{},
			Directory:   BASEOTDIRECTORY,
			EngineName:  "opentofu",
		},
	}

	config, ok := provisioners[iac]
	if !ok {
		log.Fatalf("Unknown IAC type: %s", iac)
	}

	log.Printf("Using %s provisioner", config.EngineName)
	return config.Provisioner, config.EngineName, config.Directory
}

func ValidateInputParams(resourceType string, vpcParams models.VPCParams, subnetParams models.SubnetParams, igwParams models.IGWParams, natparams models.NatParams, rtparams models.RTParams, sgparams models.SGParams) (map[string]string, error) {
	switch resourceType {
	case "vpc":
		return ValidateVPCParams(vpcParams)
	case "subnet":
		return ValidateSubnetParams(subnetParams)
	case "igw":
		return ValidateIGWParams(igwParams)
	case "nat":
		return ValidateNATParams(natparams)
	case "sg":
		return ValidateSGParams(sgparams)
	case "rt":
		return ValidateRTParams(rtparams)
	default:
		return nil, fmt.Errorf("invalid resource type: %s", resourceType)
	}

	return nil, nil
}

/*
This function reads the input provides a response for the NATS message

*/

func ParseInputFlags() (map[string]string, string, error) {
	// Define flags
	resourceType := flag.String("resource", "vpc", "Type of AWS resource to provision (e.g., vpc, subnet)")
	name := flag.String("name", "", "Name of the resource")
	cidr := flag.String("cidr", "", "CIDR block (for VPC or Subnet)")
	engine := flag.String("engine", "", "Engine used to provision the resource")
	vpcID := flag.String("vpcID", "", "VPC ID (for Subnet)")
	subnetName := flag.String("subnetName", "", "Subnet Name")
	subnetSize := flag.String("subnetSize", "", "Subnet size")
	igwName := flag.String("igwName", "", "IGW Name")
	subnetID := flag.String("subnetID", "", "Subnet id")
	natID := flag.String("natID", "", "Nat id")
	natName := flag.String("natName", "", "Nat Name")
	igwID := flag.String("igwID", "", "IGW ID")

	if *engine != "terraform" && *engine != "opentofu" && *engine != "" {
		log.Fatalf("Invalid engine provided")
		os.Exit(1)
	}

	// Parse flags
	flag.Parse()

	// Initialize parameters
	vpcParams := models.VPCParams{
		Name: *name,
		CIDR: *cidr,
	}

	subnetParams := models.SubnetParams{
		VPCID:      *vpcID,
		SubnetName: *subnetName,
		SubnetSize: *subnetSize,
	}

	igwParams := models.IGWParams{
		VPCID:   *vpcID,
		IGWName: *igwName,
	}

	natParams := models.NatParams{
		SubnetId: *subnetID,
		NatName:  *natName,
	}

	rtParams := models.RTParams{
		VPCID:    *vpcID,
		NatID:    *natID,
		IGWID:    *igwID,
		SubnetID: *subnetID,
	}

	sgParams := models.SGParams{
		VPCID: *vpcID,
	}

	//return *engine, *resourceType, vpcParams, subnetParams, igwParams

	params, err := ValidateInputParams(*resourceType, vpcParams, subnetParams, igwParams, natParams, rtParams, sgParams)
	if err != nil {
		return nil, "", err
	}
	log.Println("Received response from the validated Input function")
	params["engine"] = *engine
	return params, *resourceType, err
}
