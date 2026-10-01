---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/kb-ubiquiti-unifi-device-stuck-in-adoption-failed-b9876616
title: "kb-ubiquiti-unifi-device-stuck-in-adoption-failed-b9876616"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/kb-ubiquiti-unifi-device-stuck-in-adoption-failed-b9876616.md
source_anchor: ""
source_lines: [1, 47]
sha256: a2ce7702bae7245870d24f778f09553380d186e7d6638b9f957a55466498e86f
---

# kb-ubiquiti-unifi-device-stuck-in-adoption-failed-b9876616

Unifi device stuck in Adoption failed
What
This can happen when the wind blows the wrong way quite frankly
This error is identifiable by the below error message
How to resolve
Sometimes this issue will go away on its own. Sometimes it requires quite an aggressive approach.
How to follow this guide
Best to follow this guide one section after another. Between each step check in the Unifi Console
Restart the controller
You can either reboot the controller, or restart the service
Reset the Access point
Open Settings > System > Device SSH Authentication
Make a note of the SSH password and Username.
SSH to the access point
When prompted, type in the password from the UI.
Do not power down the device
Doing so will brick the Access Point
You will get a shell, run the below
The SSH session will time out after about 5 minutes.
Check the console, perhaps the device says it's ready for adoption
Delete the access point from the database
This could break things badly
If something goes wrong, it could break the unifi install quite badly.
SSH to the Unifi controller Not the access point
Open a connection to the Database
Use the ace database
Using the Mac address from the UI, replace <mac> in the below and run the command
The exit the database
Restart the controller
You now need to set the inform URL on the access point
Adopt the access point
Assumes you reset the access point
The below assumes you've reset the access point
SSH to the access point
Username and password are ubnt
Set the inform for the controller
In the console it will show the device as pending adoption.
What if this didn't work?
Sorry, this is what worked for me.
Try:
- Reset the access point
- Upgrade the access point
- Restart the controller
- Delete from the database
- Restart the controller
- Set Inform
If that doesn't work, it may be worth resetting the controller all together
