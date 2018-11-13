#! /usr/bin/env nix-shell
#! nix-shell -i bash -p bash hyperfine
set -euf

subcmd='sh'

hyperfine \
    --warmup 5 \
    "./subcommander $subcmd" \
    "./subcommands/$subcmd"
