---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e-2
title: "docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["liability", "license"]
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e.md
source_anchor: ""
source_lines: [134, 322]
sha256: 85c73db9eac1673924eef24b8caa1c05f069519d9a26f8f2b618d1732860b4dd
---

# docs-network-security-tutorials-opnsense-security-and-hardening-best-practice-gu-6899cb5e

1. 
Connect your OPNsense node via SSH or console.
2. 
Login using default root credentials. The default credentials are "root" and "opnsense". This will display the console menu similar to the one given below: ```
*** OPNsense.localdomain: OPNsense 23.1 ***
LAN (vtnet1) -> v4: 10.10.10.1/24
MyWireGuard (wg0) -> v4: 10.45.45.2/24
WAN (vtnet0) -> v4/DHCP4: 192.168.0.34/24
HTTPS: SHA256 3E 1A C6 00 C6 D0 FE B9 B8 39 74 07 3D AC 03 02
75 1A DE 02 92 83 7B 3E F9 AC E7 70 75 DE 89 D1
SSH: SHA256 NESuGb+GNjFA9egfXIvbmqm2FE1Og6r431BE2zZRYn0 (ECDSA)
SSH: SHA256 MVpoHjreB+CQu0kE8t3E9U0/kBYzORr+OlDlJqakaZU (ED25519)
SSH: SHA256 5q/ESIEfJsO29/1BSlD1rRg1tdV5nfj8x/p+1RpU6NQ (RSA)
0) Logout 7) Ping host
1) Assign interfaces 8) Shell
2) Set interface IP address 9) pfTop
3) Reset the root password 10) Firewall log
4) Reset to factory defaults 11) Reload all services
5) Power off system 12) Update from console
6) Reboot system 13) Restore a backup
Enter an option:
```
3. 
Type `3` for selecting`3) Reset the root password` option.```
The root user login behavior will be restored to its default.
Do you want to proceed? [y/N]: y
Type a new password:
Confirm new password:
```
4. 
Press `y` for confirmation.
5. 
Type a new password.
6. 
Retype the password for verification.

You may change the default OPNsense password via Web UI by following the steps below:

1. 
Connect your OPNsense node via your favorite browser.
2. 
Login using default root credentials. The default credentials are "root" and "opnsense".
3. 
Navigate to the **System** >**Access** >**Users** .

**Figure 1.** *System Users Settings on OPNsense*

1. 
Click the `Edit` button with a pen icon next to the`root` user.
2. 
Type a new root password in the **Password** field.

**Figure 2.** *Setting New root Password on OPNsense*

1. Click the **Save and go back** button at the end of the page.

## Enabling Two-Factor Authentication

Phishing attempts and using weak or repeated passwords can compromise traditional password-based authentication systems. A compromised password allows an attacker to obtain unauthorized access to a user account since many systems are accessible from anywhere over the Internet.

Two-factor authentication, often known as two-step verification, is intended to increase the security of user accounts. Logging into a 2FA-enabled account requires the user to provide an extra authentication factor in addition to a password.

Two-factor authentication, often known as 2FA or 2-Step Verification, is an authentication system that consists of two components: a PIN/password and a token. Two-factor authentication adds an additional layer of protection to an application or service. Additionally, it is typically simple to apply and uses minimal effort. OPNsense provides full support for two-factor authentication (2FA) throughout the system using Google Authenticator and you may easily enable 2FA on your OPNsense node. The subsequent OPNsense services support 2FA:

- 
Virtual Private Networking (OpenVPN & IPsec)
- 
Caching Proxy
- 
OPNsense Graphical User Interface
- 
Captive Portal

## Configure Country Blocking

There is no one-size-fits-all solution to country blocking, GeoIP blocking, as the countries that should be prohibited on a firewall depends on the organization's individual needs and objectives. Country blocking is used for the following reasons:

- 
To prevent users from accessing content that is illegal or prohibited in the country.
- 
To comply with government regulations or restrictions on specific types of content.
- 
To avoid potential legal liability for hosting or transmitting illegal or prohibited content.
- 
To protect against cyberattacks originating from specific countries.

