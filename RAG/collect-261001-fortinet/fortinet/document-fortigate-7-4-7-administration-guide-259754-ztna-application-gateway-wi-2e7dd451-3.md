---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-259754-ztna-application-gateway-wi-2e7dd451-3
title: "ZTNA application gateway with SAML and MFA using FortiAuthenticator example"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-259754-ztna-application-gateway-wi-2e7dd451.md
source_anchor: ""
source_lines: [269, 367]
sha256: 91621365202c5e3f89268e324f44010f42f49d75925f3cfcbd8adec6f7963930
---

# ZTNA application gateway with SAML and MFA using FortiAuthenticator example

3. 
                                                    Define the full ZTNA policy to allow access to the ZTNA server: 
  1. 
                                                            Go to *Policy & Objects > Proxy policy* and click*Create New* .
  2. 
                                                            Configure the following: Name ZTNA-Rule Type ZTNA Incoming Interface port3 Source (Address) all Source (User) ztna-saml-users Destination all ZTNA Server ZTNA-access Action Accept Log Allowed Traffic All Sessions
  3. 
                                                            Click *OK* .
4. 
                                                            

###### To configure a VIP and a firewall policy to forward IdP authentication traffic to the FortiAuthenticator:

Remote clients connect to the FortiAuthenticator IdP behind the FortiGate using a VIP. In this example, users connect to the FQDN fac.ztnademo.com that resolves to the VIP's external IP address.

1. 
                                                    Configure the VIP to forward traffic to the FortiAuthenticator: 
  1. 
                                                            Go to *Policy & Objects > Virtual IPs* and navigate to the*Virtual IP* tab.
  2. 
                                                            Click *Create new* .
  3. 
                                                            Configure the following: Name FAC-VIP Interface Any External IP address 10.0.3.7 Map to > IPv4 address/range 10.88.0.7 Port Forwarding Enabled Protocol TCP External service port 443 Map to IPv4 port 443
  4. 
                                                            Click *OK* .
2. 
                                                            
3. 
                                                    Configure a firewall policy to allow VIP: 
  1. 
                                                            Go to *Policy & Objects > Firewall Policy* and click*Create New* .
  2. 
                                                            Configure the following: Name WAN_to_FAC Type Standard Incoming Interface port3 Outgoing Interface port2 Source All ZTNA Server FAC-VIP Schedule Always Service HTTP. HTTPS Action Accept NAT disabled
  3. 
                                                            Click *OK* .
4. 
                                                            

In this HTTPS access proxy example, two real servers are implemented with round robin load balancing performed between them. The HTTPS access proxy is configured on the same ZTNA server as was configured in the authentication step. The same ZTNA rule and firewall policy also apply.

###### To configure the ZTNA server for HTTPS access proxy with load balancing:

1. 
                                                    Go to *Policy & Objects > ZTNA* and select the*ZTNA Servers* tab.
2. 
                                                    Edit the ZTNA-access server.
3. 
                                                    In the *Service/server mapping* table, click*Create New* :
  1. 
                                                            Set *Service* to*HTTPS* .
  2. 
                                                            Set *Virtual Host* to*Any Host* .
  3. 
                                                            In the *Server* section, select*IP* :
    1. 
                                                                    Set *IP address* to*10.88.0.3* .
    2. 
                                                                    Set *Port* to*9443* .
    3. 
                                                                    Click *OK* .
  4. 
                                                                    
4. 
                                                            
5. 
                                                    Click *OK* .
6. 
                                                    Use the CLI to configure the second server with IP *10.88.0.4* .config firewall access-proxy edit "ZTNA-access" config api-gateway edit 1 config realservers edit 0 set ip 10.88.0.4 set port <real server 2 port> next end next end next end
7. 
                                                    In the GUI, edit the ZTNA-access server and verify the server mapping. 
  1. In the *Service/server* mapping table, edit the entry. The*Load balancing* option is visible, and additional real servers can be added by clicking*Create new* .
  2. Click *Cancel* .
8. In the 

###### Testing and verification:

From the remote endpoint, user Tom Smith attempts to connect to the web server over ZTNA:

1. 
                                                    On the remote Windows computer, open FortiClient and register to the EMS server. It is not necessary to configure a *ZTNA Destination* on the FortiClient for the HTTPS access proxy use case. In fact, configuring a*ZTNA Destination* rule for the website may interfere with its operation.
2. 
                                                    Open a browser and attempt to connect to the web server at *https://webserver.ztnademo.com:9443* .
3. 
                                                    Device authentication prompts the user for their device certificate. Select the certificate issued by EMS and click *OK* .
4. 
                                                    FortiGate receives the SAML request and redirects the user to the IdP login screen. Enter the username and password for Tom Smith and click *Login* .
5. 
                                                    A second prompt opens asking for the *Token Code* . Enter the code then click*Verify* .
6. 
                                                    The FortiAuthenticator IdP verifies the login, then sends the SAML assertion back to the user.
7. 
                                                    The browser redirects the assertion to the FortiGate SP, which decides if the user is allowed access.
8. 
                                                    On a successful log in, FortiGate redirects the user to the web page that they are trying to access.

###### Logs and debugs:

On the FortiGate, a successful connection can be seen in *Log & Report > Forward Traffic log*, or by using the CLI:

