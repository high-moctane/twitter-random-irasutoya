require "clockwork"
require "active_support/all"

module Clockwork
  main_path = File.expand_path("../main.rb", __FILE__)

  handler do |job|
    case job
    when "update"
      load main_path
    end
  end

  every(1.day, "update", :at => "00:00")
end