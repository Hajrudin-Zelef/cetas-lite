---
id: collect-261001-general-networking/general-networking/calculer-son-alimentation-pc-1000-w-pour-rtx-5090-2026-5
title: "Exemple : RTX 5090 + Ryzen 9 9950X3D"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "gpu", "nvidia"]
source: docs/RAG/collect-261001-general-networking/calculer-son-alimentation-pc-1000-w-pour-rtx-5090-2026.md
source_anchor: ""
source_lines: [261, 275]
sha256: 3330d12cfaa6bba3ae6d04e8d8a5644c01ce7eba697fcb0e497a140ab1db5740
---

# Exemple : RTX 5090 + Ryzen 9 9950X3D

Single rail est généralement recommandé pour une configuration mono-GPU, car cela évite les erreurs de répartition de charge entre plusieurs circuits.

**Combien de temps dure une alimentation PC ?**

Entre 7 et 10 ans en usage normal pour un modèle Gold ou supérieur avec une bonne garantie constructeur, moins si elle est utilisée en permanence proche de sa charge maximale.

**Peut-on réutiliser son alimentation actuelle en changeant de GPU ?**

Cela dépend du wattage disponible et de l’âge du bloc. Recalculez systématiquement avec le TGP du nouveau GPU avant de décider, en particulier si votre alimentation a plus de 5 ans.

**Pourquoi mon alimentation 750 W ne suffit-elle pas pour une RTX 5080 avec un gros CPU ?**

Le TGP de 360 W plus un CPU haut de gamme à 230 W et la plateforme dépassent souvent 700 W de base, et la marge de sécurité de 30-40 % pousse le besoin réel au-delà de 750 W. Le passage à 850-1000 W est alors nécessaire.

Pour aller plus loin sur les cartes graphiques RTX 50 et leurs spécifications complètes, consultez la page officielle NVIDIA RTX 5090. Les seuils de rendement 80 PLUS sont détaillés sur le site de l’organisme de certification 80 PLUS. Pour l’historique de la norme ATX, la fiche ATX sur Wikipédia reste une référence utile. Tom’s Hardware publie régulièrement des tests d’alimentations PC, tandis qu’AMD détaille les spécifications de ses cartes RX 9000 sur son site officiel. Le comparateur français Les Numériques publie aussi des comparatifs de composants PC actualisés.
