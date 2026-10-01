---
id: collect-261001-ia-llm/ia-llm/hugging-face-and-cerebras-bring-gemma-4-to-real-time-voice-ai
title: "hugging-face-and-cerebras-bring-gemma-4-to-real-time-voice-ai"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Cerebras", "Google", "Hugging Face", "Nvidia"]
dates: []
keywords: ["voice", "cost", "inference", "latency", "multimodal", "nvidia", "qwen", "research"]
source: docs/RAG/collect-261001-ia-llm/hugging-face-and-cerebras-bring-gemma-4-to-real-time-voice-ai.md
source_anchor: ""
source_lines: [1, 56]
sha256: cb949444ac7eb35090f49001fd4d90a1947383499a992e30d215104ca36956b7
---

# hugging-face-and-cerebras-bring-gemma-4-to-real-time-voice-ai

#### HF Realtime Voice

Realtime voice over WebSocket or WebRTC

The result is a speech-to-speech experience that feels dramatically more natural. Instead of waiting for an AI to respond, conversations flow with the responsiveness users expect from human interaction.

The demo is built as a real-time speech-to-speech pipeline. Each part of the system is modular, open, and replaceable, making it easy for developers to adapt the stack for different assistants, robots, products, or research projects.

This creates a fully open speech-to-speech loop:

```
Speech input
  -> speech recognition with Nvidia's Parakeet
  -> Gemma 4 VLM inference on Cerebras
  -> text-to-speech with Alibaba's Qwen3TTS
  -> spoken response
```
The architecture brings together the strength of the open-source AI ecosystem: Cerebras for fast inference, Google DeepMind’s Gemma 4 31B for the language model, and Qwen for text-to-speech. Every layer can be inspected, modified, and extended by the developers

Today, some production systems see a reasonable median latency while still experiencing frustrating multi-second delays at the P95. Those delays become even more noticeable when tool calls or multimodal steps require multiple turns.

Cerebras helps solve one of the most important bottlenecks in the stack: the language-model response time. By making inference dramatically faster and more stable, Cerebras allows the rest of the Hugging Face pipeline to shine.

That stability is especially important at the long tail. Many systems can deliver acceptable median response times, but occasional slow responses still make conversations feel unreliable.

This same Hugging Face speech-to-speech pipeline already powers Reachy Mini robots, with more than 9,000 robots in the wild. For robots, voice assistants, and embodied AI, responsiveness is not a cosmetic improvement. It is what makes the interaction feel alive.

The motivation to use Cerebras is therefore not simply cost reduction. It is low latency, predictable performance, and the ability to create real-time experiences that feel natural at scale.

This collaboration reflects a shared belief that the future of AI will be both open and performant. Open-source models, open infrastructure, and breakthrough inference speed together create a foundation for the next generation of conversational AI.

We invite developers to explore the demo, experiment with the code, and help shape what comes next for real-time voice AI.

Demo: Hugging Face Space

Repository: huggingface/speech-to-speech

It's exciting to see how Hugging Face and Cerebras are pushing real-time voice technology forward. Fast responses make conversations feel much more comfortable and engaging, and that's exactly what users have been hoping for. 😊 When the delay is reduced, interactions become smoother and more natural, making voice experiences feel less like talking to software and more like having a real conversation.

What I appreciate most is the focus on an open and modular approach. That gives developers more flexibility to build solutions that fit different needs instead of being locked into a single ecosystem. The combination of strong model quality with impressive inference speed could help unlock many practical applications across education, customer support, accessibility, and beyond. 🚀

The emphasis on latency is absolutely the right direction because speed directly shapes the overall experience. Even the smartest system can feel frustrating if every response takes too long. Bringing that waiting time down creates a much more fluid interaction, and it's great to see innovation focused on something users notice immediately. 👏

Looking forward to seeing how this evolves over the coming months... Faster conversations... Better responsiveness... More natural interactions... Exciting progress for developers and users alike... 🌟🙂

Does anybody work on multipoint speech2speech interaction? I.e. when 2 or more people are talking to AI?

I think LLM can handle messages from several persons when they are properly tagged.

How about speech2text models - are there any works in this direction? The model would have to have some context and probably stereo audio input or some additional tagging to differentiate between persons.

As an example  practical usage: voice assistant in car, when we have more than one person there, not unusual.

Also some public kiosks where a group of people can interact by voice.

There is much more in it then just text tagging by voice, as in this case people could also talk to each other (not to AI) and here also some capturing of non-verbal communication will be needed.
