package config

import "os"

// Env は環境変数を取得するためのstructです
// 環境変数がどこで利用されているか把握するため、
// 環境変数を取得する際は必ずこの struct を利用してください
type Env struct{}

// NewEnv は Env のコンストラクタです
func NewEnv() *Env {
	return &Env{}
}

// GetAppSecret は Cookie のエンコード等に利用する鍵
func (e Env) GetAppSecret() string {
	return os.Getenv("APP_SECRET")
}

// GetServerPort はサーバーを起動する際のポート番号
func (e Env) GetServerPort() string {
	return os.Getenv("PORT")
}

// GetDBUser はデータベースに接続する時のユーザー
func (e Env) GetDBUser() string {
	return os.Getenv("DB_USER")
}

// GetDBPassword はデータベースに接続する時のパスワード
func (e Env) GetDBPassword() string {
	return os.Getenv("DB_PASSWORD")
}

// GetDBName は接続するデータベース
func (e Env) GetDBName() string {
	return os.Getenv("DB_NAME")
}

// GetDBHost は接続するデータベースサーバーのホスト
func (e Env) GetDBHost() string {
	return os.Getenv("DB_HOST")
}
