---
id: collect-261001-meraki/meraki/listing-meraki-network-devices-using-the-api-devnet-4f2aad96
title: "listing-meraki-network-devices-using-the-api-devnet-4f2aad96"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["decode", "sandbox"]
source: docs/RAG/collect-261001-meraki/listing-meraki-network-devices-using-the-api-devnet-4f2aad96.md
source_anchor: ""
source_lines: [1, 35]
sha256: 831498b1477fa5748abfa6b2028c8f9dc2aa2c3240db95968f5afb59d62ce00e
---

# listing-meraki-network-devices-using-the-api-devnet-4f2aad96

Listing Meraki Network Devices Using the API (DevNet)
Starting with the basics of network automation involves learning how to construct code. The DevNet Associate will test against your ability to perform this task.
In this example, I will leverage the Meraki API in order to get a list of devices part of a network within an Organization.
Anyone can get started with this as Meraki provides a DevNet Sandbox for people to test their Python capabilities.
I'll be using just three things for this lab:
Postman will help us construct our script with ease. By exploring the API with Postman, we can see which API to use that will give us the results we require. Atom is my text editor of choice.
When writing the script, we'll eventually need to include a few libraries. Here's what we'll be using:
Libraries
- Requests
- json
The requests library is exactly how we're going to make our request to the Meraki API. We want to request a list of devices from a network that is part of an organization.
When we make the request, we'll get data in return and it will be in json format. We need a way to work with that json data and that's wha the json library is for.
Meraki has their API documented very well. If our goal is to get a list of devices then we simply need to find which API call needs to be made.
There is a GET request called getNetworkDevices and we're given the URL needed to make that request.
From that URL, we need to pass a network ID. To find the network ID we can list all networks in an organization.
To make a request to the Meraki API, we need to provide an API key. In this example, the API key being used was provided by Meraki from their own examples.
We have a variable defined which will contain the GET response of our request.
I'll pass the variable into the json library and decode it into a Python object which will be a list in this case. We can tell this is a list if we pass devicesXML into the type function.
Now that we have our data in a Python object, we an iterate through it. This is where we tackle the core objective of listing devices in Meraki. We'll do this with a for loop.
import requests
import json
url = "https://api.meraki.com/api/v1/networks/L_646829496481105433/devices"
payload = {}
headers = {
     'X-Cisco-Meraki-API-Key': '093b24e85df15a3e66f1fc359f4c48493eaa1b73'
 }
response = requests.request("GET", url, headers=headers, data=payload)
allDevices = json.loads(response.text)
for device in allDevices:
    print("Model: {} \t Serial: {}".format(device["model"], device["serial"]))
The output:
% python3 getDevices.py
Model: MX65      Serial: Q2QN-9J8L-SLPD
Model: MS220-8P      Serial: Q2HP-F5K5-R88R
Model: MR53      Serial: Q2MD-BHHS-5FDL
