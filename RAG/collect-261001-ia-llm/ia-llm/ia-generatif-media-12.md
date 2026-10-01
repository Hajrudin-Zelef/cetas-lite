---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-12
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["EU", "Google", "Perplexity"]
dates: ["2026-09-24", "2026-09-27"]
keywords: ["apache", "diffusion", "open source", "open weights", "perplexity"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [1097, 1182]
sha256: 032a9cb17877289b0f569a42785ebaf4fd35136147f10e43d9924a6955dfdfc7
---

# IA générative : image, vidéo, recherche

Kling API (30/180 j), Kling web (fin de mois), Luma (fin de mois), Midjourney (fin de mois) : partout, le non-consommé est perdu. **Symptôme** : tu paies 700 $ de pack « pour l'année » et tu en utilises 200 $. **Parade** : commencer petit, mesurer 2 mois, puis dimensionner. Voir section 71.

## 82. Piège n°2 : web ≠ API, deux caisses séparées

Chez Kling (et c'est fréquent ailleurs), l'abonnement web ne crédite pas l'API et réciproquement. **Symptôme** : « je suis Pro à 25,99 $, pourquoi l'API me dit crédit insuffisant ? » **Parade** : décider dès le départ si ton usage est manuel (web) ou programmé (API), et ne payer qu'une caisse.

## 83. Piège n°3 : les générations ratées sont facturées

Un clip avec six doigts, un schéma avec du texte en charabia : tu as payé quand même. Avec un taux de rejet de ×3 au démarrage, ton coût réel est le triple du prix facial. **Parade** : brouillons sur modèles pas chers (schnell, klein, Hailuo, Veo Lite), master sur le bon modèle ; fixer un quota d'essais par livrable.

## 84. Piège n°4 : l'audio natif qui multiplie la facture

Sur Kling, l'audio natif peut multiplier le coût par 2,5 à 5 (selon version et mode). **Symptôme** : la facture triple le mois où « on a juste coché la case son ». **Parade** : ne l'activer que si le son apporte vraiment (dialogue, ambiance) ; sinon ajouter une bande-son en montage (gratuit, contrôlable).

## 85. Piège n°5 : la 4K systématique

La 4K native (Kling 3.0 : 30 crédits/s ; Veo : 0,60 $/s) coûte 3 à 5× le 1080p. Pour de l'intranet, de la formation projetée en salle ou du web, le 1080p suffit dans 90 % des cas. **Parade** : master en 1080p, 4K uniquement pour diffusion grand écran / print vidéo.

## 86. Piège n°6 : la licence « open weights » confondue avec « open source »

FLUX.2 [dev] se télécharge librement mais **interdit l'usage commercial** sans licence BFL. **Symptôme** : des visuels de doc client générés « gratuitement » en infraction. **Parade** : klein 4B ou schnell (Apache 2.0) pour le commercial local ; lire la licence avant, pas après. Voir section 18.

## 87. Piège n°7 : le fournisseur qui ferme (leçon Sora)

Le 24/09/2026, l'API Sora a disparu avec 6 mois de préavis et **zéro remplaçant**. **Parade** : isoler chaque appel IA derrière une interface interne (une fonction `generer_video()` à toi, pas des appels Kling en dur partout), pinner les versions, exporter les assets, suivre les pages de dépréciation, toujours avoir un fournisseur B identifié.

## 88. Piège n°8 : la version qui change sous tes pieds

`jev-latest`, les snapshots de modèles, les « améliorations silencieuses » : un pipeline qui marchait peut dériver sans que ton code change. **Parade** : pinner les versions exactes (`jev-1.13.0`, `flux-2-pro`, snapshot daté), logger modèle+version à chaque génération, tester en non-régression après chaque upgrade.

## 89. Piège n°9 : les données confidentielles dans le prompt

Envoyer à une API publique : une photo d'installation client, un schéma réseau annoté, un plan de local technique. **Parade** : classification en 3 niveaux (section 78), endpoint UE quand il existe, local pour le confidentiel. Et relire ce que contient une image **avant** de l'uploader (les EXIF aussi).

## 90. Piège n°10 : Perplexity pris pour une source primaire

Une réponse bien citée n'est pas une preuve. Les modèles **résument parfois de travers**, et une citation peut pointer vers un blog qui lui-même se trompe. **Parade** : pour toute valeur d'ingénierie (tension, couple, norme), ouvrir la source primaire. Perplexity = le pointeur, pas la preuve.

## 91. Piège n°11 : le texte dans l'image qui ment

« 400 V » devenu « 4000 V », un logo approximatif, une étiquette en charabia. Dans une doc technique, c'est un **risque de sécurité**, pas un défaut esthétique. **Parade** : relire chaque texte généré caractère par caractère ; pour les schémas cotés, générer l'image **sans** texte et ajouter le texte en PAO/DAO après.

## 92. Piège n°12 : l'incohérence d'un visuel à l'autre

Le même onduleur qui change de couleur, le technicien qui change de visage entre deux clips. **Parade** : FLUX.2 multi-références, Kling O1 (identité), seeds fixes, prompts avec description verrouillée du sujet (« même armoire grise RAL 7035, voyants verts »). Pour une série, valider le « personnage » d'abord, décliner ensuite.

## 93. Piège n°13 : le watermark oublié

Sortie du plan gratuit (Kling, Luma) utilisée telle quelle dans une présentation client. **Parade** : bannir les exports gratuits des dossiers « livrables » ; nommer les dossiers `brouillon/` vs `final/` ; check-list de publication (section 80).

## 94. Piège n°14 : la personne réelle générée sans consentement

Le collègue « incrusté » dans une vidéo de formation pour la blague, le client reconnaissable dans un visuel. **Parade** : consentement écrit ou anonymisation (visages non identifiables, silhouettes). Le droit à l'image ne rigole pas, et l'EU AI Act durcit la transparence.

## 95. Piège n°15 : le style d'un artiste vivant pour du commercial

« À la manière de [artiste] » pour une affiche d'entreprise : zone grise juridique et éthique, et un style reconnaissable qui vieillit mal. **Parade** : décrire le style par ses caractéristiques (« affiche rétro, aplats, typographie bold ») plutôt que par un nom propre.

## 96. Piège n°16 : la file d'attente le jour J

Présentation client à 14h, génération lancée à 13h30 sur le plan web gratuit/standard : file de 30 min. **Parade** : générer la veille ; pour l'événementiel, plan avec file prioritaire (Kling Pro/Premier) ou API ; toujours avoir un plan B (slides statiques).

## 97. Piège n°17 : l'API sans garde-fous budgétaires

Un script de veille en boucle, un batch d'images sans limite, une clé API qui fuite sur GitHub. **Parade** : alertes de dépense dès le jour 1, quotas par clé, une clé par usage, jamais de clé dans le code (variables d'environnement / gestionnaire de secrets), rotation immédiate en cas d'exposition.

## 98. Piège n°18 : le « tout gratuit » qui coûte en temps

Enchaîner les offres gratuites (66 crédits/jour ici, 8 vidéos/mois là) pour « ne rien payer » : des heures perdues en files d'attente, watermarks, qualité brouillon. **Parade** : chiffrer ton temps (ton taux horaire × heures perdues) — 26 $/mois de Kling Pro sont rentabilisés en une heure gagnée.

## 99. Piège n°19 : l'over-reliance — l'IA comme béquille

Générer au lieu de comprendre : un schéma IA « à peu près juste » dans une procédure de consignation, une réponse Perplexity recopiée dans un rapport sans vérification. **Parade** : l'IA produit le **brouillon**, l'expert **valide**. Pour le critique (sécurité, conformité), la validation est non négociable et c'est toi qui signes.

## 100. Piège n°20 : oublier que les prix bougent

Toutes les grilles de ce guide datent du **27/09/2026**. Les fournisseurs ajustent tous les trimestres (exemples récents : restructuration Luma, packs Kling, batch -50 % chez Google). **Parade** : re-vérifier les 3-4 tarifs qui comptent pour toi **chaque trimestre** (15 min), et dater tes comparaisons internes.

---

# PARTIE K — CAS PRATIQUES COMMENTÉS

## 101. Cas n°1 : visuels pour une documentation technique

**Besoin** : illustrer une procédure « Remplacement d'un module de puissance sur onduleur 40 kVA » — 12 visuels (vues d'ensemble, détails borniers, gestes).

