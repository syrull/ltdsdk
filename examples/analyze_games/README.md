# Setup

Create empty file `games.sqlite` and execute the `games-schema.sql` on it. Then run:

```console
$ go run main.go
```

# Analyze Games

```sql
-- Get percentage of games that finished at wave 15, 13 and 10
SELECT 
    AVG(
        CASE 
            WHEN ending_wave = 15
                THEN 1.0 ELSE 0 
                END
        ) AS won_at_15,
    AVG(
        CASE 
            WHEN ending_wave = 13
                THEN 1.0 ELSE 0 
                END
        ) AS won_at_13,
    AVG(
        CASE 
            WHEN ending_wave = 10
                THEN 1.0 ELSE 0 
                END
        ) AS won_at_10
FROM 
    games;
```