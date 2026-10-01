---
id: collect-261001-ia-llm/ia-llm/introducing-nvidia-nemotron-3-nano-omni-long-context-multimodal-intelligence-for-documents-3
title: "Introducing NVIDIA Nemotron 3 Nano Omni: Long-Context Multimodal Intelligence for Documents, Audio and Video Agents"
domain: ia-llm
role: reference
task: reference
actors: ["Hugging Face", "Nvidia"]
dates: []
keywords: ["nvidia", "omni", "distribution", "fp8", "license", "nvfp4", "reasoning", "training"]
source: docs/RAG/collect-261001-ia-llm/introducing-nvidia-nemotron-3-nano-omni-long-context-multimodal-intelligence-for-documents-audio-and.md
source_anchor: ""
source_lines: [147, 243]
sha256: 4eab90da51ee4706a1385e5aa2ace1b94f5899b042cf9d4909a31a85abb03232
---

# Introducing NVIDIA Nemotron 3 Nano Omni: Long-Context Multimodal Intelligence for Documents, Audio and Video Agents

```
Successfully found the Driver License Eligibility Requirements page on the Virginia DMV website. The page contains comprehensive information about driver license eligibility including:
**General Requirements:**
- Must be a resident of the Commonwealth of Virginia
- Must be at least 16 years and 3 months of age
**First Time Drivers:**
- Must complete a state-approved driver education program
- Must obtain a learner's permit at least 60 days before taking the road skills test
- Must be at least 15 years and 6 months of age for a learner's permit
- No exceptions to age requirements
**New Residents:**
- Temporary residents: must obtain license within 6 months
- Permanent residents: must obtain license within 1 year
**Military Personnel:**
- Active-duty members of the Armed Forces stationed in Virginia
- Spouses and dependents 16 years and 3 months or older may drive with valid out-of-state license
- If vehicle is co-owned and out-of-state plates are used, all co-owners must be active duty members
**Out-of-State Students:**
- Full-time students enrolled in accredited Virginia schools (not employed) may drive with valid out-of-state license
- Full-time students who are employed are considered Virginia residents for vehicle laws
**Non-Residents:**
- Non-residents temporarily living in Virginia may drive with home state license for no more than 6 months
- Note: This does not apply to commercial vehicle drivers
- If becoming gainfully employed, must hold Virginia driver's license
**Identification and Residency Requirements:**
- Proof of identity (acceptable documents listed in DMV 141)
- Proof of legal presence in the U.S.
- Proof of Virginia residency
- Proof of Social Security number
- Proof of name change (if applicable)
Additional resources available: interactive document guide, "Acceptable Documents for Obtaining a Driver's License or Photo ID Card" (DMV 141), and Spanish language version.
```
Nemotron-3 Nano Omni can analyze charts, figures shown in images, along with audio files to generate any commonalities, discrepancies across the media.

**Audio**

**Question**

```
Use both the slide content and the spoken commentary.
List out any extra points mentioned by the speaker that is not shown in the slide
```
**Answer** 

```
Based on the speaker's commentary and the visual content of the slide, here are the extra points mentioned that are not explicitly shown or detailed in the slide:
*   **The Topic:** The speaker explicitly states she will be covering **multitask learning** with neural networks, while the slide title is simply "Some notation" (though it contains the mathematical definitions for tasks).
*   **Nature of the Output:** She describes the output ($y$) as a **probability distribution** over possible labels (e.g., tiger, tiger cat, lynx, cat), whereas the slide only shows a bar chart with class labels without explicitly labeling the vector as a distribution.
*   **Specific Task Example (Paper Titles):** She introduces a concrete example where the input is the **title of a paper** and the neural network is tasked with estimating the **length of the paper**. The slide shows "length of paper" near the bar chart, but does not visually distinguish how a paper title maps to these bars versus representing animal classes.
*   **Optimization Techniques:** She mentions that the loss function is typically optimized using **stochastic gradient descent** or a **"favorite optimizer,"** which are procedural steps not written on the slide (which only defines the loss function $\mathcal{L}$ to minimize).
*   **Input Diversity:** She specifies that inputs can be a **"wide range of unstructured inputs,"** specifically mentioning **images** and **text**. The slide visually displays a single image (a structured input) and uses a generic variable $x$ for the input class.
```
Nemotron-3 Nano Omni can analyze charts, figures shown in images, along with audio files to generate any commonalities, discrepancies across the media.

**Audio**

**Question**

```
What is producing the main sound in this audio, and what does the acoustic environment suggest about where it was recorded? 
```
**Answer**

```
The main sound is a bird chirping, and the presence of multiple bird calls and a natural, open acoustic environment suggests the recording was made outdoors in a natural setting like a forest or park.
```
Nemotron-3-Nano-Omni can interpret and reason about general audio, including soundscapes and environmental sounds.

**Audio** 

**Question** 

```
What's the overall vibe of this song, and what kind of scenario would it complement well?
```
**Answer**

```
The song has a calm, reflective, and slightly melancholic vibe, with a gentle piano melody and soft string accompaniment. It would complement a quiet, introspective scene such as a rainy evening, a peaceful walk, or a moment of personal reflection.
```
| Hugging Face BF16 checkpoint | `https://huggingface.co/nvidia/Nemotron-3-Nano-Omni-30B-A3B-Reasoning-BF16` | 
|---|---|
| Hugging Face FP8 checkpoint | `https://huggingface.co/nvidia/Nemotron-3-Nano-Omni-30B-A3B-Reasoning-FP8` | 
| Hugging Face NVFP4 checkpoint | `https://huggingface.co/nvidia/Nemotron-3-Nano-Omni-30B-A3B-Reasoning-NVFP4` | 
| Technical report / PDF | `https://arxiv.org/abs/2604.24954` | 
| Dataset / collection release | `https://huggingface.co/datasets/nvidia/Nemotron-Image-Training-v3` | 
| Megatron-Bridge | `https://github.com/NVIDIA-NeMo/Megatron-Bridge/tree/main/examples/models/vlm/nemotron_3_omni` | 
| Nemo-RL | `https://github.com/NVIDIA-NeMo/RL/blob/nano-v3-omni/docs/guides/nemotron-3-nano-omni.md` | 
| NeMo Data Designer SDG recipes | `https://github.com/NVIDIA-NeMo/DataDesigner/tree/main/docs/assets/recipes/vlm_long_doc` | 

- NVIDIA Nemotron Nano V2 VL. **Technical report:** https://arxiv.org/abs/2511.03929
- NVIDIA Nemotron 3: Efficient and Open Intelligence. **Technical report:** https://arxiv.org/abs/2512.20856
- C-RADIOv4-H. **Hugging Face model page:** https://huggingface.co/nvidia/C-RADIOv4-H
- Parakeet-TDT-0.6B-v3. **Hugging Face model page:** https://huggingface.co/nvidia/parakeet-tdt-0.6b-v3
- Megatron-LM. **GitHub:** https://github.com/NVIDIA/Megatron-LM
- Transformer Engine. **GitHub:** https://github.com/NVIDIA/TransformerEngine
- Megatron Energon. **GitHub:** https://github.com/NVIDIA/Megatron-Energon
