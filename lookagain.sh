#!/bin/bash

find . -type f -name '*sh'| rev | cut -d '/' -f 1 | rev | sort -r | sed 's/\.sh//g' | sed 's/\.\///g'