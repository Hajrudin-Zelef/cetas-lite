---
id: collect-261001-meraki/meraki/johshea-meraki-cloud-simulator-1aeb9d52
title: "johshea-meraki-cloud-simulator-1aeb9d52"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/johshea-meraki-cloud-simulator-1aeb9d52.md
source_anchor: ""
source_lines: [1, 17]
sha256: 25fc891274b34ac8117d18e3d89b28339fb99146aaa8597580db34aeed77ece1
---

# johshea-meraki-cloud-simulator-1aeb9d52

Locally run Python 3.5+ Flask based application that provides HTTP POST simulations of Location Scanning, WebHook Alerts and Splash Page (Captive Portal) Integrations.
To run:
- virtual envrionment:
python -m venv simulator
source simulator/bin/activate
- Install dependencies
pip install -r requirements.txt
- Go!
python meraki_cloud_simulator.py
Navigate to http://localhost:5001/go and select the simulator you'd like to use.
All of the below require a third party application to be running for these simulations to send data to. Baseline samples can be found at:
- Location Scanning sample receiver
- Webhook Sample receiver
- Captive Portal Sample
Enter number of clients and APs for the simulator provide location data, set the location of the desired map and the receiving services url and hit Launch Simulator to run. Location data for that location, AP and client set will be generated and HTTP Posted to the url provided
Select the alerts for this service to send and configure the HTTP listening service at the bottom and set as default receiver. Upon Save, sample alerts will be sent on a 10 second basis>
Enter the URL for the captive portal to be tested and the continuation URL. Upon hitting "Simulate WiFi Connection" the service will open a new tab to the captive portal as if your local client was connecting to a Meraki SSID and serving up a captive portal
