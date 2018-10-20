Feature: Script Dispatching

    @announce-output
    Scenario: Can get root help
        When I run `subcommander -h`

        Then it should pass with "Usage:"

    @announce-output
    Scenario Outline: Can execute existing scripts

        When I run `subcommander <SUBCOMMAND>`

        Then it should pass with "Hello World!"

        Examples:
            | SUBCOMMAND    |
            | posix         |
            | nested python |
