---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-3-administration-guide-266506-ssl-vpn-with-certificate-au-28dbd146-2
title: "SSL VPN with certificate authentication"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-3-administration-guide-266506-ssl-vpn-with-certificate-au-28dbd146.md
source_anchor: ""
source_lines: [111, 263]
sha256: 983cc653d398d072f6b09123362c98e83ffe04e99dff3012a288dc439d9d9e5b
---

# SSL VPN with certificate authentication

12. 
                                                    Configure SSL VPN firewall policy. 
  1. 
                                                            Go to *Policy & Objects > Firewall Policy* .
  2. 
                                                            Fill in the firewall policy name. In this example, *sslvpn certificate auth* .
  3. 
                                                            Incoming interface must be *SSL-VPN tunnel interface(ssl.root)* .
  4. 
                                                            Set the *Source Address* to*all* and*Source User* to*sslvpngroup* .
  5. 
                                                            Set the *Outgoing Interface* to the local network interface so that the remote user can access the internal network. In this example,*port1* .
  6. 
                                                            Set *Destination Address* to the internal protected subnet*192.168.1.0* .
  7. 
                                                            Set *Schedule* to*always* ,*Service* to*ALL* , and*Action* to*Accept* .
  8. 
                                                            Enable *NAT* .
  9. 
                                                            Configure any remaining firewall and security options as needed.
  10. 
                                                            Click *OK* .
13. 
                                                            

###### To configure SSL VPN in the CLI:

1. 
                                                    Configure the protected subnet: ```
config firewall address
    edit "192.168.1.0"
        set subnet 192.168.1.0 255.255.255.0
    next
end
```
2. 
                                                    Install the server certificate: The server certificate allows the clients to authenticate the server and to encrypt the SSL VPN traffic. While it is easier to install the server certificate in the GUI, the CLI can be used to import a p12 certificate from a TFTP server. To import a p12 certificate, put the certificate *server_certificate.p12* on your TFTP server, then run following command on the FortiGate:execute vpn certificate local import tftp server_certificate.p12 <your tftp_server> p12 <your password for PKCS12 file> To check that the server certificate is installed: show vpn certificate local server_certificate
3. 
                                                    Install the CA certificate: The CA certificate is the certificate that signed both the server certificate and the user certificate. In this example, it is used to authenticate SSL VPN users. While it is easier to install the CA certificate from GUI, the CLI can be used to import a CA certificates from a TFTP server. To import a CA certificate, put the CA certificate on your TFTP server, then run following command on the FortiGate: execute vpn certificate ca import tftp <your CA certificate name> <your tftp server> To check that a new CA certificate is installed: show vpn certificate ca
4. 
                                                    Configure PKI users and a user group: ```
config user peer
    edit pki01
        set ca CA_Cert_1
        set subject "CN=User01"
        set two-factor enable
        set passwd **********
    next
end
```
```
config user group
    edit "sslvpngroup" 
        set member "pki01"
    next 
end
```
5. 
                                                    Configure SSL VPN web portal: ```
config vpn ssl web portal
    edit "full-access"
        set tunnel-mode enable
        set web-mode enable
        set ip-pools "SSLVPN_TUNNEL_ADDR1"
        set split-tunneling disable
     next
end
```
6. 
                                                    Configure SSL VPN settings: ```
config vpn ssl settings
    set servercert "server_certificate"
    set tunnel-ip-pools "SSLVPN_TUNNEL_ADDR1"
    set source-interface "wan1"
    set source-address "all"
    set default-portal "web-access"
    config authentication-rule
        edit 1
            set groups "sslvpngroup"
            set portal "full-access"
        next   
    end		
end
```
7. 
                                                    Configure one SSL VPN firewall policy to allow remote user to access the internal network: ```
config firewall policy 
    edit 1
        set name "sslvpn web mode access"
        set srcintf "ssl.root"
        set dstintf "port1"
        set srcaddr "all"
        set dstaddr "192.168.1.0"
        set groups "sslvpngroup"
        set action accept
        set schedule "always"
        set service "ALL"
        set nat enable
    next
end
```

### Installation

To use the user certificate, you must first install it on the user's PC. When the user tries to authenticate, the user certificate is checked against the CA certificate to verify that they match.

Every user should have a unique user certificate. This allows you to distinguish each user and revoke a specific user's certificate, such as if a user no longer has VPN access.

###### To install the user certificate on Windows 7, 8, and 10:

1. Double-click the certificate file to open the *Import Wizard* .
2. Use the *Import Wizard* to import the certificate into the*Personal store* of the current user.

###### To install the user certificate on Mac OS X:

1. Open the certificate file, to open *Keychain Access* .
2. Double-click the certificate.
3. Expand *Trust* and select*Always Trust* .

###### To see the results of tunnel connection:

1. Download FortiClient from www.forticlient.com.
2. Open the FortiClient Console and go to *Remote Access > Configure VPN* .
3. Add a new connection.
  - Set *VPN Type* to*SSL VPN* .
  - Set *Remote Gateway* to the IP of the listening FortiGate interface, in this example,*172.20.120.123* .
4. Set 
5. Select *Customize Port* and set it to*10443* .
6. Enable *Client Certificate* and select the authentication certificate.
7. Save your settings.
8. Use the credentials you've set up to connect to the SSL VPN tunnel.If the certificate is correct, you can connect.

###### To see the results of web portal:

1. In a web browser, log into the portal *http://172.20.120.123:10443* .A message requests a certificate for authentication.
2. Select the user certificate.
3. Enter your user credentials.If the certificate is correct, you can connect to the SSL VPN web portal.

###### To check the SSL VPN connection using the GUI:

1. Go to *VPN > Monitor > SSL-VPN Monitor* to verify the list of SSL users.
2. Go to *Log & Report > System Events* and select the*VPN Events* card to view the details for the SSL connection log.

###### To check the SSL VPN connection using the CLI:

get vpn ssl monitor
SSL VPN Login Users:
 Index   User    Auth Type      Timeout         From     HTTP in/out    HTTPS in/out
 0       pki01,cn=User01          1(1)            229     10.1.100.254   0/0     0/0
 1       pki01,cn=User01          1(1)            291     10.1.100.254   0/0     0/0
SSL VPN sessions:
 Index   User    Source IP      Duration        I/O Bytes       Tunnel/Dest IP 
 0       pki01,cn=User01          10.1.100.254    9       22099/43228    10.212.134.200
