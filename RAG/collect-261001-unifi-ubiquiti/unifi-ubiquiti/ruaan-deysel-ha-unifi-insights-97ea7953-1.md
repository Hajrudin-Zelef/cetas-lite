---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/ruaan-deysel-ha-unifi-insights-97ea7953-1
title: "Force an immediate data refresh"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "packaging", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/ruaan-deysel-ha-unifi-insights-97ea7953.md
source_anchor: ""
source_lines: [1, 110]
sha256: 9f52fe67b6714f94d22508a691c05e66c21e1afc44130c99b9672eeb4dfe0f9d
---

# Force an immediate data refresh

A Home Assistant custom integration for monitoring and controlling your UniFi Network and UniFi Protect infrastructure using the official UniFi APIs.
This project overlaps with and complements Home Assistant's official integrations:
| Area | UniFi Insights (this project) | Official integrations | 
|---|---|---|
| Packaging | Single integration covering both Network and Protect in one setup flow | Two separate core integrations ( unifi andunifiprotect ) | 
| Authentication | API key — local and cloud (remote console) connection modes | UniFi Network: local credentials. UniFi Protect: local credentials + API key | 
| Remote management | Supports UniFi cloud console discovery and selection | Primarily local controller connectivity | 
| Service surface | Adds integration-specific services for Network and Protect actions (vouchers, PTZ, chime, light) | Uses Home Assistant core entities and actions per integration | 
| Project lifecycle | Community custom component released via GitHub and HACS | Included in Home Assistant Core release cycle | 
Which to choose:
- Use the official integrations if you prefer core-maintained components with the widest documented feature surface.
- Use UniFi Insights if you want a single integration with API-key-first setup and combined Network and Protect support in one place.
- Running both side-by-side can create overlapping entities. Review and disable duplicates to avoid automation conflicts.
- Device monitoring: CPU, memory, uptime, temperature, and throughput for all adopted devices
- Per-port sensors: PoE power, port speed, link state, SFP module info, TX/RX traffic counters
- Client tracking & control: presence detection, block/allow switches, and reconnect buttons
- WiFi control & QR codes: enable/disable WiFi networks and scan-to-connect QR code images
- Firewall policy control: enable and disable user-defined firewall rules
- Traffic route management: enable and disable policy-based routing rules dynamically
- VPN client control: enable and disable VPN client interfaces (WireGuard, OpenVPN, Privado VPN)
- Firmware update management: check and initiate device upgrades
- Device and port actions: restart devices, power cycle PoE ports
- Guest & voucher services: authorize guests, generate and delete hotspot vouchers
- Camera streaming: live view, snapshots, RTSPS streams
- Motion and smart detection: person, vehicle, animal, and package binary sensors
- Doorbell ring detection
- Camera controls: microphone, privacy mode, status light, high FPS mode
- Protect light control: brightness and mode (always on, motion, off)
- PTZ cameras: move to preset position and run patrol via services
- Chime control: volume, ringtone selection, and repeat settings
- Protect sensor readings: temperature, humidity, light level, battery
- NVR storage monitoring (when available)
- Motion, ring, and smart detection events
- Home Assistant 2026.6.0 or newer
- UniFi Network Application 8.0 or newer
- UniFi Protect 3.0 or newer (required only for Protect features)
- Open HACS in Home Assistant.
- Click Integrations.
- Click the three-dot menu in the top right and select Custom repositories.
- Add https://github.com/ruaan-deysel/ha-unifi-insights as a custom repository with the category set to Integration.
- Search for "UniFi Insights" and install it.
- Restart Home Assistant.
- Download the latest release from the Releases page.
- Copy the custom_components/unifi_insights folder into your Home Assistantcustom_components directory.
- Restart Home Assistant.
- Go to UniFi Site Manager.
- Navigate to Integrations.
- Click Create API Key, give it a name (for example "Home Assistant"), and copy the generated key.
- The API key is shown on screen for you to Copy for the next step.
- In Home Assistant, go to Settings → Devices & Services.
- Click + Add Integration and search for "UniFi Insights".
- Choose your connection type:
  - Local — direct connection to your UniFi console on the local network.
  - Remote — connect via UniFi Cloud. Enter your API key and then select the console from the list of discovered devices.
