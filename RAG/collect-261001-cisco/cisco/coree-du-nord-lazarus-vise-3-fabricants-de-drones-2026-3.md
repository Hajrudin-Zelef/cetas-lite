---
id: collect-261001-cisco/cisco/coree-du-nord-lazarus-vise-3-fabricants-de-drones-2026-3
title: "coree-du-nord-lazarus-vise-3-fabricants-de-drones-2026"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["attribution", "open source", "research"]
source: docs/RAG/collect-261001-cisco/coree-du-nord-lazarus-vise-3-fabricants-de-drones-2026.md
source_anchor: ""
source_lines: [109, 173]
sha256: d4f448b24370e18dfb69c573a37fb7c594aaf6d7f3571a4f293400612d9f6c75
---

# coree-du-nord-lazarus-vise-3-fabricants-de-drones-2026

Le texte élargit considérablement le nombre d’entités soumises à des obligations de sécurité et de notification d’incidents, un périmètre qui inclut désormais de nombreux sous-traitants industriels, y compris dans la défense. Sur le papier, une entreprise victime d’une campagne comme DreamJob devrait aujourd’hui la signaler à son autorité nationale compétente. Reste que la mise en œuvre demeure inégale selon les pays, comme le montre le renvoi de la France et de trois autres pays devant la Cour de justice de l’Union européenne pour retard de transposition.

## Le silence relatif des autorités européennes

Un point frappe à la lecture du dossier public : neuf mois après la révélation d’ESET, aucune agence européenne, ni l’ENISA ni un CERT national, n’a publié d’alerte officielle spécifiquement dédiée à cette campagne contre le secteur des drones. La réaction est venue du secteur privé, pas des institutions. C’est un schéma courant en matière de cyberespionnage industriel : les victimes, souvent des PME, n’ont ni l’obligation ni toujours l’intérêt commercial à communiquer publiquement sur une intrusion qui touche à leur propriété intellectuelle.

Cette discrétion contraste avec la communication beaucoup plus musclée qui entoure généralement les rançongiciels, où la pression d’un compte à rebours et d’une fuite de données publique pousse les victimes à réagir vite. L’espionnage industriel, lui, reste largement invisible du grand public tant qu’un chercheur comme Peter Kálnai ne le documente pas noir sur blanc. Pour les rédacteurs de la directive NIS2, c’est précisément l’angle mort que le texte tente de combler en imposant un devoir de notification, même en l’absence de rançon ou de fuite spectaculaire.

## Comment les entreprises de défense peuvent se protéger

Face à une attaque qui repose avant tout sur l’ingénierie sociale, les correctifs techniques ne suffisent pas. ESET a publié une série d’indicateurs de compromission (IoC) permettant aux équipes de sécurité de vérifier si leurs systèmes ont été exposés à cette campagne.

```
Catégories d'indicateurs de compromission publiés par ESET Research :
- Hachages SHA-1 des fichiers malveillants (leurres, chargeurs, charge utile ScoringMathTea)
- Domaines de commande et contrôle (C2)
- Adresses IP associées à l'infrastructure d'attaque
- Noms des projets open source détournés et hébergés sur GitHub
Référence : blog technique ESET Research, "Gotta fly: Lazarus targets the UAV sector"
```
Au-delà de la détection technique, les recommandations restent classiques mais trop souvent négligées : vérification systématique de l’identité des recruteurs avant tout échange de fichiers, formation régulière du personnel à l’ingénierie sociale, séparation stricte des postes ayant accès aux données sensibles, et surveillance comportementale des postes de travail via des outils de détection et réponse (EDR). Aucune de ces mesures n’est coûteuse à l’échelle d’un groupe de défense. Beaucoup le sont, en revanche, pour une PME sous-traitante qui n’a ni RSSI dédié ni budget de sécurité conséquent.

## Cinq prévisions pour la cybersécurité de la défense européenne

