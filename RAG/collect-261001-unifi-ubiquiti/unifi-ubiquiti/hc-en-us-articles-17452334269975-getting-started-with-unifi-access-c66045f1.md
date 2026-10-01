---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-17452334269975-getting-started-with-unifi-access-c66045f1
title: "hc-en-us-articles-17452334269975-getting-started-with-unifi-access-c66045f1"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["ethernet", "license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-17452334269975-getting-started-with-unifi-access-c66045f1.md
source_anchor: ""
source_lines: [1, 72]
sha256: cfec111d8ba617fbafd1de710596445309fcf41782f9bef7e49073c7fa416fb3
---

# hc-en-us-articles-17452334269975-getting-started-with-unifi-access-c66045f1

Getting Started with UniFi Access
UniFi Access is a highly scalable, license-free door access system that is simple to install with just a few PoE-connected components. It features an intuitive interface and is just one component of UniFi's end-to-end IT management platform. Learn more about UniFi
Requirements
Setting up UniFi Access requires just a few simple components:
- A UniFi Console that supports the UniFi Access application.
- A UniFi Access Control Hub whose terminals connect to the door lock for entry and exit.
- A UniFi Access Reader that connects to the hub, allowing users to unlock doors or initiate doorbell calls.
- A compatible door lock.
- 
Power-over-Ethernet switches to power your hubs and readers.
  - Consider using PoE Over 2-Wire Retrofit Extender (UACC-Retrofit-PoE-2Wire) for retrofit scenarios.
Setting Up Your UniFi Access System
- Prepare Your UniFi Console: Ensure that your UniFi Console is set up and running UniFi Access. To view installed UniFi applications, click the Settings icon on the left navigation bar, then select Control Plane.
- 
Connect Your Hubs: Ensure the hub is connected to your UniFi Console's local network and powered according to its requirements. 
  - PoE+ for Access Ultra
  - PoE++ for Door Hub Mini, Door Hub, Elevator Hub, and Gate Hub
  - AC power for Enterprise Access Hub
  - DC power for Retrofit Hub
- 
Adopt Your Readers:
  - Connect your readers to your hub and add them to the UniFi Access application using one of these methods:
    - 
Mobile App
      - Download the UniFi Access mobile app (iOS | Android).
      - Log in with your UI Account.
      - Follow the prompts to adopt your readers, or navigate to the Devices page for manual adoption.
    - 
Site Manager
      - Go to unifi.ui.com and log in with your UI Account.
      - Select your UniFi Console, click the Access icon, and open the Devices tab.
      - Click Adopt to add your readers.
  - 
Mobile App
  - Connect UniFi Retrofit Readers or third-party Wiegand readers to your Retrofit Hub.
    - This two-door solution upgrades your existing access control system without replacing any wiring. Learn more
- Connect your readers to your hub and add them to the UniFi Access application using one of these methods:
- Connect Your Locks to the Hub(s): Power the lock and wire it depending on whether it is fail-safe or fail-secure.
- Add Users to UniFi. Optionally set up UniFi Fabrics to centrally manage users and their permissions across multiple UniFi sites.
- Assign Access Policies and Schedules.
- Configure the Door Unlock Methods available for users (e.g., NFC, PIN, or Mobile App).
- (Optional) Configure Doorbell Call Receivers to manage how authorized personnel handle doorbell alerts, interact with visitors, and unlock doors.
- (Optional) Configure Visitor Schedules for guests visiting a location.
Take Advantage of UniFi Access
Check out some of our articles below to get the most out of your UniFi Access deployment:
FAQs
Can I use my existing Wiegand wiring to connect UniFi Access Control Hubs and Readers?
Yes. The UniFi Retrofit Hub and Retrofit Reader are designed to upgrade your existing access control system without the need to replace any wiring. Setup is simple and helps you modernize your system with minimal effort. Learn more
Is UniFi Access compatible with my existing Wiegand readers?
Yes. The UniFi Retrofit Hub works with both UniFi Retrofit Readers and third-party Wiegand readers. See the list of compatible third-party Wiegand readers
Which network ports are used by UniFi Access and its devices?
UniFi Access uses the following ports for application services, device communication, and device adoption:
- 12443 (HTTPS): Secure communication with UniFi Access devices
- 12812 (MQTT): UniFi Access device messaging (devices only)
- 12080 (HTTP): Localhost access only
- 12455 (HTTPS): OpenAPI access (only if UniFi Access OpenAPI is enabled)
- 8080 (HTTPS): UniFi Access device adoption
- 10001 (UDP): UniFi Access device discovery
How do UniFi Access components communicate with each other?
UniFi Access components communicate using standard Ethernet and IP networking.
- Access Readers communicate directly with Access Control Hubs over Layer 2 (L2) Ethernet.
  - Access Readers and Access Control Hubs must be on the same L2 network and cannot communicate across VLANs.
- Access Control Hubs communicate with the Access application over IP (Layer 3).
  - If the UniFi Console also functions as a gateway (router) and inter-VLAN routing is allowed, Access Control Hubs can be deployed on a different VLAN from the console. To allow communication across VLANs:
    - Ensure routing and firewall rules permit traffic between the VLANs.
    - In the Access application, navigate to Settings > General > Network and select All Networks.
- If the UniFi Console also functions as a gateway (router) and inter-VLAN routing is allowed, Access Control Hubs can be deployed on a different VLAN from the console. To allow communication across VLANs:
What should I be aware of when using third-party switches?
Some third-party switches may restrict or disable broadcast, multicast, or device discovery traffic by default. If these features are blocked, UniFi Access devices may not be discovered or adopted successfully. Ensure that:
- Broadcast and discovery traffic is allowed on ports connected to UniFi Access devices
- Required discovery ports (for example, UDP 10001) are permitted
- Layer 2 discovery traffic is not blocked by switch security features
