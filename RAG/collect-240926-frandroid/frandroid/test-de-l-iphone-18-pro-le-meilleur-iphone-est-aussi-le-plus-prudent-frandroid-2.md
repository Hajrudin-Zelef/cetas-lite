---
id: collect-240926-frandroid/frandroid/test-de-l-iphone-18-pro-le-meilleur-iphone-est-aussi-le-plus-prudent-frandroid-2
title: "test-de-l-iphone-18-pro-le-meilleur-iphone-est-aussi-le-plus-prudent-frandroid"
domain: frandroid
role: reference
task: reference
actors: ["Alibaba", "Apple", "Google", "Samsung", "Xiaomi"]
dates: []
keywords: ["benchmark", "dci", "gemini", "gpu", "llama", "llama.cpp", "memory", "tokens per second"]
source: docs/RAG/clean_en/frandroid/test-de-l-iphone-18-pro-le-meilleur-iphone-est-aussi-le-plus-prudent-frandroid.md
source_anchor: ""
source_lines: [69, 114]
sha256: 348c800477291bdf20acdd72ded6865ee52137546fbaf7c6304608af664ffdbd
---

# test-de-l-iphone-18-pro-le-meilleur-iphone-est-aussi-le-plus-prudent-frandroid

Our X-Rite i1 Display Pro probe and Calman record a peak of 1,023 nits in SDR on a white test pattern, exactly the 1,000 nits promised for everyday use, and 2,740 nits in HDR, where Apple merely claims 1,600. The 3,000-nit figure isn't measurable in the lab, it only triggers in full sunlight with the brightness sensor, but at 2,740 on a test pattern, we're already at the level of the best Android screens, and a Dolby Vision video doesn't demand more.

For accuracy, in SDR, on a ColorChecker test pattern, the average Delta E 2000 comes in at 2.48, below the threshold of 3 at which the eye can distinguish a difference. In HDR, it drops to 1.25, which is excellent, and rare: most smartphone screens fall short in HDR, the iPhone does the opposite.

The only reservation is therefore the white point, measured at 6,710 K, a little above the 6,500 K reference. It's a slightly cool, somewhat bluish white, visible on grays more than on colors, and that True Tone corrects in daily use since we had disabled it for the measurement. Those who retouch photos will enable Reference mode, but chances are you'll never see the difference.

In terms of coverage, the panel covers 106.5% of BT.709, 99.5% of DCI-P3 and 72% of BT.2020. This means the screen is capable of displaying almost everything the widest Ultra HD content contains, which very few panels can do.

A useful clarification for understanding why the iPhone sometimes looks less "punchy" than a Galaxy in Vivid mode: iOS manages colors strictly, sRGB content stays within sRGB, and this gamut headroom is only called upon by content that declares it. The iPhone offers no option to change display modes, unlike Android. In short, Apple has the widest screen on the market and refuses to use it to flatter the eye. It's a choice, and it's certainly the right one.

Finally, we would have liked to see Apple adopt the privacy screen that Samsung has on the Galaxy S26 Ultra and that Xiaomi will introduce on its Xiaomi 18 Pro line. In this case, I think it's really one of the features that would be of great interest in Apple's privacy strategy. We might see it arrive on a future iPhone generation, even though, for now, no leak mentions the subject.

## Performance and heat: Apple has closed the case

The A20 Pro is the first iPhone chip etched in 2 nm, with six CPU cores, a GPU going from six to seven cores, two 16-core Neural Engines, 50% more memory bandwidth and still 12 GB of RAM.

The A20 Pro chip is the most powerful mobile chip we have tested to date.

These are numbers you don't feel when opening Safari, any recent iPhone is instantaneous, but you feel them when exporting a video, gaming for twenty minutes, or running a language model locally, where the gain approaches 60% more tokens per second. The 18 Pro is, to date, the best smartphone for on-device AI, which has a particular flavor when Siri AI isn't available in France.

