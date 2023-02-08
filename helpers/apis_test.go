package helpers

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestValidateIp(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Empty string", args{""}, false},
		{"IP with minus", args{"185.-1.102.248"}, false},
		{"IP with char", args{"185.#.102.248"}, false},
		{"IP with string", args{"185.go.102.248"}, false},
		{"No valid IP", args{"270.9.9.9"}, false},
		{"No valid IP2", args{"192.168.000.254"}, false},
		{"No valid IP3", args{"192.168.0.257"}, false},
		{"Valid IP", args{"185.220.102.248"}, true},
		{"Valid IP2", args{"5.2.69.50"}, true},
		{"Valid IP3", args{"185.220.101.1"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateIp(tt.args.ip); got != tt.want {
				t.Errorf("ValidateIp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateStr(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Empty string", args{""}, true},
		{"IP with minus", args{"185.-1.102.248"}, false},
		{"IP with char", args{"185.#.102.248"}, false},
		{"IP with string", args{"185.go.102.248"}, false},
		// here the user should have the opportunity to enter an invalid ip
		{"No valid IP", args{"270.9.9.9"}, true},
		{"No valid IP2", args{"192.168.000.254"}, true},
		{"No valid IP3", args{"192.168.0.257"}, true},
		{"Valid IP", args{"185.220.102.248"}, true},
		{"Valid IP2", args{"5.2.69.50"}, true},
		{"Valid IP3", args{"185.220.101.1"}, true},
		{"Part ip", args{"0.101.1"}, true},
		{"Part ip2", args{"0.101.1"}, true},
		{"Part ip3", args{"1"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateStr(tt.args.ip); got != tt.want {
				t.Errorf("ValidateStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemoryData_IpExists(t *testing.T) {
	var validData = []string{"185.220.102.248", "5.2.69.50"}
	var NovalidData = []string{"185.38.175.132", ""}

	type fields struct {
		RWMutex sync.RWMutex
		URL     string
		Ips     map[string]bool
		Changes map[string][]string
	}
	type args struct {
		ip []string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      bool
		wantfalse bool
	}{
		{
			name: "TestMemoryData_IpExists:True:False",
			fields: fields{
				RWMutex: sync.RWMutex{},
				URL:     "",
				Ips:     map[string]bool{"185.220.102.248": true, "5.2.69.50": true},
				Changes: map[string][]string{},
			},
			args: args{validData},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryData{
				RWMutex: tt.fields.RWMutex,
				URL:     tt.fields.URL,
				Ips:     tt.fields.Ips,
				Changes: tt.fields.Changes,
			}
			for _, value := range validData {
				if got := m.IpExists(value); got != tt.want {
					t.Errorf("MemoryData.IpExists() validData = %v, want %v", got, tt.want)
				}
			}
			for _, value := range NovalidData {
				if got := m.IpExists(value); got != tt.wantfalse {
					t.Errorf("MemoryData.IpExists() NovalidData = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestMemoryData_MatchTrue(t *testing.T) {
	var validData = []string{".22", "", "4", "."}
	var oneValue string = "185.220.102.248"
	type fields struct {
		RWMutex sync.RWMutex
		URL     string
		Ips     map[string]bool
		Changes map[string][]string
	}
	type args struct {
		ip []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantOne []string
	}{
		{
			name: "TestMemoryData_Match:True",
			fields: fields{
				RWMutex: sync.RWMutex{},
				URL:     "",
				Ips:     map[string]bool{"185.220.102.248": true, "4.22.188.50": true},
				Changes: map[string][]string{},
			},
			args:    args{validData},
			want:    []string{"185.220.102.248", "4.22.188.50"},
			wantOne: []string{"185.220.102.248"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryData{
				RWMutex: tt.fields.RWMutex,
				URL:     tt.fields.URL,
				Ips:     tt.fields.Ips,
				Changes: tt.fields.Changes,
			}
			for _, value := range validData {
				if got := m.Match(value); !reflect.DeepEqual(got, tt.want) && !reflect.DeepEqual(got, []string{"4.22.188.50", "185.220.102.248"}) {
					t.Errorf("MemoryData.Match1() = %v, want %v, value:= %v", got, tt.want, value)
				}
			}
			if got := m.Match(oneValue); !reflect.DeepEqual(got, tt.wantOne) {
				t.Errorf("MemoryData.Match1 One() = %v, want %v, value:= %v", got, tt.want, oneValue)
			}

		})
	}
}

func TestMemoryData_MatchFalse(t *testing.T) {
	var validData = []string{".102.1", "9", "a", "#"}
	type fields struct {
		RWMutex sync.RWMutex
		URL     string
		Ips     map[string]bool
		Changes map[string][]string
	}
	type args struct {
		ip []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "TestMemoryData_Match:False",
			fields: fields{
				RWMutex: sync.RWMutex{},
				URL:     "",
				Ips:     map[string]bool{"185.220.102.248": true, "4.22.188.50": true},
				Changes: map[string][]string{},
			},
			args: args{validData},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryData{
				RWMutex: tt.fields.RWMutex,
				URL:     tt.fields.URL,
				Ips:     tt.fields.Ips,
				Changes: tt.fields.Changes,
			}
			for _, value := range validData {
				if got := m.Match(value); reflect.DeepEqual(got, tt.want) {
					t.Errorf("MemoryData.Match() = %v, want %v, value:= %v", got, tt.want, value)
				}
			}

		})
	}
}

func TestMemoryData_ChangeApis(t *testing.T) {
	var wg sync.WaitGroup
	type fields struct {
		RWMutex sync.RWMutex
		URL     string
		Ips     map[string]bool
		Changes map[string][]string
	}
	type args struct {
		data JsonReader
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string][]string
	}{
		{
			name: "TestMemoryData_ChangeApis_Added_Only",
			fields: fields{
				RWMutex: sync.RWMutex{},
				URL:     "",
				Ips:     map[string]bool{"185.220.102.248": true, "4.22.188.50": true},
				Changes: map[string][]string{},
			},
			args: args{JsonReader{IpAddress: "", Add: []string{"4.22.188.50"}, Remove: []string{""}}},
			want: map[string][]string{"added": {"4.22.188.50"}},
		},
		{
			name: "TestMemoryData_ChangeApis_Added_Removed",
			fields: fields{
				RWMutex: sync.RWMutex{},
				URL:     "",
				Ips:     map[string]bool{"185.220.102.248": true, "4.22.188.50": true},
				Changes: map[string][]string{},
			},
			args: args{JsonReader{IpAddress: "", Add: []string{"4.22.188.51", "4.22.188.52"}, Remove: []string{"4.22.188.50"}}},
			want: map[string][]string{"added": {"4.22.188.51", "4.22.188.52"}, "removed": {"4.22.188.50"}},
		},
		{
			name: "TestMemoryData_ChangeApis_Added_Removed_",
			fields: fields{
				RWMutex: sync.RWMutex{},
				URL:     "",
				Ips:     map[string]bool{"185.220.102.248": true, "4.22.188.50": true},
				Changes: map[string][]string{},
			},
			args: args{JsonReader{IpAddress: "", Add: []string{"4.22.188.51", "4.22.188.52"}, Remove: []string{"4.22.188.52"}}},
			want: map[string][]string{"added": {"4.22.188.51", "4.22.188.52"}, "removed": {"4.22.188.52"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryData{
				RWMutex: tt.fields.RWMutex,
				URL:     tt.fields.URL,
				Ips:     tt.fields.Ips,
				Changes: tt.fields.Changes,
			}
			wg.Add(1)
			m.ChangeApis(tt.args.data)
			wg.Done()
			fmt.Println(m.Changes)
			if got := m.Changes; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemoryData.Match1() = %v, want %v", got, tt.want)
			}
		})
	}
}
