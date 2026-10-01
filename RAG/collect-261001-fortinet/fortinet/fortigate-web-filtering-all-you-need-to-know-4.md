---
id: collect-261001-fortinet/fortinet/fortigate-web-filtering-all-you-need-to-know-4
title: "diag webfilter fortiguard cache dump"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-web-filtering-all-you-need-to-know.md
source_anchor: ""
source_lines: [529, 706]
sha256: 9cfb62cbf90bf62e2522e72cdd5f3e9d8b3583f3068a02af354249b9f335a864
---

# diag webfilter fortiguard cache dump

Note: Also check SSL Exemption list in the applied SSL Inspection Profile - if the website you are trying to enforce Usage Quota is in the list - Quota will work but ONLY for new connections. I.e. if I set 5 minutes quota on Finance Category - all websites here are by default exempted from SSL Inspection, so Fortigate has no control over already established browser session with such website. Once user exhausts her quota to say Paypal.com - she will CONTINUE browsing to the website uninterrupted, BUT if she tries to enter any new website in the Finance Category - she will be blocked on used quota. The Paypal.com will only be blocked if she opens a new browser window to it, or closes and opens browser window completely, then try to enter Paypal.com. The funny thing - logs will show "UTM blocked" on quota, but actually user will continue browsing just fine until 1 of the things above happen. So, it is not a bug but a feature.

- 
How do I know if a website is SSL exempted? 2 ways - in the browser check the SSL/TLS certificate of the website - if it is the original certificate of the website - it is exempted, but if it is a Fortigate CA certificate - then it is not (given SSL Deep Inspection is used). The other way is in Fortigate logs - the action for such website will be "passthrough":

Check SSL Exemption list in **Security Profiles → SSL/SSH Inspection**:

- 
CLI config

```
config webfilter profile
    edit "WebProfile"
        set feature-set proxy
        config ftgd-wf
            unset options
            config filters
                edit 1
                    set category 2
                    set action warning
                next
                edit 31
                    set category 31
                next
                edit 46
                    set category 46
                    set action authenticate
                    set warn-duration 2h2m
                    set auth-usr-grp "webauthgrp"
                next
            end
            config quota
                edit 1
                    set category 46 <-- SPORTS CATEGORY
                    set type traffic
                    set value 10 <-- 10 MBYTES
                next
                edit 2
                    set category 31 <-- FINANCE CATEGORY
                next
            end
        end
    next
end
```
### Custom/local Categories and Web rating Override

We can override/reassign a specific website to another Fortiguard category or to the Local/Custom category. This way we can change what happens when a user tries to access it.

E.g. I will re-assign www.tripadvisor.com from Travel to the Local **custom1** category and then will set the Action to Block on it, while leaving action for the Travel category Allow.

Assign the website to the Local *custom1* category in **Security Profiles → Web Rating Overrides**:

Change the Action to Block in the actual Web Filterfor *custom1*:

Now, even though "Travel" category is allowed, the specific website is being blocked on "custom1" category:

CLI configuration:

```
config webfilter ftgd-local-rating
    edit "www.tripadvisor.com"
        set rating 140 <-- CUSTOM1 CATEGORY
    next
end
```
Set action in Fortiguard-based filter to Block:

```
config webfilter profile
    edit "WebProfile"
        set feature-set proxy
        config ftgd-wf
            unset options
            config filters
            edit 0
            set category 140
                    set action block
                next
            end
```
### Remote Category filter for external threat feed

This option allows us to point Fortigate to external web server that contains as a plain text file list of URLs we want to act upon - Block/Allow/Exempt.  As I already said - the priority of categories is **Local → External/Remote → Fortiguard**.

First, we create a text file with URLs and host it on external web server. The file will look like:

https://yurisk.info/2025/02/26/fortigate-dlp-file-filtering-and-more-examples/
*.sky.com
cnn.com

1st is a complete URL, 2nd is a wildcard (no regex is supported) that will also block `sky.com`, and 3rd is an exact match.

Now, we need to create **External Connector** of type **Fortiguard Category**:

| Note | Do NOT, by mistake, chose **Domain Name** as this is for DNS FIlter, not Web Filter. | 

In which we set URL to access the external feed, optional authentication and refresh rate of data from remote server (default 5 mins):

After a few seconds, the status will change to green - Connected, also we can click on "View" to actually see the URLs read from the remote server and their (URLs) status:

Now, the category "Remote" will appear in all existing and future Web Filter Profiles (by defult in status Disabled) with the external feed we just configured, make sure to change the Action to Block/Monitor/etc.:

If a user tries to enter the website from the externel feed (and I set action to Block) - it will be blocked, even though the Fortiguard category for this website is set to Allow (categories precedence):

The block message the end user will see:

Also, the FGT blocks access to the full path URL but not to the whole website (yurisk.info):

Again, it is possible because I have Deep SSL Inspection in Web Profile, otherwise FGT would not be able to see beyong 1st slash after .info in https://yurisk.info/

This Remote feed will also be available in SSL Profile for Exemption if we need to:

## Search Engines Safe Search and Vimeo

Not much to configure here - just enable or not the option to redirect end users to the Safe Search page of the listed Search Engines. The "safety" of the search is enforced by the search providers themselves, not by the Fortigate. The config is done under Static URL section (proxy-mode only feature):

CLI:

```
config webfilter profile
    edit "WebProfile"
        set feature-set proxy
        config web
            set bword-threshold 20
            set bword-table 1
            set safe-search url header <-- SAFE SEARCH CONFIG STARTS HERE
            set youtube-restrict strict
            set log-search enable
        end
    end
```
The **Vimeo** option is available in CLI only:

```
config web
    set bword-threshold 20
    set bword-table 1
    set safe-search url header
    set youtube-restrict strict
    set log-search enable
end
(web) $ get
bword-threshold     : 20
bword-table         : 1 <-- CONTENT BLOCK TABLE
urlfilter-table     : 0 <-- STATIC URL FILTER TABLE
content-header-list : 0
blocklist           : disable  <-- USE URL LIST RECEIVED FROM FSA TO BLOCK
allowlist           :
safe-search         : url header
youtube-restrict    : strict
vimeo-restrict      :    <-- VIMEO
log-search          : enable
keyword-match       :
```
Fortigate supports 2 categories for VIMEO - 7 and 134.

- 
7 - block mature content
- 
134 - block both unrated and mature content

(web) $ set vimeo-restrict 7

## Rate by both IP Address and Domain

Regularly, the FGT asks Fortiguard for categorization/rating of the **domain** the end user is trying to access only. If we enable this option, FGT will ask Fortiguard for 2 ratings - 1st of the domain (as usual), and the 2nd - of IP Address this domain resolves to. If they differ, FGT will use rating weight of each returned category - the one having higher weight will be used. I haven’t seen many admins using this option.

## Block Invalid URLs

This feature, if enabled, blocks access to websites whose SSL/TLS certificate does not contain a valid domain. Usually this feature is enabled by default.

## Content Web Filtering

Here, the Fortigate checks the contents of a web page, looking for pre-defined by us keywords or whole phrases (up to 80 characters long). We can also assign *score* to each such keyword (the default being 10), as FGT counts appearances of each keyword and sums up their scores before making final decision.

| Note | The Security Policy where such web filter is used, **has** to be in Proxy mode, not Flow. | 

