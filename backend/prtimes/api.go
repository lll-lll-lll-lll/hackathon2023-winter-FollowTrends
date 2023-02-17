// PRTimes API周りの処理全般を担当
package prtimes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// カテゴリー名・ID対応表
var Category = map[string]string{"パソコン・周辺機器": "1", "パソコンソフトウェア": "2", "プロバイダ・回線接続": "3", "ネットサービス": "4", "スマートフォンアプリ": "5",
	"サーバ・周辺機器": "6", "ネットワーク・ネットワーク機器": "7", "アプリケーション・セキュリティ": "8", "システム・Webサイト・アプリ開発": "9", "百貨店・スーパー・コンビニ・ストア": "10",
	"EC・通販": "11", "卸売・問屋": "12", "芸能": "13", "映画・演劇・DVD": "14", "音楽": "15", "テレビ・CM": "16", "スポーツ": "17", "アウトドア・登山": "18", "雑誌・本・出版物": "19",
	"漫画・アニメ": "20", "アート・カルチャー": "21", "コンシューマーゲーム": "22", "スマートフォンゲーム": "23", "アーケードゲーム": "24", "おもちゃ・遊具・人形": "25", "モバイル端末": "26", "カメラ": "27",
	"AV機器": "28", "調理・生活家電": "29", "健康・美容家電": "30", "インテリア・家具・収納": "31", "日用品・生活雑貨": "32", "ガーデン・DIY": "33", "自動車・カー用品": "34", "バイク・バイク用品": "35",
	"レディースファッション": "36", "メンズファッション": "37", "シューズ・バッグ": "38", "ジュエリー・アクセサリー": "39", "キッズ・ベビー・マタニティ": "40", "食品・お菓子": "41", "ソフトドリンク・アルコール飲料": "42",
	"レストラン・ファストフード・居酒屋": "43", "中食・宅配": "44", "スキンケア・化粧品・ヘア用品": "45", "ダイエット・健康食品・サプリメント": "46", "医療・病院": "47", "医薬・製薬": "48", "福祉・介護・リハビリ": "49",
	"経営・コンサルティング": "50", "シンクタンク": "51", "財務・経理": "52", "法務・特許・知的財産": "53", "銀行・信用金庫・信用組合": "54", "クレジットカード・ローン": "55", "証券・FX・投資信託": "56",
	"生命保険・損害保険": "57", "広告・宣伝・PR": "58", "マーケティング・リサーチ": "59", "セールス・営業": "60", "就職・転職・人材派遣・アルバイト": "61", "資格・留学・語学": "62", "学校・大学": "63",
	"学習塾・予備校・通信教育": "64", "保育・幼児教育": "65", "ホテル・旅館": "66", "旅行・観光": "67", "テーマパーク・遊園地": "68", "住宅・マンション": "69", "商業施設・オフィスビル": "70", "建築・空間デザイン": "71",
	"建設・土木": "72", "鉄鋼・金属・ガラス・土石・ゴム": "73", "化学": "74", "電気・ガス・資源・エネルギー": "75", "交通・運送・引越し": "76", "物流・倉庫・貨物": "77", "自然・天気": "78",
	"環境・エコ・リサイクル": "79", "ペット・ペット用品": "80", "ギフト・花": "81", "恋愛・結婚": "82", "出産・育児": "83", "葬儀": "84", "政治・官公庁・地方自治体": "85", "財団法人・社団法人": "86", "ボランティア": "87",
	"国際情報・国際サービス": "88", "農林・水産": "89", "その他": "90", "フィットネス・ヘルスケア": "91", "電子部品・半導体・電気機器": "92"}

// 受け取ったカテゴリー名に対応するカテゴリーIDを格納する変数宣言
var categoryID string

const NotFoundArticle = "カテゴリに属している記事が存在しませんでした"

// PRTimes APIを叩いたり、周りの処理を担当する
type PRTimes struct {
	Response Items `json:"response"`
}

func New() *PRTimes {
	return &PRTimes{}
}

type Item struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// String LINE上に表示するフォーマット
func (i Item) String() string {
	return fmt.Sprintf("%s:%s", i.Title, i.URL)
}

type Items []Item

// String スライスのItem構造体を文字列にするメソッド
func (it Items) String() string {
	var s string
	if len(it) == 0 {
		return NotFoundArticle
	}

	for i, v := range it {
		if i == 3 {
			return s
		}
		s += v.String()
	}
	return s
}

// 現在時刻を取得し、3日前の日付を返す。
// この時のフォーマットはYYYY-MM-DD
func GetFromDate() string {
	now := time.Now()
	fromDate := now.AddDate(0, 0, -4)
	return fromDate.Format("2006-01-02")
}

// GetItems PRTimes APIを叩いてItem構造体を返す
func (pt *PRTimes) GetItems(categoryName string) (Items, error) {

	//カテゴリー対応表と受け取ったカテゴリー名を比較、同じならカテゴリーIDを設定
	for i, j := range Category {
		if categoryName == i {
			categoryID = j
		}
	}

	client := &http.Client{}
	// 3日前の日付を取得
	fromDate := GetFromDate()
	// PRTimes APIを叩く
	// カテゴリーIDと3日前の日付を指定
	url := fmt.Sprintf("https://hackathon.stg-prtimes.net/api/categories/%s/releases?from_date=%s", categoryID, fromDate)

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("Authorization", "Bearer b655dffbe1b2c82ca882874670cb110995c6604151e1b781cf5c362563eb4e12")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var items Items
	if err := json.Unmarshal(body, &items); err != nil {
		return nil, err
	}
	pt.Response = items
	return items, nil
}
