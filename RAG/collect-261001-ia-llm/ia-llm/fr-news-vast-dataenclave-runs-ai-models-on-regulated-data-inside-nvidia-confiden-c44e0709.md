---
id: collect-261001-ia-llm/ia-llm/fr-news-vast-dataenclave-runs-ai-models-on-regulated-data-inside-nvidia-confiden-c44e0709
title: "fr-news-vast-dataenclave-runs-ai-models-on-regulated-data-inside-nvidia-confiden-c44e0709"
domain: ia-llm
role: reference
task: reference
actors: ["Cohere", "Nscale", "Nvidia"]
dates: []
keywords: ["nvidia", "agents", "blackwell", "cohere", "gpu", "nvlink", "rubin"]
source: docs/RAG/collect-261001-ia-llm/fr-news-vast-dataenclave-runs-ai-models-on-regulated-data-inside-nvidia-confiden-c44e0709.md
source_anchor: ""
source_lines: [1, 8]
sha256: 6f43a8ce9c931198ad2fb0faabb5642970c5e10f41e9d26716587b653b8e9f5c
---

# fr-news-vast-dataenclave-runs-ai-models-on-regulated-data-inside-nvidia-confiden-c44e0709

VAST Data a présenté en avant-première VAST DataEnclave, une fonctionnalité d'IA confidentielle de VAST DataEngine basée sur NVIDIA Confidential Computing. Cette solution permet d'exécuter des modèles d'IA propriétaires et ouverts sur des données d'entreprise au sein d'environnements contrôlés par le client : centres de données sur site, clouds souverains et sites totalement isolés du réseau. L'environnement d'exécution associe une exécution isolée matériellement à une attestation cryptographique, garantissant ainsi le traitement des pondérations des modèles et des ensembles de données réglementés tandis que les clés de chaque partie restent confinées à son propre domaine de confiance. Les opérateurs et administrateurs d'infrastructure ne peuvent donc y accéder.
Attestation cryptographique et protection en mémoire
VAST DataEnclave s'intègre à la technologie NVIDIA Confidential Computing de troisième génération sur les architectures Hopper, Blackwell et Rubin . Le système établit des environnements matériellement isolés pour les machines virtuelles et les conteneurs confidentiels, en chiffrant la mémoire du système invité, la mémoire GPU et le trafic NVLink inter-GPU. Les pipelines et modèles de données actifs restent isolés des administrateurs de la plateforme hôte, des opérateurs d'infrastructure et des charges de travail mutualisées adjacentes partageant le même matériel.
L'environnement d'exécution utilise un processus d'attestation de vérification avant déchiffrement qui valide cryptographiquement l'environnement d'exécution de confiance et le matériel GPU NVIDIA avant de divulguer les clés de déchiffrement. Ceci empêche la divulgation en clair des pondérations ou des enregistrements sources jusqu'à ce que l'enclave vérifie l'intégrité du système et l'application des politiques. La gestion des clés s'effectue via des intégrations BYOK (Bring Your Own Key Management System), conservant les clés de données client et les pondérations des modèles fournisseur dans des domaines de confiance distincts et indépendants.
Déploiements souverains, pistes d'audit et isolation des agents
DataEnclave prend en charge les environnements connectés au réseau ainsi que les centres de données totalement isolés. Les déploiements utilisent des services d'attestation basés sur la plateforme ouverte CNCF Trustee, ou l'infrastructure Confidential AI de Fortanix grâce à un partenariat pour une IA pleinement souveraine. Les événements d'attestation, les libérations de clés cryptographiques et les actions du cycle de vie de l'enclave sont consignés dans un journal d'audit infalsifiable et interrogeable de la base de données VAST.
Le même environnement d'exécution sécurisé DataEngine fournit des environnements d'exécution isolés pour les agents d'IA via VAST AgentEngine, appliquant une politique sur les données, les systèmes et les outils auxquels les agents peuvent accéder ainsi que sur les actions qu'ils peuvent entreprendre.
VAST a annoncé la préversion de DataEnclave en collaboration avec les concepteurs de modèles Cohere, CrowdStrike, Deepgram, Factory, Fundamental et TwelveLabs, les fournisseurs de cloud IA BUZZ HPC, Nscale et Sharon AI, ainsi que les partenaires matériels Cisco et Supermicro, aux côtés de Fortanix et NVIDIA. Cette préversion intervient trois semaines après l'intégration de CrowdStrike à VAST , CrowdStrike figurant désormais parmi les partenaires de DataEnclave en tant que concepteur de modèles. La commercialisation de DataEnclave est prévue pour le premier trimestre 2027 via VAST Data et ses partenaires OEM participants, dont Cisco et Supermicro.
