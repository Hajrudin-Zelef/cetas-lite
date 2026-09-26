---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-54-2
title: "Description"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-54.md
source_anchor: ""
source_lines: [84, 123]
sha256: 3e3c0ccd5819553a154365eec4ed17971ae10da9bdecf484d08f1c1ab1716116
---

# Description

| Feature | Clarification of which menus become unavailable to change | 
| bandwidth-test | */tool bandwidth-test**/tool bandwidth-server* */tool speed-test* | 
| routerboard | */system routerboard settings* (except auto-upgrade option) | 
| container | all container functionality | 
| install-any-version | RouterOS will no longer allow for you to install RouterOS version below versions listed under "allowed-versions" attribute. | 
|  | */tool e-mail*  | 
| fetch | */tool fetch* | 
| hotspot | */ip hotspot* | 
| ipsec | */ip ipsec* | 
| l2tp | */interface l2tp-server* */interface l2tp-client* | 
| partitions | */partitions*  does not allow to change count of partitions. If your router is unable to boot, it will still be able to boot into your other partitions. No restriction for crash recovery. | 
| pptp | */interface pptp-server* */interface pptp-client* | 
| proxy | */ip proxy* | 
| romon | */tool romon* | 
| scheduler | */system scheduler* | 
| smb | */ip smb* | 
| sniffer | */tool sniffer* | 
| socks | */ip socks* | 
| traffic-gen | */tool traffic-generator* */tool flood-ping* */tool ping-speed* | 
| zerotier | */zerotier* | 

# Allowed versions

Device mode lists in its parameters an argument called "allowed-versions". This is a list of versions which MikroTik considers as secure and which ones do not include any serious vulnerabilities which could be used by an attacker. 

This setting does not depend on the installed RouterOS version and works as a separate protection layer, in order to disallow attacker to downgrade version step-by-step to reach any known vulnerable RouterOS release. When you upgrade RouterOS to a release where a newer "allowed-versions" list is available, oldest list will be overwritten. If you downgrade RouterOS, "allowed-versions" list will not change and will remain updated to the latest list. The list is ignored, if device-mode "install-any-version" is enabled.

# Flagged status

Along with the device-mode feature, RouterOS now can analyze the whole configuration at system startup, to determine if there are any signs of unauthorized access to your router. If suspicious configuration is detected, the suspicious configuration will be disabled and the **flagged** parameter will be set to "yes". The device has now a Flagged state and enforces certain limitations. 

If the system has this flagged status, the current configuration works, but it is not possible to perform the following actions: 

bandwidth-test, traffic-generator, sniffer, as well as configuration actions that enable or create new configuration entries (it will still be possible to disable or delete them) for the following programs: *system scheduler, SOCKS proxy, pptp, l2tp, ipsec, proxy, smb*.

When performing the aforementioned actions while the router has the flagged state, you will receive an error message:

To exit the flagged state, you must perform the command "/system/device-mode/update flagged=no". The system will ask to either press a button, or issue a hard reboot (cut power physically or do a hard reboot of the virtual machine). **Important!** Although the system has disabled any malicious looking rules, which triggered the flagged state, it is crucial to inspect all of your configuration for other unknown things, before exiting the flagged state. If your system has been flagged, assume that your system has been compromised and do a full audit of all settings before re-enabling the system for use. After completing the audit, change all the system passwords and upgrade to the latest RouterOS version. 

Starting from RouterOS version 7.17 device-mode restricts SwOS/RouterOS transition for dual-boot; in order to enable: system/device-mode/update routerboard=yes
