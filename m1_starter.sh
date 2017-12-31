#! /bin/sh
#
# m1_starter.sh
# Copyright (C) 2017 stcarolas <stcarolas@carbon>
#
# Distributed under terms of the MIT license.
#


m1_actions | m1
export command=$(cat /tmp/m1)
$(m1 $command)
