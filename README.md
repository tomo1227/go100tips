# go100tips

[![CI](https://github.com/tomo1227/go100tips/actions/workflows/ci.yml/badge.svg)](https://github.com/tomo1227/go100tips/actions/workflows/ci.yml) [![codecov](https://codecov.io/github/tomo1227/go100tips/graph/badge.svg?token=O0NWQ58DGF)](https://codecov.io/github/tomo1227/go100tips) [![CodeQL](https://github.com/tomo1227/go100tips/actions/workflows/code_ql.yml/badge.svg)](https://github.com/tomo1227/go100tips/actions/workflows/code_ql.yml) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Go100Tipsの解説で使用するサンプルリポジトリ

## Slide

* [Go言語100Tips No.1~5 まとめ](https://speakerdeck.com/tomo1227/goyan-yu-100tips-no-dot-1-5-matome)
* [Go言語100Tips No.6~10 まとめ](https://speakerdeck.com/tomo1227/go100tips6-10)
* [ドラ🔵もんでDIを学ぶ](https://speakerdeck.com/tomo1227/dora-mondediwoxue-bu)
* [依存性の注入とは？](https://speakerdeck.com/tomo1227/yi-cun-xing-nozhu-ru-toha)

## Usage

### APIサーバーの起動

以下のコマンドでサンプルのサーバーを起動できる。

```sh
make run
```

```shell_session
$ curl localhost:8080/hello
{"data":"Hello GoFr!"}
$ curl -s localhost:2121/metrics | head -n 5
# HELP app_go_numGC Number of completed Garbage Collector cycles.
# TYPE app_go_numGC gauge
app_go_numGC{otel_scope_name="gofr-app",otel_scope_version="dev"} 5
# HELP app_go_routines Number of Go routines running.
# TYPE app_go_routines gauge
```

### 環境変数の変更

* [.env](.env)で環境変数を自分の好みで設定をしましょう。(portの割り当てに関しては[以下の説明](#portの割り当て)を参照)

### PORTの割り当て

[.env](.env)でホストIPを変更すれば、他プロジェクトとポートが被っても使用できる。

> [!IMPORTANT]
> Macの場合は、.envのHOST_IPに127.0.0.1以外のホストIP(ループバックアドレス)を指定するとき
> 以下のコマンドをターミナルで事前に叩いておく必要がある。(127.0.0.2の箇所にHOST IPを指定)
>
> ```txt
> sudo ifconfig lo0 alias 127.0.0.2
> ```

## merge

main にマージするときは`squash and merge`すること

## commit template の設定

```bash
git config --global commit.template .commit.template
```

## Commit Message ガイドライン

.commit template を参考に記述してください

## Pull Request ガイドライン

`[emoji][type]: (title)`

* ✨feat: ログイン機能を追加
* 👓fix(a11y): ナビゲーションのアクセシビリティを改善

### Pull Request Type

* fix: 🐛 バグの修正 (SemVer パッチと関連)
* feat: ✨ 新機能を追加 (SemVer のマイナーに対応)
* feat!: 💥 破壊的な新機能 (SemVer のメジャーになります)
* fix!: 💥 破壊的なバグ修正 (SemVer のメジャーになります)
* refactor: ♻️ コードの再構築
* revert: ⏪ 変更を取り消す
* test: 🧪 テストに関連する変更
* docs: 📚 ドキュメンテーションの変更
* style: 🎨 スタイルや書式の変更
* perf: ⚡ パフォーマンス改善
* build: 👷‍♀️ ビルドシステムや外部依存関係の変更
* chore: 🔧 その他の変更
* ci: 🎡 CI/CD パイプラインに関連する変更
