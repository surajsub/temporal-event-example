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

func (a *AWSImpl) SGInitActivity(ctx context.Context, prov string) (string, error) {
	log.Println("Calling Security Init Activity with engine ", prov)

	provisioner, engine, dir := utils.GetProvisioner(prov)
	output, err := provisioner.Init(dir + utils.SgDir)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info("Calling Security Group Activity with engine ", engine)
	return output, nil
}

func (a *AWSImpl) SGApplyActivity(ctx context.Context, prov string, vpcid string) (string, error) {
	activity.GetLogger(ctx).Info("The vpc input to the Security Group is %s", vpcid)
	provisioner, engine, dir := utils.GetProvisioner(prov)
	output, err := provisioner.Apply(dir+utils.SgDir,
		"-var", fmt.Sprintf("vpc_id=%s", vpcid),
	)
	if err != nil {
		return "", err
	}
	activity.GetLogger(ctx).Info("Calling Subnet Apply Activity with engine ", engine)
	return output, nil
}

func (a *AWSImpl) SGOutputActivity(ctx context.Context, prov string) (map[string]string, error) {
	log.Println("Calling Security Output Activity with engine ", prov)

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

	sgOutput := map[string]string{
		"sg_id":  tfOutput[utils.SGID].Value,
		"sg_arn": tfOutput[utils.SGARN].Value,
	}

	log.Printf("THe values from the SG activity are %v\n", sgOutput)

	return sgOutput, nil
}
