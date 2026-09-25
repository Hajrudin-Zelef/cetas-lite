---
id: collect-250926-servers-hardware/servers-hardware/modelfile-reference-ollama
title: "modelfile-reference-ollama"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["gguf", "license", "parameters", "safetensors"]
source: docs/RAG/clean4/modelfile-reference-ollama.md
source_anchor: ""
source_lines: [1, 63]
sha256: 2cf2a8d66cbb6a588a2fb7c84afb73e03f3f3bef4b133261c5b2d8f329929a94
---

# modelfile-reference-ollama

## Table of Contents

## Format

The format of the`Modelfile`:
## Examples

### Basic `Modelfile`

An example of a `Modelfile` creating a mario blueprint:
1. Save it as a file (e.g. `Modelfile` )
2. `ollama create choose-a-model-name -f <location of the file e.g. ./Modelfile>`
3. `ollama run choose-a-model-name`
4. Start using the model!

`ollama show --modelfile` command.
## Instructions

### FROM (Required)

The`FROM` instruction defines the base model to use when creating a model.
#### Build from existing model

## Model library

Browse available models

#### Build from a Safetensors model

#### Build from a GGUF model

`Modelfile` location.
For a split GGUF model, keep the original split filenames and use a wildcard that matches every shard. Shards can also be listed with separate `FROM` instructions.
### PARAMETER

The`PARAMETER` instruction defines a parameter that can be set when the model is run.
#### Valid Parameters and Values

### TEMPLATE

`TEMPLATE` of the full prompt template to be passed into the model. It may include (optionally) a system message, a user’s message and the response from the model. Note: syntax may be model specific. Templates use Go template syntax.
#### Template Variables

### SYSTEM

The`SYSTEM` instruction specifies the system message to be used in the template, if applicable.
### LICENSE

The`LICENSE` instruction allows you to specify the legal license under which the model used with this Modelfile is shared or distributed.
### MESSAGE

The`MESSAGE` instruction allows you to specify a message history for the model to use when responding. Use multiple iterations of the MESSAGE command to build up a conversation which will guide the model to answer in a similar way.
#### Valid roles

#### Example conversation

### REQUIRES

The`REQUIRES` instruction allows you to specify the minimum version of Ollama required by the model.
## Notes

- the **`Modelfile` is not case sensitive** . In the examples, uppercase instructions are used to make it easier to distinguish it from arguments.
- Instructions can be in any order. In the examples, the `FROM` instruction is first to keep it easily readable.
