---
id: collect-261001-fortinet/fortinet/fortigate-web-filtering-all-you-need-to-know-3
title: "diag webfilter fortiguard cache dump"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2025-03-06"]
keywords: ["agent", "attention"]
source: docs/RAG/collect-261001-fortinet/fortigate-web-filtering-all-you-need-to-know.md
source_anchor: ""
source_lines: [338, 528]
sha256: 0e37e38f0570b397712c874dd2ad4990355721cf48060490327805b33cc60e80
---

# diag webfilter fortiguard cache dump

If enabled, this feature allows end user, after trying to reach site that is blocked per category in the current Web Filter profile, to enter his user/pass and have Web Filter profile replaced with another profile configured beforehand. This way we can allow users in a particular user group to switch to more allowing Web Filter profile.

E.g. I create a new Web Filter profile that allows more categories than the main Web Filter, including Social Networks and News & Media named "OverrideAllowAll" (not shown here).

Then I create the main Web Filter called "WebProfile", which lists that "OverrideAllowAll" profile as a replacement after the user successfully authenticates, and in which ("WebProfile") I block Social Networks and News & Media:

In the section below, I set the user group to authenticate against (can be local and can be remote - LDAP/RADIUS) and replacement Web Filter profile for such authenticated users:

Also, note that I set override duration to 5 mins, which means after 5 minutes user will be again blocked from accessing the site with option to authenticate again.

Page with the option to override the Web Filter profile that blocks access as seen by end user:

After clicking "Override":

Pay attention to the port used by Fortigate to authenticate end user, make sure it is open:

The ports used by Fortigate for this and other features are listed here, and can be set to anything you want:
**config webfilter fortiguard**

config webfilter fortiguard
(fortiguard) # get
cache-mode          : ttl
cache-prefix-match  : enable
cache-mem-permille  : 1
ovrd-auth-port-http : 8008
ovrd-auth-port-https: 8010
ovrd-auth-port-https-flow: 8015
ovrd-auth-port-warning: 8020
ovrd-auth-https     : enable
warn-auth-https     : enable
close-ports         : disable
request-packet-size-limit: 0
embed-image         : enable

Create user group:

```
config user group
    edit "webauthgrp"
        set member "yurisk4" "localvpn1"
    next
end
```
As categories are all appear as numbers, it is easier to configure in GUI, but for completeness sake, here is the complete Web Filter profile:

```
config webfilter profile
    edit "WebProfile"
        set feature-set proxy
        set ovrd-perm bannedword-override urlfilter-override fortiguard-wf-override contenttype-check-override
        config override
            set ovrd-scope ip
            set ovrd-dur 5m
            set ovrd-user-group "webauthgrp"
            set profile "OverrideAllowAll"
        end
        config ftgd-wf
            unset options
            set ovrd 83 96 98 99 26 61 86 88 90 23 30 36 37
            config filters
                edit 1
                    set category 2
                    set action warning
                next
                edit 2
                    set category 7
                    set action warning
                next
                ... CUT ...
                edit 19
                    set category 61
                    set action block
                next
                edit 20
                    set category 86
                    set action block
                next
                edit 21
                    set category 88
                    set action block
                next
                edit 22
                    set category 90
                    set action block
                next
                edit 23
                    set category 23
                    set action block
                next
                edit 24
                    set category 96
                    set action block
                next
                edit 25
                    set category 98
                    set action block
                next
                edit 26
                    set category 99
                    set action block
                next
                edit 27
                    set category 83
                    set action block
                next
                edit 28
                    set category 4
                next
                edit 29
                    set category 1
                next
                edit 30
                    set category 3
                next
                edit 31
                    set category 31
                next
                edit 32
                    set category 59
                next
                edit 33
                    set category 62
                next
                edit 34
                    set category 6
                next
                edit 46
                    set category 46
                    set action authenticate
                    set warn-duration 2h2m
                    set auth-usr-grp "webauthgrp"
                next
                edit 37
                    set category 37
                    set action block
                next
                edit 38
                    set category 30
                    set action block
                next
                edit 39
                    set category 36
                    set action block
                next
            end
        end
    next
end
```
Pay attention - we can NOT specify which category to override - ALL blocked categories are overriden by new Web Filter profile "OverrideAllowAll" and will be blocked/allowed according to the new profile.

**Verification and debug** - I didn’t find yet much info on this, except that we can clear all overrides/warning periods and thus force users to authenticate again with **dia test app ovrd 6**:

dia test application ovrd 333
1.   This menu
2.   Display stats
5.   Clear all user override entries in all vdoms
6.   Clear all warning/authentication entries in all vdoms
99.  Restart the ovrd daemon.

For debug **dia deb app urlfilter 250** will give A LOT of output, beware. And **dia deb app urlfilter -1** will given even more .

In logs of Web Filtering, all access will be with the action Block, even after authentication and actually successfully accessing the website:

2: date=2025-03-06 time=10:19:34 eventtime=1741285174244223171 tz="-0800"
logid="0316013056" type="utm" subtype="webfilter" eventtype="ftgd_blk" level="warning"
vd="root" policyid=21 poluuid="06ac0942" policytype="policy"
sessionid=10154 user="localvpn1" group="ipsecgrp" authserver="localvpn1"
srcip=192.168.17.0 srcport=53720 srccountry="Reserved" srcintf="IKEv1" srcintfrole="undefined" srcuuid="c796b842"
dstip=23.200.96.79 dstport=443 dstcountry="Ireland" dstintf="port1"
dstintfrole="undefined" dstuuid="76ab0144"
proto=6 httpmethod="GET" service="HTTPS" hostname="news.sky.com"
agent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36
(KH" profile="WebProfile" action="blocked" reqtype="referral"
url="https://news.sky.com/world"
referralurl="https://news.sky.com/us"
sentbyte=4107 rcvdbyte=643 direction="outgoing"
msg="URL belongs to a denied category in policy"
ratemethod="domain"
cat=36 catdesc="News and Media"

### Usage Quota

When a Category has action set to one of **Warning, Authenticate**, or **Monitor**, we can limit how much bandwidth or time the end user can consume from such a category. After a user uses up her quota (for this day) she will be blocked from the websites in such category. The usage counters reset automatically each midnight. The quota is calculated per user, and for the whole category - i.e. if a user used up Sport Category by browsing espn.com, she will be blocked for any other website which is also in this category.

| Important | Quotas will only work in **Proxy-mode** webfilters/security rules. | 

The end user will see such message of using up her quota:

In logs we will see UTM block once the quota is reached (here user blocked after using up 10 Mbytes on espn.com - Sports Category):

