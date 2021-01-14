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
