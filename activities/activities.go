package activities

import (
	"context"
)

type AWSImpl struct {
}

type AWSImplActivities interface {
	VPCInitActivity(ctx context.Context, prov string) (string, error)
	VPCApplyActivity(ctx context.Context, prov string, name string, vpc string) (string, error)
	VPCOutputActivity(ctx context.Context, prov string) (map[string]string, error)
	// Subnet
	SubnetInitActivity(ctx context.Context, prov string) (string, error)
	SubnetApplyActivity(ctx context.Context, prov, vpcid, subnetname, subnetsize string) (string, error)
	SubnetOutputActivity(ctx context.Context, prov string) (map[string]string, error)
	// NAT
	NATInitActivity(ctx context.Context, prov string) (string, error)
	NATApplyActivity(ctx context.Context, prov, subnetId, igwName string) (string, error)
	NATOutputActivity(ctx context.Context, prov string) (map[string]string, error)
	// IGW
	IGWInitActivity(ctx context.Context, prov string) (string, error)
	IGWApplyActivity(ctx context.Context, prov, vpcid, name string) (string, error)
	IGWOutputActivity(ctx context.Context, prov string) (map[string]string, error)
	// RT
	RTInitActivity(ctx context.Context, prov string) (string, error)
	RTApplyActivity(ctx context.Context, prov string, vpcid string) (string, error)
	RTOutputActivity(ctx context.Context, prov string) (map[string]string, error)
	//SG
	SGInitActivity(ctx context.Context, prov string) (string, error)
	SGApplyActivity(ctx context.Context, prov string, vpcid string) (string, error)
	SGOutputActivity(ctx context.Context, prov string) (map[string]string, error)

	// EC2
	EC2InitActivity(ctx context.Context, prov string) (string, error)
	EC2ApplyActivity(ctx context.Context, prov, subnetId, sgId, instanceName string) (string, error)
	EC2OutputActivity(ctx context.Context, prov string) (map[string]string, error)
}
