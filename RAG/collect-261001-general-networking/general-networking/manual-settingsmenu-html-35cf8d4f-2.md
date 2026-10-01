---
id: collect-261001-general-networking/general-networking/manual-settingsmenu-html-35cf8d4f-2
title: "manual-settingsmenu-html-35cf8d4f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/manual-settingsmenu-html-35cf8d4f.md
source_anchor: ""
source_lines: [84, 149]
sha256: 67a68292939c0ed0a05b576ad8fa33f917fcc399f45af7f67f9ec63a9caea104
---

# manual-settingsmenu-html-35cf8d4f

| Server | Select one or more authentication servers to validate user credentials against. Multiple servers can make sense with remote authentication methods to provide a fallback during connectivity issues. When nothing is specified the default of “Local Database” is used. | 
| Disable integrated authentication | When set, console login, SSH, and other system services can only use standard UNIX account authentication. | 
| Sudo | Permit sudo usage for administrators with shell access. | 
| User OTP seed | Select groups which are allowed to generate their own OTP seed on the password page. | 
Cron
Cron is a service that is used to execute jobs periodically. Cron jobs can be viewed by navigating to
. New jobs can be added by click the + button in the lower right
corner.
When adding a new job or modifying an existing one, you will be presented with fields that directly reflect the cron file syntax and that mostly speak for themselves. A job needs a name, a command, command parameters (if applicable), a description (optional, but recommend) and most importantly, a schedule. All time-related fields share the same syntax:
- An asterisk (*) can be used to mean “any”
- Specifying multiple values is possible using the comma: 1,4,9
- Ranges can be specified using a dash: 4-9
Available cron jobs are registered in the backend to prevent command injection and privilege escalation. These can be found under Command and may use optional Parameters. Restart and reload actions are self-explanatory. They tend to take no parameters and will restart (usually slower stop and start of a process) or reload (usually a faster SIGHUP) the respective service. The availability of restart and reload is subject to their respective services as not all software will support a reload for implementational reasons.
The most common core commands are as follows:
| Command in GUI | Command in shell | Supported parameters | Background information | 
|---|---|---|---|
| Automatic firmware update | configctl firmware auto-update | N/A | Perform a minor update if applicable. | 
| Download and reload external proxy ACLs | configctl proxy fetchacls | N/A | Fetch and activate the external ACL files for configured blocklists. | 
| Firmware changelog update | configctl firmware changelog cron | N/A | Refresh current changelog status from authoritative firmware location to preview changelogs for new versions. Note this utilizes a skew interval of 25 minutes and is also performed by the firmware update check. | 
| Firmware update check | configctl firmware poll | N/A | Refresh current update status from firmware mirror for e.g. remote status check via API. Note this utilizes a skew interval of 25 minutes. | 
| HA update and reconfigure backup | configctl system ha_reconfigure_backup | N/A | Synchronize the configuration to the backup firewall and restart its services to apply the changes. | 
| Halt and power off the system | configctl system halt | N/A | Perform a power off at the specified time. | 
| Manual gateway switch | configctl interface routes alarm | N/A | Perform a manual gateway switch if applicable. Malfunctioning gateway monitors will be restarted as well | 
| Periodic interface reset | configctl interface reconfigure [identifier] | identifier: Internal name of the interface as shown in assignments or overview page, e.g. “lan”, “wan”, “optX”. | Cycle through an interface reset that removes all connectivity and reactivates it cleanly. | 
| Reboot the system | configctl system reboot | N/A | Perform a reboot at the specified time. | 
| Remote backup | configctl system remote backup | N/A | Trigger the remote backup at the specified time as opposed to its nightly default. | 
| Restart OpenVPN instance | configctl openvpn restart | instance: UUID | Restart the given OpenVPN instance. | 
| Update and reload firewall aliases | configctl filter refresh_aliases | N/A | Updates IP aliases for DNS entries and MAC addresses as well as URL tables. | 
| Update and reload intrusion detection rules | configctl ids update | N/A | Fetches remote rules and reloads the IDS instance to make use of newly fetched rules. | 
| Update Unbound DNSBLs | configctl unbound dnsbl | N/A | Update the the DNS blocklists and apply the changes to Unbound. | 
| ZFS pool trim | configctl zfs trim [pool] | pool: ZFS pool name to perform the action on | Initiates an immediate on-demand TRIM operation for all of the free space in a pool. This operation informs the underlying storage devices of all blocks in the pool which are no longer allocated and allows thinly provisioned devices to reclaim the space. | 
| ZFS pool scrub | configctl zfs scrub [pool] | pool: ZFS pool name to perform the action on | Begins a scrub or resumes a paused scrub. The scrub examines all data in the specified pools to verify that it checksums correctly. For replicated (mirror, raidz, or draid) devices, ZFS automatically repairs any damage discovered during the scrub. | 
General
The general settings mainly concern network-related settings like the hostname. The general setting can be set by going to . The following settings are available:
| Option | Description | 
|---|---|
| System |  | 
| Hostname | Hostname without domain, e.g.: firewall | 
| Domain | The domain, e.g. mycorp.com ,home ,office ,private , etc. Do not use ‘local’ as a domain name. It will cause local hosts running mDNS (avahi, bonjour, etc.) to be unable to resolve local hosts not running mDNS. | 
| Time zone | Set the time zone closest to you. | 
| Language | Default language. Can be overridden by users. | 
| Theme | More themes can be installed via plug-ins. | 
| Picture | Upload a picture for display in the Picture widget on the dashboard. The maximum file size is 10MB. Pictures are scaled automatically to the widget size. | 
| Networking |  | 
| Prefer to use IPv4 even if IPv6 is available | By default if a hostname resolves IPv6 and IPv4 addresses, the IPv6 will be used. If checked, then IPv4 addresses will be used instead of IPv6. | 
| DNS servers | A list of DNS servers, optionally with a gateway. These DNS servers are also used for the DHCP service, DNS services and for PPTP VPN clients. When using multiple WAN connections there should be at least one unique DNS server per gateway. | 
| Allow DNS server list to be overridden by DHCP/PPP on WAN | If this option is set, DNS servers assigned by a DHCP/PPP server on the WAN will be used for their own purposes (including the DNS services). However, they will not be assigned to DHCP and PPTP VPN clients. | 
| Do not use the local DNS service as a nameserver for this system | When enabling local DNS services such as Dnsmasq and Unbound, OPNsense will use these as a nameserver. Check this option to prevent this. | 
| Allow default gateway switching | If the link where the default gateway resides fails switch the default gateway to another available one. | 
Tunables
Tunables are the settings that go into the loader.conf and sysctl.conf files, which allows tweaking of low-level system
settings. They can be set by going to .
Here, the currently active settings can be viewed and new ones can be created.
A list of possible values can be obtained by issuing sysctl -a on an OPNsense shell.
Additional tunables may exist depending on boot loader capabilities and kernel module support.
Miscellaneous
As the name implies, this section contains the settings that do not fit anywhere else.
| Option | Description | 
|---|---|
| Cryptography settings |  | 
| Hardware acceleration | Select your method of hardware acceleration, if present. Check the full help for hardware-specific advice. | 
| Thermal Sensors |  | 
| Hardware | Select between No/ACPI thermal sensor driver and processor-specific drivers. | 
| Periodic Backups |  | 
| Periodic RRD Backup | Periodically backup Round Robin Database. | 
| Periodic DHCP Leases Backup | Periodically backup DHCP leases. | 
