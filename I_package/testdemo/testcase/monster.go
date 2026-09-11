package monster
import(
	"fmt"
	"encoding/json"
	"os"
)

type Monster struct {
	Name string
	Age int 
	Skill string
}
//给Moster绑定方法Store，可以将一个Monster变量（对象），序列化后保存到文件中
func (this *Monster) Store() bool{
	data,err := json.Marshal(this)
	if err != nil {
		fmt.Println("序列化失败")
		return false
	}
	filePath := "e:/monster.txt"
	err = os.WriteFile(filePath,data,0666)
	if err != nil {
		fmt.Println("write file err=",err)
		return false
	}
	return true
}
//给Moster绑定方法Restore，可以将一个序列化的Moster，从文件中读取，并发序列化为Moster对象，检查反序列化
func (this *Monster) Restore() bool {
	//1.先从文件中，读取序列化的字符串
	filePath := "e:/monster.txt"
	data,err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFilr err = ",err)
		return false
	}
	//2.使用读取到data []byte ，对反序列化
	err = json.Unmarshal(data,this)
	if err != nil {
		fmt.Println("Unmarshal err = ",err)
		return false
	}
	return true
}
