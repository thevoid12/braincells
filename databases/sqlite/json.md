# json in sqlite
```sql
create table test (
jsoncol TEXT);

INSERT INTO test (jsoncol) VALUES (
    '{
      "1": {
        "id": "a92e3a5d-33b7-4d97-aec7-93c36b16411f",
        "templateID": "99e2a1c1-efde-4c1d-8c1a-603fb2e33b73"
      },
      "2": {
        "id": "6f2c10d2-72d6-4eeb-9372-8261cf14c9b5",
        "templateID": "142fd8ea-7b49-4c5a-a5a9-cfa9d6a66d0d"
      }
    }'
);
  


 SELECT * FROM test WHERE json_valid(jsoncol);
 select * from test;
 SELECT json_extract(jsoncol,'$') FROM test; -- i am displaying everything
 SELECT json_extract(jsoncol,'$.1') FROM test; -- i am displaying only the risk program cft(1)
  SELECT json_extract(jsoncol,'$.1.id') FROM test; -- i am fetching only the id part of cft(1)
	select * from test where json_extract(jsoncol,'$.1.id')= "a92e3a5d-33b7-4d97-aec7-93c36b16411f";
SELECT jsoncol->>'$.1.id' FROM test;
select * from test where jsoncol->>'$.1.id'= "a92e3a5d-33b7-4d97-aec7-93c36b16411f";
```
### ->> Operator (Extract as String)
- Returns the extracted value as a plain SQL string (TEXT).

- If the extracted value is a number, it returns it as a string.

- If the extracted value is an object or array, it returns it as a string representation.
```sql
SELECT jsoncol->>'$.1.id' FROM test;
```
op:
```sql
a92e3a5d-33b7-4d97-aec7-93c36b16411f
```
### -> Operator (Extract as JSON)
- Returns the extracted value as a JSON value.

- If the extracted value is a string, it remains as a JSON string with quotes ("value" instead of value).

- If it's an object or array, it remains structured as JSON.
```sql
SELECT jsoncol->'$.1.id' FROM test;
```
op:
```sql
"a92e3a5d-33b7-4d97-aec7-93c36b16411f"
```
