---
id: collect-240926-frandroid/frandroid/j-ai-testa-le-mac-mini-m6-toujours-une-raussite-malgra-son-prix-1
title: "j-ai-testa-le-mac-mini-m6-toujours-une-raussite-malgra-son-prix"
domain: frandroid
role: reference
task: reference
actors: ["AMD", "Alibaba", "Apple", "Intel", "Qualcomm", "Samsung"]
dates: []
keywords: ["accelerator", "amd", "benchmark", "benchmarks", "cost", "exploit", "gpu", "inference", "latency", "memory", "nand", "panther lake"]
source: docs/RAG/clean_en/frandroid/j-ai-testa-le-mac-mini-m6-toujours-une-raussite-malgra-son-prix.md
source_anchor: ""
source_lines: [1, 87]
sha256: 25120fdb1d8afa86c4cddbeb4baac20e29fe20b00c2d6cc99c39bf0496771e9e
---

# j-ai-testa-le-mac-mini-m6-toujours-une-raussite-malgra-son-prix

<!-- source: https://www.frandroid.com/marques/apple/3255721_test-de-lapple-mac-mini-m6-la-machine-parfaite-pour-tout-le-monde -->

As it does every two years, Apple is back with a new version of the Mac mini. Every two years, it is one of our favorite computers of the year, thanks to its design, its connectivity, and above all its performance.

In 2026, Apple seems to have understood that its mini-PC (yes, PC) is used both by workers and creatives and by some curious users looking to exploit its AI capabilities. For this new model, the firm is pulling out all the stops with two different models: the one we are testing in these lines with an M6 chip and a second for the most demanding users equipped with the M5 Pro chip.

## Technical specifications

*The machine is loaned by Apple for this test.*

| Model | Apple Mac mini M6 (2026) | 
|---|---|
| Dimensions | 127 mm x 127 mm x 50 mm | 
| Processor (CPU) | M6, M5 Pro | 
| Graphics chip (GPU) | Apple GPU | 
| Random access memory (RAM) | 16, 24, 32, 48, 64 GB | 
| Internal storage | 256, 512, 1024, 2048, 4096, 8192 GB | 
| Wi-Fi standard | Wi-Fi 7 (be) | 
| Bluetooth version | 6.0 | 
| Operating system (OS) | macOS | 
| Weight | 670 grams | 
|  | Product page | 

## Design

The design has not changed one bit since the Mac mini M4. This new model keeps its ultra-compact chassis with its 12.7 cm sides and 5 cm height. The unibody enclosure made of recycled aluminum will still fit into any environment and can easily slip under or even behind a screen.

Same design also means the same position for the power button, located under the chassis at the rear. We still find it to be a design flaw, but Apple apparently wanted to keep as much minimalism as possible on the front face of its enclosure.

Note that while the Mac mini was not known for its modularity, it was still possible to replace the SSD. The process was admittedly a bit laborious, and required special storage modules, but it was possible.

The operation is no longer permitted on this Mac mini M6, as Apple has visibly soldered NAND Flash chips directly onto the computer's motherboard.

### Connectivity

The connectivity of the Mac mini M6 is still just as complete. At the rear, there are three USB-C Thunderbolt ports (40 Gb/s, 120 Gb/s on the M5 Pro model), a full-size HDMI port, an RJ-45 2.6 Gb/s port (configurable to 10 Gb/s), and finally a power port for a simple two-prong cable.

At the front, two other USB-C ports are included, but they are limited to 10 Gb/s. Finally, the 3.5 mm headphone jack completes the package.

In terms of connectivity, the Mac mini M6 moves to Wi-Fi 7, but as with the M5 versions of the MacBook that made the same transition, you have to read between the lines. Support for the 2.5, 5, and 6 GHz frequency bands is indeed included, but the channel width remains capped at 160 MHz.

This configuration therefore does not offer the optimal performance of the standard, but will still offer much lower latency. We hope that Apple will evolve the characteristics of the network chip to take advantage of the full potential of Wi-Fi 7 in the future.

## Software

After the major update that was macOS 26 Tahoe in 2025, macOS 27 Golden Gate is, so to speak, an iteration of macOS 26. You will thus benefit from an improved Liquid Glass design that is far more customizable, as well as a redesign of the Shortcuts app to boost productivity.

