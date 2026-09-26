---
id: collect-240926-storagereview/storagereview/fr-review-seven-things-to-love-about-arc-enabled-sql-managed-instances-721979cb-2
title: "fr-review-seven-things-to-love-about-arc-enabled-sql-managed-instances-721979cb"
domain: storagereview
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["compute", "cost", "intel", "license", "open source", "pricing"]
source: docs/RAG/clean_en/storagereview/fr-review-seven-things-to-love-about-arc-enabled-sql-managed-instances-721979cb.md
source_anchor: ""
source_lines: [3, 33]
sha256: cc796a59b80e4633612eeb6160583f61103f6ced5a30d227d987b16769f5be72
---

# fr-review-seven-things-to-love-about-arc-enabled-sql-managed-instances-721979cb

Following our recent article on Azure Arc-enabled managed services, we continued our exploration of the power of Azure Arc and Azure Stack HCI with DataON, a Microsoft and Intel partner. We quickly saw their benefits, and one use case stood out in particular: the Azure Arc-enabled SQL Managed Instance. This instance is a platform as a service (PaaS) that uses the latest version of SQL Server (Enterprise Edition), automatically updated and backed up. In addition, for mission-critical applications, it includes high availability.
While exploring Azure Arc-enabled SQL Managed Instance, we discovered several unique, interesting, or powerful features. These are discussed below.
Leveraging the power of hyperconverged infrastructure
Very early on, enterprises discovered the power of the Azure public cloud and the services it could provide. However, for some workloads, they need to remain on-premises for compliance reasons. Azure Stack HCI meets regulatory requirements by using the power and services offered by Azure (including Arc-enabled SQL Managed Instance), allowing these workloads to run on the company's hardware at the location of its choice.
DataON, one of the companies with which we work in partnership, was one of the first to adopt these technologies and helped us better understand them.
With Azure Arc, customers can view and manage their applications and databases consistently with a familiar set of tools and interface, regardless of where these services run, from on-premises to multicloud to the edge.
Now, each Azure Stack HCI cluster node is Arc-enabled when a cluster is registered with Azure. This means that all these powerful Azure management capabilities are available for your Azure Stack HCI nodes.
Adopting hardware-based security
Early in its development, Microsoft prioritized security when creating Azure Stack HCI, Arc, and Azure Arc-enabled SQL Managed Instance. Microsoft and Intel collaborated to provide a comprehensive security solution with Azure Stack HCI, covering the entire IT infrastructure. They also integrated Azure Arc to extend Azure-based security to hybrid and multi-cloud environments. Intel's built-in security and extensions further strengthen this solution, ensuring comprehensive protection from silicon to cloud.
Intel's security measures ensure that devices and data are trustworthy, while providing workload and encryption acceleration. This enables secure, hardware-isolated data protection and software reliability to guard against cyberthreats.
The Azure platform incorporates easily accessible and user-friendly security tools and controls. Native DevOps and Security Center controls can be customized to protect and monitor all cloud resources and all architecture levels. Microsoft developed Azure using industry-standard zero trust principles, which involve explicit verification and the assumption that a breach has occurred.
Security starts at the hardware level. Using a Secured-core server and a dashboard, available through Azure Stack HCI, enables verification and auditing of the hardware to ensure that the server meets Secured-core requirements.
Engagement with DataON (an Intel Platinum partner) ensures that the hardware foundation for an on-premises Azure Stack HCI deployment uses the latest Intel-based servers to meet Secured-core server requirements. TPM2.0, Secure Boot, Virtualization Based Security (VBS), Hypervisor-protected Code Integrity, Pre-boot DMA protection, and DRTM protection are some of the security features provided by Intel servers and verified by Azure Stack HCI.
Harnessing the power of Kubernetes
Arc-enabled SQL Managed Instance leverages Kubernetes (K8s) to host the SQL instance and provide additional management capabilities for these SQL instances. K8s is a proven technology (it has existed for about a decade) in the data center, and by using it, Microsoft capitalizes on its features and functions and on its powerful and rich ecosystem.
Azure Arc-enabled SQL Managed Instance hides the complexity of running containers through dashboards and wizards while allowing others to work directly with K8s.
Transparent and instant SQL pricing
Licensing costs for your Arc-enabled SQL Managed Instance are calculated and displayed as the instance is configured, revealing the database cost before deployment. This also allows customers to perform simulation calculations and weigh trade-offs when deciding what to deploy. For example, you can determine whether you want one, two, or three replicas for high availability or any other attribute that Azure Arc-enabled SQL Managed Instance can provide. This cost information avoids surprises at the end of the month and allows business units to configure their instances according to their budgets.
As a bonus, if you already own a SQL Server license, you can use the Azure Hybrid Benefit to save on licensing costs.
Simplifying the creation of new databases
Because Azure Arc is policy-driven, an administrator or even the end user of a database can create a new SQL managed instance using the Azure web interface. Azure Stack HCI pools all server compute and storage under its control. Thus, creating a new database involves selecting the necessary attributes, but not having to decide which individual and discrete components are used for hosting.
In just a few minutes of deployment, a highly available Arc-enabled SQL managed instance with built-in features such as automated backups, monitoring, high availability, disaster recovery, and more will be ready to use.
To use the database, Azure Arc-enabled SQL Managed Instance provides a list of connection strings for common programming languages. It is a small change, but it can save programmers a lot of frustration when trying to connect.
Migrating existing databases is a breeze
Using Microsoft's fully automated Azure Data Migration Service, moving a database to Azure Stack HCI as an Arc-enabled SQL managed instance is a breeze. Even for skilled and experienced professionals, migrating to a database can be an anxiety-inducing prospect. Microsoft created a wizard to guide users through the process, removing the stress of doing it yourself or the cost of outsourcing.
Built-in monitoring
Most often, database monitoring is an afterthought, an additional cost, or neglected because of its complexity or availability. Microsoft made a bold decision by including an open source monitoring stack that includes InfluxDB and Grafana for metrics and Elastic and Kibana for logs for its Arc-enabled SQL managed instances.
We were surprised and delighted that Microsoft decided to use well-known and easily extensible open source products for monitoring. For example, Arc provides a Grafana-compatible Azure Arc-enabled SQL Managed Instance dashboard with widgets that display key performance indicators and individual metrics.
A Grafana dashboard is also provided for hosts.
In retrospect, we should have titled this article "The seven things we liked most about Arc-enabled SQL managed instances, running on Azure Stack HCI, with Arc integration on a secured Intel-based server provided by DataON," because each of these products builds on and complements the other.
SQL Managed Instance enables easy migration or creation of a database presented and consumed as PaaS. Azure Stack HCI allows Azure Arc-enabled SQL Managed Instance and other Azure services to run on-premises. Arc allows Azure Stack HCI and Azure in the cloud to be managed from the same web interface. DataON is a valued partner of Microsoft and Intel that provides hardware to run Azure Arc-enabled SQL Managed Instance in a customer's data center, in a remote office, or at the edge. Intel-based servers provide a secure foundation for this solution.
