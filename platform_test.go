package main

import (
	"reflect"
	"testing"
)

func TestSupportedPlatforms(t *testing.T) {
	var ps []Platform

	ps = SupportedPlatforms("go1.0")
	if !reflect.DeepEqual(ps, Platforms_1_0) {
		t.Fatalf("bad: %#v", ps)
	}

	ps = SupportedPlatforms("go1.1")
	if !reflect.DeepEqual(ps, Platforms_1_1) {
		t.Fatalf("bad: %#v", ps)
	}

	ps = SupportedPlatforms("go1.2")
	if !reflect.DeepEqual(ps, Platforms_1_1) {
		t.Fatalf("bad: %#v", ps)
	}

	ps = SupportedPlatforms("go1.3")
	if !reflect.DeepEqual(ps, Platforms_1_3) {
		t.Fatalf("bad: %#v", ps)
	}

	ps = SupportedPlatforms("go1.4")
	if !reflect.DeepEqual(ps, Platforms_1_4) {
		t.Fatalf("bad: %#v", ps)
	}

	ps = SupportedPlatforms("go1.5")
	if !reflect.DeepEqual(ps, Platforms_1_5) {
		t.Fatalf("bad: %#v", ps)
	}

	ps = SupportedPlatforms("go1.6")
	if !reflect.DeepEqual(ps, Platforms_1_6) {
		t.Fatalf("bad: %#v", ps)
	}

	ps = SupportedPlatforms("go1.7")
	if !reflect.DeepEqual(ps, Platforms_1_7) {
		t.Fatalf("bad: %#v", ps)
	}

	ps = SupportedPlatforms("go1.8")
	if !reflect.DeepEqual(ps, Platforms_1_8) {
		t.Fatalf("bad: %#v", ps)
	}

	ps = SupportedPlatforms("go1.9")
	if !reflect.DeepEqual(ps, Platforms_1_9) {
		t.Fatalf("bad: %#v", ps)
	}

	ps = SupportedPlatforms("go1.10")
	if !reflect.DeepEqual(ps, Platforms_1_10) {
		t.Fatalf("bad: %#v", ps)
	}

	// Unknown
	ps = SupportedPlatforms("foo")
	if !reflect.DeepEqual(ps, PlatformsLatest) {
		t.Fatalf("bad: %#v", ps)
	}
}

func TestMIPS(t *testing.T) {
	g16 := SupportedPlatforms("go1.6")
	for _, p := range g16 {
		if p.Arch == "mips64" && p.Default {
			t.Fatal("mips64 should not be default for 1.6")
		}
	}

	g17 := SupportedPlatforms("go1.7")
	for _, p := range g17 {
		if p.Arch == "mips64" && !p.Default {
			t.Fatal("mips64 should be default for 1.7")
		}
	}

}

func Test_removeElement(t *testing.T) {
	type args struct {
		from     []Platform
		elements []Platform
	}
	tests := []struct {
		name string
		args args
		want []Platform
	}{
		{
			name: "removing existing element",
			args: args{
				from: []Platform{
					{"windows", "386", false},
					{"windows", "amd64", true},
					{"windows", "arm64", true},
					{"linux", "386", false},
					{"linux", "amd64", true},
					{"linux", "arm64", true},
				},
				elements: []Platform{{"windows", "arm64", true}},
			},
			want: []Platform{
				{"windows", "386", false},
				{"windows", "amd64", true},
				{"linux", "386", false},
				{"linux", "amd64", true},
				{"linux", "arm64", true},
			}},
		{
			name: "removing element that doesn't exist",
			args: args{
				from: []Platform{
					{"windows", "386", false},
					{"windows", "amd64", true},
					{"windows", "arm64", true},
					{"linux", "riscv64", false},
					{"linux", "amd64", true},
					{"linux", "arm64", true},
				},
				elements: []Platform{{"linux", "386", true}},
			},
			want: []Platform{
				{"windows", "386", false},
				{"windows", "amd64", true},
				{"windows", "arm64", true},
				{"linux", "riscv64", false},
				{"linux", "amd64", true},
				{"linux", "arm64", true},
			},
		},
		{
			name: "removing multiple elements",
			args: args{
				from: []Platform{
					{"windows", "386", false},
					{"windows", "amd64", true},
					{"windows", "arm64", true},
					{"linux", "riscv64", false},
					{"linux", "amd64", true},
					{"linux", "arm64", true},
				},
				elements: []Platform{
					{"windows", "386", true},
					{"windows", "amd64", true},
					{"linux", "arm64", true},
				},
			},
			want: []Platform{
				{"windows", "arm64", true},
				{"linux", "riscv64", false},
				{"linux", "amd64", true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.from, tt.args.elements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
