---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-administration-guide-018013-outbound-firewall-authentic-b398612c
title: "Outbound firewall authentication for a SAML user"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-administration-guide-018013-outbound-firewall-authentic-b398612c.md
source_anchor: ""
source_lines: [1, 61]
sha256: bce39398c85531448c4b5950712425d389732f3611cbac070917c40357746e5b
---

# Outbound firewall authentication for a SAML user

# Outbound firewall authentication for a SAML user

When you configure a FortiGate as a service provider (SP), you can create an authentication profile that uses SAML for firewall authentication.

|  | You must use the identity provider's (IdP) remote certificate on the SPs. | 

The following example uses a FortiGate as an SP and FortiAuthenticator as the IdP server:


###### To configure firewall authentication:

1. Configure the FortiGate SP to be a SAML user:```
config user saml
    edit "fac-firewall"
        set entity-id "http://10.2.2.2:1000/saml/metadata/"
        set single-sign-on-url "https://10.2.2.2:1003/saml/login/"
        set single-logout-url "https://10.2.2.2:1003/saml/logout/"
        set idp-entity-id "http://172.18.58.93:443/saml-idp/bbbbbb/metadata/"
        set idp-single-sign-on-url "https://172.18.58.93:443/saml-idp/bbbbbb/login/"
        set idp-single-logout-url "https://172.18.58.93:443/saml-idp/bbbbbb/logout/"
        set idp-cert "REMOTE_Cert_3"
        set user-name "username"
        set group-name "group"
    next
end
```
2. Add the SAML user to the user group (optionally, you can configure group matching):```
config user group
    edit "saml_firewall"
        set member "fac-firewall"
        config match
            edit 1
                set server-name "fac-firewall"
                set group-name "user_group1"
            next
        end
    next
end
```
3. Add the SAML user group to a firewall policy:```
config firewall policy
    edit 2
        set srcintf "port3"
        set dstintf "port1"
        set srcaddr "all"
        set dstaddr "pc4"
        set action accept
        set schedule "always"
        set service "ALL"
        set logtraffic all
        set fsso disable
        set groups "saml_firewall" "group_local"
        set users "first"
        set nat enable
    next
end
```
4. Configure the FortiAuthenticator IdP as needed.
5. Run HTTP/HTTPS authentication for a remote user. The SAML login page appears:
