---
id: collect-261001-ia-llm/ia-llm/introducing-nvidia-nemotron-3-nano-omni-long-context-multimodal-intelligence-for-documents-2
title: "Introducing NVIDIA Nemotron 3 Nano Omni: Long-Context Multimodal Intelligence for Documents, Audio and Video Agents"
domain: ia-llm
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["agent", "multimodal", "nvidia", "omni", "agentic", "benchmark", "embedding", "license", "memory", "parameters", "reasoning", "tool calling"]
source: docs/RAG/collect-261001-ia-llm/introducing-nvidia-nemotron-3-nano-omni-long-context-multimodal-intelligence-for-documents-audio-and.md
source_anchor: ""
source_lines: [66, 146]
sha256: 70154efde42fffedad84ec066fe95a5521ee58a0664366144a0fab9cd7a45c14
---

# Introducing NVIDIA Nemotron 3 Nano Omni: Long-Context Multimodal Intelligence for Documents, Audio and Video Agents

This represents a shift from traditional VLM pipelines by enabling native audio processing within a shared multimodal sequence, allowing audio, visual, and text tokens to be jointly modeled. This is crucial for scenarios like narrated screen recordings, video Q&A where speech alters visual meaning, long-form instructional or meeting content, and tasks requiring temporally grounded multimodal reasoning.

Each encoder is connected to the LLM with a lightweight **2-layer MLP projector** that maps encoder features into the shared embedding space. Once projected, **vision, audio, and text tokens are interleaved and processed jointly**.

This design keeps the overall system modular while still enabling genuine cross-modal reasoning inside the backbone itself.

The SFT stages are trained on **NVIDIA H100**, scaling from **32 to 128 nodes** depending on the stage. The stack uses **Megatron-LM**, **Transformer Engine**, and **Megatron Energon**, with tensor parallelism, expert parallelism, sequence parallelism, context parallelism for the long context stages, online sequence packing, and selective activation recomputation.

Post-SFT reinforcement learning uses **NeMo-RL** **and NeMo Gym** with a Megatron backend. The RL infrastructure used a **Ray-based distributed setup** across **B200 and H100 clusters**, plus multimodal deduplication, so repeated rollouts do not multiply image, video, and audio memory.

We open-source substantial parts of our training code.

We introduce multi-environment text and omni training in Nemotron 3 Nano Omni. Our text RL training stage happens across diverse environments in Nemo-Gym, which evaluates the model’s ability to perform sequences of actions such as tool calling, writing code, and multi-part planning that satisfy the verifiable criteria.

Omni RL trains the model to reason **across images, video, audio, and text** within a unified framework, covering tasks from **single-modality to fully multimodal scenarios**. A diverse verifier suite evaluates outputs across formats like multiple-choice, math, GUI grounding, and ASR, while intentionally including unanswerable cases to teach the model to abstain when evidence is insufficient rather than hallucinate.

Nemotron 3 Nano Omni is trained on an enhanced dataset that emphasizes high-quality reasoning across multiple modalities. We significantly expand task coverage and introduce synthetic data for complex reasoning scenarios where public datasets are limited. To support this, we build task-specific, multi-stage pipelines for scalable synthetic data generation.

As one example, we generated approximately 11.4M synthetic QA pairs (~45B tokens) from a large corpus of real-world PDFs using NeMo Data Designer. This dataset is used to strengthen long-context document reasoning during post-training and delivers a 2.19× improvement in overall accuracy on MMLongBench-Doc.

We detail the full pipeline evolution, including failure analysis and key lessons learned, in our Data Designer developer note. The note also includes nine runnable pipeline recipes that can serve as starting points for building your own document-understanding datasets.

Nemotron 3 Nano Omni can analyze and reason over long documents such as financial reports, academic papers, product manuals, etc. The following example retrieves financial metrics across a 100+ page document to calculate another metric:

Figure 3: Stylized example from the MMLongBench-Doc benchmark

The model was prompted to answer the question as follows:

```
Extract information from all the given images, then answer the question using a single word or phrase. Return 'Not answerable' if the answer cannot be derived from the the images.
```
The model is capable of long-context retrieval, structured extraction, table/chart reading and multi-page reasoning all in one go.

Nemotron-3 Nano Omni performs joint audio-visual analysis, both locally for specific scenes and globally across the entire video. This enables it to answer complex questions that require cross-modal reasoning—such as identifying specific visuals shown exactly when a certain topic is mentioned in the audio.

**Video**  

**Question**

```
Watch the video and listen to the narration.
1.What structure is on fire as shown in the video and how much money was being spent in its renovation project ? 
2. Describe in short what visuals are shown when the eye-witness narrates her experience?
```
**Answer** 

```
1. The structure on fire is the Notre Dame Cathedral, which was undergoing an almost $7 million renovation project.
2. When the eye-witness narrates her experience, the visuals show the cathedral engulfed in flames and thick smoke, scaffolding surrounding the burning roof, firefighters using ladders to spray water, a crowd of people watching from the riverbank, and views of the burning structure against the Paris skyline.
```
Nemotron 3 Omni can be integrated into agentic computer-use systems to reason over user intents, analyze GUI elements, and execute actions to accomplish tasks.

The model was given access to a web environment with the following system prompt:

```
You are a GUI agent. You are given an instruction, a screenshot of the screen and your previous interactions with the computer. You need to perform a series of actions to complete the task. The passoword of the computer is {password}.
For each step, provide your response in this format:
{thought}
## Action:
{action}
## Code:
{code}
In the code section, the code should be either pyautogui code or one of the following functions wrapped in the code block:
- {"name": "computer.wait", "description": "Make the computer wait for 20 seconds for installation, running code, etc.", "parameters": {"type": "object", "properties": {}, "required": []}}
- {"name": "computer.terminate", "description": "Terminate the current task and report its completion status", "parameters": {"type": "object", "properties": {"status": {"type": "string", "enum": ["success", "failure"], "description": "The status of the task"}, "answer": {"type": "string", "description": "The answer of the task"}}, "required": ["status"]}}
```
**User Instruction**

```
Find the Driver License Eligibility Requirements
```
**Model Response**:

The GIF above shows the actions that the agent takes to fulfill the user’s request

1. It clicks on the “License & IDs” tab
2. Then it clicks on “Driver’s License” from the dropdown menu
3. Then it clicks on “Driver’s License Eligibility” on the left
4. It scrolls down on this page and back up
5. Finally it produces the following response for the user:

