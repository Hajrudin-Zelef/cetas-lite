---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-6
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [849, 939]
sha256: 74f20d2ea8283cfa7013ca5cf3474336156fa0fc5cc4dce57585893b355c82a3
---

# Configure DeviceA to generate a local key pair.

If different users access the network after RADIUS over DTLS is configured, you are advised to adjust the CPCAR value of packets based on the following table to ensure the user access rate.
| RADIUS over DTLS Scenario | Possible Problem | Workaround | Impact | 
|---|---|---|---|
| Administrator | The access rate of administrators is low. | Do not adjust the CPCAR value. | There is no impact. | 
| MAC | Many users need to be concurrently authenticated, and RADIUS DTLS packets need to be retransmitted. | Run the access-user car-mode dtls mac-authen wired command to automatically adjust the CPCAR value of MAC protocol packets to 20% of the default CPCAR value. | The access rate of wired MAC address authentication users in the non-RADIUS over DTLS authentication scenario is affected. | 
| Portal 2.0 | Many users need to be concurrently authenticated, and RADIUS DTLS packets need to be retransmitted. | Run the access-user car-mode dtls portal wired command to automatically adjust the CPCAR value of Portal protocol packets to 20% of the default CPCAR value. | The access rate of wired or wireless Portal 2.0 authentication users in the non-RADIUS over DTLS authentication scenario is affected. | 
| HTTPS Portal | Many users need to be concurrently authenticated, and RADIUS DTLS packets need to be retransmitted. | Run the access-user car-mode dtls https-portal wired command to automatically adjust the CPCAR value of HTTPS Portal protocol packets to 20% of the default CPCAR value. | The access rate of wired or wireless HTTPS Portal authentication users in the non-RADIUS over DTLS authentication scenario is affected. | 
| 802.1X CHAP/PAP | Many users need to be concurrently authenticated, and RADIUS DTLS packets need to be retransmitted. | Run the access-user car-mode dtls dot1x wired command to automatically adjust the CPCAR value of 802.1X protocol packets to 20% of the default CPCAR value. | The access rate of wired 802.1X authentication users in the non-RADIUS over DTLS authentication scenario is affected. | 
| 802.1X EAP | Many users need to be concurrently authenticated, and RADIUS DTLS packets need to be retransmitted. |  | The access rate of wired 802.1X authentication users in the non-RADIUS over DTLS authentication scenario is affected. | 
| Wireless 802.1X | Many users need to be concurrently authenticated, and RADIUS DTLS packets need to be retransmitted. |  | The access rates of wireless open authentication, MAC address authentication, and Portal authentication users are affected. | 
| Wireless MAC | Many users need to be concurrently authenticated, and RADIUS DTLS packets need to be retransmitted. | Run the access-user car-mode dtls mac-authen wireless command to automatically adjust the CPCAR value in user association to 20% of the default CPCAR value. | The access rates of wireless open authentication, 802.1X authentication, and Portal authentication users are affected. | 
| Wireless Portal | Many users need to be concurrently authenticated, and RADIUS DTLS packets need to be retransmitted. | Run the access-user car-mode dtls portal wireless command to automatically adjust the CPCAR value in user association and that of Portal protocol packets to 20% of the default CPCAR value. | The access rates of wireless open authentication, 802.1X authentication, and MAC address authentication users are affected. | 
| Wireless MAC+802.1X | Many users need to be concurrently authenticated, and RADIUS DTLS packets need to be retransmitted. | Run the access-user car-mode dtls mac-authen dot1x wireless command to automatically adjust the CPCAR value in user association and that of wireless 802.1X packets to 20% of the default CPCAR value. | The access rates of wireless open authentication, 802.1X authentication, and Portal authentication users are affected. | 
| RADIUS over DTLS in active/standby or load balancing scenario | The CPCAR value is small in the RADIUS over DTLS scenario. As a result, packets may fail to be received in the multi-server scenario. |  | In the multi-server scenario, servers may fail to receive packets. | 
Ensure that the PKI realm realm-name bound to the DTLS policy has a certificate loaded. The certificate can be an initial device certificate or a digital certificate applied by a user. For details about how to load a certificate using PKI, see "PKI Configuration" in CLI Configuration Guide > Security Configuration.
[DeviceA] pki realm abcd 
[DeviceA-pki-abcd] quit 
[DeviceA] pki import-certificate ca realm abcd pem filename test.cer
[DeviceA] dtls policy huawei
[DeviceA-dtls-policy-huawei] pki-domain abcd
[DeviceA-dtls-policy-huawei] quit
# Configure a RADIUS server template named shiva.
# Configure an IP address and a port number for the RADIUS authentication and accounting server.
[DeviceA-radius-shiva] radius-server authentication 10.7.66.66 2083 weight 80 dtls-policy huawei
[DeviceA-radius-shiva] radius-server accounting 10.7.66.66 2083 weight 80 dtls-policy huawei
# Create an authentication scheme named auth and set the authentication mode to RADIUS authentication and local authentication.
Run the display radius-server configuration template template-name command on DeviceA to check the RADIUS server template configuration. The command output shows that the configuration meets the requirements.
[DeviceA] display radius-server configuration template shiva
  ------------------------------------------------------------------------------
  Server-template-name          :  shiva
  Server-template-index         :  14                                      
  Protocol-version              :  standard
  Traffic-unit                  :  B
  Shared-secret-key             :  ****************
  Group-filter                  :  class  
  Timeout-interval(in second)   :  5
  Retransmission                :  2
  EndPacketSendTime             :  0
  Dead time(in minute)          :  5
  Domain-included               :  NO
  NAS-IP-Address                :  0.0.0.0
  Calling-station-id MAC-format :  xxxx-xxxx-xxxx
  Called-station-id MAC-format  :  XX-XX-XX-XX-XX-XX
  NAS-Port-ID format            :  New 
  Service-type                  :  - 
  NAS-IPv6-Address              :  ::
  Server algorithm              :  master-backup           
  Detect-interval(in second)    :  60     
  Detect up-server(in second)   :  0              
  Detect timeout(in second)     :  3                               
  Chargeable-user-identity      :  Not Support             
  CUI Not reject                :  No               
  Force framed-ip-addr Attr     :  No 
  Detect-interval(in second)    :  60 
  Authentication Server 1       :  10.7.66.66     Port:2083  Weight:80  [up]
                                   Vrf:- 
                                   Source Interface:NULL                        
                                   Source IP: ::
                                   Shared-key: -
                                   Down time left: -
                                   Dtls-Policy name: huawei
  Accounting Server     1       :  10.7.66.66     Port:1813  Weight:80  [up]
                                   Vrf:- 
                                   Source Interface:NULL                        
                                   Source IP: ::
                                   Shared-key: -
                                   Down time left: -
                                   Dtls-Policy name: huawei
  ------------------------------------------------------------------------------ 
#
pki realm abcd  
#
sysname DeviceA
#
#
dtls policy huawei
 pki-domain abcd
#
radius-server template shiva
 radius-server authentication 10.7.66.66 2083 weight 80 dtls-policy huawei
 radius-server accounting 10.7.66.66 2083 weight 80 dtls-policy huawei
#
aaa
 authentication-scheme auth
  authentication-mode radius local
 accounting-scheme abc
  accounting-mode radius
 domain huawei
  authentication-scheme auth
  accounting-scheme abc
  radius-server shiva
