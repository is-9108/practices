# Go / Docker / CI-CD 実務レベル ハンズオン

他言語の実務経験がある人が、**Go・Docker・GitHub Actions を実務で使えるレベルまで**持っていくための
ハンズオン教材です。最終的に「Go 製 REST API + PostgreSQL + テスト + Docker + CI/CD」を一式作り切ります。

## 進め方

1. `lessons/lessonNN/README.md` を読む（解説 + 課題）
2. `lessons/lessonNN/exercise.go` の `panic("TODO")` を自分の実装に置き換える
3. テストを実行して答え合わせする

```bash
go test ./lessons/lesson01/
```

- **PASS すれば正解**です。`exercise_test.go` が採点表なので、こちらは編集しません（読むのは推奨。テストの書き方のお手本にもなります）
- 詳細を見たいときは `go test -v ./lessons/lesson01/`
- 全レッスンをまとめて回すときは `go test ./...`
- 詰まったら「lessonNN の X が分からない」と聞いてください。ヒント → 解説 → 模範解答の順で出します

## カリキュラム

### Phase 1: Go の文法（実務で使う範囲）

| # | テーマ | 主な内容 |
|---|--------|---------|
| 01 | 基本文法 | 複数戻り値・error・switch・slice・map・rune |
| 02 | slice と map の深掘り | 内部構造、append の罠、参照の共有、nil の扱い |
| 03 | struct とメソッド | 値レシーバ vs ポインタレシーバ、埋め込み、コンストラクタ |
| 04 | interface | 暗黙実装、小さく切る設計、型アサーション、`any` |
| 05 | エラーハンドリング | wrapping、`errors.Is/As`、カスタムエラー、panic を使わない設計 |
| 06 | 並行処理 | goroutine、channel、`select`、`sync`、`context` |
| 07 | 標準ライブラリ | `encoding/json`、`time`、`io`、`strings`、`log/slog` |
| 08 | パッケージ設計 | modules、可視性、ディレクトリ構成、循環参照の回避 |

### Phase 2: テストコード

| # | テーマ | 主な内容 |
|---|--------|---------|
| 09 | table-driven test | サブテスト、`t.Helper`、`t.Cleanup`、失敗メッセージの書き方 |
| 10 | テストダブル | interface による差し替え、`httptest`、ゴールデンファイル |
| 11 | 品質を測る | カバレッジ、`-race`、ベンチマーク、Fuzzing |

### Phase 3: REST API を作る

| # | テーマ | 主な内容 |
|---|--------|---------|
| 12 | net/http | Go 1.22+ の `ServeMux`、ルーティング、ハンドラ、graceful shutdown |
| 13 | レイヤード構成 | handler / service / repository、依存性注入 |
| 14 | ミドルウェア | ロギング、パニックリカバリ、リクエストID、タイムアウト |
| 15 | 入出力設計 | バリデーション、エラーレスポンス、ステータスコード設計 |

### Phase 4: DB 連携（PostgreSQL）

| # | テーマ | 主な内容 |
|---|--------|---------|
| 16 | Docker で DB を立てる | `docker run` / `docker compose`、ボリューム、ネットワーク |
| 17 | database/sql と pgx | クエリ、スキャン、コネクションプール、`context` 連携 |
| 18 | マイグレーション | golang-migrate、up/down、スキーマ管理の実務 |
| 19 | リポジトリ実装 | CRUD、トランザクション、N+1、SQL インジェクション対策 |
| 20 | DB を使ったテスト | テスト用 DB、テストの独立性、フィクスチャ |

### Phase 5: Docker を実務レベルで

| # | テーマ | 主な内容 |
|---|--------|---------|
| 21 | Dockerfile | レイヤーキャッシュ、マルチステージビルド、distroless、非 root |
| 22 | docker compose | API + DB の連携、依存関係、ヘルスチェック、開発用ホットリロード |
| 23 | 運用の観点 | イメージサイズ最適化、環境変数とシークレット、ログ設計 |

### Phase 6: CI/CD（GitHub Actions）

| # | テーマ | 主な内容 |
|---|--------|---------|
| 24 | Actions 基礎 | workflow / job / step、トリガー、マトリクス、権限 |
| 25 | CI パイプライン | `golangci-lint` + test + `-race` + カバレッジ、キャッシュ |
| 26 | 統合テスト | service containers で PostgreSQL を立ててテスト |
| 27 | イメージのビルドと配信 | GHCR への push、タグ戦略、ビルドキャッシュ |
| 28 | CD | リリース自動化、環境分離、secrets、デプロイの型 |

## 環境

- Go 1.25.6
- Docker Desktop（Phase 4 以降で使用）
- Git / GitHub アカウント（Phase 6 で使用）
