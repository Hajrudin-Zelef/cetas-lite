---
id: collect-240926-hardwarecooking/hardwarecooking/le-dlss-5-tourne-sur-une-intel-arc-140v-un-moddeur-reussit-l-improbable-mais-a-q
title: "le-dlss-5-tourne-sur-une-intel-arc-140v-un-moddeur-reussit-l-improbable-mais-a-q"
domain: hardwarecooking
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["intel", "amd", "fp8", "gpus", "nvidia"]
source: docs/RAG/clean_en/hardwarecooking/le-dlss-5-tourne-sur-une-intel-arc-140v-un-moddeur-reussit-l-improbable-mais-a-q.md
source_anchor: ""
source_lines: [1, 29]
sha256: 0d96cafd7dda3e556cd3708f501905ba85d8c95ffabbfbf39947b6d01300e88e
---

# le-dlss-5-tourne-sur-une-intel-arc-140v-un-moddeur-reussit-l-improbable-mais-a-q

<!-- source: https://www.hardwarecooking.fr/le-dlss-5-tourne-sur-une-intel-arc-140v-un-moddeur-reussit-limprobable-mais-a-quel-prix/ -->

While the neural rendering of **NVIDIA DLSS 5** is officially reserved for the most recent GeForce GPUs, a modder has managed to get this technology running on a platform that seemed very far from the manufacturer's requirements: the **Intel Arc 140V** iGPU. An impressive technical demonstration that proves Neural Rendering is not necessarily limited to NVIDIA hardware, even if the performance obtained remains, for now, very far from real-world use.

In recent years, AI-assisted upscaling and graphics enhancement technologies have profoundly transformed the world of PC gaming. After Ray Tracing and then Path Tracing, NVIDIA took a new step with DLSS 5 and its neural rendering capable of generating certain visual elements using AI models. An approach that has generated as much enthusiasm as debate, but which remains one of the most ambitious advances in the industry today.

## NVIDIA's neural rendering adapted to Intel's XMX units

The project is based on the work of an independent developer who managed to run NVIDIA's Neural Rendering on an integrated Intel Arc 140V graphics chip. To achieve this, it was necessary to circumvent NVIDIA's intended operation and adapt the calculations to the XMX acceleration units present in Intel's Xe2 architecture.

The mod exploits the official DLSS files in order to extract the neural models used by NVIDIA. The network weights, initially optimized for GeForce GPUs and stored in FP8 format, are then converted to FP16 so they can be processed by the Intel hardware. This additional step represents a significant computational load, but demonstrates that it is technically possible to run the neural rendering engine on a completely different architecture.

This experiment also confirms that certain technological building blocks of DLSS 5 are not intrinsically tied to NVIDIA's Tensor Cores. However, this does not mean they can operate efficiently on just any graphics processor.

## Performance far too low for gaming

While the demonstration is impressive on a technical level, the results obtained quickly show the limits of the exercise. In ***Tekken 7***, the Intel Arc 140V reaches only about 3 FPS at 720p and drops to just 1.5 FPS at 1080p when neural rendering is enabled.

The additional processing imposed by Neural Rendering is colossal for an integrated graphics chip in this category. In Full HD, the added computation time would exceed 400 milliseconds per frame, which explains the complete collapse of the framerate. Even by drastically reducing the resolution to 640 x 360 pixels, the platform only manages to reach about ten frames per second.

The results observed in ***Mortal Kombat 1*** point in the same direction, with only a few frames per second despite a resolution limited to 720p. In other words, the experience is completely unplayable, but that wasn't really the goal of the project.

## A proof of concept that raises questions about the future of neural rendering

Beyond performance, this experiment above all shows that neural rendering technologies could become more open than initially imagined. If a modest integrated Intel graphics chip can run, even laboriously, certain processes intended for high-end NVIDIA GPUs, that hints at interesting possibilities for future competing architectures.

This demonstration comes at a time when AMD is also reportedly preparing its own neural rendering technologies for the future RDNA 5 generation. The industry seems to be gradually converging toward a new approach where artificial intelligence is no longer limited to upscaling, but directly participates in creating the final image.

For now, running DLSS 5 on an Intel Arc 140V remains more of a reverse-engineering feat than a practical solution for gamers. But this experiment perfectly illustrates just how far the modding community is capable of going to push the limits of current hardware.
