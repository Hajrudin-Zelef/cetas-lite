---
id: collect-261001-automatisation-infra/automatisation-infra/t5-developers-apis-dashboard-api-python-library-v1-36-0-1-ef-b8-8f-e2-83-a3-ef-b-f85a7929
title: "t5-developers-apis-dashboard-api-python-library-v1-36-0-1-ef-b8-8f-e2-83-a3-ef-b-f85a7929"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/t5-developers-apis-dashboard-api-python-library-v1-36-0-1-ef-b8-8f-e2-83-a3-ef-b-f85a7929.md
source_anchor: ""
source_lines: [1, 98]
sha256: 8505036f7cdf781e2af01bbf38f8676fa7e81a12a090142836ac5c9c479e72ee
---

# t5-developers-apis-dashboard-api-python-library-v1-36-0-1-ef-b8-8f-e2-83-a3-ef-b-f85a7929

Dashboard API Python Library v1.36.0 🐍✌🏼1️⃣⏺️3️⃣6️⃣⏺️0️⃣
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-04-2023 06:30 PM
Library version 1.36.0 includes the latest capabilities released in dashboard API release 1.36.0. It also includes a number of bugfixes (read more on GitHub.)
Install or upgrade with:
pip install --upgrade meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-06-2023 01:52 PM
You seem to be powering through the updates. Are you banned form taking leave?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-09-2023 05:41 PM
So, I believe we found a bug in this version, it might have been in v1.35 as well but when running the Organization Networks report and if the Network Notes field has more than one line of text, it will cause the JSON to CVS conversion to jump several rows between the lines of text from the "Notes" field which throws the data in the columns off as well.
If you change the API version to v0, then the JSON output converts properly because v0 doesn't include the "notes" field.
Example v1 output for that field: "notes": "10.182.0.0/16\nnew office",
I could run a find and replace on the JSON data before the conversion to remedy this but thats not something you can do when running the Meraki Tool for GSheets.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-18-2023 10:47 AM
Hi @DaRusalka , thanks for reporting this!
I think the GSheets extension doesn't use the Python library. Do you think it could be specific to how Gsheets handles newline characters (e.g. \n)?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-18-2023 11:20 AM
Hi @John-on-API ,
Its not limited to GSheets. We first ran across it when one of my team members ran the Postman collection and then tried to convert the JSON data to csv via an online conversion site. The resulting csv was all messed up. Initially we thought it was him since he has had issues with Postman in the past but several others ran it as well and we were getting similar results, even when using different conversion techniques. After messing with the data for awhile I was able to trace it back to 'base url' v1 of the newer API versions including the "notes" field now (think the last time we ran this report, we were using API v1.3.4) and when data in that field included the \n because there were multiple lines of data for the notes.
So long story short, we found that if we wanted to get the JSON data to convert correctly to CSV, we had to change the 'base url' version to v0 so it would exclude the "notes" field.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-18-2023 11:24 AM
Before converting the JSON, do you see any errors in the JSON? I wonder if this is a CSV conversion issue -- the library doesn't do any CSV conversion.
If the API is returning invalid JSON, then that could be an issue with the API. If the API is returning valid JSON, but a CSV converter doesn't successfully parse it, that's probably an issue with the converter.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-18-2023 12:02 PM
No JSON errors. If anything, its the way the data is stored in the "notes" field.
And its only now an issue because the newer API versions are now including the 'notes' field in the json.response.
I have a way to work around it. Just means I have to run the reports for my team now.
And it does mean that Meraki Tools dynamic reports are useless to us now
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-25-2023 08:15 PM
Hi @DaRusalka thanks for reporting this. I talked to the developer of the plugin and he thinks he has a fix in the works. In the future, if you have issues with the Gsheets plugin please feel free to start a new thread, since it's unrelated to the Python library.
