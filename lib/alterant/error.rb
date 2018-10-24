module Alterant
	class AlterantError < StandardError; end
	class ParseError < AlterantError; end
	class RuntimeError < AlterantError; end
  end
