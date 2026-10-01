---
id: collect-261001-huawei/huawei/index-php-documentation-howtos-huawei-ar1000v-d7e0dc90
title: "index-php-documentation-howtos-huawei-ar1000v-d7e0dc90"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/index-php-documentation-howtos-huawei-ar1000v-d7e0dc90.md
source_anchor: ""
source_lines: [1, 7]
sha256: 09bbfa4cd35156a27e07d8c44e9e9b0abbd35cee77f9ebbb51d2ea525a16cccb
---

# index-php-documentation-howtos-huawei-ar1000v-d7e0dc90

Versions this guide is based on:
| EVE Image Foldername | Downloaded image | Version | vCPU | RAM | HDD Format | Console | Interfaces | 
|---|---|---|---|---|---|---|---|
| huaweiar1k-5.170 | var_allinone.img | 5.170 | 1 | 2048 | hda | vnc | virtio 6 | 
| Instructions | 
|---|
| Other versions should also be supported following bellow’s procedure. This how to is tested for image version Huawei AR1000v-5.170 1. SSH to EVE and login as root, from cli and create image directory: mkdir /opt/unetlab/addons/qemu/huaweiar1k-5.170 2. Upload the var_allinone.img image to the EVE /opt/unetlab/addons/qemu/huaweiar1k-5.170/ using, for example, FileZilla or WinSCP. 3. From cli go to created directory: cd /opt/unetlab/addons/qemu/huaweiar1k-5.170/ 4. Rename source original filename to virtioa.qcow2: mv var_allinone.img hda.qcow25. Fix permissions /opt/unetlab/wrappers/unl_wrapper -a fixpermissions |
