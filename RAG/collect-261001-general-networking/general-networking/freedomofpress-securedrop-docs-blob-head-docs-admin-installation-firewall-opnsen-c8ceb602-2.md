---
id: collect-261001-general-networking/general-networking/freedomofpress-securedrop-docs-blob-head-docs-admin-installation-firewall-opnsen-c8ceb602-2
title: "freedomofpress-securedrop-docs-blob-head-docs-admin-installation-firewall-opnsen-c8ceb602"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/freedomofpress-securedrop-docs-blob-head-docs-admin-installation-firewall-opnsen-c8ceb602.md
source_anchor: ""
source_lines: [65, 159]
sha256: df390777b2710b1aa8fe790a86417ee14de502917cf337945d917eb5a6504607
---

# freedomofpress-securedrop-docs-blob-head-docs-admin-installation-firewall-opnsen-c8ceb602

You should update OPNSense to the latest version available before proceeding with the rest of the configuration. Navigate to **Lobby ▸ Dashboard** and click **Click to check for updates** to start the process, and follow any on-screen instructions to complete the update. Note that a reboot may be required, and you may also need to apply several updates in a row to get to the latest version.

OPNSense supports two-factor authentication (2FA) via mobile apps such as Google Authenticator or FreeOTP. To set it up, first make sure you have a mobile device available with your choice of 2FA app.

Next, in the OPNSense Web GUI, navigate to **System ▸ Access ▸ Servers** and click **+** to add a new server.

Note

The time on your firewall must be set correctly for 2FA to work properly. This should happen automatically once the WAN connection is established.

On the next page, enter `TOTP Local` in the **Descriptive name** field and choose `Local + Timebased One Time Password` from the **Type** dropdown. Leave the other fields at their default values and click **Save**

Next, navigate to **System ▸ Access ▸ Users** and click the edit button for the `root` user. Scroll down the page to the **OTP seed** section and check the **Generate new secret (160bit)** checkbox. Finally, click **Save**.

Once the page has reloaded, scroll down to the **OTP QR code** section and click **Click to unhide**, then scan the generated QR code with your mobile auth application of choice.

If you wish, you may also save the OTP seed value displayed above the QR code in your Tails KeePassXC database - this isn't required, but will allow you to set up TOTP on another mobile device if you need to in the future.

To verify that your new password and OTP secret are working, navigate to **System ▸ Access ▸ Tester**. Select `TOTP Local` from the **Authentication Server** dropdown, enter the `root` username in the **Username** field, and enter your OTP token and password concatenated like `123456PASSWORD` in the **Password** field. Then click **Test**.

If the test fails, make sure you have used the correct OTP code and password, and edit the `root` user record as necessary.

Note

You must enter the OTP token and passphrase concatenated as a single string like `123456PASSWORD` in the **Password** field.

Warning

Do not skip this test, or proceed further until it passes, as you will be locked out of the firewall Web GUI and console if the account is not set up correctly!

Finally,  navigate to **System ▸ Settings ▸ Administration** and scroll down to the **Authentication** section at the bottom of the page. In the **Server** dropdown, select `TOTP Local` and deselect `Local Database.`. Click **Save**.

OPNSense runs a DHCP server on the LAN interface by default. At this stage in the documentation, the Admin Workstation likely has an IP address assigned via that DHCP server.

In order to tighten the firewall rules as much as possible, we recommend disabling the DHCP server and assigning a static IP address to the Admin Workstation instead.

To disable DHCP, navigate to **Services ▸ DHCPv4 ▸ [LAN]** in the Web GUI. Uncheck the **Enable DHCP server on the LAN interface** checkbox, scroll down, and click **Save**.

Now you will need to assign a static IP to the Admin Workstation.

You can easily check your current IP address by *clicking* the top right of the menu bar, clicking on the **Wired Connection** and then clicking **Wired Settings**.

From here you can click on the cog beside the wired network connection:

This will take you to the network settings. Change to the **IPv4** tab. Ensure that **IPv4 Method** is set to **Manual**, and that the **Automatic** switch for **DNS** is in the "off" position, as highlighted in the screenshot below:

Note

The Unsafe Browser will not launch when using a manual network configuration if it does not have DNS servers configured. This is technically unnecessary for our use case because we are only using it to access IP addresses on the LAN, and do not need to resolve anything with DNS. Nonetheless, you should configure some DNS servers here so you can continue to use the Unsafe Browser to access the WebGUI in future sessions.

We recommend keeping it simple and using the same DNS servers that you used for the network firewall in the setup wizard.

Fill in the static networking information for the Admin Workstation:

- Address: `10.20.1.2`
- Netmask: `255.255.255.0`
- Gateway : `10.20.1.1`

Click **Apply**. If the network does not come up within 15 seconds or so, try disconnecting and reconnecting your network cable to trigger the change. You will need you have succeeded in connecting with your new static IP when you are able to connect using the Tor Connection assistant, and you see the message "Connected to Tor successfully".

After saving the new network configuration, you may still encounter the "No DNS servers configured" error when trying to launch the Unsafe Browser. If you encounter this issue, you can resolve it by disconnecting from the network and then reconnecting, which causes the network configuration to be reloaded.

To do this, click the network icon in the system toolbar, and click **Disconnect** under the name of the currently active network connection, which is displayed in bold. After it disconnects, click the network icon again and click the name of the connection to reconnect. You should see a popup notification that says "Connection Established", and the Tor Connection assistant should show the message "Connected to Tor successfully".

For the next step, SecureDrop Configuration, you will manually configure the firewall for SecureDrop, using screenshots as a reference.

SecureDrop uses the firewall to achieve two primary goals:

1. Isolating SecureDrop from the existing network, which may be compromised (especially if it is a venerable network in a large organization like a newsroom).
2. Isolating the Application Server and the Monitor Server from each other as much as possible, to reduce attack surface.

In order to use the firewall to isolate the Application Server and the Monitor Server from each other, we need to connect them to separate interfaces, and then set up firewall rules that allow them to communicate.

The OPT1 and OPT2 interfaces will be used for the Application Server and Monitor Server respectively. To enable them, first connect the Application Server to the physical OPT1 port and the Monitor Server to the OPT2 port.

Next, navigate to **Interfaces ▸ Assignments**. LAN and WAN will already be enabled. Click the **+** button in the **New Interface** section to enable the OPT1 interface on the next available NIC (`igb2` in the screenshot below). Once OPT1 has been added, click **+** again to add OPT2 (on `igb3` in the screenshot below)

Finally, click **Save**.

OPT1 and OPT2 need to be configured to use the subnets defined for the Application and Monitor Servers, and some additional configuration is required for the LAN and WAN interfaces, that is not covered by the Setup Wizard.

First, navigate to **Interfaces ▸ [WAN]**. In the **Basic configuration** section, check the checkbox labeled **Prevent interface removal**.

In the **Generic configuration** section, make sure that the **Block private networks** and **Block bogon networks** checkboxes are checked.

Scroll down and click  **Save**, then click **Apply changes** when prompted.

Next, navigate to **Interfaces ▸ [LAN]**. In the **Basic configuration** section, check the checkbox labeled **Prevent interface removal**.

In the **Generic configuration** section, select `Static IPv4` in the **IPv4 Configuration Type** dropdown, and `None` in the **IPV6 Configuration Type** dropdown.

Scroll down and click **Save**, then click **Apply changes** when prompted.

Next, navigate to **Interfaces ▸ [OPT1]**. In the **Basic configuration** section, check the checkboxes labeled **Enable interface** and **Prevent interface removal**.

