module Alterant
	module Classes
		CLASSES = {}

		def self.LoadClasses
			# ::Alterant::Classes.constants.map(&::Alterant::Classes.method(:const_get)).grep(Class) do |c|
			# 	name = c.name.split('::').last
			# 	::Alterant::Classes::CLASSES[name] = c
			# end

			# load all JS files in the classes dir and construct a long text
			js_preload = []
			Dir["#{__dir__}/*.js"].each do |f|
				js_preload << File.read(f)
			end

			return js_preload
		end
	end
end
