module Alterant
	module Classes
		class YamlReader
			attr_reader :value

			def call(file)
				if @alter.basedir.nil?
					raise ::Alterant::RuntimeError, 'no basedir set'
				end

				content = File.read(File.join(@alter.basedir, file))
				return ::YAML.safe_load(content)
			end

			def initialize(alter, context)
				@context = context
				@alter = alter
			end

		end
	end
end
