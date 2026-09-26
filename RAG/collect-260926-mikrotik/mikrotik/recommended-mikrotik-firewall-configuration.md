---
id: collect-260926-mikrotik/mikrotik/recommended-mikrotik-firewall-configuration
title: "Securing MikroTik"
domain: mikrotik
role: reference
task: reference
actors: ["AWS", "Google", "Stripe"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/recommended-mikrotik-firewall-configuration.md
source_anchor: ""
source_lines: [1, 101]
sha256: af2579e0baf9ea1519624c58d16543475a75a6bcc7ea0487dde2e9a92bc91f54
---

# Securing MikroTik

If you want to make your MikroTik device more secure, you might consider using a rule like this:

```
/ip firewall filter
add action=drop chain=input comment="bottom drop" disabled=yes
```

This rule drops **all incoming traffic** on the **input chain** from any source across all interfaces. However, this also prevents redirection to the Splash Page and blocks access to essential services.

##  Minimum Requirement

To safely use such a rule, you **must allow traffic** from the following hosts on the **WireGuard/OpenVPN** and **WAN** interfaces to maintain connectivity with the Powerlynx infrastructure:

```
/ip firewall address-list
add address=172.16.0.0/12 list="allowed pool"
add address=10.112.0.0/16 list="allowed pool"
add address=64.227.114.153 list="allowed pool"
add address=157.230.77.58 list="allowed pool"
add address=5.101.109.44 list="allowed pool"
```

**Allowing these hosts ensures that users can obtain a voucher from a free data plan.**

## Examples

These IPs are essential for Powerlynx services and must be allowed through your firewall.


##  Important Notes

Allowing these hosts is critical for the following features to work properly:

- 
**Free Data Plan access**
- 
**Login using Voucher Code**

If you plan to let customers **purchase vouchers** through payment gateways (e.g. **PayPal**, **Stripe** and etc), you also need to allow access to additional hosts associated with those services.

### Example for PayPal:

Here are some common entries used in the **Walled Garden** (this is not a complete list):

```
/ip hotspot walled-garden
add dst-host=*.digitaloceanspaces.com
add dst-host=*paypal*
add dst-host=*betacdn.net
add dst-host=*.powerlynx.app
add dst-host=*.akamaiedge.net
```

These entries use wildcards to allow access to broader domains hosted on content delivery networks (like Amazon, Azure, or Google Cloud).

**Important:** These hosts should be added to both the **Walled Garden** and the **Firewall**.



##  Debugging Tip

To identify which connections are being blocked, you can temporarily enable logging:

```
/ip firewall filter
add action=log chain=input comment="Log remaining input before drop" log=yes log-prefix="DROP-INPUT "
```

This helps identify which protocols, ports, and source addresses attempted to send traffic to your router on the ‘INPUT’ chain, and this information may help you reconfigure your router’s firewall.

**Therefore, at a minimum, you need to allow input chain traffic from hosts related to the Powerlynx infrastructure, as well as from hosts associated with your payment gateway. This is essential for the hotspot to function correctly if you intend to use a drop rule.**

**As an alternative, you can review your firewall filter rules for security and consider reconfiguring them. As a basic reference, you can check the default MikroTik firewall filter configuration.**

```
/ip firewall {                                                                                                                                                  
                       filter add chain=input action=accept connection-state=established,related,untracked comment="defconf: accept established,related,untracked"                   
                       filter add chain=input action=drop connection-state=invalid comment="defconf: drop invalid"                                                                   
                       filter add chain=input action=accept protocol=icmp comment="defconf: accept ICMP"                                                                             
                       filter add chain=input action=accept dst-address=127.0.0.1 comment="defconf: accept to local loopback (for CAPsMAN)"                                          
                       filter add chain=input action=drop in-interface-list=!LAN comment="defconf: drop all not coming from LAN"                                                     
                       filter add chain=forward action=accept ipsec-policy=in,ipsec comment="defconf: accept in ipsec policy"                                                        
                       filter add chain=forward action=accept ipsec-policy=out,ipsec comment="defconf: accept out ipsec policy"                                                      
                       filter add chain=forward action=fasttrack-connection connection-state=established,related comment="defconf: fasttrack"                                        
                       filter add chain=forward action=accept connection-state=established,related,untracked comment="defconf: accept established,related, untracked"                
                       filter add chain=forward action=drop connection-state=invalid comment="defconf: drop invalid"                                                                 
                       filter add chain=forward action=drop connection-state=new connection-nat-state=!dstnat in-interface-list=WAN comment="defconf: drop all from WAN not DSTNATed"
                     }  
```

In this configuration:

- The **WAN interface is less secure** , but it’s simpler and doesn’t require manual whitelisting.
- You can enhance this with **port knocking** ,**anti-port scanning rules** , or a**honeypot redirect** , but use these only if necessary and with caution.


##  Note

Make sure **hotspot dynamic rules** are located **at the top** of your firewall rule list. Otherwise, other rules may interfere with hotspot functionality. (Refer to the previous topic for more details.)
