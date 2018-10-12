module Alterant
	class Alterant
		attr_reader :basedir

		# input is a hash
		# filename is the modifier filename use for the backtrace
		# modifier is the script in string
		def initialize(input:, modifier:, filename:, options: {})
			@modifier = modifier
			@filename = filename
			@input = input
			@basedir = options[:basedir]
			@js_preload = options[:js_preload] || []
		end

		# timeout is in ms
		# returns a hash
		def execute(timeout: 500, max_memory: 5000000)
			jpath = ::Alterant::Helpers::Jpath.new

			result = []
			snapshot = MiniRacer::Snapshot.new("$$ = #{@input.to_json};\n" + @js_preload.join("\n")) # this is more efficient but we lose debug info (filename) of helper classes

			isolate = MiniRacer::Isolate.new(snapshot)
			@input.each_with_index do |input, idx|
				ctx = ::MiniRacer::Context.new(isolate: isolate, timeout: timeout, max_memory: max_memory)
				ctx.eval("$ = #{input.to_json}")
				ctx.eval("$['fetch'] = function(key) { return jpath.fetch(JSON.stringify($), key); }")
				ctx.attach('jpath.fetch', proc{|x, y| jpath.fetch(x, y)})
				ctx.attach('console.log', proc{|x| STDERR.puts("DEBUG: #{x.inspect}") if $debug })
				ctx.attach('console.exception', proc{|x| raise ::Alterant::RuntimeError, x })
				ctx.attach('$$.push', proc{|x| result << x })
				ctx.attach('$.index', proc{ idx })
				ctx.attach('YamlReader', ::Alterant::Classes::YamlReader.new(self, ctx))
				ctx.attach('JsonReader', ::Alterant::Classes::JsonReader.new(self, ctx))

				ctx.eval(@modifier, filename: @filename)
				pre_convert = ctx.eval("JSON.stringify($)")
				converted = JSON.parse(pre_convert)
				result << converted

				ctx.dispose
				isolate.idle_notification(100)
			rescue ::MiniRacer::RuntimeError => exc
				if $debug
					raise
				else
					raise ::Alterant::ParseError, "part: #{idx} - #{exc.message}, #{exc.backtrace.first}"
				end
			rescue ::Alterant::AlterantError => exc
				STDERR.puts exc.message.red
				return nil
			end

			return result
		end

	end
end
