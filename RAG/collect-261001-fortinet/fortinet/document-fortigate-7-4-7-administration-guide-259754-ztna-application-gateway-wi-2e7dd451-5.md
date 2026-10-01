---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-259754-ztna-application-gateway-wi-2e7dd451-5
title: "ZTNA application gateway with SAML and MFA using FortiAuthenticator example"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2023-05-09"]
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-259754-ztna-application-gateway-wi-2e7dd451.md
source_anchor: ""
source_lines: [471, 597]
sha256: 67fbdf1f27d159e89c6ebfb86dc643af223f0812426f0746f186d5d2daf347fc
---

# execute log filter category 0
# execute log filter field srcip 10.0.3.2
# execute log display
...
8: date=2023-05-09 time=01:14:34 eventtime=1683620074907124357 tz="-0700" logid="0005000024" type="traffic" subtype="ztna" level="notice" vd="root" **srcip=10.0.3.2** srcport=17556 srcintf="port3" srcintfrole="wan" dstcountry="Reserved" srccountry="Reserved" **dstip=10.88.0.1** dstport=3389 dstintf="port2" dstintfrole="dmz" sessionid=142359 **service="RDP"** proxyapptype="http" proto=6 action="accept" policyid=1 policytype="proxy-policy" poluuid="08a04362-ee38-51ed-77fa-737e2656f04a" policyname="ZTNA-Rule" duration=85 **user="tsmith" group="ztna-saml-users"** gatewayid=3 vip="ZTNA-access" accessproxy="ZTNA-access" clientdevicemanageable="manageable" wanin=0 rcvdbyte=0 wanout=0 lanin=3639 sentbyte=3639 lanout=3797 appcat="unscanned"
9: date=2023-05-09 time=01:14:16 eventtime=1683620056835075857 tz="-0700" logid="0000000013" type="traffic" subtype="forward" level="notice" vd="root" **srcip=10.0.3.2** srcport=17571 srcintf="port3" srcintfrole="wan" **dstip=10.0.3.7** dstport=443 dstintf="port2" dstintfrole="dmz" srccountry="Reserved" dstcountry="Reserved" sessionid=142389 proto=6 action="client-rst" policyid=6 policytype="policy" poluuid="2e840d4e-1184-51ec-63b9-9b805a8b7344" **policyname="WAN_to_FAC"** service="HTTPS" trandisp="dnat" **tranip=10.88.0.7** tranport=443 duration=10 sentbyte=3368 rcvdbyte=4376 sentpkt=13 rcvdpkt=13 appcat="unscanned"

Use the following command to show if the FortiGate's WAD process has an active record of the SAML user login:

# diagnose wad user list
ID: 9, VDOM: root, **IPv4: 10.0.3.2**
  user name   : **tsmith**
  worker      : 1
  duration    : 164
  auth_type   : Session
  auth_method : **SAML**
  pol_id      : 1
  g_id        : 4
  user_based  : 0
  expire      : no
  LAN:
    bytes_in=80950 bytes_out=233010
  WAN:
    bytes_in=219973 bytes_out=58315

In this SSH access proxy example, SSH connections can be forwarded to the FortiAnalyzer (10.88.0.2).

###### To configure the ZTNA server for SSH access proxy:

1. 
                                                    Create a firewall address for the web server: 
  1. 
                                                            Go to *Policy & Objects > Addresses* and select*Address* .
  2. 
                                                            Click *Create new* .
  3. 
                                                            Configure the following: Name FAZ Type Subnet IP/Netmask 10.88.0.2/32 Interface any
  4. 
                                                            Click *OK* .
2. 
                                                            
3. 
                                                    Go to *Policy & Objects > ZTNA* and select the*ZTNA Servers* tab.
4. 
                                                    Edit the ZTNA-access server.
5. 
                                                    In the *Service/server mapping* table, edit the*TCP Forwarding* entry:
  1. 
                                                            In the *Servers* table click*Create New* :
    1. 
                                                                    Set *Address* to*FAZ* .
    2. 
                                                                    Set *Ports* to*22* .
    3. 
                                                                    Optionally, enable *Additional SSH Options* to configure other SSH options as needed.
    4. 
                                                                    Click *OK* .
  2. 
                                                                    
  3. 
                                                            Click *OK* .
6. 
                                                            
7. 
                                                    Click *OK* .

###### Testing and verification:

On the remote endpoint, manually configure ZTNA destination to forward SSH traffic to the ZTNA access proxy. The destination can also be pushed from the EMS server; for details see Provisioning ZTNA TCP forwarding rules via EMS.

Configure the ZTNA Destination:

1. 
                                                    On the remote Windows computer, open FortiClient.
2. 
                                                    Register to the EMS server.
3. 
                                                    On the *ZTNA Destination* tab, click*Add Destination* to add a TCP forwarding rule.
4. 
                                                    Configure the following: Rule Name SSH-FAZ Destination Host 10.88.0.2:22 Proxy Gateway 10.0.3.10:9443 Mode Transparent Encryption Disabled
5. 
                                                    Click *Create* .

Connect over SSH:

1. 
                                                    On the remote PC, open a new SSH connection.
2. 
                                                    Enter the host admin@10.88.0.2 on port 22. When the connection to the ZTNA application is established, FortiGate will redirect the SAML login request to the FortiAuthenticator IdP. A FortiClient prompt will open with the FortiAuthenticator login screen.
3. 
                                                    Enter the username and password then click *Login* .
4. 
                                                    A second prompt opens asking for the *Token Code* . Enter the code from your FortiToken app, then click*Verify* .
5. 
                                                    FortiAuthenticator verifies the token code, determines if the login is successful, then sends the SAML assertion back to the client.
6. 
                                                    The client redirects the response back to the FortiGate SP.
7. 
                                                    If the log in was successful, the user can now log on to the SSH session.

###### Logs and debugs:

On the FortiGate, a successful connection can be seen in *Log & Report > Forward Traffic log*, or by using the CLI:

# execute log filter category 0
# execute log filter field srcip 10.0.3.2
# execute log display
...
5: date=2023-05-09 time=10:06:04 eventtime=1683651963820802005 tz="-0700" logid="0005000024" type="traffic" subtype="ztna" level="notice" vd="root" **srcip=10.0.3.2** srcport=22536 srcintf="port3" srcintfrole="wan" dstcountry="Reserved" srccountry="Reserved" **dstip=10.88.0.2 dstport=22** dstintf="port2" dstintfrole="dmz" sessionid=151696 **service="SSH"** proxyapptype="http" proto=6 action="accept" policyid=1 policytype="proxy-policy" policyname="ZTNA-Rule" duration=9 **user="tsmith" group="ztna-saml-users"** gatewayid=3 vip="ZTNA-access" accessproxy="ZTNA-access" clientdevicemanageable="manageable" wanin=2981 rcvdbyte=2981 wanout=2753 lanin=4663 sentbyte=4663 lanout=5247 appcat="unscanned"

Use the following command to show if the FortiGate's WAD process has an active record of the SAML user login:

# diagnose wad user list
ID: 10, VDOM: root, **IPv4: 10.0.3.2**
  user name   : **tsmith**
  worker      : 1
  duration    : 387
  auth_type   : Session
  auth_method : **SAML**
  pol_id      : 1
  g_id        : 4
  user_based  : 0
  expire      : no
  LAN:
    bytes_in=23261 bytes_out=14024
  WAN:
    bytes_in=3242 bytes_out=2541
