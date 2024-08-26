package models

// ProvisionMessage is the structure of the message to be sent
type ProvisionMessage struct {
	ResourceType string            `json:"resource_type"`
	Params       map[string]string `json:"params"`
}

type VPCParams struct {
	Name string
	CIDR string
}

type SubnetParams struct {
	VPCID      string
	SubnetName string
	SubnetSize string
}

type IGWParams struct {
	VPCID   string
	IGWName string
}

type NatParams struct {
	SubnetId string
	NatName  string
}

type RTParams struct {
	VPCID    string
	NatID    string
	IGWID    string
	SubnetID string
}

type EC2Params struct {
}

type SGParams struct {
	VPCID string
}
type VPCMessage struct {
	VPCName            string `json:"vpc_name"`
	ProvisioningEngine string `json:"provisioning_engine"`
	CIDRBlock          string `json:"cidr_block"`
}

type CommonOutput struct {
	Value string `json:"value"`
}

type VPCOutput struct {
	Value string `json:"value"`
}

type VPCApplyOutput struct {
	VPCID   string `json:"vpc_id"`
	VPCCIDR string `json:"vpc_cidr_block"`
}

type VPCCommonOutput struct {
	Value string `json:"value"`
}

type SubnetCommonOutput struct {
	Value string `json:"value"`
}

type SubnetApplyOutput struct {
	SubnetId   string `json:"subnet_id"`
	SubnetArn  string `json:"subnet_arn"`
	SubnetCIDR string `json:"subnet_cidr"`
}

// IGW

type IGWApplyOutput struct {
	IGWId  string `json:"igw_id"`
	IGWArn string `json:"igw_arn"`
}

type SGApplyOutput struct {
	SgID  string `json:"sg_id"`
	SgArn string `json:"sg_arn"`
}

type NATApplyOutput struct {
	NatID           string `json:"nat_id"`
	NatGateway      string `json:"nat_gateway_id"`
	NatAllocationID string `json:"nat_allocation_id"`
}

type EC2ApplyOutput struct {
	//EC2 Apply Output Structure.
	InstanceID       string `json:"instance_id"`
	InstancePublicIP string `json:"instance_public_ip"`
}

type RTApplyOutput struct {
	RouteTableID      string `json:"route_table_id"`
	RouteTableARN     string `json:"route_table_arn"`
	RouteTableOwnerID string `json:"route_table_owner_id"`
}
