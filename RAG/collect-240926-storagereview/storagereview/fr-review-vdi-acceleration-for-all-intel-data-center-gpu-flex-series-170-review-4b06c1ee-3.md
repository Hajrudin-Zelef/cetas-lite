---
id: collect-240926-storagereview/storagereview/fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee-3
title: "fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee"
domain: storagereview
role: reference
task: reference
actors: ["Apple", "Google", "Intel", "Samsung"]
dates: []
keywords: ["gpu", "intel", "accelerator", "benchmarks", "compute", "gpus", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee.md
source_anchor: ""
source_lines: [58, 115]
sha256: 47a1e825bb25837dd93384bbeda4de9f1b0ff84954894b837ecb5cdfb2768219
---

# fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee

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
