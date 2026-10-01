---
id: collect-261001-general-networking/general-networking/manual-settingsmenu-html-35cf8d4f-1
title: "manual-settingsmenu-html-35cf8d4f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/manual-settingsmenu-html-35cf8d4f.md
source_anchor: ""
source_lines: [1, 83]
sha256: e7cee3a500e050bd57f67f79c2571d3f68248a35f0c5a7b4daef2b3d0823da85
---

# manual-settingsmenu-html-35cf8d4f

Settings
Besides the configuration options that every component has, OPNsense also contains a lot of general settings that you can tweak. This page contains an overview of them.
Administration
The settings on this page concerns logging into OPNsense. The “Secure Shell” settings are described under Creating Users & Groups.
Listen interfaces
Warning
Before considering the use of manual selected interfaces, make sure to read this chapter so you are aware of the pitfalls upfront. Misconfigurations likely lead to a non-accessible web interface and/or missing ssh access.
Both the WebUI and the Secure Shell server support the option to only listen on specific interfaces, the use of this option however comes with clear warnings which you do need to be aware of before deciding to use this option.
By default (our recommended settings), these services listen on all addresses (interfaces).
If for whatever reason, you do need to listen only on specific interfaces, the following rules apply:
- The interface must always be available, so do not try to bind to vpn instances of any kind (OpenVPN, Wireguard, …)
- The addressing must be fully static, so no IPv6 tracking configured for example
As the webgui is not able to predict with 100% certainty that these rules do apply, it is possible to select interfaces that don’t support binding for these services.
Note
When facing issues with the webgui (and/or ssh) and the above rules are not met, please do not bother to open a ticket as these are unsupported scenario’s.
Tip
In case (for any service) one would like to prevent binding on all interfaces, it is possible to add a loopback interface (), assign an ip address and bind to that.
When traffic is routed through the firewall, the “loopback ip” (some private address, not in the loopback range)
should be directly accessible from the network behind it. For example use an address like 192.168.130.1/32
to access the web interface while your own network is using 192.168.1.0/24.
Technologies like Network Address Translation (NAT) can also be combined if the other end is not aware of the route to this single address.
Web GUI
| Option | Description | 
|---|---|
| Protocol | It is strongly recommended to leave this on “HTTPS” | 
| SSL Certificate | By default, a self-signed certificate is used. Certificates can be added via . | 
| SSL Ciphers | Can be used to limit SSL cipher selection in case the system defaults are undesired. Note that restrictive use may lead to an inaccessible web GUI. | 
| HTTP Strict Transport Security | Enforces loading the web GUI over HTTPS, even when the connection is hijacked (man-in-the-middle attack), and do not allow the user to trust an invalid certificate for the web GUI. | 
| TCP port | Can be useful if there are other services that are reachable via port 80/443 of the external IP, for example. | 
| HTTP Redirect | If you change the port, a redirect rule from port 80/443 will be created. Check this option to disable the creation this automatic redirect rule. | 
| Login Messages | If checked, disable the successful login messages in the web GUI. | 
| Session Timeout | Time in minutes to expire idle management sessions. | 
| DNS Rebind Check | OPNsense contains protection against DNS rebinding by filtering out DNS replies with local IPs. Check this box to disable this protection if it interferes with web GUI access or name resolution in your environment. | 
| Alternate Hostnames | Alternate, valid hostnames (to avoid false positives in referrer/DNS rebinding protection). | 
| HTTP Compression | Reduces size of transfer, at the cost of slightly higher CPU usage. | 
| Access log | Log all access to the Web GUI for debugging/analysis. | 
| Server Log | Display all web GUI errors in the main system log. | 
| Listen interfaces | Can be used to limit interfaces on which the Web GUI can be accessed. This allows freeing the interface for other services, such as HAProxy. | 
| HTTP_REFERER enforcement check | The origins of requests are checked in order to provide some protection against CSRF. You can turn this off if it interferes with external scripts that interact with the Web GUI. | 
Deployment settings
| Option | Description | 
|---|---|
| Deployment type | Influences error feedback to the user | 
| Strict security | Prevent the webgui from running as root, some legacy components may not be compatible with this feature. Disabling the feature again requires console access. | 
Secure Shell
User accounts can be used for logging in to the web frontend, as well as for logging in to the console (via VGA,
serial or SSH). The latter will only work if the user shell is not set to /sbin/nologin.
In order to access OPNsense via SSH, SSH access will need to be configured via . Under the “Secure Shell” heading, the following options are available:
| Option | Description | 
|---|---|
| Secure Shell Server | Enable a secure shell service | 
| Permit Root Login | Root login is generally discouraged. It is advised to log in via another user and switch to root afterwards. | 
| Permit password login | When disabled, authorized keys need to be configured for each User that has been granted secure shell access. | 
| SSH port | Port to listen on, default is 22 | 
| Listen Interfaces | Only accept connections from the selected interfaces. Leave empty to listen globally. Use with extreme care. | 
| Key exchange algorithms | The key exchange methods that are used to generate per-connection keys | 
| Ciphers | The ciphers to encrypt the connection | 
| MACs | The message authentication codes used to detect traffic modification | 
| Host key algorithms | Specifies the host key algorithms that the server offers | 
| Public key signature algorithms | The signature algorithms that are used for public key authentication | 
Secure Shell - Advanced Settings
To configure options that are not available in the GUI one can add custom configuration files on the firewall itself.
Files can be added in /usr/local/etc/ssh/sshd_config.d/ using the .conf file extension.
When more files are placed inside the directory, they will be included in alphabetical order.
The custom configuration files are included first to provide proper overrides to options normally set by the system.
Warning
It is the responsibility of the administrator to ensure that the configuration is valid. No configuration checks will be performed and any errors may prevent you being able to login to the firewall remotely!
Console
In case of an emergency, it’s always practical to make sure to configure a console to be able to access the firewall when network connectivity is not possible.
Tip
After initial installation, always make sure to test if the console actually works. When concluding the console is not functional when you need it can be very unpractical.
| Option | Description | 
|---|---|
| Use the virtual terminal driver (vt) | When unchecked, OPNsense will use the older sc driver. | 
| Primary Console | The primary console will show boot script output. All consoles display OS boot messages, console messages, and the console menu. | 
| Secondary Console | See above. | 
| Serial Speed | Allows adjusting the baud rate. 115200 is the most common. | 
| Use USB-based serial ports | Listen on /dev/ttyU0 ,/dev/ttyU1 , … instead of/dev/ttyu0 . | 
| Password protect the console menu | Can be unchecked to allow physical console access without password. This can avoid lock-out, but at the cost of attackers being able to do anything if they gain physical access to your system. | 
Authentication
The authentication section of the Administrationm settings offers general security settings for users logging into the firewall.
| Option | Description | 
|---|---|
