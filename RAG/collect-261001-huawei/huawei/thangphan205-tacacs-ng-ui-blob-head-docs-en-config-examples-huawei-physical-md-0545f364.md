---
id: collect-261001-huawei/huawei/thangphan205-tacacs-ng-ui-blob-head-docs-en-config-examples-huawei-physical-md-0545f364
title: "Set default_admin as the default administrative domain for SSH / Console / Telnet logins"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-huawei/thangphan205-tacacs-ng-ui-blob-head-docs-en-config-examples-huawei-physical-md-0545f364.md
source_anchor: ""
source_lines: [1, 140]
sha256: b558031a4f4b5b401e4d91c2c2f4c1970d06dbfc0808d293757a2019212671a1
---

# Set default_admin as the default administrative domain for SSH / Console / Telnet logins

This guide details how to integrate a physical Huawei campus switch (e.g., S5700 / S5720 / S5735 / S6700 series running VRP V200R022) with the tacacs-ng-ui server.
Note
Physical campus switches running VRP V200R022 use a different command syntax compared to virtual/data center CloudEngine switches (CE12800 / eNSP). See Key Differences below for a breakdown.
| Parameter / Feature | Physical Campus Switch (S-Series / V200R022) | Virtual Switch (CloudEngine CE12800 / eNSP) | 
|---|---|---|
| CLI Keyword Syntax | Hyphenated: hwtacacs-server template <name> ,hwtacacs-server authentication ... | Space-separated: hwtacacs server template <name> ,hwtacacs server authentication ... | 
| Domain Binding Command | hwtacacs-server <template_name> underaaa -> domain | hwtacacs server <template_name> underaaa -> domain | 
| HWTACACS Enablement | hwtacacs enable (oftenundo hwtacacs enable by default) | hwtacacs enable | 
| Domain Name Suffix | undo hwtacacs-server user-name domain-included | hwtacacs server user-name domain-excluded | 
| Management Routing | Usually in-band via VLAN interface (e.g., Vlanif101 ) in the global routing table | Usually out-of-band via MEth0/0/0 inside a VRF (vpn-instance __MGMT_VPN__ ) | 
| Command Recording | Requires recording-scheme <name> pluscmd recording-scheme <name> underaaa | Configured directly under recording-scheme inaaa | 
| Default Admin Domain | domain default_admin admin in system-view ensures administrative SSH logins usedefault_admin | domain default_admin | 
| SSH Dynamic Auth | Mandatory: ssh authorization-type default aaa allows TACACS+ SSH users without pre-existingssh user entries | ssh authorization-type default aaa | 
| Accounting Resiliency | accounting start-fail online prevents dropping users if accounting packet exchange encounters latency | Usually default | 
Configure tacacs-ng-ui to support the physical Huawei switch:
Navigate to Hosts -> Add Host:
- Host Name: Huawei (or your switch hostname)
- IP Address: 192.168.1.2 (The switch source IP, e.g.,Vlanif101 )
- Shared Secret Key: <YOUR_TACACS_SECRET_KEY> (e.g.,Netconsole123 )
The default pre-seeded objects in tacacs-ng-ui are ready for Huawei VRP:
- Profiles:
  - tacacs_super_user_profile : Returnspriv-lvl = 15 for servicesshell andh3c_shell .
  - tacacs_read_only_profile : Returnspriv-lvl = 1 for servicesshell andh3c_shell .
- Groups: tacacs_super_user (level 15) andtacacs_read_only (level 1).
- Users: user_admin (member oftacacs_super_user ) anduser_read_only (member oftacacs_read_only ).
- Ruleset: default_ruleset permits both groups and attaches their respective profiles.
Important
Always navigate to TACACS Configs -> Generate Config -> Activate after modifying hosts or rules.
Apply the following commands to your physical switch in system-view:
system-view
hwtacacs enable
Create the template using the hyphenated hwtacacs-server syntax:
hwtacacs-server template tacacs_netadmin
 # TACACS+ Server IP address (Replace with your actual tacacs-ng-ui server IP)
 hwtacacs-server authentication <IP_TACACS_SERVER>
 hwtacacs-server authorization <IP_TACACS_SERVER>
 hwtacacs-server accounting <IP_TACACS_SERVER>
 
 # Source IP of the management VLAN interface (e.g., Vlanif101: 192.168.1.2)
 hwtacacs-server source-ip 192.168.1.2
 
 # Shared Secret Key matching tacacs-ng-ui host configuration
 hwtacacs-server shared-key simple <YOUR_TACACS_SECRET_KEY>
 
 # Strip domain suffix (@domain) when forwarding username to TACACS+
 undo hwtacacs-server user-name domain-included
