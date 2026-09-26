---
id: collect-240926-storagereview/storagereview/fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee-2
title: "fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee"
domain: storagereview
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["gpu", "intel", "accelerator", "cost", "decode", "energy", "gpus", "latency", "memory"]
source: docs/RAG/clean_en/storagereview/fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee.md
source_anchor: ""
source_lines: [3, 57]
sha256: 1909efbfd712fa75b4c8a4400d050acafde967d03b465a43efc8fc70cfd00d71
---

# fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee

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
