---
id: collect-240926-frandroid/frandroid/test-de-l-iphone-18-pro-le-meilleur-iphone-est-aussi-le-plus-prudent-frandroid-3
title: "test-de-l-iphone-18-pro-le-meilleur-iphone-est-aussi-le-plus-prudent-frandroid"
domain: frandroid
role: reference
task: reference
actors: ["Apple", "DeepSeek", "EU"]
dates: []
keywords: ["agent", "deepseek", "gpu", "latency", "memory", "parameters", "tokens per second", "transcription", "voice"]
source: docs/RAG/clean_en/frandroid/test-de-l-iphone-18-pro-le-meilleur-iphone-est-aussi-le-plus-prudent-frandroid.md
source_anchor: ""
source_lines: [115, 167]
sha256: b93eee8fc95a07032bb40be4ec5cf6ad0eb74137d68b41e9fa12124a5d73d2b6
---

# test-de-l-iphone-18-pro-le-meilleur-iphone-est-aussi-le-plus-prudent-frandroid

We continued the tests, and we won't deliver all of them here. But we wondered what the intelligence of a 4-billion-parameter model, the kind that fits in a smartphone, looks like? We submitted to Gemma 4 E4B, on all four devices, a small physics problem based on our own charging measurements, six values to calculate (the charging test is further down). The four smartphones gave different answers, with the same model and the same prompt, because GPU arithmetic doesn't accumulate in the same order.

None exceeded 4 out of 6, all succeeded at the yield, three out of four failed a subtraction. And when, by mistake, we left the answer key in the prompt, the Pixel calculated 18.25 W, saw that the expected answer was 18.2, and wrote that one should "stick to this value."

Like all the others, the chip in the 18 Pro is faster, but it isn't smarter. That's even the common thread of this section: the closer you get to the uses Apple intended, the less the chip matters; the further you move away, the more it crushes everything. The 32-core Neural Engine is an investment for tomorrow's applications, and for now, we're still far from exploiting it fully.

In short, with its A20 Pro chip, the iPhone 18 Pro remains the most gifted smartphone for on-device AI, even though Siri AI isn't available in France. That's probably where the shoe pinches.

### What's all this power for, then?

The question needs to be asked frankly, because I'm asking it myself: who, in real life, is going to run a language model in a third-party app on their smartphone? Almost no one.

Our tokens per second are the only way to measure the Neural Engine, but in everyday use, it mostly works elsewhere. What uses it is everything else, and everything else is invisible. Every photo passes through neural networks before even appearing on screen, for exposure fusion, denoising, subject separation, white balance. The keyboard uses it for correction, Photos for search, Voice Memos for transcription, the Camera for recognizing text or an object in the viewfinder, Face ID to unlock your phone, and Siri for pretty much everything it still does here.

These functions run locally, without a network, hundreds of times a day, and it's precisely for them that Apple doubled its Neural Engine. The problem is that none of them can be measured. You don't time the denoising of a photo, you don't compare keyboard correction between two chips, you don't know whether an iPhone 18 Pro finds "the blue house" in Photos faster than an Air because the search is done before you type. The gain probably exists, spread out in milliseconds across a thousand gestures, and it's impossible to demonstrate. That's the limit of this test, and of all tests: we measure power where it's visible, and it mostly serves where it isn't. Apple knows this, and that's probably why its marketing talks about "2x" rather than showing you where.

And for tomorrow, even if it isn't the subject of this test, I'm thinking of an operating system that no longer launches apps but builds them, a local agent that understands the need and constructs the tool, without sending anything to anyone. And there's nothing science-fiction about that, since it's what frontier models already do in the cloud.

iPhone 18 Pro at €99 + 250GB plan at €29.99/month for 12 months. Enjoy €50 off instantly with code NEW50 and up to €150 trade-in bonus for your old phone. Free delivery.

What's blocking is fitting them into a pocket, and there, frankly, physics decides, whatever ingenuity you throw at it. A frontier-class model has hundreds of billions of parameters; even quantized, DeepSeek in its 671-billion version requires a Mac Studio with 512 GB of unified memory and more than 800 GB/s of bandwidth to run.

An iPhone 18 Pro has 12 GB and a fraction of that bandwidth, and we've seen what that bandwidth does to tokens per second. Context is the other wall: a coding agent that loads an entire repository far exceeds the memory cache a 16 GB card can hold, so for a smartphone chip, let's not even talk about it. What remains, then, is what we tested, the 4-billion-parameter models, and they're impressive for their size, but their size is the point.

My own experience is more modest and more telling: four high-school-level physics calculations, and Gemma 4 E4B fails two out of six (mental arithmetic is a weakness of all language models, but still…).

The good news is that the frontier is coming down: a 3-to-14-billion model today does what a 70-billion did twelve to eighteen months ago. The dream will therefore come, from below, and the iPhone that will have the memory and bandwidth to welcome it hasn't been released yet.

## iOS: the good and the missing

Let's start with what doesn't make waves, because it's rare: iOS 27 on the iPhone 18 Pro is a finished system. No blocking bug for a week, no stutter in animations, no lazy screen wake, no Wi-Fi dropping out, none of those little miseries that ordinarily accompany a new iPhone until version .1 or .2.

- iPhone updated? Here are the iOS 27 settings to change right away
- iOS 27: 15 small hidden settings that save time every day

120 Hz is used everywhere, transitions are clean, the Camera app has integrated its Pro controls without weighing down the interface, and the Apple Intelligence features available here—Clean Up, Reframe, Writing Tools, Image Playground—run without a hitch. After two years of redesigns, Apple has delivered an iOS that asks for nothing more than to be used. That deserved to be said before what follows.

Because what follows is absence. Siri AI, the real software novelty of the year, the one that took up a good half-hour of WWDC as well as announcements at the September keynote, is not available in the European Union at the launch of iOS 27, and Apple has no date to give: according to Craig Federighi, Apple's software chief, no timeline exists for the iPhone and iPad in Europe. Apple cites the DMA and claims that regulators have not accepted any of its proposals, including a mechanism called Trusted System Agent designed as a technical intermediary between Siri and competing assistants, with a gradual rollout over eighteen months.

These same users will nevertheless have Siri AI on Mac and on Vision Pro, and nothing on the smartphone they just paid 1,479 euros for. And what's missing is not a detail.

In testing, what really convinces is less the conversation than the screen reading: taking a number from a document, creating a reminder from a product page, recognizing a filmed object, gestures that replace dozens of copy-pastes per day.

A French iPhone 18 Pro in France is an iPhone 18 Pro amputated of what Apple presented as its raison d'Ãªtre in 2026. You can work around this block with an American or Canadian Apple account, but that's a hack.

The second flaw is more down-to-earth and, in use, even more annoying: voice dictation has become a horror. This is not a regression of the iPhone 18 Pro, it's a regression of iOS 27 in French: random punctuation, words swallowed at the end of sentences, proper nouns butchered, corrections that come back around and replace a right word with a wrong one, and a latency that makes text arrive in bursts rather than following the voice.

On a device that packs a 32-core Neural Engine and a local transcription model capable of transcribing ten minutes of memo in a few seconds in Voice Memos, it's incomprehensible: the building block exists, it's not connected to the keyboard. It's crazy to think that basic artificial intelligence functions still work so poorly. It's like search in Settings, which works one time out of two.

Our advice is simple and a bit shameful for Apple: go through a third-party app. Keyboards and voice recorders based on Whisper or an equivalent model do better, in French, with Breton names, without a connection, and they insert text into any application.

## Photo: the diaphragm is almost useless, the aperture is everything

