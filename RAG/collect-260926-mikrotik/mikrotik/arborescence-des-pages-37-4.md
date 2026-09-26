---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-37-4
title: "arborescence-des-pages-37"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-37.md
source_anchor: ""
source_lines: [364, 505]
sha256: 02baa7b61b2cacfdd3be6f7e09989c9e452f55b934b3ee8f421ae72759ed1cf2
---

# arborescence-des-pages-37

```
- Hotspot will ask the RADIUS server whether to allow the login or not. If allowed, alogin.html page will be displayed (it can be modified to do anything). If not allowed, the flogin.html (or login.html) page will be displayed, which will redirect the client back to the external authentication server.

**Note:** as shown in these examples, HTTPS protocol and POST method can be used to secure communications.

#### HTTP header detection

The Hotspot login pages have access to HTTP headers by using **$(http-header-name);**

For example, there exists an ability to check the user agent (or browser), and will return any other content instead of the regular login page, if so desired. This can be used to disable automatic popups in phones, for example.

For example, to output "SUCCESS" for users of a specific Firefox mobile version, instead of the login page, you can these lines on the top of the **rlogin.html** page in your hotspot directory:

$(if user-agent == "Mozilla/5.0 (Android; Mobile; rv:40.0) Gecko/40.0 Firefox/40.0" ) 
<HTML><HEAD><TITLE>Success</TITLE></HEAD><BODY>Success</BODY></HTML> 
$(else)
---- regular content of rlogin.html page  ----
$(endif)

This will DISABLE the login popup for Android Firefox 40 users.

#### One-click login

It is possible to create a modified captive portal for quick one-click login for scenarios where no user or password is required.

What you need to do is:

- Create a user for this purpose. In example, it is "notsosecretuser" with password "notsosecretpass"
- Assign this user to a user profile that allows a specific/unlimited amount of simultaneous active users.
- Copy original hotspot directory that is already generated in routers file menu on root level.
- Modify the contents of this copy directory contents.
  - Only one file requires modifications for this to work, the "login.html".

Original:

```
<table width="100" style="background-color: #ffffff">
  <tr><td align="right">login</td>
      <td><input style="width: 80px" name="username" type="text" value="$(username)"/></td>
  </tr>
  <tr><td align="right">password</td>
      <td><input style="width: 80px" name="password" type="password"/></td>
  </tr>
  <tr><td> </td>
      <td><input type="submit" value="OK" /></td>
  </tr>
</table>
```
Modified:

```
<table width="100" style="background-color: #ffffff">
  <tr style="display:none;"><td align="right">login</td>
    <td><input style="width: 80px" name="username" type="text" value="notsosecretuser"/></td>
  </tr>
  <tr style="display:none;"><td align="right">password</td>
    <td><input style="width: 80px" name="password" type="password" value="notsosecretpass"/></td>
  </tr>
  <tr><td> </td>
    <td><input type="submit" value="Proceed to Internet!" /></td>
  </tr>
</table>
```
What changed:

    - User and Password "" fields are hidden.
    - Both User and Password field values contain predefined values.
    - Changed the "OK" button value(name) to something more fitting.

- Now upload this new hotspot folder back to the router, preferably with a different name.
- Change settings in the hotspot server profile to use this new html directory.

/ip hotspot profile set (profile number or name) html-directory-override=(dir path/name)

## Firewall customizations

### Summary

Apart from the obvious dynamic entries in the /ip hotspot submenu itself (like hosts and active users), some additional rules are added in the firewall tables when activating a HotSpot service.

### NAT

From **/ip firewall nat print dynamic** command, you can get something like this (comments follow after each of the rules):

 0 D chain=dstnat action=jump jump-target=hotspot hotspot=from-client

Putting all HotSpot-related tasks for packets from all HotSpot clients into a separate chain.

 1 I chain=hotspot action=jump jump-target=pre-hotspot

Any actions that should be done before HotSpot rules apply, should be put in the pre-hotspot chain. This chain is under full administrator control and does not contain any rules set by the system, hence the invalid jump rule (as the chain does not have any rules by default).

 2 D chain=hotspot action=redirect to-ports=64872 dst-port=53 protocol=udp 
 3 D chain=hotspot action=redirect to-ports=64872 dst-port=53 protocol=tcp 

Redirect all DNS requests to the HotSpot service. The 64872 port provides DNS service for all HotSpot users. If you want the HotSpot server to listen to another port, add rules here the same way, changing dst-port property.

```
 4 D chain=hotspot action=redirect to-ports=64873 hotspot=local-dst dst-port=80
     protocol=tcp
```
Redirect all HTTP login requests to the HTTP login servlet. The 64873 is HotSpot HTTP servlet port.

```
 5 D chain=hotspot action=redirect to-ports=64875 hotspot=local-dst dst-port=443
     protocol=tcp
```
Redirect all HTTPS login requests to the HTTPS login servlet. The 64875 is HotSpot HTTPS servlet port.

 6 D chain=hotspot action=jump jump-target=hs-unauth hotspot=!auth protocol=tcp

All other packets except DNS and login requests from unauthorized clients should pass through the hs-unauth chain.

 7 D chain=hotspot action=jump jump-target=hs-auth hotspot=auth protocol=tcp

And packets from the authorized clients - through the hs-auth chain.

```
 8 D ;;; www.mikrotik.com
     chain=hs-unauth action=return dst-address=66.228.113.26 dst-port=80 protocol=tcp
```
First in the **hs-unauth** chain is put everything that affects TCP protocol in the `/ip hotspot walled-garden ip` submenu (i.e., everything where either protocol is not set, or set to TCP). Here we are excluding www.mikrotik.com from being redirected to the login page.

 9 D chain=hs-unauth action=redirect to-ports=64874 dst-port=80 protocol=tcp

All other HTTP requests are redirected to the Walled Garden proxy server which listens to the 64874 port. If there is an "allow" entry in the `/ip hotspot walled-garden` menu for an HTTP request, it is being forwarded to the destination. Otherwise, the request will be automatically redirected to the HotSpot login servlet (port 64873).

10 D chain=hs-unauth action=redirect to-ports=64874 dst-port=3128 protocol=tcp 
11 D chain=hs-unauth action=redirect to-ports=64874 dst-port=8080 protocol=tcp 

HotSpot by default assumes that only these ports may be used for HTTP proxy requests. These two entries are used to "catch" client requests to unknown proxies (you can add more rules here for other ports). I.e., to make it possible for the clients with unknown proxy settings to work with the HotSpot system. This feature is called "Universal Proxy". If it is detected that a client is using some proxy server, the system will automatically mark those packets with the HTTP hotspot mark to work around the unknown proxy problem, as we will see later on. Note that the port used (64874) is the same as for HTTP requests in rule #9 (so both HTTP and HTTP proxy requests are processed by the same code).

12 D chain=hs-unauth action=redirect to-ports=64875 dst-port=443 protocol=tcp

HTTPS proxy is listening on the 64875 port.

13 I chain=hs-unauth action=jump jump-target=hs-smtp dst-port=25 protocol=tcp

Redirect for SMTP protocol may also be defined in the HotSpot configuration. In case it is, a redirect rule will be put in the hs-smtp chain. This is done so that users with unknown SMTP configuration would be able to send their mail through the service provider's (your) SMTP server instead of going to the [possibly unavailable outside their network of origin] SMTP server users have configured on their computers. The chain is empty by default, hence the invalid jump rule.

14 D chain=hs-auth action=redirect to-ports=64874 hotspot=http protocol=tcp

