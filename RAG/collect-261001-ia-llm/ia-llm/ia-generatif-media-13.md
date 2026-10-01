---
id: collect-261001-ia-llm/ia-llm/ia-generatif-media-13
title: "IA générative : image, vidéo, recherche"
domain: ia-llm
role: reference
task: reference
actors: ["Microsoft", "Moonshot", "OpenAI", "Perplexity"]
dates: []
keywords: ["apache", "copilot", "fine-tuning", "gpu", "kimi", "lora", "perplexity", "sol"]
source: docs/RAG/collect-261001-ia-llm/ia_generatif_media.md
source_anchor: ""
source_lines: [1183, 1316]
sha256: 5491dcec2d62be632fd9aff0297310afabbb041e47e1004f614ee1b31704c051
---

# IA générative : image, vidéo, recherche

**Démarche** :
1. **Inventaire** : lister les 12 visuels, classer en « schéma » (6) et « photo d'ambiance » (6).
2. **Choix d'outil** : FLUX local (klein 4B, Apache 2.0 — commercial OK, 0 $ marginal, confidentiel). Si pas de GPU : API BFL (FLUX.2 pro, 0,03 $/image).
3. **Série cohérente** : fixer une description verrouillée (« armoire RAL 7035, voyants verts, sol technique gris ») + seed de base, déclinée par prompt.
4. **Génération** : 12 prompts, ×2 essais → ~24 générations, ~30 min en local.
5. **Relecture critique** : vérifier chaque texte/étiquette (piège n°11), chaque incohérence électrique (un disjoncteur « à l'envers », un câble impossible).
6. **Finition** : ajouter flèches/numéros/légendes en DAO (pas dans l'IA), mention « visuels d'illustration générés par IA » en pied de doc.
7. **Archivage** : prompts + seeds + modèle dans le dossier de la doc (reproductibilité pour la v2).

**Coût** : ~0 $ (local) ou ~0,72 $ (API BFL). **Temps** : 2-3 h tout compris.

## 102. Cas n°2 : vidéo de présentation d'un équipement (onduleur)

**Besoin** : clip 30 s pour présenter un onduleur 60 kVA sur l'intranet / à un client.

**Démarche** :
1. **Storyboard** : 3 plans de 10 s (1. travelling vers l'équipement, 2. détail façade/voyants, 3. plan large du local).
2. **Source** : photographier le vrai matériel (ou un équivalent) → **image-to-video** (crédibilité maximale, piège n°12 évité).
3. **Génération** : Kling 3.0, 1080p, **sans audio natif** (on ajoutera une voix off + musique libre en montage — moins cher et contrôlable). Multi-prompt si besoin de continuité narrative.
4. **Essais** : budgéter ×3 par plan (9 générations ≈ 9 × 10 s × ~0,35-0,70 $ en crédits web Pro, ou ~1,1 $/clip en API).
5. **Montage** : assembler (DaVinci/Premiere), ajouter cartons de texte **en montage** (pas générés), voix off, mention « séquences générées par IA ».
6. **Validation** : faire relire par un collègue qui connaît l'équipement (un détail absurde grille la crédibilité).

**Coût** : < 10 $ sur Kling Pro web (crédits inclus dans les 25,99 $/mois). **Temps** : une demi-journée.

## 103. Cas n°3 : veille automatisée hebdomadaire (API Sonar + cron)

**Besoin** : chaque lundi 7h, un digest « onduleurs, batteries, normes, CVE énergie » dans ta boîte mail.

**Démarche** :
1. **Clé API** Perplexity (compte à part « veille », alerte de dépense à 5 $/mois).
2. **Script Python** (modèle `sonar`, pas `sonar-pro` : 10× moins cher pour de la veille) :

```python
import os, smtplib
from email.message import EmailMessage
from openai import OpenAI

client = OpenAI(api_key=os.environ["PERPLEXITY_API_KEY"],
                base_url="https://api.perplexity.ai")

SUJETS = [
    "nouveautés onduleurs triphasés 10-120 kVA cette semaine",
    "batteries lithium pour UPS : nouveautés et rappels produits",
    "CVE ou failles de sécurité touchant des onduleurs ou équipements d'énergie connectés",
    "évolution normes NF C 15-100, NF C 18-510, IEC 62040",
]

blocs = []
for s in SUJETS:
    r = client.chat.completions.create(
        model="sonar",   # économique ; passer à sonar-pro si la qualité suit mal
        messages=[{"role": "user",
                   "content": f"Actualités des 7 derniers jours : {s}. "
                              "Réponds en 5 puces max, chaque puce avec la source "
                              "(titre + URL + date). Si rien de notable : dis-le."}])
    blocs.append(f"## {s}\n{r.choices[0].message.content}")

msg = EmailMessage()
msg["Subject"] = "Veille énergie — hebdo"
msg["From"] = "veille@exemple.fr"
msg["To"] = "zelef@exemple.fr"          # adresse d'exemple : adapter
msg.set_content("\n\n".join(blocs))
with smtplib.SMTP("localhost") as smtp:  # ou ton relais SMTP
    smtp.send_message(msg)
```

3. **Cron** : `0 7 * * 1 /usr/bin/python3 /opt/veille/veille_hebdo.py` (adapter le chemin).
4. **Garde-fous** : timeout, log des appels, Jev en option pour scorer la pertinence avant envoi (section 65).
5. **Coût** : ~80 requêtes/mois en `sonar` ≈ **0,50 $/mois**.

**Variante sans code** : Space Perplexity avec tâche planifiée (section 34) — 0 ligne de code, inclus dans Pro.

## 104. Cas n°4 : série de visuels cohérents (identité produit)

**Besoin** : 20 visuels d'un même équipement (ex. une gamme d'onduleurs) pour un catalogue interne, sous plusieurs angles et décors, **sans** pouvoir photographier chaque configuration.

**Démarche** :
1. **Référence** : 3-4 vraies photos de l'équipement sous différents angles.
2. **Outil** : **FLUX.2 multi-références** (jusqu'à ~6-10 images) ou **Kling O1** côté vidéo — le verrouillage d'identité sans fine-tuning.
3. **Protocole** : valider UN visuel maître (angle 3/4, décor neutre), puis décliner (décor atelier, décor data center, détail façade) en gardant les mêmes références + description verrouillée.
4. **Contrôle** : passer chaque visuel dans la check-list (logo correct ? proportions ? texte lisible ?) ; rejeter sans pitié (le coût unitaire est de quelques centimes).
5. **Alternative lourde** (si l'identité doit être parfaite sur 200+ visuels) : fine-tuning LoRA sur FLUX en local — investissement initial (dizaines d'images d'entraînement, heures de GPU), rentabilisé au volume.

**Coût** : quelques dollars en API, ou 0 $ en local. **Leçon** : la cohérence se **conçoit** (références + protocole), elle ne s'obtient pas en relançant le même prompt en espérant mieux.

---

# PARTIE L — CHECK-LISTS & PENSE-BÊTE TERRAIN

## 105. Pense-bête : quel outil pour quel besoin (1 page)

```
BESOIN IMAGE
├── Beau visuel, affiche, créa ................. Midjourney V8.2
├── Doc technique, schéma, volume .............. FLUX (local klein 4B / API BFL)
├── Photoréalisme critique (portrait, produit) . Imagen 4 Ultra / Nano Banana Pro
├── Texte long lisible dans l'image ............ Ideogram 4.0 / Nano Banana Pro
├── Instructions complexes suivies au pied ..... GPT Image 2
├── Juridiquement sûr (entreprise frileuse) .... Adobe Firefly
└── Confidentiel / air-gap ..................... FLUX local (Apache 2.0)

BESOIN VIDÉO
├── Rapport qualité/prix, usage régulier ....... Kling 3.0 (Pro web)
├── Clip vitrine premium ....................... Veo 3.1 Standard
├── Régularité + suite de prod ................. Runway Gen-4.5
├── Ambiance (fumée, eau, feu) ................. Luma Ray 3.2
├── Brouillon rapide pas cher .................. Hailuo H3 / Veo Lite
├── Mouvement précis, clip 30 s ................ Seedance 2.5
└── Intégration par API au volume .............. Kling API / Veo API / fal.ai

BESOIN RECHERCHE
├── Fait technique sourcé, vite ............... Perplexity (+ ouvrir les sources)
├── Code / config .............................. Phind
├── Veille auto d'équipe ....................... Perplexity Spaces / API Sonar
├── Écosystème chinois ......................... Kimi
└── Déjà dans M365 ............................. Copilot

DÉCISION AUTO (pas de génération)
└── Classer / scorer / router / garde-fou ...... Jev (TypeSafe)
```

## 106. Check-list : avant de générer (2 min)

- [ ] Le besoin est-il bien de la **génération** (vs une vraie photo, un vrai schéma DAO) ?
- [ ] Le bon outil pour ce besoin (pense-bête section 105) ?
- [ ] Brouillon sur modèle **pas cher** d'abord ?
- [ ] Rien de **confidentiel** dans le prompt / l'image source ? (section 78)
- [ ] Aucune **personne réelle** / marque sans droit ?
- [ ] Budget : coût estimé × taux de rejet, dans le quota du mois ?
- [ ] Licence : usage commercial autorisé pour ce que je vais en faire ? (section 74)

## 107. Check-list : avant de publier un livrable IA (5 min)

