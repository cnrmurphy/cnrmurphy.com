# Conor Murphy

_Software Engineer

**Email:** [conor@cnrmurphy.com](mailto:conor@cnrmurphy.com)

**Github:** <https://www.github.com/cnrmurphy>

- - -

# Professional Experience
**SPINS**

*Software Engineer*

~03/2022 – 06/2024

* Diagnosed and resolved MongoDB performance issues including over-fetching (1M+ unnecessary documents), N+1 query patterns, and excessive frontend API calls, significantly reducing data transfer costs and eliminating production timeouts; coincided with ~30% infrastructure cost reduction
* Improved system performance by refactoring schema anti-patterns and enforcing MongoDB best practices
* Enhanced observability across microservices by implementing OpenTelemetry, enabling real-time latency monitoring and correlated logging
* Maintained and debugged Builder, a proprietary DSL compiler critical to data accuracy, enabling nutrition experts to author 1,000+ data transformation formulas compiled to JavaScript at runtime; improved developer experience by fixing error propagation and supporting hot-reloading for production systems processing 1M+ products
* Conducted comprehensive code reviews for 9-person engineering team, identifying MongoDB anti-patterns, data over-fetching issues, and performance bottlenecks; frequently served as primary reviewer during high-velocity delivery periods, ensuring architectural best practices and preventing production performance issues
* Facilitated post-acquisition integration of 7-person Product Intelligence team by developing automated documentation system (Mongoose schemas to Markdown with cross-references), establishing development environments, and providing ongoing architectural guidance for feature implementation
* Developed full-stack image processing pipeline managing 100K+ images with priority queuing for batches up to 10K images; implemented S3 streaming, automated zip/folder organization, and identified race conditions, implementing MongoDB document locking to prevent concurrent modification during downloads
* Built self-service script execution platform enabling operations team to run data processing and maintenance scripts via web interface with real-time logging; implemented local development environment with temporary BigQuery tables, eliminating UAT deployment requirements for testing
* Deployed and maintained production services on GCP (Cloud Run, Cloud Functions, BigQuery) using Terraform and Concourse; developed React/MobX admin tools and managed AWS S3 integrations for data storage

**Pinto** ~(Acquired by SPINS)

*Software Engineer*

~07/2021 – 03/2022

* Proposed and built a static analysis tool to transpile Mongoose schemas into Markdown documents for automatic domain documentation
* Developed real-time retailer data quality scoring system serving 100+ retailers; implemented middleware with tree traversal algorithms to merge grading feedback into retailer documents upon product data updates
* Built APIs and provided technical support for data engineering workflows; wrote documentation and conducted code reviews for Python-based data collection scripts, and developed ad-hoc queries for operations team

**Wakefern Food Corporation**

*Software Engineer*

~05/2017 – 07/2021

* Built and deployed message broker processing transaction data from 100+ stores using Go, gRPC, and Beanstalkd; parsed and transformed XML transaction data to DB2 for customer shopping trend analysis
* Built full-stack project management system serving 6 project managers and 100+ department staff; implemented JIRA proxy with Vue.js frontend and PostgreSQL backend to enforce PM-only project state controls and provide executive-level reporting with data rollups for department director
* Created an ETL pipeline in Python to ingest and clean IBM Weather API data
* Developed custom JavaScript SPA framework with routing, state management, and component system using regex-based template parsing; implemented custom directive syntax ($$for loops, dynamic replacements) for internal applications in restricted open-source environment
* Implemented shop-from-home feature across 100+ Shoprite stores using geofencing technology; developed cross-platform location detection with unified iOS/Android permission handling and state management for automatic store notification upon customer arrival

- - -

# Personal Projects

**Algorithmic Trading System** ~(Go, Python, MongoDB)

* Event-driven low-latency trading engine built with components like Order Manager and Execution Manager. Uses Go channels to implement Actor-like model to minimize contention. Includes
  custom REST/WebSocket clients, supports 80+ active orders and 8 concurrent markets, and features execution algorithms like TWAP.

**VoilaDB** ~(Go)

<https://github.com/cnrmurphy/voila>

* A key-value store using fixed-page layout and offset-based navigation for performance and simplicity.

**Mail Server** ~(Rust)

<https://github.com/cnrmurphy/pop3-rs>

* A mail server fully compliant with the POP3 protocol. It uses MailDir as the underlying storage layout, Sled to store users, and Argon2 to secure passwords.

**B+ Tree** ~(Node.js)

<https://github.com/cnrmurphy/B-Plus-Tree>

* B+ Tree with binary search for key placement, internal node splitting, and efficient insertion logic.

**Go-HashMap** (Go)

<https://github.com/cnrmurphy/go-hashmap>

* Basic hashmap using FNV-1a hashing and bitmasking to derive bucket indices.

**Sporcle “US City Triplets” Solver** ~(Rust)

<https://github.com/cnrmurphy/us-city-triplets-solver>

* A simple Rust program to help solve the Sporcle quiz “US City Triplets” utilizing web scraping to retrieve required inputs and file parsing/scanning to identify potential solutions based
  on those inputs

**FTX-US.js** ~(Node.js)

<https://github.com/cnrmurphy/ftx-us-js>

* API client for the FTX Exchange supporting market data access, trading, and asset transfers.