Before determining which countries to restrict on their firewall, organizations should thoroughly evaluate their unique needs and objectives.

You may configure GeoIP blocking on your OPNsense firewall by following the 3 main steps:

1. 
Generate MaxMind GeoIP License Key
2. 
Define GeoIP Alias
3. 
Define Firewall Rule for Country Blocking

Each of these steps will be explained in more detail below.

### How to Generate MaxMind GeoIP License Key?

On OPNsense, you can block or allow one or more countries or entire continents. OPNsense accomplishes this by utilizing the MaxMind GeoIP database, which requires a license key. MaxMind, an industry leader in the accuracy of IP geolocation provides and maintains lists that are used by OPNsense. The MaxMind license key is completely free. The MaxMind License Key field description includes a link to the MaxMind registration page.

Websites host content and media on servers all over the world, so be cautious about blocking too much. Inadvertently blocking some of these IP addresses may result in broken websites or unavailable downloads.

To obtain your license key, fill out the registration form on the MaxMind sign-up page.

**Figure 3.** *MaxMind GeoLite2 Sign Up page*

After creating your account and verifying your email, you may proceed to generate your license key. The license key is necessary for OPNsense to automatically download the database.

Under your MaxMind account's settings, there should be a Manage License Keys option.

**Figure 4.** *MaxMind Managing license keys*

Here, you may produce a new key, describe it, and verify it.

**Figure 5.** *Generating MaxMind License Key*

### How to Define GeoIP Alias?

To download the GeoIP database from MaxMind you may follow these steps:

1. Navigate to the **Firewall** >**Aliases** >**GeoIP Settings** .

**Figure 6.** *GeoIP Alias Settings on OPNsense*

1. 
Type `https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country-CSV&license_key=YOUR_LICENSE_KEY&suffix=zip` where`YOUR_LICENSE_KEY` is your newly generated MaxMind License key in the URL field.
2. 
Click **Apply** button to start downloading the GeoIP database.**Last updated** and**Total number of ranges** fields should be updated.

**Figure 7.** *Downloading MaxMind GeoIP database*

1. 
Switch back to the **Aliases** tab on OPNsense Web UI.
2. 
Click **Add** button with`+` icon to create an alias.
3. 
Type a descriptive name, such as `NorthKorea` in the**Name** field.
4. 
Select **GeoIP** from the**Type** drop-down menu. You may select both`IPv4` and`IPv6` .
5. 
You may leave **Category** field empty or select a category if you have defined one before.

**Figure 8.** *Enabling GeoIP alias*

1. 
Select **Korea (North)** in the**Countries** drop-down menu next to the`Asia` Region.
2. 
You may click the **Statistics** checkbox to maintain a set of counters for each table entry.
3. 
Type a descriptive name, such as `Block North Korea` in the**Description** field.

**Figure 9.** *Adding alias for country blocking*

1. Click **Save** to save the alias.

**Figure 10.** *Applying GeoIP alias*

1. Click **Apply** to activate the settings.

### How to Define Firewall Rule for Country Blocking?

To define a firewall rule for country blocking you may follow the following steps given below:

1. 
Navigate to the **Firewalls** >**Rules** >*WAN* *.
2. 
Click Add button with `+` icon to add a rule.
3. 
Select **Block** in the**Action** option.
4. 
Select newly defined alias, such as `NorthKorean` in the**Source** field.
5. 
Type a descriptive name, like `GeoIP Blocking` in the**Description** field.
6. 
You may leave other options as default.
7. 
Click **Save**

**Figure 11.** *Applying GeoIP Firewall Rule*

1. Click **Apply Changes** to activate the rule.

## Disable SSH Connections

If you do not need remote access, disable any services, such as SSH, that are not necessary for your network to function. Disabling SSH connection is one of the first steps you can do to strengthen OPNsense. Disabling SSH increases system security overall by limiting the number of open ports and potential attack vectors. Some rules or business policies may necessitate the deactivation of unwanted or extra services, such as SSH.

