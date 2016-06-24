The .editorconfig rules applied are according to the documentation found under "Supported Properties" here: http://docs.editorconfig.org/en/master/editorconfig-format.html#properties

indent_style: set to "tab" or "space" to use hard tabs or soft tabs respectively. The values are case insensitive.

indent_size: a whole number defining the number of columns used for each indentation level and the width of soft tabs (when supported). If this equals to "tab", the indent_size will be set to the tab size, which should be tab_width if tab_width is specified, or the tab size set by editor if tab_width is not specified. The values are case insensitive.

tab_width: a whole number defining the number of columns used to represent a tab character. This defaults to the value of indent_size and should not usually need to be specified.

end_of_line: set to "lf", "cr", or "crlf" to control how line breaks are represented. The values are case insensitive.

charset: set to "latin1", "utf-8", "utf-8-bom", "utf-16be" or "utf-16le" to control the character set. Use of "utf-8-bom" is discouraged.

trim_trailing_whitespace: set to "true" to remove any whitespace characters preceeding newline characters and "false" to ensure it doesn't.

insert_final_newline: set to "true" ensure file ends with a newline when saving and "false" to ensure it doesn't.

root: special property that should be specified at the top of the file outside of any sections. Set to "true" to stop .editorconfig files search on current file. The value is case insensitive.
