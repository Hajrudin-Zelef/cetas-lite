---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/aioue-ansible-unifi-inventory-8ce0b2d3-1
title: "Example: prod.unifi.yml"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/aioue-ansible-unifi-inventory-8ce0b2d3.md
source_anchor: ""
source_lines: [1, 295]
sha256: e128372657c6070766ef55768cf48560343f7c0e3cef4ff2c3416bdd0e76e898
---

# Example: prod.unifi.yml

Dynamic inventory plugin for Ansible that discovers hosts from a UniFi OS controller (UDM, UCG, etc.). Built on top of aiounifi (v91+ required; tested with v92).
$ ansible-inventory -i inventory/unifi.yaml all --graph
@all:
  |--@ungrouped:
  |--@unifi_clients:
  |  |--Kitchen_Echo
  |  |--nas-server
  |  |--Study_Proxmox
  |  |--phone
  |  |--homeassistant
  |--@unifi_wireless_clients:
  |  |--Kitchen_Echo
  |  |--phone
  |--@unifi_wired_clients:
  |  |--nas-server
  |  |--Study_Proxmox
  |  |--homeassistant
  |--@network_default:
  |  |--nas-server
  |  |--Study_Proxmox
  |  |--homeassistant
  |  |--phone
  |--@network_iot:
  |  |--Kitchen_Echo
  |--@vlan_30:
  |  |--Kitchen_Echo
  |--@vlan_iot:
  |  |--Kitchen_Echo
  |--@ssid_home_iot:
  |  |--Kitchen_Echo
  |--@ssid_home_wifi:
  |  |--phone
  |--@unifi_devices:
  |  |--U6_Pro
  |  |--USW_Flex
  |  |--USW_Ultra
  |  |--Dream_Machine
  |--@unifi_uap:
  |  |--U6_Pro
  |--@unifi_usw:
  |  |--USW_Flex
  |  |--USW_Ultra
  |--@device_state_connected:
  |  |--U6_Pro
  |  |--USW_Flex
  |  |--USW_Ultra
  |  |--Dream_Machine
  |--@unifi_poe_powered:
  |  |--USW_Flex
  |--@unifi_udm:
  |  |--Dream_Machine
