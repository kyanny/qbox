package main

import (
	"fmt"
	"encoding/base64"
)

func main() {
	str := "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	str = "lorem ipsum（ロレム・イプサム、略してリプサム lipsum ともいう）とは、出版、ウェブデザイン、グラフィックデザインなどの諸分野において使用されている典型的なダミーテキスト。書籍やウェブページや広告などのデザインのプロトタイプを制作したり顧客にプレゼンテーションしたりする際に、まだ正式な文章の出来上がっていないテキスト部分の書体（フォント）、タイポグラフィ、レイアウトなどといった視覚的なデザインを調整したりわかりやすく見せるために用いられる。"
	fmt.Println(base64.URLEncoding.EncodeToString([]byte(str)))
}
