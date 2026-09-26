---
id: collect-260926-mikrotik/mikrotik/how-to-add-mikrotik-to-eve-ng-techbloat-2
title: "how-to-add-mikrotik-to-eve-ng-techbloat"
domain: mikrotik
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "memory"]
source: docs/RAG/lot-mikrotik/forum/misc/how-to-add-mikrotik-to-eve-ng-techbloat.md
source_anchor: ""
source_lines: [64, 200]
sha256: 048a525af86ee0b83c9fa9ef023a22529df57f82be70efa75e424a2ff8764397
---

# how-to-add-mikrotik-to-eve-ng-techbloat

- **Verify EVE-NG Installation:** Confirm that EVE-NG is installed and operational on your server or VM. Access the EVE-NG web interface through your browser to check availability.
- **Update EVE-NG:** Keep your EVE-NG instance up-to-date. Use the command line or web interface to apply the latest updates, ensuring compatibility with MikroTik images and features.
- **Obtain MikroTik RouterOS Image:** Download the appropriate MikroTik RouterOS image in QEMU or ISO format from the official MikroTik website or trusted sources. Ensure the image is compatible with EVE-NG.
- **Transfer Image to EVE-NG Server:** Use secure methods such as SCP or FTP to upload the MikroTik image to your EVE-NG server. Place the image into the correct directory, typically`/opt/unetlab/addons/qemu/` or a relevant subfolder.
- **Set Proper Permissions:** After uploading, adjust permissions to ensure EVE-NG can access and run the image. Use commands like`chmod +x` and set ownership to the`unl` user as needed.
- **Configure Network Settings:** Verify that your EVE-NG environment has proper network connectivity and that virtualization features (VT-x/AMD-V) are enabled in your BIOS to support running MikroTik images.
- **Test Image Recognition:** Launch your EVE-NG web interface, create a new node, and verify that the MikroTik image appears in the list of available devices. If not, revisit the image placement and permissions setup.

Proper preparation ensures a smooth process when adding MikroTik to EVE-NG, minimizing troubleshooting and enabling a seamless virtual lab environment.

## Importing MikroTik Image into EVE-NG

To add MikroTik to EVE-NG, first, ensure you have a compatible MikroTik image. Typically, MikroTik RouterOS images are available in .img or .iso formats. It’s recommended to use the Cloud Core Router (CCR) images for virtualization.

## Prepare the MikroTik Image

- Download the desired MikroTik image from the official MikroTik website or your licensed source.
- Rename the image file to a simple name without spaces or special characters, e.g., *mikrotik-routeros.x86.qcow2* .

## Transfer the Image to EVE-NG Server

