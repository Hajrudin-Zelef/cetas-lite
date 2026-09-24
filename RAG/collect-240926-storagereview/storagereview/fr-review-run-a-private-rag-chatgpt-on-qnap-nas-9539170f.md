---
id: collect-240926-storagereview/storagereview/fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f
title: "fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f"
domain: storagereview
role: reference
task: reference
actors: ["AMD", "Nvidia", "OpenAI", "TensorRT-LLM"]
dates: []
keywords: ["chatgpt", "amd", "cost", "gpu", "gpus", "memory", "nvidia", "parameters", "tensorrt"]
source: docs/RAG/clean_en/storagereview/fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f.md
source_anchor: ""
source_lines: [1, 79]
sha256: 2dcf814ddced774c52832751282f0e15c81a2d205ef96e02d9eca1de3011c9b5
---

# fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f

<!-- source: https://www.storagereview.com/fr/review/run-a-private-rag-chatgpt-on-qnap-nas -->

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
This turnkey application leverages the power of retrieval-augmented generation (RAG), combined with the efficiency of the TensorRT-optimized LLM and the high-speed capabilities of RTX acceleration. These provide contextual responses that are both fast and highly relevant. Running directly on your RTX Windows desktop or workstation, this setup ensures quick access to information and a high degree of privacy and security, as all processing is handled locally.
Implementing an LLM with RAG capabilities offers an excellent solution for professionals and advanced users who prioritize privacy, security, and personalized efficiency. Unlike public models such as ChatGPT, which process queries over the Internet, a local LLM operates entirely within the confines of your QNAP NAS.
This offline functionality ensures that all interactions remain private and secure. This allows users to customize the AI's knowledge base according to their specific needs, whether it's confidential business documents, specialized databases, or personal notes. This approach greatly improves the relevance and speed of AI responses, making it a valuable tool for those who need immediate, contextual information without compromising privacy or data security.
Also worth noting, and this may seem obvious, adding a GPU to the NAS directly simplifies the link between a company's data and the LLM. There is no need to move data to take advantage of this particular model, and the process is as simple and cost-effective as installing a mid-range GPU in the NAS. Moreover, at present, all this software is free, which greatly democratizes the potential of AI for small organizations.
Chat with RTX is still a beta program and at the time of writing, we were using version 0.2. But the ease of installing it and getting the web interface running was refreshing. Anyone who knows how to download and install an application can now get a local LLM with RAG in a few clicks.
Enabling remote access to chat with RTX via a universally accessible URL
We took our scenario to the next level and made it available to the entire office.
Step 1: Locate the configuration file
Start by navigating to the folder containing the configuration file:
- File path: C:\Users\{YourUserDir}\AppData\Local\NVIDIA\ChatWithRTX\RAG\trt-llm-rag-windows-main\ui\user_interface.py
Step 2: Update the launch code
Open the user_interface.py file and Ctrl-F for interface.launch. Locate the correct segment, which will appear by default as follows:
interface.launch(
    favicon_path=os.path.join(os.path.dirname(__file__), 'assets/nvidia_logo.png'),
    show_api=False,
    server_port=port
)
To enable network access, you need to add share=True as well:
interface.launch(
    favicon_path=os.path.join(os.path.dirname(__file__), 'assets/nvidia_logo.png'),
    show_api=False,
    share=True,
    server_port=port
)
Save the changes to the user_interface.py file. Then launch Chat with RTX via the Start menu, which will open a command prompt window and activate the interface.
Step 3: Find the public URL
The command prompt window will display both a local and public URL. To create a working public URL accessible from any device, merge the elements of both URLs. It would be best to take the public URL and add the local cookie information to the end:
- Public URL: https://62e1db9de99021560f.gradio.live
- Local URL with parameters: http://127.0.0.1:16852?cookie=4a56dd55-72a1-49c1-a6de-453fc5dba8f3&__theme=dark
Your combined URL should look like this, with the ?cookie added to the public URL:
https://62e1db9de99021560f.gradio.live?cookie=4a56dd55-72a1-49c1-a6de-453fc5dba8f3&__theme=dark
This URL allows access to Chat with RTX from any device on your network, extending its usability beyond local constraints.
Final thoughts
We have long been fans of QNAP's leadership in NAS hardware design, but QNAP customers have far more value at their disposal than they probably realize. To be honest, Virtualization Station is a great starting point, but why not take it to the next level and try GPU Passthrough? At the very least, organizations can provide a high-end GPU-powered virtual machine without having to set up a dedicated workstation. There are also the apparent benefits of a VM placed next to a huge internal storage pool with native performance levels. In this case, we had shared storage performance of nearly 10 GB/s, without worrying about a single 100 GbE connection or switch, all because the GPU-accelerated VM was inside the NAS itself.
Why not go even further to realize the benefits of AI for the organization? We have shown that adding a decent GPU to a QNAP NAS is relatively simple and inexpensive. We put an A4000 to work, and with a list price of around $1,050, that's not bad considering Virtualization Station is free and NVIDIA Chat with RTX is available for free. Being able to securely point this powerful LLM at a company's private data should provide actionable insights while making the business more agile.
Another goal to consider here is a file store for models that can be external to the QNAP system itself. This is ideal for small businesses that need a fast place to store their working data. With advanced networking capabilities, you could eventually use the NAS as a data storage location for RAG work on a larger GPU server, enabling an easily shareable data store from which to infer.
This is just one example of AI. The industry is evolving rapidly, so tools will continue to be made available. Smart businesses need to learn how to leverage AI, and this simple QNAP feature is a great way to start.
QNAP Virtualization Station
