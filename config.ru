require "rack"
require "logger"

module Rack
  class Lint
    def call(env=nil)
      @app.call(env)
    end
  end

  class CommonLogger
    def call(env)
      @app.call(env)
    end
  end
end

app = proc { |env|
  now = Time.now
  warn "> #{now}.#{now.usec}"
  _ = env["rack.input"].read
  now = Time.now
  warn "  #{now}.#{now.usec}"
  [200, { 'Content-Type' => 'text/plain' }, ['ok']]
}

run app
