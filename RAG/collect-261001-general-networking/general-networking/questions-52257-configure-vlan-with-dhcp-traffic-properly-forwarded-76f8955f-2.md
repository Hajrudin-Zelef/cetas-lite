---
id: collect-261001-general-networking/general-networking/questions-52257-configure-vlan-with-dhcp-traffic-properly-forwarded-76f8955f-2
title: "questions-52257-configure-vlan-with-dhcp-traffic-properly-forwarded-76f8955f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-52257-configure-vlan-with-dhcp-traffic-properly-forwarded-76f8955f.md
source_anchor: ""
source_lines: [351, 690]
sha256: b8541d39a6e54046ff766436d957ab5f2b5ada1a63bbcdbb090b16eca00e7516
---

# questions-52257-configure-vlan-with-dhcp-traffic-properly-forwarded-76f8955f

                ingress {
                    interface ge-1/0/0.0;
                }
                egress {
                    interface ge-1/0/0.0;
                }
            }
            output {
                interface ge-1/0/2.0;
            }
        }
    }
    dhcp-relay {
        server-group {
            dhcp_servers {
                192.168.75.12;
            }
        }
        group group1 {
            active-server-group dhcp_servers;
            interface irb.1;
            interface irb.2;
            interface irb.3;
            interface irb.4;
            interface irb.5;
            interface irb.6;
            interface irb.8;
            interface irb.9;
            interface irb.10;
            interface irb.11;
            interface irb.15;
            interface irb.100;
            interface irb.111;
            interface irb.130;
            interface irb.132;
            interface irb.133;
            interface irb.136;
            interface irb.138;
            interface irb.170;
            interface irb.180;
            interface irb.188;
            interface irb.210;
            interface irb.211;
            interface irb.212;
            interface irb.213;
            interface irb.230;
            interface irb.231;
            interface irb.232;
            interface irb.234;
            interface irb.236;
            interface irb.238;
            interface irb.270;
            interface irb.280;
            interface irb.288;
            interface irb.311;
            interface irb.331;
            interface irb.332;
            interface irb.334;
            interface irb.338;
            interface irb.370;
            interface irb.380;
            interface irb.388;
        }
    }
}
routing-options {
    nonstop-routing;
    static {
        route 0.0.0.0/0 next-hop 192.168.75.13;
        route 172.25.0.18/32 next-hop 172.16.41.182;
        route 172.25.0.19/32 next-hop 172.16.41.182;
    }
    router-id 192.168.100.1;
}
protocols {
    ospf {
        area 0.0.0.0 {
            interface irb.10;
            interface irb.1 {
                passive;
            }
            interface irb.4 {
                passive;
            }
            interface irb.2 {
                passive;
            }
            interface irb.5 {
                passive;
            }
            interface irb.6 {
                passive;
            }
            interface irb.8 {
                passive;
            }
            interface irb.9 {
                passive;
            }
            interface irb.15 {
                passive;
            }
            interface irb.111 {
                passive;
            }
            interface irb.100 {
                passive;
            }
            interface irb.172 {
                passive;
            }
            interface irb.141 {
                passive;
            }
            interface irb.11 {
                passive;
            }
            interface irb.130 {
                passive;
            }
            interface irb.311 {
                passive;
            }
            interface irb.140 {
                passive;
            }
            interface irb.133 {
                passive;
            }
            interface irb.132 {
                passive;
            }
            interface irb.176 {
                passive;
            }
            interface irb.331 {
                passive;
            }
            interface irb.332 {
                passive;
            }
            interface irb.334 {
                passive;
            }
            interface irb.136 {
                passive;
            }
            interface irb.20 {
                passive;
            }
            interface irb.212 {
                passive;
            }
            interface irb.3 {
                passive;
            }
            interface irb.211 {
                passive;
            }
            interface irb.213 {
                passive;
            }
            interface irb.230 {
                passive;
            }
            interface irb.231 {
                passive;
            }
            interface irb.232 {
                passive;
            }
            interface irb.234 {
                passive;
            }
            interface irb.236 {
                passive;
            }
            interface irb.210 {
                passive;
            }
            interface irb.138 {
                passive;
            }
            interface irb.238 {
                passive;
            }
            interface irb.338 {
                passive;
            }
            interface irb.170 {
                passive;
            }
            interface irb.180 {
                passive;
            }
            interface irb.188 {
                passive;
            }
            interface irb.288 {
                passive;
            }
            interface irb.388 {
                passive;
            }
            interface irb.380 {
                passive;
            }
            interface irb.370 {
                passive;
            }
            interface irb.270 {
                passive;
            }
            interface irb.280 {
                passive;
            }
        }
    }
    lldp {
        interface all;
    }
    lldp-med {
        interface all;
    }
    igmp-snooping {
        vlan default;
    }
    inactive: rstp {
        interface xe-0/0/0;
        interface xe-0/0/1;
        interface xe-0/0/2;
        interface xe-0/0/3;
        interface xe-0/0/4;
        interface xe-0/0/5;
        interface xe-0/0/6;
        interface xe-0/0/7;
        interface xe-0/0/8;
        interface xe-0/0/9;
        interface xe-0/0/10;
        interface xe-0/0/11;
        interface xe-0/0/12;
        interface xe-0/0/13;
        interface xe-0/0/14;
        interface xe-0/0/15;
        interface xe-0/0/16;
        interface xe-0/0/17;
        interface xe-0/0/18;
        interface xe-0/0/19;
        interface xe-0/0/20;
        interface xe-0/0/21;
        interface xe-0/0/22;
        interface xe-0/0/23;
        interface et-0/0/24;
        interface xe-0/0/24:0;
        interface xe-0/0/24:1;
        interface xe-0/0/24:2;
        interface xe-0/0/24:3;
        interface et-0/0/25;
        interface xe-0/0/25:0;
        interface xe-0/0/25:1;
        interface xe-0/0/25:2;
        interface xe-0/0/25:3;
        interface et-0/0/26;
        interface xe-0/0/26:0;
        interface xe-0/0/26:1;
        interface xe-0/0/26:2;
        interface xe-0/0/26:3;
        interface et-0/0/27;
        interface xe-0/0/27:0;
        interface xe-0/0/27:1;
        interface xe-0/0/27:2;
        interface xe-0/0/27:3;
    }
}
virtual-chassis {
    preprovisioned;
    member 0 {
        role routing-engine;
        serial-number TC3716430102;
    }
    member 1 {
        role routing-engine;
        serial-number PE3716350055;
    }
    member 2 {
        role line-card;
        serial-number PE3716350466;
    }
}
vlans {
    Backbone {
        vlan-id 10;
        l3-interface irb.10;
    }
    Central {
        vlan-id 5;
        l3-interface irb.5;
    }
    DEFAULT_VLAN {
        vlan-id 1;
        l3-interface irb.1;
    }
    DLVIDEO {
        vlan-id 141;
        l3-interface irb.141;
    }
    ES4thLab {
        vlan-id 311;
        l3-interface irb.311;
    }
    ESFac {
        vlan-id 380;
        l3-interface irb.380;
    }
    ESFacBYOD {
        vlan-id 388;
        l3-interface irb.388;
    }
    ESStu {
        vlan-id 370;
        l3-interface irb.370;
    }
    ES_BYOD {
        vlan-id 331;
        l3-interface irb.331;
    }
    ES_BYOD_Fac {
        vlan-id 332;
        l3-interface irb.332;
    }
    ES_Main_Wifi {
        vlan-id 334;
        l3-interface irb.334;
    }
    Fog {
        vlan-id 11;
        l3-interface irb.11;
    }
    HSFac {
        vlan-id 180;
