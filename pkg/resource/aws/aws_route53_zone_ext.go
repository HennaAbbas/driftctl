package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
)

func (r *AwsRoute53Zone) String() string {
	return fmt.Sprintf("%s (%s)", aws.StringValue(r.Name), r.Id)
}
