---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-8-administration-guide-300428-18a509cc
title: "Explicit web proxy"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-8-administration-guide-300428-18a509cc.md
source_anchor: ""
source_lines: [1, 131]
sha256: 235e524ba9f53302ead2d8b4f83f4de0ed20adb5cb95f01dabab41fff04a3917
---

# Explicit web proxy

# Explicit web proxy

Explicit web proxy can be configured on FortiGate for proxying HTTP and HTTPS traffic.

To deploy explicit proxy, individual client browsers can be manually configured to send requests directly to the proxy, or they can be configured to download proxy configuration instructions from a Proxy Auto-Configuration (PAC) file.

When explicit proxy is configured on an interface, the interface IP address can be used by client browsers to forward requests directly to the FortiGate. FortiGate also supports PAC file configuration.

|  | For FortiOS 7.4.0, SSL VPN web mode, explicit web proxy, and interface mode IPsec VPN features will not work with the following configuration:  Configuring an IP pool as the source NAT IP address in a regular firewall policy works as before. For details, see Technical Tip: IP pool and virtual IP behaviour changes in FortiOS 6.4, 7.0, 7.2, and 7.4. | 

###### To configure explicit web proxy in the GUI:

1. 
                                                    Enable and configure explicit web proxy: 
  1. 
                                                            Go to *Network > Explicit Proxy* .
  2. 
                                                            Enable *Explicit Web Proxy* .
  3. 
                                                            Select *port2* as the*Listen on Interfaces* and set the*HTTP Port* to*8080* .
  4. 
                                                            Configure the remaining settings as needed.
  5. 
                                                            Click *Apply* .
2. 
                                                            
3. 
                                                    Create an explicit web proxy policy: 
  1. 
                                                            Go to *Policy & Objects > Proxy Policy* .
  2. 
                                                            Click *Create New* .
  3. 
                                                            Set *Proxy Type* to*Explicit Web* and*Outgoing Interface* to*port1* .
  4. 
                                                            Also set *Source* and*Destination* to*all* ,*Schedule* to*always* ,*Service* to*webproxy* , and*Action* to*ACCEPT* .
  5. 
                                                            Click *OK* to create the policy.
4. 
                                                            
5. 
                                                    Configure a client to use the FortiGate explicit proxy: Set the FortiGate IP address as the proxy IP address in the browser, or use an automatic configuration script for the PAC file.

|  | This example creates a basic policy. If required, security profiles can be enabled, and deep SSL inspection can be selected to inspect HTTPS traffic. | 

###### To configure explicit web proxy in the CLI:

1. Enable and configure explicit web proxy:```
config web-proxy explicit
    set status enable
    set ftp-over-http enable
    set socks enable
    set http-incoming-port 8080
    set ipv6-status enable
    set unknown-http-version best-effort
end
config system interface
    edit "port2"
        set vdom "vdom1"
        set ip 10.1.100.1 255.255.255.0
        set allowaccess ping https ssh snmp http telnet
        set type physical
        set explicit-web-proxy enable
        set snmp-index 12
        end
    next
end
```
2. Create an explicit web proxy policy:```
config firewall proxy-policy
    edit 1
        set name "proxy-policy-explicit"
        set proxy explicit-web
        set dstintf "port1"
        set srcaddr "all"
        set dstaddr "all"
        set service "webproxy"
        set action accept
        set schedule "always"
        set logtraffic all
    next
end
```
3. Configure a client to use the FortiGate explicit web proxy:Set the FortiGate IP address as the proxy IP address in the browser, or use an automatic configuration script for the PAC file.

|  | This example creates a basic policy. If required, security profiles can be enabled, and deep SSL inspection can be selected to inspect HTTPS traffic. | 

## Downloading a PAC file using HTTPS

PAC files can be downloaded for an explicit proxy through the FortiGate's captive portal using HTTPS to ensure a secure download.

In this example, a Windows PC has an HTTPS URL configured in its proxy settings to download a PAC file from a FortiGate by using a download link, https://cp.myqalab.local:7831/proxy.pac, through a captive portal. Once the PAC file is securely downloaded using HTTPS, browsers installed on the PC can use the proxy in the PAC file to visit a website.

The global web proxy settings must be configured to use a customized SSL certificate because the default Fortinet_Factory certificate will not be accepted by Windows due to security restrictions. The customized SSL certificate is used as the HTTPS server's certificate on the FortiGate. All CA certificates in the server certificate must be installed and trusted on the Windows PC.

###### To download a PAC file using HTTPS:

1. 
                                                    Configure the explicit web proxy to get a PAC file through HTTPS: ```
config web-proxy explicit
    set pac-file-server-status enable
    unset pac-file-server-port
    set pac-file-name "proxy.pac"
    set pac-file-data "function FindProxyForURL(url, host) {
   // testtest
   return \"PROXY 10.1.100.1:8080\";
}
"
    
```
**set pac-file-through-https enable** end
2. 
                                                    Configure the captive portal to be used as an HTTPS server to provide the service to download the PAC file: ```
config authentication setting
    set captive-portal-type ip
    
```
**set captive-portal-ip 10.1.100.1
    set captive-portal-ssl-port 7831** end
3. 
                                                    Configure the global web proxy settings to use a customized SSL certificate: ```
config web-proxy global
    set ssl-cert "server_cert"
end
```
4. 
                                                    On the Windows PC, go to *Settings > Network & Internet > Proxy* .
5. 
                                                    In the *Automatic proxy setup* section, click*Save* to trigger the PAC file download from the HTTPS URL.
