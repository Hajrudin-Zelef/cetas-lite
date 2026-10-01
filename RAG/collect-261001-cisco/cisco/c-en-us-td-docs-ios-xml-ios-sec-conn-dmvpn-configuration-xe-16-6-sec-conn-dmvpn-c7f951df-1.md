---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-xe-16-6-sec-conn-dmvpn-c7f951df-1
title: "c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-xe-16-6-sec-conn-dmvpn--c7f951df"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-xe-16-6-sec-conn-dmvpn--c7f951df.md
source_anchor: ""
source_lines: [1, 101]
sha256: ed9b0161bd0d28243898727dcbc3cc25f0aa94845eae09037ac5a547eb5ea3d8
---

# c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-xe-16-6-sec-conn-dmvpn--c7f951df

SPOKE 1 has the following results for its DMVPN configuration: 
                                    		
                                 
 
                                 		
Spoke1# show ip nhrp
10.0.0.1/32 via 10.0.0.1, Tunnel 0 created 00:06:52, never expire
  Type: static, Flags: used
  NBMA address: 172.17.0.1
10.0.0.12/32 via 10.0.0.12, Tunnel 0 created 00:03:17, expire 00:01:52
  Type: dynamic, Flags: router
  NBMA address: 172.17.0.12
10.0.1.1/32 via 10.0.1.1, Tunnel 1 created 00:13:45, never expire
  Type: static, Flags: used
  NBMA address: 172.17.0.5
10.0.1.12/32 via 10.0.1.12, Tunnel 1 created 00:00:02, expire 00:04:57
  Type: dynamic, Flags: router
  NBMA address: 172.17.0.12
Spoke1# show crypto socket
 
                                 		
                                    
                                       | Note |   There are only three crypto connections (172.17.0.12, 172.17.0.5 and 172.17.0.1). The two NHRP sessions (10.0.0.12, Tunnel                                                 0) and (10.0.1.12, Tunnel 1) represent the same IPsec session because they both have the same nonbroadcast multiaccess (NBMA)                                                 IPsec peer address.                                                  		                                                | 
                                 
Number of Crypto Socket connections 3
   Shd Peers (local/remote): 172.17.0.11
/172.17.0.12
       Local Ident  (addr/mask/port/prot): (172.17.0.11/255.255.255.255/0/47)
       Remote Ident (addr/mask/port/prot): (172.17.0.12/255.255.255.255/0/47)
       Flags: shared
       ipsec Profile: "vpnprof"
       Socket State: Open
       Client: "TUNNEL SEC" (Client State: Active)
   Shd Peers (local/remote): 172.17.0.11
/172.17.0.5
       Local Ident  (addr/mask/port/prot): (172.17.0.11/255.255.255.255/0/47)
       Remote Ident (addr/mask/port/prot): (172.17.0.5/255.255.255.255/0/47)
       Flags: shared
       ipsec Profile: "vpnprof"
       Socket State: Open
       Client: "TUNNEL SEC" (Client State: Active)
   Shd Peers (local/remote): 172.17.0.11
/172.17.0.1
       Local Ident  (addr/mask/port/prot): (172.17.0.11/255.255.255.255/0/47)
       Remote Ident (addr/mask/port/prot): (172.17.0.1/255.255.255.255/0/47)
       Flags: shared
       ipsec Profile: "vpnprof"
       Socket State: Open
       Client: "TUNNEL SEC" (Client State: Active)
Crypto Sockets in Listen state:
Client: "TUNNEL SEC" Profile: "vpnprof" Map-name: "vpnprof-head-1"
Spoke1# show crypto map
Crypto Map: "vpnprof-head-1" idb: FastEthernet0/0/0 local address: 172.17.0.11
Crypto Map "vpnprof-head-1" 65536 ipsec-isakmp
        Profile name: vpnprof
        Security association lifetime: 4608000 kilobytes/3600 seconds
        PFS (Y/N): N
        Transform sets={
                trans2,
        }
Crypto Map "vpnprof-head-1" 65537 ipsec-isakmp
        Map is a PROFILE INSTANCE.
        Peer = 172.17.0.5
        Extended IP access list
            access-list  permit gre host 172.17.0.11 host 172.17.0.5
        Current peer: 172.17.0.5
        Security association lifetime: 4608000 kilobytes/3600 seconds
        PFS (Y/N): N
        Transform sets={
                trans2,
        }
Crypto Map "vpnprof-head-1" 65538 ipsec-isakmp
        Map is a PROFILE INSTANCE.
        Peer = 172.17.0.1
        Extended IP access list
            access-list  permit gre host 172.17.0.11 host 172.17.0.1
        Current peer: 172.17.0.1
        Security association lifetime: 4608000 kilobytes/3600 seconds
        PFS (Y/N): N
        Transform sets={
                trans2,
        }
Crypto Map "vpnprof-head-1" 65539 ipsec-isakmp
        Map is a PROFILE INSTANCE.
        Peer = 172.17.0.12
        Extended IP access list
            access-list  permit gre host 172.17.0.11 host 172.17.0.12
        Current peer: 172.17.0.12
        Security association lifetime: 4608000 kilobytes/3600 seconds
        PFS (Y/N): N
        Transform sets={
                trans2,
        }
        Interfaces using crypto map vpnprof-head-1:
                Tunnel1
                Tunnel0
 
                                 		
                                    
                                       | Note |   The three crypto sessions are shown under both tunnel interface (three entries, twice) in the                                                  		  show                                                     			 crypto                                                     			 ipsec                                                     			 sa  output because both interfaces are mapped to the same IPsec SADB, which has three entries. This duplication of output is                                                 expected in this case.                                                  		                                                | 
                                 
