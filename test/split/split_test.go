package split

import (
	"testing"
	"reflect"
)

// 测试

// func TestSplit(t *testing.T){
// 	got := Split("我爱你", "爱")
// 	want := []string{"我", "你"}
// 	// got := Split("a:b", ":")
// 	// want := []string{"a", "b"}
	
// 	if !reflect.DeepEqual(got, want){
// 		t.Errorf("want:%v  got:%v", want, got)
	
// 	}
// }


func TestSplit(t *testing.T){
	type test struct{
		input string
		sep string
		want []string
	}

	tests := map[string]test{
		"simple":test{ 
			input: "我爱你",
			sep:"爱",
			want:[]string{"我", "你"},
		},
		"multi sep":test{ 
			input: "a:b:c",
			sep:":",
			want:[]string{"a", "b", "c"},
		},

	}
	for name, tc := range tests{
		t.Run(name, func(t *testing.T){
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want){
				t.Errorf("name:%v failed, want:%v got:%v", name, tc.want, got)
		}
		})
	}

}