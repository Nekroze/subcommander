# Subcommander

This single file script works like a dispatcher, distributing commands and parameters to scripts.

The primary use case for `subcommander` is to simplify developing applications with many (perhaps nested like a tree) subcommands. When using `subcommander` you need only write your subcommand scripts following some simple rules and it will generate usage help and version information.

## Developing with Subcommander

First thing you should do is set some environment variables.

- `APPLICATION` the executable/command that the user is running, displayed in usage.
- `SUBCOMMANDS` the path (absolute or relative) to your subcommands, defaults to `./subcommands`.
- `DESCRIPTION` should be set to the long description of your application displayed in usage output on the root command.
- `VERSION` the current version of your application.
- `HOOK` the path to a shell script to be sourced (or noop for empty/undefined string) before dispatching.

Once set we pass any parameters the user has given our application to subcommand for dispathching to our subcommands.

```bash
#!/usr/bin/env bash
set -eu
subcommander "$@"
```

Subcommander will automatically generate `--help` `-h` `help` `-v` `--version` `version` subcommands and switches for all inputs, letting you get to work on the real features as soon as possible.

### Developing a subcommand

In order to do something useful you need to define a subcommand. In your `SUBCOMMANDS` directory add a new executable file that defines a description and usage comment.

```bash
$ echo '#!/bin/sh
# Description: echoes out all given params
# Usage: [<PARAM>...]
echo given $*' > subcommands/params
$ chmod +x subcommands/params
```

Note that subcommands/param uses a shebang, as such you can use anything that supports a shebang like python.

Now we can set some environment variables for our test application and execute our new application.

```bash
$ export APPLICATION=echoer
$ export DESCRIPTION='echos various things and stuff'
$ export VERSION='1.0.0'
$ subcommander params --foo
given --foo
```

You can even use nested subcommands where a directory is the namespace.

```bash
$ mkdir subcommands/env
$ echo '#!/bin/sh
# Description: echoes out all environment variables
printenv' > subcommands/env/all
$ chmod +x subcommands/env/all
```
