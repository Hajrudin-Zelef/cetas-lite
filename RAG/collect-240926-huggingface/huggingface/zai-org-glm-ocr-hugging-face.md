---
id: collect-240926-huggingface/huggingface/zai-org-glm-ocr-hugging-face
title: "zai-org-glm-ocr-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["SGLang", "Z.ai", "vLLM"]
dates: []
keywords: ["glm", "apache", "benchmarks", "compute", "cost", "inference", "latency", "leaderboard", "license", "licenses", "mit license", "multimodal"]
source: docs/RAG/clean_en/huggingface/zai-org-glm-ocr-hugging-face.md
source_anchor: ""
source_lines: [1, 107]
sha256: ebc0c116344efd38a22fc7ef81c930c60cdf0930ddb8adacdd6cdc83ee84d27e
---

# zai-org-glm-ocr-hugging-face

<!-- source: https://huggingface.co/zai-org/GLM-OCR -->

👋 Join our WeChat and Discord community
    

    📍 Use GLM-OCR's API
    

    👉 GLM-OCR SDK Recommended
    

    📖  Technical Report

GLM-OCR is a multimodal OCR model for complex document understanding, built on the GLM-V encoder–decoder architecture. It introduces Multi-Token Prediction (MTP) loss and stable full-task reinforcement learning to improve training efficiency, recognition accuracy, and generalization. The model integrates the CogViT visual encoder pre-trained on large-scale image–text data, a lightweight cross-modal connector with efficient token downsampling, and a GLM-0.5B language decoder. Combined with a two-stage pipeline of layout analysis and parallel recognition based on PP-DocLayout-V3, GLM-OCR delivers robust and high-quality OCR performance across diverse document layouts.

**Key Features**

- **State-of-the-Art Performance** : Achieves a score of 94.62 on OmniDocBench V1.5, ranking #1 overall, and delivers state-of-the-art results across major document understanding benchmarks, including formula recognition, table recognition, and information extraction.
- **Optimized for Real-World Scenarios** : Designed and optimized for practical business use cases, maintaining robust performance on complex tables, code-heavy documents, seals, and other challenging real-world layouts.
- **Efficient Inference** : With only 0.9B parameters, GLM-OCR supports deployment via vLLM, SGLang, and Ollama, significantly reducing inference latency and compute cost, making it ideal for high-concurrency services and edge deployments.
- **Easy to Use** : Fully open-sourced and equipped with a comprehensive SDK and inference toolchain, offering simple installation, one-line invocation, and smooth integration into existing production pipelines.

- Document Parsing & Information Extraction

- Real-World Scenarios Performance

- Speed Test

For speed, we compared different OCR methods under identical hardware and testing conditions (single replica, single concurrency), evaluating their performance in parsing and exporting Markdown files from both image and PDF inputs. Results show GLM-OCR achieves a throughput of 1.86 pages/second for PDF documents and 0.67 images/second for images, significantly outperforming comparable models.

For document parsing tasks, we strongly recommend using our official SDK. Compared with model-only inference, the SDK integrates PP-DocLayoutV3 and provides a complete, easy-to-use pipeline for document parsing, including layout analysis and structured output generation. This significantly reduces the engineering overhead required to build end-to-end document intelligence systems.

Note that the SDK is currently designed for document parsing tasks only. For information extraction tasks, please refer to the following section and run inference directly with the model.

GLM-OCR supports deployment with the following frameworks. Feel free to try them out:

- SGLang — see cookbook
- vLLM — see recipes
- Transformers — see transformers docs

GLM-OCR currently supports two types of prompt scenarios:

1. **Document Parsing** – extract raw content from documents. Supported tasks include:

```
{
    "text": "Text Recognition:",
    "formula": "Formula Recognition:",
    "table": "Table Recognition:"
}
```
1. **Information Extraction** – extract structured information from documents. Prompts must follow a strict JSON schema. For example, to extract personal ID information:

```
请按下列JSON格式输出图中信息:
{
    "id_number": "",
    "last_name": "",
    "first_name": "",
    "date_of_birth": "",
    "address": {
        "street": "",
        "city": "",
        "state": "",
        "zip_code": ""
    },
    "dates": {
        "issue_date": "",
        "expiration_date": ""
    },
    "sex": ""
}
```
⚠️ Note: When using information extraction, the output must strictly adhere to the defined JSON schema to ensure downstream processing compatibility.

This project is inspired by the excellent work of the following projects and communities:

The GLM-OCR model is released under the MIT License.

The complete OCR pipeline integrates PP-DocLayoutV3 for document layout analysis, which is licensed under the Apache License 2.0. Users should comply with both licenses when using this project.

If you find GLM-OCR useful in your research, please cite our technical report:

```
@misc{duan2026glmocrtechnicalreport,
      title={GLM-OCR Technical Report},
      author={Shuaiqi Duan and Yadong Xue and Weihan Wang and Zhe Su and Huan Liu and Sheng Yang and Guobing Gan and Guo Wang and Zihan Wang and Shengdong Yan and Dexin Jin and Yuxuan Zhang and Guohong Wen and Yanfeng Wang and Yutao Zhang and Xiaohan Zhang and Wenyi Hong and Yukuo Cen and Da Yin and Bin Chen and Wenmeng Yu and Xiaotao Gu and Jie Tang},
      year={2026},
      eprint={2603.10910},
      archivePrefix={arXiv},
      primaryClass={cs.CL},
      url={https://arxiv.org/abs/2603.10910},
}
```
- Downloads last month
- 1,712,126

## Spaces using zai-org/GLM-OCR 98

## Paper for zai-org/GLM-OCR

- allenai/olmOCR-bench leaderboard
- Overall View evaluation resultsExcluding Headers & Footers category. Using ZAI API.75.2<sup>*</sup>
- Arxiv Math View evaluation results80.7
- Old Scans Math View evaluation results68.3
- llamaindex/ParseBench leaderboard
- Mean View evaluation resultssourcePipeline name: glmocr_pipeline29.6<sup>*</sup>
