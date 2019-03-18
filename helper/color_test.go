package helper

import (
	"testing"
)

func TestColorForStatus(t *testing.T) {
	tests := []struct {
		name string
		args int
		want string
	}{
		{
			name: "Testcase #1: Return Green",
			args: 200,
			want: Green,
		},
		{
			name: "Testcase #2: Return White",
			args: 300,
			want: White,
		},
		{
			name: "Testcase #3: Return Yellow",
			args: 400,
			want: Yellow,
		},
		{
			name: "Testcase #4: Return Red",
			args: 500,
			want: Red,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ColorForStatus(tt.args); got != tt.want {
				t.Errorf("ColorForStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColorForMethod(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Testcase #1: Return Blue",
			args: "GET",
			want: Blue,
		},
		{
			name: "Testcase #2: Return Cyan",
			args: "POST",
			want: Cyan,
		},
		{
			name: "Testcase #3: Return Yellow",
			args: "PUT",
			want: Yellow,
		},
		{
			name: "Testcase #4: Return Red",
			args: "DELETE",
			want: Red,
		},
		{
			name: "Testcase #5: Return Green",
			args: "PATCH",
			want: Green,
		},
		{
			name: "Testcase #6: Return Magenta",
			args: "HEAD",
			want: Magenta,
		},
		{
			name: "Testcase #7: Return White",
			args: "OPTIONS",
			want: White,
		},
		{
			name: "Testcase #8: Return RESET",
			args: "RESET",
			want: Reset,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ColorForMethod(tt.args); got != tt.want {
				t.Errorf("ColorForMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColorString(t *testing.T) {
	t.Run("Test String Red", func(t *testing.T) {
		if got := StringRed("String Red"); got != "\x1b[31;1mString Red\x1b[0m" {
			t.Errorf("StringRed() = %v, want %v", got, "String Red")
		}
	})

	t.Run("Test String Green", func(t *testing.T) {
		if got := StringGreen("String Green"); got != "\x1b[32;1mString Green\x1b[0m" {
			t.Errorf("StringGreen() = %v, want %v", got, "String Green")
		}
	})

	t.Run("Test String Yellow", func(t *testing.T) {
		if got := StringYellow("String Yellow"); got != "\x1b[33;1mString Yellow\x1b[0m" {
			t.Errorf("StringYellow() = %v, want %v", got, "String Yellow")
		}
	})
}
