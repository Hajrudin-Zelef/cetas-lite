---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/lazro-securing-reolink-ip-camera-systems-nvr-vlan-isolation-with-unifi-os-4-4-9-da21d48c-2
title: "lazro-securing-reolink-ip-camera-systems-nvr-vlan-isolation-with-unifi-os-4-4-9--da21d48c"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/lazro-securing-reolink-ip-camera-systems-nvr-vlan-isolation-with-unifi-os-4-4-9--da21d48c.md
source_anchor: ""
source_lines: [155, 337]
sha256: 81569492794ca9f201395a3ae62f85d0b19b5db82f20aaa84d462a2ae762b168
---

# lazro-securing-reolink-ip-camera-systems-nvr-vlan-isolation-with-unifi-os-4-4-9--da21d48c

Most Reolink cameras connect via PoE directly to the NVR and require no additional configuration. However, if you have WiFi-enabled cameras (battery-powered, plug-in WiFi, or PoE cameras used wirelessly):
Step 1: Initial Camera Setup
Initialize your WiFi camera on your network first:
- Follow Reolink’s initial setup process for your camera
- Connect the camera to your WiFi network
- Set a strong password for the camera
- Verify the camera is functioning properly
For detailed setup instructions, refer to: How to Initially Set up Reolink Devices
Step 2: Add Camera to NVR via LAN
Important: Both the camera and NVR must be on the same network for discovery. Since your NVR is now on the Camera VLAN, you need to temporarily assign the camera to the Camera VLAN as well.
- Click Client Devices on the left sidebar in UniFi
- Find and select the WiFi camera
- Click Settings in the right-hand menu
- Under IP Settings:
- Check Virtual Network Override
- Set Network to Camera
- Check Fixed IP Address (recommended)
- Assign a static IP address
5. Click Apply
Step 3: Auto-Add Camera to NVR
- Log in to your Reolink NVR
- Navigate to Channel → Channel Management
- Turn on Auto Add
- All cameras on the same LAN (Camera VLAN) will automatically appear in the list
- You may see “Incorrect account name or password” — this is normal
Step 4: Enter Camera Credentials
- Click Modify next to the camera
- Enter the login password you set for the camera during initial setup
- Click OK
- The camera should now appear with live view on the NVR
Step 5: Assign to Channel
- Select the camera you want to add
- Choose which channel to assign it to
- Click Next to complete the addition
Note: Each camera must be on a unique channel. You cannot have two cameras on the same channel.
Important Considerations for Battery-Powered WiFi Cameras
NVR Takes Over Management: Once a battery-powered WiFi camera is added to the NVR:
- The NVR takes full control of the camera
- Standalone access through the Reolink App/Client is no longer possible
- All viewing and management must be done through the NVR
To Restore Standalone Access: If you need standalone access to the battery-powered camera again (without NVR):
- A hard reset of the camera is required
- The camera will need to be set up from scratch
- You’ll lose the NVR integration
For complete instructions on adding WiFi cameras to Reolink NVRs, see: How to Add Reolink IP Cameras to Reolink WiFi NVR
Part 5: Configuring Firewall Policy
Create IP-Based Policy for NVR Access
- Navigate to Settings → Policy Engine
- Click Create New Policy
- Select Firewall
- Configure:
- Name: Internal to Reolink NVR
- Source Zone: Internal
- Source Networks: Select all networks that need access (Apple, Windows VLANs, etc.)
- Source Port: Any
- Action: Allow
- Auto Allow Return Traffic: Checked
- Destination Zone: IP Address
- Destination IP: Your NVR’s static IP address
- Destination Port: 443,9000
- Protocol: All
5. Click Save
Understanding the Ports
We specify only the ports configured in the NVR settings:
Get Lazro’s stories in your inbox
Join Medium for free to get updates from this writer.
Port 443 (TCP) — HTTPS:
- Secure web interface access
- Required for browser-based management
- Must be enabled in NVR settings (Part 1, Step 2)
Port 9000 (TCP) — Basic Service Port:
- Reolink client software connection
- Default port for client application access
- Configured in NVR settings (Part 1, Step 2)
Note: If you changed the default ports in the NVR configuration (Part 1, Step 2), update the port numbers in this policy to match your custom settings. For example, if you set HTTPS to port 8443 instead of 443, use 8443,9000 in the policy.
Still IP-Based for Maximum Security
We’re targeting the NVR’s specific IP address rather than the entire Camera network. This means:
- Only the NVR is accessible from approved VLANs
- WiFi cameras on the Camera VLAN remain isolated
- Adding cameras doesn’t expand the accessible surface
- Clear, maintainable security posture
Part 6: Testing and Verification
Verify NVR Access from Approved Networks
From Windows PC (on Windows VLAN):
- Open web browser
- Navigate to: https://nvr.local (or static IP address)
- Should see Reolink login page
- Log in with NVR credentials
- Test live view from cameras
- Test playback of recorded footage
From macOS/iOS (on Apple VLAN):
- Open Reolink app on iPhone/iPad
- NVR should be accessible
- Test live view
- Test playback
- Test remote notifications (if enabled)
From Any Approved Network:
- Access NVR web interface
- Verify all cameras appear and are recording
- Test playback controls
- Check motion detection settings
- Verify recording schedules
Verify Internet Connectivity
Check NVR Can Reach Internet:
- Log into NVR web interface
- Navigate to Settings → System → About
- Check for firmware updates — should be able to check for updates
- If available, test downloading an update
Test Remote Access (If Using Reolink Cloud):
- Disconnect from your home network (use cellular data)
- Open Reolink app
- NVR should be accessible remotely
- Test live view over internet
- Verify push notifications work
Verify Time Synchronization:
- Check NVR system time
- Should match actual time (synced via NTP)
- Important for accurate recording timestamps
Verify Isolation
Test That Other VLANs Cannot Access NVR:
- Devices on VLANs NOT listed in the Source Networks should be unable to reach the NVR
- Try accessing NVR from a guest network or other isolated VLAN
- Should timeout or be blocked
Check Camera VLAN Cannot Access Other Networks:
- The NVR should not be able to access devices on your main networks
- Isolation prevents the NVR from scanning or accessing other network resources
- This is automatic due to the “Isolate Network” setting
Part 7: Troubleshooting
Cannot Access NVR Web Interface
Possible Causes:
- Firewall policy not active or misconfigured
- Wrong IP address in policy
- NVR didn’t pick up static IP
- Source network not included in policy
Solutions:
- Verify required ports are in firewall policy (443,9000 if using defaults)
- Check if you changed default ports in NVR — update policy to match
- Confirm policy destination IP matches NVR’s actual IP
- Check NVR IP in UniFi Client Devices
- Verify your device’s VLAN is listed in Source Networks
- Try accessing from a different device on an approved VLAN
NVR Shows as Offline in App
Possible Causes:
- NVR doesn’t have internet access
- DNS resolution issues
- Reolink cloud service disruption
- Port forwarding issues (if manually configured)
Solutions:
- Verify “Allow Internet Access” is checked on Camera VLAN
- Check NVR can ping internet addresses (8.8.8.8)
- Test web interface access locally (should work even if cloud is down)
- Check Reolink service status on their website
- Restart NVR and test again
Cameras Not Appearing in NVR
For PoE Cameras:
- Verify camera is connected to NVR’s PoE ports
- Check cable connections
- Try different PoE port on NVR
- Power cycle NVR
- Check camera power indicator lights
For WiFi Cameras:
- Verify camera is on Camera VLAN (check UniFi Client Devices)
- Confirm camera has static IP assigned
- Check WiFi signal strength
- Camera and NVR must be on same VLAN for discovery
- Try re-adding camera through NVR interface
Part 8: Security Best Practices
NVR-Specific Security
- Change Default Password Immediately:
- Access NVR web interface
- Change the admin password from default
- Use strong, unique password (20+ characters)
- Store in password manager
- Never use default credentials like “admin/admin”
2. Keep Firmware Updated:
- Enable Auto Update in NVR settings
- Updates often include security patches
- NVR will automatically check for and install updates
- Test functionality after automatic updates if needed
3. Disable Unnecessary Features:
- Turn off UPnP if not needed for remote access
- Disable P2P if using VPN for remote access
- Turn off cloud features if not used
