---
id: collect-261001-huawei/huawei/r-huawei-comments-1bdoc9s-huawei-s5735l48t4xeav2-no-web-interface-c54529c2-2
title: "r-huawei-comments-1bdoc9s-huawei-s5735l48t4xeav2-no-web-interface-c54529c2"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2023-11-02", "2024-03-13"]
keywords: ["agent", "license"]
source: docs/RAG/collect-261001-huawei/r-huawei-comments-1bdoc9s-huawei-s5735l48t4xeav2-no-web-interface-c54529c2.md
source_anchor: ""
source_lines: [7, 306]
sha256: 674398229303c2e2111e5d936ef8fe414eb423a302325cd2d06f7c2ba43df763
---

# r-huawei-comments-1bdoc9s-huawei-s5735l48t4xeav2-no-web-interface-c54529c2

    Hi All
I've got a Huawei S5735-L48T4XE-A-V2.
It is running the following System File & Patch File:
System: S5735-L-V2_V600R022C10SPC500.cc
Patch: S5735-L-V2_V600R022SPH151.PAT
EDIT: For anyone who stumbles on this post,updating to the following firmware solves the problem:
Huawei Solved it with the current Firmware
Release: V600R023C00SPC500 Patch: V600R023SPH120
Just in case you are still looking for it.
Now here's the problem. I cannot enable the Web Interface.
On the underside of the device is a sticker with basic instructions on how to do this.
(Press mode button for more than 6 seconds and then access the switch at the IP 192.168.1.253)
This worked perfectly on previous S5720-28X-LI-AC & S5735-L24T4X-A1 models but does not work at all for the S5735-L48T4XE-A-V2.
After this failed I connected the switch via serial and then manually set up an IP and enabled the web interface. However, it does not give me any configuration settings whatsoever.
I believed it was a user permission level setting but the web user already has the highest privilege level.
What on earth is going wrong or what am I doing wrong ?
Current config file export:
display current-configuration > 1710325530201.cfg
!Software Version V600R022C10SPC500
!Last configuration was updated at 2023-11-02 22:06:30+02:00 by administrator
!Last configuration was saved at 2024-03-13 12:23:46+02:00 by administrator
!md_tlm VRPV800R006C00B016D0127-0.0.1
#
pki realm default
#
language character-set ISO8859-1
#
clock timezone Bucharest add 02:00:00
#
sysname HUAWEI
#
undo ftp server source all-interface
undo ftp ipv6 server source all-interface
#
ssl policy default
 pki-domain default
 ssl minimum version tls1.2
 cipher-suite exclude key-exchange rsa
 cipher-suite exclude cipher mode cbc
 cipher-suite exclude hmac sha1
 diffie-hellman modulus 3072
 ecdh group curve brainpool
 signature algorithm-list ed25519 ed448 rsa-pss-pss-sha256 rsa-pss-pss-sha384 rsa-pss-pss-sha512 rsa-pss-rsae-sha256 rsa-pss-rsae-sha384 rsa-pss-rsae-sha512
#
info-center logfile compression lzma
#
device board 1 board-type S5735-L48T4XE-A-V2
#
authentication-profile name default_authen_profile
authentication-profile name dot1x_authen_profile
authentication-profile name dot1xmac_authen_profile
authentication-profile name mac_authen_profile
#
access-user dot1x-identity speed-limit 60
#
drop-profile default
#
ntp server source-interface all disable
ntp ipv6 server source-interface all disable
#
error-down auto-recovery cause link-flap interval 60
#
undo telnet server-source all-interface
undo telnet ipv6 server-source all-interface
#
mac-address update arp enable
#
qos schedule-profile default
#
diffserv domain default
#
ip vpn-instance _management_vpn_
 ipv4-family
#
aaa
 authentication-scheme default
  authentication-mode local
 authentication-scheme radius
  authentication-mode radius
 authorization-scheme default
  authorization-mode local
 accounting-scheme default
  accounting-mode none
 local-aaa-user password policy administrator
  password expire 999
 domain default
  authentication-scheme default
  accounting-scheme default
 domain default_admin
  authentication-scheme default
  accounting-scheme default
 local-user administrator password irreversible-cipher $1d$4yZl~e[pM))cLb:E$r&wyGm,py9'~(`A;YpVPFYPl<H=;A0=&A<Ilk-"L$
 local-user administrator privilege level 3
 local-user administrator ftp-directory flash:
 local-user administrator service-type telnet terminal ssh ftp http
 local-user mtnadmin password irreversible-cipher $1d$Y$zM/WK7XBskI}G/$_WAO:20!b~NS^<,Gs=12+bKT#FDOJ2N+o;Fv<xR#$
 local-user mtnadmin ftp-directory flash:
 local-user mtnadmin service-type telnet terminal ssh ftp http
 local-user mtnadmin user-group manage-ug
