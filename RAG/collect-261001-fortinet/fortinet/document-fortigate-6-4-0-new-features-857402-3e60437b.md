---
id: collect-261001-fortinet/fortinet/document-fortigate-6-4-0-new-features-857402-3e60437b
title: "document-fortigate-6-4-0-new-features-857402-3e60437b"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-6-4-0-new-features-857402-3e60437b.md
source_anchor: ""
source_lines: [1, 148]
sha256: 8e875db85aaa909b13db5690b033621d9915876613a83c6b99b32a32d610c5dc
---

# document-fortigate-6-4-0-new-features-857402-3e60437b

Security Profiles enhancements
Security Profiles enhancements
Feature set option
To more clearly show the features specific to proxy-based mode, use the new Feature set option to select Flow-based or Proxy-based. When you select Flow-based or Proxy-based, only the features for that mode are available.
The following pages have the Feature set option:
- Security Profiles > AntiVirus
- Security Profiles > Web Filter
- Security Profiles > Email Filter
- Security Profiles > Data Leak (CLI only)
- Policy & Objects > Protocol Options
Example of the Feature set option in Security Profiles > AntiVirus:
If you select Proxy-based, a red P icon indicates the proxy-only features. FortiOS.
When you configure firewall policies:
- If the inspection mode is flow-based, dropdown menus only display profiles with flow-based feature sets.
- If the inspection mode is proxy-based, dropdown menus display profiles with flow-based or proxy-based feature sets.
If a flow-based inspection policy has a proxy-based profile assigned, a warning icon and tooltip informs you that proxy features do not work in a flow-based policy. This warning also appears when you use the CLI to assign security profiles.
Upgrade support
Upgrading from 6.2.x to 6.4.0 causes the following changes to security profiles.
| Upgrade scenario | Result after upgrade | 
|---|---|
| Profile was assigned exclusively to flow-base firewall policies in 6.2.x. | feature-set = flow | 
| Profile was assigned exclusively to proxy-base firewall policies in 6.2.x. | feature-set = proxy | 
| Profile was assigned to both flow-base and proxy-base firewall policies in 6.2.x. | feature-set = proxy | 
| Profile was not assigned to any firewall policies in 6.2.x. | feature-set = flow | 
Configure security profiles using CLI
To configure the Antivirus security profile using the CLI:
config antivirus profile
    edit new-av-profile
        set comment <string>
        set feature-set {flow | proxy}
        set ftgd-analytics {disable | suspicious | everything}
        ...
    next
end
                                            See Configure Antivirus profiles in the FortiOS CLI Reference for more information.
To configure the Web Filter security profile using the CLI:
config webfilter profile
    edit "new-wf-profile"
        set comment <string>
        set feature-set {flow | proxy}
        ...
        config ftgd-wf
            unset options
            config filters
                ...
            end
        end
    next
end
                                            See Configure Web filter profiles in the FortiOS CLI Reference for more information.
To configure the Email Filter security profile using the CLI:
config emailfilter profile
    edit "new-ef-profile"
        set comment <string>
        set feature-set {flow | proxy}
        ...
    next
end
                                            See Configure Email filter profiles in the FortiOS CLI Reference for more information.
To configure the DLP security profile using the CLI:
config dlp sensor
    edit "new-dlp-profile"
        set comment <string>
        set feature-set {flow | proxy}
        ...
    next
end
                                            See Configure DLP sensors in the FortiOS CLI Reference for more information.
To configure Protocol Options in Policy & Objects using the CLI:
config firewall profile-protocol-options
    edit "new-protocol-options"
        set feature-set {flow | proxy}
        config http
            set ports 80
            unset options
            unset post-lang
        end
        config ftp
            set ports 21
            set options splice
        end
        config imap
            set ports 143
            set options fragmail
        end
        ...
    next
end
                                            See Configure protocol options in the FortiOS CLI Reference for more information.
Antivirus profiles use hybrid scanning as default
In flow-based Antivirus profiles, the scan-mode option is removed. Flow-based Antivirus profiles use the default hybrid scanning method to process traffic. Legacy mode is available for diagnostics only.
|  | When upgrading from 6.2.x to 6.4.0, Antivirus profiles assigned to flow-based firewall policies only operate in the default hybrid mode regardless of the previous scan-mode setting. | 
In CLI, scan-mode options are only available for proxy-based Antivirus profiles. The scan-mode options are not available for flow-based Antivirus profiles.
config antivirus profile
    edit "new-av-profile"
        set comment ''
        set replacemsg-group ''
        set feature-set proxy
        set mobile-malware-db enable
        config http
            unset options
            unset archive-block
            unset archive-log
            set emulator enable
            set outbreak-prevention disabled
        end
        ...
        set av-virus-log enable
        set av-block-log enable
        set extended-log disable
        set scan-mode default
    next
end
set ?
comment              Comment.
replacemsg-group     Replacement message group customized for this profile.
feature-set          Flow/proxy feature set.
mobile-malware-db    Enable/disable using the mobile malware signature database.
av-virus-log         Enable/disable AntiVirus logging.
av-block-log         Enable/disable logging for AntiVirus file blocking.
extended-log         Enable/disable extended logging for antivirus.
scan-mode            Choose between default scan mode and legacy scan mode.
                                            Diagnostics
The following diagnostic commands are meant for troubleshooting only.
diagnose ips av mode ?
    hybrid    Enable/disable hybrid scan mode.
    show      Show status of hybrid scan mode.
                                            To check flow-base AV scan mode status:
diagnose ips av mode show
    Flow-av hybrid scan: Enabled
    Flow-av hybrid scan: Enabled
    Flow-av hybrid scan: Enabled
    Flow-av hybrid scan: Enabled
                                            To disable hybrid scan for flow-base AV and enable full scan:
|  | This command does not persist over a reboot. Flow-av hybrid scan is enabled by default. | 
diagnose ips av mode hybrid disable
diagnose ips av mode show
    Flow-av hybrid scan: Disabled
    Flow-av hybrid scan: Disabled
    Flow-av hybrid scan: Disabled
    Flow-av hybrid scan: Disabled
                                            To enable hybrid scan for flow-base AV and disable full scan to go back to default:
diagnose ips av mode hybrid enable
diagnose ips av mode show
    Flow-av hybrid scan: Enabled
    Flow-av hybrid scan: Enabled
    Flow-av hybrid scan: Enabled
    Flow-av hybrid scan: Enabled
