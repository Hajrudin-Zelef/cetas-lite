---
id: collect-260926-mikrotik/mikrotik/questions-1014226-mikrotik-set-scan-list-in-ros-to-only-use-2-4-ghz-non-overlapp-e4b3ffce
title: "questions-1014226-mikrotik-set-scan-list-in-ros-to-only-use-2-4-ghz-non-overlapp-e4b3ffce"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-1014226-mikrotik-set-scan-list-in-ros-to-only-use-2-4-ghz-non-overlapp-e4b3ffce.md
source_anchor: ""
source_lines: [1, 14]
sha256: d9a27d1ee17dc3ab3d816f2f6b10a916114d9436d368c187aa37fc5b3e93be62
---

# questions-1014226-mikrotik-set-scan-list-in-ros-to-only-use-2-4-ghz-non-overlapp-e4b3ffce

How do I set the scan-list in MikroTik RouterOS to use only 2.4 GHz non-overlapping channel frequencies?
                    
                        Add a comment
                    
                 | 
            
                
            
        
         
    1 Answer 1
To set the scan-list to ONLY use 2.4 GHz non-overlapping channel 1,6 & 11 in MikroTik RouterOS (currently v6.46.5):
/interface wireless set scan-list=2412,2437,2462 numbers=wlan1
Obviously change wlan1 to the correct value as necessary for the radio you're working with in your MikroTik-
