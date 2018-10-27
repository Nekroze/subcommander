Feature: Script Dispatching

    Background:
        Given I set the environment variable "SUBCOMMANDS" to "/usr/src/app/subcommands"

    Scenario: Usage contains application name
        Given I set the environment variable "APPLICATION" to "Retro Encabulator"

        When I run `subcommander help`

        Then it should pass with "Retro Encabulator"

    Scenario: Usage contains description
        Given I set the environment variable "DESCRIPTION" to "42"

        When I run `subcommander help`

        Then it should pass with "42"

    Scenario Outline: Can get root help
        When I run `subcommander <ARG>`

        Then it should pass with "Usage:"

        Examples:
            | ARG                  |
            | -h                   |
            | help                 |
            | --help               |
            | posix -h             |
            | posix help           |
            | posix --help         |
            | sh -h                |
            | sh help              |
            | sh --help            |
            | nested python -h     |
            | nested python help   |
            | nested python --help |

    Scenario Outline: can get version info
        Given I set the environment variable "VERSION" to "0.1.1"

        When I run `subcommander <ARG>`

        Then it should pass with "0.1.1"

        Examples:
            | ARG                     |
            | -v                      |
            | version                 |
            | --version               |
            | posix -v                |
            | posix version           |
            | posix --version         |
            | sh -v                   |
            | sh version              |
            | sh --version            |
            | nested python -v        |
            | nested python version   |
            | nested python --version |

    Scenario Outline: Can execute existing scripts

        When I run `subcommander <SUBCOMMAND>`

        Then it should pass with "Hello World!"

        Examples:
            | SUBCOMMAND    |
            | posix         |
            | nested python |

    Scenario: Can see subcommands listed for help on root namespace

        When I run `subcommander help`

        Then it should pass with "Usage:"
        And the output should match /^SUBCOMMAND\s*|\s*ALIASES\s*|\s*DESCRIPTION$/
        And the output should match /^\s*posix\s*|\s*sh\s*|\s*Does things$/

    Scenario: Can see subcommands listed for help on namespaces

        When I run `subcommander help`

        Then it should pass with "Usage:"
        And the output should match /^SUBCOMMAND\s*|\s*ALIASES\s*|\s*DESCRIPTION$/
        And the output should match /^\s*python\s*|\s*|\s*Does things with python$/
