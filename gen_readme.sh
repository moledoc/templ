#!/bin/bash

set -e
echo "# templ

$(go doc -all templ)

## Author

Meelis Utt" > README.md