#
free-rule-template name default_free_rule
#
dot1x-access-profile name dot1x_access_profile
#
mac-access-profile name mac_access_profile
#
stack
#
license
#
warranty
#
interface Vlanif1
 ip address 10.0.44.23 255.255.255.0
#
interface Stack-Port1/1
#
interface Stack-Port1/2
#
interface GE1/0/1
#
interface GE1/0/2
#
interface GE1/0/3
#
interface GE1/0/4
#
interface GE1/0/5
#
interface GE1/0/6
#
interface GE1/0/7
#
interface GE1/0/8
#
interface GE1/0/9
#
interface GE1/0/10
#
interface GE1/0/11
#
interface GE1/0/12
#
interface GE1/0/13
#
interface GE1/0/14
#
interface GE1/0/15
#
interface GE1/0/16
#
interface GE1/0/17
#
interface GE1/0/18
#
interface GE1/0/19
#
interface GE1/0/20
#
interface GE1/0/21
#
interface GE1/0/22
#
interface GE1/0/23
#
interface GE1/0/24
#
interface GE1/0/25
#
interface GE1/0/26
#
interface GE1/0/27
#
interface GE1/0/28
#
interface GE1/0/29
#
interface GE1/0/30
#
interface GE1/0/31
#
interface GE1/0/32
#
interface GE1/0/33
#
interface GE1/0/34
#
interface GE1/0/35
#
interface GE1/0/36
#
interface GE1/0/37
#
interface GE1/0/38
#
interface GE1/0/39
#
interface GE1/0/40
#
interface GE1/0/41
#
interface GE1/0/42
#
interface GE1/0/43
#
interface GE1/0/44
#
interface GE1/0/45
#
interface GE1/0/46
#
interface GE1/0/47
#
interface GE1/0/48
#
interface 10GE1/0/1
#
interface 10GE1/0/2
#
interface 10GE1/0/3
#
interface 10GE1/0/4
#
interface 10GE1/0/5
#
interface 10GE1/0/6
#
interface NULL0
#
ip route-static 0.0.0.0 255.255.255.0 10.0.44.1
#
snmp-agent local-engineid 800007DB0348B25DBBEB94
#
snmp-agent sys-info version v3
#
undo snmp-agent protocol source-status all-interface
undo snmp-agent protocol source-status ipv6 all-interface
#
undo snmp-agent proxy protocol source-status all-interface
undo snmp-agent proxy protocol source-status ipv6 all-interface
#
#
ssh server rsa-key min-length 3072
undo ssh authentication-type default password
ssh user administrator
ssh user administrator authentication-type password
ssh user administrator service-type all
ssh user administrator sftp-directory flash:
ssh user mtnadmin
ssh user mtnadmin authentication-type password
ssh user mtnadmin service-type all
ssh user mtnadmin sftp-directory flash:
ssh server-source all-interface
undo ssh ipv6 server-source all-interface
ssh authorization-type default aaa
#
ssh server cipher aes256_gcm aes128_gcm aes256_ctr aes192_ctr aes128_ctr
ssh server hmac sha2_512 sha2_256
ssh server key-exchange dh_group_exchange_sha256 dh_group16_sha512 curve25519_sha256
#
ssh server publickey rsa_sha2_256 rsa_sha2_512
#
ssh server dh-exchange min-len 3072
#
ssh client publickey rsa_sha2_256 rsa_sha2_512
#
ssh client cipher aes256_gcm aes128_gcm aes256_ctr aes192_ctr aes128_ctr
ssh client hmac sha2_512 sha2_256
ssh client key-exchange dh_group_exchange_sha256 dh_group16_sha512 curve25519_sha256
#
user-interface maximum-vty 5
#
user-interface con 0
 authentication-mode password
 set authentication password cipher $1d$k78>-jE]>3JyWU;d$&oBn3)+MF:$WctJ;^p(6)1{t>2K|f2uJ.fF2\E9S$
#
user-interface vty 0 4
 authentication-mode aaa
 user privilege level 3
#
http
#
web-manager enable port 443
web-manager http forward enable
web-manager server-source all-interface
web-manager ipv4 server-source -a 10.0.44.23 vpn-instance public
undo web-manager captcha enable
#
return
Section des commentaires
Hi
We have exactly the same problem with S5735-L48P4XE-A-V2. Have you found a solution in the meantime?
Thx
No, Huawei forum has yielded no answers
Some of their documents even state that only ‘Object’ and ‘System’ are available on the UI but that just doesn’t make sense to me
I have gotten a response on the Huawei Community forum.
"The operating system of S5735-L48T4XE-A-V2 is YunShan OS, while the operating system of S5735-L24T4X-A1 is VRP. So there may be some differences in configuration."
So it appears that there is no traditional web interface like on the S5720 series.
Last Huawei product I will ever purchase. What a steaming pile of shit
Hi, there is currently no web interface on the current Huawei switches. But this is to come.
