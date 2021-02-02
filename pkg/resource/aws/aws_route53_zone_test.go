package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

func TestAwsRoute53Zone_String(t *testing.T) {
	tests := []struct {
		name string
		zone AwsRoute53Zone
		want string
	}{
		{name: "test route53 zone stringer with name and id",
			zone: AwsRoute53Zone{
				Name: aws.String("example.com"),
				Id:   "01",
			},
			want: "example.com (01)",
		},
		{name: "test route53 zone stringer without name",
			zone: AwsRoute53Zone{
				Name: nil,
				Id:   "02",
			},
			want: " (02)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.zone.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
