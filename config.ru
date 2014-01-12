require "rack"

module Rack
  class Lint
    def call(env=nil)
      @app.call(env)
    end
  end
end

app = proc { |env|
  input = env['rack.input']
  out = open('zero.dat', 'wb')
  while buf = input.read(1024*16)
    out.write(buf)
  end
  [200, { 'Content-Type' => 'text/plain' }, ['ok']]
}

run app
