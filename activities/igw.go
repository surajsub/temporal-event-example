package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/temporal-event-example/models"
	"github.com/surajsub/temporal-event-example/utils"
	"go.temporal.io/sdk/activity"
)

// This is the common IGW provisioner

func (a *AWSImpl) IGWInitActivity(ctx context.Context, prov string) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(prov)
	output, err := provisioner.Init(dir + utils.IgwDir)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info("Calling IGW Init Activity with engine ", engine)
	return output, nil
}

func (a *AWSImpl) IGWApplyActivity(ctx context.Context, prov, vpcid, name string) (string, error) {
	activity.GetLogger(ctx).Info("The vpc input to the subnet is %s", vpcid)
	provisioner, engine, dir := utils.GetProvisioner(prov)
	output, err := provisioner.Apply(dir+utils.IgwDir,
		"-var", fmt.Sprintf("vpc_id=%s", vpcid),
		"-var", fmt.Sprintf("igw_name=%s", name),
	)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info("Calling IGW Apply Activity with engine ", engine)
	return output, nil
}

func (a *AWSImpl) IGWOutputActivity(ctx context.Context, prov string) (map[string]string, error) {
	provisioner, engine, dir := utils.GetProvisioner(prov)
	outputValues, err := provisioner.Output(dir + utils.IgwDir)
	if err != nil {
		return nil, err
	}
	activity.GetLogger(ctx).Info("Calling IGW Output Activity with engine ", engine)

	var tfOutput map[string]models.CommonOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling  output: %w", err)
	}

	igwOutput := map[string]string{
		"igw_id":  tfOutput[utils.IGWID].Value,
		"igw_arn": tfOutput[utils.IGWARN].Value,
	}

	return igwOutput, nil
}
