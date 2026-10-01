---
id: collect-261001-meraki/meraki/r-meraki-comments-kmn9nm-idiots-guide-to-using-the-api-postman-andor-python-be57b482
title: "r-meraki-comments-kmn9nm-idiots-guide-to-using-the-api-postman-andor-python-be57b482"
domain: meraki
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-kmn9nm-idiots-guide-to-using-the-api-postman-andor-python-be57b482.md
source_anchor: ""
source_lines: [1, 29]
sha256: f14903d10b65f0e180097017dd9a3a172d8a59113767fb7f0f6f16910b99116c
---

# r-meraki-comments-kmn9nm-idiots-guide-to-using-the-api-postman-andor-python-be57b482

Idiots guide to using the API - Postman and/or Python 
        
        
        
    
    
    So about a year and a half ago I started digging into the (v0) API to see if we could automate some basic things and it was working reasonably well for what little we were doing with it.
Of course between then and now, Meraki releases the v1 API, what little I learned I lost, and now I feel like I'm starting from scratch and my old snippets of code don't seem to play ball with the v1 API.
I vaguely remember there being more idiot-friendly guides to getting started with the API on Cisco's site but damned if I can find them anymore. I swear they walked you through getting the library loaded into Postman and doing basic calls in Postman, too.
Long story short, I'm looking for a script that will spit out device name, serial number, (street) address, and notes from the dashboard and while I would appreciate someone giving me the fish, I'd also like to see if there's some good resources that can teach me to fish too.
Section des commentaires
How about a Google sheet that uses the API to get you that info with a click of a button?
https://developer.cisco.com/meraki/build/meraki-dashboard-reports-with-google-sheets/
https://community.meraki.com/t5/Developers-APIs/Meraki-Tools-for-Google-Sheets-add-on/td-p/50838
Oh that's slick! That actually perfectly solves my immediate need to get this data into a report for some people to validate.
Check out the python notebooks on github, as well as the examples if you want some good learning resources straight from Meraki:
https://github.com/meraki/dashboard-api-python
https://github.com/meraki/dashboard-api-python/tree/master/examples
https://github.com/meraki/dashboard-api-python/tree/master/notebooks
Send me an IM and I may send you over a copy of my Meraki Network Configuration tool built in Powershell that offers this and a whole lot more. I've also got some powershell report collections for Meraki that I can share with you.
This API call gives you all you need:
https://developer.cisco.com/meraki/api-v1/#!get-network-devices
I would suggest also looking int OpenPyXL so you can write the output to an Excel-file.
Spend some time using Python 3 with a Linter and use the Meraki Library to send requests to the API. A Linter can show you which arguments the Meraki Library has and makes using the library very easy.
After spending some time with Python with little to no programming experience, I’ve written a complete app in Python that allows me to stage a network from the ground up in about 30 minutes!
Did you find Meraki.io? That is where everything has moved.
This is the collection for postman. Once you launch it click 'Run in Postman' and it will put the collection into the postman client.
I normally run the command in postman to make sure I am getting what I need and then incorporate the GET/POST requests in python to automate what I need.
Most of the stuff I've done so far is templating to ensure settings are the same across the estate (content filtering/IDS on/etc)
