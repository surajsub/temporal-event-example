package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/temporal-event-example/models"
	"github.com/surajsub/temporal-event-example/utils"
	"go.temporal.io/sdk/activity"
)

// This is the common EC2 Provisioner code

func (a *AWSImpl) EC2InitActivity(ctx context.Context, prov string) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(prov)
	output, err := provisioner.Init(dir + utils.Ec2Dir)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info("Calling EC2 Init Activity with engine ", engine)
	return output, nil
}

func (a *AWSImpl) EC2ApplyActivity(ctx context.Context, prov, subnetId, sgId, instanceName string) (string, error) {
	activity.GetLogger(ctx).Info("The vpc input to the subnet is %s", subnetId)
	provisioner, engine, dir := utils.GetProvisioner(prov)
	output, err := provisioner.Apply(dir+utils.Ec2Dir,
		"-var", fmt.Sprintf("subnet_id=%s", subnetId),
		"-var", fmt.Sprintf("sg_id=%s", sgId),
		"-var", fmt.Sprintf("instance_name=%s", instanceName),
	)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info("Calling EC2 Apply Activity with engine ", engine)
	return output, nil
}
func (a *AWSImpl) EC2OutputActivity(ctx context.Context, prov string) (map[string]string, error) {
	provisioner, engine, dir := utils.GetProvisioner(prov)
	outputValues, err := provisioner.Output(dir + utils.Ec2Dir)
	if err != nil {
		return nil, err
	}
	activity.GetLogger(ctx).Info("Calling EC2 Output Activity with engine ", engine)

	var tfOutput map[string]models.CommonOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling  output: %w", err)
	}

	ec2Output := map[string]string{
		"ec2_id":       tfOutput[utils.EC2ID].Value,
		"ec2_instance": tfOutput[utils.EC2PUBLIC].Value,
	}

	return ec2Output, nil
}
