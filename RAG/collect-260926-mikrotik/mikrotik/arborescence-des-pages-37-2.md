---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-37-2
title: "arborescence-des-pages-37"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-37.md
source_anchor: ""
source_lines: [77, 204]
sha256: 0d5b09d4e87bf3ec38066ce718337464be41ae196f0de801fef61171a5b98bf9
---

# arborescence-des-pages-37

All of the Servlet HTML pages use variables to show user specific values. Variable names appear only in the HTML source of the servlet pages - they are automatically replaced with the respective values by the HotSpot Servlet. For most variables, there is an example of their possible value included in brackets. All the described variables are valid in all servlet pages, but some of them just might be empty at the time they are accessed (for example, there is no uptime before a user has logged in).

#### List of available variables

**Note:** Some of the variables use hard-coded http URL, if you are using https, you can construct the link in some other way, for example for $link-status, you can use https://$(hostname)/$(target-dir)status

**Common server variables:**

- **hostname** - DNS name or IP address (if DNS name is not given) of the HotSpot Servlet ("hotspot.example.net")
- **identity** - RouterOS identity name ("MikroTik")
- **login-by** - authentication method used by user
- **plain-passwd** - a "yes/no" representation of whether HTTP-PAP login method is allowed ("no")
- **server-address** - HotSpot server address ("10.5.50.1:80")
- **ssl-login** - a "yes/no" representation of whether HTTPS method was used to access that servlet page ("no")
- **server-name** - HotSpot server name (set in the /ip hotspot menu, as the name property)

**Links:**

- **link-login** - link to login page including original URL requested ("http://10.5.50.1/login?dst=http://www.example.com/")
- **link-login-only** - link to login page, not including original URL requested ("http://10.5.50.1/login")
- **link-logout** - link to logout page ("http://10.5.50.1/logout")
- **link-status** - link to status page ("http://10.5.50.1/status")
- **link-orig** - original URL requested ("http://www.example.com/")

**General client information:**

- **domain** - domain name of the user ("example.com")
- **interface-name** - physical HotSpot interface name (in case of bridged interfaces, this will return the actual bridge port name)
- **ip** - IP address of the client ("10.5.50.2")
- **logged-in** - "yes" if the user is logged in, otherwise - "no" ("yes")
- **mac** - MAC address of the user ("01:23:45:67:89:AB")
- **trial** - a "yes/no" representation of whether the user has access to trial time. If user's trial time has expired, the value is "no"
- **username** - the name of the user ("John")
- **host-ip** - client IP address from /ip hotspot host table
- **vlan-id** - Represents ID of a VLAN interface from which the client is connected

**User status information:**

- **idle-timeout** - idle timeout ("20m" or "" if none)
- **idle-timeout-secs** - idle timeout in seconds ("88" or "0" if there is such timeout)
- **limit-bytes-in** - byte limit for send ("1000000" or "---" if there is no limit)
- **limit-bytes-out** - byte limit for receive ("1000000" or "---" if there is no limit)
- **refresh-timeout** - status page refresh timeout ("1m30s" or "" if none)
- **refresh-timeout-secs** - status page refresh timeout in seconds ("90s" or "0" if none)
- **session-timeout** - session time left for the user ("5h" or "" if none)
- **session-timeout-secs** - session time left for the user, in seconds ("3475" or "0" if there is such timeout)
- **session-time-left** - session time left for the user ("5h" or "" if none)
- **session-time-left-secs** - session time left for the user, in seconds ("3475" or "0" if there is such timeout)
- **uptime** - current session uptime ("10h2m33s")
- **uptime-secs** - current session uptime in seconds ("125")

**Traffic counters, which are available only on the status page:**

- **bytes-in** - number of bytes received from the user ("15423")
- **bytes-in-nice** - user-friendly form of number of bytes received from the user ("15423")
- **bytes-out** - number of bytes sent to the user ("11352")
- **bytes-out-nice** - user-friendly form of number of bytes sent to the user ("11352")
- **packets-in** - number of packets received from the user ("251")
- **packets-out** - number of packets sent to the user ("211")
- **remain-bytes-in** - remaining bytes until limit-bytes-in will be reached ("337465" or "---" if there is no limit)
- **remain-bytes-out** - remaining bytes until limit-bytes-out will be reached ("124455" or "---" if there is no limit)

**Miscellaneous variables:**

- **session-id** - value of 'session-id' parameter in the last request
- **var** - value of 'var' parameter in the last request
- **error** - error message, if something failed ("invalid username or password")
- **error-orig** - original error message (without translations retrieved from errors.txt), if something failed ("invalid username or password")
- **chap-id** - value of chap ID ("\371")
- **chap-challenge** - value of chap challenge ("\357\015\330\013\021\234\145\245\303\253\142\246\133\175\375\316")
- **popup** - whether to pop-up checkbox ("true" or "false")
- **advert-pending** - whether an advertisement is pending to be displayed ("yes" or "no")
- **http-status** - allows the setting of the http status code and message
- **http-header** - allows the setting of the http header

**RADIUS-related variables:**

- **radius<id>** - show the attribute identified with <id> in text string form (in case RADIUS authentication was used; "" otherwise)
- **radius<id>u** - show the attribute identified with <id> in unsigned integer form (in case RADIUS authentication was used; "0" otherwise)
- **radius<id>-<vnd-id>** - show the attribute identified with <id> and vendor ID <vnd-id> in text string form (in case RADIUS authentication was used; "" otherwise)
- **radius<id>-<vnd-id>u** - show the attribute identified with <id> and vendor ID <vnd-id> in unsigned integer form (in case RADIUS authentication was used; "0" otherwise)

#### Working with variables

$(if <var_name>) statements can be used in these pages. The following content will be included, if value of <var_name> will not be an empty string. It is an equivalent to $(if <var_name> != "") It is possible to compare on equivalence as well: $(if <var_name> == <value>) These statements have effect until $(elif <var_name>), $(else) or $(endif). In general case it looks like this:

some content, which will always be displayed
$(if username == john)
Hey, your username is john
$(elif username == dizzy)
Hello, Dizzy! How are you? Your administrator.
$(elif ip == 10.1.2.3)
You are sitting at that old computer, which is so slow...
$(elif mac == 00:01:02:03:04:05)
This is an ethernet card, which was stolen few months ago...
$(else)
I don't know who you are, so lets live in peace.
$(endif)
other content, which will always be displayed

Only one of those expressions will be shown. Which one - depends on the values of those variables for each client.

#### Redirects and custom Headers

	$(if http-status == 302)Hotspot login required$(endif)
	$(if http-header == "Location")$(link-redirect)$(endif)

**Note:** Although the above appears to use the conditional expression 'if' it is in fact setting the 'http-status' to '302' not testing for it. Also the same for the variable 'http-header'. Once again, even though it uses an 'if' it is in fact setting the variable to 'Location' followed by the URL set from the variable 'link-redirect'.

For example, in the case where $(link-redirect) evaluates to "http://192.168.88.1/login", then the HTTP response returned to the client will be changed to:

HTTP/1.0 302 Hotspot login required
<regular HTTP headers>
Location: http://192.168.88.1/login

**http-status syntax**:

	$(if http-status == XYZ)HTTP_STATUS_MESSAGE$(endif)

- *XYZ* - The status code you wish to return. Should be 3 decimal digits, the first one must not be 0
- *HTTP_STATUS_MESSAGE* - any text you wish to return to the client that will follow the above status code in the HTTP reply

In any HTTP response it will be on the first line and will be as follows:

HTTP/1.0 XYZ HTTP_STATUS_MESSAGE

**http-header syntax:**

