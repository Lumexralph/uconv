# uconv

A general purpose unit-conversion CLI tool. uconv is analogous to the the google search  unit conversion tool only that it is used in your favourite terminal and save you the stress of going to the browsers to get some conversion done for you.

It helps you focus by staying in the terminal and maintain your productivity. It takes input with flags and converts the given input into units like temperature in Celsius, Fahrenheit and Kelvin, length in feet and meters, weight in pounds and kilograms, and the likes.

## Getting started

Clone the repository

    git clone https://github.com/Lumexralph/uconv.git


Change into the directory

    cd uconv

Install all the necessary packages, run

    go mod tidy

Build and install the tool in your computer

    go install .

uconv gets added as a binary in your computer and you can produce using it.

## Usage Example

    uncon temperature 100 --from=c --to=k

and get the following output

    temperature: 30°C ==> 303K

The anatomy of the command is this

    [command] [subcommand] [argument] --flags

    uconv temperature 100 --from=c --to=k

## Errors and bugs

If something is not behaving intuitively, it is a bug and should be reported.
Report it here by creating an issue: https://github.com/Lumexralph/uconv/issues

Help me fix the problem as quickly as possible by following [Mozilla's guidelines for reporting bugs.](https://developer.mozilla.org/en-US/docs/Mozilla/QA/Bug_writing_guidelines#General_Outline_of_a_Bug_Report)

## Patches and pull requests

Your patches are welcome. Here's a suggested workflow:

* Fork the project.
* Make your feature addition or bug fix.
* Send a pull request with a description of your work.

## Copyright

Copyright (c) 2020 LumexRalph. Released under the [MIT License](https://github.com/Lumexralph/uconv/blob/master/LICENSE).
