---
id: collect-261001-general-networking/general-networking/hc-en-us-articles-17784610114199-configuring-access-policies-and-schedules-in-un-595c6f92
title: "hc-en-us-articles-17784610114199-configuring-access-policies-and-schedules-in-un-595c6f92"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/hc-en-us-articles-17784610114199-configuring-access-policies-and-schedules-in-un-595c6f92.md
source_anchor: ""
source_lines: [1, 34]
sha256: a78aab007c2bdec471abc070c9ee9a17f9f2cadb221abcd1566d7b5434dcadfa
---

# hc-en-us-articles-17784610114199-configuring-access-policies-and-schedules-in-un-595c6f92

Configuring Access Policies and Schedules in UniFi Access
UniFi makes it easy to create tailored door access policies that seamlessly align with each user's schedule—whether for shifts, holidays, or weekends. If you're looking to create policies for guests or visitors, click here. To create door unlock schedules, click here.
| Update your UniFi OS, Access application, and Access devices to the latest versions for the newest features and optimal performance. | 
Creating Access Policies
Customize access times, locations, schedules, and user permissions.
- Navigate to Access application > Settings > Policies & Schedules > Policies > Create New.
  - Locations: Define which doors, gates, or elevators users can access
  - People: Define who the policy applies to. Learn more about adding users and groups
  - Schedule: Define when access is allowed.
  - Save as a Predefined Schedule: Save the settings for future use.
- Click Create.
Example 1: For Daily Routines
For employees with daily routines, grant all users at the New York Office access to all locations from Monday to Friday, between 9:00 AM and 5:00 PM.
Example 2: For Work Shifts
For employees working shifts, grant specific users and user groups access to specific locations on Monday, Wednesday, and Friday from 8:00 AM to 12:00 PM and 7:00 PM to 11:59 PM.
Example 3: For Holidays
Grant all users access to all locations from Monday to Friday, between 9:00 AM and 5:00 PM. However, on January 7 and May 1, only allow access from 10:00 AM to 12:00 PM.
Creating Access Schedules
Create a predefined access schedule and apply it to access policies or other schedules for easy setup.
- Navigate to Access application > Settings > Policies & Schedules > Predefined Schedules > Create New.
- Specify the following:
  - Recurring Schedule: Define the access times.
  - Holidays: Turn it On to create holidays and set custom access times.
FAQs
Is there a limit to how many access policies and schedules can be created?
No. You can create as many access policies and schedules as needed.
Can users unlock doors without an assigned access policy?
No. Users must be assigned at least one access policy to unlock doors using door unlock methods.
Can I assign access policies to visitors?
No. An access policy can only be assigned to admins and users. To configure visit schedules for one-time or recurring visitors, click here.
If a user has multiple access policies with different schedules, will access be allowed during all scheduled times?
Yes. If a user has multiple access policies with different access schedules, the system combines all allowed time periods. For example, if one policy allows access from 9 AM to 12 PM and another from 2 PM to 5 PM, the user can unlock the door during both periods (9 AM–12 PM and 2 PM–5 PM).
How do I set a custom access policy as the default, and what are its benefits?
Go to Access application > Settings > Policies & Schedules > Policies and hover over the custom policy you want to use. Click Set as Default. When a new user is created, this default access policy is automatically applied to the user.
