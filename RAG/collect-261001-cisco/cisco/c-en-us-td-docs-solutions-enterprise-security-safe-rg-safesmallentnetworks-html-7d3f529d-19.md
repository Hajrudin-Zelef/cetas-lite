---
id: collect-261001-cisco/cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d-19
title: "c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d.md
source_anchor: ""
source_lines: [635, 663]
sha256: 516009c6b06732abe1b2b56c5c289f830b21e316c5010c97bb4de9068ebd142a
---

# c-en-us-td-docs-solutions-enterprise-security-safe-rg-safesmallentnetworks-html-7d3f529d

! Physical interface and Ethernet parametersinterface GigabitEthernet0/0description Connected to cr24-4507-DOno nameifno security-levelno ip address!interface GigabitEthernet0/1description backup to cr24-4507-DOno nameifno security-levelno ip address!! Defines logical redundant interface associated with physical! interfaces. Configures IP and logical interface parameters.interface Redundant1description Connected to cr24-4507-DOmember-interface GigabitEthernet0/0member-interface GigabitEthernet0/1nameif insidesecurity-level 100ip address 10.125.33.10 255.255.255.0authentication key eigrp 100 <removed> key-id 1authentication mode eigrp 100 md5!
Routing
An interior gateway protocol, EIGRP in our configuration examples, is used for dynamic routing. The Internet firewall may participate in routing by learning the internal routes and by injecting a default route pointing to the Internet. The default route should be removed dynamically if the Internet connection becomes unavailable.
As part of the small enterprise network design, two different approaches were validated for the injection of the default route:
•OSPF—The Cisco ASA appliance learns the default route from the Internet border router using OSPF. The default route is then redistributed into EIGRP and from there propagated into the rest of the internal network.
•Static Route—The Cisco ASA appliance is configured with a static default route pointing to the Internet gateway. Object tracking is configured to dynamically remove the default route when the Internet connection becomes unavailable. The default route is redistributed into EIGRP and from there propagated into the rest of the internal network.
Injecting a default route with OSPF requires the configuration of an OSPF process between the Cisco ASA and the Internet border router, as illustrated in Figure 25. If the router is managed by the ISP, the configuration will require coordination with the service provider. This scenario also requires the default route to be propagated over OSPF. The actual default route may originate from the Internet border router itself or somewhere in the ISP network.
Figure 25 OSPF Default Route Injection
The following are the guidelines for using OSPF for the injection of a default route:
•Whenever possible, use MD5 authentication to secure the routing session between the Cisco ASA and the Internet border router.
•Since NAT is configured on the Cisco ASA and the inside address space is not visible outside the firewall, there is no need to redistribute routes from the internal EIGRP into OSPF.
•Route redistribution from OSPF into the internal EIGRP should be limited to the default route only. No other routes should be propagated into EIGRP.
The following configuration snippet illustrates the routing configuration of the Cisco ASA appliance. The configuration includes the route redistribution from OSPF into EIGRP with the enforcement of a route-map allowing only the injection of the default route. MD5 authentication is used for OSPF, and the logging of neighbor status changes is enabled.
! Permit default onlyaccess-list Inbound-Routes standard permit host 0.0.0.0!interface GigabitEthernet0/2ospf message-digest-key 1 md5 <removed>ospf authentication message-digest!route-map Inbound-EIGRP permit 10match ip address Inbound-Routes!router eigrp 100no auto-summarynetwork 10.125.33.0 255.255.255.0passive-interface defaultno passive-interface insideredistribute ospf 200 metric 1000000 2000 255 1 1500 route-map Inbound-EIGRP!router ospf 200network 198.133.219.0 255.255.255.0 area 100area 100 authentication message-digestlog-adj-changes!
Note The hello-interval and dead-interval OSPF timers can be adjusted to detect topological changes faster.
The other validated alternative for the default route injection is the definition of a static default route, which then can be redistributed into the internal EIGRP process. This is shown in Figure 26. This option does not require the configuration of the Internet border router.
Figure 26 Static Default Route with Object Tracking
It is highly recommended to use object tracking so the default route is removed when the Internet connection becomes unavailable. Without object tracking, the default route will be removed only if the outside interface of the appliance goes down. So there is a possibility that the default route may remain in the routing table even if the Internet border router becomes unavailable. To avoid that problem, the static default route can be configured with object tracking. This consists in associating the default route with a monitoring target. The Cisco ASA appliance monitors the target using ICMP echo requests. If an echo reply is not received within a specified time period, the object is considered down and the associated default route is removed from the routing table.
The monitoring target needs to be carefully selected. First, pick one that can receive and respond to ICMP echo requests sent by the Cisco ASA. Second, it is better to use a persistent network object. In the configuration example below the Cisco ASA monitors the IP address of the next hop gateway, which helps identifying if the Internet gateway goes down, but it will not help if the connection is lost upstream. If available, you may want to monitor a persistent network object located somewhere in the ISP network. Static route tracking can also be configured for default routes obtained through DHCP or PPPoE.
In the following configuration the IP address of the next hop gateway (198.133.219.1) is used as the monitoring target. The static default route is then redistributed into EIGRP.
router eigrp 100no auto-summarynetwork 10.125.33.0 255.255.255.0passive-interface defaultno passive-interface insideredistribute static metric 1000000 2000 255 1 1500!route outside 0.0.0.0 0.0.0.0 198.133.219.1 1 track 10!sla monitor 1type echo protocol ipIcmpEcho 198.133.219.1 interface outsidesla monitor schedule 1 life forever start-time now!track 10 rtr 1 reachability
Note The frequency and timeout parameters of object tracking can be adjusted to detect topological changes faster.
Intrusion Prevention Deployment
The small enterprise network design implements Intrusion Prevention using an Advanced Inspection and Prevention Security Services Module (AIP SSM) on the Cisco ASA appliance deployed at the Internet perimeter. This section describes the best practices for integrating and configuring the IPS module for maximum threat control and visibility, as well as the deployment of the IPS Global Correlation feature.
Deploying IPS with the Cisco ASA
The Advanced Inspection and Prevention Security Services Module (AIP SSM) is supported on Cisco ASA 5510 and higher platforms. The AIP SSM runs advanced IPS software that provides proactive, full-featured intrusion prevention services to stop malicious traffic, including worms and network viruses, before they can affect your network.
As described in Intrusion Prevention Guidelines, the AIP SSM may be deployed in inline or promiscuous mode. In inline mode the AIP SSM is placed directly in the traffic flow, while in promiscuous mode the Cisco ASA sends a duplicate stream of traffic to the AIP SSM.
When deploying the AIP SSM in inline mode it is particularly important to determine how traffic will be treated in case of a module failure. The AIP SSM card may be configured to fail open or close when the module becomes unavailable. When configured to fail open, the Cisco ASA appliance allows all traffic through, uninspected, if the AIP SSM becomes unavailable. In contrast, when configured to fail close, the adaptive security appliance blocks all traffic in case of an AIP SSM failure.
The following example illustrates how a Cisco ASA can be configured to divert all IP traffic to the AIP SSM in inline mode and to block all IP traffic if the AIP SSM card fails for any reason:
