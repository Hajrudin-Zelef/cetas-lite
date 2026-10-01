---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-8-administration-guide-951346-68fbba0a-2
title: "document-fortigate-7-4-8-administration-guide-951346-68fbba0a"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-8-administration-guide-951346-68fbba0a.md
source_anchor: ""
source_lines: [81, 133]
sha256: b88f63e315a6954e306578de8337916814ea55bac6dccc161bffbf5466bcb6cd
---

# document-fortigate-7-4-8-administration-guide-951346-68fbba0a

                                                    Configure the following options: Name FCT_SAML Comments (Optional) Network IP Version IPv4 Remote Gateway Dialup User Interface port1 Select the IPsec tunnel gateway interface. Mode Config Enable Use system DNS in mode config (Optional) Enable FortiClient to use the host's DNS server after it connects to VPN. Assign IP From Enable Select Address/Address Group from the dropdown list. IPv4 mode config Client Address Range VPN_Client_IP_Range VPN_Client_IP_Range is configured from 10.212.134.1 to 10.212.134.200. If it is not already created, select Create > Address from the dropdown menu to create a new address object. See Subnet for more information. Subnet Mask 255.255.255.255 DNS Server 8.8.8.8 Authentication Method Pre-shared key Pre-shared key Enter the pre-shared key of at least six characters. IKE Version 2 Peer Options Accept Types Any peer ID Phase 1 Proposal Encryption AES128 Authentication SHA256 Select the desired Encryption and Authentication algorithms that should also match with Phase1 Proposals configured on FortiClient. See Configuring IPsec VPN profile on FortiClient.
- 
                                                    Keep other configurations as defaults.
- 
                                                    Click OK. The newly created IPsec tunnel would be now visible under VPN > IPsec Tunnels.
- 
                                                    As IKEv2 uses EAP for user authentication, enable EAP using the CLI inside the configured IPsec tunnel for user authentication, as follows: config vpn ipsec phase1-interface
    edit "FCT_SAML"
        set eap enable
        set eap-identity send-request
    next
end
For other advanced custom configurations as per your requirement, see Remote access.
|  | The SAML group configured, <group-name> , must be either configured inside the IPsec Phase 1 setting,set authusrgrp <group-name> , or in the firewall policy,set groups <group-name> , to allow the traffic to flow through the IPsec tunnel. If the SAML group is configured in both IPsec Phase 1 and firewall policy, the traffic stops to flow through the IPsec tunnel. In the example discussed, it is configured it in the firewall policy. See Using single or multiple user groups for user authentication for more information. | 
Configuring firewall policies for IPsec tunnel
To configure firewall policies for IPsec tunnel:
- 
                                                    Go to Policy & Object > Firewall Policy.
- 
                                                    Click Create New.
- 
                                                    Enter the following configuration: Name IPsec to DMZ Enter the desired name. Incoming Interface FCT_SAML Select the configured IPsec tunnel. outgoing Interface DMZ Select the interfaces that FortiClient needs access to when it connects to VPN. Source Under Address, select VPN_Client_IP_Range. Under User, select SAML-FAC-Group (or SAML-ENTRA-ID-Group). The group under User is the SAML user group configured in the earlier steps. Destination DMZ subnet Click Create if it is not already created. See Subnet for more information. Service ALL
- 
                                                    Click OK.
- 
                                                    As IPsec tunnel configured as full-tunnel, create another policy to allow traffic from IPsec to Internet, to allow FortiClient to access Internet through IPsec tunnel.
For additional custom settings as per your requirement, see Firewall policy.
Configuring IPsec VPN profile on FortiClient
To configure an IPsec VPN profile on FortiClient:
- 
                                                    In FortiClient, go to Remote Access > Configure VPN or Add a new connection.
- 
                                                    Set the following settings to configure an IPsec IKEv2 profile on FortiClient: Connection Name VPN-Tunnel Remote Gateway <VPN Gateway FQDN> or <VPN Gateway IP> Authentication Method Pre-shared key with Enable Single Sign On (SSO) for VPN Tunnel enabled. Customize port 9443 Advanced Settings > VPN Settings IKE Version 2 Options Mode Config
To explore additional custom options to configure IPsec VPN profile, see Configuring an IPsec VPN connection.
Verifying IPsec connection
To verify the IPsec connection in the GUI:
- 
                                                    On the client PC, open FortiClient and select the Remote Access tab.
- 
                                                    Select the VPN tunnel, VPN-Tunnel, and click Connect.
- 
                                                    If the connection is successful, a FortiClient pop-up will appear briefly indicating that the IKE negotiation succeeded. The Remote Access window now displays VPN Connected and the associated VPN tunnel details.
- 
                                                    In FortiOS, go to Dashboard > Network and locate the IPsec widget. Click the widget to expand to full view and view more details.
To verify the IPsec connection in the CLI:
The following debugs are from FortiGate when used with FortiAuthenticator as the IdP. The debugs should be similar for other IdPs depending on the SAML attributes supported and sent by the IdP.
- 
                                                    Verify the IKE gateway list: # diagnose vpn ike gateway list vd: root/0 name: FCT_SAML_0 version: 2 interface: port1 3 addr: 10.100.66.99:4500 -> 208.91.115.30:64917 tun_id: 10.212.134.1/::10.0.0.18 remote_location: 0.0.0.0 network-id: 0 transport: UDP created: 33s ago eap-user: testuser 2FA: no groups: SAML-FAC-Group 5 peer-id: 172.19.50.196 peer-id-auth: no FortiClient UID: 19E1FA565259468FB46EDAA9D595176F assigned IPv4 address: 10.212.134.1/255.255.255.255 nat: me peer PPK: no IKE SA: created 1/1 established 1/1 time 1680/1680/1680 ms IPsec SA: created 1/1 established 1/1 time 40/40/40 ms id/spi: 1049 f883b783547b0c64/f45745cd8b228850 direction: responder status: established 33-31s ago = 1680ms proposal: aes256-sha256 child: no SK_ei: 09d0e99e4ee86518-82da5e46c7ef0425-0816ef283fed3ca6-3fa0eeb56ac863a5 SK_er: 50e94be11ece32f8-aa13e54400e29531-684473a924ff04c5-8ebf45d854a59412 SK_ai: 3d95eec2deb54cf1-a59a945f0156c214-fe9aa188a96dd70c-f2394e1f7bb647b0 SK_ar: 0c0a478b800c7c9c-9dc56c05e9657200-7399b15d13ab8ad9-13984182abea936c PPK: no message-id sent/recv: 0/12 QKD: no lifetime/rekey: 86400/86098 DPD sent/recv: 00000000/00000000 peer-id: 172.19.50.196
- 
                                                    Verify the authd daemon debug output: # diagnose debug application authd -1 ... [authd_http_on_method_post:5151]: src 10.1.100.253 flag 00008000 [authd_local_saml_auth:5602]: SAML login with UID '19E1FA565259468FB46EDAA9D595176F'. [authd_http_prepare_javascript_redir:3852]: https://<VPN Gateway FQDN>:9443/saml?0704048f9683e491 ...
- 
                                                    Verify the samld daemon debug output: # diagnose debug application samld -1 ... </Session> samld_send_common_reply [99]: Attr: 17, 31, magic=040c07809dafc13e samld_send_common_reply [99]: Attr: 18, 29, 2024-03-19T21:42:21Z samld_send_common_reply [95]: Attr: 10, 26, 'username' 'testuser' samld_send_common_reply [95]: Attr: 10, 17, 'group' 'IT' ...
- 
