---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-xe-16-6-sec-conn-dmvpn-c7f951df-2
title: "c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-xe-16-6-sec-conn-dmvpn--c7f951df"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-xe-16-6-sec-conn-dmvpn--c7f951df.md
source_anchor: ""
source_lines: [102, 287]
sha256: 023c29b489fb64204c957c35c123d34f00b42571e6363cb1afb1e60807df02fc
---

# c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-xe-16-6-sec-conn-dmvpn--c7f951df

Spoke1# show crypto ipsec sa
interface: Tunnel 0
			Crypto map tag: vpnprof-head-1, local addr 172.17.0.11
   protected vrf: (none)
   								local  ident (addr/mask/prot/port): (172.17.0.11/255.255.255.255/47/0)
											remote ident (addr/mask/prot/port): (172.17.0.1/255.255.255.255/47/0)
											current_peer 172.17.0.1 port 500
     PERMIT, flags={origin_is_acl,}
    #pkts encaps: 134, #pkts encrypt: 134, #pkts digest: 134
    #pkts decaps: 118, #pkts decrypt: 118, #pkts verify: 118
    #pkts compressed: 0, #pkts decompressed: 0
    #pkts not compressed: 0, #pkts compr. failed: 0
    #pkts not decompressed: 0, #pkts decompress failed: 0
    #send errors 22, #recv errors 0
     local crypto endpt.: 172.17.0.11, remote crypto endpt.: 172.17.0.1
     path mtu 1500, ip mtu 1500, ip mtu idb FastEthernet0/0/0
     current outbound spi: 0xA75421B1(2807308721)
     inbound esp sas:
      spi: 0x96185188(2518176136)
        transform: esp-des esp-md5-hmac ,
        in use settings ={Transport, }
        conn id: 3, flow_id: SW:3, crypto map: vpnprof-head-1
        sa timing: remaining key lifetime (k/sec): (4569747/3242)
        IV size: 8 bytes
        replay detection support: Y
        Status: ACTIVE
     inbound ah sas:
     inbound pcp sas:
     outbound esp sas:
      spi: 0xA75421B1(2807308721)
        transform: esp-des esp-md5-hmac ,
        in use settings ={Transport, }
        conn id: 4, flow_id: SW:4, crypto map: vpnprof-head-1
        sa timing: remaining key lifetime (k/sec): (4569745/3242)
        IV size: 8 bytes
        replay detection support: Y
        Status: ACTIVE
     outbound ah sas:
     outbound pcp sas:
   protected vrf: (none)
   								local  ident (addr/mask/prot/port): (172.17.0.11/255.255.255.255/47/0)	
											remote ident (addr/mask/prot/port): (172.17.0.5/255.255.255.255/47/0)
											current_peer 172.17.0.5 port 500
     PERMIT, flags={origin_is_acl,}
    #pkts encaps: 244, #pkts encrypt: 244, #pkts digest: 244
    #pkts decaps: 253, #pkts decrypt: 253, #pkts verify: 253
    #pkts compressed: 0, #pkts decompressed: 0
    #pkts not compressed: 0, #pkts compr. failed: 0
    #pkts not decompressed: 0, #pkts decompress failed: 0
    #send errors 1, #recv errors 0
     local crypto endpt.: 172.17.0.11, remote crypto endpt.: 172.17.0.5
     path mtu 1500, ip mtu 1500, ip mtu idb FastEthernet0/0/0
     current outbound spi: 0x3C50B3AB(1011921835)
     inbound esp sas:
      spi: 0x3EBE84EF(1052673263)
        transform: esp-des esp-md5-hmac ,
        in use settings ={Transport, }
        conn id: 1, flow_id: SW:1, crypto map: vpnprof-head-1
        sa timing: remaining key lifetime (k/sec): (4549326/2779)
        IV size: 8 bytes
        replay detection support: Y
        Status: ACTIVE
     inbound ah sas:
     inbound pcp sas:
     outbound esp sas:
      spi: 0x3C50B3AB(1011921835)
        transform: esp-des esp-md5-hmac ,
        in use settings ={Transport, }
        conn id: 2, flow_id: SW:2, crypto map: vpnprof-head-1
        sa timing: remaining key lifetime (k/sec): (4549327/2779)
        IV size: 8 bytes
        replay detection support: Y
        Status: ACTIVE
     outbound ah sas:
     outbound pcp sas:
   protected vrf: (none)
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
								interface: Tunnel 1
								Crypto map tag: vpnprof-head-1, local addr 172.17.0.11
   protected vrf: (none)
   								local  ident (addr/mask/prot/port): (172.17.0.11/255.255.255.255/47/0)
											remote ident (addr/mask/prot/port): (172.17.0.1/255.255.255.255/47/0)
											current_peer 172.17.0.1 port 500
     PERMIT, flags={origin_is_acl,}
    #pkts encaps: 134, #pkts encrypt: 134, #pkts digest: 134
    #pkts decaps: 118, #pkts decrypt: 118, #pkts verify: 118
    #pkts compressed: 0, #pkts decompressed: 0
    #pkts not compressed: 0, #pkts compr. failed: 0
    #pkts not decompressed: 0, #pkts decompress failed: 0
    #send errors 22, #recv errors 0
     local crypto endpt.: 172.17.0.11, remote crypto endpt.: 172.17.0.1
     path mtu 1500, ip mtu 1500, ip mtu idb FastEthernet0/0/0
     current outbound spi: 0xA75421B1(2807308721)
     inbound esp sas:
      spi: 0x96185188(2518176136)
        transform: esp-des esp-md5-hmac ,
        in use settings ={Transport, }
        conn id: 3, flow_id: SW:3, crypto map: vpnprof-head-1
        sa timing: remaining key lifetime (k/sec): (4569747/3242)
        IV size: 8 bytes
        replay detection support: Y
        Status: ACTIVE
     inbound ah sas:
     inbound pcp sas:
     outbound esp sas:
      spi: 0xA75421B1(2807308721)
        transform: esp-des esp-md5-hmac ,
        in use settings ={Transport, }
        conn id: 4, flow_id: SW:4, crypto map: vpnprof-head-1
        sa timing: remaining key lifetime (k/sec): (4569745/3242)
        IV size: 8 bytes
        replay detection support: Y
        Status: ACTIVE
     outbound ah sas:
     outbound pcp sas:
   protected vrf: (none)
   								local  ident (addr/mask/prot/port): (172.17.0.11/255.255.255.255/47/0)
											remote ident (addr/mask/prot/port): (172.17.0.5/255.255.255.255/47/0)
											current_peer 172.17.0.5 port 500
     PERMIT, flags={origin_is_acl,}
    #pkts encaps: 244, #pkts encrypt: 244, #pkts digest: 244
    #pkts decaps: 253, #pkts decrypt: 253, #pkts verify: 253
    #pkts compressed: 0, #pkts decompressed: 0
    #pkts not compressed: 0, #pkts compr. failed: 0
    #pkts not decompressed: 0, #pkts decompress failed: 0
    #send errors 1, #recv errors 0
     local crypto endpt.: 172.17.0.11, remote crypto endpt.: 172.17.0.5
     path mtu 1500, ip mtu 1500, ip mtu idb FastEthernet0/0/0
     current outbound spi: 0x3C50B3AB(1011921835)
     inbound esp sas:
      spi: 0x3EBE84EF(1052673263)
        transform: esp-des esp-md5-hmac ,
        in use settings ={Transport, }
        conn id: 1, flow_id: SW:1, crypto map: vpnprof-head-1
        sa timing: remaining key lifetime (k/sec): (4549326/2779)
        IV size: 8 bytes
        replay detection support: Y
        Status: ACTIVE
     inbound ah sas:
     inbound pcp sas:
     outbound esp sas:
      spi: 0x3C50B3AB(1011921835)
        transform: esp-des esp-md5-hmac ,
        in use settings ={Transport, }
        conn id: 2, flow_id: SW:2, crypto map: vpnprof-head-1
        sa timing: remaining key lifetime (k/sec): (4549327/2779)
        IV size: 8 bytes
        replay detection support: Y
        Status: ACTIVE
     outbound ah sas:
     outbound pcp sas:
   protected vrf: (none)
