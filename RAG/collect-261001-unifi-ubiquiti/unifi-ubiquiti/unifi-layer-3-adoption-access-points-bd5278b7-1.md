---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/unifi-layer-3-adoption-access-points-bd5278b7-1
title: "unifi-layer-3-adoption-access-points-bd5278b7"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["data centre"]
source: docs/RAG/collect-261001-unifi-ubiquiti/unifi-layer-3-adoption-access-points-bd5278b7.md
source_anchor: ""
source_lines: [1, 60]
sha256: 1bb089551adba6130864c43cd308a52e9396a06da230822fe6b424870ef2ca44
---

# unifi-layer-3-adoption-access-points-bd5278b7

When an access point and your controller sit on the same network, adoption is a single click: the device shows up on its own, you press “Adopt”, and you are done. The moment your controller lives somewhere else – a different VLAN, another site, or a hosted cloud controller in a data centre – the device no longer finds it. That is where Layer 3 adoption comes in. This guide walks through the available methods, the ports they depend on, and the mistakes that most often break the process.
Layer 2 or Layer 3 – which applies to you?
Adoption is the process of connecting a UniFi device – an access point, switch, camera or sensor – to the UniFi application that will manage it. Only after adoption can the device be configured, updated and monitored. A brand-new or factory-reset access point actively looks for a controller and reports in as Pending Adoption.
That search relies on broadcast discovery inside the local network. Devices on the same subnet therefore usually appear in the controller without you doing anything at all – that is Layer 2 adoption. Put a routing boundary in between and the broadcast never arrives, so you need Layer 3 adoption. Ubiquiti describes it as a process for advanced users, or for adopting devices to a UniFi Cloud Console.
| Situation | Method | 
|---|---|
| AP and controller in the same subnet | Layer 2 – device appears automatically | 
| AP in a management VLAN, controller in another VLAN | Layer 3 | 
| Controller at a different site | Layer 3 | 
| Hosted cloud controller in a data centre | Layer 3 | 
Prerequisite: get the ports right
Before you try any method, check your network permissions. None of the approaches below work through a closed port. Between the UniFi host and your UniFi devices, TCP 8080 and UDP 10001 must be open across all gateways, firewalls and antivirus software. For Layer 3 adoption specifically, unrestricted connectivity over TCP port 8080 is the critical requirement – that is the port the device uses to inform the controller.
| Port | Protocol | Purpose | 
|---|---|---|
| 8080 | TCP | Inform port – device to controller communication | 
| 10001 | UDP | Device discovery on the local network | 
| 8443 | TCP | Controller web interface and API | 
| 3478 | UDP | STUN | 
One detail trips people up regularly: the inform URL uses HTTP, not HTTPS. The default format is http://unifi:8080/inform. Guess your way to https here and the device will never check in.
The Layer 3 adoption methods
Ubiquiti documents several approaches. Which one fits depends on how many devices you are rolling out and how much control you have over the local network.
1. SSH and set-inform
The most direct route for individual devices. SSH is enabled by default on UniFi network devices, so you can connect right after a factory reset. You do need to be on the same local network as the access point.
- Find the access point’s IP address – for example from your router’s DHCP leases.
- Connect to the device over SSH. Before adoption, the default credentials are ui /ui , orubnt /ubnt on older devices.
- Run the command on the device:set-inform http://<ip-or-hostname>:8080/inform
- The device then shows up in the remote controller and can be adopted there just like a standard Layer 2 adoption.
Worth knowing: after adoption the device’s SSH credentials are no longer ui/ui but a random string of characters assigned by the controller.
2. DHCP option 43
The method of choice for larger rollouts: your DHCP server tells each UniFi device where the controller lives. On most third-party firewalls, option 43 is encoded in hex, and there are two ways to express it:
- By IP address: start with 0104 and append the static IP encoded in hex. 192.168.3.10 becomes0104c0a8030a – no spaces.
- By FQDN: start with 02 , then the length of the FQDN in bytes (hex), then the inform URL converted to hex.
For added redundancy, both the IP and the FQDN can be specified in the same DHCP option. Either way, the device and the UniFi Network application still need unrestricted connectivity over TCP port 8080.
3. A DNS record
The tidiest option if you control local DNS. UniFi devices query the hostname unifi by default. Point a unifi A record at your controller’s IP – or a CNAME at its FQDN – and devices find the remote controller with no per-device work. That makes it a genuine set-and-forget solution: every new device lands in the right place automatically.
4. The UniFi mobile app
The iOS and Android app can streamline adoption if your network is already broadcasting WiFi on the same VLAN as the unadopted device. For the occasional straggler, it is often the quickest path.
5. Zero-touch provisioning (ZTP)
ZTP is Ubiquiti’s answer to the overhead of manual adoption: plug the device in and it adopts itself, with no admin clicking “Adopt”. That removes a lot of friction when adopting across VLANs and remote networks. Support is still limited, though – according to Ubiquiti, ZTP is currently supported by U7 Pro Max access points, with more devices to follow.
One pitfall to keep in mind: ZTP codes remain valid unless the device is reset using the physical hardware reset button. To preserve the code, always perform a factory reset through the UniFi management interface rather than the button on the housing. And if you already run a local UniFi Cloud Gateway or another UniFi console on site, ZTP offers no additional benefit over the standard adoption process.
Methods compared
| Method | Best for | Requires | 
|---|---|---|
| SSH / set-inform | Single devices, troubleshooting | Local network access, SSH credentials | 
| DHCP option 43 | Larger rollouts | Control over the DHCP server | 
| DNS record “unifi” | A permanent default | Control over local DNS | 
| Mobile app | Individual devices on site | WiFi on the same VLAN as the device | 
| ZTP | Remote rollouts without an engineer on site | A supported model (currently U7 Pro Max) | 
After adoption: overriding the inform host
Once a device is adopted, it needs to keep reaching the controller reliably. The UniFi application offers an option to override the inform host with a specific hostname or IP. When it is set, the controller uses exactly that address and overrides the devices’ existing inform URLs.
Two things to watch. First, name resolution for that hostname has to work properly, otherwise your devices lose contact. Second, editing the configuration files on the devices themselves gets you nowhere – the controller simply overwrites those changes. The setting belongs in the controller, not on the access point.
When it goes wrong: common stumbling blocks
- The device never appears: verify that TCP 8080 and UDP 10001 really are open between the device and the UniFi host, including local firewalls and antivirus software.
- “Managed by other”: the device is still adopted by a different controller. It has to be released there or reset.
- The device stops responding: hold the reset button for 10 seconds; the LED should then turn steady white. As a last resort, factory reset the device and start the Layer 3 adoption again.
- Wrong protocol in the inform URL: it is http:// , nothttps:// , and the port defaults to 8080.
Conclusion
Layer 3 adoption is not difficult, but it demands care in three places: open ports, a correct inform URL, and a properly configured inform host after adoption. For one-off devices, set-inform over SSH is the fastest route; for larger environments, DHCP option 43 or a DNS record pays off as a permanent solution. If you would rather not run the controller yourself – including availability, updates and backups – talk to us about a managed UniFi cloud controller hosted in a German, GDPR-compliant data centre, with a stable address your access points can keep informing to.
Frequently asked questions
What is the difference between Layer 2 and Layer 3 adoption?
