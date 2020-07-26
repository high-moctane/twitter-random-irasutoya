require "net/http"
require "uri"
require "json"
require "open-uri"
require "twitter"
require "pp"
require "tmpdir"

# いらすとやのURLを返す
def irasutoya_url(idx)
  "http://www.irasutoya.com/feeds/posts/summary?start-index=#{idx}&max-results=1&alt=json-in-script"
end

# いらすとやのJSONをとってくる
def fetch_irasutoya_json(idx)
  resp = Net::HTTP.get(URI.parse(irasutoya_url(idx)))
  raw_json = /{.*}/.match(resp).to_s
  JSON.parse(raw_json)
end

# いらすとやの最大インデックスをとってくる
def fetch_irasutoya_max_idx
  # 現状でインデックスは20000以上はあるので，1から20000で取ってくる
  json = fetch_irasutoya_json(rand(1..20000))
  json["feed"]["openSearch$totalResults"]["$t"]
end

# いらすとやのランダムなJSONをとってくる
def fetch_random_irasutoya_json
  rand_idx = fetch_irasutoya_max_idx.to_i
  fetch_irasutoya_json(rand(1..rand_idx))
end

# JSON からタイトルをひろう
def json_title(json)
  json["feed"]["entry"][0]["title"]["$t"]
end

# JSON から要約をひろう
def json_summary(json)
  json["feed"]["entry"][0]["summary"]["$t"]
end

# JSON からURLをひろう
def json_url(json)
  json["feed"]["entry"][0]["link"][-1]["href"]
end

# JSON からサムネイルのURLをとってくる
def json_thumbnail_url(json)
  json["feed"]["entry"][0]["media$thumbnail"]["url"]
end


# ---------------------------------------------------------------------------
# main
# ---------------------------------------------------------------------------

# ランダムいらすとやをとってくる
json = fetch_random_irasutoya_json
title = json_title(json)
summary = json_summary(json)
url = json_url(json)
thumbnail_url = json_thumbnail_url(json)

# twitter クライアントの作成
twitter = Twitter::REST::Client.new do |config|
  config.consumer_key = ENV["CONSUMER_KEY"]
  config.consumer_secret = ENV["CONSUMER_SECRET"]
  config.access_token = ENV["ACCESS_TOKEN"]
  config.access_token_secret = ENV["ACCESS_TOKEN_SECRET"]
end

# 画像をフェッチする都合上一時フォルダの中で操作します
Dir.mktmpdir do |dir|
  thumb_name = dir + "/thumbnail.png"
  URI.open(thumbnail_url, "rb") { |f| open(thumb_name, "wb") { |png| png.write(f.read) } }

  # プロフィールのアップデート
  twitter.update_profile({name: title[0..20], description: summary})
  open(thumb_name) { |f| twitter.update_profile_image(f) }
  twitter.update(summary[0...(140-url.length)] + url)
end
