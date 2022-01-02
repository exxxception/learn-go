package pkg

import "testing"

func TestToCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	testTable := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Simple Test",
			args: args{
				s: "the-stealth-warrior",
			},
			want: "theStealthWarrior",
		},
	}
	for _, testCase := range testTable {
		result := ToCamelCase(testCase.args.s)
		if result != testCase.want {
			t.Errorf("Incorrect result. Expect %s, got %s", testCase.want, result)
		}
	}
}
