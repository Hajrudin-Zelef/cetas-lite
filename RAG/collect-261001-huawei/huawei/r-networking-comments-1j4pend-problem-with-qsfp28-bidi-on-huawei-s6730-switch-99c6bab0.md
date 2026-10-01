---
id: collect-261001-huawei/huawei/r-networking-comments-1j4pend-problem-with-qsfp28-bidi-on-huawei-s6730-switch-99c6bab0
title: "r-networking-comments-1j4pend-problem-with-qsfp28-bidi-on-huawei-s6730-switch-99c6bab0"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["wavelength"]
source: docs/RAG/collect-261001-huawei/r-networking-comments-1j4pend-problem-with-qsfp28-bidi-on-huawei-s6730-switch-99c6bab0.md
source_anchor: ""
source_lines: [1, 18]
sha256: 7d1b3c8d03b2c0a1657fa1f094b8a204d36b8b80a083365f34051bd63d1587db
---

# r-networking-comments-1j4pend-problem-with-qsfp28-bidi-on-huawei-s6730-switch-99c6bab0

Problem with QSFP28 BIDI on Huawei S6730 Switch 
        
        
        
    
    
    Hello, i have a problem with running HUAWEI QSFP28 100G BIDI on a HUAWEI S6730 Cloud Engine. Patch Version is V600R024HP0021 The Bidi is correctly displayed in the switch: 100GE1/0/4 transceiver information:
Common information: Transceiver Type :100GBASE_LR4 Connector Type :LC Wavelength (nm) :1309 Transfer Distance (m) :30000(9um/125um SMF) Digital Diagnostic Monitoring :YES Vendor Name :HUAWEI Vendor Part Number :02311KNU Ordering Name :
Manufacture information: Manu. Serial Number :G4O2022623 Manufacturing Date :2016-3-23 Vendor Name :HUAWEI
Alarm information:
Warning information:
Diagnostic information: Temperature (Celsius) :28.99 Voltage (V) :3.41 Bias Current (mA) :0.00|0.00 (Lane0|Lane1) 0.00|0.00 (Lane2|Lane3) Bias High Threshold (mA) :120.00 Bias Low Threshold (mA) :5.00 Current RX Power (dBm) :-40.00|-40.00(Lane0|Lane1) -40.00|-40.00(Lane2|Lane3) Default RX Power High Threshold (dBm) :-2.50 Default RX Power Low Threshold (dBm) :-16.00 Current TX Power (dBm) :-40.00|-40.00(Lane0|Lane1) -40.00|-40.00(Lane2|Lane3) Default TX Power High Threshold (dBm) :7.00 Default TX Power Low Threshold (dBm) :0.00
Following config on the port, but also tested with default settings: <bh-s6730-iscsi-1-rz1>display current-configuration interface 100GE1/0/1
interface 100GE1/0/1 port link-type access device transceiver 100GBASE-FIBER fec mode none
return I noticed, that there is no light in the bidi, as when i plug the bidi into a HPE switch, i can see the laser. Does anyone have an idea how to troubleshoot this issue or what could be the problem? Thank you in advance!
Section des commentaires
Soyez la première personne à commenter
Personne n’a encore répondu à cette publication. Partage ton avis pour lancer la conversation.
