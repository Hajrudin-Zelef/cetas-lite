---
id: collect-261001-fortinet/fortinet/document-fortigate-7-2-9-administration-guide-443323-dialup-ipsec-vpn-with-certi-8d455b86-2
title: "Dialup IPsec VPN with certificate authentication"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2021-08-23"]
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-2-9-administration-guide-443323-dialup-ipsec-vpn-with-certi-8d455b86.md
source_anchor: ""
source_lines: [167, 213]
sha256: 9ef14eb5c4d22376e650d4b7dd89271c7a97bcf9757b5b3bd0caaedd5964b656
---

# Dialup IPsec VPN with certificate authentication

1. Go to *VPN > IPsec Tunnels* and click*Create New > IPsec Tunnel* .
2. Enter a name for the tunnel, *Dialup-cert_0* .
3. For *Template type* , select*Custom* then click*Next* .
4. In the *Network* section, enter the following:Remote Gateway *Dialup User*Interface *port1*Mode Config Enable Assign IP From *Range*IPv4 mode config > Client Address Range *172.18.200.10-172.18.200.99*Enable IPv4 Split Tunnel Enable Accessible Networks *192.168.20.0*
5. In the *Authentication* section, enter the following:Method *Signature*Certificate Name Select the server certificate that was imported. Mode *Aggressive*Peer Options > Accept Types *Peer certificate group*Peer Options > Peer certificate group Select the group based on the preferred method: 
  - For subject verification, select *pki-users* .
  - For LDAP integration, select *pki-ldap* .
 When IKEv1 is used, aggressive mode should be selected so that the connecting endpoint will provide its peer ID in the first message of the IKE exchange. The peer identifier allows the FortiGate to match the correct tunnel when multiple dialup tunnels are defined.
6. For subject verification, select 
7. For *Phase 2 Selectors* , leave the local and remote selectors as*0.0.0.0/0.0.0.0* .
8. Click *OK* .

###### To configure the firewall policy:

1. Go to *Policy & Objects > Firewall Policy* and click*Create New* .
2. Configure the following:Name Enter a policy name. Incoming interface *Dialup-cert_0*Outgoing Interface *port3*Source *remote-user-range*Destination *192.168.20.0*Schedule *always*Service *ALL*Action *ACCEPT*
3. Configure the other settings as needed.
4. Click *OK* .

The following example is configured on a Windows PC with FortiClient 7.0.0. Other configurations may differ slightly.

The user certificate and CA certificate must be installed on the endpoint device. They may be pushed by the administrator through group policies or another method. This example assumes that the user certificate and CA certificate are already installed on the endpoint.

###### To verify the user and CA certificates:

1. Open the Windows certificate manager (certmgr):
  1. In the Control Panel, type *Manage user certificate* in the search box.
  2. Click the result, *Manage user certificates* .
2. In the Control Panel, type 
3. Go to *Personal > Certificate* . The user certificate should be listed.
4. Go to *Trusted Root Certification Authorities > Certificates* . The company CA certificate should be listed.

###### To configure the FortiClient endpoint settings:

1. In FortiClient, click the *Remote Access* tab and add a new connection:
  1. If there are no existing connections, click *Configure VPN* .
  2. If there are existing connections, click the menu icon and select *Add a new connection* .
2. If there are no existing connections, click 
3. Configure the following:VPN *IPsec VPN*Connection Name *Dialup-cert_0*Remote Gateway *192.168.2.5*Authentication Method *X.509 Certificate*Select the user certificate, *tgerber/root CA* , from the dropdown.Authentication (XAuth) *Disable*
4. Click *Save* .

1. On the client PC, open FortiClient and click the *Remote Access* tab.
2. Select the VPN tunnel, *Dialup-cert_0* , and click*Connect* .If the connection is successful, a FortiClient pop-up will appear briefly indicating that the IKE negotiation succeeded. The *Remote Access* window now displays*VPN Connected* and the associated VPN tunnel details.
3. On the FortiGate, go to *Dashboard > Network* and locate the*IPsec* widget to view the VPN tunnel monitor. Click the widget to expand to full view.The widget displays tunnel information, including the *Peer ID* containing the subject field of the user certificate.
4. Go to *Log & Report > System Events* and select the*VPN Events* card. Several tunnel related logs are recorded.
5.  The same logs can be viewed in the CLI:# execute log filter category 1 # execute log filter field subtype vpn # execute log display 7: date=2021-08-23 time=15:53:08 eventtime=1629759188862005740 tz="-0700" logid="0101037138" type="event" subtype="vpn" level="notice" vd="root" logdesc="IPsec connection status changed" msg="IPsec connection status change" **action="tunnel-up"** remip=192.168.2.1 locip=192.168.2.5 remport=64916 locport=4500 outintf="port1" cookies="19f05ebc8c2f7a0d/7716190005538db5"**user="C = CA, ST = British Columbia, L = Burnaby, O = FortiKeith, OU = TAC, CN = tgerber" group="pki-ldap"** useralt="C = CA, ST = British Columbia, L = Burnaby, O = FortiKeith, OU = TAC, CN = tgerber" xauthuser="N/A" xauthgroup="N/A" assignip=172.18.200.10 vpntunnel="Dialup-cert_0" tunnelip=172.18.200.10 tunnelid=3418215253 tunneltype="ipsec" duration=0 sentbyte=0 rcvdbyte=0 nextstat=0
6. If any issues arise during the connection, run the following debug commands to troubleshoot the issue:# diagnose debug application ike -1 # diagnose debug application fnbamd -1 # diagnose debug enable
