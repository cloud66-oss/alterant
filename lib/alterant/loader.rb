require 'mini_racer'
require 'yaml'
require 'json'
require 'diffy'

Dir.glob File.join(__dir__, 'helpers', '**', '*.rb'), &method(:require)
Dir.glob File.join(__dir__, 'classes', '**', '*.rb'), &method(:require)

# Load other Alterant classes.
Dir.glob File.join(__dir__, '**', '*.rb'), &method(:require)

