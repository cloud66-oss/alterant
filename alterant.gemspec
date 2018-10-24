
# coding: utf-8
lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)

require 'alterant/version'

Gem::Specification.new do |spec|
  spec.name          = "alterant"
  spec.version       = Alterant::VERSION
  spec.authors       = ["Khash Sajadi"]
  spec.email         = ["khash@cloud66.com"]

  spec.summary       = %q{Alterant gem and command line}
  spec.description   = %q{Alterant is a tool to alter configuration files}
  spec.homepage      = "https://github.com/cloud66/alterant"
  spec.license       = 'Nonstandard'

  spec.files         = Dir.glob("{bin,lib}/**/*") + %w(README.md)
  spec.bindir        = "bin"
  spec.executables   = spec.files.grep(%r{^bin/}) { |f| File.basename(f) }
  spec.require_paths = ["lib"]

  spec.add_development_dependency "bundler", "~> 1.14"
  spec.add_development_dependency "rake", "~> 10.0"
  spec.add_development_dependency "rerun", "~> 0.13"

  spec.add_dependency 'jsonpath', '~>0.9'
  spec.add_dependency 'json', '~> 1.4'
  spec.add_dependency 'thor', '~> 0.20'
  spec.add_dependency 'mini_racer', '~> 0.2'
  spec.add_dependency 'colorize', '~> 0.8'
  spec.add_dependency 'diffy', '~> 3.2'
end
