---
id: collect-261001-general-networking/general-networking/manual-how-tos-guestnet-html-9557083d-2
title: "manual-how-tos-guestnet-html-9557083d"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-how-tos-guestnet-html-9557083d.md
source_anchor: ""
source_lines: [126, 255]
sha256: ddd6853e090cda7936b1ad9d6883391e776059448370869a00c5942b8467b7a0
---

# manual-how-tos-guestnet-html-9557083d

      <h1>Welcome to My Company Guest Network.</h1>
      <h2>Feel free to use the guest network for professional usage</h2>
      <h3>See our website for more details: <a href="https://www.opnsense.org">My Company</a></h3>
    </div>
</header>
Copy the company logo to the image directory. Now zip the template directory and upload the new template by pressing the + on the Template tab.
Download the example Template (full)
Enter a Template Name, for this example we use Company. Hit Upload ( )
To enable the captive portal on the GUESTNET interface just click on Apply.
Step 6 - Limit Guests Bandwidth
For our example we will reserve 10 Mbps down and 1 Mbps Up for the Guest Network’s Internet Access. This bandwidth will be shared evenly between connected clients.
Note
With sharing evenly we mean that if 10 users at the same time try to use as much bandwidth as possible then everyone gets 1/10th. So in our example that would be 1 Mbps down stream (download). It is also possible to limit the traffic per user see also Setup Traffic Shaping
Go to: .
Create a pipe for the Download by pressing the + in the lower right corner of the form and enter the following details:
| Enabled | Checked | 
| bandwidth | 10 | 
| bandwidth Metric | Mbit/s | 
| mask | Destination | 
| Description | pipe_10Mbps_down | 
Click Save changes. And add another pipe for the upload traffic.
| Enabled | Checked | 
| bandwidth | 1 | 
| bandwidth Metric | Mbit/s | 
| mask | Destination | 
| Description | pipe_1Mbps_up | 
Click on Save changes.
Create the traffic shaper rules. Click on the tab Rules and press the + to do so.
First toggle the advanced mode (upper left corner of the form) and then fill in the following details (leave everything not specified on defaults):
| sequence | (leave on default) | 
| interface | WAN | 
| interface 2 | GUESTNET | 
| direction | in | 
| target | pipe_10Mbps_down | 
| description | Limit Guests download to 10 Mbps | 
Click Save changes.
| sequence | (leave on default) | 
| interface | WAN | 
| interface 2 | GUESTNET | 
| direction | out | 
| target | pipe_1Mbps_up | 
| description | Limit Guests upload to 1 Mbps | 
Click Save changes.
Now click on Apply to apply the changes.
Step 7 - Test Business GuestNet
Connect your PC or laptop to the Guest Network and start your favourite browser. Enter an address to browse to and you will be presented with the Login form we created with the template in the previous step. Click on login and start browsing.
To test your traffic shaper go to a speed test site such as http://www.speedtest.net/ After testing your result should be similar to this (if your internet connection has sufficient bandwidth).
Note
Keep in mind we have only one connected client in this test, so all reserved bandwidth will be available for our client.
Royal Hotel Example
From this point we will implement the Hotel/RV Park solution. You need to follow step 1-7 first and choose the template you like to use for your guests.
This example will be for our “Royal Hotel”.
Step 8 - Add Voucher Server
To add a Voucher Server go to: and click on Add server in the top right corner of the screen.
Fill in:
| Descriptive name | Vouchers | The name for your voucher server | 
| Type | Voucher |  | 
Click on Save.
Step 9 - Create Vouchers
Go back to the Captive portal and select Vouchers (). Click on Create Vouchers in the lower right corner of the form.
Let’s create 1-day vouchers for our guests:
Enter the Validity (1 day), the number of Vouchers and a Groupname (Wi-Fi day pass, for example). For the example we create 10 vouchers. Click on Generate.
A file will be generated called Wi-Fi day pass.csv. The content of this file looks like this:
username,password,vouchergroup,validity
"IgJw@Pqf","MLi+Sb7Ak#","Wi-Fi day pass","86400"
"++?f[@i[","!m*)e(@;F,","Wi-Fi day pass","86400"
"bbtK9mBk","f/jCDL3:)b","Wi-Fi day pass","86400"
"iD%L[jLJ","I#FoZ#g!AY","Wi-Fi day pass","86400"
"+4bA\E[I","CNavt@0ck+","Wi-Fi day pass","86400"
"+,fg/\Sv","#22iIL-iQA","Wi-Fi day pass","86400"
":;Pc\N#s","Y\HuG9vAN$","Wi-Fi day pass","86400"
"00nLb=0Q","0*C_\_Nb_x","Wi-Fi day pass","86400"
"PA$J0YHF","kp!q%9;m)g","Wi-Fi day pass","86400"
"a,mCxbya","LcnCb#g/di","Wi-Fi day pass","86400"
The content are:
| username | Username the guest needs to login with | 
| password | Password the guest needs to login with | 
| vouchergroup | The name of the group you created | 
| validity | The time the voucher will be valid in seconds | 
Warning
For security reasons the plain text passwords for the vouchers are NOT stored on the firewall.
This file can be used for creating nice guest vouchers (on paper) by just merging the CSV data with Microsoft Word, LibreOffice or any other DTP/text editor.
Create something like this:
You can select a database to and remove it entirely. This way you can create a voucher database for the arrival date of guest per guest group (week, midweek, weekend, etc.) and delete the full database when the guests have left.
Note
When a voucher is activated the time will be used regardless of the user being logged in or out. For a “used time” solution use a Radius server look at Setup FreeRADIUS for accounting
Step 10 - Voucher Authentication
Enable the voucher authentication by changing the zone settings. Go to the tab Zones and select the Guest Network by clicking on the pencil icon right next to it.
Change Authenticate using from an empty field to Vouchers.
When done click Save changes and the Apply to apply the new settings.
Now users will see the login form as part of your template:
Check Sessions
To check the active sessions go to Our current session looks like this:
You can drop an active session by clicking on the trashcan.
Note
Notice the selection box at the upper right corner, with this you can select the right zone when you have configured more than one.
Check Voucher Status
You can check the validity and active status of a voucher by going to the voucher page of the captive portal () and select the correct database (Wi-Fi day pass in our example).
Note
The state valid means it is activated but still valid.
Advanced - Session popup
Let’s create a Session Popup so users can see some details about their session and Logout. For this feature we will use OPNsense’s built-in API calls.
In particular we will use the following API call (for zone id 0):
/api/captiveportal/access/status/0/
The response on this API call looks like this (for an active session):
{"userName":"IgJw@Pqf",
"macAddress":"10:dd:b1:bc:75:46",
"acc_session_timeout":14095,
"authenticated_via":"Vouchers",
"packets_out":2834,
"bytes_in":512869,
"last_accessed":1457527526,
"zoneid":0,
"sessionId":"npd5bd6SIVQeMfIbWBdong==","
startTime":1457526930.1719,
"bytes_out":1322351,
"ipAddress":"192.168.200.100",
"packets_in":3181,
"clientState":"AUTHORIZED"}
It would go a bit to far to explain standard HTML and JavaScript used for our simple popup, but a full demo template can be downloaded:
Download the example Template (with popup)
The demo includes a new file called session_popup.html with all the logic to show the time left on the voucher and a logout button. As well as a simple update to our index.html page to call the popup on a successful login. The latter looks like this (shown with a bit of context):
// redirect on successful login
if (data['clientState'] == 'AUTHORIZED') {
    window.open("session_popup.html","Session Status & Logout","width=400, height=400");
Advanced - CLI Session Status
OPNsense has a very powerful CLI that is particularly useful for debugging purposes. For this example we will use the cli to list the status off all active sessions.
Type the following on the cli prompt to do so (for zone id 0):
configctl captiveportal list_clients 0
The output will be something similar to this:
