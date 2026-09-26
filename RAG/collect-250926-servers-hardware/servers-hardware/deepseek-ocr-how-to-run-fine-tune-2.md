---
id: collect-250926-servers-hardware/servers-hardware/deepseek-ocr-how-to-run-fine-tune-2
title: "DeepSeek-OCR: How to Run & Fine-tune"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["deepseek", "agent", "agents"]
source: docs/RAG/clean4/DeepSeek-OCR How to Run & Fine-tune.md
source_anchor: ""
source_lines: [155, 201]
sha256: 4545519aa1a0747da67d8088181fcef6107a0a3b85e5f97c19d4c4b0b96b5efb
---

# DeepSeek-OCR: How to Run & Fine-tune

============================================================
Fine-tuned Model Performance
============================================================
Number of samples: 200
Mean CER: 60.43%
Median CER: 50.00%
Std Dev: 80.63%
Min CER: 0.00%
Max CER: 916.67%
============================================================
Best Predictions (Lowest CER):
Sample 301 (CER: 0.00%)
Reference: باشه بابا تو لاکچری، تو خاص، تو خفن...
Prediction: باشه بابا تو لاکچری، تو خاص، تو خفن...
Sample 2512 (CER: 0.00%)
Reference: از شخص حاج عبدالله زنجبیلی میگیرنش...
Prediction: از شخص حاج عبدالله زنجبیلی میگیرنش...
Sample 2713 (CER: 0.00%)
Reference: نمی دونم والا تحمل نقد ندارن ظاهرا...
Prediction: نمی دونم والا تحمل نقد ندارن ظاهرا...
Worst Predictions (Highest CER):
Sample 14270 (CER: 916.67%)
Reference: ۴۳۵۹۴۷۴۷۳۸۹۰...
Prediction: پروپریپریپریپریپریپریپریپریپریپریپریپریپریپریپریپریپریپریپیپریپریپریپریپریپریپریپریپریپریپریپریپریپر...
Sample 3919 (CER: 380.00%)
Reference: ۷۵۵۰۷۱۰۶۵۹...
Prediction: وادووووووووووووووووووووووووووووووووووو...
Sample 3718 (CER: 333.33%)
Reference: ۳۲۶۷۲۲۶۵۵۸۴۶...
Prediction: پُپُسوپُسوپُسوپُسوپُسوپُسوپُسوپُسوپُسوپُ...

{% endcolumn %}
{% endcolumns %}
An example from the 200K Persian dataset we used (you may use your own), showing the image on the left and the corresponding text on the right.
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/models/tutorials/deepseek-ocr-how-to-run-and-fine-tune.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
