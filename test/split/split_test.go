package split

import (
	"testing"
	"reflect"
)

// 测试

func TestSplit(t *testing.T){
	got := Split("我爱你", "爱")
	want := []string{"我", "你"}
	// got := Split("a:b", ":")
	// want := []string{"a", "b"}
	
	if !reflect.DeepEqual(got, want){
		t.Errorf("want:%v  got:%v", want, got)
	
	}
}