Sur la base des éléments rassemblés par ESET Research et du schéma habituel des campagnes Lazarus, plusieurs évolutions semblent probables dans les prochains mois.

- **Un ciblage prolongé.** Tant que la guerre en Ukraine se poursuit et que Pyongyang approfondit sa coopération militaire avec Moscou, Lazarus devrait continuer à cibler les sous-traitants de défense européens, en particulier dans le secteur des drones.
- **Une extension géographique.** Après l’Europe centrale et du Sud-Est, d’autres marchés producteurs de systèmes UAV pourraient devenir des cibles, en cohérence avec le schéma d’expansion sectorielle habituel de l’opération DreamJob.
- **Davantage de signalements sous NIS2.** L’élargissement du périmètre de la directive devrait faire remonter un nombre croissant de campagnes similaires dans les statistiques officielles des autorités nationales.
- **De nouvelles sanctions.** Si l’attribution à Lazarus s’étend à d’autres victimes confirmées, une nouvelle vague de désignations, à l’image de celles déjà prononcées par l’OFAC, est plausible.
- **Un durcissement des procédures RH.** Les entreprises de défense devraient renforcer la vérification des offres d’emploi non sollicitées, ce type d’attaque reposant presque entièrement sur la crédulité humaine plutôt que sur une faille technique.

## Questions fréquentes

**Qu’est-ce que l’opération DreamJob ?**

Il s’agit d’une série de campagnes de cyberespionnage menées par le groupe Lazarus, qui utilise de fausses offres d’emploi prestigieuses pour infiltrer des organisations, principalement dans les secteurs de l’aérospatiale et de la défense.

**Qui est le groupe Lazarus ?**

Lazarus, également appelé HIDDEN COBRA, est un groupe de menace persistante avancée aligné sur les intérêts de la Corée du Nord, actif depuis au moins 2009 selon ESET (2007 selon Group-IB), connu pour ses opérations d’espionnage, de sabotage et de cybercriminalité financière.

**Quelles entreprises ont été touchées par la campagne visant les drones ?**

ESET Research indique que trois entreprises situées en Europe centrale et du Sud-Est ont été visées successivement entre mars et octobre 2025. Leurs identités n’ont pas été rendues publiques, mais elles fabriquent des équipements militaires, dont des drones, déployés en Ukraine.

**Comment fonctionne le malware ScoringMathTea ?**

C’est un cheval de Troie d’accès à distance capable d’exécuter une quarantaine de commandes : manipulation de fichiers et de processus, collecte d’informations système, ouverture de connexions réseau et téléchargement de nouvelles charges depuis un serveur de commande et contrôle.

**Pourquoi la Corée du Nord cible-t-elle les fabricants de drones ?**

Pyongyang investit massivement dans ses propres capacités de fabrication de drones militaires et s’appuie largement sur la rétro-ingénierie et le vol de propriété intellectuelle pour combler son retard technologique, selon les chercheurs d’ESET.

**Cette campagne est-elle liée à la guerre en Ukraine ?**

ESET souligne que les entreprises visées produisent du matériel militaire utilisé en Ukraine dans le cadre de l’aide européenne, et que la campagne coïncide avec le déploiement de troupes nord-coréennes en Russie, dans la région de Koursk.

**Comment les entreprises peuvent-elles se protéger contre ce type d’attaque ?**

Les recommandations de base restent les plus efficaces : vérification systématique des recruteurs et des offres d’emploi non sollicitées, sensibilisation du personnel à l’ingénierie sociale, détection comportementale sur les postes de travail, et cloisonnement des systèmes sensibles.

**La France et l’Europe occidentale sont-elles concernées ?**

Les victimes identifiées se situent en Europe centrale et du Sud-Est, mais la France dispose d’une base industrielle de défense et de drones significative et reste soumise aux mêmes obligations de vigilance au titre de la directive NIS2. Le schéma d’expansion sectorielle habituel de Lazarus n’exclut pas de futures cibles plus à l’ouest.
