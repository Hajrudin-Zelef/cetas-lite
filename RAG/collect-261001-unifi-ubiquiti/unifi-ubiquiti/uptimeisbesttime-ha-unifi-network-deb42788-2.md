---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/uptimeisbesttime-ha-unifi-network-deb42788-2
title: "Install Ruff (user env or venv)"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/uptimeisbesttime-ha-unifi-network-deb42788.md
source_anchor: ""
source_lines: [108, 165]
sha256: b65cc06d58b56bdf40ad70bbfdea98449e7fe6a92dd5146ee5a6445ae8f69d20
---

# Install Ruff (user env or venv)

  - Clients: Present in connected clients list → home , otherwise →not_home
- Clients: Present in connected clients list → 
- Dynamic Entity Creation:
  - Radio sensors only appear for devices that expose radio interface statistics (Access Points, Gateways with Wi-Fi)
  - Port sensors are created for all physical ports on devices with port interfaces (Switches, Gateways)
  - PoE sensors and buttons only appear for ports with PoE capability
  - Client trackers are created for all connected clients and automatically updated as new clients connect
  - Update entities show firmware information for all devices with available firmware data
- Device Capabilities: Different UniFi devices expose different sensor sets based on their hardware capabilities (e.g., switches vs access points vs gateways).
- Entity Organization:
  - Device entities are grouped under their respective UniFi device in the Device Registry
  - Client entities are grouped under their respective client device
  - All entities use the device/client name as the device name with specific sensor names
  - Entity categories are set appropriately (Diagnostic for monitoring, Config for controls)
- Service Usage: Use the unifi_network.remove_stale_clients service to clean up the device registry when clients are no longer present in your network
- 
unifi_network/ : Main integration code
  - Core integration logic, coordinators, and entity platforms
  - Entity platforms: device_tracker.py ,sensor.py ,button.py ,update.py
  - Configuration flow: config_flow.py
  - Data coordinators: coordinator.py
  - Device/client wrappers: unifi_device.py ,unifi_client.py
  - Services: services.py (stale client cleanup)
- 
unifi_network/api_client/ : Generated API client (excluded from linting/formatting)
  - Auto-generated from UniFi Network Integration API OpenAPI specification
  - Models, API endpoints, and type definitions
  - Located in openapi_client_generator/ for regeneration scripts
- 
unifi_network/translations/ : Internationalization files
  - Entity names, configuration flow text
  - Currently supports English (en.json )
- Update interval: Controlled by DEFAULT_UPDATE_INTERVAL inconst.py (30 seconds)
- Platforms: Defined in PLATFORMS inconst.py (sensor, device_tracker, button, update)
- Domain: unifi_network
- Services: remove_stale_clients for device registry cleanup
This is a standard Home Assistant custom component. For development:
- Install in Home Assistant as described above
- Make code changes
- Restart Home Assistant to reload the integration
- Use Home Assistant logs to debug issues
This project uses Ruff for both formatting and linting, aligned with Home Assistant Core's standards:
- Formatting: Black-compatible via ruff format with a line length of 88.
- Imports: Sorted via Ruff's import sorter (I ).
- Lint rules: A pragmatic selection similar to Home Assistant Core (see pyproject.toml ).
- Generated code is excluded from lint/format to avoid churn: openapi_client_generator/ andunifi_network/api_client/ .
Editor setup (VS Code): The workspace sets Ruff as the default Python formatter and organizes imports on save. You can also install the "Ruff" extension by Astral.
Quick commands (optional):
# Install Ruff (user env or venv)
pip install ruff
# Format code
ruff format
# Lint and auto-fix safe issues
ruff check --fix
# Enable pre-commit hooks (recommended)
pip install pre-commit
pre-commit install
This project is not affiliated with Ubiquiti Inc. Use at your own risk.
