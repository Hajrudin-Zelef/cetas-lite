---
id: collect-261001-general-networking/general-networking/questions-52257-configure-vlan-with-dhcp-traffic-properly-forwarded-76f8955f-3
title: "questions-52257-configure-vlan-with-dhcp-traffic-properly-forwarded-76f8955f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-52257-configure-vlan-with-dhcp-traffic-properly-forwarded-76f8955f.md
source_anchor: ""
source_lines: [691, 855]
sha256: 36516000bd571aa851e79b14fef7db3869a9758f58328f64f6ac97e4a5d11cb4
---

# questions-52257-configure-vlan-with-dhcp-traffic-properly-forwarded-76f8955f

        l3-interface irb.180;
    }
    HSFacBYOD {
        vlan-id 188;
        l3-interface irb.188;
    }
    HSPhone {
        vlan-id 140;
        l3-interface irb.140;
    }
    HSStu {
        vlan-id 170;
        l3-interface irb.170;
    }
    HS_BYOD_Fac {
        vlan-id 133;
        l3-interface irb.133;
    }
    HS_BYOD_Stu {
        vlan-id 132;
        l3-interface irb.132;
    }
    HighSchool {
        vlan-id 4;
        l3-interface irb.4;
    }
    JH {
        vlan-id 3;
        l3-interface irb.3;
    }
    JH107 {
        vlan-id 211;
        l3-interface irb.211;
    }
    JH116 {
        vlan-id 212;
        l3-interface irb.212;
    }
    JHFac {
        vlan-id 280;
        l3-interface irb.280;
    }
    JHFacBYOD {
        vlan-id 288;
        l3-interface irb.288;
    }
    JHOffice {
        vlan-id 213;
        l3-interface irb.213;
    }
    JHStu {
        vlan-id 270;
        l3-interface irb.270;
    }
    JH_BYOD_Fac {
        vlan-id 232;
        l3-interface irb.232;
    }
    JH_BYOD_Stu {
        vlan-id 231;
        l3-interface irb.231;
    }
    JH_Data {
        vlan-id 210;
        l3-interface irb.210;
    }
    JH_Faculty {
        vlan-id 234;
        l3-interface irb.234;
    }
    JH_Student {
        vlan-id 236;
        l3-interface irb.236;
    }
    JH_Wifi_APs {
        vlan-id 230;
        l3-interface irb.230;
    }
    KCISDFac_ES {
        vlan-id 338;
        l3-interface irb.338;
    }
    KCISDFac_HS {
        vlan-id 138;
        l3-interface irb.138;
    }
    KCISDFac_JH {
        vlan-id 238;
        l3-interface irb.238;
    }
    KHFac {
        vlan-id 130;
        l3-interface irb.130;
    }
    KHStu {
        vlan-id 136;
        l3-interface irb.136;
    }
    Scale {
        vlan-id 20;
        l3-interface irb.20;
    }
    SciAutoAGath {
        vlan-id 6;
        l3-interface irb.6;
    }
    VLAN172 {
        vlan-id 172;
        l3-interface irb.172;
    }
    VLAN176 {
        vlan-id 176;
        l3-interface irb.176;
    }
    VideoSur {
        vlan-id 100;
        l3-interface irb.100;
    }
    busbarn {
        vlan-id 9;
        l3-interface irb.9;
    }
    elem456 {
        vlan-id 15;
        l3-interface irb.15;
    }
    elementary {
        vlan-id 2;
        l3-interface irb.2;
    }
    hsmacser {
        vlan-id 8;
        l3-interface irb.8;
    }
    weather {
        vlan-id 111;
        l3-interface irb.111;
    }
}
The juniper at the junior high MDF is on port xe-0/0/23.
root@HSServerRoom-4600> show dhcp relay statistics
Packets dropped:
    Total                      127849
    Bootp packets              126744
    Interface not configured   588
    Bad UDP checksum           1
    No binding found           512
    dhcp-service total         4
Messages received:
    BOOTREQUEST                5571847
    DHCPDECLINE                142
    DHCPDISCOVER               1751818
    DHCPINFORM                 996327
    DHCPRELEASE                99599
    DHCPREQUEST                2723961
Messages sent:
    BOOTREPLY                  3999872
    DHCPOFFER                  968432
    DHCPACK                    3030802
    DHCPNAK                    638
    DHCPFORCERENEW             0
Packets forwarded:
    Total                      5737
    BOOTREQUEST                0
    BOOTREPLY                  5737
