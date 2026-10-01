---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-xe-16-6-sec-conn-dmvpn-c7f951df-3
title: "c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-xe-16-6-sec-conn-dmvpn--c7f951df"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-xe-16-6-sec-conn-dmvpn--c7f951df.md
source_anchor: ""
source_lines: [288, 322]
sha256: dba5494d5c019bfa21c8baaee959f829b650dcf4fb54e3ca566f1cbe9e35f3c3
---

# c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-xe-16-6-sec-conn-dmvpn--c7f951df

 										local  ident (addr/mask/prot/port): (172.17.0.11/255.255.255.255/47/0)
											remote ident (addr/mask/prot/port): (172.17.0.12/255.255.255.255/47/0)
											current_peer 172.17.0.12 port 500
     PERMIT, flags={origin_is_acl,}
    #pkts encaps: 0, #pkts encrypt: 0, #pkts digest: 0
    #pkts decaps: 2, #pkts decrypt: 2, #pkts verify: 2
    #pkts compressed: 0, #pkts decompressed: 0
    #pkts not compressed: 0, #pkts compr. failed: 0
    #pkts not decompressed: 0, #pkts decompress failed: 0
    #send errors 0, #recv errors 0
     local crypto endpt.: 172.17.0.11, remote crypto endpt.: 172.17.0.12
     path mtu 1500, ip mtu 1500, ip mtu idb FastEthernet0/0/0
     current outbound spi: 0x38C04B36(952126262)
     inbound esp sas:
      spi: 0xA2EC557(170837335)
        transform: esp-des esp-md5-hmac ,
        in use settings ={Transport, }
        conn id: 5, flow_id: SW:5, crypto map: vpnprof-head-1
        sa timing: remaining key lifetime (k/sec): (4515510/3395)
        IV size: 8 bytes
        replay detection support: Y
        Status: ACTIVE
     inbound ah sas:
     inbound pcp sas:
     outbound esp sas:
      spi: 0x38C04B36(952126262)
        transform: esp-des esp-md5-hmac ,
        in use settings ={Transport, }
        conn id: 6, flow_id: SW:6, crypto map: vpnprof-head-1
        sa timing: remaining key lifetime (k/sec): (4515511/3395)
        IV size: 8 bytes
        replay detection support: Y
        Status: ACTIVE
     outbound ah sas:
     outbound pcp sas:
