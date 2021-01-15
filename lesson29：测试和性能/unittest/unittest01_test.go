package unittest

import (
	"net/http"
	"testing"
)

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
