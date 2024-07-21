package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 日付の入力を促す
	fmt.Print("イベントの日付を入力してください (形式、 yyyymmdd): ")
	date, _ := reader.ReadString('\n')
	date = strings.TrimSpace(date)

	// タイトルの入力を促す
	fmt.Print("イベントのタイトルを入力してください: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	// イベントURLの入力を促す
	fmt.Print("イベントのURLを入力してください: ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)

	// 画像ファイルパスの入力を促す
	fmt.Print("使用するバナーのファイル名を拡張子付きで入力してください（例、 schedule-yyyymmdd-bg.png）: ")
	imagePath, _ := reader.ReadString('\n')
	imagePath = strings.TrimSpace(imagePath)
	imageFileName := filepath.Base(imagePath)

	// HTMLテンプレートの作成
	template := generateTemplate(url, date, title, imageFileName)

	// 出力ファイルパスの設定
	outputPath := filepath.Join(".", "output.txt")

	// HTMLテンプレートをファイルに書き込む
	err := os.WriteFile(outputPath, []byte(template), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Printf("テンプレートが生成されました: %s\n", outputPath)
}

func generateTemplate(url, date, title, imageFileName string) string {
	// 年、月、日の部分を抽出
	year := date[:4]
	month := date[4:6]
	day := date[6:8]

	// "yyyy/mm/dd"形式に変換
	formattedDate := year + "/" + month + "/" + day

	// 月と日を数字で取得
	monthInt, err := strconv.Atoi(month)
	if err != nil {
		fmt.Println("Error converting month:", err)
		return ""
	}
	dayInt, err := strconv.Atoi(day)
	if err != nil {
		fmt.Println("Error converting day:", err)
		return ""
	}

	return fmt.Sprintf(`
/* events.html */
<a class="schedule-cont" href="%s" target="_blank" data-category="program new">
    <div class="schedule-img-cont schedule-%s-img">
        <div class="schedule-dates-cont js-event" data-day="%s">
            <p class="manth-cont">%d月</p>
            <p class="days-cont">%d</p>
        </div>
        <p class="end-cont js-end">終了</p>
    </div>
    <p class="schedule-comment">%s</p>
</a>

/* style.css */
#event .event-schedule .schedule-cont .schedule-%s-img {
  background: url(../img/event_schedule_img/%s)
    center center no-repeat;
  background-size: cover;
}`, url, date, formattedDate, monthInt, dayInt, title, date, imageFileName)
}
