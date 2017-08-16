package resources

// AWSEC2Instance_NoDevice AWS CloudFormation Resource (AWS::EC2::Instance.NoDevice)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-nodevice.html
type AWSEC2Instance_NoDevice struct {
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance_NoDevice) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.NoDevice"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2Instance_NoDevice) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}