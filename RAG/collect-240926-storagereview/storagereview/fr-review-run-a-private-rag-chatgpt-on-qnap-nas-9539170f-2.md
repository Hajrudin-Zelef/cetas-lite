---
id: collect-240926-storagereview/storagereview/fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f-2
title: "fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Nvidia", "OpenAI"]
dates: []
keywords: ["chatgpt", "amd", "gpu", "gpus", "memory", "nvidia"]
source: docs/RAG/clean_en/storagereview/fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f.md
source_anchor: ""
source_lines: [3, 41]
sha256: 0eabcc9f9313fac0e32183a00f43ffe783c89bc0c8f5579f06ba4563a6e12118
---

# fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f

QNAP is renowned for its hardware design, and in particular for its ability to integrate power, scalability, and flexibility beyond that of its competitors. We recently tested the TS-h1290FX, a 12-bay NVMe NAS equipped with an AMD EPYC 7302P processor (16 cores/32 threads), 256 GB of RAM, built-in 25 GbE connectivity, and numerous PCI ports. With such power and so many pre-installed applications, what happens if we add a graphics card and see how far we can push this NAS to run artificial intelligence applications, such as a private ChatGPT server?
NAS storage potential for AI
The QNAP TS-h1290FX has a lot to offer businesses looking to get into AI. This NAS has a unique advantage: it supports an internal GPU and offers massive storage potential. Large AI models require a significant amount of data, which must be stored and accessed efficiently. This can prove complex for storage platforms using hard drives, but the TS-h1290FX, thanks to its U.2 NVMe compatibility, meets every challenge.
When thinking of high-capacity NAS, one often imagines platforms with 3.5-inch hard drives compatible with drives up to 24 TB. That sounds enormous, but it's little compared to U.2 QLC SSDs. QNAP recently added compatibility with the Solidigm P5336 range, which reaches an incredible capacity of 61.44 TB per drive. For a 12-bay model like the TS-h1290FX, users get 737 TB of raw storage before data reduction. When it comes to compact desktop NAS, few systems can compete.
As businesses rapidly adopt AI, having a system capable of providing storage capacity for AI workloads and running models is a huge advantage. The impressive feat, however, is that this QNAP NAS can run these AI workloads while fulfilling its primary storage-sharing tasks in the SMB environment.
It should also be said that AI is not monolithic. Different AI projects require different types of storage to support them. While we are focusing here on the desktop unit, QNAP offers many other NAS systems supporting flash and high-speed networking, elements essential to meeting a more ambitious AI need than what we have covered here.
How does QNAP support GPUs?
QNAP supports GPUs in many of their NAS systems. They also offer a few applications that also support GPUs. For this article, we are primarily looking at the GPU through the prism of Virtualization Station. Virtualization Station is a hypervisor for QNAP NAS, which allows users to create a variety of virtual machines. Virtualization Station also has a set of in-depth features that support VM backups, snapshots, clones and, more importantly, GPU passthrough in the context of this article.
Inside our test unit, the QNAP TS-h1290FX is equipped with a typical server board with several PCIe slots available for expansion. QNAP also provides the necessary GPU power cables inside the chassis, so no funny business is required for cards that need more than the PCIe slot power. We found that the single-slot NVIDIA RTX A4000 fit perfectly with enough room for cooling. On this platform, a GPU with an active cooler is preferred. Your choice of GPU will be determined by the workload and by what the NAS can physically support and cool.
Configuring QNAP for AI
Configuring a virtual machine (VM) with direct GPU passthrough on a QNAP NAS requires several steps. This requires a QNAP NAS compatible with virtualization and equipped with the necessary hardware capabilities. Below is a guide explaining how to configure a QNAP NAS with direct GPU passthrough.
1. Check hardware compatibility
Make sure your QNAP NAS supports Virtualization Station, which is QNAP's virtualization application.
- Confirm that the NAS has an available PCIe slot for a GPU and that the GPU supports passthrough. Compatibility lists are often available on the QNAP website. Although the current compatibility list does not officially support the NVIDIA A4000, we had no problems with the features.
2. Install the GPU
- Turn off the NAS and unplug it from power. Open the case and insert the GPU into an available PCIe slot. Connect all necessary power cables to the GPU. Close the case, reconnect power, and turn on the NAS.
3. Update your QNAP firmware and software
Make sure your QNAP NAS is running the latest version of QTS (QNAP's operating system). We used Virtualization Station 4, which is an open beta version from QNAP, to provide better support and better performance for GPU work. Virtualization Station 4 is a self-installing package, unlike others that are installed directly via QNAP App Center.
4. Install the operating system on the VM
After installing QNAP's Virtualization Station on your NAS, you can access the management interface to deploy your virtual machine (VM). When you click "Create," a prompt window will appear allowing you to provide the VM name and select the location on the NAS where the VM will run. In most cases, you may need to make a few minor adjustments to the operating system and version information.
Next, adjust the resources and CPU compatibility type that the VM will see at the guest operating system level. In our case, we gave our VM 64 GB of memory and 8 processors. We selected the passthrough processor type for the model and changed the BIOS to UEFI.
To boot and install the operating system, you need to download and mount an ISO file as a virtual CD/DVD drive. Once the installation process is complete, enable RDP for management before moving to the next step. QNAP's virtual machine management functionality changes once GPU passthrough is enabled, and RDP greatly simplifies this process. At this point, shut down the VM.
5. Configure GPU passthrough
In Virtualization Station:
- With the existing VM powered off, edit your VM.
- In the VM settings menu, look for the Physical Devices tab. From there, select PCIe. You will see a device available for passthrough. In our case, it was the NVIDIA RTX A4000. Apply this change.
- If you need to allocate other resources to your VM, such as CPU cores, RAM, and storage, this is the time to do so.
- Power the VM back on.
6. Install GPU drivers in the VM
Once back in the VM using RDP with the GPU connected, download and install the appropriate drivers for your GPU in the VM. This step is crucial for the GPU to function properly and provide the expected performance improvements.
7. Verify GPU Passthrough functionality
After installing the drivers, verify that the GPU is recognized and working properly within the VM. You can use Device Manager under Windows or the relevant command-line tools under Linux to check the GPU status.
Troubleshooting and tips
- Compatibility: Check the QNAP and GPU manufacturer websites for specific compatibility notes or firmware updates that may affect passthrough functionality.
- Performance: Monitor your VM's performance and adjust resource allocations if necessary. Make sure your NAS has enough room for cooling, especially after adding a high-performance GPU.
- Networking and storage: Optimize network settings and storage configurations to avoid bottlenecks that could impact VM application performance.
NVIDIA Chat with RTX – Private ChatGPT
While it's easy to stop there (creating a Windows virtual machine with GPU access), we pushed further in this experiment to offer businesses a unique way to leverage AI securely, by harnessing the performance of the NVMe-based NAS. In our case, the VM was using RAID5-protected storage that delivered performance of 9.4 GB/s read and 2.1 GB/s write.
NVIDIA recently launched software called Chat with RTX. Chat with RTX revolutionizes interaction with AI by offering a personalized experience through the integration of a GPT-based large language model (LLM) and a unique local dataset. It notably allows processing documents, notes, multimedia content, YouTube videos, playlists, and much more.
