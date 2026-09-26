---
id: vague2-nerdykings/nerdykings/gemma-4-transformer-unique-vision-audio-1
title: "Gemma 4 12B : Un Seul Transformer Pour Voir, Entendre Et Raisonner"
domain: nerdykings
role: reference
task: article
actors: ["China", "DeepSeek", "Google"]
dates: ["2026-09-23"]
keywords: ["agent", "agents", "benchmarks", "consumer", "datacenter", "deepseek", "multimodal", "parameters", "reasoning"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/gemma-4-transformer-unique-vision-audio.md
source_anchor: ""
source_lines: [1, 57]
sha256: b53ed5f27a577aaa26e3cc951d4fbf61763c7d19f63be459b8c2a7e5168cef88
---

# Gemma 4 12B : Un Seul Transformer Pour Voir, Entendre Et Raisonner

## Metadata

- **Source** : https://www.nerdykings.com/blog/gemma-4-transformer-unique-vision-audio.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

When you send an image to an AI and ask what it sees, one might think the model looks directly at the image like a human. In reality it has never worked that way — and Google DeepMind proposes a radically different approach with **Gemma 4 12B**, a multimodal model that almost no longer needs a specialized system to explain what it sees or hears.

How multimodal AI normally works: a language model receives tokens (text chunks turned into numbers). Give it the text "there is a cat on the table" and it's fine; give it the photo directly and it struggles, because an image is pixels and a language model doesn't know what to do with them. The traditional solution is to place another neural network in front. For images, a **vision transformer** analyzes the image and extracts comprehensible features; an adapter then converts this into a format compatible with the language model — like a translator between a vision specialist and a reasoning brain. For audio, the same logic: a specialized encoder processes the sound signal before sending the result to the main model. This works very well, but the "multimodal" AI ends up looking like a small team of specialized neural networks wired together rather than one unified brain.

Gemma 4: the end of specialized encoders. Gemma 4 12B changes philosophy radically: it removes these specialized encoders. Taking an image, instead of sending it into a huge vision-dedicated network, Gemma first cuts it into small squares called **patches** — like a photo cut into a mosaic. Each piece contains the RGB values of the pixels at that location. Then the interesting part: these values are projected mathematically directly into the internal dimension used by the model, with a system indicating where each piece was in the original image. After that, that's almost all. These representations are sent directly into the main Transformer. There's no longer a big vision model "looking" at the image before it: the main model itself must learn to understand what it receives. Instead of a specialist analyzing the image then passing the result to the brain, the image pieces are given directly to the brain with the instruction: figure it out.

Sound follows exactly the same path. The audio signal is cut into tiny slices of **40 milliseconds**. Each slice contains the signal values at that instant, and this information is also projected into the model's space before arriving directly in the Transformer. Result: text arrives in the Transformer, image arrives in the Transformer, audio arrives in the Transformer — and it's the same network that must progressively learn the relations between all of them. If shown enough images of a dog associated with the word "dog," the model must learn by itself that certain pixel configurations correspond to that concept. Same for speech: the Transformer must learn that certain structures in the sound signal correspond to certain sounds, words, and finally meanings.

The boundary between perceiving and reasoning fades. This architecture blurs a separation almost taken for granted. Normally the work is roughly split into three: a network that sees, another that listens, and an LLM that reasons about the transmitted information. With Gemma 4, much of this work is directly in the same Transformer. The brain becomes, so to speak, also the eyes and ears. And removing these specialized networks also removes a huge number of parameters. In DeepMind's described architecture, the big visual encoder — which could have represented about **550 million parameters** — is replaced by a much smaller integration module. On the audio side, a specialized encoder of several hundred million parameters is also dropped.

Does it really hold up? On paper this is great, but a specialized visual encoder remains extremely powerful for extracting complex details from an image. Facing very dense documents, very similar textures, or particularly complicated scenes, this simplification may have limits. What's impressive is that despite this much simpler architecture, Gemma 4 12B retains solid multimodal capabilities. And this idea shouldn't be seen only as a way to save a few hundred million parameters: the philosophy matters. Gemma 4 12B shows multimodal models can be built small enough to run locally on consumer hardware — 12 billion parameters, largely within reach of a good PC. This reflects a trend also seen in some Chinese labs like DeepSeek: stop simply stacking more parameters and seek more efficient models.

Why it matters for AI agents: an agent must read text, look at an interface, understand images or audio, then act. If all this information can be processed directly by the same Transformer, the architecture becomes much more unified — particularly suited to future agents that must run locally without constantly depending on a datacenter. This doesn't mean all models will abandon specialized encoders: such architecture keeps huge advantages, especially for the most visually demanding cases. But the experiment proves one thing: a single brain can be given text, pixels, and sound, and forced to learn to understand all of it by itself.

The author's view: on paper, "removing specialized encoders" looks like a somewhat risky simplification. But that's precisely what's appealing about DeepMind's work. It's not just an optimization exercise to gain a few benchmarks — it's a real architectural bet: can a single network really learn to see, hear, and reason without having the work pre-chewed? And the answer, at only 12 billion parameters, is surprisingly yes. What interests the author most isn't Gemma 4 itself but the direction it draws: multimodal models light enough to run on your machine without depending on a cloud for each request. For AI agents able to read your screen and act locally, this kind of unified architecture changes a lot. To be taken with caution for the most demanding use cases — but the trend is clear.

## Key points

- **Gemma 4 12B** (Google DeepMind) removes specialized vision/audio encoders from a multimodal model.
- Traditional approach: vision transformer + adapter (and audio encoder) translate modalities before the LLM.
- Gemma's approach: cut images into **patches** (RGB values) projected directly into the model's internal dimension with position info.
- Audio is sliced into **40 ms** segments and projected the same way into the Transformer.
- All modalities (text, image, audio) enter the same Transformer, which learns cross-modal relations itself.
- The ~**550M-parameter** visual encoder (plus a several-hundred-million audio encoder) is replaced by much smaller integration modules.
- The boundary between perceiving and reasoning fades; the "brain" also becomes eyes and ears.
- Runs locally: 12B parameters, within reach of a good PC; suited to local AI agents.
- Limits possible on dense documents, similar textures, and complex scenes where specialized encoders excel.
- Reflects the industry trend (also at DeepSeek) toward efficiency over raw scale.

## Technical data / figures

| Item | Value |
|---|---|
| Model | Gemma 4 12B |
| Developer | Google DeepMind |
| Parameters | 12 billion |
| Removed visual encoder | ~550 million parameters |
| Removed audio encoder | several hundred million parameters |
| Image handling | Patch mosaic with RGB values, position-tagged |
| Audio handling | 40-millisecond slices |
| Main network | Single Transformer (text + image + audio) |
| Target | Local / consumer hardware |
| Comparison | Classic vision transformer + adapter architecture |

