package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/temporal-event-example/models"
	"github.com/surajsub/temporal-event-example/utils"
	"go.temporal.io/sdk/activity"
)

// This is the common subnet provisioner

func (a *AWSImpl) SubnetInitActivity(ctx context.Context, prov string) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(prov)
	output, err := provisioner.Init(dir + utils.SubnetDir)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info("Calling Subnet Init Activity with engine ", engine)
	return output, nil
}

// SubnetApplyActivity /* TODO - For subnet also pass the availability zone. We need to make sure the zone is in the same region as the vpc is created

func (a *AWSImpl) SubnetApplyActivity(ctx context.Context, prov, vpcid, subnetname, subnetsize string) (string, error) {
	activity.GetLogger(ctx).Info("The vpc input to the subnet is %s", vpcid)
	provisioner, engine, dir := utils.GetProvisioner(prov)
	output, err := provisioner.Apply(dir+utils.SubnetDir,
		"-var", fmt.Sprintf("vpc_id=%s", vpcid),
		"-var", fmt.Sprintf("subnet_name=%s", subnetname),
		"-var", fmt.Sprintf("subnet_size=%s", subnetsize),
	)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info("Calling Subnet Apply Activity with engine ", engine)
	return output, nil
}

func (a *AWSImpl) SubnetOutputActivity(ctx context.Context, prov string) (map[string]string, error) {
	provisioner, engine, dir := utils.GetProvisioner(prov)
	outputValues, err := provisioner.Output(dir + utils.SubnetDir)
	if err != nil {
		return nil, err
	}
	activity.GetLogger(ctx).Info("Calling Subnet Output Activity with engine ", engine)

	var tfOutput map[string]models.CommonOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling  output: %w", err)
	}

	subnetOutput := map[string]string{
		"subnet_id":  tfOutput[utils.SUBNETID].Value,
		"subnet_arn": tfOutput[utils.SUBNETARN].Value,
	}

	return subnetOutput, nil
}
