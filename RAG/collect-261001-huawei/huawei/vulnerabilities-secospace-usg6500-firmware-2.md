---
id: collect-261001-huawei/huawei/vulnerabilities-secospace-usg6500-firmware-2
title: "vulnerabilities-secospace-usg6500-firmware"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["exploit", "memory"]
source: docs/RAG/collect-261001-huawei/vulnerabilities-secospace-usg6500-firmware.md
source_anchor: ""
source_lines: [50, 78]
sha256: d19dadbc2c172d217072f3eb7ca2451ae41039ad869f5368b3477488135448d8
---

# vulnerabilities-secospace-usg6500-firmware

        | There is a weak secure algorithm vulnerability in Huawei products. A weak secure algorithm is used in a module. Attackers can exploit this vulnerability by capturing and analyzing the messages between devices to obtain information. This can lead to information leak.Affected product versions include: IPS Module V500R005C00SPC100, V500R005C00SPC200; NGFW Module V500R005C00SPC100, V500R005C00SPC200; Secospace USG6300 V500R001C30SPC200, V500R001C30SPC600, V500R001C60SPC500, V500R005C00SPC100, V500R005C00SPC200; Secospace USG6500 V500R001C30SPC200, V500R001C30SPC600, V500R001C60SPC500, V500R005C00SPC100, V500R005C00SPC200; Secospace USG6600 V500R001C30SPC200, V500R001C30SPC600, V500R001C60SPC500, V500R005C00SPC100, V500R005C00SPC200; USG9500 V500R001C30SPC200, V500R001C30SPC600, V500R001C60SPC500, V500R005C00SPC100, V500R005C00SPC200. | 
    
    
    
        | There is a memory leak vulnerability in Huawei products. A resource management weakness exists in a module. Attackers with high privilege can exploit this vulnerability by performing some operations. This can lead to memory leak. Affected product versions include:IPS Module V500R005C00SPC100,V500R005C00SPC200;NGFW Module V500R005C00SPC100,V500R005C00SPC200;NIP6300 V500R005C00SPC100,V500R005C10SPC200;NIP6600 V500R005C00SPC100,V500R005C00SPC200;Secospace USG6300 V500R005C00SPC100,V500R005C00SPC200;Secospace USG6500 V500R005C00SPC100,V500R005C10SPC200;Secospace USG6600 V500R005C00SPC100,V500R005C00SPC200. | 
    
    
    
        | There is a use-after-free vulnerability in a Huawei product. A module cannot deal with specific operations in special scenarios. Attackers can exploit this vulnerability by performing malicious operations. This can cause memory use-after-free, compromising normal service. Affected product include some versions of NIP6300, NIP6600, NIP6800, S1700, S2700, S5700, S6700 , S7700, S9700, Secospace USG6300, Secospace USG6500, Secospace USG6600 and USG9500. | 
    
    
    
        | There is a denial of service vulnerability in Huawei products. A module cannot deal with specific messages correctly. Attackers can exploit this vulnerability by sending malicious messages to an affected module. This can lead to denial of service. Affected product include some versions of IPS Module, NGFW Module, NIP6600, NIP6800, Secospace USG6300, Secospace USG6500 and Secospace USG6600. | 
    
    
    
        | There is a memory leak vulnerability in some Huawei products. An authenticated remote attacker may exploit this vulnerability by sending specific message to the affected product. Due to not release the allocated memory properly, successful exploit may cause some service abnormal. Affected product include some versions of IPS Module, NGFW Module, Secospace USG6300, Secospace USG6500, Secospace USG6600 and USG9500. | 
    
    
    
        | There is an information leakage vulnerability in some huawei products. Due to the properly storage of specific information in the log file, the attacker can obtain the information when a user logs in to the device. Successful exploit may cause an information leak. Affected product versions include: NIP6300 versions V500R001C00,V500R001C20,V500R001C30;NIP6600 versions V500R001C00,V500R001C20,V500R001C30;Secospace USG6300 versions V500R001C00,V500R001C20,V500R001C30;Secospace USG6500 versions V500R001C00,V500R001C20,V500R001C30;Secospace USG6600 versions V500R001C00,V500R001C20,V500R001C30,V500R001C50,V500R001C60,V500R001C80;USG9500 versions V500R005C00,V500R005C10. | 
    
    
    
        | There is a denial of service vulnerability in some huawei products. In specific scenarios, due to the improper handling of the packets, an attacker may craft many specific packets. Successful exploit may cause some services to be abnormal. Affected products include some versions of NGFW Module, NIP6300, NIP6600, NIP6800, Secospace USG6300, Secospace USG6500, Secospace USG6600 and SG9500. | 
    
    
    
        | Some Huawei products have a command injection vulnerability. Due to insufficient input validation, an attacker with high privilege may inject some malicious codes in some files of the affected products. Successful exploit may cause command injection.Affected product versions include:NIP6300 versions V500R001C30,V500R001C60;NIP6600 versions V500R001C30,V500R001C60;Secospace USG6300 versions V500R001C30,V500R001C60;Secospace USG6500 versions V500R001C30,V500R001C60;Secospace USG6600 versions V500R001C30,V500R001C60;USG9500 versions V500R001C30,V500R001C60. |
