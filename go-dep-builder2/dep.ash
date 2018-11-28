#!/bin/ash
. /builder/prepare_workspace.inc
prepare_workspace || exit
/root/go/bin/dep "$@"