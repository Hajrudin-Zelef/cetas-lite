---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-ipsec-vpn-with-sophos-is-not-working-181232-e2ff36af
title: "fortigate-3-troubleshooting-tip-ipsec-vpn-with-sophos-is-not-working-181232-e2ff36af"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-ipsec-vpn-with-sophos-is-not-working-181232-e2ff36af.md
source_anchor: ""
source_lines: [1, 4]
sha256: a97573adff1867d6adaea1b5734b027413145ffa6fa782acc051d8eaed425ecd
---

# fortigate-3-troubleshooting-tip-ipsec-vpn-with-sophos-is-not-working-181232-e2ff36af

Troubleshooting Tip: IPsec VPN with Sophos is not working
| Description | This article describes workarounds when a VPN tunnel cannot be established between a FortiGate and Sophos. | 
| Scope | FortiGate. | 
| Solution | After configuring both using IKE v1, it is verified that the configuration is correct on both sides. However, phases 1 and 2 are still down. If there are no restrictions in the tunnel configuration, change the IKE version from 1 to 2.  For configuring FortiGate and Sophos using IKE v2, refer to the following document: Technical Tip: Set up IPsec VPN between FortiGate and Sophos XG using IKEv2.  If the tunnel is still down after configuring both devices correctly, run the following commands:  diagnose debug disable diagnose debug reset diagnose vpn ike log filter clear diagnose vpn ike log filter rem-addr4 <Public IP> <----- Sophos/Arista public IP. diagnose debug application ike -1 diagnose debug enable  Notes:   If the following error is encountered when running the debug:  ike 0: IKEv2 exchange=AUTH_RESPONSE id=554498e2804b4c46/9352e4e98f9c4c1a:00000001 len=72 ike 0: in D4A1F7692BCDE5A453EF234D14B7891E2F213450123456780123D561234A1BC97543EF98A5 6C9D8A43BC5D9A79C5B7A42BDFE498DC3AFDB10C23456789ABCDEF0 ike 0:IPSECVPN_Test:1535625: dec 65D3AE182A9BF8E453CD353C74A2861E2E2023200000000100000028290000040000000801000018 ike 0:IPSECVPN_Test:1535625: initiator received AUTH msg ike 0:IPSECVPN_Test:1535625: received notify type AUTHENTICATION_FAILED  If the IKE debug message shows a 'malformed responder cookie', verify whether the FortiGate has a local ID configured and whether the Sophos device has the corresponding peer ID set correctly.  Make the following changes:    config vpn ipsec phase1-interface     edit " Tunnel" set localid-type address set localid 79.88.88.88 <- Public IP address of the FortiGate. end  Notes:    2026-01-16T17:30:39.382 \|3336791\| ERR  [VPN] vc_ike_process_msg:456 IKE DS look fail on packet X.X.X.X:4500 ==> Y.Y.Y.Y:4500 ike_cky={2acc9d05a207c39d 2052078b4ba21f64} 2026-01-16T17:30:39.382 \|3336791\| ERR  [VPN] vc_ike_process_msg:456 IKE DS look fail on packet Y.Y.Y.Y:4500 ==> X.X.X.X:4500 ike_cky={d49a7f4f142f4605 d15626edb2d9758c}  After the changes, the authentication error will be fixed, and both phases of the tunnel will be up.  Note: Authentication failure can also be related to a pre-shared key (PSK) mismatch. |
