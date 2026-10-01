---
id: collect-261001-fortinet/fortinet/index-php-document-fortigate-7-2-4-cli-reference-399620-config-webfilter-profile-c776463e-1
title: "index-php-document-fortigate-7-2-4-cli-reference-399620-config-webfilter-profile-c776463e"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/index-php-document-fortigate-7-2-4-cli-reference-399620-config-webfilter-profile-c776463e.md
source_anchor: ""
source_lines: [1, 178]
sha256: 1106c9abfcc0fdcd21ff0d3416d2ca39d7688ec47c8407fda6bd7823543e1dc5
---

# index-php-document-fortigate-7-2-4-cli-reference-399620-config-webfilter-profile-c776463e

config webfilter profile
config webfilter profile
Configure Web filter profiles.
config webfilter profile
    Description: Configure Web filter profiles.
    edit <name>
        config antiphish
            Description: AntiPhishing profile.
            set status [enable|disable]
            set default-action [exempt|log|...]
            set check-uri [enable|disable]
            set check-basic-auth [enable|disable]
            set check-username-only [enable|disable]
            set max-body-len {integer}
            config inspection-entries
                Description: AntiPhishing entries.
                edit <name>
                    set fortiguard-category {user}
                    set action [exempt|log|...]
                next
            end
            config custom-patterns
                Description: Custom username and password regex patterns.
                edit <pattern>
                    set category [username|password]
                    set type [regex|literal]
                next
            end
            set authentication [domain-controller|ldap]
            set domain-controller {string}
            set ldap {string}
        end
        set comment {var-string}
        set extended-log [enable|disable]
        set feature-set [flow|proxy]
        config ftgd-wf
            Description: FortiGuard Web Filter settings.
            set options {option1}, {option2}, ...
            set exempt-quota {user}
            set ovrd {user}
            config filters
                Description: FortiGuard filters.
                edit <id>
                    set category {integer}
                    set action [block|authenticate|...]
                    set warn-duration {user}
                    set auth-usr-grp <name1>, <name2>, ...
                    set log [enable|disable]
                    set override-replacemsg {string}
                    set warning-prompt [per-domain|per-category]
                    set warning-duration-type [session|timeout]
                next
            end
            config quota
                Description: FortiGuard traffic quota settings.
                edit <id>
                    set category {user}
                    set type [time|traffic]
                    set unit [B|KB|...]
                    set value {integer}
                    set duration {user}
                    set override-replacemsg {string}
                next
            end
            set max-quota-timeout {integer}
            set rate-javascript-urls [disable|enable]
            set rate-css-urls [disable|enable]
            set rate-crl-urls [disable|enable]
        end
        set https-replacemsg [enable|disable]
        set log-all-url [enable|disable]
        set options {option1}, {option2}, ...
        config override
            Description: Web Filter override settings.
            set ovrd-cookie [allow|deny]
            set ovrd-scope [user|user-group|...]
            set profile-type [list|radius]
            set ovrd-dur-mode [constant|ask]
            set ovrd-dur {user}
            set profile-attribute [User-Name|NAS-IP-Address|...]
            set ovrd-user-group <name1>, <name2>, ...
            set profile <name1>, <name2>, ...
        end
        set ovrd-perm {option1}, {option2}, ...
        set post-action [normal|block]
        set replacemsg-group {string}
        config web
            Description: Web content filtering settings.
            set bword-threshold {integer}
            set bword-table {integer}
            set urlfilter-table {integer}
            set content-header-list {integer}
            set blocklist [enable|disable]
            set allowlist {option1}, {option2}, ...
            set safe-search {option1}, {option2}, ...
            set youtube-restrict [none|strict|...]
            set vimeo-restrict {string}
            set log-search [enable|disable]
            set keyword-match <pattern1>, <pattern2>, ...
        end
        set web-antiphishing-log [enable|disable]
        set web-content-log [enable|disable]
        set web-extended-all-action-log [enable|disable]
        set web-filter-activex-log [enable|disable]
        set web-filter-applet-log [enable|disable]
        set web-filter-command-block-log [enable|disable]
        set web-filter-cookie-log [enable|disable]
        set web-filter-cookie-removal-log [enable|disable]
        set web-filter-js-log [enable|disable]
        set web-filter-jscript-log [enable|disable]
        set web-filter-referer-log [enable|disable]
        set web-filter-unknown-log [enable|disable]
        set web-filter-vbs-log [enable|disable]
        set web-ftgd-err-log [enable|disable]
        set web-ftgd-quota-usage [enable|disable]
        set web-invalid-domain-log [enable|disable]
        set web-url-log [enable|disable]
        set wisp [enable|disable]
        set wisp-algorithm [primary-secondary|round-robin|...]
        set wisp-servers <name1>, <name2>, ...
    next
