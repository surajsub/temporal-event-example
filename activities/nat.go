package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/surajsub/temporal-event-example/models"
	"github.com/surajsub/temporal-event-example/utils"
	"go.temporal.io/sdk/activity"
)

// This is the common NAT provisioner

func (a *AWSImpl) NATInitActivity(ctx context.Context, prov string) (string, error) {
	provisioner, engine, dir := utils.GetProvisioner(prov)
	output, err := provisioner.Init(dir + utils.NatDir)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info("Calling IGW Init Activity with engine ", engine)
	return output, nil
}

func (a *AWSImpl) NATApplyActivity(ctx context.Context, prov, subnetId, name string) (string, error) {
	activity.GetLogger(ctx).Info("The vpc input to the subnet is %s", subnetId)
	provisioner, engine, dir := utils.GetProvisioner(prov)
	output, err := provisioner.Apply(dir+utils.NatDir,
		"-var", fmt.Sprintf("subnet_id=%s", subnetId),
		"-var", fmt.Sprintf("nat_name=%s", name),
	)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info("Calling NAT Apply Activity with engine ", engine)
	return output, nil
}
func (a *AWSImpl) NATOutputActivity(ctx context.Context, prov string) (map[string]string, error) {
	provisioner, engine, dir := utils.GetProvisioner(prov)
	outputValues, err := provisioner.Output(dir + utils.NatDir)
	if err != nil {
		return nil, err
	}
	activity.GetLogger(ctx).Info("Calling NAT Output Activity with engine ", engine)

	var tfOutput map[string]models.CommonOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling  output: %w", err)
	}

	natOutput := map[string]string{
		"nat_id":            tfOutput[utils.NATID].Value,
		"nat_gateway_id":    tfOutput[utils.NATGATEWAYID].Value,
		"nat_allocation_id": tfOutput[utils.NATALLOCATIONID].Value,
	}

	return natOutput, nil
}
