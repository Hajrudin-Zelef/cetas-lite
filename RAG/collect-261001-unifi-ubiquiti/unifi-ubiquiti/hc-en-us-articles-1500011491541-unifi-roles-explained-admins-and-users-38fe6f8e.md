---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-1500011491541-unifi-roles-explained-admins-and-users-38fe6f8e
title: "hc-en-us-articles-1500011491541-unifi-roles-explained-admins-and-users-38fe6f8e"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-1500011491541-unifi-roles-explained-admins-and-users-38fe6f8e.md
source_anchor: ""
source_lines: [1, 22]
sha256: 0a9c1524dec5ab187e53f7bc063fd50b895933f0a6f518fc920e20e54121a661
---

# hc-en-us-articles-1500011491541-unifi-roles-explained-admins-and-users-38fe6f8e

UniFi Roles Explained: Admins and Users/Members
UniFi uses a role-based access control (RBAC) model that separates Admins, who manage UniFi infrastructure, from Users and Members, who interact with UniFi services such as WiFi, VPN, and Door Access. How these roles and permissions are managed depends on whether you are using UniFi Fabrics for centralized, multi-site access control or operating individual sites.
UniFi Fabrics & Roles
UniFi Fabrics introduce a centralized way to manage people, roles, and permissions across multiple sites. Instead of assigning access on a per-site basis, Fabrics allow roles to be defined once and applied consistently across all applicable sites. Furthermore, it enables zero-trust network security through the use of UniFi Endpoint, which supports secure SAML SSO when using services like one-click WiFi, one-click VPN, or smart door access.
For an overview of Fabrics, check out this UniFi Academy Topic.
To learn how roles and permissions work in Fabrics, see UniFi Fabrics: Managing People, Roles, and Permissions.
Admins (Single-Site / Non-Fabric)
Admins are accounts that can access the UniFi management interface from Site Manager, or directly from the offline, local interface.
Common admin roles include:
- Owner: The account that originally set up a UniFi Console and has the highest level of access.
- Super Admin: Full administrative access to a site. Super Admins can perform actions like backup restoration and configuring SSH, but some features in UniFi OS and the UniFi mobile app are exclusive to the owner. Super Admins should be considered equivalent to Owners from a security perspective.
- Hotspot Operator (Network): Manage guest WiFi hotspots.
- Door Attendant (Access): Open doors and communicate with visitors.
To learn more about adding Admins, click here.
Users (Single-Site / Non-Fabric)
Users are people configured in UniFi to interact with UniFi services as part of daily operations but do not access the UniFi management interface.
Examples include:
- Connecting to WiFi
- Accessing a VPN
- Unlocking doors
- Using other UniFi services assigned to them
User permissions are granted by Admins, and determine which services they can use. To learn about creating Members, click here.
