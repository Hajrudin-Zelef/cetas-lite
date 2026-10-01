---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-8-administration-guide-951346-68fbba0a-3
title: "document-fortigate-7-4-8-administration-guide-951346-68fbba0a"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-8-administration-guide-951346-68fbba0a.md
source_anchor: ""
source_lines: [134, 136]
sha256: 712797d8653b8ddf6d09069a5a2a8ccb1941e8873026e6b7ef7e7581e2228498
---

# document-fortigate-7-4-8-administration-guide-951346-68fbba0a

                                                    Verify the fnbamd daemon debug output: # diagnose debug application fnbamd -1 ... [2426] handle_req-Rcvd auth cache message [133] __saml_auth_cache_push-Auth cache created, user='19E1FA565259468FB46EDAA9D595176F', SAML_server='saml-fac', vfid=0 [140] __saml_auth_cache_push-Hash bucket 227 [182] __saml_auth_cache_push-New auth cache entry is created, user='19E1FA565259468FB46EDAA9D595176F', expires=1648598587, SAML_server='saml-fac', vfid=0 [1918] handle_req-Rcvd auth req 994781475 for 19E1FA565259468FB46EDAA9D595176F in ipsec opt=00000000 prot=5 [466] __compose_group_list_from_req-Group 'saml-fac', type 1 [971] fnbamd_saml_auth_cache_lookup-Authneticating '19E1FA565259468FB46EDAA9D595176F'. [1005] fnbamd_saml_auth_cache_lookup-Authentication passed.
- 
                                                    Verify the IPsec daemon debug output: # diagnose debug application ike -1 ... ike V=root:0:FCT_SAML: user 'testuser' authenticated group 'SAML-FAC-Group' 5 ike V=root:0:FCT_SAML:1180: responder preparing EAP pass through message ... ike V=root:0:FCT_SAML_0:1180: mode-cfg assigned (1) IPv4 address 10.212.134.1 ike V=root:0:FCT_SAML_0:1180: mode-cfg assigned (2) IPv4 netmask 255.255.255.255 ike V=root:0:FCT_SAML_0:1180: mode-cfg send (13) 0:0.0.0.0/0.0.0.0:0 ike V=root:0:FCT_SAML_0:1180: mode-cfg send (3) IPv4 DNS(1) 8.8.8.8 ... ike V=root:0:FCT_SAML_0: sent tunnel-up message to EMS: (fct-uid=19E1FA565259468FB46EDAA9D595176F, intf=FCT_SAML_0, addr=10.212.134.1, vdom=root) ike V=root:0:FCT_SAML_0: user 'testuser' 10.212.134.1 groups 1 ...
