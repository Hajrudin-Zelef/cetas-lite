---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-3-administration-guide-266506-ssl-vpn-with-certificate-au-28dbd146-1
title: "SSL VPN with certificate authentication"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-3-administration-guide-266506-ssl-vpn-with-certificate-au-28dbd146.md
source_anchor: ""
source_lines: [1, 110]
sha256: 4e901c6f611f5d2c049c1786c1c324de06ffe0317cea9c720a8225bbccace954
---

# SSL VPN with certificate authentication

# SSL VPN with certificate authentication

                                                

There are two ways to configure certificate authentication:

In this example, the server and client certificates are signed by the same Certificate Authority (CA).

|  | Self-signed certificates are provided by default to simplify initial installation and testing. It is **HIGHLY** recommended that you acquire a signed certificate for your installation. Continuing to use these certificates can result in your connection being compromised, allowing attackers to steal your information, such as credit card details. For more information, please review the Use a non-factory SSL certificate for the SSL VPN portal and learn how to Procuring and importing a signed SSL certificate. | 


When using PKI users, the FortiGate authenticates the user based on there identity in the subject or the common name on the certificate. The certificate must be signed by a CA that is known by the FortiGate, either through the default CA certificates or through importing a CA certificate.

The user can either match a static subject or common name defined in the PKI user settings, or match an LDAP user in the LDAP server defined in the PKI user settings. Multi-factor authentication can also be enabled with the password as the second factor.

Using this method, the user is authenticated based on their regular username and password, but SSL VPN will still require an additional certificate check. The client certificate only needs to be signed by a known CA in order to pass authentication.

This method can be configured by enabling *Require Client Certificate* (`reqclientcert`) in the SSL-VPN settings.

### Configuration

In the following example, SSL VPN users are authenticated using the first method. A PKI user is configured with multi-factor authentication

Pre-requisites:

- 
                                                    The CA has already issued a client certificate to the user.
- 
                                                    The CA has issued a server certificate for the FortiGate's SSL VPN portal.
- 
                                                    The CA certificate is available to be imported on the FortiGate.

###### To configure SSL VPN in the GUI:

1. 
                                                    Install the server certificate. The server certificate allows the clients to authenticate the server and to encrypt the SSL VPN traffic. 
  1. 
                                                            Go to *System > Feature Visibility* and ensure*Certificates* is enabled.
  2. 
                                                            Go to *System > Certificates* and select*Import > Local Certificate* .
    - 
                                                                    Set *Type* to*Certificate* .
    - 
                                                                    Choose the *Certificate file* and the*Key file* for your certificate, and enter the*Password* .
    - 
                                                                    If required, you can change the *Certificate Name* .
  3. 
                                                                    
 The server certificate now appears in the list of *Certificates* .
2. 
                                                            
3. 
                                                    Install the CA certificate. The CA certificate is the certificate that signed both the server certificate and the user certificate. In this example, it is used to authenticate SSL VPN users. 
  1. 
                                                            Go to *System > Certificates* and select*Import > CA Certificate* .
  2. 
                                                            Select *Local PC* and then select the certificate file.
 The CA certificate now appears in the list of *External CA Certificates* . In this example, it is called*CA_Cert_1* .
4. 
                                                            
5. 
                                                    Configure PKI users and a user group. To use certificate authentication, use the CLI to create PKI users. ```
config user peer
    edit pki01
        set ca CA_Cert_1
        set subject "CN=User01"
    next
end
```
Ensure that the subject matches the name of the user certificate. In this example, *User01* .
6. 
                                                    After you have create a PKI user, a new menu is added to the GUI: 
  1. 
                                                            Go to *User & Authentication > PKI* to see the new user.
  2. 
                                                            Edit the user account.
  3. 
                                                            Enable *Two-factor authentication* and set a password for the account.
  4. 
                                                            Go to *User & Authentication > User Groups* and create a group  called*sslvpngroup* .
  5. 
                                                            Add the PKI user *pki01* to the group.
7. 
                                                            
8. 
                                                    Configure SSL VPN web portal. 
  1. 
                                                            Go to *VPN > SSL-VPN Portals* to edit the*full-access* portal.This portal supports both web and tunnel mode.
  2. 
                                                            Disable *Enable Split Tunneling* so that all SSL VPN traffic goes through the FortiGate.
9. 
                                                            
10. 
                                                    Configure SSL VPN settings. 
  1. 
                                                            Go to *VPN > SSL-VPN Settings* and enable SSL-VPN.
  2. 
                                                            Set the *Listen on Interface(s)* to*wan1* .
  3. 
                                                            Set *Listen on Port* to*10443* .
  4. 
                                                            Set *Server Certificate* to the local certificate that was imported.
  5. 
                                                            Under *Authentication/Portal Mapping* , set default Portal*web-access* for*All Other Users/Groups* .
  6. 
                                                            Create new *Authentication/Portal Mapping* for group*sslvpngroup* mapping portal*full-access* .
11. 
                                                            
