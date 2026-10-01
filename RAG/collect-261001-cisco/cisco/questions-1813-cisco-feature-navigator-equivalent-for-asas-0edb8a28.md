---
id: collect-261001-cisco/cisco/questions-1813-cisco-feature-navigator-equivalent-for-asas-0edb8a28
title: "questions-1813-cisco-feature-navigator-equivalent-for-asas-0edb8a28"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-1813-cisco-feature-navigator-equivalent-for-asas-0edb8a28.md
source_anchor: ""
source_lines: [1, 24]
sha256: e850bca08e2f28a5c79a901ba6f75ea2dc7a42429c563a5ea958c7a6bf597e60
---

# questions-1813-cisco-feature-navigator-equivalent-for-asas-0edb8a28

I've recently been using the Cisco Feature Navigator to assess whether or not to upgrade IOS images (and compare them), but it doesn't support ASA software. Is there an easy way to compare images for ASA's like with CFN? If not, how do you do this without manually combing release notes?
1 Answer 1
The feature set of the ASA changes so very rarely, it's not something Cisco has ever put any effort into creating. The few changes that have occurred are between major revisions, eg.
- 8.2 to 8.3 -> major nat configuration change
- somewhere in the v7 era "hair-pinning" became possible
- SSLVPN (webvpn) [v7+]
- IPv6 support (the joke that it is)
- transparent mode
(I'm sure there's more, but that all I can think of!)
The command reference generally tells in which versions it's valid.
The only reason I've updated any over the last ~5 years is security bug fixes.
- 
        Hey thanks Ricky, yeah all I was aware of was the changes to how NAT was configured, but I couldn't find an authoritative way to see the other versions. Where do you go to see ASA software releases?A L– A L2013-06-12 17:27:14 +00:00Commented Jun 12, 2013 at 17:27
- 
            
            
- 
        1software.cisco.com/download/…Ricky– Ricky2013-06-12 19:56:14 +00:00Commented Jun 12, 2013 at 19:56
- 
        3Cisco has a document summarizing all major new features from 7.0 up to the current 9.1 release.James Sneeringer– James Sneeringer2013-06-13 18:29:45 +00:00Commented Jun 13, 2013 at 18:29
- 
        Thanks Ricky/James, I actually came to post the software download section (which for reference to others reading this - shows every release for a given product, so you can easily identify the latest and greatest) - which you posted. The roadmap is a cool document as well - thanks for that! Marking answered.A L– A L2013-06-13 18:47:11 +00:00Commented Jun 13, 2013 at 18:47
- 
        Once you get an idea of what is around it is work looking at Cisco365 sessions from Cisco Lives. BRKSEC-2021,1661,2020. Also - Architecture changes. 8.6 brought multiple modes, multiple contexts. 9.0 brought clustering and the pros and cons brought with it.Pandom– Pandom2013-06-14 06:41:18 +00:00Commented Jun 14, 2013 at 6:41
