#! /bin/sh
#
# m1-starter.sh
# Copyright (C) 2017 stcarolas <stcarolas@homeGround>
#
# Distributed under terms of the MIT license.
#

export command=$(m1 | fzf)
$(m1 $command)
