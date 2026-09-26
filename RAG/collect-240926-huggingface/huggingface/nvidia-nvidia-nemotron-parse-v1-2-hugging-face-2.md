---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-parse-v1-2-hugging-face-2
title: "Load model and processor"
domain: huggingface
role: reference
task: reference
actors: ["Hugging Face", "Nvidia", "OpenAI", "TensorRT-LLM", "vLLM"]
dates: ["2026-17-02"]
keywords: ["attention", "blackwell", "decode", "gpu", "inference", "license", "nvidia", "parameters", "tensorrt", "tensorrt-llm", "training", "vllm"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-parse-v1-2-hugging-face.md
source_anchor: ""
source_lines: [149, 292]
sha256: d9cbdc7453b9a005c8dd6ab136685a42fd0afe45fa68c63d1c4bc32340fa5545
---

# Load model and processor

```
import base64
from openai import OpenAI
client = OpenAI(
    base_url="http://localhost:8000/v1",
)
# Read and base64-encode the image
with open(<your-image-path>, "rb") as f:
    img_b64 = base64.b64encode(f.read()).decode("utf-8")
prompt_text = "</s><s><predict_bbox><predict_classes><output_markdown><predict_no_text_in_pic>"
resp = client.chat.completions.create(
    model="nvidia/NVIDIA-Nemotron-Parse-v1.2",
    messages=[
        {
            "role": "user",
            "content": [
                {
                    "type": "text",
                    "text": prompt_text,
                },
                {
                    "type": "image_url",
                    "image_url": {
                        "url": f"data:image/png;base64,{img_b64}",
                    },
                },
            ],
        }
    ],
    max_tokens=9000,
    temperature=0.0,
    extra_body={
        "repetition_penalty": 1.1,
        "top_k": 1,
        "skip_special_tokens": False,
    },
)
print(resp.choices[0].message.content)
```
*Note:* we recommend using the default prompt that extracts bounding boxes, classes, and text in markdown formatting for all use cases (`</s><s><predict_bbox><predict_classes><output_markdown><predict_no_text_in_pic>` or `</s><s><predict_bbox><predict_classes><output_markdown><predict_text_in_pic>`). If necessary, optionally the prompt that omits text extraction and only outputs bounding boxes and classes could be used: `</s><s><predict_bbox><predict_classes><output_no_text><predict_no_text_in_pic>`.

With Nemotron-Parse-v1.2 we share 2 logits processors available in logitsprocessors/ dir for vllm and in hf_logits_processor.py for the python model. NemotronParseRepetitionStopProcessor - when used during generation, detects repeating n-grams and forces the model to close the block when detecting potential hallucination. NemotronParseTableInsertionLogitsProcessor - forces every block to follow a table structure (useful if, e.g., you are running the model on table image crops)

Please refer to the example_with_processor.py for example usage with python model. With vllm, you can provide these as arguments to vllm serve, after exporting logitsprocs/ to PYTHONPATH, e.g.:

```
vllm serve nvidia/NVIDIA-Nemotron-Parse-v1.2 \
  --dtype bfloat16 \
  --max-num-seqs 4 \
  --limit-mm-per-prompt '{"image": 1}' \
  --attention-backend=TRITON_ATTN \
  --trust-remote-code \
  --logits-processors nemotron_parse_vllm_logitprocs:NemotronParseTableInsertionLogitsProcessor \
  --port 8000
```
An example of inference with vllm openai server is available in vllm_example.py

Governing Terms: Your use of this model is governed by the: NVIDIA Nemotron Open Model License. Use of the tokenizer included in this model is governed by the CC-BY-4.0 license.

Global

NVIDIA Nemotron Parse v1.2 will be capable of comprehensive text understanding and document structure understanding. It will be used in retriever and curator solutions. Its text extraction datasets and capabilities will help with LLM and VLM training, as well as improve run-time inference accuracy of VLMs. The NVIDIA Nemotron Parse v1.2 model will perform text extraction from PDF and PPT documents. The NVIDIA Nemotron Parse v1.2 can classify the objects (title, section, caption, index, footnote, lists, tables, bibliography, image) in a given document, and provide bounding boxes with coordinates.

Hugging Face [02/17/2026] via [URL] 

**Architecture Type:** Transformer-based vision-encoder-decoder model

**Network Architecture:** 

- Vision Encoder: ViT-H model (https://huggingface.co/nvidia/C-RADIO)
- Adapter Layer: 1D convolutions & norms to compress dimensionality and sequence length of the latent space (1280 tokens to 320 tokens)
- Decoder: mBart [1] 10 blocks
- Tokenizer: Use of the tokenizer included in this model is governed by the CC-BY-4.0 license
- Number of Parameters: < 1B

- Input Type: Image, Text
- Input Type(s): Red, Green, Blue (RGB) + Prompt (String)
- Input Parameters: Two-Dimensional (2D), One-Dimensional (1D)
- Other Properties Related to Input:
  - Max Input Resolution (Width, Height): 1664, 2048
  - Min Input Resolution (Width, Height): 1024, 1280
- Channel Count: 3

- Output Type: Text
- Output Format: String
- Output Parameters: One-Dimensional (1D)
- Other Properties Related to Output: Nemotron-parse output format is a string which encodes text content (formatted or not) as well as bounding boxes and class attributes.

Our AI models are designed and/or optimized to run on NVIDIA GPU-accelerated systems. By leveraging NVIDIA’s hardware (e.g. GPU cores) and software frameworks (e.g., CUDA libraries), the model achieves faster training and inference times compared to CPU-only solutions. 

**Runtime Engine(s):**

- TensorRT-LLM 
- vLLM 

**Supported Hardware Microarchitecture Compatibility:** 

- NVIDIA Ampere 
- NVIDIA Blackwell  
- NVIDIA Hopper 
- NVIDIA Turing 

**Supported Operating System(s):**

- [Linux] 

The integration of foundation and fine-tuned models into AI systems requires additional testing using use-case-specific data to ensure safe and effective deployment. Following the V-model methodology, iterative testing and validation at both unit and system levels are essential to mitigate risks, meet technical and functional requirements, and ensure compliance with safety and ethical standards before deployment. 

Nemotron Parse 1.2 

** Image Training Data Size 

- [1 Million to 1 Billion Images] 

** Text Training Data Size 

- [1 Billion to 10 Trillion Tokens] 

** Data Collection Method by dataset 

- Hybrid: Automated, Human, Synthetic 

** Labeling Method by dataset 

- Hybrid: Automated, Human, Synthetic 

**Properties (Quantity, Dataset Descriptions, Sensor(s)):** The training set contains millions of image–text items, aggregated across many large document and table datasets totaling several terabytes of data. The data consists of document-page and table images paired with OCR text, bounding boxes, and layout labels, drawn from real-world sources (scientific papers, PDFs, Wikipedia pages) as well as fully synthetic tables and word/character renderings. Modalities are primarily images plus associated text and structural annotations; content spans public-domain resources, and synthetic data. Images are obtained by rendering digital documents or generating synthetic layouts, and annotations come from OCR/layout models, third-party OCR services, and human labeling. 

**Acceleration Engine:** Tensor(RT)-LLM, vLLM 

**Test Hardware:** 

- H100 
- A100 

NVIDIA believes Trustworthy AI is a shared responsibility and we have established policies and practices to enable development for a wide array of AI applications.  When downloaded or used in accordance with our terms of service, developers should work with their internal model team to ensure this model meets requirements for the relevant industry and use case and addresses unforeseen product misuse. 

Please make sure you have proper rights and permissions for all input image and video content; if image or video includes people, personal health information, or intellectual property, the image or video generated will not blur or maintain proportions of image subjects included. 

For more detailed information on ethical considerations for this model, please see the Model Card++ Bias, Explainability, Safety & Security, and Privacy Subcards
Please report model quality, risk, security vulnerabilities or NVIDIA AI Concerns here.

- Downloads last month
- 66,995
