---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-parse-v1-2-hugging-face
title: "Load model and processor"
domain: huggingface
role: reference
task: reference
actors: ["Hugging Face", "Nvidia", "OpenAI", "TensorRT-LLM", "vLLM"]
dates: ["2026-17-02"]
keywords: ["agentic", "attention", "blackwell", "decode", "gpu", "inference", "license", "nvidia", "parameters", "tensorrt", "tensorrt-llm", "training"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-parse-v1-2-hugging-face.md
source_anchor: ""
source_lines: [1, 292]
sha256: 09dc0c67a59c3eb4eeb03eee9cb4f4aa2c32389d33f96f6e3bab2e3715f97544
---

# Load model and processor

<!-- source: https://huggingface.co/nvidia/NVIDIA-Nemotron-Parse-v1.2 -->

NVIDIA Nemotron Parse v1.2 is designed to understand document semantics and extract text and tables elements with spatial grounding. Given an image, NVIDIA Nemotron Parse v1.2 produces structured annotations, including formatted text, bounding-boxes and the corresponding semantic classes, ordered according to the document's reading flow. It overcomes the shortcomings of traditional OCR technologies that struggle with complex document layouts with structural variability, and helps transform unstructured documents into actionable and machine-usable representations. This has several downstream benefits such as increasing the availability of training-data for Large Language Models (LLMs), improving the accuracy of extractor, curator, retriever and AI agentic applications, and enhancing document understanding pipelines.

This model is ready for commercial use. 

Compared to NVIDIA-Nemotron-Parse v1.1 the following changes are made in v1.2:

NVIDIA-Nemotron-Parse v1.2 extends the prompt interface with a fourth prompt token category — **Text-in-picture prompts** — controlling whether the model extracts text from embedded images (e.g. photographs, figures):

- `<predict_text_in_pic>` : extract text found inside pictures.
- `<predict_no_text_in_pic>` : skip text extraction from pictures, reducing generation time by avoiding unnecessary tokens.

The full v1.2 prompt now consists of four prompt tokens (one from each category):

`</s><s><predict_bbox><predict_classes><output_markdown><predict_no_text_in_pic>`

**Important:** NVIDIA-Nemotron-Parse v1.1 used a three-token prompt (e.g. `</s><s><predict_bbox><predict_classes><output_markdown>`). In v1.2, the fourth token is required. The prompt is not validated at runtime — the model will still generate output with the old three-token prompt — but because all v1.2 training data included four prompt tokens, generation quality will be significantly degraded without it.


Output elements now follow the document's natural reading order across all semantic classes, including Footnotes, Page-Footers, Tables, Pictures, and Captions.

Note that input image size in v1.2 changed from 2048x1648 to 2048x1664 compared to v1.1

You can use a public image *nvcr.io/nvidia/pytorch:25.03-py3* with the following library versions installed on top:

```
pip install accelerate==1.12.0
pip install albumentations==2.0.8
pip install transformers==5.6.1
pip install timm==1.0.22
```
```
import torch
from PIL import Image, ImageDraw
from transformers import AutoModel, AutoProcessor, AutoTokenizer, AutoConfig, AutoImageProcessor, GenerationConfig
from postprocessing import extract_classes_bboxes, transform_bbox_to_original, postprocess_text
# Load model and processor
model_path = "nvidia/NVIDIA-Nemotron-Parse-v1.2"  # Or use a local path
device = "cuda:0"
model = AutoModel.from_pretrained(
    model_path,
    trust_remote_code=True,
    torch_dtype=torch.bfloat16
).to(device).eval()
tokenizer = AutoTokenizer.from_pretrained(model_path)
processor = AutoProcessor.from_pretrained(model_path, trust_remote_code=True)
# Load image
image = Image.open("path/to/your/image.jpg")
task_prompt = "</s><s><predict_bbox><predict_classes><output_markdown><predict_no_text_in_pic>"
# task_prompt = "</s><s><predict_bbox><predict_classes><output_markdown><predict_text_in_pic>"
# Process image
inputs = processor(images=[image], text=task_prompt, return_tensors="pt", add_special_tokens=False).to(device)
generation_config = GenerationConfig.from_pretrained(model_path, trust_remote_code=True)
# Generate text
outputs = model.generate(**inputs,  generation_config=generation_config)
# Decode the generated text
generated_text = processor.batch_decode(outputs, skip_special_tokens=True)[0]
```
```
from PIL import Image, ImageDraw
from postprocessing import extract_classes_bboxes, transform_bbox_to_original, postprocess_text
classes, bboxes, texts = extract_classes_bboxes(generated_text)
bboxes = [transform_bbox_to_original(bbox, image.width, image.height) for bbox in bboxes]
# Specify output formats for postprocessing
table_format = 'latex' # latex | HTML | markdown | json | csv
text_format = 'markdown' # markdown | plain
blank_text_in_figures = False # remove text inside 'Picture' class
texts = [postprocess_text(text, cls = cls, table_format=table_format, text_format=text_format, blank_text_in_figures=blank_text_in_figures) for text, cls in zip(texts, classes)]
for cl, bb, txt in zip(classes, bboxes, texts):
    print(cl, ': ', txt)
draw = ImageDraw.Draw(image)
for bbox in bboxes:
  draw.rectangle((bbox[0], bbox[1], bbox[2], bbox[3]), outline="red")
```
Supported values for `table_format`: `'latex'` | `'HTML'` | `'markdown'` | `'json'` | `'csv'`

Nemotron-Parse-v1.2 is available in vllm main and can be found in vllm/vllm-openai:v0.14.1 docker image.

Note: when running on A100/A10 we recommend running vllm serve with *--attention-backend=TRITON_ATTN*

You will need to install the following dependencies on top, and then follow the VLLM Inference example below:

```
pip install albumentations timm open_clip_torch
```
```
from vllm import LLM, SamplingParams
from PIL import Image
def main():
    sampling_params = SamplingParams(
        temperature=0,
        top_k=1,
        repetition_penalty=1.1,
        max_tokens=9000,
        skip_special_tokens=False,
    )
    
    llm = LLM(
        model="nvidia/NVIDIA-Nemotron-Parse-v1.2",
        max_num_seqs=64,
        limit_mm_per_prompt={"image": 1},
        dtype="bfloat16",
        trust_remote_code=True,
    )
    
    image = Image.open("<YOUR-IMAGE-PATH>")
    
    prompts = [
        {  # Implicit prompt
            "prompt": "</s><s><predict_bbox><predict_classes><output_markdown><predict_no_text_in_pic>",
            "multi_modal_data": {
                "image": image
            },
        },
        {  # Explicit encoder/decoder prompt
            "encoder_prompt": {
                "prompt": "",
                "multi_modal_data": {
                    "image": image
                },
            },
            "decoder_prompt": "</s><s><predict_bbox><predict_classes><output_markdown><predict_no_text_in_pic>",
        },
    ]
    
    outputs = llm.generate(prompts, sampling_params)
    
    for output in outputs:
        prompt = output.prompt
        generated_text = output.outputs[0].text
        print(f"Decoder prompt: {prompt!r}, Generated text: {generated_text!r}")
if __name__ == "__main__":
    main()
```
Alternatively, you can start a vllm server as:

```
vllm serve nvidia/NVIDIA-Nemotron-Parse-v1.2 \
    --dtype bfloat16 \
    --max-num-seqs 8 \
    --limit-mm-per-prompt '{"image": 1}' \
    --trust-remote-code \
    --port 8000 \
    --chat-template chat_template.jinja
```
with *chat_template.jinja* provided in this repository. Then, you can run inference as:

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
