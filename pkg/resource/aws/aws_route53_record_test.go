package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

func TestAwsRoute53Record_String(t *testing.T) {
	tests := []struct {
		name   string
		record AwsRoute53Record
		want   string
	}{
		{name: "test route53 record stringer with fqdn and type and zoneid",
			record: AwsRoute53Record{
				Fqdn:   aws.String("true"),
				Type:   aws.String("type"),
				ZoneId: aws.String("zone_id"),
			},
			want: "true (type) (zone: zone_id)",
		},
		{name: "test route53 record stringer without values",
			record: AwsRoute53Record{
				Fqdn: nil,
			},
			want: " () (zone: )",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.record.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
