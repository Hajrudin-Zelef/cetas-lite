---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-5-administration-guide-150448-861edde2
title: "Troubleshooting for DNS filter"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2023-10-28"]
keywords: ["latency", "license"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-5-administration-guide-150448-861edde2.md
source_anchor: ""
source_lines: [1, 86]
sha256: 581652e1ca2a3d5a96f983c62d32d005a6b2870e7d98b14443cb9721ca0e1253
---

# Troubleshooting for DNS filter

# Troubleshooting for DNS filter

If you have trouble with the DNS filter profile in your policy, start with the following troubleshooting steps:

- 
                                                    Check the connection between the FortiGate and FortiGuard DNS rating server (SDNS server).
- 
                                                    Check that the FortiGate has a valid FortiGuard web filter license.
- 
                                                    Check the FortiGate DNS filter configuration.

You need to ensure the FortiGate can connect to the FortiGuard SDNS server. By default, the FortiGate uses DNS over TLS (DoT, TCP port 853) to connect to the SDNS server. See DNS over TLS and HTTPS for more information.

###### To check the connection between the FortiGate and SDNS server:

1. 
                                                    Verify the FortiGuard SDNS server information: # diagnose test application dnsproxy 3 ... **SDNS servers:
173.243.140.53:853** vrf=0 tz=-480 encrypt=dot req=0 to=0 res=0 rt=34 ready=1 timer=0 probe=0 failure=0 last_failed=0The SDNS server IP address might be different depending on location (in this example, it is 173.243.140.53:853).
2. 
                                                    In the management VDOM, check the communication between the FortiGate and the SDNS server: # execute ping 173.243.140.53
3. 
                                                    If FortiGuard is not reachable using anycast, configure the default FortiGuard SDNS (unicast) server (208.91.112.220): ```
config system fortiguard
    set fortiguard-anycast disable
    set sdns-server-ip "208.91.112.220"
end
```
4. 
                                                    Verify the list of SDNS servers again: # diagnose test application dnsproxy 3 FGD_DNS_SERVICE_LICENSE: server=208.91.112.220:53, expiry=2023-10-28, expired=0, type=2 server=83.231.212.53:53, expiry=2023-10-28, expired=0, type=2 The default FortiGuard SDNS server should work in most cases; however, you can switch to another server to see if it improves latency.

|  | By default, DNS filtering connects to the FortiGuard secure DNS server over anycast and uses DoT (TCP port 853) when the default settings of `fortiguard-anycast enable` and`fortiguard-anycast-source fortinet` are configured. Disabling`fortiguard-anycast` will force the FortiGate to use cleartext (UDP port 53) instead of DoT (TCP port 853) in addition to disabling FortiGuard secure DNS over anycast. | 

The FortiGuard DNS rating service shares the license with the FortiGuard web filter, so you must have a valid web filter license for the DNS rating service to work. While the license is shared, the DNS rating service uses a separate connection mechanism from the web filter rating.

###### To check the DNS rating service license in the CLI:

1. 
                                                    View the DNS settings: # diagnose test application dnsproxy 3
2. 
                                                    Find the `FGD_DNS_SERVICE_LICENSE` line and check that the license has not expired:FGD_DNS_SERVICE_LICENSE: server=173.243.140.53:853, expiry=2023-10-28, **expired=0** , type=2
3. 
                                                    Find the `SDNS servers` line to view the functioning servers:SDNS servers: 173.243.140.53:853 vrf=0 tz=-480 encrypt=dot req=0 to=0 res=0 rt=34 **ready=1** timer=0 probe=0 failure=0 last_failed=0

###### To check the DNS filter profile configuration:

1. 
                                                    In FortiOS, create a local domain filter and set the *Action* to*Redirect to Block Portal* (see Local domain filter).
2. 
                                                    Apply this DNS filter profile to the policy.
3. 
                                                    From the client PC, perform a DNS query on this domain. If you get the profile's redirected portal address, this means that the DNS filter profile works as expected.

## Additional troubleshooting

Use  `diagnose test application dnsproxy <test level>` to troubleshoot further DNS proxy information, where:

| Test level | Action | 
|---|---|
| 1 | Clear DNS cache | 
| 2 | Show statistics | 
| 3 | Dump DNS setting | 
| 4 | Reload FQDN | 
| 5 | Requery FQDN | 
| 6 | Dump FQDN | 
| 7 | Dump DNS cache | 
| 8 | Dump DNS database | 
| 9 | Reload DNS database | 
| 10 | Dump secure DNS policy/profile | 
| 11 | Dump botnet domain | 
| 12 | Reload secure DNS setting | 
| 13 | Show hostname cache | 
| 14 | Clear hostname cache | 
| 15 | Show SDNS rating cache | 
| 16 | Clear SDNS rating cache | 
| 17 | Show DNS debug bit mask | 
| 18 | Show DNS debug object members | 
| 99 | Restart the dnsproxy worker | 

###### To debug DNS proxy details:

```
# diagnose debug application dnsproxy -1
# diagnose debug {enable | disable}
```
