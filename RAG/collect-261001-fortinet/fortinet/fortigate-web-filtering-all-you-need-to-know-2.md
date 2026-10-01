---
id: collect-261001-fortinet/fortinet/fortigate-web-filtering-all-you-need-to-know-2
title: "diag webfilter fortiguard cache dump"
domain: fortinet
role: reference
task: reference
actors: ["Apple", "Google", "United States"]
dates: ["2025-03-05"]
keywords: ["agent", "distribution", "license"]
source: docs/RAG/collect-261001-fortinet/fortigate-web-filtering-all-you-need-to-know.md
source_anchor: ""
source_lines: [175, 337]
sha256: 72dd9aff9d28fb37d0d2cdc51765338ce9dd4046dbea6a03ab75007c6177b7f9
---

# diag webfilter fortiguard cache dump

The result of the above configuration will look in GUI as:

In logs this Exempt will appear as "Passthrough":

## FortiGuard Category based Web filtering

For this filtering, the Fortigate on each web site request by users (responses are cached) queries the FortiGuard Distribution Network (FDN). The FDN collects billions of web sites categorizing them, by both domain names and IP addresses.

Prerequisites:

- 
Valid Web Filtering license, see in Dashboard or **dia debug rating**
- 
Being able to connect to FortiGuard servers, again start with **dia debug rating** to verify.
- 
As an option, Fortimanager that does have access to Fortiguard can be configured as FDS to be used by Fortigates.
- 
You have to use SSL inspection Profile in security rule. When using Certitificate-only profile, the Fortigate will only be able to check **domain name** of the website, not page contents, nor the full path in URL (part after the 1st slash in URL). For maximum efficiency, the Deep SSL Inspection SSL profile should be used (but beware of browser error on MiTM - have to 1st install the Fortigate CA certificate on all end stations).
- 
Watch out for **SSL Exemption** list inside the SSL Inspection Profile - categories/domains/IP addresses listed there will NOT be inspected if they use SSL (most of the websites) and thus features like Usage Quota/Warning/etc. will work but in a bit different way. See Usage Quota.

| Tip | To see category of a domain, you either check Fortiguard site https://www.fortiguard.com/webfilter or can go to Security Profiles → Web Rating Overrides → Select Create New → URL put the domain in question and click "Lookup Rating" | 

To list all categories on CLI: **get webfilter categories**.

Configuring it is simple - just enable it in Web Filter Profile as Category Based Filter, choose Categories and their Action and use this Web Filter in security rule(s) with UTM enabled.

- 
**Block** - conneciton to the web site is blocked, no further security processing is done, the verdict is final.
- 
**Allow** - allow connection from Web Filtering point of view, the connection will still  be checked by other security profiles/features if available - IPS/AppControl, etc and thus may be blocked later.
- 
**Monitor** - monitor (and log if logs are enabled in the security rule) but do NOT block the connection. Mostly useful to get detailed info on web sites users are visiting.
- 
**Authenticate** - access will be granted after the end user successfully authenticates to Fortigate.
- 
**Warning** - users will see a web page warning them that they are entering restricted domains, and if the user clicks on "Proceed" she will be redirected to the original web site.

Categories order of precedence: Local Category > Remote Category > Fortiguard Category.

To change the contents of **Replacement Page**, go to System → Replacement messages → click on Extended view → click on the needed template → Edit. We can change the message for actions that intercept user’s connection: Block, Warning, Authenticate.

### Category cache verification

All mappings of domains to the Category from FortiGuard Fortigate keeps in its cache, so if we want to verify what category actually a given visited domain got, we run **diag webfilter fortiguard cache dump**:

