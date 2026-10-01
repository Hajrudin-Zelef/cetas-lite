---
id: collect-261001-general-networking/general-networking/cryptographie-post-quantique-l-anssi-impose-2027-2026-3
title: "OpenSSL 3.5+ intègre nativement ML-KEM, ML-DSA et SLH-DSA"
domain: general-networking
role: reference
task: reference
actors: ["CISA", "Oracle"]
dates: []
keywords: ["cyber", "mai", "valuation"]
source: docs/RAG/collect-261001-general-networking/cryptographie-post-quantique-l-anssi-impose-2027-2026.md
source_anchor: ""
source_lines: [91, 133]
sha256: 0213fd08a7e68f4dfa51ea11703127d44819772c7d2ecbb81d825001ff36de6f
---

# OpenSSL 3.5+ intègre nativement ML-KEM, ML-DSA et SLH-DSA

L’annonce française prolonge un mouvement continental. En avril 2024, la Commission européenne a publié une recommandation pour une feuille de route coordonnée de transition vers la PQC. Le Groupe de coopération NIS en a livré la première déclinaison opérationnelle en juin 2025, cosignée par l’ANSSI et disponible sur le portail de l’agence.

Cette feuille de route fixe trois échéances : fin 2026 pour les premières stratégies nationales et la sensibilisation ; fin 2030 pour la migration des cas d’usage à haut risque (eau, énergie, santé, finance, transports) ; fin 2035 pour une transition complète, dans la mesure du possible. Le message de l’ANSSI aligne donc la contrainte de certification française sur la cible européenne de 2030 pour les systèmes les plus sensibles.

Le défi reste immense. Une évaluation de l’ENISA portant sur plus de 1 350 organisations dans les 27 États membres a conclu que la plupart des acteurs européens demeurent largement impréparés. Plus frappant encore : une étude commandée par l’ANSSI en mai 2025 a montré qu’aucune des organisations sondées ne disposait alors d’un plan de transition post-quantique formalisé. Entre l’ambition réglementaire et la réalité opérationnelle, le fossé est considérable – et c’est précisément pour le combler que l’agence a choisi de brandir l’arme de la certification. Les ressources de l’ENISA sur la cryptographie offrent un point de départ aux RSSI qui cherchent à structurer leur plan.

## Impact sur le marché : éditeurs, intégrateurs, secteurs critiques

L’onde de choc se propage à toute la chaîne de valeur, et les grands équipementiers comme les investisseurs l’ont déjà intégrée : Cisco a annoncé en septembre 2026 qu’il déploierait des communications quantum-safe sur la quasi-totalité de son portefeuille de cœur de réseau d’ici décembre 2026, tandis qu’une nouvelle génération d’éditeurs spécialisés émerge, à l’image de Quantum Secure Encryption Corp, qui a lancé dès mars 2026 sa plateforme de migration PQC pour entreprises, QPA v2. Le britannique PQShield, spécialiste pur-player de la cryptographie post-quantique, a porté son financement cumulé à 63 millions de dollars, dont une levée Série B de 37 millions de dollars bouclée en avril 2026 selon Sig.ai – signe que les fonds parient sur une demande européenne durable, portée par des échéances comme celle de l’ANSSI. Pour les éditeurs français de sécurité – fabricants de HSM, de cartes à puce, de VPN souverains, de solutions de signature électronique – l’échéance 2027 impose un chantier de réécriture cryptographique immédiat. Ceux qui ont anticipé, souvent des acteurs historiques rompus aux exigences de l’ANSSI, disposent d’une avance décisive. Les autres risquent l’exclusion pure et simple du marché public, sanction commerciale plus redoutable qu’aucune amende.

La décision s’inscrit dans une stratégie nationale de long terme, mais aussi dans une course mondiale aux financements. La France avait lancé en 2021 un plan quantique doté de 1,8 milliard d’euros, et la souveraineté cryptographique en constitue un pilier ; selon SNS Insider (décembre 2025), plus de 55 pays ont désormais alloué collectivement plus de 4 milliards de dollars à des initiatives post-quantiques, preuve que Paris n’est qu’une capitale parmi beaucoup d’autres à transformer la PQC en priorité budgétaire. En liant certification et post-quantique, l’ANSSI transforme un investissement de recherche en levier industriel : elle crée une demande captive pour les technologies qu’elle a contribué à financer.

