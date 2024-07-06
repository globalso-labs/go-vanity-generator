# Go Vanity Import Path Generator

This project simplifies the creation of vanity import paths for Go packages 
using a YAML configuration file. It allows developers to define custom import
paths for their Go packages, making them more accessible and easier to use in
Go projects.

## Features

- **Easy Configuration**: Define your vanity import paths in a simple YAML file.
- **Automatic Updates**: Automatically updates your Go project with the specified
vanity import paths.
- **Customizable**: Flexible configuration options to meet your project's needs.

## Getting Started

### Prerequisites

- Go 1.15 or later
- Git

### Installation

Download the latest release from the 
[releases page](https://github.com/globalso-labs/go-vanity-generator/releases) 
and extract the binary to your desired location.

### Configuration
Create a go-vanity.yml file in the root directory of your project with the
following structure:

```yaml 
# yaml-language-server: $schema=https://schemas.globalso.dev/go-vanity/config.schema.json
domain: "go.globalso.dev"
author: "Global Solutions L.A."
packages:
  - name: x/tools/vanity
    provider: github
    url: "https://github.com/globalso-labs/go-vanity-generator"
    branch: main
  - name: x/tools/logs
    provider: github
    url: "https://github.com/globalso-labs/go-logger"
    branch: main

```
