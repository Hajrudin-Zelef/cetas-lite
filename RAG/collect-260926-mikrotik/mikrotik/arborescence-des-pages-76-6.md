---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-76-6
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-76.md
source_anchor: ""
source_lines: [476, 543]
sha256: 48a0a5beccf2fc5e7c3b6b03f776b4da0fb25edb9a34867bf8a69128cc98bcd7
---

# Overview

When the payment is completed, the User Manager will ask PayPal to approve the transaction. After approval, the profile is assigned to the user and is ready to use.

## Migrating from RouterOS v6


When you upgrade your User Manager router from RouterOS v6 to the v7 the new User Manager will work with new database files and configuration. To continue using the old user, router, profile, etc. configuration you must manually execute the migrate command. To do so you must have files from the old User Manager server folder "user-manager" present. The folder can be renamed, but all the contents from the old installation must be transferred to the new v7 installation (you can move the old configuration from one router to another router with v7, you must copy "user-manager" folder). After that, all you need to do is execute this command - "/user-manager/database/migrate-legacy-db database-path=<path_to_old_user_manager_folder>".

The import process will try to convert such configuration - users, profiles, user-profiles, limitations, profile-limitations, user-counters, routers, and sessions.

# Application Examples

## Basic L2TP/IPsec server with User Manager authentication

**User Manager configuration**

Start off by enabling User Manager functionality.

Allow receiving RADIUS requests from the localhost (the router itself).

Next, add users and their credentials that clients will use to authenticate to the server.

**Configuring RADIUS client**

For the router to use the RADIUS server for user authentication, it is required to add a new RADIUS client that has the same shared secret that we already configured on User Manager.

**L2TP/IPsec server configuration**

Configure the IP pool from which IP addresses will be assigned to the users and assign it to the PPP Profile.

Enable the use of RADIUS for PPP authentication.

Enable the L2TP server with IPsec encryption.

That is it. Your router is now ready to accept L2TP/IPsec connections and authenticate them to the internal User Manager.

## Two factor authentication for RouterOS user login (MFA/2FA)

As User-Manager supports TOTP (time-based-one-time password), it is possible to setup so called MFA authentication for different services. Here will will look into RouterOS user authentication over User Manager (radius) with time based password, that is changed every 30 seconds.

**Here are the necessary configuration options on your MikroTik router,**

Enable to use RADIUS for /user menu, and set default-grop from /user group menu. Keep in mind that local /user database is checked first, and then RADIUS is contacted.

 

/user/aaa/set use-radius=yes default-group=full

Enable radius client for login service. As we run User Manager on the same router, 127.0.0.1 is used,

/radius/add address=127.0.0.1 service=login secret=mystrongsecret

**Here are the configuration steps for User Manager**,

Make sure you have added your managed devices to "Routers" menu,

/user-manager/router/add name=myrouter address=127.0.0.1 shared-secret=mystrongsecret

Add user to User-Manager user table with OTP secret parameter. Few more steps are required for proper OTP configuration.

Pick OTP-secret name and convert it to base32 format (there are plenty of online converters from utf-8 to base32 format). For my configuration I use "mysupersecret", that in base32 would be NV4XG5LQMVZHGZLDOJSXI===

/user-manager/user/add name=mikrotik password=mysuperpassword otp-secret="NV4XG5LQMVZHGZLDOJSXI==="

Note, than in your favorite authenticator app, you will need to set user this key manually "NV4XG5LQMVZHGZLDOJSXI===", when adding new time password instance.

To login to your MikroTik device, open Winbox/Console and connect to your router address, use login:mikrotik and password:mysuperpasswordxxxxxx, where xxxxxx is 6 digit code from your favorite authentication app.

Password is changed every 30 seconds and it is available from your favorite app.
