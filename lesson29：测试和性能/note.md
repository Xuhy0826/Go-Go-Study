# 测试和性能

## 单元测试

### 基础单元测试
看一个单元测试例子1。
> unittest/unittest01_test.go
```
package unittest

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713" //对号（ √）
const ballotX = "\u2717"   //叉号（× ）

//TestDownload 测试函数需要满足：
//(1) 公开的函数，并且以 Test 单词开头。
//(2) 函数入参是一个指向 testing.T 类型的指针，并且没有返回值。
func TestDownload(t *testing.T) {
	url := "https://raw.githubusercontent.com/Xuhy0826/Go-Go-Study/master/lesson29%EF%BC%9A%E6%B5%8B%E8%AF%95%E5%92%8C%E6%80%A7%E8%83%BD/resource/data/index.html"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)
			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}

```
该测试文件展示了http包的Get函数的单元测试。在当前工作目录下执行`go test -v`来运行测试，`-v`参数表示提供冗余输出。  
一个合法的单元测试需要符合几个条件
1. 测试文件必须是**_test.go**结尾的文件
2. 测试函数必须是公开的函数，并且以 Test 单词开头，函数入参是一个指向 testing.T 类型的指针，并且没有返回值。  
测试代码中的`t.Log`和`t.Logf`即输出测试消息，后者是格式化输出版本。如果执行`go test`的时候没有加入冗余选项`-v`，除非测试失败，否则看不到任何测试输出。`t.Fatal`和`t.Error`函数是汇报测试失败的方法，但凡执行了任一就表示本次测试失败。其中`t.Fatal`和`t.Fatalf`执行后会停止执行，`t.Error`和`t.Errorf`则不会终止执行。

### 表组测试
表则测试的示例代码如下，就是将基础的单元测试进行不同条件多次测试。代码类似上一个示例，无需多言。
```
package unittest

import (
	"net/http"
	"testing"
)

//TestDownload 确认 http 包的 Get 函数可以下载内容
func TestDownload02(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"https://raw.githubusercontent.com/Xuhy0826/Go-Go-Study/master/lesson29%EF%BC%9A%E6%B5%8B%E8%AF%95%E5%92%8C%E6%80%A7%E8%83%BD/resource/data/index.html",
			http.StatusOK,
		},
		{
			"https://raw.githubusercontent.com/Xuhy0826/Go-Go-Study/master/lesson29%EF%BC%9A%E6%B5%8B%E8%AF%95%E5%92%8C%E6%80%A7%E8%83%BD/resource/data/index1.html",
			http.StatusNotFound,
		},
	}

	t.Log("Given the need to test downloading content.")
	{
		for _, u := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"", u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
				}
				t.Log("\t\tShould be able to make the Get call.", checkMark)
				defer resp.Body.Close()

				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould receive a \"%d\" status. %v", u.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould receive a \"%d\" status. %v %v", u.statusCode, ballotX, resp.StatusCode)
				}
			}
		}
	}
}
```