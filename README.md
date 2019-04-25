# Alterant

Alterant is a configuration transformer. It reads and converts configuration files in YAML using Javascript you write. The purpose of Alterant is to provide a transparent and auditable way to apply changes to configuration files, like Kubernetes manifest files. To read more about Alterant, please visit [Alterant Website](https://help.cloud66.com/alterant/).

## IMPORTANT NOTICE

The original Alterant code was written in Ruby and is now retired. Alterant v2 is written in Go and is the one that will be supported going forward. Alterant v1 scripts can be used with Alterant v2 in most cases. However, since Alterant v2 doesn't support ES6 Javascript format, you might have to change your scripts if you are using ES6 features in them.

### Why the change?

Alterant v1 used v8 Javascript engine which required native C bindings in Ruby. This made installation of Alterant difficult for users with little or no experience in Ruby and its gem system. The new Alterant supports the same functionality but doesn't have any external dependencies.

## Install

To install Alterant, download it from here: https://github.com/cloud66-oss/alterant/releases/latest

Once you install Alterant, you can use `alterant update` to update to the latest version.

## Usage

For more information on usage, please visit the [Alterant Website](https://help.cloud66.com/alterant/).
