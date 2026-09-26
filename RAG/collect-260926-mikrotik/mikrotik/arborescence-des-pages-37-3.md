---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-37-3
title: "arborescence-des-pages-37"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-37.md
source_anchor: ""
source_lines: [205, 363]
sha256: 4acf294642eef7d60e657da0a33f91223e35b0d20a09839614bc3c6935114589
---

# arborescence-des-pages-37

$(if http-header == HTTP_HEADER_NAME)HTTP_HEADER_VALUE$(endif)

- *HTTP_HEADER_NAME* - name of the HTTP header to be sent in the response
- *HTTP_HEADER_VALUE* - the value of the HTTP header with the name HTTP_HEADER_NAME to be sent in the response

The HTTP response will appear as:

HTTP_HEADER_NAME: HTTP_HEADER_VALUE


All variables and conditional expressions within HTTP_HEADER_VALUE and HTTP_STATUS_MESSAGE are processed as usual.

In case multiple headers with the same name are added, then only the last one will be used (previous ones will be discarded). It allows the system to override regular HTTP headers (for example, Content-Type and Cache-Control).

### Customizing Error Messages

All error messages are stored in the errors.txt file within the respective HotSpot servlet directory. You can change and translate all these messages to your native language. To do so, edit the errors.txt file. You can also use variables in the messages. All instructions are given in that file.

### Multiple Versions of HotSpot Pages

Multiple HotSpot page sets for the same HotSpot server are supported. They can be chosen by the user (to select language) or automatically by JavaScript (to select PDA/regular version of HTML pages).

To utilize this feature, create subdirectories in the HotSpot HTML directory, and place those HTML files, which are different, in that subdirectory. For example, to translate everything in Latvian, the subdirectory "lv" can be created with login.html, logout.html, status.html, alogin.html, radvert.html and errors.txt files, which are translated into Latvian. If the requested HTML page can not be found in the requested subdirectory, the corresponding HTML file from the main directory will be used. Then main login.html file would contain a link to "/lv/login?dst=$(link-orig-esc)", which then displays Latvian version of login page: <a href="/lv/login?dst=$(link-orig-esc)">Latviski</a> . And Latvian version would contain a link to English version: <a href="/login?dst=$(link-orig-esc)">English</a>

Another way of referencing directories is to specify 'target' variable:

```
        <a href="$(link-login-only)?dst=$(link-orig-esc)&target=lv">Latviski</a>
        <a href="$(link-login-only)?dst=$(link-orig-esc)&target=%2F">English</a>
```
After the preferred directory has been selected (for example, "lv"), all links to local HotSpot pages will contain that path (for example, $(link-status) = "http://hotspot.mt.lv/lv/status"). So, if all HotSpot pages reference links using "$(link-xxx)" variables, then no more changes are to be made - each client will stay within the selected directory all the time.

### Misc

If you want to use the HTTP-CHAP authentication method, you need to include the **doLogin()** function (which references to the **md5.js** which must be already loaded) before the **Submit action** of the login form. Otherwise, CHAP login will fail.

The resulting password to be sent to the HotSpot gateway in case of HTTP-CHAP method, it is formed by MD5-hashing the concatenation of the following: chap-id, the password of the user, and chap-challenge (in the given order)

In case variables are to be used in the link directly, then they must be escaped accordingly. For example, in login page, **<a href="https://login.example.com/login?mac=$(mac)&user=$(username)">link</a>** will not work as intended, if username will be "123&456=1 2". In this case instead of $(user), its escaped version must be used: $(user-esc): **<a href="https://login.server.serv/login?mac=$(mac-esc)&user=$(user-esc)">link</a>**. Now the same username will be converted to "123%26456%3D1+2", which is the valid representation of "123&456=1 2" in URL. This trick may be used with any variables, not only with $(username).

There is a boolean parameter "erase-cookie" to the logout page, which may be either "on" or "true" to delete user cookie on logout (so that the user would not be automatically logged on when he/she opens a browser next time.

### Examples

With basic HTML language knowledge and the examples below it should be easy to implement the ideas described above.

- To provide predefined value as username, in login.html change:

<type="text" value="$(username)>

to this line:

<input type="hidden" name="username" value="hsuser">

(where hsuser is the username you are providing)

- To provide a predefined value as a password, in login.html change:

<input type="password">

to this line:

<input type="hidden" name="password" value="hspass">

(where hspass is the password you are providing)

- To send the client's MAC address to a registration server in the form of:

change the Login button link in login.html to:

https://www.example.com/register.html?mac=$(mac)

(you should correct the link to point to your server)

- To show a banner after user login, in alogin.html after

$(if popup == 'true') add the following line:

```
open('http://www.example.com/your-banner-page.html', 'my-banner-name','');
```
(you should correct the link to point to the page you want to show)

- To choose a different page shown after login, in login.html change:

<input type="hidden" name="dst" value="$(link-orig)">

to this line:

<input type="hidden" name="dst" value="http://www.example.com">

(you should correct the link to point to your server)

- To erase the cookie on logoff, in the page containing a link to the logout (for example, in status.html) change:

```
open('$(link-logout)', 'hotspot_logout', ...
```
to this:

```
open('$(link-logout)?erase-cookie=on', 'hotspot_logout', ...
```
or alternatively add this line:

<input type="hidden" name="erase-cookie" value="on">

before this one:

<input type="submit" value="log off">

#### External authentication

Another example is making HotSpot to authenticate on a remote server (which may, for example, perform credit card charging):

- Allow direct access to the external server in walled-garden (either HTTP-based or IP-based)
- Modify the login page of the HotSpot servlet to redirect to the external authentication server. The external server should modify the RADIUS database as needed

Here is an example of such a login page to put on the HotSpot router (it is redirecting to https://auth.example.com/login.php, replace with the actual address of an external authentication server):

```
<html>
<title>...</title>
<body>
<form name="redirect" action="https://auth.example.com/login.php" method="post">
<input type="hidden" name="mac" value="$(mac)">
<input type="hidden" name="ip" value="$(ip)">
<input type="hidden" name="username" value="$(username)">
<input type="hidden" name="link-login" value="$(link-login)">
<input type="hidden" name="link-orig" value="$(link-orig)">
<input type="hidden" name="error" value="$(error)">
</form>
<script language="JavaScript">
<!--
	document.redirect.submit();
//-->
</script>
</body>
</html>
          
```
- The external server can log in a HotSpot client by redirecting it back to the original HotSpot servlet login page, specifying the correct username and password

Here is an example of such a page (it is redirecting to https://hotspot.example.com/login, replace with the actual address of a HotSpot router; also, it is displaying www.mikrotik.com after successful login, replace with what is needed):

```
<html>
<title>Hotspot login page</title>
<body>
<form name="login" action="https://hotspot.example.com/login" method="post">
<input type="text" name="username" value="demo">
<input type="password" name="password" value="none">
<input type="hidden" name="domain" value="">
<input type="hidden" name="dst" value="http://www.mikrotik.com/">
<input type="submit" name="login" value="log in">
</form>
</body>
</html>
          
