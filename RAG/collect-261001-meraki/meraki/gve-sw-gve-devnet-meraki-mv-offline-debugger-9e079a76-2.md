---
id: collect-261001-meraki/meraki/gve-sw-gve-devnet-meraki-mv-offline-debugger-9e079a76-2
title: ".env"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-meraki/gve-sw-gve-devnet-meraki-mv-offline-debugger-9e079a76.md
source_anchor: ""
source_lines: [87, 109]
sha256: 7637922386728d62baeecf15ed8e64b133182a0a0ff341c3059c50fded35156b
---

# .env

$ docker-compose up -d --build
This will build the containers in the background and launch them. If you have docker desktop, you should see (assuming no errors):
The docker-compose file runs 3 main python scripts:
- db.py
This script creates the database that will hold the connection and status information of the routers, switches, and cameras in the network (critical for intelligent ticketing).
Once this code completes, it will print out three empty lists representing the items in the routers, switches, and cameras tables. It will also create the database file sqlite.db.
- populate.py
The next step is to populate the database with the devices in your Meraki network, which is done with this script.
Once the code completes, it will print out the contents of the routers, switches, and camera tables in the database representing (Source Serial, Connected Device Serial, Source Status)
Optional Arguments include:
-h: Help Message
--print {camera,switch,router,ticket}: Print contents of specified table
Note: this code is written with the assumption that the switches are connected to routers and the access points are connected to switches. If this is not the case in your environment, it will likely generate an error. To remedy this, add more conditions in the file. For example, there is a condition to check if the device types connected to one another are an 'Camera' and 'switch'. Suppose you have switches connected to switches in your environment, you will need to add a condition to the code that checks if the device types are both 'switch' and then write the code to handle this condition.
- app.py (flask run )
The main web app will start and begin listening on default port 5000.
Once a webhook is received, the contents will be printed out:
A down camera will trigger the troubleshooting automation. The results of the automation are written to a log file in flask_app/logs/[alerting serial].log:
A Ticket is created in ServiceNow if the troubleshooting process is unable to revive the down camera.
The ticket content is written to a CSV File as a backup. A new CSV file is created every week in /flask_app/csv_reports with a week and year stamp containing the entries for that week.
Provided under Cisco Sample Code License, for details see LICENSE
Our code of conduct is available here
See our contributing guidelines here
Please note: This script is meant for demo purposes only. All tools/ scripts in this repo are released for use "AS IS" without any warranties of any kind, including, but not limited to their installation, use, or performance. Any use of these scripts and tools is at your own risk. There is no guarantee that they have been through thorough testing in a comparable environment and we are not responsible for any damage or data loss incurred with their use. You are responsible for reviewing and testing any scripts you run thoroughly before use in any non-testing environment.
