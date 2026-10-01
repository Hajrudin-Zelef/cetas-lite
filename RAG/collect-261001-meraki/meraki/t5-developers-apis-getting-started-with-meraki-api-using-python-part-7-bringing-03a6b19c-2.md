---
id: collect-261001-meraki/meraki/t5-developers-apis-getting-started-with-meraki-api-using-python-part-7-bringing-03a6b19c-2
title: "t5-developers-apis-getting-started-with-meraki-api-using-python-part-7-bringing--03a6b19c"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-meraki/t5-developers-apis-getting-started-with-meraki-api-using-python-part-7-bringing--03a6b19c.md
source_anchor: ""
source_lines: [15, 113]
sha256: d0e25bf3c82abcc651874bd90f25b153451ce160ad222bc496d858896f4b996e
---

# t5-developers-apis-getting-started-with-meraki-api-using-python-part-7-bringing--03a6b19c

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-02-2023 01:55 PM
Welcome to Part 7 of this Getting Started series! We hope you have gotten plenty of time to practice the Python concepts and got a good grasp on how to incorporate those concepts into a script to automate your basic workflows. In this part, we will bring all the concepts we learned thus far together and write a script to solve a problem.
If you missed the earlier parts of our series, please make sure you go back here first before diving in. Well, let's get started!
Consider the following scenario: You are an administrator that has access to multiple different dashboard organizations. Your manager has asked you to write a Python script to get a list of Wireless Access Points from one of those orgs called “DeLab” and display only their serial numbers and connectivity status (dormant, online, offline, or alerting). To make it more readable, the information must be displayed in a tabular format.
We have created the following flowchart to break down the different tasks involved for your reference.
Please stop here, give it some thought, and try to write a script on your own using all the concepts we have discussed in this series. There is definitely more than one way to complete this task so there is no right or wrong answer here.
We have come up with a simple solution as follows:
import meraki
from prettytable import PrettyTable
API_KEY = '6bec40cf957de430a6f1f2baa056b99a4fac9ea0'
dashboard = meraki.DashboardAPI(API_KEY, suppress_logging=True)
We start by importing a couple of libraries, namely meraki and prettytable. The meraki library is used to make API calls to the dashboard. The prettytable library will help you create a table around the data. Please note prettytable needs to be installed using pip install prettytable. We then set the API key so we can use it to make API calls throughout our script. Another thing to note is while setting up the API key using meraki.DashboardAPI() function, we have used suppress_logging=True as one of its parameters. This is to omit all the API logs from appearing on the screen so we can only see a clean output i.e. the data we expect from this script.
list_of_orgs = dashboard.organizations.getOrganizations()
def get_orgID(list_of_orgs):
for org in list_of_orgs:
if org['name'] == 'DeLab':
org_id = org['id']
return org_id
org_id = get_orgID(list_of_orgs)
Based on the flow chart above, the first task is to get a list of all organizations and then get org ID for the DeLab organization. We can do that by making an API call to get all orgs and store it in the list_of_orgs variable which will be passed as an input to a function called get_orgID(). We can then extract the org ID of DeLab org by simply running a for loop and if statement within. Remember, a function needs to be called in order for it to be executed so the last line of the code above does just that. Now the org_id variable will hold the org ID of the DeLab organization.
def get_devices(org_id):
list_of_devices = dashboard.organizations.getOrganizationDevicesAvailabilities(
org_id, total_pages='all')
return list_of_devices
org_devices = get_devices(org_id)
The next task is to get a list of devices in the DeLab organization. We can do that by defining another function called get_devices() and passing the org_id variable that we received from the previous function as its input so we only get the devices from only the DeLab organization.
You will need to use this API endpoint to get all the information about the devices in a given org along with their status and other information. Check out the Request Parameters section for this endpoint to know what sort of information you will get as a response. We are using a variable called list_of_devices to hold that response and return it back to the caller which is what the final line is for. So ultimately the org_devices variable has the list of all devices in the DeLab organization.
table = PrettyTable()
table.field_names = ['Serial', 'Status']
for device in org_devices:
if device['productType'] == 'wireless':
table.add_row([device['serial'], device['status']])
print(table)
The next task is to filter only the Wireless APs from the rest and then display their status in a tabular format. Let's start by creating a table using PrettyTable and store it in a table variable. We then add the fields to it as columns, namely Serial and Status as per our requirement. We can use the for loop to iterate through all the devices and from the Request Parameters section on the API endpoint page here we can see the serial number and device type is referred to by the serial and productTypes parameters respectively. So within the for loop, we can use the if statement to filter only the devices that have productType of wireless and then add both serial and status as a row to our existing table. Finally, we print the table and you will get the desired output.
The full code is as follows:
import meraki
from prettytable import PrettyTable
API_KEY = '6bec40cf957de430a6f1f2baa056b99a4fac9ea0'
dashboard = meraki.DashboardAPI(API_KEY, suppress_logging=True)
list_of_orgs = dashboard.organizations.getOrganizations()
def get_orgID(list_of_orgs):
for org in list_of_orgs:
if org['name'] == 'DeLab':
org_id = org['id']
return org_id
org_id = get_orgID(list_of_orgs)
def get_devices(org_id):
list_of_devices = dashboard.organizations.getOrganizationDevicesAvailabilities(
org_id, total_pages='all')
return list_of_devices
org_devices = get_devices(org_id)
table = PrettyTable()
table.field_names = ['Serial', 'Status']
for device in org_devices:
if device['productType'] == 'wireless':
table.add_row([device['serial'], device['status']])
print(table)
The output should be as follows:
That’s it for this post! I will leave it here to let you try it out yourself! Please try to come up with another creative solution for the same. Better yet, try to build something on top of it like add names of the APs to the table or count the total number of APs with status other than online for troubleshooting or anything you want for your own network. Most importantly, have fun building the scripts.
This post marks the end of our series on Getting Started with the Meraki API. It is just the beginning of your own personal journey with APIs though! Thank you for following along. We hope you found the series to be an exciting, fun, and informative introduction to the world of writing scripts. Now you have the tools you need to continue the adventure on your own! You’ll find plenty of support and resources to help you in the API & Developers board should you get stuck.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-27-2023 04:07 PM
This is a nice little series, one thing I did as I got in to it was create an example script of ways to filter and search the responses.
Like, this will create a new list from the response where the key "AP" = WAP1, that you may want to loop though and use to search a different response.
subtest = [r for r in ls_dict if (r["AP"] == "WAP1")]
and
This will just count for you how many entries there are based on a search.
number_off = len([r for r in ls_dict if (r["clientMac"] == "FFFFFFFF")])
How about one last leg of the journy to give some examples like this to help people getting in to Python / API / JSON to take the journy to the next level. May be sourced from the community, not lengthy code but single line well described examples that carry out specific by common operations.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-26-2023 02:02 PM
This is a very good series to start Meraki API programing. Thanks for the hard work.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-29-2023 06:22 PM
