---
id: collect-260926-mikrotik/mikrotik/how-to-add-mikrotik-to-eve-ng-techbloat-3
title: "how-to-add-mikrotik-to-eve-ng-techbloat"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "parameters"]
source: docs/RAG/lot-mikrotik/forum/misc/how-to-add-mikrotik-to-eve-ng-techbloat.md
source_anchor: ""
source_lines: [201, 246]
sha256: 0ff11aa90e92e8c4f43ca96f7860c979a3024acf4b47a912972df4bc8fc5b3aa
---

# how-to-add-mikrotik-to-eve-ng-techbloat

## 1. Verify Image Compatibility and Placement

- Ensure you have the correct MikroTik image (e.g., CHR) compatible with your EVE-NG version.
- Place the image in the correct directory: **/opt/unetlab/addons/qemu/** .
- Make sure the image has the proper permissions: **chmod +x** your image file.

## 2. Correctly Configure the VM Profile

- Use the MikroTik image as a QEMU VM, selecting the right template.
- Verify the CPU, RAM, and disk settings match MikroTik’s recommended specs.
- Check that the NICs are configured appropriately for your network design.

## 3. Check for Proper Licensing and Image Integrity

- Ensure the MikroTik CHR license is valid and properly installed.
- Verify the image wasn’t corrupted during download by comparing checksum values.
- Re-download the image if you suspect corruption or incompatibility.

## 4. Confirm Network and Console Settings

- Configure network adapters correctly within EVE-NG to allow communication with MikroTik.
- Use the console option in EVE-NG to access MikroTik’s CLI; ensure the console port is correctly mapped.
- Adjust terminal settings if connection issues occur, including baud rate and port number.

## 5. Review Logs and Error Messages

- Check the EVE-NG logs for startup errors related to the MikroTik VM.
- Inspect the QEMU command line for misconfigurations or missing parameters.
- Address specific error messages by consulting MikroTik or EVE-NG documentation.

By systematically checking these common issues, you can efficiently troubleshoot problems with adding MikroTik to EVE-NG, ensuring a reliable virtual network environment.

## Conclusion

Integrating MikroTik into your EVE-NG environment is a straightforward process that significantly enhances your network simulation capabilities. By following the steps outlined—downloading the appropriate MikroTik image, configuring the EVE-NG server, and properly setting up the VM—you can create a robust lab environment for testing, learning, and troubleshooting.

Ensuring that you select the correct image type and match it with your EVE-NG version is critical for seamless operation. Keep in mind that MikroTik RouterOS images are often provided in .img or .iso formats, and you may need to convert or prepare these images to function correctly within EVE-NG. Proper resource allocation, such as CPU, RAM, and storage, will optimize performance and stability of your MikroTik nodes.

Additionally, understanding the networking setup—such as assigning appropriate interfaces, bridges, and VLANs—is vital to replicate real-world scenarios accurately. Utilize the EVE-NG documentation and MikroTik resources to troubleshoot common issues like connectivity errors or image recognition problems.

The Tool Desk

Outbyte PC Repair FREEClear out junk files and repair common Windows errorsFree Scan →Outbyte Driver Updater FREEFix the driver behind crashes, sound loss and screen glitchesFind Drivers →
Regular updates to both EVE-NG and your MikroTik images help maintain security and compatibility. Remember to back up your configurations and project files periodically to prevent data loss.

In summary, adding MikroTik to EVE-NG empowers network professionals and enthusiasts to design comprehensive lab environments. With careful preparation and adherence to best practices, you can leverage this integration to deepen your understanding of MikroTik networking and improve your troubleshooting skills. This setup not only supports learning but also prepares you for real-world deployments, making it a valuable addition to any network simulation toolkit.
