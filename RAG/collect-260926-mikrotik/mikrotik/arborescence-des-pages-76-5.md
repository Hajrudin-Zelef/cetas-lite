---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-76-5
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-76.md
source_anchor: ""
source_lines: [347, 475]
sha256: 06bf69b4d547e684cb49794975c2f8b8c17ec81501ab334791c0f7d8226386f2
---

# Overview

| Property | Description | 
|---|---|
| **attributes** (*array of attributes* ; Default: ) | Custom set of **Attributes** with their values that will additionally be added to Access-Accept messages for users in this group. | 
| **comment** (*string* ; Default: ) | Short description of the group. | 
| **inner-auths** (*list of auths* ; Default: ) | List of allowed authentication methods for tunneled (outer) authentication methods. Supported inner authentication methods - *ttls-pap* ,*ttls-chap* ,*ttls-mschap1* ,*ttls-mschap2* ,*peap-mschap2* . | 
| **name** (*string* ; Default: ) | Unique name of the group. | 
| **outer-auths** (*list of auths* ; Default: ) | List of allowed authentication methods. Supported outer authentication methods - *pap* ,*chap* ,*mschap1* ,*mschap2* ,*eap-tls* ,*eap-ttls* ,*eap-peap* ,*eap-mschap2* . | 

# User Profiles

**Sub-menu:** `/user-manager user-profile`

This menu assigns users a profile and tracks the status of the profile. A single user can have multiple profiles assigned, however, only one can be used at the same time. A user will seamlessly be switched to the next profile when the currently active profile expires without dropping the user's session.

**Properties**

| Property | Description | 
|---|---|
| **profile** (*profile* ; Default: ) | Name of the profile to assign for user. | 
| **user** (*user* ; Default: ) | Name of the user to use particular profile. | 

**Read-only properties**

| Property | Description | 
|---|---|
| **end-time** (*datetime* )  | Date and time the **User Profile** will expire. | 
| **state** (*running active* \| running \|*used* )  | The current state of the **User Profile** .*Running active -* currently used profile by the user.*Running* - a profile is ready to be used.*Used* - an expired profile that can no longer be activated. | 

**Commands**

| Property | Description | 
|---|---|
| **activate-user-profile** () | Make a **User Profile** entry active immediately. | 

# WEB Interface

Each user has access to his personal profile using a WEB interface. The WEB interface can be accessed by adding "/um/" directory to the router's IP or domain, for example, http://example.com/um/. Note that the WEB interface is affected by IP Services "www" and "www-ssl". The WEB interface can be customized using CSS, JavaScript, and HTML.

**Customizable file reference**

| File | Description | 
|---|---|
| **css/login.css** | Cascading style sheet file used in login prompt page. | 
| **css/user.css** | Cascading style sheet file used in user's profile page. | 
| **img/PayPal_mark_37x23.gif** | PayPal logo image. | 
| **img/ajax-loader.gif** | Loading gif while processing page switching. | 
| **img/mikrotik_logo.png** | MikroTik logo that is displayed on all pages. | 
| **js/generic.js** | Javascript file used on all pages. | 
| **js/login.js** | Javascript file used in login prompt page. | 
| **js/user.js** | Javascript file used in user's profile page. | 
| **user/login_dynamic.html** | Layout of the login prompt page. | 
| **user/user_dynamic.html** | Layout of the user's profile page. | 

# Application Guides

## Batch user creation

It is possible to create multiple new users with randomly generated usernames and passwords. For example, the following command will generate 3 new users with 6 lowercase symbols as the username and 6 lowercase, uppercase, and numbers as the password.

The command generated users can be seen by printing the user's table:

## Providing NAS with custom RADIUS attributes

It is possible to send additional RADIUS attributes during the authentication process to provide NAS with custom information about the session, such as what IP address should be assigned to the supplicant or what address pool to use for address assigning.

### Static IP address for a user

To assign the end user a static IP address, *Framed-IP-Address* attribute can be used. When using static IP address allocation, *shared-sessions* must be set to 1 to prevent cases when a user has multiple simultaneous sessions, but there is only one IP address. For example:

### Specifying address pool for a group of users

We can group up multiple similar users and assign RADIUS attributes to all of them at once. First of all, create a new group:

The next step is to assign a user to the group:

In this case, an IP address from *pool1* will be assigned to the user upon authentication - make sure *pool1* is created on the NAS device.

## Using TOTP (time-based one-time password) for user authentication

User Manager supports time-based authentication token addition to the user's password field that is regenerated every 30 seconds.

OTP depends on the clock, so make sure time settings are configured correctly.

TOTP works by having a shared secret on the supplicant (client) and the authentication server (User Manager). To configure TOTP on RouterOS, simply set the *otp-secret* for the user. For example:

To calculate the TOTP token on the supplicant side, many widely available applications can be used, for example, Google Authenticator or https://totp.app/. Adding mysecret to the TOTP token generator will provide a new unique 6-digit code that must be added to the user password.

The following example will accept the user's authentication with a calculated TOTP token added to the common password until a new TOTP token is generated, for example,

## Exporting user credentials

### **Printable login credentials for a single user**

To generate a single user's printable voucher card, simply use the *generate-voucher* command. Specify the RouterOS ID number of the user or use the *find* command to specify a username. A template is already included in User Manager's installation available in the Files section of your device. You can customize the template for your needs.

The generated voucher card is available by accessing the router using a WEB browser and navigating to */um/PRIVATE/GENERATED/vouchers/gen_printable_vouchers.html*

By default, the printable card looks like this:

To access the PRIVATE path of the /um/ directory by the WEB browser, *private-username* and *private-password* must be configured. See **Settings** section.

It is possible to use different variables when generating vouchers. Currently, supported variables are:

$(username) - Represents User Manager username

$(password) - Password of the username

$(userprofname) - Profile that is active for the particular user

$(userprofendtime) - Profile validity end time if specified

### Multiple user credential export

It is possible to generate a CSV or XML file with multiple or all user credentials at once by using the *export.xml* or *export.csv* as *voucher-template*.

The command generates an XML file *um5files/PRIVATE/GENERATED/vouchers/gen_export.xml* which can either be accessible by the WEB browser or any other file access tools.

## Generating usage report

In cases where presentable network usage information is required by companies billing or legal team an automated session export can be created using the *generate-report* command. The command requires an input of the report template - an example of the template is available in *um5files/PRIVATE/TEMPLATES/reports/report_default.html*. Example of the report generation:

The generated report is available by accessing the router using a WEB browser and navigating to */um/PRIVATE/GENERATED/reports/gen_report_default.html*

## Purchasing a profile

After logging into the user's private profile by accessing the router's */um/* directory using a WEB browser, for example, http://example.com/um/, he will be able to see all available **Profiles** in the respective menu. Profiles that have specified *price* values will have a *Buy this Profile* button available.

After pressing the *Buy this Profile* button, the user will be asked to choose from available transaction service providers (currently only PayPal is available) and later redirected to PayPal's payment processing page.

