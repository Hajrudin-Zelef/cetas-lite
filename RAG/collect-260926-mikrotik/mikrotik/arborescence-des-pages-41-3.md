---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-41-3
title: "arborescence-des-pages-41"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["1970-01-01"]
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-41.md
source_anchor: ""
source_lines: [294, 339]
sha256: fa092c91a7c0d8084e1602757ae9525787f60ff8f428409938108b5feb8fdfbe
---

# arborescence-des-pages-41

| Command | Syntax | Description | Example | 
|---|---|---|---|
| **/** |  | go to the root menu |  | 
| **..** |  | go back by one menu level |  | 
| **?** |  | list all available menu commands and brief descriptions |  | 
| **global** | `:global <var> [<value>]` | define a global variable | `:global myVar "something"; :put $myVar;` | 
| **local** | `:local <var> [<value>]` | define the local variable | `{ :local myLocalVar "I am local"; :put $myVar; }` | 
| **beep** | `:beep frequency=[num] length=[num]` | beep built-in speaker |  | 
| **break** | :break | Breaks a loop. |  | 
| **convert** | `:convert from=[arg] to=[arg]` | Converts specified value from one format to another. By default uses an automatically parsed value, if the "from" format is not specified (for example, "001" becomes "1", "10.1" becomes "10.0.0.1", etc.). **from** specifies the format of the value -*base32, base64, byte-array, hex, num, raw, url* . **to** specifies the format of the output value -*base32, base64, bit-array-lsb, bit-array-msb, byte-array, hex, num, raw, url* . **transform** to transform values -*lc (transforms value to be in lowercases), uc (uppercases), lcfirst (first value to lowercase), ucfirst (first value to uppercase), crlf, ed25519-private-to-x25519-private, none, rot 13, x25519-private-to-x25519-public, ed25519-private-to-ed25519-public, ed25519-public-to-x25519-public, md5, reverse (reverses text), sha512.* | `:put [:convert 001 to=hex ]` `31` `:put [:convert [/ip dhcp-client/option/get hostname raw-value] from=hex to=raw ]` `MikroTik` `:put [convert transform=lc "AAA"]`          `aaa` | 
| **continue** | :continue | Skips the rest of the current loop and continues with the next iteration. |  | 
| **delay** | `:delay <time>` | do nothing for a given period of time |  | 
| **environment** | `:environment print <start>` | print initialized variable information | `:global myVar true; :environment print;`  | 
| **error** | `:error <output>` | Generate console error and stop executing the script |  | 
| **execute** | `:execute <expression>` | Execute the script in the background. The result can be written in the file by setting a "file" parameter or printed to the CLI by setting "as-string". When using the "as-string" parameter executed script is blocked (not executed in the background). Executed script can not be larger than 64kB |  | 
| **exit** | `:exit` | Gracefully stop executing the script |  | 
| **find** | `:find <arg> <arg> <start>` | return position of a substring or array element | `:put [:find "abc" "a" -1];` | 
| **grep** | `:grep script=[str] pattern=[expression] after=[num] before=[num] filename=[str]` | Execute provided `script` in the terminal and print the lines matched by a given `pattern`. Parameters `after` and `before` sets how many lines to print additionaly before and after the matched line. | `:grep script="/interface print" pattern="ether" after=1 before=1 filename=results.txt` | 
| **jobname**  | :jobname | return current script name |  | 
| **len** | `:len <expression>` | return string length or array element count | `:put [:len "length=8"];` | 
| **log** | `:log <topic> <message>` | write a message to the system log. Available topics are `"debug, error, info and warning"` | `:log info "Hello from script";` | 
| **onerror** | `:onerror <var_name> in={<command>} do={<expression>}` | The command used to catch errors and get error details. The **do={...}** block is executed, when**in={...}** block has an error,  and error details are written in <var_name> variable. Parameter order is important. The "error" parameter must be set before "do" block, otherwise do block will not see the local variable. :onerror will return *false* (if there is no error) and*true* (if there is an error) unless otherwise specified (with commands such as :return or :error), so it can be used in**:if** condition statement scripts. |  `:onerror errorName in={ :error "failure" } do={ :put "Critical $errorName" }` | 
| **parse** | `:parse <expression>` | parse the string and return parsed console commands. Can be used as a function. | ``` :global myFunc [:parse ":put hello!"]; $myFunc; ```  | 
| **pick** | `:pick <var> <start> [<end>]` | return range of elements or substring. If the count is not specified, will return only one element from an array.  |  | 
| **put** | `:put <expression>` | put the supplied argument into the console | :put "Hello world" | 
| **range** | `:range <var> <var>` | creates an array from the specified range | :put [:range 2 8] 2;3;4;5;6;7;8 | 
| **resolve** | `:resolve <arg> [<domain-name>][<server>@vrf][<server-port>][<type>]`  | return the IP address of the given DNS name  | `:put [:resolve "www.mikrotik.com"];` `:put [:resolve domain-name="www.mikrotik.com"];` `:put [:resolve domain-name="www.mikrotik.com" server=192.168.88.1 port=53];` `:put [:resolve domain-name="www.mikrotik.com" type=ipv6];` | 
| **retry** | :onerror e {:retry command=<expr> delay=[num] max=[num]} do={<expr>} | Try to execute the given command "max" amount of times with a given "delay" in seconds between tries. On failure, execute the command in the do={} block. |  | 
| **typeof** | `:typeof <var>` | the return data type of variable | `:put [:typeof 4];` | 
| **rndnum** | `:rndnum from=[num] to=[num]` | random number generator | `:put [:rndnum from=1 to=99];` | 
| **rndstr** | `:rndstr from=[str] length=[num]` | Random string generator. **from** specifies characters to construct the string from and defaults to all ASCII letters and numerals.**length** specifies the length of the string to create and defaults to 16. | `:put [:rndstr from="abcdef%^&``" length=33];` | 
| **set** | `:set <var> [<value>]` | assign value to a declared variable. | `:global a; :set a true;` | 
| **serialize** | `:serialize [<value>] to=[arg]` | Serialize specified value/array to JSON or dsv (delimiter separated values) format. **value**  specifues which values to process. **to** specifies the format -*json, dsv* **delimiter** sets the "separator". **order** specifies the order for variables. **options** specifies additional options*:*  **file-name**  enables the option to generate command's output into a file (available for download in the "/files" section). |  | 
| **deserialize** | `:deserialize [<value>] from=[arg]` | Deserialize specified value/array from JSON or dsv (delimiter separated values) format. **from** specifies the format -*json, dsv* **delimiter** sets the "separator". **options** specifies additional options*:*  |  | 
| **time** | `:time <expression>` | return interval of time needed to execute the command | `:put [:time {:for i from=1 to=10 do={ :delay 100ms }}];` | 
| **timestamp** | `:timestamp` | returns the time since epoch, where epoch is January 1, 1970 (Thursday), not counting leap seconds |  | 
| **toarray** | `:toarray <var>` | convert a variable to the array |  | 
| **tobool** | `:tobool <var>` | convert a variable to boolean |  | 
| **toid** | `:toid <var>` | convert a variable to internal ID |  | 
| **toip** | `:toip <var>` | convert a variable to IP address |  | 
| **toip6** | `:toip6 <var>` | convert a variable to IPv6 address |  | 
| **tonum** | `:tonum <var>` | convert a variable to an integer |  | 
| **tostr** | `:tostr <var>` | convert a variable to a string |  | 
| **totime** | `:totime <var>` | convert a variable to time |  | 
| **tonsec** | `:tonsec <var>` | convert time to nanoseconds | :put [:tonsec value=10:00]               36000000000000 | 
| **tocrlf** | `:tocrlf <var>` | converts line endings to CRLF | ``` :put [:tocrlf  "AAA\r\nBBB\r\nCCC" ]                                                                                                              AAA                                               BBB CCC ```  | 
