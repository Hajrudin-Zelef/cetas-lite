---
id: collect-261001-general-networking/general-networking/genius0x1-how-i-automated-firewall-management-with-python-and-rest-apis-fortigat-a27f4356
title: "Discard ssl warnings"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/genius0x1-how-i-automated-firewall-management-with-python-and-rest-apis-fortigat-a27f4356.md
source_anchor: ""
source_lines: [1, 84]
sha256: 51de57ee4fa4128b7b4295acf8e21c087813c926ec39f03f0193c883e51799e5
---

# Discard ssl warnings

How I Automated Firewall Management with Python and REST APIs (FortiGate Example)
Introduction
Hello, I am Ahmed Abdelslam, currently working as Network Security Engineer at SEE.
Get genius0x1’s stories in your inbox
Join Medium for free to get updates from this writer.
I was assigned the task of migration between a very old firewall called Stonegate and FortiGate. My task is to move 914 address and 32 Address group from the old firewall to the new one. It would be time consuming and prone to human error if done manually. so i decided to use Python and FortiGate REST APIs to automate the process.
Preparing the Data (CSV Files)
First i opened Stonegate and copied all Addresses, IP Addresses and Groups and put them into two CSV Files.
- Addresses.csv → contains all the IP addresses I needed to migrate.
Name IP Address
a    1.1.1.1
b    1.1.1.2
2. Groups.csv → maps each address to the group it belongs to.
Name Group
a     x
b     y
FortiGate REST API
First i Created a suitable API profile in Fortigate From System=>Admin Profiles and enabled Read/Write Permission in Address Only (API_Address).
Then i created REST API Administrator from System=>Administrators=> Create New=> REST API Admin and chose Administrator profile (API_Address)
these step generated a token which we will use later.
Using Python
I created a simple script which will read columns from Addresses.csv and create them direct in FortiGate. Please don’t forget to replace FGT_IP and API_KEY with their Value.
import csv
import requests
import urllib3
# Discard ssl warnings
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
FGT_IP = "https://<FIREWALL_IP>"        #You can use HTTP also
API_KEY = "<API_TOKEN>"                #The token we generated in FortiGate
headers = {
    "Authorization": f"Bearer {API_KEY}",
    "Content-Type": "application/json"
}
with open('addresses.csv', 'r') as file:
    reader = csv.DictReader(file)
    for row in reader:
        name = row['Name'].strip()
        ip = row['IP Address'].strip()
        #Address object
        data = {
            "name": name,
            "subnet": f"{ip} 255.255.255.255"
        }
        url = f"{FGT_IP}/api/v2/cmdb/firewall/address"
        response = requests.post(url, headers=headers, json=data, verify=False)
        if response.status_code == 200:
            print(f"[✓] Added {name} ({ip}) successfully")
        else:
            print(f"[✗] Failed to add {name}: {response.status_code} - {response.text}")
After we created the Addresses successfully, the second step is to map these Addresses with their groups so we will create another script for this. This script will read columns from Groups.csv and maps every IP Address to its Group. Please don’t forget to replace FGT_IP and API_KEY with their Value.
import csv
import requests
import urllib3
from collections import defaultdict
# Discard ssl warnings
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
FGT_IP = "https://<FIREWALL_IP>"   #You can use HTTP also
API_KEY = "<API_TOKEN>"            #The token we generated in FortiGate
headers = {
    "Authorization": f"Bearer {API_KEY}",
    "Content-Type": "application/json"
}
# Group -> list of members
groups = defaultdict(list)
# Read CSV
with open('Groups.csv', 'r') as file:
    reader = csv.DictReader(file)
    for row in reader:
        group = row['Group'].strip()
        name = row['Name'].strip()
        groups[group].append(name)
# Send API Requests
for group, members in groups.items():
    data = {
        "name": group,
        "member": [{"name": m} for m in members]
    }
    url = f"{FGT_IP}/api/v2/cmdb/firewall/addrgrp"
    response = requests.post(url, headers=headers, json=data, verify=False)
    if response.status_code == 200:
        print(f"[✓] Group '{group}' created with members: {members}")
    else:
        print(f"[✗] Failed to create group '{group}': {response.status_code} - {response.text}")
Automating this migration saved me hours of manual work and significantly reduced the chance of human error. Instead of clicking through hundreds of address objects and groups, a few Python scripts and CSV files did the job in minutes. If you found this article useful, feel free to share it or connect with me.
