package collect

import (
	"bufio"
	"fmt"
	"github.com/1152545264/goSpyder/util"
	"golang.org/x/text/transform"
	"io"
	"net/http"
)

type BrowserFetcher struct {
}

// Get 模拟浏览器访问
func (BrowserFetcher) Get(url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		//fmt.Println(err)
		return nil, fmt.Errorf("get url failed:%v", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Mobile Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := util.DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return io.ReadAll(utf8Reader)
}
