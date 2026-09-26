---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-65-2
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-65.md
source_anchor: ""
source_lines: [126, 148]
sha256: 82919446055a47c615b37c3d2d6d25fa916442736b239052cf9613644aa84ce6
---

# Summary

It is possible to add a swap space to your RouterOS device. This is useful when using containers on RouterOS to be able to run containers that require much more RAM than you RouterOS device has. You can use a disk (or a partition) as a swap space or file as as swap space.

### Swap partition

Swap partition requires you to have a disk (or partition) connected to your RouterOS device. All the disk (or partition) will be used as swap space and cannot be used for other purposes. Make sure you a high speed disk for your swap partition. Using a swap partition has a better performance than using a swap file. |

To use a disk (or partition) as a swap partition, you can use the following command:

Make sure you change `disk1` to your correct disk's name!

### Swap file

Swap file requires you to have a disk that is formatted with a file system, for example, Btrfs. Compared to the swap partition option, the whole disk (or partition) will not be used as swap space, only the swap file's size will be used on your disk (or partition). Using a swap file will have a lower performance than using a swap partition.

To create a swap file on your existing file system , you can use the following command:

Make sure you change `disk1` to your correct path, where your disk is mounted.

## Mount images

RouterOS can also mount **.iso** and **.squashfs** images directly. To mount an image, use the following command:

Make sure you change `disk1/mycopy.iso` to your correct path, where the image is located.
