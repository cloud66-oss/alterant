module Alterant
	module Helpers
		class Jpath
			require 'jsonpath'

			def fetch(context, key)
				return ::JsonPath.on(context, key)
			end

		end
	end
end
