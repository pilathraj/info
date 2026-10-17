# System Design
## Frameworks - 5 (60 mins)
1. Clarify Requirements - (20 mins)
2. Data + API - (10 mins)
3. Architecture (15 mins)
4. Scale & Edge Cases (10 mins)
5. Wrap Up. (5 mins)

## Dimensions - 5 
1. Communications
2. Structured Thinking
3. Technical Judgment
4. Practical Depth
5. Time Management

### 1 Clarify the Requirments
- Clarify the problems first
- Don't jump into the designing anything.
- **Functional requirements**
    - What are the main use cases?
    - What is explicitly out of scope?
    - Who are the users?
    - What type of platform we want to building?
- **Non Functional requirements**
    - Daily/Monthly users?
    - Request per seconds
    - Read and write Radio?
    - Latency target 200ms or sub 100 ms?
    - Acceptable downtime?
- Does this scope and these requirements make sense before I move forward?
  
### Design Core Data & APIs
- The Model what you are storing & how clients will interact with your system.
- Data Model
      - The Main entities & Relationship
      - Primary Key, Foreign Key and Indexes
      - Which data needs strong consistency?
  - API Design
      - Main endpoints with request & response formats
      - Pagination & error handling
      - Rate limit & versioning
      - Best Practices
  - Database Choice
       - Pick the database & justify it
       - Mention alternatives considered
       - Walk through the trade-off explicitly.