diag webfilter fortiguard cache dump
# diag webfilter fortiguard cache dump
Caution: This command is for diagnostic purposes ONLY.
The bigger the cache size is set, the more impact on
performance the command has.
Do you want to continue? (y/n)
Saving to file [/tmp/urcCache.txt]
Cache Contents:
-=-=-=-=-=-=-=-
Cache Mode:   TTL
Cache DB Ver: 234.25972
Rating            DB Ver   DOT  SLASH ORIG_FLAG T URL
00000000|00000000 234.25972     2    0 00000001 P Dhttps://104.26.8.62/
2e000000|2e000000 234.25970     1    0 00000102 E Dhttps://www.espn.com/
00000000|00000000 234.25970     1    0 00000001 P Dhttps://4.207.247.139/
34000000|34000000 234.25970     1    0 00000001 P Dhttps://client.wns.windows.com/
34000000|34000000 234.25970     1    0 00000001 P Dhttps://8.8.4.4/
34000000|34000000 234.25970     1    0 00000001 P Dhttps://dns.google/

The 1st number is the category in hex. After we translate it to decimal, we can compare with the built-in categories of the Fortigate **get webfilter categories | grep *n***:

E.g. for espn.com the hex is 2e = 46:

```
get webfilter categories | grep 46
     46 Sports
```
As we can see, most of the IP addresses do not get their rating at all, except the well-known Google DNS, that is because I have option "Rate URLs by domain and IP Address" disabled.

### Action - Authenticate

Here we are not blocking or allowing access to a Category immediately, but require additional authentication from the user to proceed and view the web site. Users can be local or remote (e.g. RADIUS/LDAP, but no FSSO/SAML meantime). Once we have user/group, we can set them in the Web Filter → Category → Authenticate. Each time user tries to access such category, she will be asked to enter user/pass, and will be allowed access for the Warning period of time (default = 5 minutes), after which the access again is blocked until the user again authenticates.

| Note | FortiOS up to 7.4.4 require proxy-mode policy for this feature to work. | 

GUI Config:

- 
Enable Category filtering, find the needed Category (here Sport), and click on Action = Authenticate.

- 
On trying to enter the website belonging to the restricted Category, the user will see:

And after clicking on "Proceed":

And after entering the user/pass combo correctly, the user will be redirected to the original website, here espn.com.

- 
In logs the above will be seen as:

# exe log filter category 3
# exe log display
899 logs found.
10 logs returned.
1: date=2025-03-05 time=15:56:37 eventtime=1741182997112761866 tz="+0200"
logid="0316013057" type="utm" subtype="webfilter" eventtype="ftgd_blk" level="warning"
vd="VMVDOM" policyid=9 poluuid="6d0debbe-1755-51ef-ca7d-d938a76591a0" policytype="policy"
sessionid=1999198506 user="yurisk1" authserver="FortiAuth" srcip=192.168.101.0
srcport=54983 srccountry="Reserved" srcintf="Peer1P1" srcintfrole="undefined"
srcuuid="55594f04-1755-51ef-321a-01bd763d3f03" dstip=3.169.71.125 dstport=443
dstcountry="United States" dstintf="VDOM-EXT" dstintfrole="undefined"
dstuuid="99376ffe-9e90-51ea-7ca2-915d04cabdb7" proto=6 httpmethod="GET" service="HTTPS"
hostname="www.espn.com" agent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.
36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36 Edg/133.0.0.0" profile="WebProfile"
action="blocked" reqtype="referral" url="https://www.espn.com/service-worker.js"
direction="outgoing" msg="URL belongs to a category with warnings enabled"
ratemethod="domain" cat=46 catdesc="Sports"

- 
CLI configuration

Create a local user:

```
config user local
    edit "yurisk4"
        set type password
        set passwd-time 2025-03-05 11:46:59
        set passwd ENC +UY9yWP==
    next
end
```
Create a Firewall user group and add the user to it:

```
config user group
    edit "webauthgrp"
        set member "yurisk4"
    next
end
```
And finally, configure the Web Filter profile:

```
config webfilter profile
    edit "WebProfile"
        set feature-set proxy
        config web <-- CONTENT KEYWORD BLOCKING, UNRELATED
            set bword-threshold 20
            set bword-table 1
        end
        config ftgd-wf
            unset options
            config filters
                edit 46  <-- 46 IF SPORTS CATEGORY
                    set category 46
                    set action authenticate
                    set warn-duration 2h2m <-- INCREASE WARN TIME TO 2H
                    set auth-usr-grp "webauthgrp"
                next
            end
        end
    next
end
```
### Allow User Override

