package utils

import (
	"time"
)

const TASK_QUEUE_NAME = "event-driven-task-queue"

const WORKFLOW_TASK_QUEUE = "EKS_STACK_QUEUE"
const BASEOTDIRECTORY = "../opentofu"
const BASETFDIRECTORY = "../terraform"

const TemporalQueueName = "provision-task-queue"

// Terraform file locations

const VpcDir = "/vpc"
const VPCTIMEOUT = 1 * time.Minute

const SubnetDir = "/subnet"
const SubnetTimeOut = 2 * time.Minute

const IgwDir = "/igw"
const IgwTimeOut = 5 * time.Minute

const NatDir = "/nat"
const RtDir = "/rt"
const SgDir = "/sg"
const Ec2Dir = "/ec2"
const EKS_DIR = "/eks"

const VPC_RESOURCE = "VPC"
const SUBNET_RESOURCE = "Subnet"
const IGW_RESOURCE = "Internet Gateway"
const SG_RESOURCE = "Security Group"
const EC2_RESOURCE = "EC2 Instance"

const VpcWorkflow = "AWS VPC"
const IGW_WorkflowName = "AWS_Internet_Gateway"
const SubnetWorkflow = "AWS VPC Subnet"
const NatWorkflow = "AWS Nat Service"
const RtWorkflow = "AWS Route Table Service"
const SgWorkflow = "AWS Security Group"
const Ec2Workflow = "AWS EC2 Instance"

const EKS_WorkflowName = "AWS EKS"
const NODE_WorkflowName = "AWS EKS Nodes"

// Define the constants for the variables

const VpcInit = "Starting the VPC Init Activity:"
const SubnetInit = "Subnet Init Activity:"
const IgwInit = "Internet Gateway Init Activity:"
const NatInit = "NAT Init Activity"
const RtInit = "Route Table Init Activity"
const SgInit = "Security Group Init Activity:"
const Ec2Init = "EC2 Init Activity"
const EKS_INIT = "EKS Init Activity"
const NODE_INIT = "EKS Node Init Activity"

const VpcApply = "VPC Apply Activity:"
const SgApply = "Security Group Apply Activity:"
const IgwApply = "Internet Gateway Apply Activity"
const SubnetApply = "AWS Subnet Apply Activity"
const NatApply = "NAT Apply Activity"
const RtApply = "Route Table Apply Activity"
const Ec2Apply = "EC2 Apply Activity"
const EKS_APPLY = "EKS Apply Activity:"
const NODE_APPLY = "EKS Node Apply Activity:"

const VpcOutput = "VPC Output Activity"
const SubnetOutput = "Subnet Output Activity"
const IgwOutput = "Internet Gateway Output Activity"
const NatOutput = "Network Gateway Output Activity"
const RtOutput = "Route Table Output Activity"
const SgOutput = "Security Group Output Activity"
const Ec2Output = "EC2 Output Activity"

const VPCID = "vpc_id"
const VPCCIDR = "vpc_cidr_block"

const SUBNETID = "subnet_id"
const SUBNETARN = "subnet_arn"

const PRIVATE_SUBNET_ID = "private_subnet_id"
const PUBLIC_SUBNET_ID = "public_subnet_id"

const IGWID = "igw_id"
const IGWARN = "igw_arn"

const SGID = "sg_id"
const SGARN = "sg_arn"

const NATID = "nat_id"
const NATGATEWAYID = "nat_gateway_id"
const NATALLOCATIONID = "nat_allocation_id"

const RTPRIVATEID = "rt_private_id"
const RTPUBLICID = "rt_public_id"
const RTTABLEARN = "rt_table_arn"

// These need to map to the output from the tf file
const EC2ID = "instance_id"
const EC2PUBLIC = "instance_public_ip"

const EKS_ID = "eks_id"
const EKS_ARN = "eks_arn"
const EKS_ENDPOINT = "eks_endpoint"

const ENGINE = "engine"
const TERRAFORM = "terraform"
const OPENTOFU = "tofu"

// NATS Variables
const NATS_AWS_PROVISION_QUEUE = "provision-aws"
const NATS_AZ_PROVISION_QUEUE = "provision-az"
