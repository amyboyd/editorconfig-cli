.editorconfig CLI
=================

This CLI tool is built to automate validating and applying [.editorconfig](http://editorconfig.org/)
rules against files and directories.

It is built in Go, so installation is as simple as downloading one file and running it. It is also
super fast.

The features documented below all work, but some additonal features remain to be done
(see `docs/to-do.md`).

This project's homepage is on GitHub at [github.com/amyboyd/editorconfig-cli](https://github.com/amyboyd/editorconfig-cli)

To downloads the pre-built executables, head to the [releases page](https://github.com/amyboyd/editorconfig-cli/releases/tag/0.1.0)
and look under the "Downloads" title.

Features
--------

* Command `editorconfig-cli check [PATH]` - check if the files within `[PATH]` satisfy the rules
defined in `.editorconfig`. You can use this in a continuous integration process, like Jenkins,
to fail pull requests that don't satify the rules. Or you could use this in a Git pre-commit hook.

* Command `editorconfig-cli fix [PATH]` - fix the files within `[PATH]` to satify the rules.

* Command `editorconfig-cli ls [PATH]` - list the files found within `[PATH]` and the .editorconfig
files that would be applied to them.

* Command `editorconfig-cli rules [PATH]` - list the rules that would be applied to `[PATH]`.

* There is a single binary you can download that works on Mac, Linx and Windows. No installation
needed!

* It's open source.

* It's super fast. The `check` command finishes small codebases in well under 1 second, and a 250k
line codebase is checked in under 3 seconds.

Use in a Git pre-commit hook
----------------------------

```
editorconfig-cli check src/ tests/
if [[ $? != '0' ]]; then
    echo 'Code is not aligned with .editorconfig'
    echo 'Review the output and commit your fixes'
    exit 1
fi
```

How to contribute
-----------------

There is some work still to be done. Refer to the file `docs/to-do.md` for a list.

To run the tests, execute `bin/test`.

License
-------

MIT License

Copyright (c) 2016 Amy Boyd

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
