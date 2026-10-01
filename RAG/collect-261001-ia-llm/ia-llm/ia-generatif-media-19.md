---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-19
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "gpu", "lora", "sol"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [1940, 2036]
sha256: 526dfd1867c65de0e2d1aca59abfddf62bcddb4b0e39f68d0c1394420b90f601
---

# IA générative : image, vidéo, recherche

Si tu dois produire **des dizaines de visuels** d'un même équipement (gamme d'onduleurs, armoires types) :

1. Constituer un **dataset** : 20-50 photos réelles de l'équipement (angles variés, bonne lumière, arrière-plans variés).
2. Légender chaque image (captioning : description courte et factuelle).
3. Entraîner un **LoRA** (outils : kohya_sd, ai-toolkit — quelques heures sur RTX 4090).
4. Utiliser : charger le LoRA dans ComfyUI (`Load LoRA`, force 0,7-1,0) + mot-déclencheur dans le prompt.

Coût : une demi-journée de setup + électricité. Gain : une **identité visuelle verrouillée** sur des centaines de générations, sans payer de multi-références à chaque appel. Pour < 20 visuels, le multi-référence FLUX.2 suffit (section 104).

## 174. ComfyUI : hygiène d'exploitation

- **Versionner les workflows** (git) : un workflow = un fichier JSON versionné avec la doc qu'il illustre.
- **Épingler les versions** : ComfyUI + custom nodes + poids — noter les versions dans un `ENVIRONNEMENT.md` ; une mise à jour peut casser un workflow (même leçon que section 88).
- **File et concurrence** : une génération à la fois par GPU ; pour du batch nocturne, une file (script section 170) + `--lowvram` si besoin.
- **Sauvegarde des poids** : 70-90 Go — les sauvegarder (disque externe), re-télécharger à chaque fois est une perte de temps.
- **Sécurité** : ComfyUI écoute en local par défaut (`127.0.0.1`) — ne pas l'exposer tel quel sur le réseau sans authentification (l'API permet l'exécution de code via custom nodes).

## 175. Dix prompts image commentés (copiables, à adapter)

**P1 — Photo documentaire d'armoire (FLUX, doc technique)**
> « photographie technique d'une armoire électrique industrielle ouverte, borniers à vis alignés sur rail DIN, repérage par étiquettes, câbles gris et bleus en goulottes, lumière neutre d'atelier, cadrage frontal, netteté maximale, style documentation constructeur »
*Pourquoi ça marche : sujet précis + environnement + lumière + cadrage + style. Le « style documentation constructeur » verrouille le registre.*

**P2 — Salle serveurs (ambiance)**
> « travelling photographique : allée centrale d'un data center, rangées de baies 19 pouces noires avec LEDs bleues et vertes, sol technique gris clair, plafond avec chemins de câbles, lumière froide, profondeur de champ, photoréaliste, 35mm »
*Le détail « 35mm » et la lumière guident le rendu photo ; les LEDs donnent la vie.*

**P3 — Schéma de principe (style épuré)**
> « illustration technique épurée en isométrique d'un local énergie : onduleur 40 kVA, armoire TGBT, climatisation de précision, flèches de flux d'air bleu et rouge, fond blanc cassé, lignes fines bleu marine, sans texte, style manuel d'ingénierie »
*« sans texte » : on ajoute les légendes en DAO après (piège n°11).*

**P4 — Détail macro connectique**
> « macrophotographie d'un bornier de puissance avec câbles 25mm² sertis, cosses tubulaires, repérage jaune-vert de la terre, métal cuivré net, profondeur de champ courte, lumière latérale d'atelier »
*Le macro cache la complexité globale : parfait pour illustrer un geste précis.*

**P5 — Avant/après réaménagement (Kontext, édition)**
> Image source : photo réelle du local. Instruction : « ajoute une deuxième armoire électrique grise identique à droite de l'existante, conserve le sol, la lumière et la perspective »
*Kontext excelle là où la génération from-scratch dériverait.*

**P6 — Visuel formation sécurité**
> « illustration pédagogique : technicien avec EPI (gants isolants, visière) devant une armoire consignée, cadenas et pancarte de consignation, style affiche de prévention, couleurs franches, composition lisible de loin »
*Penser « lisible de loin » : une affiche de sécurité se lit à 3 mètres.*

