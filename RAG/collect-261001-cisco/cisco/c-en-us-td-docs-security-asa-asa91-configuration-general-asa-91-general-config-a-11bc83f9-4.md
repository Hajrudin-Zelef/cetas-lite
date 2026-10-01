---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9-4
title: "c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2006-01-01"]
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9.md
source_anchor: ""
source_lines: [149, 226]
sha256: d1766b2049a6406373188f50677d79593ab2303e67baf2be6d9b5d014e861036
---

# c-en-us-td-docs-security-asa-asa91-configuration-general-asa-91-general-config-a-11bc83f9

Note As an optimization, the ASA searches on the deobfuscated URL. Deobfuscation compresses multiple forward slashes (/) into a single slash. For strings that commonly use double slashes, like “http://”, be sure to search for “http:/” instead.
Table 17-1 lists the metacharacters that have special meanings.
|  |  |  | 
|---|---|---|
| . | Dot | Matches any single character. For example, d.g matches dog, dag, dtg, and any word that contains those characters, such as doggonnit. | 
| ( exp ) | Subexpression | A subexpression segregates characters from surrounding characters, so that you can use other metacharacters on the subexpression. For example, d(o\|a)g matches dog and dag, but do\|ag matches do and ag. A subexpression can also be used with repeat quantifiers to differentiate the characters meant for repetition. For example, ab(xy){3}z matches abxyxyxyz. | 
| \| | Alternation | Matches either expression it separates. For example, dog\|cat matches dog or cat. | 
| ? | Question mark | A quantifier that indicates that there are 0 or 1 of the previous expression. For example, lo?se matches lse or lose. Note You must enter Ctrl+V and then the question mark or else the help function is invoked. | 
| * | Asterisk | A quantifier that indicates that there are 0, 1 or any number of the previous expression. For example, lo*se matches lse, lose, loose, and so on. | 
| + | Plus | A quantifier that indicates that there is at least 1 of the previous expression. For example, lo+se matches lose and loose, but not lse. | 
| { x } or { x ,} | Minimum repeat quantifier | Repeat at least x times. For example, ab(xy){2,}z matches abxyxyz, abxyxyxyz, and so on. | 
| [ abc ] | Character class | Matches any character in the brackets. For example, [abc] matches a, b, or c. | 
| [^ abc ] | Negated character class | Matches a single character that is not contained within the brackets. For example, [^abc] matches any character other than a, b, or c. [^A-Z] matches any single character that is not an uppercase letter. | 
| [ a - c ] | Character range class | Matches any character in the range. [a-z] matches any lowercase letter. You can mix characters and ranges: [abcq-z] matches a, b, c, q, r, s, t, u, v, w, x, y, z, and so does [ a-cq-z] . The dash (-) character is literal only if it is the last or the first character within the brackets: [abc-] or [-abc] . | 
| “” | Quotation marks | Preserves trailing or leading spaces in the string. For example, “ test” preserves the leading space when it looks for a match. | 
| ^ | Caret | Specifies the beginning of a line. | 
| \ | Escape character | When used with a metacharacter, matches a literal character. For example, \[ matches the left square bracket. | 
| char | Character | When character is not a metacharacter, matches the literal character. | 
| \r | Carriage return | Matches a carriage return 0x0d. | 
| \n | Newline | Matches a new line 0x0a. | 
| \t | Tab | Matches a tab 0x09. | 
| \f | Formfeed | Matches a form feed 0x0c. | 
| \x NN | Escaped hexadecimal number | Matches an ASCII character using hexadecimal (exactly two digits). | 
| \ NNN | Escaped octal number | Matches an ASCII character as octal (exactly three digits). For example, the character 040 represents a space. | 
Detailed Steps
Step 1 To test a regular expression to make sure it matches what you think it will match, enter the following command:
Where the input_text argument is a string you want to match using the regular expression, up to 201 characters in length.
The regular_expression argument can be up to 100 characters in length.
Use Ctrl+V to escape all of the special characters in the CLI. For example, to enter a tab in the input text in the test regex command, you must enter test regex “test[Ctrl+V Tab]” “test\t” .
If the regular expression matches the input text, you see the following message:
If the regular expression does not match the input text, you see the following message:
Step 2 To add a regular expression after you tested it, enter the following command:
Where the name argument can be up to 40 characters in length.
The regular_expression argument can be up to 100 characters in length.
Examples
The following example creates two regular expressions for use in an inspection policy map:
Creating a Regular Expression Class Map
A regular expression class map identifies one or more regular expressions. You can use a regular expression class map to match the content of certain traffic; for example, you can match URL strings inside HTTP packets.
Prerequisites
Create one or more regular expressions according to the “Creating a Regular Expression” section.
Detailed Steps
Step 1 Create a class map by entering the following command:
Where class_map_name is a string up to 40 characters in length. The name “class-default” is reserved. All types of class maps use the same name space, so you cannot reuse a name already used by another type of class map.
The match-any keyword specifies that the traffic matches the class map if it matches at least one of the regular expressions.
The CLI enters class-map configuration mode.
Step 2 (Optional) Add a description to the class map by entering the following command:
Step 3 Identify the regular expressions you want to include by entering the following command for each regular expression:
Examples
The following example creates two regular expressions, and adds them to a regular expression class map. Traffic matches the class map if it includes the string “example.com” or “example2.com.”
Configuring Time Ranges
Create a reusable component that defines starting and ending times that can be applied to various security features. Once you have defined a time range, you can select the time range and apply it to different options that require scheduling.
The time range feature lets you define a time range that you can attach to traffic rules, or an action. For example, you can attach an ACL to a time range to restrict access to the ASA.
A time range consists of a start time, an end time, and optional recurring entries.
Guidelines
- Multiple periodic entries are allowed per time range. If a time range has both absolute and periodic values specified, then the periodic values are evaluated only after the absolute start time is reached, and they are not further evaluated after the absolute end time is reached.
- Creating a time range does not restrict access to the device. This procedure defines the time range only.
Detailed Steps
|  |  |  | 
|---|---|---|
| Step 1 |  | Identifies the time-range name. | 
| Step 2 | Do one of the following: |  | 
|  |  | Specifies a recurring time range. You can specify the following values for days-of-the-week : The time is in the format hh : mm . For example, 8:00 is 8:00 a.m. and 20:00 is 8:00 p.m. | 
|  |  | Specifies an absolute time range. The time is in the format hh : mm . For example, 8:00 is 8:00 a.m. and 20:00 is 8:00 p.m. The date is in the format day month year ; for example, 1 january 2006 . | 
Examples
The following is an example of an absolute time range beginning at 8:00 a.m. on January 1, 2006. Because no end time and date are specified, the time range is in effect indefinitely.
The following is an example of a weekly periodic time range from 8:00 a.m. to 6:00 p.m on weekdays:
Monitoring Objects
To monitor objects and groups, enter the following commands:
|  |  | 
|---|---|
|  | Displays the access list entries that are expanded out into individual entries without their object groupings. | 
|  | Displays all current object groups. | 
|  | Displays the current object groups by their group ID. | 
|  | Displays the current object groups by their group type. | 
Feature History for Objects
Table 17-2 lists each feature change and the platform release in which it was implemented.
|  |  |  | 
|---|---|---|
