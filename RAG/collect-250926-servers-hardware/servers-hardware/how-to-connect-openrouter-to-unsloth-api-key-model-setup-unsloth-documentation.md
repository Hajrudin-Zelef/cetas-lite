---
id: collect-250926-servers-hardware/servers-hardware/how-to-connect-openrouter-to-unsloth-api-key-model-setup-unsloth-documentation
title: "how-to-connect-openrouter-to-unsloth-api-key-model-setup-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI", "OpenRouter", "Unsloth"]
dates: []
keywords: ["tool calling"]
source: docs/RAG/clean4/how-to-connect-openrouter-to-unsloth-api-key-model-setup-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 21]
sha256: d79ea191c5b3c484dff1590789de76543b59de5efc73e83cfd23fcb43fe63950
---

# how-to-connect-openrouter-to-unsloth-api-key-model-setup-unsloth-documentation

This guide explains how to connect **OpenRouter to** **Unsloth** so you can access hosted AI models from providers like **OpenAI, Anthropic,** and **Google** through an open-source local UI chat interface. You’ll learn how to create an OpenRouter API key, add OpenRouter as a provider in Unsloth, load or manually enter model IDs, and enable external models for chat.

Once a single API key is connected, OpenRouter models in Unsloth can provide advanced features such as thinking, web search, tool calling, code execution, and customizable generation settings directly from the chat page.

Sign in to your OpenRouter account. Create an API key from the OpenRouter dashboard:

Copy the key. You will paste it into Unsloth in the next step.

When creating the key, you can optionally set a credit limit or expiration date.

OpenRouter provides access to many models from different providers. If **Load Models** does not return the models you want to select, enter the model IDs you want enabled.

Example model IDs:

If OpenRouter fails to connect, check that the API key is valid and belongs to the correct OpenRouter account.

If a model does not appear after clicking **Load Models**, it may not be available for your account or region. You can enter the model ID manually or choose another model.

Last updated

Was this helpful?
