---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-37-1
title: "arborescence-des-pages-37"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-37.md
source_anchor: ""
source_lines: [1, 76]
sha256: 7232224bde6398dcbbc92b5e908fdc08484d53c2251c8e131ea8cd3ffd19b3fb
---

# arborescence-des-pages-37

### Requirements

### Introduction

You can create a completely different set of servlet pages for each HotSpot server you have, specifying the directory in the "html-override-directory" property of a HotSpot server profile /ip hotspot profile. The default servlet pages are copied in the directory "hotspot" directory right after you create the server profile. This directory can be accessed by connecting to the router with an FTP client. You can copy this directory and modify the pages as you like using the information from this section of the manual. Note that it is suggested to edit the files manually, as automated HTML editing tools may corrupt the pages by removing variables or other vital parts. After you are finished with content modification you need to upload this modified content to some custom directory on the hotspot router and point previously mentioned property "html-override-directory" value as path to this new custom HTML directory.

**Note:** If "html-override-directory" value path is missing or empty then the hotspot server will revert to default HTML files.

### Available Pages

Main HTML servlet pages, which are shown to the user:

- **redirect.html** - redirects the user to another URL (for example, to the login page)
- **login.html** - login page shown to a user to ask for a username and password. This page may take the following parameters:
  - **username** - username
  - **password** - either plain-text password (in case of PAP authentication) or MD5 hash of chap-id variable, password, and CHAP challenge (in case of CHAP authentication). This value is used as e-mail address for trial users
  - **dst** - original URL requested before the redirect. This will be opened on successful login
  - **popup** - whether to pop-up a status window on successful login
  - **radius<id>** - send the attribute identified with <id> in text string form to the RADIUS server (in case RADIUS authentication is used; lost otherwise)
  - **radius<id>u** - send the attribute identified with <id> in unsigned integer form to the RADIUS server (in case RADIUS authentication is used; lost otherwise)
  - **radius<id>-<vnd-id>** - send the attribute identified with <id> and vendor ID <vnd-id> in text string form to the RADIUS server (in case RADIUS authentication is used; lost otherwise)
  - **radius<id>-<vnd-id>u** - send the attribute identified with <id> and vendor ID <vnd-id> in unsigned integer form to the RADIUS server (in case RADIUS authentication is used; lost otherwise)
- **md5.js** - JavaScript for MD5 password hashing. Used together with http-chap login method
- **alogin.html** - page shown after a client has logged in. It pops-up status page and redirects the browser to the originally requested page (before he/she was redirected to the HotSpot login page)
- **status.html** - status page, shows statistics for the client. It is also able to display advertisements automatically
- **logout.html** - logout page, shown after a user is logged out. Shows final statistics about the finished session. This page may take the following additional parameters:
  - **erase-cookie** - whether to erase cookies from the HotSpot server on logout (makes it impossible to log in with cookie next time from the same browser, might be useful in multiuser environments)
- **error.html** - error page, shown on fatal errors only

Some other pages are available as well, if more control is needed:

- **rlogin.html** - page, which redirects the client from some other URL to the login page, if authorization of the client is required to access that URL
- **rstatus.html** - similar to rlogin.html, only in case if the client is already logged in and the original URL is not known
- **radvert.html** - redirects the client to the scheduled advertisement link
- **flogin.html** - shown instead of login.html, if some error has happened (invalid username or password, for example)
- **fstatus.html** - shown instead of redirect, if a status page is requested, but the client is not logged in
- **flogout.html** - shown instead of redirect, if logout page is requested, but the client is not logged in

### Serving Servlet Pages

The HotSpot servlet recognizes 5 different request types:

1. **request for a remote host**
  - if user is logged in and advertisement is due to be displayed, radvert.html is displayed. This page redirects to the scheduled advertisement page
  - if user is logged in and advertisement is not scheduled for this user, the requested page is served
  - if user is not logged in, but the destination host is allowed by the walled garden, then the request is also served
  - if user is not logged in, and the destination host is disallowed by the walled garden, rlogin.html is displayed; if rlogin.html is not found, redirect.html is used to redirect to the login page
2. **request for "/" on the HotSpot host**
  - if user is logged in, rstatus.html is displayed; if rstatus.html is not found, redirect.html is used to redirect to the status page
  - if user is not logged in, rlogin.html is displayed; if rlogin.html is not found, redirect.html is used to redirect to the login page
3. **request for "/login" page**
  - if user has successfully logged in (or is already logged in), alogin.html is displayed; if alogin.html is not found, redirect.html is used to redirect to the originally requested page or the status page (in case, the original destination page was not given)
  - if user is not logged in (username was not supplied, no error message appeared), login.html is showed
  - if login procedure has failed (an error message is supplied), flogin.html is displayed; if flogin.html is not found, login.html is used
  - in case of fatal errors, error.html is showed
4. **request for "/status" page**
  - if user is logged in, status.html is displayed
  - if user is not logged in, fstatus.html is displayed; if fstatus.html is not found, redirect.html is used to redirect to the login page
5. **request for '/logout' page**
  - if user is logged in, logout.html is displayed
  - if user is not logged in, flogout.html is displayed; if flogout.html is not found, redirect.html is used to redirect to the login page

**Note:** If it is not possible to meet a request using the pages stored on the router's FTP server, Error 404 is displayed

There are many ways to customize what the HotSpot authentication pages look like:

- The pages are easily modifiable. They are stored on the router's FTP server in the directory you choose for the respective HotSpot server profile.
- By changing the variables, which client sends to the HotSpot servlet, it is possible to reduce the keyword count to one (username or password; for example, the client's MAC address may be used as the other value) or even to zero (License Agreement; some predefined values general for all users or client's MAC address may be used as username and password)
- Registration may occur on a different server (for example, on a server that is able to charge Credit Cards). Client's MAC address may be passed to it, so that this information doesn't have to be entered manually. After the registration, the server should change RADIUS database enabling client to log in for some amount of time.

To insert a variable in some place in the HTML file, the $(var_name) syntax is used, where the "var_name" is the name of the variable (without quotes). This construction may be used in any HotSpot HTML file accessed as '/', '/login', '/status' or '/logout', as well as any text or HTML (.txt, .htm or .html) file stored on the HotSpot server (with the exception of traffic counters, which are available in status page only, and **error**, **error-orig**, **chap-id**, **chap-challenge** and **popup** variables, which are available in login page only). For example, to show a link to the login page, following construction can be used:

<a href="$(link-login)">login</a>

### Variables

