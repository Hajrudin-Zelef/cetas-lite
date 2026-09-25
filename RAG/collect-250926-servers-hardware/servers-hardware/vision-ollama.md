---
id: collect-250926-servers-hardware/servers-hardware/vision-ollama
title: "1. Download a sample imagecurl -L -o test.jpg \"https://upload.wikimedia.org/wikipedia/commons/3/3a/Cat03.jpg\"# 2. Encode the imageIMG=$(base64 < test.jpg | tr -d '\\n')# 3. Send it to Ollamacurl -X POST http://localhost:11434/api/chat \\-H \"Content-Type: application/json\" \\-d '{ \"model\": \"gemma4\", \"messages\": [{ \"role\": \"user\", \"content\": \"What is in this image?\", \"images\": [\"'\"$IMG\"'\"] }], \"stream\": false}'"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["decode"]
source: docs/RAG/clean4/vision-ollama.md
source_anchor: ""
source_lines: [1, 13]
sha256: 91c5277df98e0ce43e99c93e72170f2fbb2351add958a23b47ba519254c2b643
---

# 1. Download a sample imagecurl -L -o test.jpg "https://upload.wikimedia.org/wikipedia/commons/3/3a/Cat03.jpg"# 2. Encode the imageIMG=$(base64 < test.jpg | tr -d '\n')# 3. Send it to Ollamacurl -X POST http://localhost:11434/api/chat \-H "Content-Type: application/json" \-d '{ "model": "gemma4", "messages": [{ "role": "user", "content": "What is in this image?", "images": ["'"$IMG"'"] }], "stream": false}'

Provide an images array. SDKs accept file paths, URLs or raw bytes while the REST API expects base64-encoded image data.

cURL

Python

JavaScript

# 1. Download a sample imagecurl -L -o test.jpg "https://upload.wikimedia.org/wikipedia/commons/3/3a/Cat03.jpg"# 2. Encode the imageIMG=$(base64 < test.jpg | tr -d '\n')# 3. Send it to Ollamacurl -X POST http://localhost:11434/api/chat \-H "Content-Type: application/json" \-d '{ "model": "gemma4", "messages": [{ "role": "user", "content": "What is in this image?", "images": ["'"$IMG"'"] }], "stream": false}'

from ollama import chat# from pathlib import Path# Pass in the path to the imagepath = input('Please enter the path to the image: ')# You can also pass in base64 encoded image data# img = base64.b64encode(Path(path).read_bytes()).decode()# or the raw bytes# img = Path(path).read_bytes()response = chat( model='gemma4', messages=[ { 'role': 'user', 'content': 'What is in this image? Be concise.', 'images': [path], } ],)print(response.message.content)

import ollama from 'ollama'const imagePath = '/absolute/path/to/image.jpg'const response = await ollama.chat({ model: 'gemma4', messages: [ { role: 'user', content: 'What is in this image?', images: [imagePath] } ], stream: false,})console.log(response.message.content)
