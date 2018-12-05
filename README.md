# twitter-random-irasutoya
Twitter のプロフィールをランダムないらすとやにする(｀･ω･´)

## 概要

このプログラムを実行すると Twitter のプロフィールがランダムに更新されます。

更新内容は，アイコン，アカウント名，bio 欄です。

ついでに更新ツイートもします。

## 使い方

### 準備

`.env` ファイルに Twitter の各種キーを書く。

### 手動更新

`ruby main.rb`

### 毎日 00:00 に更新

`clockwork clockwork.rb`
