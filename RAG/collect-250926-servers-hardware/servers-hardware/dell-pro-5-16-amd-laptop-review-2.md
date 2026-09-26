---
id: collect-250926-servers-hardware/servers-hardware/dell-pro-5-16-amd-laptop-review-2
title: "dell-pro-5-16-amd-laptop-review"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "intel", "latency", "lpddr5x", "memory", "power delivery"]
source: docs/RAG/clean4/dell-pro-5-16-amd-laptop-review.md
source_anchor: ""
source_lines: [54, 84]
sha256: 8379ea84c698c39a9b51a69b22ae62db30a12f7f66e4a5888a37c6940f4ef60b
---

# dell-pro-5-16-amd-laptop-review

Shifting over to the left side of the laptop, we find the system’s faster ports. This includes a pair of 40Gbps USB-C ports driven by the Gorgon Point SoC. Notably, the ports/laptop have even been been Thunderbolt 4 certified, which is a step a lot of AMD laptop vendors do not take, and helps to elevate this laptop from its competition. Though for Dell, the matter may be more one of pragmatism: TB4 certification keeps the overall laptop capabilities at parity with the Intel-based Dell Pro 5 models.

Moving on, we have another 5Gbps USB-A port hanging off of the system’s SoC, and then finally an HDMI port. External monitor users will want to make note that even though it is an HDMI 2.1 port, Dell is only supporting TDMS signaling here, so the port can only drive a display up to 4K@60Hz without DSC. Which is sufficient for TVs, but a bit shy of what the newest monitors can do. That means the USB-C ports with DisplayPort Alt Mode will be the prime choice for high-end external monitors. Especially if you can get one with USB power delivery, as the Dell Pro 5 16 charges exclusively via USB-C.

A quick look at the top of the laptop reveals the familiar Dell logo.

Finally, along the bottom we find the intakes for the laptop’s ventilation system, along with Dell’s usual collection of stickers, serial numbers, and QR codes. It is also the only place you will find labeling that identifies the specific model, or that it is a Pro laptop in the first place. Dell uses rubber strips along both the front and back sides of the laptops to elevate them for airflow, and these are surprisingly tall strips as well, with the rear strip adding a bit over 4mm to the final height of the 19mm laptop.

Now, let us go ahead and take a look at the internals of the Dell Pro 5 16 (AMD) laptop.

LatencyMon results would have been interesting to check whether this notebook is suitable to be used as a DAW for audio processing.

Please add LatencyMon results measuring “interrupt to user process latency”.

Results of a 5-10 minute run with no other applications running.

LatencyMon issues a typical DAW load itself.

The internal layout is a bit disappointing. They clearly used a battery meant for a 14″ device, while they could have very easily fitted in a larger one to improve the runtime and while doing so they also opted for a single 2230 m.2 drive, which is just madness. Also, using SODIMMs with a CPU that could have greatly benefitted from an LPCAMM module or soldered memory is just sad.

It feels like Dell deliberately wants their AMD machines to be inferior to their Intel ones, for some reason, and this isn’t the only example of that. I’m genuinely surprised they included a USB4 controller in the first place.

“Also, using SODIMMs with a CPU that could have greatly benefitted from an LPCAMM module or soldered memory is just sad.”

To note: AMD’s Strix/Gorgon Point SoCs do not support LPCAMM2, unfortunately. They are not quite a drop-in replacement for soldered LPDDR5X, and require some design considerations on the host side. So if removable memory is desired, Dell’s only option here was DDR5 SO-DIMM.

The CPU in this laptop supports two USB4 ports natively. AMDs mobile only silicon has done so since the 6000 series. I’d be very surprised if Dell spent money on an external controller instead of using the native function in the CPU.

@Anonymous

Oh goodness, you are right! This is definitely native. I have gone back and corrected the article. Thank you for pointing that out.
