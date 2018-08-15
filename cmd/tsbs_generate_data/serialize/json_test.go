package serialize

import "testing"

func TestJSONSerializerSerialize(t *testing.T) {
	cases := []serializeCase{
		{
			desc:       "a regular Point",
			inputPoint: testPointDefault,
			output:     `{"_ts":1451606400000000000,"datacenter":"eu-west-1b","hostname":"host_0","region":"eu-west-1","usage_guest_nice":38.24311829}` + "\n"},
		{
			desc:       "a regular Point using int as value",
			inputPoint: testPointInt,
			output:     `{"_ts":1451606400000000000,"datacenter":"eu-west-1b","hostname":"host_0","region":"eu-west-1","usage_guest":38}` + "\n",
		},
		{
			desc:       "a regular Point with multiple fields",
			inputPoint: testPointMultiField,
			output:     `{"_ts":1451606400000000000,"big_usage_guest":5000000000,"datacenter":"eu-west-1b","hostname":"host_0","region":"eu-west-1","usage_guest":38,"usage_guest_nice":38.24311829}` + "\n",
		},
		{
			desc:       "a Point with no tags",
			inputPoint: testPointNoTags,
			output:     `{"_ts":1451606400000000000,"usage_guest_nice":38.24311829}` + "\n",
		},
	}

	testSerializer(t, cases, &JSONSerializer{})
}
