package workflows

import (
	"fmt"
	"github.com/surajsub/temporal-event-example/activities"
	"github.com/surajsub/temporal-event-example/utils"
	"go.temporal.io/sdk/workflow"
	"log"
	"time"
)

func ProvisionAWSResources(ctx workflow.Context, resourceType string, params map[string]string) (map[string]string, error) {
	// log.Printf("Printing the data for the resources type [ %s ] and the params [%v\n] ", resourceType, params)

	results := make(map[string]interface{})
	var proacs activities.AWSImpl
	templog := workflow.GetLogger(ctx)
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * 10 * time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	switch resourceType {

	case "vpc":
		var vpcOutput map[string]string

		err := workflow.ExecuteActivity(ctx, proacs.VPCInitActivity, params["engine"]).Get(ctx, nil)
		if err != nil {
			templog.Info(utils.VpcInit, "failed")
		}

		err = workflow.ExecuteActivity(ctx, proacs.VPCApplyActivity, params["engine"], params["name"], params["cidr"]).Get(ctx, nil)
		if err != nil {
			templog.Info(utils.VpcApply, "failed")
		}
		err = workflow.ExecuteActivity(ctx, proacs.VPCOutputActivity, params["engine"]).Get(ctx, &vpcOutput)
		if err != nil {
			templog.Info(utils.VpcOutput, "failed")
		}

		log.Printf("the vpc output from the activity  is %v\n", vpcOutput)
		// Marshal the VPC details into a JSON string
		vpcid, err := extractVPCID(vpcOutput)
		if err != nil {
			return nil, err
		}
		log.Printf("the vpcid from the extracvpcid is %s\n", vpcid)
		results["vpcResult"] = vpcOutput
		return vpcOutput, nil

		// start of the Subnet
	case "subnet":
		var subnetOutput map[string]string
		err := workflow.ExecuteActivity(ctx, proacs.SubnetInitActivity, params["engine"], params["vpcID"]).Get(ctx, nil)
		if err != nil {
			templog.Info(utils.SubnetInit, "failed")
		}
		err = workflow.ExecuteActivity(ctx, proacs.SubnetApplyActivity, params["engine"], params["vpcID"], params["subnetName"], params["subnetSize"]).Get(ctx, nil)
		if err != nil {
			templog.Info(utils.SubnetApply, "failed")
		}
		err = workflow.ExecuteActivity(ctx, proacs.SubnetOutputActivity, params["engine"]).Get(ctx, &subnetOutput)
		if err != nil {
			templog.Info(utils.SubnetOutput, "failed")
		}

		if err != nil {
			return nil, err
		}
		log.Printf("the Subnet output from the activity  is %v\n", subnetOutput)
		results["subnetResult"] = subnetOutput
		return subnetOutput, nil

	case "nat":
		var natOutput map[string]string
		err := workflow.ExecuteActivity(ctx, proacs.NATInitActivity, params["engine"]).Get(ctx, nil)
		if err != nil {
			templog.Info(utils.NatInit, "failed")
		}

		err = workflow.ExecuteActivity(ctx, proacs.NATApplyActivity, params["engine"], params["subnetID"], params["natName"]).Get(ctx, nil)
		if err != nil {
			templog.Info(utils.NatApply, "failed")
		}
		err = workflow.ExecuteActivity(ctx, proacs.NATOutputActivity, params["engine"]).Get(ctx, &natOutput)
		if err != nil {
			templog.Info(utils.NatOutput, "failed")
		}

		log.Printf("the NAT output from the activity  is %v\n", natOutput)
		if err != nil {
			return nil, err
		}

		results["natOuput"] = natOutput
		return natOutput, nil

	case "igw":
		var igwOutput map[string]string
		err := workflow.ExecuteActivity(ctx, proacs.IGWInitActivity, params["engine"]).Get(ctx, nil)
		if err != nil {
			templog.Info(utils.IgwInit, "failed")
		}

		err = workflow.ExecuteActivity(ctx, proacs.IGWApplyActivity, params["engine"], params["vpcID"], params["igwName"]).Get(ctx, nil)
		if err != nil {
			templog.Info(utils.IgwApply, "failed")
		}
		err = workflow.ExecuteActivity(ctx, proacs.IGWOutputActivity, params["engine"]).Get(ctx, &igwOutput)
		if err != nil {
			templog.Info(utils.IgwOutput, "failed")
		}

		log.Printf("the IGW output from the activity  is %v\n", igwOutput)
		if err != nil {
			return nil, err
		}

		results["igwResult"] = igwOutput
		return igwOutput, nil

	case "sg":
		var sgOutput map[string]string
		err := workflow.ExecuteActivity(ctx, proacs.SGInitActivity, params["engine"]).Get(ctx, nil)
		if err != nil {
			templog.Info(utils.SgInit, "failed")
		}

		err = workflow.ExecuteActivity(ctx, proacs.SGApplyActivity, params["engine"], params["vpcID"]).Get(ctx, nil)
		if err != nil {
			templog.Info(utils.SgApply, "failed")
		}
		err = workflow.ExecuteActivity(ctx, proacs.SGOutputActivity, params["engine"]).Get(ctx, &sgOutput)
		if err != nil {
			templog.Info(utils.SgOutput, "failed")
		}

		log.Printf("the SG output from the activity  is %v\n", sgOutput)
		if err != nil {
			return nil, err
		}
		results["sgResult"] = sgOutput
		return sgOutput, nil

	case "rt":
		var rtOutput map[string]string
		err := workflow.ExecuteActivity(ctx, proacs.RTInitActivity, params["engine"]).Get(ctx, nil)
		if err != nil {
			templog.Info(utils.RtInit, "failed")
		}

		err = workflow.ExecuteActivity(ctx, proacs.RTApplyActivity, params["engine"], params["vpcID"], params["subnetID"], params["igwID"], params["natID"]).Get(ctx, nil)
		if err != nil {
			templog.Info(utils.RtApply, "failed")
		}
		err = workflow.ExecuteActivity(ctx, proacs.RTOutputActivity, params["engine"]).Get(ctx, &rtOutput)
		if err != nil {
			templog.Info(utils.RtOutput, "failed")
		}

		log.Printf("the RT output from the activity  is %v\n", rtOutput)
		if err != nil {
			return nil, err
		}
		results["rtResult"] = rtOutput
		log.Printf("The results are %v\n", results)
		return rtOutput, nil
	case "ec2":
		var ec2Output map[string]string
		err := workflow.ExecuteActivity(ctx, proacs.EC2InitActivity, params["engine"]).Get(ctx, nil)
		if err != nil {
			templog.Info(utils.Ec2Init, "failed")
		}

		err = workflow.ExecuteActivity(ctx, proacs.EC2ApplyActivity, params["engine"], params["vpcID"], params["subnetID"], params["igwID"], params["natID"]).Get(ctx, nil)
		if err != nil {
			templog.Info(utils.Ec2Apply, "failed")
		}
		err = workflow.ExecuteActivity(ctx, proacs.EC2OutputActivity, params["engine"]).Get(ctx, &ec2Output)
		if err != nil {
			templog.Info(utils.Ec2Output, "failed")
		}

		log.Printf("the EC2 output from the activity  is %v\n", ec2Output)
		if err != nil {
			return nil, err
		}
		results["ec2Result"] = ec2Output
		log.Printf("The results are %v\n", results)
		return ec2Output, nil
	default:
		return nil, fmt.Errorf("unknown resource type: %s", resourceType)
	}
}

func extractVPCID(output map[string]string) (string, error) {
	vpcID, ok := output["vpc_id"]
	if !ok {
		return "", fmt.Errorf("vpc_id not found in output")
	}
	return vpcID, nil
}

func extractSubnetId(output map[string]string) (string, error) {
	subnetId, ok := output["subnet_id"]
	if !ok {
		return "", fmt.Errorf("vpc_id not found in output")
	}
	return subnetId, nil
}
