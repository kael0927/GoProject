package monster

import (
	"testing"
)

// 测试用例 调用Store方法
func TestStore(t *testing.T) {
	monster := &Monster{
		Name:  "JT",
		Age:   20,
		Skill: "sing",
	}
	ok := monster.Store()
	if !ok {
		t.Fatalf("monster.Store() error 期望为%v，实际为%v", true, ok)
	}
	t.Logf("monster.Store() 测试成功")
}
func TestRestore(t *testing.T) {
	//创建空实例
	var monster Monster
	ok := monster.Restore()
	if !ok {
		t.Fatalf("monster.Restore() 错误 希望为%v,实际为%v", true, ok)
	}
	//进一步判断
	if monster.Name != "JT" {
		t.Fatalf("monster.Restore() 错误 希望为%v,实际为%v","JT", monster.Name)
	}
	t.Logf("monster.ReStore() 测试成功")
}
