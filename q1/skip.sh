#!/bin/bash

ls -l | awk 'NR != 1' | awk 'NR % 2 == 1'
