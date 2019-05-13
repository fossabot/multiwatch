# multiwatch

[![Build Status](https://travis-ci.org/Enapiuz/multiwatch.svg?branch=master)](https://travis-ci.org/Enapiuz/multiwatch)
[![Go Report Card](https://goreportcard.com/badge/github.com/Enapiuz/multiwatch)](https://goreportcard.com/report/github.com/Enapiuz/multiwatch)
[![codecov](https://codecov.io/gh/Enapiuz/multiwatch/branch/master/graph/badge.svg)](https://codecov.io/gh/Enapiuz/multiwatch)
[![Maintainability](https://api.codeclimate.com/v1/badges/61bf67df2cdf15e5262f/maintainability)](https://codeclimate.com/github/Enapiuz/multiwatch/maintainability)
[![Open Source Helpers](https://www.codetriage.com/enapiuz/multiwatch/badges/users.svg)](https://www.codetriage.com/enapiuz/multiwatch)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://github.com/Enapiuz/multiwatch/blob/master/LICENSE)

Simple task runner on directory changes.
[![asciicast](https://asciinema.org/a/245987.svg)](https://asciinema.org/a/245987)

## Installation
### Manual
1. Download multiwatch
    - `git clone https://github.com/Enapiuz/multiwatch.git`
2. Install via go
    - `cd multiwatch && go install`
    
### Distros
#### macOS
`brew install Enapiuz/tap/multiwatch`
#### Other systems
Work in progress

## Config
```toml
# debounce time for file change events in milliseconds
delay=500

[[watch]]
name = "linter"
paths = ["src"]
commands = ["npm run lint"]

[[watch]]
name = "tests"
paths = ["src", "tests"]
ignorePrefixes=["vendor"] # ignore "src/vendor/*" and "tests/vendor/*" 
commands = ["npm run test"]
```
