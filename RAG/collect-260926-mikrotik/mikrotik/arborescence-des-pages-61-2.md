---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-61-2
title: "System Requirements"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2024-08-25", "2024-09-24"]
keywords: ["license", "licenses", "training"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-61.md
source_anchor: ""
source_lines: [142, 290]
sha256: e85e8e756fcb8b0fca0dcfd836323b33159f8b8a52ca6e5498878d121a0e35f3
---

# System Requirements

It is only possible to license an expired CHR instance using a Prepaid key.

## Prepaid Key

A Prepaid Key is a type of license key you can purchase in advance for MikroTik products, such as the CHR, or convert into a license key to apply to an x86 system's Software ID. It allows you to buy a license without immediately assigning it to a specific device. Once you have a Prepaid Key, you can use it to upgrade a CHR or later convert it into a license key by providing the device's Software ID.

## How to Purchase a Prepaid Key to License a CHR

1. Go to mikrotik.com and log in to your account.
2. Access the "Purchase a RouterOS License Key" Section.
3. Choose the desired license Key Level;
4. Select Key Type.
5. Select the key type: "Prepaid key";
6. Input the quantity of prepaid keys you wish to purchase;
7. Select Optional Key Features:
  - Choose any additional features you might need for your key.
8. Press the "Place key in the cart" button.
9. Click "Proceed to checkout" to complete your purchase.

**Review and Complete Your Purchase**

- Review your order details.
- Proceed with payment using **Credit Card** (CC) or**PayPal** .

Congratulations! You have successfully purchased a Prepaid Key.

## Getting and Upgrading the License

After the initial setup, a CHR instance will be assigned a free trial license. You can upgrade this license to a higher tier through your MikroTik account. All license management, including upgrades, is handled on the account server.

Note that you can upgrade to any tier except for *p-unlimited*, which is already the highest tier.

## Initial Upgrade from Free to P1 License Level or Higher

Initial upgrade from the *free* tier to anything higher than that incurs CHR instance registration on the account server.

To do that you have to enter your MikroTik.com username and password and the desired license level you want to acquire.

To upgrade from the *free* tier to a higher license level, you need to register the CHR instance on the account server. Enter your MikroTik username and password, then select the desired license level to complete the upgrade.

As a result, a CHR System ID will be assigned to your account on the account server, and a 60-day trial will be created for that System ID. There are two ways to obtain a license: using WinBox or the RouterOS command-line interface.

### Upgrade license level using WinBox

**(System -> License menu):**

### Upgrade license level using the command-line interface

```
[admin@MikroTik] > /system license print 
  system-id: 6lR1ZP/utuJ
      level: free
[admin@MikroTik] > /system/license/renew
account: mymikrotikcomaccount
password: *********************
level: p1
  status: done
[admin@MikroTik] > /system/license/print 
         system-id: 6lR1ZP/utuJ
             level: p1
  limited-upgrades: no
   next-renewal-at: 2024-08-25 13:18:06
       deadline-at: 2024-09-24 13:18:06
```
## Payment

To acquire a higher-level trial, set up a new CHR instance, renew the license, and select the desired level.

To upgrade from a Trial license to a Paid one, go to the MikroTik account server and choose "All CHR keys" in the "CHR LICENCES" section.

The list of your CHR instances and their corresponding licenses will be displayed.

To upgrade from a Trial to a Paid license, click "Upgrade", select the desired license level (which can differ from the trial license level), and click "Upgrade" button.

If there are **Prepaid keys** available, it is possible to use it for CHR - press "**Pay using Prepaid key**". If there are no Prepaid keys or you do not want to use them, press "Proceed to checkout".

Choose the payment method: It is possible to pay using a credit card (CC) or PayPal.

## License Update

In the System-License menu, the router will indicate "**next-renewal-at**" - the time when it will reattempt to contact the server located on licence.mikrotik.com.

Communication attempts will be performed once an hour after the date on "**next-renewal-at**" and will not cease until the server responds with an error.

If the "**deadline-at**" date is reached without successfully contacting the account server, the router will consider that the license has expired and will disallow further software updates or package changes. However, the router will continue to work with the same license tier as before.

After successful communication with the license server, the dates will be updated.

## Upgrading the Level of Perpetual License

It is possible to upgrade the Level of Perpetual License from P1 to P10 or P-Unlimited. Once the upgrade is purchased at the full price, the former license will become available for later use on your account.

It is also possible to upgrade the Level of Perpetual License from P10 to P-Unlimited. Once the upgrade is purchased at the full price, the former license will become available for later use on your account.

The P-Unlimited (perpetual-unlimited) license level allows CHR to run indefinitely. It is the highest-tier license and it has no enforced limitations.

To upgrade the license level, follow these steps:

- Go to the "All CHR keys" section on your mikrotik.com account.
- Choose the CHR instance you want to upgrade and press "Upgrade".

If you want to upgrade a perpetual license to a higher level, please transfer the previous perpetual license to another CHR first. This will prevent the previous perpetual license from being lost during the upgrade process.

- Select the desired license level to upgrade to (P10 or P-Unlimited) and press "Upgrade".

- **Payment Options** :
  - If you have prepaid keys available, you can use them for the upgrade by pressing "Pay using Prepaid key".
  - If you do not have prepaid keys or prefer not to use them, press "Proceed to checkout".
- **Choose Payment Method** :
  - Choose your preferred payment method. You can pay using a credit card (CC) or PayPal.

After completing these steps, your CHR license will be upgraded to the selected level, and the previous license will be available for later use on your account.

## License Transfer

CHR installations are tied directly to the account on our website. It is possible to transfer a perpetual license to another CHR instance registered under the same account.

Licenses cannot be transferred to another account. The license transfer process requires that both the old and new CHR instances are registered under the same MikroTik account. If you need to use the CHR on a different account, a new license must be purchased for that account.

The only kind of licenses, that could be transferred to another Account is a Prepaid key. Prepaid keys got as a gift from the Training are not transferable.

To transfer purchased prepaid key navigate to "Transfer prepaid keys" in the section "ROUTEROS KEYS" on your MikroTik Account.

It is not possible to transfer the Perpetual license to an expired instance. 

If you see the error: "*This key is not eligible for transfer as there is no other valid CHR key that could be upgraded to the license level of this key."*, it means you don't have any CHR instances in Trial mode to which the current license can be transferred. 

You need to create a new CHR instance and add it to your account. Once added, you will be able to transfer the existing license to the new CHR instance.

First, register the new machine under the same MikroTik account where the old CHR is registered using the CLI command "/system license renew".

Once both the old and new CHR machines are visible in the "All CHR keys" section of your account, use the "Transfer" button to transfer the license.

- Press the "Transfer" button for the System ID you need to transfer.

- Select the System ID you are transferring to from the list.

- Press "Transfer subscription".

# Virtual Network Adapters

Fast Path is supported in RouterOS v7 for "vmxnet3" and "virtio-net" adapters.

RouterOS v6 does not support Fast Path.

# Troubleshooting

## Running on VMware ESXi

