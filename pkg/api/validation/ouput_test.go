package validation

import (
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"reflect"
	"testing"
)

func Test_transToConfig(t *testing.T) {
	type args struct {
		cfg interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    v1.Config
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "file_output",
			args: args{
				kubev1.FileConfig{
					File:      "demo_file",
					Dir:       "demo_Dir",
					Expire:    0,
					Formatter: kubev1.FormatterConfig{"a": {"1", "2", "3"}},
				},
			},
			want: v1.Config{
				"file":      "demo_file",
				"dir":       "demo_Dir",
				"period":    "",
				"expire":    0,
				"type":      "",
				"formatter": map[string][]string{"a": {"1", "2", "3"}},
			},
			wantErr: false,
		},
		{
			name: "nsqd",
			args: args{
				kubev1.NsqdConfig{
					Topic:      "demo_topic",
					Address:    nil,
					ClientConf: kubev1.Config{"a": "b"},
				},
			},
			want: v1.Config{
				"topic":     "demo_topic",
				"address":   nil,
				"nsq_conf":  map[string]interface{}{"a": "b"},
				"type":      "",
				"formatter": nil,
			},
			wantErr: false,
		},
		{
			name: "http_output",
			args: args{
				kubev1.HttpConfig{
					Method:    "GET",
					Url:       "127.0.0.1:4444/output",
					Type:      "",
					Formatter: nil,
				},
			},
			want: v1.Config{
				"method":    "GET",
				"url":       "127.0.0.1:4444/output",
				"headers":   nil,
				"type":      "",
				"formatter": nil,
			},
			wantErr: false,
		},
		{
			name: "syslog_output",
			args: args{
				kubev1.SysConfig{
					Network:   "demo",
					Address:   "demo",
					Level:     "demo",
					Type:      "demo",
					Formatter: nil,
				},
			},
			want: v1.Config{
				"network":   "demo",
				"address":   "demo",
				"level":     "demo",
				"tpye":      "demo",
				"formatter": nil,
			},
			wantErr: false,
		},
		{
			name: "kafka_output",
			args: args{
				kubev1.KafkaConfig{
					Topic:         "topic",
					Address:       "address",
					Version:       "",
					PartitionType: "",
					Partition:     0,
					Formatter:     nil,
				},
			},
			want: v1.Config{
				"topic":           "topic",
				"address":         "address",
				"timeout":         0,
				"version":         "",
				"paritition_type": "",
				"partition":       0,
				"partition_key":   "",
				"type":            "",
				"formatter":       nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := transToConfig(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("transToConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transToConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