The real subject is heat, and here Apple has finished the job it started last year. The vapor chamber has tripled in surface area and extends all the way to the camera block, the chip has been moved to the top of the motherboard and bonded to the chamber, the memory, until now stacked on top of it, has been relocated to the side, and a nano-etched copper plate isolates the screen. In extended benchmarking, the 18 Pro heats up as much as a 17 Pro, the heat concentrates under the camera block instead of rising into the screen, and it holds out longer before throttling. In daily use, intensive photography, navigation, video, it heats up less and cools down faster. And above all, it no longer becomes scorching hot while charging, which, as we will see, changes everything. Three years ago, heat was the iPhone's primary flaw. It no longer is one.

## AI: it runs locally

Apple has doubled the Neural Engine of the A20 Pro, going from 16 to 32 cores, and speaks of AI power doubled. This is the kind of promise that can be verified, provided you don't stop at a benchmark score. We therefore tested at three levels: a third-party language model in an app that anyone can install, the same model in Google's engine to compare with Android on equal footing, and the native Apple Intelligence features that people actually use.

Three levels, three different answers. I had the iPhone Air on hand, so it will serve as the reference: its A19 Pro shares the CPU and Neural Engine with the iPhone 17 Pro, so it plays the role of the previous generation.

So, before anyone jumps on me, I know its chassis isn't that of a Pro, without a vapor chamber and with a thickness that forgives nothing when it comes to heat. That's why all the tests in this section are short bursts, launched cold, and the Air was left to rest between each run: what we're measuring here is the chip itself, rather than the smartphone's ability to cool it. Sustained performance is the subject of the previous performance section, where the Air instead plays the role of the worst thermal case, and that's where you need to go to know what the vapor chamber changes.

First level, PocketPal, a llama.cpp app, with Qwen3 4B quantized in Q4_K_M and a prompt of over 2,000 tokens.

The iPhone 18 Pro reads the prompt in 4.4 seconds and generates at 32.4 tokens per second, the Air takes 7.1 seconds to read and generates at 16.6 tokens per second. Reading is 1.6 times faster, writing twice as fast. The ratio is instructive, because it's inverted relative to Apple's messaging: generating a token involves re-reading the model's 2.5 GB of weights from memory, and it's bandwidth that limits it.

The 50% extra bandwidth of the new memory architecture shows up here more than anywhere else. Besides, llama.cpp runs on the Metal GPU and doesn't use the Neural Engine. This test therefore sets aside the 32 cores and measures what a developer who has done nothing special for the iPhone gets, and that's already a lot.

Against Android, the comparison becomes a story of ecosystem before being a story of silicon. In the same PocketPal app, the Xiaomi 17 Ultra and its Snapdragon 8 Elite Gen 5 cap out at 15 tokens per second, the Pixel 11 Pro Fold at 2.5: on Android, the app runs on CPU, and enabling the Xiaomi's Hexagon NPU divides the speed by three for lack of a compatible format. For an honest measurement, we went through Google AI Edge Gallery, Google's app, with Google's model, Gemma 4 E4B, on GPU everywhere: the 18 Pro decodes between 25 and 34 tokens per second depending on the run, the Air and the Xiaomi run at 19, the Pixel at 7.5. The Tensor G6 is excellent with Gemini Nano in Google apps, and average as soon as you leave the Google ecosystem, the Air and the Xiaomi are evenly matched, the 18 Pro is alone out front.

And on iPhone, a third-party app gets hardware acceleration without configuring anything. On Android, the same model can go six times faster or slower depending on the app, the setting (CPU, GPU, etc.) and the smartphone.

Third level, the native features, and there the narrative flips. We took the same photo, taken with the Air, and subjected it to the same edits on both iPhones: erasing a real estate sign on a balcony with Clean Up, removing the drain grate, straightening the perspective with the new Reframe, expanding the frame.

The result is identical down to the pixel, which is logical since it's the same Apple Intelligence model on both sides, and the 18 Pro is a bit faster, on the order of a few seconds on a reframe, nothing you'd notice without a stopwatch.

It's explained quite simply, in fact: Apple calibrates its features for the least powerful iPhone entitled to them. Incidentally, Reframe leaves empty corners, which the smartphone fills in by inventing. In our photo, the building on the left gained a window that doesn't exist. It's well aligned, well lit, and totally fake.

