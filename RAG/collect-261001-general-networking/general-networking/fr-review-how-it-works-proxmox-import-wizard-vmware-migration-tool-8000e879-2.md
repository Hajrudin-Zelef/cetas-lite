---
id: collect-261001-general-networking/general-networking/fr-review-how-it-works-proxmox-import-wizard-vmware-migration-tool-8000e879-2
title: "fr-review-how-it-works-proxmox-import-wizard-vmware-migration-tool-8000e879"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fr-review-how-it-works-proxmox-import-wizard-vmware-migration-tool-8000e879.md
source_anchor: ""
source_lines: [43, 50]
sha256: c14bb0fd10b09325eeae46ca7b0b471aab3f1d9107ac18a4f6e575a6ea5db8ef
---

# fr-review-how-it-works-proxmox-import-wizard-vmware-migration-tool-8000e879

- Détachez, puis double-cliquez pour rattacher le disque de démarrage et tout autre disque que vous souhaitez utiliser comme VirtIO SCSI/Block. Assurez-vous de modifier l'option Bus/Périphérique sur SCSI.
- Accédez à l'onglet Options et sélectionnez le lecteur SCSI comme option de démarrage.
- La VM peut être démarrée et devrait fonctionner normalement. Si la machine virtuelle ne démarre pas, le disque SCSI peut être détaché et reconnecté en tant qu'IDE ou SATA pour refaire la procédure de disque factice et réessayer.
Ces pilotes peuvent également être installés lors de l'installation de Windows pour les nouvelles machines virtuelles. La documentation Proxmox est disponible ici.
Conclusion
Même si de nombreux clients VMware savent que Proxmox ne remplace pas complètement ESXi, beaucoup réalisent que des charges de travail spécifiques pourraient y migrer pour réduire les coûts de licence. Le plus grand obstacle pour de nombreux centres de données virtualisés est le manque de portabilité, ce qui signifie qu'une fois que vous vous engagez dans une plate-forme d'hyperviseur, il devient difficile de s'en éloigner. Cependant, l'entreprise fait d'énormes progrès avec des outils tels que l'assistant d'importation Proxmox, qui allégeront le fardeau des administrateurs informatiques.
Certains outils et paramètres spécifiques à l'hyperviseur peuvent rendre la migration des machines virtuelles un peu complexe. Toutefois, le processus devrait se simplifier et être mieux documenté à mesure que davantage d'entreprises l'adopteront. D'autres options que l'importation automatique peuvent être plus adaptées à votre environnement, mais nous pensons que cette méthode restera la plus utilisée. Vous trouverez des informations sur les alternatives dans la documentation Proxmox.
D’autres hyperviseurs faciliteront probablement le passage de VMware aux clients potentiels. Il existe déjà une documentation pour le faire avec Hyper-V.
