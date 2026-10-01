---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-259754-ztna-application-gateway-wi-2e7dd451-4
title: "ZTNA application gateway with SAML and MFA using FortiAuthenticator example"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2023-05-09"]
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-259754-ztna-application-gateway-wi-2e7dd451.md
source_anchor: ""
source_lines: [368, 470]
sha256: 79ea5d317838ab560c9735af7ccb3d40ffc503d0a07c13c221341b31742984f9
---

# execute log filter category 0
# execute log filter field subtype srcip 
# execute log display
...
2: date=2023-05-09 time=00:47:56 eventtime=1683618475812847145 tz="-0700" logid="0005000024" type="traffic" subtype="ztna" level="notice" vd="root" **srcip=10.0.3.2** srcport=17201 srcintf="port3" srcintfrole="wan" dstcountry="Reserved" srccountry="Reserved" **dstip=10.88.0.4** dstport=9443 dstintf="port2" dstintfrole="dmz" sessionid=141588 service="tcp/9443" proxyapptype="http" proto=6 action="accept" policyid=1 policytype="proxy-policy" **policyname="ZTNA-Rule"** duration=83 **user="tsmith" group="ztna-saml-users"** authserver="ZTNA-FAC-SAML" gatewayid=1 **realserverid=2** vip="ZTNA-access" accessproxy="ZTNA-access" clientdevicemanageable="manageable" wanin=316227 rcvdbyte=316227 wanout=6477 lanin=14843 sentbyte=14843 lanout=315122 appcat="unscanned"
...
8: date=2023-05-09 time=00:46:09 eventtime=1683618369392379538 tz="-0700" logid="0005000024" type="traffic" subtype="ztna" level="notice" vd="root" **srcip=10.0.3.2** srcport=17177 srcintf="port3" srcintfrole="wan" dstcountry="Reserved" srccountry="Reserved" **dstip=10.88.0.3** dstport=9443 dstintf="port2" dstintfrole="dmz" sessionid=141526 service="tcp/9443" proxyapptype="http" proto=6 action="accept" policyid=1 policytype="proxy-policy" poluuid="08a04362-ee38-51ed-77fa-737e2656f04a" **policyname="ZTNA-Rule"** duration=63 **user="tsmith" group="ztna-saml-users"** authserver="ZTNA-FAC-SAML" gatewayid=1 **realserverid=1** vip="ZTNA-access" accessproxy="ZTNA-access" clientdevicemanageable="manageable" wanin=313997 rcvdbyte=313997 wanout=4894 lanin=14676 sentbyte=14676 lanout=314865 appcat="unscanned"
...
10: date=2023-05-09 time=00:45:26 eventtime=1683618326285366024 tz="-0700" logid="0000000013" type="traffic" subtype="forward" level="notice" vd="root" **srcip=10.0.3.2** srcport=17181 srcintf="port3" srcintfrole="wan" dstip=10.0.3.7 dstport=443 dstintf="port2" dstintfrole="dmz" srccountry="Reserved" dstcountry="Reserved" sessionid=141539 proto=6 action="client-rst" policyid=6 policytype="policy" poluuid="2e840d4e-1184-51ec-63b9-9b805a8b7344" **policyname="WAN_to_FAC"** service="HTTPS" trandisp="dnat" tranip=10.88.0.7 tranport=443 duration=10 sentbyte=3449 rcvdbyte=4372 sentpkt=13 rcvdpkt=13 appcat="unscanned"

Log number ten shows that authentication first passes through the WAN_to_FAC policy. Log numbers two and eight show the traffic allowed through the ZTNA proxy-policy over two successive sessions. Note that they have different destination IP addresses (`dstip`), indicating that ZTNA was performing server load balancing.

Use the following command to show if the FortiGate's WAD process has an active record of the SAML user login:

# diagnose wad user list 
ID: 8, VDOM: root, **IPv4: 10.0.3.2**
  user name   : **tsmith**
  worker      : 1
  duration    : 332
  auth_type   : Session
  auth_method : **SAML**
  pol_id      : 1
  g_id        : 4
  user_based  : 0
  expire      : 390
  LAN:
    bytes_in=29519 bytes_out=629987
  WAN:
    bytes_in=958636 bytes_out=19201

In this TCP forwarding access proxy example, RDP connections are allowed to be forwarded to the Windows/EMS server. Traffic to TCP/3389 is allowed through the ZTNA proxy.

###### To configure the ZTNA server for TCP forwarding on TCP/3389:

1. 
                                                    Create a firewall address for the Windows/EMS server: 
  1. 
                                                            Go to *Policy & Objects > Addresses* and select*Address* .
  2. 
                                                            Click *Create new* .
  3. 
                                                            Configure the following: Name winserver Type Subnet IP/Netmask 10.88.0.1/32 Interface any
  4. 
                                                            Click *OK* .
2. 
                                                            
3. 
                                                    Go to *Policy & Objects > ZTNA* and select the*ZTNA Servers* tab.
4. 
                                                    Edit the ZTNA-access server.
5. 
                                                    In the *Service/server mapping* table, click*Create New* :
  1. 
                                                            Set *Service* to*TCP Forwarding* .
  2. 
                                                            In the *Server* section, set*Address* to*winserver* .
  3. 
                                                            Set *Ports* to*3389* .
  4. 
                                                            Click *OK* .
6. 
                                                            
7. 
                                                    Click *OK* .

###### Testing and verification:

On the remote endpoint, manually configure a ZTNA destination to forward RDP traffic to the ZTNA application gateway. The rules can also be pushed from the EMS server; for details see Provisioning ZTNA TCP forwarding rules via EMS.

Configure the ZTNA Destination:

1. 
                                                    On the remote Windows computer, open FortiClient.
2. 
                                                    Register to the EMS server.
3. 
                                                    On the *ZTNA Destination* tab, click*Add Destination* to add a TCP forwarding rule.
4. 
                                                    Configure the following: Rule Name RDP-server Destination Host 10.88.0.1:3389 Proxy Gateway 10.0.3.10:9443 Mode Transparent Encryption Disabled *Encryption* can be enabled or disabled. When it is disabled, the client to access proxy connection is not encrypted in HTTPS. Because RDP is encrypted by default, disabling*Encryption* does not reduce security.
5. 
                                                    Click *Create* .

Connect over RDP:

1. 
                                                    On the remote PC, open a new RDP connection.
2. 
                                                    Enter the IP address 10.88.0.1. By default, RDP session use port 3389. When the connection to the ZTNA application gateway is established, FortiGate will redirect the SAML login request to the FortiAuthenticator IdP. A FortiClient prompt will open with the FortiAuthenticator login screen.
3. 
                                                    Enter the username and password then click *Login* .
4. 
                                                    A second prompt opens asking for the *Token Code* . Enter the code from your FortiToken app, then click*Verify* .
5. 
                                                    FortiAuthenticator verifies the token code, determines if the login is successful, then sends the SAML assertion back to the client.
6. 
                                                    The client redirects the response back to the FortiGate SP.
7. 
                                                    If the log in was successful, the user can now log on to the RDP session.

###### Logs and debugs:

On the FortiGate, a successful connection can be seen in *Log & Report > Forward Traffic log*, or by using the CLI:

