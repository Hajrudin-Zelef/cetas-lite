---
id: collect-korben/korben/gpt-6-astra-dechiffre-un-message-allemand-de-1918
title: "GPT-6 Astra déchiffre un message allemand de 1918"
domain: korben
role: reference
task: article
actors: ["OpenAI", "United States"]
dates: ["1918-11", "2026-09-23"]
keywords: ["astra", "gpt-6", "transcription"]
source: docs/RAG/Collect RAG/01_korben/gpt-6-astra-dechiffre-un-message-allemand-de-1918.md
source_anchor: ""
source_lines: [1, 54]
sha256: 29f2d3a110d8dd8dc9589b9dd1fb50c1973202c87114d29ea98f8aca67b5895f
---

# GPT-6 Astra déchiffre un message allemand de 1918

## Metadata

- **Source** : https://korben.info/gpt-6-astra-dechiffre-un-message-allemand-de-1918.html
- **Site** : Korben.info
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article reports that a German radio message dated 27 November 1918, sent after the armistice and preserved among unresolved encrypted texts, was deciphered using GPT-6 Astra, OpenAI's AI model. The developer Prinz recounted obtaining a reading of the message, which reported the movement of British ships at Sevastopol, in Crimea.

The document belonged to a set of military transmissions listed among the 50 cryptographic puzzles on the German blog of Klaus Schmeh. Hundreds of similar messages had already been decrypted through the work of specialist George Lasry and his colleagues, but some texts had resisted, sometimes due to missing or badly transcribed characters.

The Germans used the ADFGVX cipher, named after the six letters used to encode the text. The alphabet and digits are placed in a 6x6 grid, and each character becomes a pair of coordinates, much like the game of Battleship. A second mixing step changes the order of the symbols by rearranging columns according to a key.

Astra used TRUPPENVERSCHIEBUNG, a key already documented in the work of American officer James Rives Childs. Its 19 letters make it possible to reorder the message's 170 symbols before recovering the German plaintext with the grid. The author notes the finding is genuinely nice even though the key was already available in the archives.

The resulting text describes the arrival of an English cruiser at Sevastopol, followed by an allied squadron on the 26th. The first digit of the cruiser's arrival date is illegible in this reading, which yields a day ending in 4, a small imperfection retained in the published result. The logbooks of the British cruiser Canterbury provide a convincing cross-check: they indicate its arrival at Sevastopol on 24 November and that of an allied squadron on the 26th. Prinz explains that Astra went to find these archives to verify its answer, and both events are indeed present in the online transcription. Prinz states he knows of no prior decryption of this message.

## Key points

- A German radio message of 27 November 1918, sent after the armistice, was deciphered with GPT-6 Astra (OpenAI).
- The document was part of the 50 cryptographic puzzles listed on Klaus Schmeh's German blog.
- The cipher is ADFGVX, using a 6x6 grid and a column-rearranging transposition key.
- Astra reused the key TRUPPENVERSCHIEBUNG, already documented by US officer James Rives Childs.
- Its 19 letters reordered 170 encrypted symbols to recover the German plaintext.
- The message reports a British cruiser and then an allied squadron arriving at Sevastopol (Crimea).
- Cross-check: logs of the cruiser Canterbury show arrival on 24 November and allied squadron on the 26th.
- Prinz reports no known prior decryption of this message.

## Technical data / figures

| Item | Value |
|---|---|
| Model used | GPT-6 Astra (OpenAI) |
| Message date | 27 November 1918 |
| Cipher | ADFGVX |
| Grid | 6x6 (alphabet + digits) |
| Key used | TRUPPENVERSCHIEBUNG |
| Key length | 19 letters |
| Encrypted symbols | 170 |
| Documented by | James Rives Childs (US officer) |
| Puzzle list | 50 cryptographic puzzles (Klaus Schmeh blog) |
| Cross-check ship | HMS Canterbury |
| Reported arrivals | Cruiser 24 Nov, allied squadron 26 Nov |
| Developer | Prinz |

## Why this source matters for the RAG

This source illustrates a concrete application of a frontier AI model (GPT-6 Astra) to historical cryptanalysis, including the ADFGVX cipher mechanics and the role of archival cross-verification. It provides specific technical figures (key, symbol count, dates) useful for questions on AI-assisted cryptography and World War I communications.
