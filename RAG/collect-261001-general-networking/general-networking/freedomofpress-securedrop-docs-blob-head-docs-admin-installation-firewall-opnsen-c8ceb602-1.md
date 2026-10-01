---
id: collect-261001-general-networking/general-networking/freedomofpress-securedrop-docs-blob-head-docs-admin-installation-firewall-opnsen-c8ceb602-1
title: "freedomofpress-securedrop-docs-blob-head-docs-admin-installation-firewall-opnsen-c8ceb602"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/freedomofpress-securedrop-docs-blob-head-docs-admin-installation-firewall-opnsen-c8ceb602.md
source_anchor: ""
source_lines: [1, 64]
sha256: 05110eba8e0a2b0d2000189247bd5a808200052488a657805c29361a24c5ae22
---

# freedomofpress-securedrop-docs-blob-head-docs-admin-installation-firewall-opnsen-c8ceb602

First, consider how the firewall will be connected to the Internet. You will need to provision several unique subnets, which should not conflict with the network configuration on the WAN interface. If you are unsure, consult your local system administrator.

Many firewalls, including the recommended OPNSense device, automatically set up the LAN interface on `192.168.1.1/24`. This particular private network is also a very common choice for home and office routers. If you are connecting the firewall to a router with the same subnet (common in a small office, home, or testing environment), you will probably be unable to connect to the network at first. However, you will be able to connect from the LAN to the firewall's Web GUI, and from there you will be able to configure the network so it is working correctly.

The recommended TekLager APU4D4 has 4 NICs: WAN, LAN, OPT1, and OPT2. This allows for a dedicated port on the network firewall for each component of SecureDrop (Application Server, Monitor Server, and Admin Workstation).

Depending on your network configuration, you should define the following values before continuing.

- Admin Subnet: `10.20.1.0/24`
- Admin Gateway: `10.20.1.1`
- Admin Workstation: `10.20.1.2`

- Application Subnet: `10.20.2.0/24`
- Application Gateway: `10.20.2.1`
- Application Server (OPT1): `10.20.2.2`

- Monitor Subnet: `10.20.3.0/24`
- Monitor Gateway: `10.20.3.1`
- Monitor Server (OPT2) : `10.20.3.2`

Unpack the firewall, connect the power, and power on the device.

We will use the OPNSense Web GUI to do the initial configuration of the network firewall.

1. If you have not already done so, boot the Admin Workstation.
2. Connect the Admin Workstation to the LAN interface. You should see a popup notification in Tails that says "Connection Established". If you click on the network icon in the upper right of the Tails Desktop, you should see that the "Wired Connection" is active: Warning Make sure your *only* active connection is the one you just established with the network firewall. If you are connected to another network at the same time (e.g. a wireless network), you may encounter problems trying to connect the firewall's Web GUI.
3. Launch the Unsafe Browser from the menu bar: **Apps ▸ Internet ▸ Unsafe Browser** .Note The Unsafe Browser is, as the name suggests, **unsafe** (its traffic is not routed through Tor). However, it is the only option because Tails intentionally disables LAN access in the**Tor Browser** .
4. You will see a pop-up notification that says "Starting the Unsafe Browser..."
5. After a few seconds, the Unsafe Browser should launch. The window has a bright red border to remind you to be careful when using it. You should close it once you're done configuring the firewall and use Tor Browser for any other web browsing you might do on the Admin Workstation.
6. Navigate to the OPNSense Web GUI in the Unsafe Browser: `https://192.168.1.1`Note If you have trouble connecting, go to your network settings and make sure that you have an IPv4 address in the `192.168.1.1/24` range. You may need to turn on DHCP, else you can manually configure a static IPv4 address of`192.168.1.x` with a subnet mask of`255.255.255.0` . However, make sure not to configure your Tails device to have the same IP as the firewall (`192.168.1.1` ).
7. The firewall uses a self-signed certificate, so you will see a "This Connection Is Untrusted" warning when you connect. This is expected. You can safely continue by clicking **Advanced** and**Accept the Risk and Continue** .
8. You should see the login page for the OPNSense GUI. Log in with the default username and passphrase ( `root` /`opnsense` ).

If this is your first time logging in to the firewall, the setup wizard will be displayed. You should not step through it at this point, however, as there are other tasks to complete. To exit, click the OPNSense logo in the top left corner of the screen.

Navigate to **System ▸ Access ▸ Users** and click the edit button for the `root` user. On the subsequent page, set a strong admin password. We recommend generating a strong passphrase with KeePassXC and saving it in the Tails Persistent folder using the provided KeePassXC database template. Two-factor authentication will be enabled in a later step.

Before you can set up the hardware firewall, you will need to set the **Alternate Hostnames** setting.

First, navigate to **System ▸ Settings ▸ Administration**.  In the **Web GUI** section, update the **Alternate Hostnames** field with the values `192.168.1.1` and the IP address of the *Admin Gateway* (`10.20.1.1` if you are using the recommended default values), separated by a space.

Finally, scroll to the bottom of the page and click **Save**.

To start the OPNSense Setup Wizard, navigate to **System ▸ Wizard** and click **Next**.

1. **General Information** : Leave your hostname as the default,`OPNsense` . There is no relevant domain for SecureDrop, so we recommend setting this to`securedrop.local` or something similar. Use your preferred DNS servers. If you don't know what DNS servers to use, we recommend using Google's DNS servers:`8.8.8.8` and`8.8.4.4` . Uncheck the**Override DNS** checkbox.In the **Unbound DNS** section, uncheck**Enable Resolver** .Click **Next** .
2. **Time Server Information** : Leave the default settings unchanged and  click**Next** .
3. **Configure WAN Interface** : Enter the appropriate configuration for your network. Consult your local sysadmin if you are unsure what to enter here. For many environments, the default of DHCP will work and the rest of the fields can be left at their default values.Click **Next** to proceed.
4. **Configure LAN Interface** : Use the IP address of the*Admin Gateway* (`10.20.1.1` ) and the subnet mask (`/24` ) of the*Admin Subnet* . Click**Next** .
5. **Set Root Password** : If the password was already reset during the 2FA setup, you don't need to set it again. If it was not, then set a strong password now and store it in the Admin Workstation's KeePassXC database. Click**Next** to continue.
6. **Reload Configuration** : Click**Reload** to apply the changes you made in the Setup Wizard.

At this point, since the LAN subnet settings were changed from their defaults, you will no longer be able to connect after reloading the firewall and the reload will time out. This is not an error - the firewall has reloaded and is working correctly.

To connect to the new LAN interface, unplug and reconnect your network cable to get a new network address assigned via DHCP. Note that if you used a subnet with fewer addresses than `/24`, the default DHCP configuration in OPNSense may not work. In this case, you should assign the Admin Workstation a static IP address that is known to be in the subnet to continue.

The Web GUI will now be available on the *Admin Gateway* IP address. Navigate to `https://<Admin Gateway IP>` in the Unsafe Browser and log in to the `root` account using an OTP token and the passphrase you just set.

Once you've logged in to the Web GUI, you are ready to continue configuring the firewall.

Now that the initial configuration is completed, you can connect the WAN port without potentially conflicting with the default LAN settings (as explained earlier). Connect the WAN port to the external network. You can watch the WAN entry in the Interfaces table on the OPNSense Dashboard homepage to see as it changes from down (red arrow pointing down) to up (green arrow pointing up). This usually takes several seconds. The WAN's IP address will be shown once it comes up.

Finally, test connectivity to make sure you are able to connect to the Internet through the WAN. The easiest way to do this is to open another tab in the Unsafe Browser and visit a host that you expect to be up (e.g. `google.com`).

