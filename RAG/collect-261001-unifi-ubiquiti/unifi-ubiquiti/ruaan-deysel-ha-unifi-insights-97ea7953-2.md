---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/ruaan-deysel-ha-unifi-insights-97ea7953-2
title: "Force an immediate data refresh"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/ruaan-deysel-ha-unifi-insights-97ea7953.md
source_anchor: ""
source_lines: [111, 222]
sha256: e30fbce99bfefd912d0e6ec4475dfc647437699ad4b8894d5e76791190c8ccef
---

# Force an immediate data refresh

| Platform | Description | 
|---|---|
| Button | Restart device, reconnect client, play chime, PTZ patrol start/stop | 
| Camera | Live view, snapshots, RTSPS streaming | 
| Device Tracker | Client presence detection | 
| Event | Motion, doorbell ring, and smart detection events | 
| Image | WiFi QR codes for each broadcast network | 
| Light | Protect floodlight brightness control | 
| Number | Microphone volume, chime volume, light brightness level | 
| Select | Recording mode, HDR mode, video mode, ringtone, PTZ preset, live view | 
| Update | Firmware update management | 
# Force an immediate data refresh
service: unifi_insights.refresh_data
# Restart a network device
service: unifi_insights.restart_device
data:
  site_id: "your-site-id"
  device_id: "device-id"# Set recording mode
service: unifi_insights.set_recording_mode
data:
  camera_id: "camera-id"
  mode: "motion"  # always, motion, smart, never
# Set HDR mode
service: unifi_insights.set_hdr_mode
data:
  camera_id: "camera-id"
  mode: "auto"  # auto, on, off
# Move PTZ camera to a preset position
service: unifi_insights.ptz_move
data:
  camera_id: "camera-id"
  preset: 0  # 0–15
# Start or stop PTZ patrol
service: unifi_insights.ptz_patrol
data:
  camera_id: "camera-id"
  action: "start"  # start, stop
  slot: 0  # 0–15# Set Protect light mode
service: unifi_insights.set_light_mode
data:
  light_id: "light-id"
  mode: "motion"  # always, motion, off
# Set Protect light brightness
service: unifi_insights.set_light_level
data:
  light_id: "light-id"
  level: 50  # 0–100# Play a ringtone on a chime
service: unifi_insights.play_chime_ringtone
data:
  chime_id: "chime-id"
  ringtone_id: "default"  # default, mechanical, digital, christmas, traditional
# Set chime volume
service: unifi_insights.set_chime_volume
data:
  chime_id: "chime-id"
  volume: 50  # 0–100# Authorize a guest client
service: unifi_insights.authorize_guest
data:
  site_id: "your-site-id"
  client_id: "client-id"
  duration_minutes: 480
# Generate a hotspot voucher
service: unifi_insights.generate_voucher
data:
  site_id: "your-site-id"
  count: 1
  duration_minutes: 480
The integration ships a Lovelace card that draws each site's network: the gateway, switches and access points, and the clients connected to them. It is installed with the integration — there is no separate download and no dashboard resource to add.
Add it: edit a dashboard → Add card → search for UniFi Insights Topology. With a single UniFi site the card works without any configuration; with several, pick the site in the card editor.
Use it:
- Graph view: drag to pan, pinch or Ctrl + scroll to zoom (plain scroll
zooms in panel view), and use the zoom buttons or + /- /0 keys.
Select a device to see its uplink port, link speed, PoE draw and client
counts, with a link to its Home Assistant device page.
- List view: the same network as an indented list with search. It is the recommended view for screen readers.
- Clients are grouped under their switch or access point ("12 clients"); select a group to expand it.
- The filter buttons hide gateways, switches, access points, clients or other devices; devices under a hidden one stay visible, linked with a dashed line.
Options (all optional, all available in the card editor):
| Option | Values | Default | 
|---|---|---|
| entry_id +site_id | the site to show | the only site, if there is one | 
| title | text | the site name | 
| view | graph ,list | graph | 
| show_site_selector | true ,false | false | 
| clients | collapsed ,expanded ,hidden | collapsed | 
| kinds | any of gateway ,switch ,access_point ,client ,other | all | 
| density | comfortable ,compact | automatic (compact on narrow cards) | 
| orientation | vertical ,horizontal | vertical | 
| show_labels | true ,false | true | 
| max_clients | 1–500 | 500 | 
The card only receives what the topology API sends — names, models, states, ports and VLANs. MAC addresses, IP addresses, hostnames and firmware versions are never included.
Add the following to your configuration.yaml and restart Home Assistant:
logger:
  default: info
  logs:
    custom_components.unifi_insights: debug
| Problem | Solution | 
|---|---|
| Cannot connect | Verify the host URL is reachable from Home Assistant. For self-signed certificates, disable SSL verification in the integration options. | 
| Authentication failed | Confirm the API key is valid and was not revoked in the UniFi Site Manager. | 
| No Protect entities | UniFi Protect must be running on the same console. Verify your API key has access to it. | 
| Entities missing | Confirm the devices are adopted and online in the UniFi controller. | 
| Storage sensors unavailable | The public Protect API does not expose NVR storage data on all firmware versions. | 
| Many orphaned client entities | Disable the Enable Client Control option in the integration's settings. | 
To download a sanitized diagnostic report for troubleshooting:
- Go to Settings → Devices & Services.
- Select UniFi Insights.
- Click the three-dot menu and choose Download diagnostics.
The report is sanitized before it is written: API keys, passwords and Wi-Fi secrets, host names and IP addresses, SSIDs, and the names of your clients are redacted, and every MAC address is replaced with a placeholder that stays consistent within a single report. Device, site and camera names are kept so the report remains readable.
Contributions are welcome. Please open an issue before submitting significant changes. See CONTRIBUTING.md for guidelines.
This project is licensed under the Apache License 2.0. See the LICENSE file for details.
This integration is not affiliated with or endorsed by Ubiquiti Inc. Use at your own risk.
