---
id: collect-261001-general-networking/general-networking/freedomofpress-securedrop-docs-blob-head-docs-admin-installation-firewall-opnsen-c8ceb602-3
title: "freedomofpress-securedrop-docs-blob-head-docs-admin-installation-firewall-opnsen-c8ceb602"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-general-networking/freedomofpress-securedrop-docs-blob-head-docs-admin-installation-firewall-opnsen-c8ceb602.md
source_anchor: ""
source_lines: [160, 250]
sha256: 509b7db5ae58913538e033c05b49b7a74369803a19353da0ed73041d391b47a5
---

# freedomofpress-securedrop-docs-blob-head-docs-admin-installation-firewall-opnsen-c8ceb602

In the **Generic configuration** section, select `Static IPv4` in the **IPv4 Configuration Type** dropdown, and `None` in the **IPV6 Configuration Type** dropdown.

Scroll down. In the **Static IPv4 Configuration** section, enter the *Application Gateway* IP address and routing prefix (`10.20.2.1` and `24` if you are using the recommended values).

Click **Save**, then click **Apply changes** when prompted.

Finally, navigate to **Interfaces ▸ [OPT2]**. In the **Basic configuration** section, check the checkboxes labeled **Enable interface** and **Prevent interface removal**.

In the **Generic configuration** section, select `Static IPv4` in the **IPv4 Configuration Type** dropdown, and `None` in the **IPV6 Configuration Type** dropdown.

Scroll down. In the **Static IPv4 Configuration** section, enter the *Monitor Gateway* IP address and routing prefix (`10.20.3.1` and `24` if you are using the recommended values).

Click **Save**, then click **Apply changes** when prompted.

In order to simplify firewall rule setup, the next step is to configure aliases for hosts and ports referred to in the rules.

To start, first navigate to **Firewall ▸ Aliases**. You should see some system-defined aliases as shown below:

Click the **+** button to add new aliases. You should add the aliases defined in the table below (assuming recommended values for IP addresses):

| Name | Type | Content | 
|---|---|---|
| admin_workstation | Host(s) | `10.20.1.2` | 
| app_server | Host(s) | `10.20.2.2` | 
| external_dns_servers | Host(s) | `8.8.8.8` ,`8.8.4.4` | 
| monitor_server | Host(s) | `10.20.3.2` | 
| local_servers | Host(s) | `app_server` ,`monitor_server` | 
| OSSEC | Port(s) | `1514` | 
| ossec_agent_auth | Port(s) | `1515` | 
| antilockout_ports | Port(s) | `80` ,`443` | 

When complete, the **Aliases** page should look like this:

Scroll down and click **Apply** to save and apply your new aliases.

Next, configure firewall rules for each interface.

First, navigate to **Firewall ▸ Rules ▸ LAN**.  The LAN interface should have one automatically-generated anti-lockout rule in place, in addition to two default-allow rules. The default-allow rules should be removed once the SecureDrop-specific rules below have been added. The anti-lockout feature should be disabled as a last step.

The rules needed are described in this table:

| Action | TCP/IP Version | Protocol | Src | Src port | Dest | Dest port | Description | 
|---|---|---|---|---|---|---|---|
| Pass | IPv4 | TCP | admin_workstation |  | local_servers | 22 (SSH) | SSH access for initial install | 
| Pass | IPv4 | TCP | admin_workstation |  |  |  | Tor from Tails | 

Add or remove rules until they match the following screenshot including ordering. Click the **+** button to add a rule.

Once the rules match, click **Apply Changes.**

Finally, remove the default anti-lockout rule. First, navigate to **Firewall ▸ Settings ▸ Advanced**. Scroll down to the **Miscellaneous** section and check the **Disable anti-lockout** checkbox. Then, click **Save**.

Next, navigate to **Firewall ▸ Rules ▸ OPT1**. There should be no rules defined on this interface. Add the rules below:

| Action | TCP/IP Version | Protocol | Src | Src port | Dest | Dest port | Description | 
|---|---|---|---|---|---|---|---|
| Pass | IPv4 | UDP | app_server |  | monitor_server | OSSEC | OSSEC Agent | 
| Pass | IPv4 | TCP | app_server |  | monitor_server | ossec_agent_auth | OSSEC initial auth | 
| **Block** | IPv4 | any | OPT1 net |  | LAN net |  | Block between OPT1 and LAN by default | 
| **Block** | IPv4 | any | OPT1 net |  | OPT2 net |  | Block between OPT1 and OPT2 by default | 
| Pass | IPv4 | TCP | app_server |  |  |  | Tor from App Server | 
| Pass | IPv4 | TCP/UDP | app_server |  | external_dns_servers | 53 (DNS) | Allow DNS | 
| Pass | IPv4 | UDP | app_server |  |  | 123 (NTP) | Allow NTP | 

Once they match the screenshot below, click **Apply Changes**.

Next, navigate to **Firewall ▸ Rules ▸ OPT2**. Similarly to OPT1, there should be no rules defined on this interface. Add the rules below until the rules in the Web GUI match those in the screenshot:

| Action | TCP/IP Version | Protocol | Src | Src port | Dest | Dest port | Description | 
|---|---|---|---|---|---|---|---|
| **Block** | IPv4 | any | OPT2 net |  | LAN net |  | Block between OPT2 and LAN by default | 
| **Block** | IPv4 | any | OPT2 net |  | OPT1 net |  | Block between OPT2 and OPT1 by default | 
| Pass | IPv4 | TCP | monitor_server |  |  |  | Tor, SMTP from Monitor Server | 
| Pass | IPv4 | TCP/UDP | monitor_server |  | external_dns_servers | 53 (DNS) | Allow DNS | 
| Pass | IPv4 | UDP | monitor_server |  |  | 123 (NTP) | Allow NTP | 

Finally, click **Apply Changes**.

The *Network Firewall* configuration is now complete, allowing you to move to the next step: :doc:`setting up the servers. <prepare_servers>`

Here are some general tips for setting up OPNSense firewall rules:

1. Create aliases for the repeated values (IPs and ports).
2. OPNSense is a stateful firewall, which means that you don't need corresponding rules to allow incoming traffic in response to outgoing traffic (like you would in, e.g. iptables with `--state ESTABLISHED,RELATED` ).
3. You should create the rules *on the interface where the traffic originates* .
4. Make sure you delete the default "allow all" rule on the LAN interface.
5. If you are troubleshooting connectivity, the firewall logs can be very helpful. You can find them in the Web GUI in **Firewall ▸ Log Files**

Periodically, the OPNSense project maintainers release an update to the OPNSense software running on your firewall. You can check for updates using the link on the OPNSense dashboard.

If you see that an update is available, we recommend installing it. Most of these updates are for minor bugfixes, but occasionally they can contain important security fixes. You should keep apprised of updates yourself by checking the OPNSense Blog or subscribing to the OPNSense Blog RSS feed.
