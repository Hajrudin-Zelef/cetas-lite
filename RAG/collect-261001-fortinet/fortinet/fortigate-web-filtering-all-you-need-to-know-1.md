---
id: collect-261001-fortinet/fortinet/fortigate-web-filtering-all-you-need-to-know-1
title: "diag webfilter fortiguard cache dump"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-fortinet/fortigate-web-filtering-all-you-need-to-know.md
source_anchor: ""
source_lines: [1, 174]
sha256: ca07fb84482908754a27ec95ac1f24bdb1aa31f420ab6d27d64bc0756db883e4
---

# diag webfilter fortiguard cache dump

## Important facts to know

Main parts of the Web Filtering in Fortigate:

- 
**SSL Profile** - either Certificate-only or Deep SSL Inspection, tells Fortigate whether to decrypt completely SSL communication or look just at domain names in the SSL Certificates. The*no-inspection* profile disables SSL inspection altogether, meaning any HTTPS websites will not be scanned.
- 
**FortiGuard Web Filtering** service - enables us to filter web sites/URLs by Category, instead of static URLs. You need Web Filter license for that.
- 
**Static URL Filter** - checks URL of the web site a user tries to enter, can be used together with Fortiguard Web Filtering.
- 
**Web Content Filtering** - looks inside the**Contents** of a web page to find predefined by us "banned" words and makes decision based on their existence and their count on the page.

Order of processing:

1. 
Static URL Filter
2. 
Fortiguard Category Filter (Local/custom Categories → Remote/external feed Category → FortiGuard Category)
3. 
Web Content Filter
4. 
Advanced options filter - proxy mode only (ActiveX, Java Applets etc.)
5. 
AntiVirus Scanning

**Work Flow** - Fine tune SSL Profile if built-ins are not enough, create webfilter profile in Security Profiles, fine tune it, apply to Security rule, together with SSL profile and enable UTM in such rule.

| Note | Web Filter-related databases on Fortigate are not updated by default, only after you create 1st Web Filter Profile and use it in the security rule, will Fortigate update dbs. | 

The yet-to-be-configured Web Filter Profile will look like:

## Static URL Filter

With this filter we create entries in the filter list, each matching a URL/domain together with the desired action. The list is processed in top to down order, 1st match stops further processing.

The actions:

- 
**Block** - block connection, no other processing (by AV/IPS signatures/etc.) is done. User sees custom or default block page that access was blocked by the policy.
- 
**Allow** - allow connection from URL Filtering point of view, the connection will still  be checked by other security profiles/features if available - IPS/AppControl/AV, etc. This includes FortiGuard/Category based filtering - if Static URL is set to Allow this URL, but in Category it is in the blocked category - the connection will be blocked.
- 
**Exempt** - allow connection AND stop any further security checks like AV/AppControl/IPS. This will exempt the connection even from Category based URL filtering. We can decide what exact further checks to exempt, see below.
- 
**Monitor** - monitor (and log if logs are enabled in the security rule) but do NOT block the connection.

In Static URL filter, we match domain/URL by either **Simple**, **Wildcard**, or **Regex** expressions.

**Simple** is what it says - matches exactly what you put in it. It is the least flexible but also least resource-intensive for the Fortigate. Also, its matching depends on the inspection mode - Proxy or Flow. In proxy the match is exact - e.g. if we set as Simple "yurisk.com" , it will match  just yurisk.com, but not anything else, like any subdomain - www.yurisk.com, test.yurisk.com will NOT be matched. In the Flow mode, it may also do partial matching, including sub-domains.

We may indicate the protocol as well - HTTP or HTTPS, but Fortigate will remove it anyway.

The whole URLs can be matched as well - if I want to block only say a specific page https://yurisk.info/2020/05/20/fortigate-bgp-cookbook-of-example-configuration-and-debug/ but allow anything else in yurisk.info, this will work too:

- 
CLI configuration: create static URL list, then reference it by its table number in the webfilter profile:

```
config webfilter urlfilter
    edit 1   <-- TABLE ID TO BE REFERENCED IN PROFILE
        set name "Auto-webfilter-urlfilter_12gn7p0os"
        config entries
            edit 2
                set url "yurisk.com"
                set action block
            next
            edit 3
                set url "yurisk.info/2020/05/20/
                fortigate-bgp-cookbook-of-example-configuration-and-debug/"
                set action block
            next
        end
    next
end
```
```
config webfilter profile
    edit "WebProfile"
        set feature-set proxy
        config web
            set urlfilter-table 1 
```
**(1)**
        end
1. 
Number 1 is a reference to the URL filter table entry, this way (and only in CLI) we can reference the same URLs table in different web filtering profiles, so not to re-create the same table of URLs for each new profile.

**Wildcard** match - does what the names says, we can replace any part of the domain/URL with * . When using it in the form of *.example.com, this will match all subdomains of the example.com  as well as the root domain example.com itself. Some examples: `*facebook.com`, `forti*.com`

**Regex** matching - Fortinet uses PCRE standard regex syntax. Few notes:

- 
You do not have to escape dot "." used to separate domain name parts, e.g. www.example.com is OK, but "proper" www\.example\.com would work too.
- 
By default, the regex is NOT case-sensitive, but if for some strange reason you need it to be - add after the regex "i" option, e.g. /EXAMPLE.COM/i and it will only match EXAMPLE.COM, not example.com.
- 
We can use "^" to anchor the regex to the beginning of an domain/URL string.
- 
When entering regex on CLI, escape any special symbol twice, see below

E.g. of `urlfilter`  in GUI and on CLI:

CLI:

```
config webfilter urlfilter
    edit 1
        set name "Auto-webfilter-urlfilter_12gn7p0os"
        config entries
            edit 9
                set url "yurisk.info"
            next
            edit 2
                set url "*.yurisk.com"
                set type wildcard
                set action block
            next
            edit 3
                set url "yurisk.info/2020/05/20/fortigate-bgp-cookbook-of-example-configuration-and-debug/"
                set action block
            next
            edit 4
                set url "espn.com"
                set exempt fortiguard
            next
            edit 5
                set url "*.msn.com"
                set type wildcard
                set action block
            next
            edit 6
                set url "stackoverflow.com"
                set type regex
                set action block
            next
            edit 7
                set url "sky\\.com"
                set type regex
                set action block
            next
        end
end
```
**Exempt** - this action allows the stated URL and skips any or specific further processing. The option to specify what further checks to skip is available on CLI only. E.g. here I exempt "espn.com" only from FortiGuard checks, but not from AV or other checks if they are set via Security Profiles in the security rule. As I have the category "Sports" inside Fortigaurd Category filtering set to action Authenticate, this exempt will disable Authentication when entering "espn.com" (and any URL path), but will leave Authentication enabled for any other website in the Sports category, including say espn.co.uk.

CLI config:

```
config webfilter urlfilter
    edit 1
        set name "Auto-webfilter-urlfilter_12gn7p0os"
        config entries
                    edit 4
                set url "espn.com"
                set exempt fortiguard
            next
        end
    next
end
```
Besides `fortiguard` we can exempt any of the below:

 (0) # set exempt
av                     AntiVirus scanning.
web-content            Web filter content matching.
activex-java-cookie    ActiveX, Java, and cookie filtering.
dlp                    DLP scanning.
fortiguard             FortiGuard web filtering.
range-block            Range block feature.
pass                   Pass single connection from all.
antiphish              AntiPhish credential checking.
all                    Exempt from all security profiles. <-- DEFAULT IN GUI

