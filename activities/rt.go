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

func (a *AWSImpl) RTInitActivity(ctx context.Context, prov string) (string, error) {
	log.Println("Calling Route Table Init Activity with engine ", prov)

	provisioner, engine, dir := utils.GetProvisioner(prov)
	output, err := provisioner.Init(dir + utils.RtDir)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info("Calling Route Table Init Activity with engine ", engine)
	return output, nil
}

func (a *AWSImpl) RTApplyActivity(ctx context.Context, prov string, vpcid, subnetID, igwID, natID string) (string, error) {
	activity.GetLogger(ctx).Info("The vpc input to the RT activity is %s", vpcid)
	provisioner, engine, dir := utils.GetProvisioner(prov)
	output, err := provisioner.Apply(dir+utils.RtDir,
		"-var", fmt.Sprintf("vpc_id=%s", vpcid),
		"-var", fmt.Sprintf("subnet_id=%s", subnetID),
		"-var", fmt.Sprintf("igw_id=%s", igwID),
		"-var", fmt.Sprintf("nat_id=%s", natID),
	)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info("Calling Route Table Apply Activity with engine ", engine)
	return output, nil
}

func (a *AWSImpl) RTOutputActivity(ctx context.Context, prov string) (map[string]string, error) {
	log.Println("Calling Route Table Output Activity with engine ", prov)

	provisioner, engine, dir := utils.GetProvisioner(prov)
	outputValues, err := provisioner.Output(dir + utils.SubnetDir)
	if err != nil {
		return nil, err
	}
	activity.GetLogger(ctx).Info("Calling Route Table Output Activity with engine ", engine)

	var tfOutput map[string]models.CommonOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return nil, fmt.Errorf("error unmarshaling  output: %w", err)
	}

	rtOutput := map[string]string{
		"route_table_id":  tfOutput[utils.RTPUBLICID].Value,
		"route_table_arn": tfOutput[utils.RTTABLEARN].Value,
	}

	return rtOutput, nil
}
