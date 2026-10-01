---
id: collect-261001-automatisation-infra/automatisation-infra/t5-developers-apis-python-import-csv-api-m-p-230951-highlight-true-b28f8337
title: "Access data"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/t5-developers-apis-python-import-csv-api-m-p-230951-highlight-true-b28f8337.md
source_anchor: ""
source_lines: [1, 159]
sha256: 0a3191d75f10b443cb969ed923278d1bdc65e563413dc4780ec82a0a8e1d5518
---

# Access data

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-02-2024 01:40 PM
I am trying to import a csv of switch port configs.
I am able to print the csv from the below script
def read_csv_to_dict_list(filename):
data = []
with open(filename, 'r') as csvfile:
reader = csv.DictReader(csvfile)
for row in reader:
data.append(row)
return data
data = read_csv_to_dict_list("C:\\temp\\Sample Switch.csv")
# Access data
for row in data:
print(f"Port: {row['portId']}, Name: {row['name']}, Tags: {row['tags']}, Type: {row['type']}, Vlan: {row['vlan']}, VoiceVlan: {row['voiceVlan']}")"
how do I use this with "updateDeviceSwitchPort"
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-04-2024 11:28 AM
!!UPDATE!!
I got it to work.
I added an if statement to replace empty cells with None
def read_csv_to_dict_list(filename):
data = []
with open(filename, 'r') as csvfile:
reader = csv.DictReader(csvfile)
for row in reader:
if row['Port_voiceVlan'] == '':
row['Port_voiceVlan'] = None
data.append(row)
return data
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-02-2024 01:47 PM
Take a look on those links.
https://github.com/meraki/automation-scripts
https://github.com/meraki/automation-scripts/blob/master/update_ports.py
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-02-2024 10:50 PM
You'll need to reference each field, individually.
Based on your print statement, something like
destSerial = 'Q2QN-9J8L-SLPD'
for row in data:
    response = dashboard.switch.updateDeviceSwitchPort(
        destSerial, row['portId'], 
        name=row['name'],
        tags=row['tags'],
        type=row['type'],
        vlan=row['vlan'], 
        voiceVlan=row['voiceVlan']
    )
LinkedIn ::: https://blog.rhbirkelund.dk/
Like what you see? - Mark as helpful ## Did it answer your question? - Mark it as a Solution
All code examples are provided as is. Responsibility for Code execution is solely your own.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-03-2024 07:37 AM
Thank you.
I get an error
line 41, in <module>
serial, Port=row["portId"],
~~~^^^^^^^^^^
KeyError: 'portId'
** I changed the serial line. I have the SN as a variable already
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-03-2024 07:47 AM
The error message KeyError: 'portId' indicates that the key 'portId' is not found in the dictionary row. The CSV file does not have a column named 'portId'. You need to check the CSV and ensure that the column names match the keys you are using in your code.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-03-2024 07:53 AM
line 1 of the csv is Port,name,tags,type,vlan,voiceVlan
**Edit** I updated the csv. was getting confusing.
If I print the dictionary keys..
dict_keys(['Port_num', 'Port_name', 'Port_tag', 'Port_type', 'Port_vlan', 'Port_voiceVlan'])
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-03-2024 08:37 AM
I figured it out, kinda
I had the dict keys swapped with the update switch API requirements
response = dashboard.switch.updateDeviceSwitchPort(
Serial, portId=row['Port_num'],
name=row['Port_name'],
type=row['Port_type'],
vlan=row['Port_vlan'],
)
I removed tags as it wants an array, and voiceVlan as it must be an integer or null
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-04-2024 11:28 AM
!!UPDATE!!
I got it to work.
I added an if statement to replace empty cells with None
def read_csv_to_dict_list(filename):
data = []
with open(filename, 'r') as csvfile:
reader = csv.DictReader(csvfile)
for row in reader:
if row['Port_voiceVlan'] == '':
row['Port_voiceVlan'] = None
data.append(row)
return data
