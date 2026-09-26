---
id: collect-240926-storagereview/storagereview/fr-review-after-the-pi-record-serving-a-130-tb-dataset-with-backblaze-b2-c0672672-3
title: "fr-review-after-the-pi-record-serving-a-130-tb-dataset-with-backblaze-b2-c0672672"
domain: storagereview
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["aws", "compute", "research", "throughput"]
source: docs/RAG/clean_en/storagereview/fr-review-after-the-pi-record-serving-a-130-tb-dataset-with-backblaze-b2-c0672672.md
source_anchor: ""
source_lines: [25, 37]
sha256: 4672debbfaea22c376b9bab23829a9cbb392c1371a02aaa88d2d60ad1bd4ce88
---

# fr-review-after-the-pi-record-serving-a-130-tb-dataset-with-backblaze-b2-c0672672

Backblaze created the pi-314-trillion bucket, which contains all 628 files and whose confirmed size is 132,210.5 GB. This bucket is configured as private, retains all file versions, and is accessible via the S3-compatible endpoint at s3.us-west-004.backblazeb2.com. Object storage simplifies the management of a dataset of this magnitude. Each file is individually addressable, the complete listing can be retrieved programmatically, and there is no file system hierarchy or volume limit to work around.
Accessing the dataset
Researchers are asking us to access this data; in fact, a project is already underway. Michael Kleber is a senior software engineer at Google, but he has been studying the digits of pi since earning his doctorate in mathematics in 1999. Mathematicians consider pi to be a normal number , so it is reasonable to ask: "Among the 10^d sequences of d digits, which one takes the longest to appear in pi, and how many digits are needed?" Kleber conducted the research up to d = 7, and when Fabrice Bellard computed 2.7 trillion digits of pi in 2009, Kleber encouraged him to extend his research to d = 11, the limit of what was possible at the time. "With 314 trillion random digits, there is about a 79% chance of finding all sequences of length 13," Kleber explains, "so I hope we get lucky!" We expect this to happen much more often now that Backblaze has made the data accessible to everyone.
The PI dataset is hosted on Backblaze B2 and can be downloaded by anyone who wishes to use it. Access is via a request link to obtain credentials or download instructions to retrieve the files. Backblaze will host the dataset to ensure its availability for research and verification for the duration of the hosting.
Download options
Users can retrieve individual files or the entire 130 TB of data, depending on their needs. The bucket is structured so that each object can be accessed and downloaded directly, without having to retrieve the entire dataset. For those who wish to retrieve everything, a full bucket sync can be performed using the tools described below. This option requires 135 TB of free space.
Recommended tools
- Rclone is the recommended tool for accessing the dataset. It integrates seamlessly with Backblaze B2 and allows users to adapt the download process to their bandwidth needs.
- S3-compatible APIThus, any S3-compatible download tool can be used to retrieve the data. The only condition is that this tool allows the default S3 endpoint URL to be modified to point to the B2 endpoint rather than AWS.
A concrete example of hybrid infrastructure
Our computation of Pi to 314 trillion digits perfectly illustrates what hybrid infrastructure looks like in practice. The computation was entirely executed on a single Dell PowerEdge R7725 server in the StorageReview lab. However, because the system needed to be reassigned to other projects and tasks once the computation was complete, permanently keeping 130 TB of results within the lab was not a viable solution.
The desire to make the data available to those who wish to use it, whether for rigorous scientific work or simple curiosity, is very real. However, hosting a dataset of this size within the lab quickly becomes a significant burden on operations. Bandwidth is saturated, infrastructure is constantly stressed, and the lab's daily tasks are disrupted by every download request.
A solution like Backblaze B2 eliminates all of these obstacles. The data resides in cloud infrastructure designed specifically for data volumes of this magnitude, with scalable throughput, multiple redundancy points to ensure integrity, and the security and operational expertise specific to enterprise storage platforms. The compute resources were hosted on-premises because the project required it. The storage is in the cloud because that is simply the most suitable solution for what comes next.
