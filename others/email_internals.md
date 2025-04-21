# Email Address Internals and Validations
## Email Address Internals
- An email address consists of 4 parts
    - **UserName/Local Part:**    ***abc***@auditcue.com
    - **@ Symbol** 
    - **Domain:**   abc@***auditcue***.com
    - **Top Level Domain:** abc@auditcue.***com***
  - #### Rough Overview on these Internals:
      - The left segment of the email address, which is typically unique, is the username/local Part.
      - Restrictions include allowing alphanumeric characters, underscores, and dots, with specific length constraints. The specific length constraints for email addresses can vary depending on the email service provider and their implementation. However, a common standard is defined by the Internet Engineering Task Force (IETF) in RFC 5321 and RFC 5322. 
      - According to these standards:
                - The overall length of an email address should not exceed 320 characters.
                - The local part (username) should not be longer than 64 characters.
                - The domain part (including the "@" symbol) should not be longer than 255 characters.
      - **Even though** these standards seem to be widely accepted, some email providers or systems may have their own specific limitations or variations.
      - While email addresses are ***not case-sensitive***, some systems may treat them as such. But most of the popular mail providers consider them as case-sensitive.
      - Domain names usually consist of alphanumeric characters (A-Z, 0-9), hyphens (-), and periods (.) for separating subdomains. Spaces and special characters are generally not allowed.
      - Hyphens are allowed in domain names, but they cannot be used consecutively and cannot be placed at the beginning or end of the domain name. e.g., - -auditcue- -
      - Obviously, Domain names must adhere to the Domain Name System (DNS) format, like it should have a hierarchical structure with subdomains (if exists), each separated by a period. The rightmost part of the domain is the top-level domain (TLD).
      - [_I cited the point below from Wikipedia, which is worth noting. **Click** on this link to read the full article_](https://en.wikipedia.org/wiki/Email_address) If quoted, it may contain **Space**, Horizontal Tab (HT), any ASCII graphic except Backslash and Quote and a quoted-pair consisting of a Backslash followed by HT, Space or any ASCII graphic; it may also be split between lines anywhere that HT or Space appears. In contrast to unquoted local-parts, the addresses ".John.Doe"@example.com, "John.Doe."@example.com and "John..Doe"@example.com are allowed. **Space** and special characters "(),:;<>@[\] are allowed with restrictions (they are only allowed inside a quoted string, as described in the paragraph below, and in that quoted string, any backslash or double-quote must be preceded once by a backslash); Comments are allowed with parentheses at either end of the local-part; e.g., john.smith(comment)@example.com and (comment)john.smith@example.com are both equivalent to john.smith@example.com.
## Email Address Validations:
The problem we have is in email case sensitivity validation. Here are the few ways to solve this problem:
1. ### SQL COLLATE No case in create table:
        
   - In this approach, we change the create statement column from email NOT NULL to email COLLATE NO CASE NOT NULL.
       **Pros:**
            - Since we add them up at the DB level, it doesn't affect indexing. While indexing, the indexes will be constructed in such a way that it considers case insensitivity as well. So, faster querying.
        **Cons:**
            - DB change
            - No idea if COLLATE NO CASE is possible if we migrate to another DBMS
2. ### Adding a new column in the DB:

   - In this approach, we add a new column in the DB which will always be in lowercase, and our original column can store whatever our customer enters. So for verifying, we can convert our customer's input email to lowercase and compare with the lowercase email column.
       **Pros:**
            - We are not messing around with COLLATE here.
        **Cons:**
            - DB change
            - Code change
            - Either we change it in every table which contains an email or change it only in the tables which have a UI interaction for email. If we choose the second option, then every time we add an email column elsewhere, we need to think about adding this column or not.

3. ### Adding COLLATE NO CASE in query:

   - In this approach, we will add COLLATE NO CASE in the WHERE clause and no DB change.
       **Pros:**
            - No DB change and only code change, so quicker to implement.
        **Cons:**
            - Performance issue. Here, since we are not indexing on case insensitivity, the query will do a table scan instead of an index scan, but this performance issue is negligible because the number of records will be less (customers table).
