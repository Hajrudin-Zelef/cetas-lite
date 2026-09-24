---
id: collect-240926-nerdykings/nerdykings/gemma-4-12b-un-seul-transformer-pour-voir-entendre-et-raisonner
title: "Gemma 4 12B: A Single Transformer to See, Hear, and Reason"
domain: nerdykings
role: reference
task: reference
actors: ["China", "DeepSeek", "Google"]
dates: []
keywords: ["agent", "agents", "benchmarks", "consumer", "datacenter", "deepseek", "multimodal", "parameters", "reasoning"]
source: docs/RAG/clean_en/nerdykings/gemma-4-12b-un-seul-transformer-pour-voir-entendre-et-raisonner.md
source_anchor: ""
source_lines: [1, 57]
sha256: 780806a6f1836555cb6ee2b936f5fb8c606a04fc226059cfb881caafb14b2b0d
---

# Gemma 4 12B: A Single Transformer to See, Hear, and Reason

<!-- source: https://www.nerdykings.com/blog/gemma-4-transformer-unique-vision-audio.html -->

# Gemma 4 12B: A Single Transformer to See, Hear, and Reason

When you send an image to an artificial intelligence and ask it what it sees, one might think the model looks directly at the image, a bit like we do. **In reality, it has never worked that way.** And Google DeepMind has just proposed a radically different approach with **Gemma 4 12B**: a multimodal model that barely needs any specialized system to explain to it what it sees or hears.

## How a multimodal AI normally works

At its core, a language model remains a language model: it receives **tokens**, pieces of text transformed into numbers. If you write "there is a cat on the table," no problem, the model knows how to process that information. But if you give it the photo of a cat directly, things get stuck: an image is made of pixels, and a language model has no idea what to do with that.

The traditional solution is to place another neural network in front of it. For images, a vision transformer is typically used: its job is to analyze the image and extract comprehensible features from it. An **adapter** then transforms this information into a format compatible with the language model — a bit like a translator between a vision specialist and a brain that reasons. For audio, the same logic applies: a specialized encoder is added that processes the sound signal before transmitting the result to the main model.

This works very well. But in the end, your "multimodal" AI mostly looks like a **small team of specialized neural networks** connected to one another, rather than a single unified brain.

## Gemma 4: the end of specialized encoders

This is exactly where Gemma 4 12B shifts philosophy, proposing something much more radical: **doing away with these specialized encoders entirely**.

Let's take an image. Instead of sending it into a huge network dedicated to vision, Gemma starts by cutting it into small squares, what are called **patches** — imagine a photo cut up like a mosaic. Each piece contains the RGB values of the pixels present at that location. And this is where the interesting part comes in: these values are **projected mathematically, directly into the internal dimension used by the model**, with a system that indicates where each piece was located in the original image.

After that, that's basically it. These representations are sent *directly* into the main Transformer. There is no longer a big vision model tasked with "looking" at the image before it: it's the main model itself that must learn to understand what it receives. Concretely, instead of having a specialist analyze the image and then transmit the result to the brain, the pieces of the image are given directly to the brain and it's told: figure it out.

## Sound follows exactly the same path

For audio, the principle is identical. The sound signal is cut into very small slices of **40 milliseconds**. Each slice contains the values of the signal at that instant, and this information is also projected into the space used by the model before arriving directly in the Transformer.

The result: text arrives in the Transformer, images arrive in the Transformer, audio also arrives in the Transformer — and it's this **same network** that must gradually learn the relationships between all of it. If you show it enough images of dogs associated with the word "dog," the model must learn on its own that certain pixel configurations correspond to that concept. Same for speech: the Transformer must learn that certain structures in the sound signal correspond to certain sounds, certain words, and ultimately certain meanings.

## The boundary between perceiving and reasoning fades

This architecture is fascinating because it blurs a separation that was considered almost a given. Normally, one could roughly divide the work into three: a network that sees, another that listens, and an LLM that reasons about the information transmitted to it. With Gemma 4, a large part of that work ends up directly in the **same Transformer**. The brain becomes, in a sense, also the eyes and ears.

And inevitably, removing these specialized networks also makes it possible to remove an enormous number of parameters. In the architecture described by DeepMind, the large visual encoder — which could have represented around **550 million parameters** — is replaced by a much smaller integration module. On the audio side, a specialized encoder of several hundred million parameters is also done away with.

## Does it really hold up?

All of this is great on paper, but a specialized visual encoder remains extremely powerful when it comes to extracting complex details from an image. When faced with **very dense documents**, very similar textures, or particularly complicated scenes, this simplification can have limits.

What's impressive is that despite this much simpler architecture, **Gemma 4 12B retains solid multimodal capabilities**. And this idea shouldn't be seen only as a way to save a few hundred million parameters: what matters most is the philosophy behind it. Gemma 4 12B shows that multimodal models can be built **small enough to run locally**, on consumer hardware — 12 billion parameters, therefore well within reach of a good PC.

Here we find a trend also observed among certain Chinese labs like DeepSeek: stop simply stacking more and more parameters, and seek to make models more *efficient*.

## Why it really matters for AI agents

This is particularly interesting for **AI agents**. An agent needs to be able to read text, look at an interface, understand images or audio, and then act. If all this information can be processed directly by the same Transformer, you get a much more **unified** architecture — particularly suited to future agents that will need to run locally, without constantly depending on a datacenter.

This doesn't mean all models will abandon their specialized encoders: this kind of architecture retains huge advantages, especially for the most demanding cases requiring visual precision. But the experiment proves one thing: you can take a single brain, give it text, pixels, sound, and force it to learn to understand all of that *by itself*.

## My take

Let's be honest: on paper, "removing specialized encoders" looks like a somewhat risky simplification. But that's exactly what I like about what DeepMind is doing here. It's not just an optimization exercise to squeeze out a few benchmarks — it's a real architectural bet: can a single network really learn to see, hear *and* reason, without having the work chewed up for it? And the answer, at only 12 billion parameters, is surprisingly yes.

What interests me most isn't even Gemma 4 itself, it's the direction it's charting: multimodal models light enough to run on your machine, without depending on a cloud for every request. For AI agents capable of reading your screen and acting locally, this kind of unified architecture changes a lot. Take it with a grain of salt for the most demanding use cases — but the trend itself is clear.

### 🛠️ Tools you can try related to this article

A selection of my tested tools, relevant for going further.
