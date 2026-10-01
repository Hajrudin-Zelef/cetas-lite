---
id: collect-261001-general-networking/general-networking/c-en-us-support-docs-security-vpn-ipsec-negotiation-ike-protocols-117337-config-fb1867ae-1
title: "c-en-us-support-docs-security-vpn-ipsec-negotiation-ike-protocols-117337-config--fb1867ae"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/c-en-us-support-docs-security-vpn-ipsec-negotiation-ike-protocols-117337-config--fb1867ae.md
source_anchor: ""
source_lines: [1, 84]
sha256: e69a0adde2184e68bffa939f6c954a503f930235fc1d18db0df409d562b3e0e0
---

# c-en-us-support-docs-security-vpn-ipsec-negotiation-ike-protocols-117337-config--fb1867ae

The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
This document describes how to set up a site-to-site IKEv2 tunnel between a Cisco ASA and a router that runs Cisco IOS® software.
Cisco recommends that you have knowledge of these topics:
The information in this document is based on these software and hardware versions:
The information in this document was created from the devices in a specific lab environment. All of the devices used in this document started with a cleared (default) configuration. If your network is live, ensure that you understand the potential impact of any command.
This document can also be used with these hardware and software versions:
Configuration of an IKEv2 tunnel between an ASA and a router with the use of pre-shared keys is straightforward. However, when you use certificate authentication, there are certain caveats to keep in mind.
Certificate authentication requires that the clocks on all devices used must be synchronized to a common source. While the clock can be set manually on each device, this is not very accurate and can be cumbersome. The easiest method to synchronize the clocks on all devices is to use NTP. NTP synchronizes the time among a set of distributed time servers and clients. This synchronization allows events to be correlated when system logs are created and when other time-specific events occur. For more information on how to configure NTP, refer to Network Time Protocol: Best Practices White Paper.
Tip: When a Cisco IOS software Certificate Authority (CA) server is used, it is common practice to configure the same device as the NTP server. In this example, the CA server also serves as the NTP server.
Certificate lookup based on the HTTP URL avoids the fragmentation that results when large certificates are transferred. This feature is enabled on Cisco IOS software devices by default, so the cert req type 12 is used by Cisco IOS software.
If software versions that do not have the fix for Cisco bug ID CSCul48246 are used on the ASA, then the HTTP-URL-based lookup is not negotiated on the ASA, and Cisco IOS software causes the authorization attempt to fail.
On the ASA, if IKEv2 protocol debugs are enabled, these messages appear:
IKEv2-PROTO-1: (139): Auth exchange failed
IKEv2-PROTO-1: (140): Unsupported cert encoding found or Peer requested
HTTP URL but never sent
HTTP_LOOKUP_SUPPORTED Notification
In order to avoid this issue, use the no crypto ikev2 http-url cert command in order to disable this feature on the router when it peers with an ASA.
During IKE AUTH stage Internet Security Association and Key Management Protocol (ISAKMP) negotiations, the peers must identify themselves to each other. However, there is a difference in the way routers and ASAs select their local identity.
When IKEv2 tunnels are used on routers, the local identity used in the negotiation is determined by the identity local command under the IKEv2 profile:
R1(config-ikev2-profile)#identity local ?
address address
dn Distinguished Name
email Fully qualified email string
fqdn Fully qualified domain name string
key-id key-id opaque string - proprietary types of identification
By default, the router uses the address as the local identity.
The expected peer ID is also configured manually in the same profile with the match identity remote command:
R1(config-ikev2-profile)#match identity remote ?
address IP Address(es)
any match any peer identity
email Fully qualified email string [Max. 255 char(s)]
fqdn Fully qualified domain name string [Max. 255 char(s)]
key-id key-id opaque string
On ASAs, the ISAKMP identity is selected globally with the crypto isakmp identity command:
ciscoasa/vpn(config)# crypto isakmp identity ?
configure mode commands/options:
address Use the IP address of the interface for the identity
auto Identity automatically determined by the connection type: IP
address for preshared key and Cert DN for Cert based connections
hostname Use the hostname of the router for the identity
key-id Use the specified key-id for the identity
By default, the command mode is set to auto, which means that the ASA determines ISAKMP negotiation by connection type:
Note: Cisco bug ID CSCul48099 is an enhancement request for the ability to configure on a per-tunnel-group basis rather than in the global configuration.
Remote ID validation is done automatically (determined by the connection type) and cannot be changed. Validation can be enabled or disabled on a per-tunnel-group basis with the peer-id-validate command:
ciscoasa/vpn(config-tunnel-ipsec)# peer-id-validate ?
tunnel-group-ipsec mode commands/options:
cert If supported by certificate
nocheck Do not check
req Required
The difference in ID selection/validation causes two separate interoperability issues:
When cert auth is used on the ASA, the ASA tries to validate the peer ID from the Subject Alternative Name (SAN) on the received certificate. If peer ID validation is enabled and if IKEv2 platform debugs are enabled on the ASA, these debugs appear:
IKEv2-PROTO-3: (172): Getting configured policies
IKEv2-PLAT-3: attempting to find tunnel group for ID: 172.16.1.1
IKEv2-PLAT-3: mapped to tunnel group 172.16.1.1 using phase 1 ID
IKEv2-PLAT-3: (172) tg_name set to: 172.16.1.1
IKEv2-PLAT-3: (172) tunn grp type set to: L2L
IKEv2-PLAT-3: Peer ID check started, received ID type: IPv4 address
IKEv2-PLAT-2: Peer ID check: failed to retreive IP from SAN
IKEv2-PLAT-2: Peer ID check: failed to retreive DNS name from SAN
IKEv2-PLAT-2: Peer ID check: failed to retreive RFC822 name from SAN
IKEv2-PLAT-1: retrieving SAN for peer ID check
IKEv2-PLAT-1: Peer ID check failed
IKEv2-PROTO-1: (172): Failed to locate an item in the database
IKEv2-PROTO-1: (172):
IKEv2-PROTO-5: (172): SM Trace-> SA: I_SPI=833D2323FCB46093
R_SPI=F0B4D318DDDDB783 (I) MsgID = 00000001 CurState: I_PROC_AUTH
Event: EV_AUTH_FAIL
IKEv2-PROTO-3: (172): Verify auth failed
IKEv2-PROTO-5: (172): SM Trace-> SA: I_SPI=833D2323FCB46093
R_SPI=F0B4D318DDDDB783 (I) MsgID = 00000001 CurState: AUTH_DONE
Event: EV_FAIL
IKEv2-PROTO-3: (172): Auth exchange failed
For this issue, either the IP address of the certificate needs to be included in the peer certificate, or peer ID validation needs to be disabled on the ASA.
Similarly, by default the ASA selects the local ID automatically so, when cert auth is used, it sends the Distinguished Name (DN) as the identity. If the router is configured to receive the address as the remote ID, the peer ID validation fails on the router. If IKEv2 debugs are enabled on the router, these debugs appear:
Nov 30 22:49:14.464: IKEv2:(SESSION ID = 172,SA ID = 1):SM Trace-> SA: 
I_SPI=E9E4B7FD0A336C97 R_SPI=F2CF438C0CCA281C (R) MsgID = 1 CurState:
R_WAIT_AUTH Event: EV_GET_POLICY_BY_PEERID
Nov 30 22:49:14.464: IKEv2:(SESSION ID = 172,SA ID = 1):Searching policy
based on peer's identity 'hostname=asa.cisco.com' of type 'DER ASN1 DN'
Nov 30 22:49:14.464: IKEv2:%Profile could not be found by peer certificate.
Nov 30 22:49:14.468: IKEv2:% IKEv2 profile not found
Nov 30 22:49:14.468: IKEv2:(SESSION ID = 172,SA ID = 1):: Failed to
locate an item in the database
For this issue, either configure the router in order to validate the fully qualified domain name (FQDN) or configure the ASA in order to use address as the ISAKMP ID.
