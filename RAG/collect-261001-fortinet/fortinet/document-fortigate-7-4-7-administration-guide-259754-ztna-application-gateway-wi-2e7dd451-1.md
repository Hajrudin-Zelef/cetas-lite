---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-259754-ztna-application-gateway-wi-2e7dd451-1
title: "ZTNA application gateway with SAML and MFA using FortiAuthenticator example"
domain: fortinet
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-259754-ztna-application-gateway-wi-2e7dd451.md
source_anchor: ""
source_lines: [1, 111]
sha256: 88eee63cd8045f3b11feaee3451621d6db2ebba1da1c57683aacfbfde36150ff
---

# ZTNA application gateway with SAML and MFA using FortiAuthenticator example

ZTNA application gateway supports device verification using device certificates that are issued by EMS. To authenticate users, administrators can use either basic or SAML authentication. An advantage of SAML authentication is that multi-factor authentication (MFA) can be provided by the SAML Identity Provider (IdP).

In these examples, a FortiAuthenticator is used as the IdP, and MFA is applied to user authentication for remote users accessing the web, RDP, and SSH resources over the ZTNA application gateway. It is assumed that the FortiGate EMS fabric connector has already been successfully connected.


DNS resolutions:

- 
                                                    webserver.ztnademo.com:9443 -> 10.0.3.10:9443
- 
                                                    fac.ztnademo.com -> 10.0.3.7

The FortiAuthenticator (FAC) integrates with Active Directory (AD) on the Windows Domain Controller, which is also acting as the EMS server. Users are synchronized from the AD to the FAC, and remote users are configured with token-based authentication. SAML authentication is configured on the FortiGate, pointing to the FAC as the SAML IdP. The SAML server is applied to the ZTNA application gateway authentication scheme and rule, to provide the foundation for applying user authentication on individual ZTNA policies.

First configure FortiAuthenticator to synchronize users from AD using LDAP, apply MFA to individual remote users, and be the IdP.

###### To create a remote authentication server pointing to the Windows AD:

1. 
                                                    Go to *Authentication > Remote Auth. Servers > LDAP* and click*Create New* .
2. 
                                                    Configure the following: Name FortiAD Primary server name / IP 10.88.0.1 Port 389 (or another port if using LDAPS) Based distinguished name DC=fortiad,DC=info Bind type Regular Username <user account used for LDAP bind> Password <password of user> Server Type Microsoft Active Directory User object class person (default) Username attribute sAMAccountName (default) Group object class group (default) Obtain group membership from User attribute Group membership attribute memberOf (default) Secure connection Enable if using LDAPS or STARTTLS
3. 
                                                    Click *OK* .
4. 
                                                    Edit the *FortiAD* entry.
5. 
                                                    At the bottom, click *Import users* .
6. 
                                                    On the *Import Remote LDAP User* screen, configure the following settings:
  1. 
                                                            For *Remote LDAP Server* , select*FortiAD* .
  2. 
                                                            For *Action* , select*Import users* .
    1. 
                                                                    Click *Go* .
    2. 
                                                                    Select the users to import.
    3. 
                                                                    Click *OK* .
  3. 
                                                                    
 The screen displays the users that are imported. For more details, see LDAP in the FortiAuthenticator Administration Guide.
7. 
                                                            

###### To configure a remote LDAP user to use MFA:

1. 
                                                    Go to *Authentication > User Management > Remote Users* , and edit a user.
2. 
                                                    Enable *One-Time Password (OTP) authentication* .
  1. 
                                                            Set *Deliver token codes from* .
  2. 
                                                            Set *Deliver token by* .
 For this example, select *FortiToken > Mobile* , select the Token from the drop-down list, and set the*Activation delivery method* to email.
3. 
                                                            
4. 
                                                    In the *User Information* section, add the email address that will be used for the FortiToken activation.
5. 
                                                    Click *OK* .An activation email is sent to the user that they can use to install the token to their FortiToken Mobile app. For more details, see Remote users in the FortiAuthenticator Administration Guide.

###### To configure SAML IdP:

1. 
                                                    Go to *Authentication > SAML IdP > General* and enable*Enable SAML Identity Provider portal* .
2. 
                                                    The *Server address* is the device FQDN or IP address (configured in the System Information widget at*System > Dashboard > Status* ). In this example, it is*fac.ztnademo.com* .
3. 
                                                    Set *Username input format* to*username@realm* .
4. 
                                                    Click *Add a realm* in the*Realms* table:
  1. 
                                                            Set *Realm* to the just created LDAP realm (*FortiAD* ).
  2. 
                                                            Optionally, enable *Filter* and select the required users groups. In this example,*Sales* is configured.
5. 
                                                            
6. 
                                                    Set *Default IdP certificate* to the certificate that will be used in the HTTPS connection to the IdP portal.
7. 
                                                    Click *OK* .
8. 
                                                    Go to *Authentication > SAML IdP > Service Providers* , and click*Create New* to create a service provider (SP) for the FortiGate SP.
9. 
                                                    Configure the following, which must match what will be configured on the FortiGate: SP name saml-service-provider IdP prefix ztna Server certificate Use default setting in *SAML IdP > General* page.SP entity ID http://webserver.ztnademo.com:9443/remote/saml/metadata/ SP ACS (login) URL https://webserver.ztnademo.com:9443/remote/saml/login SP SLS (logout) URL https://webserver.ztnademo.com:9443/remote/saml/logout Participate in single logout Enable Where the *SP entity ID* ,*SP ACS (login) URL* , and*SP SLS (logout) URL* break down as follows:
  - 
                                                            *webserver.ztnademo.com* - The FQDN that resolves to the FortiGate SP.
  - 
                                                            *9443* - The port that is used to map to the FortiGate's SAML SP service.
  - 
                                                            */remote/saml* - The custom, user defined fields.
  - 
                                                            */metadata* ,*/login* , and*/logout* - The standard convention used to identify the SP entity, log in portal, and log out portal.
10. 
                                                            
11. 
                                                    Click *OK* .
12. 
                                                    Edit the just created SP object and, under *Assertion Attribute* , click*Add Assertion Attribute* .
13. 
                                                    Set *SAML attribute* to the username and set*User attribute* to*Username* , then click*OK* .
14. 
                                                    Click *OK* .

For more details, see Configuring SAML settings in the SAML Interoperability Guide.

