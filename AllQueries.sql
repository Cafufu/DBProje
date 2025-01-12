CREATE TABLE "public"."bill_types" (
  "id" integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "type" varchar(144)
)

CREATE TABLE "public"."bills" (
  "id" integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "user_id" integer,
  "type_id" integer,
  "bill_name" varchar(144),
  "year" integer,
  "month" integer,
  "amount" real,
  CONSTRAINT "bills_to_users" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id"),
  CONSTRAINT "bills_to_types" FOREIGN KEY ("type_id") REFERENCES "public"."bill_types" ("id")
)

CREATE TABLE "public"."carbon_footprint" (
  "id" integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "user_id" integer,
  "value" varchar(144),
  CONSTRAINT "footprint_to_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id")
)

CREATE TABLE "public"."users" (
  "id" integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "name" varchar(144) NOT NULL,
  "surname" varchar(144) NOT NULL,
  "password" varchar(144),
  "user_name" varchar(144),
  "phone_number" varchar(144)
);

ALTER TABLE "public"."users" ADD COLUMN "avg_amount" real;
ALTER TABLE "public"."bills" ALTER COLUMN "year" TYPE varchar(144)
ALTER TABLE "public"."bills" ALTER COLUMN "month" TYPE varchar(144)
ALTER TABLE bills
ADD CONSTRAINT check_amount_max
CHECK (amount <= 10000);

CREATE TYPE fee_record AS (
    amount NUMERIC,
    type_id INT
);

//this function is created to check if any tuple exist with the given user_id in the carbon_footprint table
CREATE OR REPLACE FUNCTION carbon_exist(myid users.id%type)
RETURNS BOOLEAN AS $$
 DECLARE
     sonuc BOOLEAN;
 BEGIN
     SELECT EXISTS (
        SELECT 1 
         FROM carbon_footprint
         WHERE id = user_id
     ) INTO sonuc;
     
     RETURN sonuc;
 END;
 $$ LANGUAGE plpgsql;

 CREATE OR REPLACE FUNCTION update_carbon_footprint()
RETURNS TRIGGER AS $$
DECLARE
    total NUMERIC := 0;  -- Toplam karbon ayak izi
    carbon_cursor CURSOR FOR
        WITH combined_fees AS (
            -- Elektrik faturaları
            SELECT rec.amount, rec.type_id
            FROM bills rec
            WHERE rec.user_id = NEW.user_id AND rec.type_id = 1
            UNION
            -- Su faturaları
            SELECT rec.amount, rec.type_id
            FROM bills rec
            WHERE rec.user_id = NEW.user_id AND rec.type_id = 2
            UNION
            -- Doğalgaz faturaları
            SELECT rec.amount, rec.type_id
            FROM bills rec
            WHERE rec.user_id = NEW.user_id AND rec.type_id = 3
        )
        SELECT amount, type_id FROM combined_fees;  -- CTE'den verileri seçiyoruz
    rec fee_record;  -- fee_record türünde değişken, cursor ile her satır için kullanılacak
BEGIN
    -- Cursor'u aç ve satırları işle
    OPEN carbon_cursor;
    LOOP
        FETCH carbon_cursor INTO rec;
        EXIT WHEN NOT FOUND;

        -- Fatura türüne göre karbon hesaplama
        IF rec.type_id = 1 THEN
            total := total + (rec.amount * 0.35 / 4);  -- Elektrik
        ELSIF rec.type_id = 2 THEN
            total := total + (rec.amount * 0.4 / 25);  -- Su
        ELSIF rec.type_id = 3 THEN
            total := total + (rec.amount * 1.92 / 15);  -- Doğalgaz
        END IF;
    END LOOP;
    
    -- Cursor'u kapat
    CLOSE carbon_cursor;

    -- Sonuç değeri 3 basamağa yuvarla
    total := ROUND(total, 3);

    -- carbon_footprint tablosunu güncelle
    UPDATE carbon_footprint
    SET value = total
    WHERE user_id = NEW.user_id;

    -- Eğer kayıt yoksa, ekle
    INSERT INTO carbon_footprint (user_id, value)
    SELECT NEW.user_id, total
    ON CONFLICT (user_id) DO UPDATE
    SET value = EXCLUDED.value;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION update_user_avg()
RETURNS TRIGGER AS $$
BEGIN
    -- Her kullanıcı için ortalama fatura miktarını hesapla
    UPDATE users
    SET avg_amount = (
        SELECT AVG(amount)
        FROM bills
        WHERE user_id = users.user_id
    );

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_update_carbon_footprint
AFTER INSERT OR UPDATE ON bills
FOR EACH ROW
EXECUTE FUNCTION update_carbon_footprint();

CREATE TRIGGER trg_update_user_avg
AFTER INSERT OR UPDATE OR DELETE ON bills
FOR EACH STATEMENT
EXECUTE FUNCTION update_user_avg();

CREATE VIEW analiz AS 
SELECT type_id,avg(amount),user_id
FROM bills
GROUP BY type_id, user_id
having avg(amount) > (select avg_amount from users where id = user_id)

CREATE VIEW avg_amount_view AS 
SELECT avg(amount) AS average_amount FROM users