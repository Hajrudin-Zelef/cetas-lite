---
id: collect-261001-huawei/huawei/knowledge-base-design-and-configure-access-list-example-huawei-f2613854
title: "knowledge-base-design-and-configure-access-list-example-huawei-f2613854"
domain: huawei
role: reference
task: reference
actors: ["Huawei", "United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/knowledge-base-design-and-configure-access-list-example-huawei-f2613854.md
source_anchor: ""
source_lines: [1, 41]
sha256: 66643d85dcc498482322056c5724baa1a98bb3eadaf66ac22d098dfebea595df
---

# knowledge-base-design-and-configure-access-list-example-huawei-f2613854

Poland
GRANDMETRIC Sp. z o.o.
ul. Metalowa 5, 60-118 Poznań, Poland
NIP 7792433527
+48 61 271 04 43
info@grandmetric.com
UK
Grandmetric LTD
Office 584b
182-184 High Street North
London
E6 2JA
+44 20 3321 5276
info@grandmetric.com
US Region
Grandmetric LLC
Lewes DE 19958
16192 Coastal Hwy USA
EIN: 98-1615498
+1 302 691 94 10 
info@grandmetric.com
Technology: Network Security
Area: ACL
Vendor: Huawei
Software: eNSP
Platform: Huawei switches and routers
An Access Control List (ACL) is a set of rules that classify packets to filter them. An ACL can contain multiple rules. A rule is identified by a rule ID, which can be set by a user or automatically generated based on the ACL step. All rules in an ACL are arranged in ascending order of rule IDs.
To configure ACL, use the following command:
<HuaweiI>system-view
[Huawei]interface GigabitEtherent 0/0/1
[Huawei-interface GigabitEtherent 0/0/1]ip address 10.0.4.1 24
[Huawei-interface GigabitEtherent 0/0/0]quit
[[Huawei] acl 3000
[[Huawei-acl-adv-3000]rule 5 permit tcp source 10.0.4.254 0.0.0.0 destination 119.84.111.3 0.0.0.0 destination-port eq 23
[[Huawei-acl-adv-3000]rule 10 permit ip source 10.0.4.0 0.0.0.0 destination any
[Huawei-acl-adv-3000]rule 15 deny ip[[Huawei-acl-adv-3000]]quit
[Huawei]display acl 3000  
Total quantity of nonempty ACL number is 2
Advanced ACL 3000, 1 rule
Acl's step is 5
rule 15 deny ip source 10.164.2.0 0.0.0.255 destination 119.84.111.3 0 time-range satime(Active)
