package discord

import "testing"

func Test_stringHasPrefix(t *testing.T) {
	type args struct {
		str        string
		prefixes   []string
		ignoreCase bool
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		{
			"Exist",
			args{
				"!command help arg",
				[]string{"cmd!", "!command", "!cmd"},
				false,
			},
			true,
			"help arg",
		},
		{
			"Not exist",
			args{
				"!notcommand help arg",
				[]string{"cmd!", "!command", "!cmd"},
				false,
			},
			false,
			"!notcommand help arg",
		},
		{
			"Ignore case",
			args{
				"!Command help arg",
				[]string{"cmd!", "!commanD", "!cmd"},
				true,
			},
			true,
			"help arg",
		},
		{
			"Not ignore case",
			args{
				"!Command help arg",
				[]string{"cmd!", "!commanD", "!cmd"},
				false,
			},
			false,
			"!Command help arg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := stringHasPrefix(tt.args.str, tt.args.prefixes, tt.args.ignoreCase)
			if got != tt.want {
				t.Errorf("stringHasPrefix() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("stringHasPrefix() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_stringArrayContains(t *testing.T) {
	type args struct {
		array      []string
		str        string
		ignoreCase bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Exist",
			args{
				[]string{"command", "arg1", "arg2"},
				"arg1",
				false,
			},
			true,
		},
		{
			"Not exist",
			args{
				[]string{"command", "arg1", "arg2"},
				"arg3",
				false,
			},
			false,
		},
		{
			"Ignore case",
			args{
				[]string{"command", "arg1", "arg2"},
				"Arg1",
				true,
			},
			true,
		},
		{
			"Not ignore case",
			args{
				[]string{"command", "arg1", "arg2"},
				"Arg1",
				false,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringArrayContains(tt.args.array, tt.args.str, tt.args.ignoreCase); got != tt.want {
				t.Errorf("stringArrayContains() = %v, want %v", got, tt.want)
			}
		})
	}
}
