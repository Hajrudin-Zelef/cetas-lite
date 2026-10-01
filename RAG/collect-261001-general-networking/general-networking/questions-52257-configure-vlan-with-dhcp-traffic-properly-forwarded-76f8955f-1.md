---
id: collect-261001-general-networking/general-networking/questions-52257-configure-vlan-with-dhcp-traffic-properly-forwarded-76f8955f-1
title: "questions-52257-configure-vlan-with-dhcp-traffic-properly-forwarded-76f8955f"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2018-07-26"]
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/questions-52257-configure-vlan-with-dhcp-traffic-properly-forwarded-76f8955f.md
source_anchor: ""
source_lines: [1, 350]
sha256: 870fefe04b6ba989def2b38c04e0dc8f6bffea4c8d132eb4db8c56c737e56a6f
---

# questions-52257-configure-vlan-with-dhcp-traffic-properly-forwarded-76f8955f

I have a windows DHCP server in VLAN 10 and I have wireless clients in VLAN 288. The VLANs are defined in my Juniper EX-4600 and the VLANs are tagged/trunked all the way to the Meraki APs.
[DHCP]---[EX-4600]---[EX-4300]--[HP5130]---[Meraki APs]
There are VLANs that are working and VLANs that do not and they all seem to have the same configuration with the exception of the subnet.
The only thing that I have found in searching for a solution was on the juniper KB that said to configure IP helpers, but none of the other VLANs have that and they work, so I don't think that is the solution.
## Last commit: 2018-07-26 01:20:45 UTC by root
version 14.1X53-D40.8;
system {
    processes {
        dhcp-service {
            traceoptions {
                file dhcp_logfile size 10m;
                level all;
                flag all;
            }
        }
        app-engine-virtual-machine-management-service {
            traceoptions {
                level notice;
                flag all;
            }
        }
    }
}
chassis {
    redundancy {
        graceful-switchover;
    }
}
interfaces {
    interface-range VLAN4 {
        member-range ge-1/0/5 to ge-1/0/11;
        unit 0 {
            family ethernet-switching {
                vlan {
                    members 4;
                }
            }
        }
    }
    xe-0/0/23 {
        ether-options {
            flow-control;
        }
        unit 0 {
            family ethernet-switching {
                interface-mode trunk;
                vlan {
                    members [ 1 3 10 100 140-141 172 176 210-213 230-232 234 236 238 270 280 288 ];
                }
                storm-control default;
            }
        }
    }
    em1 {
        unit 0 {
            family inet;
        }
    }
    irb {
        unit 0 {
            family inet;
        }
        unit 1 {
            family inet {
                address 10.1.1.100/24;
            }
        }
        unit 2 {
            family inet {
                address 192.168.11.1/24;
            }
        }
        unit 3 {
            family inet {
                address 192.168.10.1/24;
            }
        }
        unit 4 {
            family inet {
                address 192.168.9.1/24;
            }
        }
        unit 5 {
            family inet {
                address 192.168.15.1/24;
            }
        }
        unit 6 {
            family inet {
                address 192.168.35.1/24;
            }
        }
        unit 8 {
            family inet {
                address 192.168.36.1/24;
            }
        }
        unit 9 {
            family inet {
                address 192.168.37.1/24;
            }
        }
        unit 10 {
            family inet {
                address 192.168.75.2/24;
            }
        }
        unit 11 {
            family inet {
                address 192.168.76.1/24;
            }
        }
        unit 15 {
            family inet {
                address 192.168.8.1/24;
            }
        }
        unit 20 {
            family inet {
                address 10.0.0.1/24;
            }
        }
        unit 100 {
            family inet {
                address 192.168.100.1/24;
            }
        }
        unit 111 {
            family inet {
                address 192.168.111.6/24;
            }
        }
        unit 130 {
            family inet {
                address 10.1.30.1/23;
            }
        }
        unit 132 {
            family inet {
                address 10.1.32.1/24;
            }
        }
        unit 133 {
            family inet {
                address 10.1.33.1/24;
            }
        }
        unit 136 {
            family inet {
                address 10.1.36.1/22;
            }
        }
        unit 138 {
            family inet {
                address 10.1.56.1/21;
            }
        }
        unit 140 {
            family inet {
                address 10.1.40.1/24;
            }
        }
        unit 141 {
            family inet {
                address 172.16.41.142/29;
            }
        }
        unit 170 {
            family inet {
                address 10.1.72.1/22;
            }
        }
        unit 171 {
            family inet {
                address 10.1.76.1/24;
            }
        }
        unit 172 {
            family inet {
                address 172.26.88.10/21;
            }
        }
        unit 176 {
            family inet {
                address 172.16.41.177/29;
            }
        }
        unit 180 {
            family inet {
                address 10.1.80.1/23;
            }
        }
        unit 188 {
            family inet {
                address 10.1.88.1/21;
            }
        }
        unit 210 {
            family inet {
                address 10.41.10.1/24;
            }
        }
        unit 211 {
            family inet {
                address 10.41.11.1/24;
            }
        }
        unit 212 {
            family inet {
                address 10.41.12.1/24;
            }
        }
        unit 213 {
            family inet {
                address 10.41.13.1/24;
            }
        }
        unit 230 {
            family inet {
                address 10.41.30.1/24;
            }
        }
        unit 231 {
            family inet {
                address 10.41.31.1/24;
            }
        }
        unit 232 {
            family inet {
                address 10.41.32.1/24;
            }
        }
        unit 234 {
            family inet {
                address 10.41.34.1/23;
            }
        }
        unit 236 {
            family inet {
                address 10.41.36.1/23;
            }
        }
        unit 238 {
            family inet {
                address 10.41.56.1/21;
            }
        }
        unit 270 {
            family inet {
                address 10.41.72.1/22;
            }
        }
        unit 271 {
            family inet {
                address 10.41.76.1/24;
            }
        }
        unit 280 {
            family inet {
                address 10.41.80.1/23;
            }
        }
        unit 288 {
            family inet {
                address 10.41.88.1/21;
            }
        }
        unit 311 {
            family inet {
                address 10.103.11.1/24;
            }
        }
        unit 331 {
            family inet {
                address 10.103.31.1/24;
            }
        }
        unit 332 {
            family inet {
                address 10.103.32.1/24;
            }
        }
        unit 334 {
            family inet {
                address 10.103.34.1/23;
            }
        }
        unit 336 {
            family inet {
                address 10.103.36.1/23;
            }
        }
        unit 338 {
            family inet {
                address 10.103.56.1/21;
            }
        }
        unit 370 {
            family inet {
                address 10.103.72.1/22;
            }
        }
        unit 371 {
            family inet {
                address 10.103.76.1/24;
            }
        }
        unit 380 {
            family inet {
                address 10.103.80.1/23;
            }
        }
        unit 388 {
            family inet {
                address 10.103.88.1/21;
            }
        }
    }
    vme {
        unit 0 {
            family inet;
        }
    }
}
snmp {
    name HSServerRoom-4600;
    description HSServerRoom-4600;
    location "HS Server Room";
    contact "jgaspard@kirbyvillecisd.org";
    client-list list0 {
        192.168.75.0/24;
    }
    community public {
        authorization read-only;
        client-list-name list0;
    }
    trap-group kcisd-traps {
        destination-port 162;
        targets {
            192.168.75.29;
        }
    }
}
forwarding-options {
    storm-control-profiles default {
        all;
    }
    analyzer {
        INTERNET_IN {
            input {
