## Postgres SQL
 ```
  cd compose
  docker-compose up -d
 ```
 ### pg admin
 1. localhost:5050
    * Add a new server in PgAdmin:
       - Host name/address postgres
       - Port 5432
       - Username as POSTGRES_USER, by default: postgres
       - Password as POSTGRES_PASSWORD, by default changeme
    * Click Database > Query Tool > Select Version() > Run
       - "PostgreSQL 13.2 (Debian 13.2-1.pgdg100+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 8.3.0-6) 8.3.0, 64-bit"
 3. docker exec -it ae8f6b2bbe4b bash
### What is Postgres Sql?
  - RDBMS
  - Advanced, enterprise-class, Open source Database.
  - Support both Relational(SQL) and Non-Relational(JSON) Queryies.
  - Primary database for many
     * Web application
     * Mobile Application
     * Data analytics Application

### Feature highlights
   - User defined types
   - Table inheritance
   - Sophisticated Locking mechanism 
   - Foreign Key referential integrity 
   - Views, rules, and subquery
   - Nested transcations (savepoint)
   - Multi-version concurrency control(MVCC)
   - Asynchronous Replication
   - Tablespaces
   - Point in time recovery
   - Native Microsoft windows server version. 