Example host variables (include_devices: true; all fields from a live run with aiounifi v92, identifying values sanitized):
$ ansible-inventory -i inventory/unifi.yaml --host Kitchen_Echo
{
  "ansible_host": "192.168.30.13",
  "association_time": 1783656843,
  "first_seen": 1735624765,
  "fixed_ip": "192.168.30.13",
  "ip": "192.168.30.13",
  "ipv4": "192.168.30.13",
  "ipv6": "2001:db8:1::772",
  "ipv6_addresses": [
    "2001:db8:1::772",
    "fe80::4a78:5eff:fefa:7ce1"
  ],
  "is_wired": false,
  "last_seen_iso": "2026-07-22T18:33:32Z",
  "last_seen_unix": 1784745212,
  "latest_association_time": 1784294898,
  "mac": "48:78:5e:fa:7c:e1",
  "network": "IoT",
  "network_id": "670ef99bba911339bf2b894c",
  "oui": "Amazon Technologies Inc.",
  "powersave_enabled": false,
  "site": "default",
  "ssid": "home.iot",
  "unifi_name": "Kitchen Echo",
  "vlan": 30,
  "vlan_name": "IoT"
}
$ ansible-inventory -i inventory/unifi.yaml --host nas-server
{
  "ansible_host": "192.168.1.148",
  "association_time": 1781365738,
  "first_seen": 1773981600,
  "ip": "192.168.1.148",
  "ipv4": "192.168.1.148",
  "ipv6": "fe80::be24:11ff:feaf:77dd",
  "is_wired": true,
  "last_seen_iso": "2026-07-22T18:33:42Z",
  "last_seen_unix": 1784745222,
  "latest_association_time": 1781925104,
  "mac": "bc:24:11:af:77:dd",
  "network": "Default",
  "network_id": "670ed85d2ce59e0ea329eff1",
  "oui": "Example Vendor Inc.",
  "site": "default",
  "switch_depth": 1,
  "unifi_hostname": "nas-server",
  "unifi_name": "nas-server",
  "wired_rate_mbps": 1000
}
$ ansible-inventory -i inventory/unifi.yaml --host U6_Pro
{
  "ansible_host": "192.168.1.252",
  "client_count": 7,
  "cpu_percent": "7.9",
  "device_id": "671113bfba911339bf2be6c8",
  "disabled": false,
  "firmware_version": "6.8.2.15592",
  "has_fan": false,
  "has_temperature": false,
  "ip": "192.168.1.252",
  "last_seen": 1784745212,
  "led_override": "off",
  "led_override_color": "#0000ff",
  "mac": "ac:8b:a9:43:b5:cd",
  "mem_percent": "65.7",
  "model": "UAP6MP",
  "overheating": false,
  "site": "default",
  "state": "CONNECTED",
  "supports_led_ring": false,
  "system_uptime": "4768028",
  "type": "uap",
  "unifi_name": "U6 Pro",
  "upgradable": false,
  "uplink": {
    "full_duplex": true,
    "max_speed": 1000,
    "name": "eth0",
    "port_idx": 1,
    "speed": 1000,
    "type": "wire",
    "up": true,
    "uplink_device_name": "USW Ultra",
    "uplink_mac": "28:70:4e:6d:f9:32",
    "uplink_remote_port": 1,
    "uplink_source": "lldp_uplink"
  },
  "uptime": 4768028
}
$ ansible-inventory -i inventory/unifi.yaml --host USW_Flex
{
  "ansible_host": "192.168.1.231",
  "client_count": 5,
  "cpu_percent": "11.0",
  "device_id": "67d1d99d2b37f907a11a58ed",
  "disabled": false,
  "firmware_version": "2.1.8.971",
  "has_fan": false,
  "has_temperature": false,
  "ip": "192.168.1.231",
  "last_seen": 1784745222,
  "led_override": "on",
  "led_override_color": "#0000ff",
  "mac": "94:2a:6f:fe:0e:e5",
  "mem_percent": "82.8",
  "model": "USWED37",
  "overheating": false,
  "poe_ports": [
    {
      "is_uplink": false,
      "name": "Port 1",
      "poe_enable": false,
      "poe_good": false,
      "poe_mode": "auto",
      "poe_power": "0.00",
      "poe_voltage": "0.00",
      "port_idx": 1,
      "up": false
    },
    {
      "is_uplink": false,
      "name": "Port 2",
      "poe_enable": false,
      "poe_good": false,
      "poe_mode": "auto",
      "poe_power": "0.00",
      "poe_voltage": "0.00",
      "port_idx": 2,
      "up": false
    },
    {
      "is_uplink": false,
      "name": "Port 3",
      "poe_enable": false,
      "poe_good": false,
      "poe_mode": "auto",
      "poe_power": "0.00",
      "poe_voltage": "0.00",
      "port_idx": 3,
      "up": false
    },
    {
      "is_uplink": false,
      "name": "Port 4",
      "poe_enable": false,
      "poe_good": false,
      "poe_mode": "auto",
      "poe_power": "0.00",
      "poe_voltage": "0.00",
      "port_idx": 4,
      "up": true
    },
    {
      "is_uplink": false,
      "name": "Port 5",
      "poe_enable": false,
      "poe_good": false,
      "poe_mode": "auto",
      "poe_power": "0.00",
      "poe_voltage": "0.00",
      "port_idx": 5,
      "up": true
    },
    {
      "is_uplink": false,
      "name": "Port 6",
      "poe_enable": true,
      "poe_good": true,
      "poe_mode": "auto",
      "poe_power": "5.91",
      "poe_voltage": "47.24",
      "port_idx": 6,
      "up": true
    },
    {
      "is_uplink": false,
      "name": "Port 7",
      "poe_enable": true,
      "poe_good": true,
      "poe_mode": "auto",
      "poe_power": "13.60",
      "poe_voltage": "47.05",
      "port_idx": 7,
      "up": true
    },
    {
      "is_uplink": false,
      "name": "Port 8",
      "poe_enable": false,
      "poe_good": false,
      "poe_mode": "auto",
      "poe_power": "0.00",
      "poe_voltage": "0.00",
      "port_idx": 8,
      "up": false
    }
  ],
  "site": "default",
  "state": "CONNECTED",
  "supports_led_ring": false,
  "type": "usw",
  "unifi_name": "USW Flex 2.5G 8 PoE",
  "upgradable": false,
  "uplink": {
    "full_duplex": true,
    "max_speed": 10000,
    "media": "10GE",
    "name": "eth0",
    "port_idx": 9,
    "speed": 2500,
    "type": "wire",
    "up": true,
    "uplink_device_name": "Dream Machine",
    "uplink_mac": "28:70:4e:6e:44:a7",
    "uplink_remote_port": 4,
    "uplink_source": "lldp_uplink"
  },
  "uptime": 4768031
}
- Dynamic inventory plugin for Ansible that fetches UniFi network clients as inventory hosts.
- Supports UniFi OS controllers (modern UniFi Dream Machine, Cloud Gateway, etc.).
- Discovers clients connected to your network (wired and wireless).
- Optionally includes UniFi devices (access points, switches, gateways).
- Not a UniFi controller configuration tool.
- Not compatible with legacy UniFi controllers (pre-UniFi OS) without modification.
- Python 3.12+ (newer aiounifi releases may require 3.13+; checkpip install output)
- Ansible 2.15+
- UniFi OS controller accessible via network (UDM, UCG, etc.).
- API credentials: API token (preferred), local admin without 2FA, or username/password with totp_secret for 2FA accounts (aiounifi v91+)
- Python dependencies: Install in the same Python environment as Ansible:
pip install -r requirements.txt
Install the aioue.network collection from this GitHub repository:
ansible-galaxy collection install git+https://github.com/aioue/ansible-unifi-inventory.git
You can also include it in a requirements.yml file:
---
collections:
  - name: aioue.network
    source: https://github.com/aioue/ansible-unifi-inventory.git
    type: git
    # If you need a specific version, you can specify a branch or tag:
    # version: v1.1.0
Then install with ansible-galaxy collection install -r requirements.yml.