- Use SCP, WinSCP, or other file transfer tools to upload the image to the EVE-NG server.
- Navigate to the images directory: */opt/unetlab/addons/qemu/* .
- Create a new folder for MikroTik, e.g., *mikrotik* , if it doesn’t already exist:

`sudo mkdir -p /opt/unetlab/addons/qemu/mikrotik`
- Place the image inside this folder.

## Configure the Image and Metadata

- Rename your image to *disk.img* or similar for consistency.
- Create a *meta.conf* file in the same folder with the following content:

```
name = MikroTik RouterOS
label = mikrotik
arch = x86_64
ram = 512
cpu = 1
```
## Update Permissions and Register the Image

- Set appropriate permissions:

`sudo chmod -R 755 /opt/unetlab/addons/qemu/mikrotik`
- Restart the EVE-NG server or run the command:

`sudo /opt/unetlab/wrappers/unl_wrapper -a fixpermissions`
## Add MikroTik to Your EVE-NG Lab

- Log into the EVE-NG web interface.
- Create a new lab or open an existing one.
- Click “Add Object” and select MikroTik from the list.
- Configure the instance as needed, then start the node.

Following these steps will seamlessly integrate MikroTik RouterOS images into your EVE-NG environment, enabling advanced network simulations and testing.

## Configuring the MikroTik VM within EVE-NG

Setting up a MikroTik VM in EVE-NG provides a powerful environment for lab testing and network simulation. Follow these straightforward steps to integrate MikroTik into your EVE-NG topology successfully.

## 1. Prepare the MikroTik Image

Obtain a compatible MikroTik image, such as the “Cloud Core Router” (CCR) ISO or VDI image. Ensure the image is licensed and legally obtained. Once downloaded, upload the image to your EVE-NG server, typically placing it in the appropriate directory (e.g., /opt/unetlab/addons/qemu/).

## 2. Create a New MikroTik VM in EVE-NG

- Log into the EVE-NG web interface and select “Add Object” > “QEMU”.
- Configure basic VM details:
  - Name: Enter a descriptive name, e.g., MikroTik-1.
  - Type: Choose “qemu”.
  - Image: Select the uploaded MikroTik image.
  - Memory: Allocate appropriate RAM (e.g., 512MB to 1GB).
  - CPUs: Assign 1 or more CPUs based on your lab needs.

## 3. Configure Network Interfaces

Set up the network interfaces within the VM creation popup:

- For each interface, specify the device type (typically virtio for efficiency) and connect to the appropriate EVE-NG switch or network segment.
- Ensure at least one interface is connected for management and additional interfaces for data traffic.

## 4. Adjust Advanced Settings

Navigate to the “Advanced” tab:

- Set the “QEMU Command” if needed for additional customization, such as enabling specific hardware options.
- Configure the boot options if necessary, especially if using ISO images that require manual intervention.

## 5. Save and Deploy

Once all configurations are complete, click “Save”. EVE-NG will deploy the MikroTik VM, which may take a few moments. Power on the VM and access it via the console to complete initial setup.

## 6. Initial Configuration

Use MikroTik’s Winbox or CLI via the console port to configure basic settings, such as IP addresses, routing, and security policies. Your MikroTik VM is now ready for integration into complex network topologies.

## Verifying the MikroTik Setup

After successfully adding MikroTik to your EVE-NG environment, it’s essential to verify that the setup functions correctly. Proper verification ensures that your MikroTik router is accessible, properly configured, and ready for further network simulations. Follow these steps to confirm the setup’s integrity.

## 1. Check the VM Status

Begin by verifying that the MikroTik VM is powered on and running within EVE-NG. In the EVE-NG interface, ensure the device icon is active and shows a running status. If it’s offline, right-click and select “Start” or “Power On”.

### Outdated Drivers Are Slowing You Down

One free scan finds every outdated or missing driver and matches the right update for your exact hardware.Free scan · exact hardware match
### PC Slower Than It Used to Be?

A free scan shows the junk files, broken settings and background clutter dragging Windows down - then fixes them in one click.Free scan · Windows 10 & 11
## 2. Access Console

Connect to the MikroTik via the EVE-NG console feature. Right-click the device icon, choose “Console,” and open an SSH or telnet session if configured. If the MikroTik OS boots correctly, you should see the MikroTik RouterOS prompt (*[admin@MikroTik]*).

## 3. Verify Network Interfaces

- Run **interface print** to check available interfaces and their status. Confirm that interfaces are operational and connected to your topology’s switches or hosts.
- Ensure the assigned IP addresses are correct by executing **ip address print** . You should see the configured IPs on the relevant interfaces.

## 4. Test Connectivity

- Ping from the MikroTik CLI to other devices in your EVE-NG topology to verify internal connectivity.
- Test external connectivity by pinging public IPs or DNS servers, such as 8.8.8.8 or google.com.

## 5. Review Logs and System Status

Check system logs for errors or misconfigurations by executing **log print**. Also, verify system health and resource utilization to ensure the VM runs smoothly.

## Conclusion

Completing these verification steps confirms your MikroTik router is correctly integrated into EVE-NG. Troubleshoot any connectivity or configuration issues before proceeding with more advanced network simulations. Proper validation keeps your lab environment reliable and efficient.

## Common Troubleshooting Tips for Adding MikroTik to EVE-NG

Integrating MikroTik into EVE-NG can sometimes present challenges. Here are essential troubleshooting tips to ensure a smooth setup process.

