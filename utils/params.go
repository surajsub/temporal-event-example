package utils

import (
	"fmt"
	"github.com/surajsub/temporal-event-example/models"
	"net"
)

var (
	vpcMessage        = "vpc Name must be provided"
	vpcIdMessage      = "vpc ID must be provided"
	cidrMessage       = "valid CIDR Must be Provided"
	subnetMessage     = "subnet name must be provided"
	subnetIdMessage   = "valid subnet id must be provided"
	subnetSizeMessage = "valid subnet size must be provided"
	igwMessage        = "igw name must be provided"
	igwIdMessage      = "valid IGW Id must be provided"
	natMessage        = "valid nat name must be provided"
	natIdMessage      = "valid Nat Id must be provided"
	sgIDMessage       = "valid VPC ID must be provided"
	instanceMessage   = "valid EC2 instance name must be provided"
)

// ValidateVPCParams validates the parameters for creating a VPC.
// It checks that the Name and CIDR fields are properly set and valid.
// It returns a map of parameters if validation is successful, or an error if validation fails.

func ValidateVPCParams(params models.VPCParams) (map[string]string, error) {
	if params.Name == "" {
		return nil, fmt.Errorf(vpcMessage)
	}
	if params.CIDR == "" {
		return nil, fmt.Errorf(cidrMessage)
	}
	if !isValidCIDR(params.CIDR) {
		return nil, fmt.Errorf("invalid CIDR block: %s", params.CIDR)
	}

	// Convert SubnetParams to map[string]*string
	paramsMap := map[string]string{
		"name": params.Name,
		"cidr": params.CIDR,
	}

	return paramsMap, nil
}

// ValidateSubnetParams validates the parameters for creating a Subnet.
// It checks that the SubnetName, SubnetSize and the VPCID are properly set and valid.
// It returns a map of parameters if validation is successful, or an error if validation fails.
func ValidateSubnetParams(params models.SubnetParams) (map[string]string, error) {
	if params.VPCID == "" {
		return nil, fmt.Errorf(vpcIdMessage)
	}
	if params.SubnetName == "" {
		return nil, fmt.Errorf(subnetMessage)
	}
	if params.SubnetSize == "" {
		return nil, fmt.Errorf(subnetSizeMessage)
	}

	// Convert SubnetParams to map[string]*string
	paramsMap := map[string]string{
		"vpcID":      params.VPCID,
		"subnetSize": params.SubnetSize,
		"subnetName": params.SubnetName,
	}
	return paramsMap, nil
}

// ValidateIGWParams validates IGW parameters.
func ValidateIGWParams(params models.IGWParams) (map[string]string, error) {
	if params.VPCID == "" {
		return nil, fmt.Errorf(vpcIdMessage)
	}
	if params.IGWName == "" {
		return nil, fmt.Errorf(igwMessage)
	}

	// Convert SubnetParams to map[string]*string
	paramsMap := map[string]string{
		"vpcID":   params.VPCID,
		"igwName": params.IGWName,
	}
	return paramsMap, nil
}

// Validate NAT

func ValidateNATParams(natparams models.NatParams) (map[string]string, error) {
	if natparams.NatName == "" {
		return nil, fmt.Errorf(natMessage)
	}
	if natparams.SubnetId == "" {
		return nil, fmt.Errorf("SubNetID is required")
	}

	// Convert SubnetParams to map[string]*string
	paramsMap := map[string]string{
		"natName":  natparams.NatName,
		"subnetID": natparams.SubnetId,
	}
	return paramsMap, nil
}

// Validate Security Group
func ValidateSGParams(sgparams models.SGParams) (map[string]string, error) {
	if sgparams.VPCID == "" {
		return nil, fmt.Errorf(vpcIdMessage)
	}
	// Convert SGparams to map[string]*string
	paramsMap := map[string]string{
		"vpcID": sgparams.VPCID,
	}
	return paramsMap, nil
}

// Validate RT

func ValidateRTParams(rtparams models.RTParams) (map[string]string, error) {
	if rtparams.IGWID == "" {
		return nil, fmt.Errorf(igwIdMessage)
	}
	if rtparams.VPCID == "" {
		return nil, fmt.Errorf(vpcIdMessage)
	}
	if rtparams.NatID == "" {
		return nil, fmt.Errorf(natIdMessage)
	}
	if rtparams.SubnetID == "" {
		return nil, fmt.Errorf(subnetIdMessage)
	}

	paramsMap := map[string]string{
		"vpcID":    rtparams.VPCID,
		"subnetID": rtparams.SubnetID,
		"natID":    rtparams.NatID,
		"igwID":    rtparams.IGWID,
	}
	return paramsMap, nil
}

// ValidateVPCParams validates the parameters for a VPC.

// ValidateEC2Params - TODO

// isValidCIDR checks if a given string is a valid CIDR block.
func isValidCIDR(cidr string) bool {
	_, _, err := net.ParseCIDR(cidr)
	return err == nil
}