end
                                            config webfilter profile
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| comment | Optional comments. | var-string | Maximum length: 255 |  | 
| extended-log | Enable/disable extended logging for web filtering. | option | - | disable | 
|  |  |  |  |  | 
| feature-set | Flow/proxy feature set. | option | - | flow | 
|  |  |  |  |  | 
| https-replacemsg | Enable replacement messages for HTTPS. | option | - | enable | 
|  |  |  |  |  | 
| log-all-url | Enable/disable logging all URLs visited. | option | - | disable | 
|  |  |  |  |  | 
| name | Profile name. | string | Maximum length: 35 |  | 
| options | Options. | option | - |  | 
|  |  |  |  |  | 
| ovrd-perm | Permitted override types. | option | - |  | 
|  |  |  |  |  | 
| post-action | Action taken for HTTP POST traffic. | option | - | normal | 
|  |  |  |  |  | 
| replacemsg-group | Replacement message group. | string | Maximum length: 35 |  | 
| web-antiphishing-log | Enable/disable logging of AntiPhishing checks. | option | - | enable | 
|  |  |  |  |  | 
| web-content-log | Enable/disable logging logging blocked web content. | option | - | enable | 
|  |  |  |  |  | 
| web-extended-all-action-log | Enable/disable extended any filter action logging for web filtering. | option | - | disable | 
|  |  |  |  |  | 
| web-filter-activex-log | Enable/disable logging ActiveX. | option | - | enable | 
|  |  |  |  |  | 
| web-filter-applet-log | Enable/disable logging Java applets. | option | - | enable | 
|  |  |  |  |  | 
| web-filter-command-block-log | Enable/disable logging blocked commands. | option | - | enable | 
|  |  |  |  |  | 
| web-filter-cookie-log | Enable/disable logging cookie filtering. | option | - | enable | 
|  |  |  |  |  | 
| web-filter-cookie-removal-log | Enable/disable logging blocked cookies. | option | - | enable | 
|  |  |  |  |  | 
| web-filter-js-log | Enable/disable logging Java scripts. | option | - | enable | 
|  |  |  |  |  | 
| web-filter-jscript-log | Enable/disable logging JScripts. | option | - | enable | 
|  |  |  |  |  | 
| web-filter-referer-log | Enable/disable logging referrers. | option | - | enable | 
|  |  |  |  |  | 
| web-filter-unknown-log | Enable/disable logging unknown scripts. | option | - | enable | 
|  |  |  |  |  | 
| web-filter-vbs-log | Enable/disable logging VBS scripts. | option | - | enable | 
|  |  |  |  |  | 
| web-ftgd-err-log | Enable/disable logging rating errors. | option | - | enable | 
|  |  |  |  |  | 
| web-ftgd-quota-usage | Enable/disable logging daily quota usage. | option | - | enable | 
|  |  |  |  |  | 
| web-invalid-domain-log | Enable/disable logging invalid domain names. | option | - | enable | 
|  |  |  |  |  | 
| web-url-log | Enable/disable logging URL filtering. | option | - | enable | 
|  |  |  |  |  | 
| wisp | Enable/disable web proxy WISP. | option | - | disable | 
|  |  |  |  |  | 
