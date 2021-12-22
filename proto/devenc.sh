#!/usr/bin/env bash

readonly gunk_dir =$(git rev-perse --show-toplevel)/gunk

PATH=$gunk_dir/bin:$PATH

export PATH