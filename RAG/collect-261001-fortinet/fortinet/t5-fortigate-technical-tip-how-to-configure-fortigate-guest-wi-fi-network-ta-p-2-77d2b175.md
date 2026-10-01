---
id: collect-261001-fortinet/fortinet/t5-fortigate-technical-tip-how-to-configure-fortigate-guest-wi-fi-network-ta-p-2-77d2b175
title: "t5-fortigate-technical-tip-how-to-configure-fortigate-guest-wi-fi-network-ta-p-2-77d2b175"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["cybersecurity"]
source: docs/RAG/collect-261001-fortinet/t5-fortigate-technical-tip-how-to-configure-fortigate-guest-wi-fi-network-ta-p-2-77d2b175.md
source_anchor: ""
source_lines: [1, 5]
sha256: 9ef21868476fcce39bba3593478c243a6e676833779b1e7eed23842c0d51ef10
---

# t5-fortigate-technical-tip-how-to-configure-fortigate-guest-wi-fi-network-ta-p-2-77d2b175

| Description | This article describes step-by-step instructions on how to implement a guest network solution based on FortiOS using a bridge mode SSID with HPE Aruba ClearPass version 6.11.6 as a captive portal. | 
| Scope | FortiGate 7.0 and above. | 
| Solution | **Configuration on FortiGate:**   In the GUI:     In the CLI:  **config system interface** edit "VLAN41" set vdom "root" set ip 10.10.41.41 255.255.255.0 set allowaccess ping radius-acct set security-mode captive-portal set security-mac-auth-bypass enable **<----- This is to allow MAC Caching (MAC authentication bypass).** set security-external-web "fqdn-to-clearpass-guest-portal/guest/pagename.php" **<- The URL of the web server hosting login page.** set auth-cert "name-of-ssl-cert" set auth-portal-addr "fqdn-to-dns-name-of-fortigate-vlan-ip" set security-exempt-list "VLAN41-exempt-list" **<----- This should allow http/https, dns, etc. to resolve the ClearPass server and any other exempt destinations/services being allowed before signing on to the guest network.** set security-groups "Guest-Users" set device-identification enable set role lan set snmp-index 21 set interface "port3" set vlanid 41 next end  Login/splash page hosted on an External Web Server:   **Note regarding security-mac-auth-bypass:**     In the GUI:   In the CLI:  **config wireless-controller vap** edit "Bridge_Guest" set ssid "FortiBridgeGuest" set security owe (or open) set pmf enable set mbo enable set local-bridging enable set intra-vap-privacy enable set schedule "always" set vlanid 41 set probe-resp-suppression enable set sticky-client-remove enable next end    **Note:**   It is recommended to use OWE (Opportunistic Wireless Encryption) on the SSID to enable strong encryption between the wireless client and the access point, even with an 'open' network. However, older clients may not support OWE, so alternatives include changing the**security** setting to**open** OR enabling**owe-transition** and specifying a fallback open SSID (see: FortiWiFi & FortiAP Config Guide - OWE Transition)    **Configuration on the ClearPass policy server:**  Two services (MAC Caching and Guest Registration) must be configured on the ClearPass Policy Server.      Note that NAS-IP-Address is where the FortiGate is included in the group.        It is still possible to use the previously used 'Custom' settings, but it is necessary to add the details in the **'Extra Fields'** settings like this:    Note that the **'Submit URL:'** section is the IP or FQDN of the FortiGate Guest Interface. It is highly recommended to use**https port 1003** for the captive portal. However, when using**HTTP** , ensure that the port number being used is**1000** .  If a different port needs to be configured, change it using the following commands and the relevant article linked below:  **config system global** set auth-http-port X (default value is port 1000 for http) set auth-https-port Y (default value is port 1003 for https) end  **Related article** : Technical Tip: Change the captive portal port  **Troubleshooting commands:**  **diagnose debug reset** **diagnose debug disable** **diagnose debug application fnbamd -1** **diagnose debug enable** | 

The Fortinet Security Fabric brings together the concepts of convergence and consolidation to provide comprehensive cybersecurity protection for all users, devices, and applications and across all network edges.
