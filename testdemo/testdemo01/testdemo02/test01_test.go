package main
import(
	_"fmt"
	"testing"
)
func TestAddpper(t *testing.T){
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("yuqizhi %v,fanhuizhi %v",55,res)
	}
	t.Logf("ok")
}