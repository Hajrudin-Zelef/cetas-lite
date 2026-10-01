---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-1-administration-guide-477578-ztna-ip-mac-based-access-co-4cb62d33-3
title: "ZTNA IP MAC based access control example"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2023-05-10"]
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-1-administration-guide-477578-ztna-ip-mac-based-access-co-4cb62d33.md
source_anchor: ""
source_lines: [291, 337]
sha256: f54660391e50ef108eb668d15a7fd55c2cc31aa635d3b6dfde77c3cb6a74225a
---

# diagnose wad dev query-by uid 9A016B5A6E914B42AD4168C066EB04CA FCTEMS8822001975 00000000000000000000000000000000
Attr of type=0, length=83, value(ascii)=9A016B5A6E914B42AD4168C066EB04CA
Attr of type=4, length=0, value(ascii)=
Attr of type=6, length=1, value(ascii)=true
Attr of type=5, length=40, value(ascii)=2B8D4FF0E71FE7E064288FE1B4F87E25232092D0
Attr of type=3, length=66, value(ascii)=ZTNA_Domain-Users_FCTEMS882200197500000000000000000000000000000000
Attr of type=3, length=68, value(ascii)=ZTNA_Remote-Allowed_FCTEMS882200197500000000000000000000000000000000
Attr of type=3, length=83, value(ascii)=ZTNA_Group-Membership-Domain-Users_FCTEMS882200197500000000000000000000000000000000
Attr of type=3, length=59, value(ascii)=CLASS_High_FCTEMS882200197500000000000000000000000000000000
Attr of type=3, length=77, value(ascii)=**ZTNA_Malicious-File-Detected**_FCTEMS882200197500000000000000000000000000000000
Attr of type=3, length=61, value(ascii)=CLASS_Remote_FCTEMS882200197500000000000000000000000000000000
Attr of type=3, length=76, value(ascii)=ZTNA_all_registered_clients_FCTEMS882200197500000000000000000000000000000000

# diagnose firewall dynamic list
 List all dynamic addresses:
 ...
CMDB name: EMS1_ZTNA_Malicious-File-Detected
TAG name: **Malicious-File-Detected**
EMS1_ZTNA_Malicious-File-Detected: ID(205)
        RANGE(10.0.1.0-10.0.0.255)
        **ADDR(10.0.1.2)**
Total IP dynamic range blocks: 1.
Total IP dynamic addresses: 0.

# diagnose test application fcnacd 7
Entry #1:
...
State:   sysinfo:1, tag:1, tagsz:1, out-of-sync:0
Owner:   
Cert SN: 2B8D4FF0E71FE7E064288FE1B4F87E25232092D0
online:  Yes
Route IP:**10.0.1.2**
vfid:    0
has more:No
Tags:
idx:0, ttdl:1   name:Domain-Users
idx:1, ttdl:1   name:Remote-Allowed
idx:2, ttdl:1   name:Group-Membership-Domain-Users
idx:3, ttdl:2   name:High
idx:4, ttdl:1   name:**Malicious-File-Detected**
idx:5, ttdl:2   name:Remote
idx:6, ttdl:1   name:all_registered_clients

# execute log filter field srcip 10.0.1.2
# execute log display 
 
1: date=2023-05-10 time=23:37:02 eventtime=1683787022146761572 tz="-0700" logid="0000000013" type="traffic" subtype="forward" level="notice" vd="root" **srcip=10.0.1.2** srcport=14609 srcintf="port1" srcintfrole="undefined" **dstip=10.88.0.3 dstport=9443** dstintf="port2" dstintfrole="dmz" srcuuid="b458a65a-f759-51ea-d7df-ef2e750026d1" dstuuid="592dfb72-0775-51ec-aa79-94bd9894388c" srccountry="Reserved" dstcountry="Reserved" sessionid=177409 proto=6 **action="deny" policyid=10** policytype="policy" poluuid="92938512-ef9a-51ed-6a39-bafb9147e9aa" **policyname="block-internal-malicious-access"** service="tcp/9443" trandisp="noop" duration=0 sentbyte=0 rcvdbyte=0 sentpkt=0 rcvdpkt=0 appcat="unscanned" crscore=30 craction=131072 crlevel="high"
