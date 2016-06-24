This is how we convert a file pattern string found in a .editorconfig to a Go-compatible regex. This
is done mostly according to the rules documented under "Wildcard Patterns" here:
http://docs.editorconfig.org/en/master/editorconfig-format.html#patterns

The are some differences to the official documentation, however.

`*` Matches any string of characters, except path separators (/).

`**` Matches any string of characters.

`?` Matches any single character.

`[seq]` Matches any single character in seq.

`[!seq]` Matches any single character not in seq.

`{s1,s2,s3}` Matches any of the strings given (separated by commas, can be nested).

`{num1..num2}` Matches any integer numbers between num1 and num2, where num1 and num2 can be either positive or negative.

**Differences from the official document**

If file pattern is entirely `*`, according to the official documentation it should match files in the
same directory and not in sub-directories (because `*` excludes path separators). However, in open
source repositories I reviewed, a single `*` seems to be universally used to mean every file,
instead of the technically-correct `**`. We adapt to what is used in the real world to be practical.
