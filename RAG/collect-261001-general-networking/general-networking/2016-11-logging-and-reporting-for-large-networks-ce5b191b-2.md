---
id: collect-261001-general-networking/general-networking/2016-11-logging-and-reporting-for-large-networks-ce5b191b-2
title: "2016-11-logging-and-reporting-for-large-networks-ce5b191b"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-general-networking/2016-11-logging-and-reporting-for-large-networks-ce5b191b.md
source_anchor: ""
source_lines: [194, 346]
sha256: c38678bd90503069e2d8ceeccbb7686f328de5aeaf9e6946aa50e1c85131c3c7
---

# 2016-11-logging-and-reporting-for-large-networks-ce5b191b

**C****on****f****i****gu****r****i****n****g logging to the FortiCloud server**

The FortiCloud server can be used as a redundant backup, or your primary logging solution. The following assumes that this service has already been registered, and a subscription has been purchased for expanded space. The following is an example of how to these settings are configured for a network’s log configuration. You need to have access to both the CLI and the web-based manager when configuring uploading of logs. The upload time and interval settings can be configured in the web-based interface.


**T****o configure logging to the FortiCloud server**

**1****.** Go to **S****ys****t****e****m > Dashboard > Status** and click **Log****i****n** next to **Fo****r****t****i****C****l****ou****d** in the License Information widget.

**2****.** Enter your username and password, and click **O****K**. (Or register, if you have not yet done so.)

**3****.** Logs will automatically be uploaded to FortiCloud as long as your FortiGate is linked to your FortiCloud account.

**4****.** To configure the upload time and interval, go to **Lo****g & Report > Log Config > Log Settings**.

**5****.** Under the Logging and Archiving header, you can select your desired upload time.

**6****.** With FortiCloud you can easily store and access FortiGate logs that can give you valuable insight into the health and security of your network.


**M****od****i****f****y****i****n****g the default FortiOS report**

The default FortiOS report is provided to help you quickly and easily configure and generate a report. Below is a sample configuration with multiple examples of significant customizations that you can make to tailor reports for larger networks.


**C****r****ea****t****i****n****g datasets**

You need to create a new dataset for gathering information about HA, admin activity and configuration changes.

Creating datasets requires SQL knowledge.


**T****o create the datasets**

**1****.** Log in to the CLI.

**2****.** Enter the following command syntax:

config report dataset edit ha

set query “select subtype_ha count(*) as totalnum from event_log

where timestamp >= F_TIMESTAMP (‘now’, ‘hour’, ‘-23’) and group by subtype_ha order by totalnum desc”

next

**3****.** Create a dataset for the admin activity, that includes log ins and log outs from the three FortiGate administrators.

set query “select subtype_admin count(*) as totalnum from event_log

where timestamp >= F_TIMESTAMP (‘now’, ‘hour’, ‘-23’) and group by subtype_

admin order by totalnum desc”

next

**4****.** Create a dataset for the configuration changes that the administrators did for the past 24 hours.

set query “select subtype_config count(*) as totalnum from event_log

where timestamp >= F_TIMESTAMP (‘now’, ‘hour’, ‘-23’) and group by subtype_

config order by totalnum desc”

end

next


**C****r****ea****t****i****n****g charts for the datasets**

**1****.** Log in to the CLI.

**2****.** Enter the following to create a new chart:

config report chart edit ha.24h

set type table

set period last24h set dataset ha

set category event set favorite no

set style auto

set title “24 Hour HA Admin Activity”

end


**U****p****l****o****a****d****i****n****g the corporate images**

You need to upload the corporate images so that they appear on the report’s pages, as well as on the cover page. Uploading images is only available in the web-based manager.


**T****o upload corporate images**

**1****.** Go to **Lo****g & Report > Report > Local**.

**2****.** Select the Image icon and drag it to a place on the page.

**3****.** The Graphic Chooser window appears.

**4****.** Select Upload and then locate the image that you want to upload and upload the image.

The images are automatically uploaded and saved.

**5****.** Repeat step 4 until the other corporate images are uploaded.

**6****.** Select Cancel to close the Graphic Chooser window and return to the page.

The images can then be placed as you like by reopening the Graphic Chooser as in step 2.


**A****dd****i****n****g a new report cover and page**

You need to add a new cover for the report, as well as a new page that will display the HA activity, admin activity and configuration changes.


**T****o add and customize a new report cover**

**1****.** Go to **Lo****g & Report > Report > Local**.

**2****.** Select **C****u****s****t****o****m****iz****e**.

**3****.** In **S****ec****t****i****on****s**, select the current default report section, and enter Report Cover in the field that appears; then press Enter to save the change.

**4****.** Remove all content from the Report Cover section, and select the image icon and drag it into the main portion of the cover page; select a cover page image and then select **O****K**.

**5****.** Select the font size you want, and drag the text icon into the area beneath the image to add a title or explanation for the cover page.

**6****.** Select **S****av****e** to save the new report cover.


**T****o add and customize a new page**

**1****.** Go to **Lo****g & Report > Report > Local**.

**2****.** Select **C****u****s****t****o****m****iz****e**.

**3****.** Select **S****ec****t****i****on****s**, and select **C****r****ea****t****e New** to add a new section to the report. Name it Report Content, and press **E****n****t****e****r**, and **O****K** to close the menu.

**4****.** At the bottom of the editing window is the **S****ec****t****i****o****n** selection, where each **S****ec****t****i****o****n** is represented by a box. Select the second box.

**5****.** Edit the content for the report as you like.

For a simpler report structure, make use of the ‘FortiGate UTM Security Analysis Report’ charts, which automatically format themselves and fill in all necessary information.

For more complex reports, add headings, default and custom charts, and explanatory text.

**6****.** Select **S****av****e** to save the new report content.

The report will automatically combine all sections. You can use headers and text to more clearly separate parts of the report, and all properly configured charts have titles built-in.