**P7 — Photo produit catalogue (Nano Banana Pro / Imagen Ultra)**
> « photographie produit studio d'un onduleur tour 10 kVA noir, façade avec afficheur LCD bleu, fond gris clair dégradé, lumière douce à 45 degrés, reflet discret au sol, ultra net, style catalogue constructeur »
*Pour le premium, le modèle premium : la différence se voit sur les matières.*

**P8 — Affiche avec titre (Ideogram)**
> « affiche A3 portrait pour une journée portes ouvertes 'ÉNERGIE & MAINTENANCE', titre en grandes lettres blanches en haut, illustration d'un technicien devant des baies serveurs au centre, bandeau bleu marine en bas avec la date, style corporate moderne »
*Le texte long : Ideogram ou Nano Banana Pro ; toujours relire le titre généré.*

**P9 — Pictogramme / icône (GPT Image)**
> « pictogramme minimaliste en ligne bleue sur fond blanc : onduleur avec éclair, style flat design, traits de 3px, aucune ombre, aucun texte »
*GPT Image suit bien les contraintes « aucun/aucune ».*

**P10 — Moodboard projet (Midjourney)**
> « série de 4 visuels d'ambiance pour un data center éco-responsable : lumière naturelle, végétalisation discrète, baies blanches, atmosphère calme et lumineuse --style raw »
*Midjourney pour l'intention esthétique ; `--style raw` pour calmer l'embellissement.*

## 176. Dix prompts vidéo commentés (copiables, à adapter)

**V1 — Présentation équipement (Kling I2V, 10 s)**
> Photo réelle de l'onduleur + « léger travelling avant vers la façade, les voyants verts clignotent doucement, reflet subtil sur le sol technique, 10 secondes, lumière froide stable »
*L'I2V depuis le réel : la crédibilité est déjà là, l'IA n'ajoute que le mouvement.*

**V2 — Geste technique (Kling T2V, 5 s)**
> « gros plan sur des mains gantées qui serrent un bornier avec un tournevis dynamométrique, geste lent et précis, atelier lumineux, 5 secondes, photoréaliste »
*Une action, un plan, 5 secondes : la règle d'or.*

**V3 — Visite de site (multi-prompt Kling, 3×5 s)**
> S1 : « plan large du local technique, porte qui s'ouvre, lumière qui s'allume » / S2 : « travelling le long des baies, LEDs qui défilent » / S3 : « arrêt sur l'onduleur, zoom lent sur l'afficheur »
*Le multi-prompt raconte ; garder « même local, même lumière » dans chaque segment.*

**V4 — Ambiance data center (Luma, atmosphérique)**
> « allée de data center dans une légère brume froide, nappes de lumière bleue, mouvement d'air subtil, caméra qui avance lentement, cinématographique »
*Le créneau Luma : l'atmosphère. Ne pas lui demander un geste précis.*

**V5 — Clip vitrine premium (Veo 3.1 Standard, 8 s)**
> « lever de soleil sur un site industriel, travelling aérien vers le bâtiment technique, reflets dorés sur les façades, musique non incluse, cinématographique, 8 secondes »
*8 secondes à 3,20 $ : on ne les dépense que pour du vitrine.*

**V6 — Brouillon storyboard (Hailuo H3, 6 s)**
> « schéma animé simplifié : flux d'énergie du réseau vers l'onduleur puis vers les baies, flèches lumineuses, fond sombre, style motion design »
*Hailuo pour valider l'idée à 0,70 $ avant de masteriser ailleurs.*

**V7 — Mouvement précis (Seedance 2.5, 10 s)**
> « technicien qui déroule un câble depuis un touret, geste ample et régulier, caméra fixe en plan moyen, atelier, 10 secondes »
*Seedance quand le geste doit être juste, pas juste « beau ».*

**V8 — Détail produit e-commerce (Kling, motion transfer)**
> Vidéo de référence : rotation produit sur plateau. Sujet : ton onduleur (photo). « applique le mouvement de rotation de la vidéo de référence au produit de l'image »
*Le motion transfer : le mouvement pro sans le studio.*

**V9 — Sensibilisation sécurité (Kling multi-prompt, 2×5 s)**
> S1 : « technicien qui pose un cadenas de consignation sur un disjoncteur, geste net » / S2 : « pancarte 'NE PAS MANŒUVRER' accrochée, zoom lent »
*Deux gestes, deux plans : la sécurité se montre, elle ne se raconte pas.*

