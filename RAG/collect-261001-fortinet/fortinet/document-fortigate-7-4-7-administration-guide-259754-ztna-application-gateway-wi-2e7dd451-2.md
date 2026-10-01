---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-259754-ztna-application-gateway-wi-2e7dd451-2
title: "ZTNA application gateway with SAML and MFA using FortiAuthenticator example"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-259754-ztna-application-gateway-wi-2e7dd451.md
source_anchor: ""
source_lines: [112, 268]
sha256: 8c94ae4894aafba4fa89b7489a4b7a7b9e55ad174f8d9c4acbb8f156990dd44d
---

# ZTNA application gateway with SAML and MFA using FortiAuthenticator example

On FortiGate, a SAML user is used to define the SAML SP and IdP settings. This user is then applied to the ZTNA proxy using an authentication scheme, rule, and settings. A ZTNA server is then created to allow access to the SAML SP server so that end users can reach the FortiGate SP's captive portal. The SAML user must then be added to a ZTNA policy to trigger authentication when accessing the ZTNA application gateway.

###### To create a new SAML user/server from the GUI:

1. 
                                                    On the FortiGate, go to *User & Authentication > Single Sign-On* .
2. 
                                                    Click *Create New* .
3. 
                                                    Set *Name* to*ZTNA-FAC-SAML* .
4. 
                                                    Set *Address* to*webserver.ztnademo.com:9443* .The *Entity ID* ,*Assertion consumer service URL* and*Single logout service URL* will be updated. These should be the same URLs you inputted into FAC.
5. 
                                                    Enable *Certificate* , then select the certificate used for the client.In this example, the *ztna-wildcard* certificate is a local certificate that is used to sign SAML messages that are exchanged between the client and the FortiGate SP.
6. 
                                                    Click *Next* .
7. 
                                                    Use the settings from the Identity Provider to fill the custom *Identity Provider Details* . In this example, we select*Fortinet Product* , and fill in the fields as follows:Address fac.ztnademo.com Prefix ztna IdP certificate REMOTE_Cert_1 
  - 
                                                            Where the `REMOTE_Cert_1` certificate is a remote certificate that is used to identify the IdP; in this example, fac.ztnademo.com.
8. 
                                                            
9. 
                                                    Set *Attribute used to identify users* to*username* . Attributes to identify users and groups are case sensitive.
10. 
                                                    Click *Submit* to save the settings.

###### To create a new SAML user/server from the CLI:

```
config user saml
    edit "ZTNA-FAC-SAML"
        set cert "ztna-wildcard"
        set entity-id "http://webserver.ztnademo.com:9443/remote/saml/metadata/"
        set single-sign-on-url "https://webserver.ztnademo.com:9443/remote/saml/login"
        set single-logout-url "https://webserver.ztnademo.com:9443/remote/saml/logout"
        set idp-entity-id "http://fac.ztnademo.com/saml-idp/ztna/metadata/"
        set idp-single-sign-on-url "https://fac.ztnademo.com/saml-idp/ztna/login/"
        set idp-single-logout-url "https://fac.ztnademo.com/saml-idp/ztna/logout/"
        set idp-cert "REMOTE_Cert_1"
        set user-name "username"
        set digest-method sha1
    next
end
```
                                            ###### To create a user group for the SAML user object:

1. 
                                                    Under *User & Authentication > User Groups* , click*Create New* .
2. 
                                                    Set *Name* to*ztna-saml-users* .
3. 
                                                    Under *Remote Groups* , click*Add* .
4. 
                                                    For *Remote Server* , select*ZTNA-FAC-SAML* .
5. 
                                                    Click *OK* .
6. 
                                                    Click *OK* again to save.

###### To apply the SAML server to proxy authentication from the GUI:

1. 
                                                    Go to *Policy & Objects > Authentication Rules* .
2. 
                                                    Click *Create New > Authentication Scheme* .
3. 
                                                    Set the name to *ZTNA-SAML-scheme* .
4. 
                                                    Set *Method* to*SAML* .
5. 
                                                    Set *SAML SSO server* to*ZTNA-FAC-SAML* .
6. 
                                                    Click *OK* .
7. 
                                                    Go to *Policy & Objects > Authentication Rules* .
8. 
                                                    Click *Create New > Authentication Rule* .
9. 
                                                    Set the *Name* to*ZTNA-SAML-rule* .
10. 
                                                    Set *Source Address* to*all* .
11. 
                                                    Set *Incoming interface* to*port3* .
12. 
                                                    Set *Protocol* to*HTTP* .
13. 
                                                    Enable *Authentication Scheme* and select*ZTNA-SAML-scheme* .
14. 
                                                    Set *IP-based Authentication* to*Disable* .
15. 
                                                    Click *OK* .

###### To apply the SAML server to proxy authentication from the CLI:

```
config authentication scheme
    edit "ZTNA-SAML-scheme"
        set method saml
        set saml-server "ZTNA-FAC-SAML"
    next
end
config authentication rule
    edit "ZTNA-SAML-rule"
        set srcintf "port3"
        set srcaddr "all"
        set ip-based disable
        set active-auth-method "ZTNA-SAML-scheme"
        set web-auth-cookie enable
    next
end
```
                                            Assign an active authentication scheme and captive portal to serve the log in page for the SAML requests.

###### To configure the active authentication scheme and captive portal from the GUI:

1. 
                                                    Go to *User & Authentication > Authentication Settings* .
2. 
                                                    Enable *Authentication scheme* .
3. 
                                                    Select *ZTNA-SAML-scheme* .
4. 
                                                    Set *Captive portal type* to*FQDN* .
5. 
                                                    Enable *Captive Portal* .
6. 
                                                    Select the firewall address *webserver.ztnademo.com* . If this has not be created, create this firewall address.
7. 
                                                    Click *Apply* to save.

###### To configure the active authentication scheme and captive portal from the CLI:

```
config firewall address
    edit "webserver.ztnademo.com"
        set type fqdn
        set fqdn "webserver.ztnademo.com"
    next
end
config authentication setting
    set active-auth-scheme "ZTNA-SAML-scheme"
    set captive-portal "webserver.ztnademo.com"
end
```
                                            ###### To configure a ZTNA application gateway to allow SAML authentication requests to the SP:

1. 
                                                    Configure the ZTNA server: 
  1. 
                                                            Go to *Policy & Objects > ZTNA* , select the*ZTNA Servers* tab, and click*Create New* .
  2. 
                                                            Configure the following: Name ZTNA-access Interface Any IP 10.0.3.10 Port 9443 SAML Enabled SAML SSO Server ZTNA-FAC-SAML Default certificate ztna-wildcard
  3. 
                                                            Click *OK* .
2. 
                                                            
