---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/infrastructure-ubiquiti-udm-dream-machine-wireguard-vpn-30991b8d
title: "infrastructure-ubiquiti-udm-dream-machine-wireguard-vpn-30991b8d"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/infrastructure-ubiquiti-udm-dream-machine-wireguard-vpn-30991b8d.md
source_anchor: ""
source_lines: [1, 33]
sha256: b39c8ce525fdc867c128ff79dc49dfbfbb8befa1f6c0b48c950f75d8c725a283
---

# infrastructure-ubiquiti-udm-dream-machine-wireguard-vpn-30991b8d

WireGuard® is an extremely simple yet fast and modern VPN that utilizes state-of-the-art cryptography. It aims to be faster, simpler, leaner, and more useful than IPsec, while avoiding the massive headache. It intends to be considerably more performant than OpenVPN. WireGuard is designed as a general purpose VPN for running on embedded interfaces and super computers alike, fit for many different circumstances. Initially released for the Linux kernel, it is now cross-platform (Windows, macOS, BSD, iOS, Android) and widely deployable. It is currently under heavy development, but already it might be regarded as the most secure, easiest to use, and simplest VPN solution in the industry.
The setup script will load the wireguard module, and setup the symbolic links for the wireguard tools (wg-quick and wg). You can run dmesg to verify the kernel module was loaded. You should see something like the following:
The wireguard module and tools included in this package have been tested on the following Ubiquiti devices:
Unifi Dream Machine (UDM) and UDM-Pro 0.5.x, 1.9.x, 1.10.x, 1.11.x.
UDM-SE and Unifi Dream Router (UDR) 2.2.x
UniFi Next-Gen Gateway (UXG-Pro) 1.11.x
Note that for the UDM, UDM Pro, and UXG-Pro, Ubiquiti includes the wireguard module in the official kernel since firmware 1.11.0-14, but doesn't include the WireGuard tools. The setup script in this package will try to load the built-in wireguard module if it exists first.
Read the documentation on WireGuard.com for general WireGuard concepts. Here is a simple example of a wireguard server configuration for UnifiOS.
Create the server and client public/private key pairs by running the following. This will create the files privatekey_server, publickey_server and privatekey_client1, publickey_client1. These contain the public and private keys. Store these files somewhere safe.
On your UDM/UDR, create a wireguard config under /etc/wireguard named wg0.conf. Here is an example server config. Remember to use the correct server private key and the client public key.
Adjust AllowedIPs to set what your client should route through the tunnel. Set to 0.0.0.0/0,::/0 to route all the client's Internet through the tunnel. See the WireGuard documentation for more information.
Note each different client requires their own private/public key pair, and the public key must be added to the server's WireGuard config as a separate Peer.
To bring the tunnel up, run wg-quick up <config>. Verify the tunnel received a handshake by running wg.
wg-quickup/etc/wireguard/wg0.conf
To bring down the tunnel, run wg-quick down <config>.
wg-quickdown/etc/wireguard/wg0.conf
In your UniFi Network settings, add a WAN_LOCAL (or Internet Local) firewall rule to ACCEPT traffic destined to UDP port 51820 (or your ListenPort if different). Opening this port in the firewall is needed so remote clients can access the WireGuard server.
The AllowedIPs parameter in the wireguard config allows you to specify which destination subnets to route through the tunnel.
If you want to route router-connected clients through the wireguard tunnel based on source subnet or source VLAN, you need to set up policy-based routing. Currently, there is no GUI support for policy-based routing in UnifiOS, but it can be set up in SSH by using ip route to create a custom routing table, and ip rule to select which clients to route through the custom table.
For a script that makes it easy to set-up policy-based routing rules on UnifiOS, see the split-vpn project.
The setup script must be run every time the system is rebooted to link the wireguard tools and load the module. This can be accomplished with a boot script.
For the UDM or UDM Pro, install UDM Utilities on-boot-script by following the instructions here, then create a boot script under /mnt/data/on_boot.d/99-setup-wireguard.sh and fill it with the following contents. Remember to run chmod +x /mnt/data/on_boot.d/99-setup-wireguard.sh afterwards.
Click here to see the boot script.
#!/bin/sh
/mnt/data/wireguard/setup_wireguard.sh
For the UDM-SE or UDR, create a systemd boot service to run the setup script at boot. Create a service file under /etc/systemd/system/setup-wireguard.service and fill it with the following contents. After creating the service, run systemctl daemon-reload && systemctl enable setup-wireguard to enable the service on boot.
Click here to see the boot service.
[Unit]Description=Run wireguard setup scriptWants=network.targetAfter=network.target[Service]Type=oneshotExecStart=sh -c 'WGDIR="$(find /mnt/data/wireguard /data/wireguard -maxdepth 1 -type d -name "wireguard" 2>/dev/null | head -n1)"; "$WGDIR/setup_wireguard.sh"'[Install]WantedBy=multi-user.target
Note this only adds the setup script to start at boot. If you also want to bring your wireguard interface up at boot, you will need to add another boot script with your wg-quick up command.
Setup script returns error "Unsupported Kernel version XXX" * The wireguard package does not contain a wireguard module built for your firmware or kernel version, nor is there a built-in module in your kernel. Please open an issue and report your version so we can try to update the module. wg-quick up returns error "unable to initialize table 'raw'" * Your kernel does not have the iptables raw module. The raw module is only required if you use `0.0.0.0/0` or `::/0` in your wireguard config's AllowedIPs. A workaround is to instead set AllowedIPs to `0.0.0.0/1,128.0.0.0/1` for IPv4 or `::/1,8000::/1` for IPv6. These subnets cover the same range but do not invoke wg-quick's use of the iptables raw module.
"WireGuard" and the "WireGuard" logo are registered trademarks of Jason A. Donenfeld.
The built-in gateway DNS does not reply to requests from the WireGuard tunnel¶
The built-in dnsmasq on UnifiOS is configured to only listen for requests from specific interfaces. The wireguard interface name (e.g.: wg0) needs to be added to the dnsmasq config so it can respond to requests from the tunnel. You can run the following to add wg0 to the dnsmasq interface list:
