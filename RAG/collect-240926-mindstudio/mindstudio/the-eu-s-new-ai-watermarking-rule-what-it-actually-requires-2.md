---
id: collect-240926-mindstudio/mindstudio/the-eu-s-new-ai-watermarking-rule-what-it-actually-requires-2
title: "the-eu-s-new-ai-watermarking-rule-what-it-actually-requires"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "EU", "Google", "OpenAI", "United States"]
dates: []
keywords: ["watermarking", "chatgpt", "claude"]
source: docs/RAG/clean_en/mindstudio/the-eu-s-new-ai-watermarking-rule-what-it-actually-requires.md
source_anchor: ""
source_lines: [68, 91]
sha256: fa68f0b36a692288aca89b478cf4f9a56a3b71322a68b35d02d7f42e98c05f5e
---

# the-eu-s-new-ai-watermarking-rule-what-it-actually-requires

The obligation is aimed at providers, the companies building and deploying AI systems, not individual users. A person generating an image with a compliant tool benefits from (or is bound by) whatever labeling that tool applies automatically, but they aren’t independently required to watermark their own output.

### Is SynthID required by EU law?

No. SynthID is Google’s specific watermarking technology. The EU AI Act sets a transparency obligation but doesn’t mandate one specific technical standard, so different labs can and do use different methods to satisfy it.

### Can AI watermarks be removed?

Often, yes. Metadata-based tags are easy to strip by resaving a file. Embedded watermarks in images and audio are more resilient to common edits like cropping or compression but aren’t guaranteed to survive determined removal attempts or transformation through another AI tool.

## 
      Plans first.
      *Then code.*
    

    Remy writes the spec, manages the build, and ships the app.

### Why are US-based companies like OpenAI and Anthropic complying with an EU law?

Because building and maintaining separate versions of a model or product for different regulatory regions is costly and complex, many labs apply the stricter standard globally rather than segmenting by geography. This mirrors how GDPR shaped global privacy defaults well beyond EU borders.

### Does this affect AI chatbot conversations, like what you see from ChatGPT or Claude?

The labeling rules mostly target generated media output, images, audio, video, and in some cases long-form text, rather than the back-and-forth of a chat interface itself. Most users will notice this as labels on generated images or media files rather than any visible change inside a chatbot conversation.
