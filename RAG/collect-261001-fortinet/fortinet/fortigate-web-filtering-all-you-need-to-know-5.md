---
id: collect-261001-fortinet/fortinet/fortigate-web-filtering-all-you-need-to-know-5
title: "diag webfilter fortiguard cache dump"
domain: fortinet
role: reference
task: reference
actors: ["United States"]
dates: ["2025-03-05"]
keywords: ["agent", "license"]
source: docs/RAG/collect-261001-fortinet/fortigate-web-filtering-all-you-need-to-know.md
source_anchor: ""
source_lines: [707, 893]
sha256: 03675a06e85b6016b4bcc58ec73a2369c71413b65c96f66c24757e960e2a6d8f
---

# diag webfilter fortiguard cache dump

Obviously, the Deep SSL Inspection profile has to be used for it to be anything effective.

Web Filter Content, complete - blocks 2 keywords "cisco/Cisco" and "palo*alto" (wildcard)

```
config webfilter content
    edit 1
        set name "Traitors"
        config entries
            edit "palo*alto" <-- KEYWORD TO LOOK FOR, WILDCARD, CASE-SENSITIVE
                set status enable
                set score 10
            next
            edit "cisco"
                set status enable
                set score 10
            next
                        edit "[Cc]isco" 
```
**(1)**
                set pattern-type regexp
                set status enable
            next
        end
    next
end
1. 
Wildcards are **case-sensitive** so just listing "cisco" would only block pages with exact word "cisco", not "Cisco" nor "CISCO". I added**regex** as pattern to account for "cisco" and "Cisco".

```
config entries
    edit "palo*alto"
        set status enable
    next
end
(palo*alto) $ get
name                : palo*alto
pattern-type        : wildcard
status              : enable
lang                : western
score               : 10
action              : block
```
Scoring:

- 
Each separte keyword entry is counted only **once** on the page. In the example above, there may be 10 words "palo alto" on the same page, but Web Filter will only count 1 word/phrase worth score of 10 points.
- 
The default score for each keyword entry is **10** . By "each keyword entry" I mean in the keywords list above - e.g. I have 2 keywords for Cisco, one matches "Cisco", 2nd "cisco" - if on the same page there both "Cisco" and "cisco", it will count as 2 words found, each adding 10 points to score, making it 20 for the page.
- 
The default score for the page to be blocked is **10** , and the default score for each individual keyword entry being also 10 means, by default, any single word from the keywords will cause blocking of the page.
- 
We can change both `bword-threshold` and`score` so that only when multiple keywords appear on the page, the page will be blocked. E.g. if we set ban-word-threshold to 20, then only pages with 2 keywords (each worth 10 points) in them will be blocked.
- 
When using **regex** they seem to count as a single appearance any number of matches. E.g. if I remove from the above keywords list the wildcard "cisco" and leave only regex "[Cc]isco" then page with "Cisco" and "cisco" will count just 1 appearance of this keyword, not 2.
- 
To force FGT to look for the whole phrases only, enclose them in double quotes.

E.g. increasing the banned words threshold per page:

```
config webfilter profile
    edit "WebProfile"
        set feature-set proxy
        config web
            set bword-threshold 20
            set bword-table 1
        end
    next
end
```
Now only the page that has multiple different banned words AND when they add up to 20, will it be blocked.

The block can be seen in Logs → Security Events → Web Filtering:

The end user will see the default block page:

Log on the CLI can be seen with **exe log filter category 3**, **exe log display**:

exe log filter category 3
exe log display
25 logs found.
10 logs returned.
1: date=2025-03-05 time=11:18:31 eventtime=1741166311470738081 tz="+0200"
logid="0314012288" type="utm" subtype="webfilter" eventtype="content" level="warning"
vd="VMVDOM" policyid=9 poluuid="6d0debbe-1755-51ef-ca7d-d938a76591a0" policytype="policy"
sessionid=1880758048 user="yurisk1" authserver="FortiAuth" srcip=192.168.101.0
srcport=54401 srccountry="Reserved" srcintf="Peer1P1" srcintfrole="undefined"
srcuuid="55594f04-1755-51ef-321a-01bd763d3f03" dstip=104.20.4.235 dstport=443
dstcountry="United States" dstintf="VDOM-EXT" dstintfrole="undefined"
dstuuid="99376ffe-9e90-51ea-7ca2-915d04cabdb7" proto=6 httpmethod="GET" service="HTTPS"
hostname="pastebin.com" agent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.
36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36 Edg/133.0.0.0" profile="WebProfile"
reqtype="direct" url="https://pastebin.com/DbquvfL1" sentbyte=4740 rcvdbyte=6576
direction="incoming" action="blocked" banword="[Cc]isco,cisco" msg="URL was blocked
because it contained banned word(s)."

## Proxy Options

Features that work, obviously from the name, in Proxy mode only.

- 
CLI config

```
config web-proxy profile
    edit "Auto-web-proxy-profile_qp9j5z2tc"
        config headers
            edit 1
                set name "X-GoogApps-Allowed-Domains"
                set content "github.com" <---
            next
        end
    next
end
```
```
config webfilter profile
    edit "WebProfile"
        set feature-set proxy
        set options activexfilter cookiefilter javafilter <-- BLOCK THESE
        set post-action block <-- BLOCK "POST" HTTP ACTION
        next
    end
```
- 
Block POST action - will prevent uploading files, authenticating to the web sites.
- 
`cookiefilter` - strips all cookies sent by website, will make 90% of web sites unusable today.
- 
`activexfilter` - strips all ActiveX applets from a web page, quite redundant today - no reputable site uses ActiveX any more, and all browsers have it disabled anyway.
- 
`javafilter` - strip all Java applets, the same - no one uses them anymore and browsers have it blocked by default.

## Video Filter (not part of Web Filtering)

Allows to filter Youtube content by Fortiguard categories or by Youtube Channel ID you put manually, requires Channel ID or/and Youtube API key.

## Debug and Verification

- 
First, let’s check that the Fortigate has the Web Filtering license active (for Fortiguard-based categories) with **dia deb rating** :

```
# diagnose debug rating
Locale       : english
Service      : Web-filter
Status       : Enable <-- WILL BE DISABLED IF NO WEB FILTER IS USED
                        IN SECURITY RULES
License      : Contract
Service      : Antispam
Status       : Disable
Service      : Virus Outbreak Prevention
Status       : Disable
Num. of servers : 1 <-- WITH ANYCAST, WITH UNICAST MANY SERVERS WILL
                        BE LISTED
Protocol        : https
Port            : 443
Anycast         : Enable
Default servers : Included
-=- Server List (Thu Mar 13 02:55:21 2025) -=-
IP
Weight
RTT Flags
TZ
FortiGuard-requests
Curr
Lost
Total Lost
Updated Time
173.243.141.16
-11420
89
DI
0
5710
0
0
Thu Mar 13 02:55:20 2025
```
Also, make sure under **config sys fortiguard** there is no **set webfilter-force-off enable** which turns off the FortiGuard web filtering service.

In GUI, System → Fortiguard, the valid license looks like:

- 
Make sure there are no other strange configs under sys fortiguard:

