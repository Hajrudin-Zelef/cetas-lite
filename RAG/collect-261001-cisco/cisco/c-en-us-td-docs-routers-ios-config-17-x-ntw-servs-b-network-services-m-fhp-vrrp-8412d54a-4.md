---
id: collect-261001-cisco/cisco/c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp-8412d54a-4
title: "c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp--8412d54a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp--8412d54a.md
source_anchor: ""
source_lines: [429, 452]
sha256: 2e657d8740a9d3f7c6d3efbef471b94bf982a07b3e2de0e294df48ac5d46dd0a
---

# c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp--8412d54a

The following example shows how to configure VRRP text authentication using a text string:
Router(config)#GigabitEthernet 0/0/0interface GigabitEthernet 0/0/0
Router(config)# ip address 10.21.8.32 255.255.255.0
Router(config-if)# vrrp 10 authentication text stringxyz
Router(config-if)# vrrp 10 ip 10.21.8.10
The
Cisco Support and Documentation website provides online resources to download
documentation, software, and tools. Use these resources to install and
configure the software and to troubleshoot and resolve technical issues with
Cisco products and technologies. Access to most tools on the Cisco Support and
Documentation website requires a Cisco.com user ID and password.
The following table provides release information about the feature or features described in this module. This table lists
only the software release that introduced support for a given feature in a given software release train. Unless noted otherwise,
subsequent releases of that software release train also support that feature.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco
Feature Navigator, go to www.cisco.com/go/cfn. An account on Cisco.com is not required.
Glossary
virtualIPaddressowner—The VRRP router that owns the IP address of the virtual router. The owner is the router that has the virtual router address
as its physical interface address.
virtualrouter—One or more VRRP routers that form a group. The virtual router acts as the default gateway router for LAN clients. Also known
as a VRRP group.
virtualrouterbackup—One or more VRRP routers that are available to assume the role of forwarding packets if the virtual primary router fails.
virtualprimaryrouter—The VRRP router that is currently responsible for forwarding packets sent to the IP addresses of the virtual router. Usually
the virtual primary router also functions as the IP address owner.
