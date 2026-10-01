---
id: collect-261001-huawei/huawei/ai-campus-white-paper-26h1-0f8c526c-47
title: "ai-campus-white-paper-26h1-0f8c526c"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["alignment", "energy", "inference"]
source: docs/RAG/collect-261001-huawei/ai-campus-white-paper-26h1-0f8c526c.md
source_anchor: ""
source_lines: [3982, 4069]
sha256: 0a32156f1d5830cf3cc5bb4a5c6eb23624bab24e4c7d18bc5d097dcea162de8b
---

# ai-campus-white-paper-26h1-0f8c526c

Diverse devices and protocols create a Tower of Babel                                                as popular ontologies) and standards like IFC
where data can't interoperate. Data standardization is                                               and Brick Schema to semantically tag terms
a linguistic foundation of sorts for AI campuses, in this                                            within models. For example, energy consumption
way. Without unified semantics, data becomes digital                                                 consistently means kWh, and zone links to
waste. Only a consistent thing model lets thousands                                                  GIS coordinates. This semantic layer decouples
of devices speak the same language, providing a                                                      business logic from data formats, enabling precise
trustworthy base for intelligent applications. This                                                  cross-system queries and inference.
section examines how native protocol unification
and semantic modeling enable data to be universally                                                  For legacy non-standard devices, lightweight expert
interpretable from the point of generation.                                                          models at the edge identify proprietary messages in
                                                                                                     real time and automatically map them to a unified
3.2.1.1 One Thing, One Model                                                                         campus dictionary. This AI-driven format conversion
                                                                                                     replaces manual plugin development, greatly
Each physical entity (e.g., access control, meters,                                                  reducing digital transformation effort and ensuring
cameras) is abstracted into a structured digital                                                     a highly standardized data foundation.
template capturing their core attribute, state, and
behavior. For example, the intelligent lighting                                                      Semantic alignment also involves spatial modeling,
model standardizes fields like on/off status, color                                                  upgrading traditional key-value data into attribute
temperature, power consumption, and location.                                                        sets enriched with spatial and temporal context. For
Regardless of their vendors, all lights are mapped into                                              instance, a model not only records AC temperature
a common logical structure. This avoids redundant                                                    25°C but also links it to its building information
modeling and enhances system interoperability.                                                       modeling (BIM) coordinates, workstation ID, and
                                                                                                     energy consumption weight. This self-identifying
3.2.1.2 Semantic Alignment                                                                           data structure provides ready-to-use inputs for AI-
                                                                                                     driven indoor path planning and precise energy-
We integrate vertical industry modeling (such                                                        efficiency-based scheduling.

                                                   Energy
                                                                                         Asset               Network
                                                consumption                                                                                               ...
                                                                                       application          management
                                                 application

                       Digital world                                                      Spatial data openness
                                                              (spatial relationship, devices, indicators, events, and intelligent interaction)

                                                              Spatial assets                                             Soft spatial indicators
                                                    Region, location, description,                                     People ﬂow, energy consumption,
                                                          and association                                                      asset, and event                         Data task
                                                                            Floor
                                                   Building




                                                                                            Room
                                                                            Floor




                                                                                                                                    Physical device points
                                                                                                         Seat




                                                                                            Room                                  Status, battery level, and location




                                                 Building             Floor          Room     Location          IH      IPC            Tablet               AP
                                                                                                                                                                        Device
                                        Space




                                                   G1                  06            D25       01S          ...         ...              ...                ...




                       Physical space




                              Figure 31 Spatial semantic alignment of campus data


                                                                                            72
03 Building with AI: Key Technical Features of AI Campuses