- For a local connection, enter the host URL (for example https://192.168.1.1 ) and your API key.
- Click Submit.
Remote entries also use the same API key to refresh Site Manager host, site, and device inventory, five-minute ISP metrics, and SD-WAN configuration types every 10 minutes. Entries using the same key share one account refresh. This data is available in the integration's coordinator and as a limited summary in the downloadable diagnostics report; it does not create additional entities. A Site Manager error does not stop the console connection, and collection availability is shown in diagnostics.
After setup, open the integration's options flow (Settings → Devices & Services → UniFi Insights → Configure) to adjust these settings:
| Option | Default | Description | 
|---|---|---|
| Track WiFi Clients | Off | Creates device tracker entities for connected wireless clients. May add a large number of entities on busy networks. | 
| Track Wired Clients | Off | Creates device tracker entities for connected wired clients. | 
| Enable Client Control | On | Creates allow/block switch and reconnect button entities for each connected client. Disable this if you only need read-only monitoring — it prevents orphaned unavailable entities from accumulating when clients leave the network. | 
| Sites | All | Only shown when the console has more than one site. Pick the sites to poll; unselected sites are not queried at all, which cuts API traffic on multi-site consoles. Leave empty to include every site. | 
| Entity | Description | 
|---|---|
| CPU Usage | Device CPU utilization (%) | 
| Memory Usage | Device memory utilization (%) | 
| Uptime | Device uptime | 
| TX Rate | Uplink transmit rate (Mbit/s) | 
| RX Rate | Uplink receive rate (Mbit/s) | 
| Firmware Version | Installed firmware version | 
| Wired Clients | Count of wired clients (switches) | 
| Wireless Clients | Count of wireless clients (access points) | 
| Total Clients | Total client count (site-level) | 
| Port PoE Power | PoE power consumption per switch port (W) | 
| Port Speed | Link speed per port (Mbps) | 
| Port TX / RX | Traffic counters per port (bytes) | 
| Temperature | Protect sensor temperature (°C) | 
| Humidity | Protect sensor humidity (%) | 
| Light Level | Protect sensor ambient light (lux) | 
| Battery | Protect sensor battery level (%) | 
| Storage Used / Total / Available | NVR storage metrics (GB, when available) | 
| Entity | Description | 
|---|---|
| Device Status | Network device online/offline state | 
| WAN Status | Gateway online state (for link state use WAN Connection) | 
| WAN Connection | Per-WAN internet connection state (DHCP, static or PPPoE) as the gateway reports it | 
| Site-to-Site VPN | Connection state of each site-to-site VPN tunnel (one sensor per tunnel) | 
| Motion Detection | Camera or sensor motion activity | 
| Person Detection | AI person detection | 
| Vehicle Detection | AI vehicle detection | 
| Animal Detection | AI animal detection | 
| Package Detection | AI package detection | 
| Doorbell Ring | Doorbell ring activity | 
| Door / Window | Protect sensor open/close state | 
| Tamper | Protect sensor tamper detection | 
| Leak | Protect sensor water leak detection | 
| Recording | Camera actively recording | 
| Entity | Description | 
|---|---|
| WiFi Network | Enable or disable a WiFi broadcast | 
| Firewall Rule | Enable or disable a user-defined firewall policy | 
| Client Allow | Block or allow a connected network client | 
| Traffic Route | Enable or disable a policy-based traffic route | 
| VPN Client | Enable or disable a VPN client interface | 
| Camera Microphone | Enable or disable the camera microphone | 
| Camera Privacy Mode | Enable or disable privacy mode | 
| Camera Status Light | Enable or disable the status LED | 
| Camera High FPS | Enable or disable high frame rate mode | 
