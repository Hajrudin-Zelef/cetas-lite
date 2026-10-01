---
id: collect-261001-huawei/huawei/biofects-ha-unifi-speedtest-1b87e159-2
title: "View Home Assistant logs"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "mit license", "parameters"]
source: docs/RAG/collect-261001-huawei/biofects-ha-unifi-speedtest-1b87e159.md
source_anchor: ""
source_lines: [117, 270]
sha256: d0058c6ebfbd5a85d8afd409823c43b6d63db5d7d8e8667ef90ed1ecfe3fff66
---

# View Home Assistant logs

Per-WAN Buttons (created dynamically for each active WAN):
- HA Unifi Speedtest Run Speedtest (WAN - eth9) - Trigger test on specific WAN
- HA Unifi Speedtest Run Speedtest (WAN2 - eth8) - Trigger test on specific WAN
- etc.
Initiates a speed test on your UniFi network.
Parameters:
- config_entry_id (optional): Specific integration instance
- interface_name (optional): Specific WAN interface to test
Manually refreshes speed test data from your UniFi controller.
Parameters:
- config_entry_id (optional): Specific integration instance
In Automations:
action:
  - service: ha_unifi_speedtest.start_speed_test
    data:
      interface_name: "eth9"  # Optional: test specific interface
In Scripts:
test_network_speed:
  sequence:
    - service: ha_unifi_speedtest.start_speed_test
    - delay: "00:02:00"  # Wait for test to complete
    - service: notify.mobile_app
      data:
        message: "Speed test completed. Download: {{ states('sensor.download_speed_wan') }} Mbps"
Lovelace Button Card:
type: button
name: Start Speed Test
icon: mdi:speedometer
tap_action:
  action: call-service
  service: ha_unifi_speedtest.start_speed_test
Multi-WAN Dashboard:
type: vertical-stack
cards:
  - type: entities
    title: Primary WAN (eth9)
    entities:
      - entity: sensor.download_speed_wan
        name: Download Speed
      - entity: sensor.upload_speed_wan
        name: Upload Speed
      - entity: sensor.ping_wan
        name: Ping
      - entity: button.ha_unifi_speedtest_run_speedtest_wan_eth9
        name: Run Speed Test
  
  - type: entities
    title: Secondary WAN (eth8)
    entities:
      - entity: sensor.download_speed_wan2
        name: Download Speed
      - entity: sensor.upload_speed_wan2
        name: Upload Speed
      - entity: sensor.ping_wan2
        name: Ping
      - entity: button.ha_unifi_speedtest_run_speedtest_wan2_eth8
        name: Run Speed Test
Single WAN Dashboard:
type: entities
title: Network Speed Test
entities:
  - entity: sensor.unifi_speed_test_download_speed
    name: Download Speed
  - entity: sensor.unifi_speed_test_upload_speed  
    name: Upload Speed
  - entity: sensor.unifi_speed_test_ping
    name: Ping
  - entity: button.ha_unifi_speedtest_run_speedtest_all_wans
    name: Run Speed Test
Fixes:
- Detects controller-initiated speed tests when UniFi reuses a result ID
- v4.1.0 - Device registry migration and stale device cleanup
- v4.0.0 - UniFi OS API-key-only integration
- v2.2.0 - User-controlled controller type selection
- v2.1.1 - UDM Pro 404 error fixes
- v2.1.0 - Intelligent primary WAN detection
- v2.0.1 - Initial multi-WAN support
Fixes:
- Preserves existing devices, area assignments, and user customizations when upgrading from v3
- Reconciles duplicate devices already created by v4.0.0
- Cleans up inactive WAN entities for interface names such as ppp0 andwan0
- Enables manual removal of stale integration devices
See CHANGELOG.md for complete details.
Invalid Authentication Error:
- Verify your local UniFi OS API key is correct and hasn't expired
  - Generate a new key if needed: Network → Settings → Control Plane → Integrations
- Confirm the key was created through the local Network application, not unifi.ui.com , Site Manager, or Protect
Cannot Connect:
- Check the URL format:
  - UniFi OS console: https://192.168.1.1
  - UniFi OS Server: use its locally configured URL and port
- UniFi OS console: 
- Verify the controller is accessible from Home Assistant
- Check firewall settings allow connections
- Disable "Verify SSL" when connecting by local IP with the default self-signed certificate
- Keep using https:// ; disabling verification does not enable plain HTTP
No Speed Test Data / No Sensors Appear:
- Ensure at least one speed test has been run on your controller
- For multi-WAN: Sensors appear dynamically after speed test data exists
- Use the "Run Speedtest (All WANs)" button to seed initial data
- Check that WANs are active (have link and IP address)
- Enable "Show Inactive WANs" in options if you want to see all WANs
Only One WAN Showing (Expected Multiple):
- Verify multi-WAN is enabled in integration options
- Check that multiple WANs are configured in UniFi controller
- Ensure WANs are active (connected with IP addresses)
- Run a manual speed test to trigger entity creation
- Check Developer Tools → States for all sensor.*wan* entities
API Key Not Working:
- Verify the target runs UniFi OS, not the legacy standalone Network Application
- Ensure the API key was copied completely (no spaces or truncation)
- Check the API key hasn't been revoked in the controller
- Try generating a new API key
Where to Find API Key:
- Local console → Network → Settings → Control Plane → Integrations → "Create API Key"
- Save the key immediately - you can't view it again after creation
Primary WAN Detection Methods (in order of priority):
- Routing Table Analysis: Checks the default route (0.0.0.0/0)
- WAN Group Priority: "WAN" group is preferred as primary
- Network Configuration: Looks for explicit primary WAN settings
- Speed Test Data: Uses most recent and complete data
- Fallback: First detected interface
If Primary WAN Detection is Incorrect:
- 
Check UniFi Network Settings: 
  - Verify routing configuration in UniFi Network application
  - Ensure the desired primary WAN has the default route
- 
Check Integration Logs (enable debug logging below): 
  - Look for: "Primary WAN determined from routing table: eth9_WAN"
  - Shows which detection method was used
- 
Check Sensor Attributes: 
  - Developer Tools → States → Find your speed test sensors
  - Verify is_primary_wan: true on correct interface
Expected Behavior:
- Interface configured as primary in UniFi shows is_primary_wan: true
- Devices named "Primary WAN [WAN - eth9]" and "Secondary WAN [WAN2 - eth8]"
- Each WAN gets separate device with its own sensors
Check for Errors:
# View Home Assistant logs
docker logs homeassistant | grep ha_unifi_speedtest
Common Causes:
- Missing required files (ensure all .py files are present)
- Python import errors (check file permissions)
- Configuration issues (try removing and re-adding integration)
To enable debug logging, add the following to your configuration.yaml:
logger:
  default: info
  logs:
    custom_components.ha_unifi_speedtest: debug
Feel free to contribute to this project. Please read the contributing guidelines before making a pull request.
This project is licensed under the MIT License - see the LICENSE file for details.
This integration is not affiliated with Ubiquiti Inc. or UI.com. All product names, logos, and brands are property of their respective owners.