### Les secteurs en première ligne

Tous les secteurs ne sont pas logés à la même enseigne. La banque et l’assurance, qui manipulent des données financières à longue durée de vie et des signatures engageantes, figurent en tête des priorités – d’autant que les régulateurs financiers (dans le sillage de DORA) scrutent désormais la résilience opérationnelle. La santé, avec des dossiers patients devant rester confidentiels des décennies, est tout aussi exposée. Viennent ensuite l’énergie, les télécommunications et les transports, colonne vertébrale des infrastructures critiques visées par la feuille de route européenne. Pour ces filières, le chiffrement post-quantique n’est plus une option de veille technologique : c’est une échéance de conformité.

## Le défi de la migration : crypto-agilité et inventaire

Migrer vers la PQC ne se résume pas à changer une bibliothèque. La plupart des organisations ignorent où, exactement, leur chiffrement est utilisé : certificats TLS, tunnels VPN, signatures de code, clés SSH, secrets applicatifs enfouis dans du code hérité. La première étape, unanimement recommandée, est l’inventaire cryptographique – cartographier chaque usage avant de le remplacer. C’est le préalable à la « crypto-agilité », cette capacité à changer d’algorithme sans réécrire toute l’application – un concept que le NIST a lui-même formalisé en juin 2026 dans la mise à jour CSWP 39upd1, consacrée aux bonnes pratiques de crypto-agilité.

Les organisations les plus matures adoptent une approche progressive : d’abord les systèmes exposés à Internet et les données à longue durée de vie, puis les systèmes internes. Le chiffrement hybride sert de pont, permettant de déployer la PQC sans casser la compatibilité avec les systèmes classiques. Concrètement, un audit peut commencer par recenser les algorithmes en usage sur un parc de serveurs.

```
# Recenser les algorithmes de clés publiques d'un parc de certificats
for host in $(cat serveurs.txt); do
  echo "== $host =="
  echo | openssl s_client -connect "$host":443 2>/dev/null \
    | openssl x509 -noout -text \
    | grep -E "Public Key Algorithm|Signature Algorithm"
done
# Repérer les clés RSA/ECDSA à migrer en priorité (longue durée de vie)
grep -rlE "BEGIN (RSA|EC) PRIVATE KEY" /etc/ssl /home 2>/dev/null
```
Ce type d’inventaire, trivial en apparence, révèle souvent des centaines de dépendances cryptographiques oubliées. C’est là que se joue le vrai coût de la transition : non dans les algorithmes eux-mêmes, publics et gratuits, mais dans l’archéologie logicielle nécessaire pour les déployer partout. Une estimation conjointe NSA, NIST et ENISA publiée en juin 2026 chiffre d’ailleurs ce déséquilibre : les dépenses de services liées à la migration post-quantique (audit, intégration, accompagnement) représenteraient de 8 à 12 fois le montant des licences logicielles PQC elles-mêmes.

## Contexte : une France sous forte pression cyber en 2026

L’annonce de l’ANSSI ne surgit pas dans un ciel serein. L’année 2026 s’est ouverte sur une vague d’incidents qui a placé la France en tête des pays européens les plus attaqués, comme nous l’avons documenté dans notre analyse de la France, deuxième pays le plus piraté. La cyberattaque contre l’ANTS, avec ses 11,7 millions de comptes exposés, ou encore l’exploitation de failles critiques comme la vulnérabilité Oracle PeopleSoft (CVSS 9,8), rappellent que la menace « classique » n’a pas attendu l’ordinateur quantique pour frapper.

Ce contexte donne du poids à la décision. En durcissant les exigences de certification, l’ANSSI ne se contente pas d’anticiper une menace future : elle relève le niveau de sécurité de tout l’écosystème face à des attaques qui, elles, sont bien présentes. La **cryptographie post-quantique** devient l’un des chantiers d’une doctrine plus large de durcissement, aux côtés du renseignement sur les menaces et de l’hygiène cyber promue par des outils comme CrowdSec ou l’auto-hébergement de coffres de mots de passe via Vaultwarden.

## Souveraineté : le vrai enjeu derrière la technique