Above all, macOS 27 brings a brand-new Siri AI, for which we are still waiting for the French version next October. Compared with the Mac mini M4, the experience here is much more optimized, particularly in terms of stability. We no longer have much to criticize about macOS this year, even if we would like greater video game compatibility. On this point, the ball is in the developers' court.

## Performance

Like the iPhone 18 Pro, the M6 chip in the Mac mini is based on an entirely new architecture etched at 2 nm. The processor includes 12 cores in total this year, including 2 very high-performance cores, 4 performance cores, and 6 efficiency cores. Apple promises significant gains in single-core compared with the M4.

The GPU is also made up of 12 cores which, this year, each host their own neural accelerator for AI workloads. The Neural Engine indeed sees the biggest technical evolution this year with two 16-core neural engines, double that of the previous model, to run AI models locally, but also for the AI features of certain creative software.

In short, everything has changed in this Mac mini, and we put the machine through the wringer of our many benchmarks. We divided our test into three categories: synthetic benchmarks for an initial estimate of the machine's power (CPU, GPU), creative tests on Photoshop and Premiere, and finally AI benchmarks, to evaluate the Mac mini's robustness in launching and running local models.

### Benchmarks

Let's start with a Cinebench 2026 test. On this universal benchmark, the Mac mini M6 comes out with a score of 5,472 in multi-core and 802 in single-core. It thus sits between Qualcomm's X2 and X2 Extreme architectures, but still far behind the M5 Pro, unsurprisingly. Compared with other mini-PCs on the market, the Mac mini M6 manages to surpass the Panther Lake chip in the GMKtec NucBox EVO-T2S, but still remains behind the Ryzen AI Max+ 395.

Apple did not mess with us: it is the most powerful mobile processor in single-core that we have tested to date, both compared with laptops and the other mini-PCs we have been able to test so far.

Same bell curve on Geekbench 7: the M6 chip far surpasses Panther Lake, and notably the Core Ultra X9 of a Framework PC, but can't do anything against AMD's most powerful (and more imposing, it must be said) chip.

Compared to the Mac mini M4, we measured a 45% gain in multicore and 25% in single-core. These gains are directly felt in the heaviest renders and compilations.

The graphics side is not left out either, and the Mac mini M6 is simply the 2nd most powerful iGPU after the... M5 Pro.

### For creatives

To evaluate the power of this Mac mini M6 in creative applications, we rely on the PugetBench benchmark tool, which automates a series of increasingly demanding processing and rendering tasks to push the machine to its limits.

In our tests, this new Mac mini manages to beat the MacBook M5 Pro on Adobe Photoshop to take the top spot in the ranking. On Premiere Pro, it's more complicated, and the machine closes the top trio led by the MacBook Pro M5 Pro and the Samsung Galaxy Book 6 Ultra and its Panther Lake / RTX 5070 Mobile pairing.

### Artificial intelligence

Apple highlighted it in its communication: the biggest technical advancement of this Mac mini M6 lies in its brand-new Neural Engine, whose cores are this time attached to the GPU cores.

We therefore tested the machine's capabilities to run several local AI models of different types (from 3B to 30B parameters) with the Ollama tool. The goal here is above all to test the device's limits against more powerful machines and on which types of models the Mac mini M6 performs best in terms of inference.

In terms of performance, we measure the tokens produced per second, an idea of the fluidity of using the models locally. Up to 9B parameters, performance is very decent and surpasses the mini-PCs with Panther Lake chips from our other tests with measurements at 69 tok/s in 3B, 33 tok/s in 7B and 25 tok/s in 9B.

But we very quickly notice a glass ceiling. On our version tested with 24 GB of unified memory, only 17.8 GB are addressable for inference, the rest being reserved for the system. In our tests, the heavier models like Qwen3, which require 18 GB, therefore overflow and cause GPU errors. It is possible to force it by bypassing this VRAM limit to run these larger models, at the cost of some system instability. For "*science*", we were able to measure performance in this configuration:

Qwen3 30B runs at 61.6 tok/s and Qwen3-coder at 58 tok/s.