quit
Configure fallback to local accounts so administrators are never locked out if TACACS+ is unreachable:
aaa
 # 1. Authentication Scheme (TACACS+ primary, Local fallback)
 authentication-scheme tac_auth
  authentication-mode hwtacacs local
 quit
 # 2. Authorization Scheme (TACACS+ primary, Local fallback)
 authorization-scheme tac_author
  authorization-mode hwtacacs local
 quit
 # 3. Accounting Scheme (Send session accounting to TACACS+)
 accounting-scheme tac_acct
  accounting-mode hwtacacs
  accounting start-fail online
 quit
 # 4. Command Recording Scheme (Audit CLI commands executed by users)
 recording-scheme tac_record
  recording-mode hwtacacs tacacs_netadmin
 quit
 cmd recording-scheme tac_record
quit
aaa
 domain default_admin
  authentication-scheme tac_auth
  authorization-scheme tac_author
  accounting-scheme tac_acct
  hwtacacs-server tacacs_netadmin
 quit
quit
# Set default_admin as the default administrative domain for SSH / Console / Telnet logins
domain default_admin admin
On physical Huawei switches, SSH by default only allows users configured with ssh user <username>. You must enable default AAA authorization:
# Allow any valid TACACS+ authenticated user to log in via SSH
ssh authorization-type default aaa
# Apply AAA to remote VTY lines
user-interface vty 0 4
 authentication-mode aaa
 idle-timeout 15 0
 protocol inbound ssh
# Verify Console keeps AAA with local fallback
user-interface con 0
 authentication-mode aaa
system-view
hwtacacs enable
#
hwtacacs-server template tacacs_netadmin
 hwtacacs-server authentication <IP_TACACS_SERVER>
 hwtacacs-server authorization <IP_TACACS_SERVER>
 hwtacacs-server accounting <IP_TACACS_SERVER>
 hwtacacs-server source-ip 192.168.1.2
 hwtacacs-server shared-key simple <YOUR_TACACS_SECRET_KEY>
 undo hwtacacs-server user-name domain-included
#
aaa
 authentication-scheme tac_auth
  authentication-mode hwtacacs local
 #
 authorization-scheme tac_author
  authorization-mode hwtacacs local
 #
 accounting-scheme tac_acct
  accounting-mode hwtacacs
  accounting start-fail online
 #
 recording-scheme tac_record
  recording-mode hwtacacs tacacs_netadmin
 cmd recording-scheme tac_record
 #
 domain default_admin
  authentication-scheme tac_auth
  authorization-scheme tac_author
  accounting-scheme tac_acct
  hwtacacs-server tacacs_netadmin
#
domain default_admin admin
ssh authorization-type default aaa
#
user-interface vty 0 4
 authentication-mode aaa
 idle-timeout 15 0
 protocol inbound ssh
#
return
save
Before logging out of your active session, test connectivity to TACACS+ from the user view (<Huawei>):
test-aaa <username> <password> hwtacacs-template tacacs_netadmin
- Success output: Info: Account test succeeded.
- Failure output: Indicates an incorrect password, mismatched shared key, or network routing/firewall block between 192.168.1.2 and<IP_TACACS_SERVER>:49 .
display hwtacacs-server template tacacs_netadmin
Verify that the server status shows active and packet counters (Request/Reply) increment.
display aaa online-user
Verify that logged-in users show Domain: default_admin and Authen-Mode: HWTACACS.
If tacacs-ng-ui is stopped or unreachable, the switch immediately checks the local user database (local-user admin, local-user pescoadmin), ensuring emergency access is never lost.
