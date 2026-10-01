---
id: collect-261001-general-networking/general-networking/manual-how-tos-caddy-html-561e74b1-3
title: "Reverse Proxy Domain: \"531e7877-0b58-4f93-a9f0-54beee58bdea\""
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-how-tos-caddy-html-561e74b1.md
source_anchor: ""
source_lines: [280, 366]
sha256: 545e057d96aca9260c4a269ba33208e3ce99b32371b5ede386d513b0b8fbb190
---

# Reverse Proxy Domain: "531e7877-0b58-4f93-a9f0-54beee58bdea"

- Go into the OPNsense WebGUI and restart CrowdSec.
High Availability Setups
There are a few possible configurations to run Caddy successfully in a High Availability Setup with two OPNsense firewalls.
The main issue is the certificate handling. If a CARP VIP is used on the WAN interface, and the A and AAAA Records of all domains point to this CARP VIP, the backup Caddy will not be able to issue ACME certificates without some additional configuration.
There are three methods that support XMLRPC sync:
Note
These methods can be mixed, just make sure to use a coherent configuration. It is best to decide for one method. Only Domains need configuration, Subdomains do not need any configuration for HA.
- Using custom certificates from the OPNsense Trust store for all Domains.
- Using the DNS-01 Challenge in the settings of Domains.
- Using the HTTP-01 Challenge Redirection option in the advanced settings of Domains.
Since the HTTP-01 Challenge Redirection needs some additional steps to work, it should be set up as followed:
- Configure Caddy on the master OPNsense until the whole initial configuration is completed.
- On the master OPNsense, select each Domain, and set the IP Address in HTTP-01 Challenge Redirection to the same value as in Synchronize Config to IP found in .
- Create a new Firewall rule on the master OPNsense that allows Port80 and443 toThis Firewall on the interface that has the prior selected IP Address (most likely a LAN or VLAN interface).
- Sync this configuration with XMLRPC sync.
Now both Caddy instances will be able to issue ACME certificates at the same time. Caddy on the master OPNsense uses the TLS-ALPN-01 challenge for itself and reverse proxies the HTTP-01 challenge to the Caddy of the backup OPNsense. Please make sure, that the master and backup OPNsense are both listening on their WAN and LAN (or VLAN) interfaces on port 80 and 443, since both ports are required for these challenges to work.
Tip
Check the Logfile on both Caddy instances for successful challenges. Look for certificate obtained successfully informational messages.
Forward Auth
Delegating authentication to Authelia or Authentik is a very advanced usecase. The Forward Auth Documentation should be used for inspiration.
To attach the Forward Auth directive to a handler, the Auth Provider has to be filled out in the General Settings. Afterwards, the Forward Auth checkbox in a Handler can be selected in advanced mode. This will prepend the forward_auth directive in front of the reverse_proxy directive in the scope of that Handler. Headers are set automatically.
Using Access Lists and Basic Auth in the Domain this Handler matches on is not recommended.
An example Caddyfile could look like this:
app1.example.com {
    handle {
        forward_auth authelia:9091 {
            uri /api/verify?rd=https://auth.example.com
            copy_headers Remote-User Remote-Groups Remote-Name Remote-Email
        }
        reverse_proxy 192.168.10.1:8080 {
        }
    }
}
Requests from clients to app1.example.com will be sent to Authelia via the forward_auth directive. Then, after the authentication has been completed, the reverse_proxy directive sends the traffic to the Upstream.
Run Caddy Process Unprivileged
In this plugin, Caddy runs as root. This is required when well-known ports are used. Since the default ports are 80 and 443, Caddy will be started as superuser.
For higher security demands, there is the option to run Caddy as www user and group. This comes with the restriction of only being able to use upper ports (≥ 1024).
Make sure all of the domains have empty ports, or ports above the well-known port range before continuing. There is a validation that will prevent configuring well-known ports when the www user is active.
Go to
- Add custom upper HTTP Port, e.g.8080
- Add custom upper HTTPS Port, e.g.8443
- Selectwww as System User
- Restart Caddy completely. Disable it and press Apply, then enable it and press Apply.
From now on, Caddy will run as www user and group. This can be verified by checking the user of the Caddy process.
Note
With this configuration, Destination NAT (Port Forward) should be used to forward port 80 and 443 to the new alternative HTTP and HTTPS Ports. For IPv6 additional steps could be required.
Bind Caddy to Interfaces
Warning
Binding a service to a specific interface via IP address can cause lots of issues. If the IP address is dynamic, the service can crash or refuse to start. During boot, the service can refuse to start if the interface IP addresses are assigned too late. Configuration changes on the interfaces can cause the service to crash. Only use this with static IP addresses! There is no OPNsense community support for this configuration.
This configuration is only useful if there are two or more WAN interfaces, and Caddy should only respond on one of them. It can also solve port conflicts, for example if one interface should DNAT or host a different service with the default webserver ports.
- Create the following files with the following content in the OPNsense filesystem:
- /usr/local/etc/caddy/caddy.d/defaultbind.global
default_bind 203.0.113.1 192.168.1.1
- /usr/local/etc/caddy/caddy.d/defaultbind.conf
http:// {
bind 203.0.113.1 192.168.1.1
}
Now Caddy will only bind to 203.0.113.1 and 192.168.1.1. It can still be configured in the GUI without restrictions.
Read more about the default_bind directive: Default Bind
Custom Configuration Files
- The Caddyfile has an additional import from the path/usr/local/etc/caddy/caddy.d/ . Place custom configuration files inside that adhere to the Caddyfile syntax.
- *.global files will be imported into theglobal block .
- *.conf files will be imported into thesite block .
- *.layer4global and*.layer4listener files will be imported into their respectivelayer4 directive .
- Don’t forget to test the custom configuration withcaddy validate --config /usr/local/etc/caddy/Caddyfile .
With these imports, the full potential of Caddy can be unlocked. The GUI options will remain focused on the reverse proxy. There is no OPNsense community support for configurations that have not been created with the offered GUI. For customized configurations, the Caddy community is the right place to ask.
Caddy: Layer4 Proxy
Enable Layer4 Proxy
- Go to
- Enable the checkbox Enable Layer4 Proxy
- Press Apply, then go to
Routing Type
The implementation has two different modes for layer4 routes:
- listener_wrappers will match traffic on the default HTTP and HTTPS ports that Caddy listens on. With a Layer 7 matcher, selected protocols can be proxied to an upstream. This can be used to multiplex protocols on the default ports and still use the Reverse Proxy at the same time. The most popular usecase is proxying HTTPS without TLS termination. As default route, all unmatched traffic will be sent to the Reverse Proxy.
- global can match any TCP/UDP traffic on any free local port. Additionally, one or multiple Layer 7 matchers can be created under the same protocol port combination. The sequence can be set manually by changing the sequence number.
Routes for both modes can be used at the same time.
Layer 4 Matchers
For the global routing type, the protocol can be set to either TCP or UDP. A local port must be selected, this port must be free and not used by any other service. Port ranges are not supported.
Any IP traffic that matches the port and protocol can proxied to one or multiple upstreams. If raw Layer 4 traffic should be proxied, select ANY as Layer 7 matcher.
Layer 7 Matchers
A Layer 7 matcher checks the first bytes of a TCP/UDP packet and decides which protocol it could be. When TLS or HTTP is detected, they can inspect the contents of the Client Hello at the start of a TLS handshake, or the Host Header in case of HTTP traffic.
There are additional matchers for all kinds of protocols, including:
- DNS
- HTTP (with and without Host Header evaluation)
- OpenVPN
- Postgres
- Proxy Protocol
