---
id: collect-240926-storagereview/storagereview/fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee
title: "fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee"
domain: storagereview
role: reference
task: reference
actors: ["Apple", "Google", "Intel", "Samsung"]
dates: []
keywords: ["gpu", "intel", "accelerator", "benchmarks", "compute", "cost", "decode", "energy", "gpus", "latency", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee.md
source_anchor: ""
source_lines: [1, 115]
sha256: d751296e3ce7830fd259e3683a0e9e80dabe1e9854d693e7ff11691597c3d769
---

# fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee

<!-- source: https://www.storagereview.com/fr/review/vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review -->

The Virtual Desktop Infrastructure (VDI) enterprise community is calling for more robust solutions. As enterprises strive to improve efficiency and user experience, the role of dedicated hardware accelerators has become essential. Intel's Data Center GPU Flex series occupies an important place in this space, offering tailored solutions for cloud gaming, media, VDI, and graphics acceleration within data centers.
What are Intel Flex Series GPUs?
Intel's offering is primarily based on two flagship products: the Flex 140 and Flex 170 GPUs. The Flex Series 140, a low-profile PCIe Gen4 card, integrates two GPUs, each with eight Xe cores and 6 GB of GDDR6 memory. This configuration is ideal for handling up to 12 VDI sessions and meets the needs of users with moderate graphics requirements.
For more graphics-intensive applications, the Flex Series 170 evolves with a single GPU node featuring 32 Xe cores and 16 GB of GDDR6 memory, all on a full-size PCIe card, offering substantial firepower for high-resolution tasks.
|  | Intel Flex 170 Data Center GPU | Intel Flex 140 Data Center GPU | 
|---|---|---|
| Essentials |  |  | 
| Microarchitecture | Xe HPG | Xe HPG | 
| Embedded options available | No | No | 
| Usage conditions | Server/Enterprise | Server/Enterprise | 
| Use cases | Cloud Computing | Cloud Computing | 
| Graphics processor specifications |  |  | 
| Xe-cores | 32 | 16 | 
| Render slices | 8 | 4 | 
| Ray tracing units | 32 | 16 | 
| Intel® Xe Matrix Extensions (Intel® XMX) engines | 512 | 256 | 
| Execution units | 512 | 256 | 
| Graphics Max Dynamic Clock | 2050 MHz | 1950 MHz | 
| Intel® Xe Matrix Extensions (Intel® XMX) Max Dynamic Clock | 1950 MHz | 1600 MHz | 
| TBP | 150 W | 75 W | 
| Memory characteristics |  |  | 
| Memory size | 16 GB | 12 GB | 
| Memory type | GDDR6 | GDDR6 | 
| Graphics memory interface | 256 Bits | 192 Bits | 
| Graphics memory bandwidth | 576 GB / s | 336 GB / s | 
| Supported technologies |  |  | 
| Ray tracing | Yes | Yes | 
| API support | Yes | Yes | 
| OpenVINO™ support | Yes | Yes | 
| DirectX* support | DirectX 12 Ultimate | DirectX 12 Ultimate | 
| Vulkan* support | 1.3 | 1.3 | 
| OpenGL* support | Up to 4.6 | Up to 4.6 | 
| OpenCL* support | 3 | 3 | 
| Multiformat codec engines | 2 | 4 | 
| Feature |  |  | 
| H.264 hardware encode/decode | Yes | Yes | 
| H.265 (HEVC) hardware encode/decode | Yes | Yes | 
| AV1 encode/decode | Yes | Yes | 
| VP9 bitstream and decode | Yes | Yes | 
The foundation of these graphics processing units is the Xe-core, with the Flex Series 170 featuring twice as many cores and render slices as the Flex 140. This paves a direct path to doubled ray tracing prowess, reinforcing the Flex 170's ability to handle complex rendering tasks and simulations with ease.
On the memory front, we also see the Flex Series 170 pulling ahead with a 16 GB GDDR6 allocation, outpacing the Flex 140's 12 GB. This additional memory, combined with a wider 256-bit interface, delivers graphics bandwidth peaking at 576 GB/s, compared to 336 GB/s for its counterpart.
The debate over efficiency persists regarding the GPUs' thermal design power (TDP). The Flex Series 170's 150-watt TDP indicates a focus on performance, while the Flex Series 140's 75-watt TDP emphasizes its penchant for power-efficient and energy-saving applications.
Support for cutting-edge technologies such as Ray Tracing, oneAPI, and OpenVINO is a common thread across both models, ensuring a scalable development platform. Meanwhile, DirectX 12 Ultimate compatibility supports ultra-realistic graphics, a nod to potential cross-applications in fields such as professional visualization and cloud gaming.
The magic of Intel Flex GPUs: SR-IOV
The magic of the Flex Series GPU comes from SR-IOV. If you know about it, feel free to skip to the next section; there's nothing new here. If not, or if you need a refresher, buckle up; this is cool stuff.
Single Root I/O Virtualization (SR-IOV) is a technology that enhances the manageability and efficiency of virtualized environments by allowing a single physical device, such as a network interface card (NIC) or a graphics processing unit (GPU), to appear as multiple distinct virtual devices. This is particularly useful in data centers for improving virtual machine (VM) performance and maximizing utilization of the underlying hardware resources.
SR-IOV technology relies on two fundamental concepts: physical functions (PF) and virtual functions (VF). The PF is the primary interface of the physical device and manages SR-IOV functionality, including the creation and management of VFs. These VFs are lightweight versions of the PF, equipped with the resources needed to move data but with reduced configuration capabilities. Each VF can be directly assigned to a VM, providing direct, high-performance access to the device's capabilities without the typical overhead of virtualized devices.
When SR-IOV is used with GPUs, it allows each virtual machine (VM) to directly access a portion of the GPU's resources. This direct access is facilitated through virtual functions (VFs), which are lightweight representations of the GPU that can be individually assigned to VMs.
Although VFs are controlled by the VMs to which they are assigned, the PF retains overall control, managing resources and enforcing policies at the device level. This configuration is invaluable in scenarios where performance and low latency are critical, such as in complex VDI environments, high-performance computing tasks, and large-scale web services, significantly improving the operational efficiency of virtualized systems.
This configuration allows VMs to bypass traditional hypervisor-based resource sharing methods, thereby reducing overhead and improving performance. GPU-intensive tasks, such as 3D rendering, video processing, or machine learning applications, can experience significantly lower latency, more efficient GPU resource utilization, and improved overall performance in virtualized environments.
Intel Flex GPUs offer "free" VDI acceleration
The section title is not a joke. The Intel Data Center GPU Flex series arrives on the accelerator scene with a significant advantage in VDI deployments: no licensing costs for configuring virtual GPU (vGPU) configurations. Leveraging the aforementioned Single Root I/O Virtualization (SR-IOV)-based GPU virtualization, this Intel Flex series eliminates the traditional financial barriers associated with vGPU provisioning. The absence of licensing fees reduces initial setup costs and decreases ongoing operational expenses, enabling significant long-term savings.
This cost-effective approach to vGPU administration improves the scalability of VDI server deployments. Organizations can dynamically provision and adjust vGPU resources in virtual environments without worrying about additional licensing costs. This flexibility is crucial for adapting to evolving workload demands and supporting a growing number of users, particularly in high-density environments typical of knowledge worker segments.
From a business perspective, Intel Flex Series GPUs offer substantial financial benefits. By removing licensing complexities and costs, Intel enables organizations to streamline their VDI infrastructure. This simplification accelerates deployment times and reduces the need for large budget allocations for GPU capacities, making the Intel Flex series an attractive option for enterprises seeking to optimize their VDI investments while maintaining high performance and reliability.
Hands-on: Intel Flex Series 170 with Supermicro SuperBlade
In our lab, we tested the Intel Flex Series 170 graphics card with VMware on our Supermicro SuperBlade X13 GPU Blade system. Installing the Intel dedicated graphics cards on VMware ESXi was streamlined to optimize the user experience and system performance. It was enough to download the driver via SCP to the host using a ZIP file, enable SSH access, and launch the installation. After a quick host reboot, the Flex Series 170 appeared in the hardware list and the SR-IOV options (from 0 to 31) were available for configuration.
The tests were carried out using the Supermicro SuperBlade system, essential for validating the performance of the Flex Series 170 GPU as a VDI accelerator. The SuperBlade system is designed to optimize compute density and efficiency while minimizing power consumption. It is thus an ideal platform for testing high-density use cases intended for knowledge workers, such as those offered by Intel Flex Series GPUs.
| Part | SuperMicro SuperBlade Configuration | 
|---|---|
| Processor | 1xIntel Xeon 8562Y+ | 
| Memory | 256GB DDR5 | 
| per chain | 2x Samsung M.3840 2G drive | 
| GPU | Intel Flex series 170 | 
The Supermicro SuperBlade X13 system is an ideal platform for optimizing deployment density. With the ability to integrate 10 nodes in a single 8U chassis, it is possible to host up to 320 accelerated VDIs in a compact and easy-to-manage chassis. Conversely, by opting for more cores and VRAM per Flex Series 170 card, one can deploy 80 to 160 VDIs for power users. Additionally, the ability to integrate up to 10 GPUs in this system offers great flexibility, and its ultra-fast internal network enables the implementation of innovative failover solutions. See our full test to learn more about this SuperBlade server, a true Swiss army knife of a system.
Intel Data Center GPU Flex Series 170 Performance
It is important to consider here that the tests were selected based on what would officially produce a complete result, a limitation of testing in virtual environments without particular adjustments or hacks. We selected a few SR-IOV slices and then a few non-standard slices to see what it would look like.
3D Rendering
We turned to the 3D Mark Wildlife benchmarks to show the cloud gaming and 3D rendering capabilities of the Flex Series GPU.
3DMark Wild Life offers a cross-platform benchmarking tool compatible with Windows, Android, and Apple iOS systems. This tool evaluates and compares the graphics performance of various devices, including laptops, tablets, and smartphones. Wild Life uses the Vulkan graphics API for Windows and Android devices, while it uses Metal for iOS devices. Since this test runs on integrated graphics according to different scores, it can illustrate the power of a pure graphics score from the Intel Flex Series 170.
| Test/SR-IOV Slice | 2GB | 4GB | 7GB | 14GB | 
|---|---|---|---|---|
| 3D Mark Wild Life | 29,062 | 42,466 | 49,671 | 45,908 | 
| 3D Mark Wild Life Extreme | 9,023 | 14,948 | 17,661 | 16,959 | 
LuxMark
Next up is LuxMark, an OpenCL GPU benchmarking utility. The Flex Series 170 truly proved flexible in this test, displaying impressive numbers and scaling.
| Test/SR-IOV Slice | 2GB | 4GB | 7GB | 14GB | 
|---|---|---|---|---|
| Luxmark Room | 2,961 | 4,382 | 11,002 | 11,202 | 
| Luxmark Food | N/A | 1,316 | 4,502 | 4,525 | 
PCMark 10 Express
You turn to PCMark 10, equipped with a vast suite of tests that accurately reflect the wide range of tasks encountered in the workplace. This benchmarking tool includes various performance assessments, custom test options, a battery life profile, and new storage tests, making it a comprehensive solution for evaluating the performance of modern desktop PCs.
| Test/SR-IOV Slice | 2GB | 4GB | 7GB | 14GB | 
|---|---|---|---|---|
| PCMark 10 Express overall | 5,111 | 5,146 | 5,311 | 5,218 | 
| Essentials EVER | 10,269 | 10,318 | 10,734 | 10,364 | 
| App startup score | 17,833 | 17,664 | 19,034 | 17,340 | 
| Video conferencing score | 7,798 | 7,933 | 8,095 | 7,980 | 
| Web browsing score | 7,789 | 7,839 | 8,028 | 8,046 | 
| Productivity | 6,952 | 7,004 | 7,181 | 7,180 | 
| Spreadsheet score | 6,924 | 6,953 | 7,186 | 7,184 | 
| Writing partition | 6,981 | 7,057 | 7,178 | 7,177 | 
Looking at our results here, even though they are not as spectacular, we can see clear scaling across various tasks that leverage acceleration.
To facilitate comparison, I compiled the results of all tests into a single table.
| Test/SR-IOV Slice | 2GB | 4GB | 7GB | 14GB | 
|---|---|---|---|---|
| 3D Mark Wild Life | 29,062 | 42,466 | 49,671 | 45,908 | 
| 3D Mark Wild Life Extreme | 9,023 | 14,948 | 17,661 | 16,959 | 
| Luxmark Room | 2,961 | 4,382 | 11,002 | 11,202 | 
| Luxmark Food | N/A | 1,316 | 4,502 | 4,525 | 
| PCMark 10 Express overall | 5,111 | 5,146 | 5,311 | 5,218 | 
| Essentials EVER | 10,269 | 10,318 | 10,734 | 10,364 | 
| App startup score | 17,833 | 17,664 | 19,034 | 17,340 | 
| Video conferencing score | 7,798 | 7,933 | 8,095 | 7,980 | 
| Web browsing score | 7,789 | 7,839 | 8,028 | 8,046 | 
| Productivity | 6,952 | 7,004 | 7,181 | 7,180 | 
| Spreadsheet score | 6,924 | 6,953 | 7,186 | 7,184 | 
| Writing partition | 6,981 | 7,057 | 7,178 | 7,177 | 
Closing Thoughts
Analyzing all this data and taking into account the human factor, I was constantly amazed, even impressed, by the simplicity and power of these cards. After allocating 1/32nd of the GPU to the VM and installing Intel's Windows driver, features like Remote Desktop worked better. Using Google Earth in Chrome, a flagship Intel feature, demonstrated that even with 512 MB of VRAM and a single Xe core, the VDI experience is surprisingly better. Having used other VDI products in the past, once the driver was installed, the experience was simply enjoyable.
The Intel Data Center GPU Flex series, particularly the Flex Series 170 model, has proven to be a reliable and scalable VDI solution that delivers exceptional performance while improving VM density and scalability. The combination of low CPU utilization, consistent frame rates, and high encoding performance per watt makes it an ideal choice for data centers seeking to transform their VDI infrastructure into high-performance, reliable, and scalable solutions. With the Intel Flex 170 GPU at its core, VDI administrators can confidently deliver exceptional graphical user experiences to end users while maintaining high quality standards across different display resolutions.
By combining them with the density achievable with the Supermicro SuperBlade system, you can create a super-dense and highly efficient VDI appliance. Thanks to the flexibility of the blades, the chassis could accommodate more powerful Max cards alongside Flex Series GPUs, bringing all users and service interactions closer together. We will discuss this again soon.
Although the scope of these tests does not include comparisons with most professional laptops, you can draw those lines yourself. When taking into account data compression for sending pixels over the wire, the security and ease of management provided by a VDI, as well as the licensing model (which, hopefully, will not tip the scales), the Flex series positions itself surprisingly well. I can summarize the Flex Series cards into two benefits that make enterprises and users happy: better performance and cheaper client PCs to deploy. For us, it becomes a no-brainer when we seek to improve aging or lackluster VDI setups that organizations struggle with.
Intel Data Center GPU Flex Series 170 Processor Product Page
