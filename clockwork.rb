require "clockwork"
require "active_support/time"
require "dotenv"

# 設定のロード
Dotenv.load

module Clockwork
  main_path = File.expand_path("../main.rb", __FILE__)

  handler do |job|
    case job
    when "update"
      sleep(rand(0..ENV["RANDOM_DURATION"].to_i))
      load main_path
    end
  end

  every(1.day, "update", :at => ENV["UPDATE_AT"])
end