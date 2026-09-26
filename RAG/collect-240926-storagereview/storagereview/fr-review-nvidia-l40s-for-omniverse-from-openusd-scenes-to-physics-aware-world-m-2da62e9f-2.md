---
id: collect-240926-storagereview/storagereview/fr-review-nvidia-l40s-for-omniverse-from-openusd-scenes-to-physics-aware-world-m-2da62e9f-2
title: "fr-review-nvidia-l40s-for-omniverse-from-openusd-scenes-to-physics-aware-world-m-2da62e9f"
domain: storagereview
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["nvidia", "compute", "cost", "fp8", "gpu", "gpus", "inference", "memory", "multimodal", "nvlink", "throughput", "training"]
source: docs/RAG/clean_en/storagereview/fr-review-nvidia-l40s-for-omniverse-from-openusd-scenes-to-physics-aware-world-m-2da62e9f.md
source_anchor: ""
source_lines: [3, 45]
sha256: 4b453baa05907ed8389670c87fea3562f1e1af7bacb805fda618d0b86290c035
---

# fr-review-nvidia-l40s-for-omniverse-from-openusd-scenes-to-physics-aware-world-m-2da62e9f

The L40S fills a crucial gap in the data center GPU ecosystem. While compute-focused AI training GPUs like the Nvidia H100 prioritize raw performance without graphics acceleration, traditional professional visualization cards typically lack the AI compute capabilities required by modern inference workloads and new AI-driven graphics applications. This positioning makes it particularly valuable for synthetic data generation, multimodal AI development, and Omniverse applications, where both compute and graphics performance are essential.
NVIDIA L40S Specifications
| Specifications | L40 | L40S | PCIe H100 80G | 
| GPU Architecture | Ada Lovelace | Ada Lovelace | Hopper | 
| GPU Matrix | AD102 | AD102 | GH100 | 
| CUDA Cores | 18,176 | 18,176 | 14,592 | 
| Tensor Cores | 568 (4th generation) | 568 (4th generation) | 456 | 
| RT Cores | 142 (3rd generation) | 142 (3rd generation) | - | 
| GPU Memory | 48 GB GDDR6 with ECC | 48 GB GDDR6 with ECC | 80 GB HBM2e | 
| Memory Bandwidth | 864 GB / s | 864 GB / s | 2 TB / s | 
| Memory Interface | 384-bits | 384-bits | 5120-bits | 
| Maximum Power Consumption | 300W | 350W | 350W | 
| Form Factor | 4.4″ H x 10.5″ L : Dual Slot | 4.4″ H x 10.5″ L : Dual Slot | 4.4″ H x 10.5″ L : Dual Slot | 
| Thermal Solution | Passive | Passive | Passive | 
| Display Connectors | 4x DisplayPort 1.4a | 4x DisplayPort 1.4a | - | 
| PCIe Interface | Gen4 x16 | Gen4 x16 | Gen5 x16 | 
| Power Connector | 16-pin | 16-pin | 16-pin | 
| vGPU Support | Yes | Yes | No | 
| Multi-Instance GPU (MiG) | No | No | Yes | 
| NVLink Support | No | No | No | 
Performance Characteristics
| Metric | L40 Performance | L40S Performance | H100 Performance | 
| FP32 Performance | 90.5 TFLOPS | 91.6 TFLOPS | 51.2 TFLOPS | 
| TF32 Tensor Core | 362.1 TFLOPS | 366 TFLOPS | 756 TFLOPS | 
| FP16 Tensor Core | 724 TFLOPS | 733 TFLOPS | 1513 TFLOPS | 
| FP8 Tensor Core | 1,448 TFLOPS | 1,466 TFLOPS | 3026 TFLOPS | 
| Peak INT8 Tensor TOPS | 1,448 TFLOPS | 1,466 TFLOPS | 3026 TOPS | 
| Base RT Performance | 209 TFLOPS | 212 TFLOPS | - | 
(Performance figures are with sparsity)
NVIDIA L40S vs H100
A detailed examination of the specifications reveals the distinct design philosophies of the NVIDIA L40S and H100. The L40S is built on the Ada Lovelace architecture, using the same AD102 matrix as high-end NVIDIA workstation graphics cards. This heritage gives it an impressive 18,176 CUDA cores, delivering excellent FP32 single-precision performance, the cornerstone of traditional graphics rendering and scientific computing. In contrast, the Hopper architecture and GH100 matrix of the H100 are primarily designed for AI and high-performance computing workloads.
The most striking differentiator lies in the core configuration. The L40S includes 142 third-generation RT cores and 568 fourth-generation Tensor cores. RT cores are specialized components for ray tracing acceleration, a feature entirely absent from the H100, giving the L40S a unique photorealistic rendering capability. Although the H100 has fewer Tensor cores, they are faster and more advanced, optimized for new AI data formats like FP8, giving it a considerable lead in raw AI performance.
This trade-off is also evident in the memory subsystems. The H100 uses 80 GB of expensive HBM2e memory, offering an impressive bandwidth of 2 TB/s. This is essential for feeding the SMs during training and inference of large-scale AI models. The L40S uses more conventional 48 GB GDDR6 memory, offering 864 GB/s of bandwidth. Although less than half that of the H100, this capacity remains substantial, perfectly suited for loading large 3D scenes, high-resolution textures, and large AI models for inference.
Finally, the feature set describes their respective roles. The L40S includes four DisplayPort 1.4a outputs and robust vGPU support, making it ideal for virtualized workstations, render farms, and cloud gaming deployments. The H100, lacking display outputs and vGPU capabilities, incorporates Multi-Instance GPU (MiG) technology, which allows it to be partitioned into multiple smaller, isolated GPU instances to simultaneously handle multiple compute-intensive workloads. The shared dual-slot passive thermal design and 350 W power envelope of the L40S and PCIe H100 provide deployment flexibility across a wide range of standard servers.
When comparing the L40S to the H100, NVIDIA's flagship product, the differences between their respective roles become evident. The H100 is the undisputed leader in raw AI training performance for this generation of GPUs, but this comparison only reveals part of the story. Interestingly, the L40 is also part of the NVIDIA lineup, with a lower TDP of 300 W compared to the L40S's 350 W, offering similar capabilities for reduced power consumption.
The H100 we are examining here is the original PCIe version with 80 GB of HBM2e memory. NVIDIA has since expanded its lineup with variants like the H100 NVL, replacing the original 80 GB PCIe model. It offers 94 GB of memory and a higher TDP of 400 W, while incorporating NVLink support for dual-GPU configurations. The H100 family also extends to 8-GPU SXM configurations and the new H200, which shares the same GPU chip but offers improved performance in both PCIe and SXM formats.
The distinction between the L40S and H100 models highlights a strategic divergence in data center GPU design. The H100 is a purely compute-focused card, optimized exclusively for AI and high-performance computing workloads. The goal is large-scale AI training, where maximum compute throughput is the primary objective. Every aspect of its design, from the enormous HBM2e memory bandwidth to the specialized Tensor cores, is engineered for the most demanding neural network training scenarios.
In contrast, the L40S is a universal GPU designed for versatility, targeting AI inference, graphics-intensive workloads, and NVIDIA vGPU deployments. While it is a high-performance and cost-effective solution for AI inference, its true strength lies in its support for a wide range of graphics-intensive applications, for which the H100 is simply not designed.
Graphics-Intensive Workloads and Professional Applications
The L40S excels in many demanding graphics domains, requiring both compute power and advanced rendering capabilities. In 3D rendering and animation, studios rely on the L40S for rendering complex scenes, where its RT cores accelerate ray tracing calculations, impossible with GPUs dedicated solely to compute. Film and television production companies use these capabilities for real-time previews, allowing directors and cinematographers to visualize photorealistic renders of CGI scenes during filming, significantly reducing post-production timelines and costs.
Architecture and engineering firms use the L40S for architectural visualization and real-time product design. Clients can thus explore photorealistic renders of buildings or examine detailed prototypes before construction begins. The card's professional graphics capabilities also make it ideal for CAD workstations, where engineers need both compute power for complex simulations and graphics acceleration for smooth, high-fidelity display performance.
In the media and entertainment sector, the L40S powers render farms that process final-quality animation frames. Meanwhile, its vGPU capabilities enable cloud-based creative workflows, allowing artists to remotely access powerful graphics workstations. Video editing and post-production studios use the L40S for real-time effects processing, color grading, and compositing, which require both graphics acceleration and significant compute resources.
The virtual desktop infrastructure (VDI) market represents another key application area. The L40S's vGPU technology allows multiple users to share GPU resources for graphically accelerated virtual desktops. This makes professional graphics capabilities accessible in the enterprise without having to dedicate a GPU to each user.
