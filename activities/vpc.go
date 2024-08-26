package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/temporal-event-example/models"
	"github.com/surajsub/temporal-event-example/utils"
	"go.temporal.io/sdk/activity"
	"log"
)

// This is the common vpc provisioner

func (a *AWSImpl) VPCInitActivity(ctx context.Context, prov string) (string, error) {
	log.Println("Calling Init Activity with engine ", prov)

	provisioner, engine, dir := utils.GetProvisioner(prov)
	output, err := provisioner.Init(dir + utils.VpcDir)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info("Calling Init Activity with engine ", engine)
	return output, nil
}

func (a *AWSImpl) VPCApplyActivity(ctx context.Context, prov string, name string, vpc string) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(prov)
	output, err := provisioner.Apply(dir+utils.VpcDir,
		"-var", fmt.Sprintf("cidr_block=%s", vpc),
		"-var", fmt.Sprintf("name=%s", name),
	)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info("Calling Init Activity with engine ", engine)
	return output, nil
}

func (a *AWSImpl) VPCOutputActivity(ctx context.Context, prov string) (map[string]string, error) {
	provisioner, engine, dir := utils.GetProvisioner(prov)
	outputValues, err := provisioner.Output(dir + utils.VpcDir)
	if err != nil {
		return nil, err
	}
	activity.GetLogger(ctx).Info("Calling Init Activity with engine ", engine)

	var tfOutput map[string]models.CommonOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling terraform output: %w", err)
	}

	vpcOutput := map[string]string{
		"vpc_id":         tfOutput[utils.VPCID].Value,
		"vpc_cidr_block": tfOutput[utils.VPCCIDR].Value,
	}

	return vpcOutput, nil
}

func VPCDestroyActivity(ctx context.Context, prov string, name string, vpc string) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(prov)
	output, err := provisioner.Destroy(dir+utils.VpcDir,
		"-var", fmt.Sprintf("cidr_block=%s", vpc),
		"-var", fmt.Sprintf("name=%s", name),
	)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info("Calling Destroy Activity with engine ", engine)
	return output, nil
}
