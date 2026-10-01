---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-administration-guide-826729-explicit-proxy-authenticati-13403a20-2
title: "Explicit proxy authentication"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-administration-guide-826729-explicit-proxy-authenticati-13403a20.md
source_anchor: ""
source_lines: [163, 319]
sha256: 209263eb988d2df39149855cd15d4403c7b2e224c944529d3f6adf1354d5d75e
---

# Explicit proxy authentication

1. 
                                                    Create an authentication scheme: 
  1. 
                                                            Go to *Policy & Objects > Authentication Rules* .
  2. 
                                                            Click *Create New > Authentication Schemes* .
  3. 
                                                            Set the *Name* to*Auth-scheme-Negotiate* and select*Negotiate* as the*Method* .
  4. 
                                                            Click *OK* .
2. 
                                                            
3. 
                                                    Create an authentication rule: 
  1. 
                                                            Go to *Policy & Objects > Authentication Rules* .
  2. 
                                                            Click *Create New > Authentication Rules* .
  3. 
                                                            Set the *Name* to*Auth-Rule* ,*Source Address* to*all* , and*Protocol* to*HTTP* .
  4. 
                                                            Enable *Authentication Scheme* , and select the just created*Auth-scheme-Negotiate* scheme.
  5. 
                                                            Click *OK* .
4. 
                                                            

###### To create an authentication scheme and rules in the CLI:

1. 
                                                    Create an authentication scheme: ```
config authentication scheme
    edit "Auth-scheme-Negotiate"
        set method negotiate       <<< Accepts both Kerberos and NTLM as fallback
    next
end
```
2. 
                                                    Create an authentication rule: ```
config authentication rule
    edit "Auth-Rule"
        set status enable
        set protocol http
        set srcaddr "all"
        set ip-based enable
        set active-auth-method "Auth-scheme-Negotiate"
        set comments "Testing"
    next
end
```

###### To create an explicit proxy policy and assign a user group to it in the GUI:

1. 
                                                    Go to *Policy & Objects > Proxy Policy* .
2. 
                                                    Click *Create New* .
3. 
                                                    Set *Proxy Type* to*Explicit Web* and*Outgoing Interface* to*port1* .
4. 
                                                    Set *Source* to*all* , and the just created user groups*NTLM-FSSO-Group* and*Ldap-Group* .
5. 
                                                    Also set *Destination* to*all* ,*Schedule* to*always* ,*Service* to*webproxy* , and*Action* to*ACCEPT* .
6. 
                                                    Click *OK* .

###### To create an explicit proxy policy and assign a user group to it in the CLI:

```
config firewall proxy-policy
    edit 1
        set proxy explicit-web
        set dstintf "port1"
        set srcaddr "all"
        set dstaddr "all"
        set service "web"
        set action accept
        set schedule "always"
        set logtraffic all
        set groups "NTLM-FSSO-Group" "Ldap-Group"
        set av-profile "av"
        set ssl-ssh-profile "deep-custom"
    next
end
```
                                            
                                            Log in using a domain and system that would be authenticated using the Kerberos server, then enter the `diagnose wad user list` CLI command to verify:

```
# diagnose wad user list
ID: 8, IP: 10.1.100.71, VDOM: vdom1
  user name   : test1@FORTINETQA.LOCAL
  duration    : 389
  auth_type   : IP
  auth_method : Negotiate
  pol_id      : 1
  g_id        : 1
  user_based  : 0
  expire      : no
  LAN:
    bytes_in=4862 bytes_out=11893
  WAN:
    bytes_in=7844 bytes_out=1023
```
                                            Log in using a system that is not part of the domain. The NTLM fallback server should be used:

```
# diagnose wad user list
ID: 2, IP: 10.1.100.202, VDOM: vdom1
  user name   : TEST31@FORTINETQA
  duration    : 7
  auth_type   : IP
  auth_method : NTLM
  pol_id      : 1
  g_id        : 5
  user_based  : 0
  expire      : no
  LAN:
    bytes_in=6156 bytes_out=16149
  WAN:
    bytes_in=7618 bytes_out=1917
```
                                            
                                            A keytab is used to allow services that are not running Windows to be configured with service instance accounts in the Active Directory Domain Service (AD DS). This allows Kerberos clients to authenticate to the service through Windows Key Distribution Centers (KDCs).

For an explanation of the process, see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/ktpass.

###### To generate a keytab on a Windows server:

1. 
                                                    On the server, create a user for the FortiGate: 
  - 
                                                            The service name is the FQDN for the explicit proxy interface, such as the hostname in the client browser proxy configuration. In this example, the service name is *FGT* .
  - 
                                                            The account only requires *domain users* membership.
  - 
                                                            The password must be very strong.
  - 
                                                            The password is set to never expire.
2. 
                                                            
3. 
                                                    Add the FortiGate FQDN in to the Windows DNS domain, as well as in-addr.arpa.
4. 
                                                    Generate the Kerberos keytab using the `ktpass` command on Windows servers and many domain workstations:# ktpass -princ HTTP/<domain name of test fgt>@realm -mapuser <user> -pass <password> -crypto all -ptype KRB5_NT_PRINCIPAL -out fgt.keytab For example: ktpass -princ HTTP/FGT.FORTINETQA.LOCAL@FORTINETQA.LOCAL -mapuser FGT -pass *********** -crypto all -ptype KRB5_NT_PRINCIPAL -out fgt.keytab If the FortiGate is handling multiple keytabs in Kerberos authentication, use different passwords when generating each keytab.
5. 
                                                    Encode the keytab to base64 in a text file: 
  - 
                                                            On Windows: `certutil -encode fgt.keytab tmp.b64 && findstr /v /c:- tmp.b64 > fgt.txt`
  - 
                                                            On Linux: `base64 fgt.keytab > fgt.txt`
  - 
                                                            On MacOS: `base64 -i fgt.keytab -o fgt.txt`
6. 
                                                            
7. 
                                                    Use the code in `fgt.txt` as the keytab parameter when configuring the FortiGate.
