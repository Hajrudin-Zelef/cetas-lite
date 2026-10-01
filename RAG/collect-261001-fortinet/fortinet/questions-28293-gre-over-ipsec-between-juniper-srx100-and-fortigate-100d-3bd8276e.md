---
id: collect-261001-fortinet/fortinet/questions-28293-gre-over-ipsec-between-juniper-srx100-and-fortigate-100d-3bd8276e
title: "questions-28293-gre-over-ipsec-between-juniper-srx100-and-fortigate-100d-3bd8276e"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-28293-gre-over-ipsec-between-juniper-srx100-and-fortigate-100d-3bd8276e.md
source_anchor: ""
source_lines: [1, 195]
sha256: 4141f5382d43de0a650a01213da68b422050e3908aacf429cf487cd5af23f5fb
---

# questions-28293-gre-over-ipsec-between-juniper-srx100-and-fortigate-100d-3bd8276e

I need assistance configuring VPN setup between Fortigate and Juniper devices (GRE over IPSec).
I have the topology as follows:
What I need to do is create a route-based IPSec tunnel between the devices to encapsulate another tunnel (GRE tunnel) within. I had similar setup between this Juniper and Cisco 2921 and it worked just fine, however, I am having real trouble setting this one up as I don't have necessary expertise with either.
Don't know if it is supposed to work that way, but...
- I want IPSec tunnel to be between 203.0.113.1 and 203.0.113.2 public addresses;
- I want GRE tunnel to initiate from loopback interface and communicate to remote endpoint's loopback (10.255.255.1 and 10.255.255.2) which I have routes for to reach via IPSec tunnel (st0.0 and ipsec interfaces accordingly);
- I want GRE interfaces to have IP addresses 192.168.0.1 and 192.168.0.2 respectively (and have OSPF running between these, configuration not included).
Juniper side
Interface configuration:
interface fe-0/0/0 {
    unit 0 {
        family inet {
            address 203.0.113.1/24;
        }
    }
}
interface lo0 {
    unit 0 {
        family inet {
            address 10.255.255.1/32;
        }
    }
}
interface st0
    unit 0 {
        family inet;
    }
}
interface gr-0/0/0 {
    unit 0 {
        clear-dont-fragment-bit;
        tunnel {
            source 10.255.255.1;
            destination 10.255.255.2;
            allow-fragmentation;
        }
        family inet {
            mtu 1400;
            address 192.168.0.1/30;
        }
    }
}
IKE configuration (edit security ike):
proposal ike-prop {
    authentication-method pre-shared-keys;
    dh-group group5;
    authentication-algorithm sha1;
    encryption-algorithm aes-256-cbc;
    lifetime-seconds 28800;
}
policy ike-policy {
    mode main;
    proposals test-ike-prop;
    pre-shared-key ascii-text "test999";
}
gateway ike-gw {
    ike-policy test-ike-policy;
    address 203.0.113.2;
    dead-peer-detection;
    nat-keepalive 10;
    local-identity inet 203.0.113.1;
    remote-identity inet 203.0.113.2;
    external-interface fe-0/0/0.0;
    version v1-only;
}
IPSec configuration (edit security ipsec):
proposal test-ipsec-prop {
    protocol esp;
    authentication-algorithm hmac-sha1-96;
    encryption-algorithm aes-256-cbc;
    lifetime-seconds 3600;
}
policy ipsec-policy {
    perfect-forward-secrecy {
        keys group5;
    }
    proposals ipsec-prop;
}
vpn ipsec-vpn {
    bind-interface st0.0;
    df-bit clear;
    ike {
        gateway ike-gw;
        proxy-identity {
            local 10.255.255.1/32;
            remote 10.255.255.2/32;
            service junos-gre;
        }
        ipsec-policy ipsec-policy;
    }
    establish-tunnels immediately;
}
Route configuration:
routing-options {
    static {
        route 10.255.255.2/32 next-hop st0.0;
        #Default route exists, but not included
    }
}
Firewall configuration:
Firewall configuration not included, however, all traffic is permitted as follows:
From GRE to IPSec interface
From IPSec to GRE interface
From Wan1 to IPSec interface
From IPSec to Wan1 interface
From Loopback to IPSec interface
From IPsec to Loopback interface
From GRE to Protected network
From Protected network to GRE
Fortigate side
Interface configuration:
config system settings
    set allow-subnet-overlap enable
end
config system interface
    edit "wan1"
        set vdom "root"
        set ip 203.0.113.2 255.255.255.0
        set allowaccess ping
        set type physical
        set weight 1
        set alias "WAN"
    next
    edit "ipsec"
        set vdom "root"
        set ip 203.0.113.2 255.255.255.255
        set type tunnel
        set remote-ip 203.0.113.1
        set interface "wan1"
    next
    edit "gre"
        set vdom "root"
        set ip 192.168.0.2 255.255.255.255
        set allowaccess ping ssh
        set type tunnel
        set remote-ip 192.168.0.1
        set mtu 1400
        set interface "ipsec"
    next
end
config system gre-tunnel
    edit "gre"
        set interface "ipsec"
        set local-gw 203.0.113.2
        set remote-gw 203.0.113.1
    next
end
IKE configuration:
config vpn ipsec phase1-interface
    edit "ike"
        set interface "wan1"
        set proposal aes256-sha1
        set peertype one
        set mode aggressive
        set localid "203.0.113.2"
        set remote-gw 203.0.113.1
        set peerid "203.0.113.1"
        set psksecret test999
    next
end
IPSec configuration:
config vpn ipsec phase2-interface
    edit "ipsec"
        set dst-addr-type ip
        set keepalive enable
        set phase1name "ike"
        set proposal aes256-sha1
        set protocol 47
        set src-addr-type ip
        set dst-start-ip 203.0.113.1
        set keylifeseconds 3600
        set src-start-ip 203.0.113.2
    next
end
Route configuration:
config router static
    edit 2
        set device ipsec
        set dst 10.255.255.1 255.255.255.255
    next
end
Firewall configuration:
Firewall configuration not included, however, all traffic is permitted as follows:
From GRE to IPSec interface
From IPSec to GRE interface
From Wan1 to IPSec interface
From IPSec to Wan1 interface
From Loopback to IPSec interface
From IPsec to Loopback interface
From GRE to Protected network
From Protected network to GRE
Please, advise if this kind of setup could at all work and what could be wrong here? I played with local/remote identities and could get IPSec to come up, show up in security-association table, however, never could get GRE interface to come up and have pingable remote address. Neither I could ping remote loopback via IPSec tunnel.
If any debug/output needed to troubleshoot just post in comments. Thanks!
Maybe you have some good resource on setting this up? Was not able to find anything on Fortinet/Juniper websites that would be helpful to setup between these two vendors. You can definitely find ones to setup with Cisco.
Note: addresses used are not public, but allocated for use in example code: https://www.rfc-editor.org/rfc/rfc5735
