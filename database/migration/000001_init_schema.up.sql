CREATE TABLE IF NOT EXISTS "users" (
                                       "id" uuid DEFAULT gen_random_uuid() PRIMARY KEY,
"name" VARCHAR(255) NOT NULL);

CREATE TABLE IF NOT EXISTS "posts" (
     "id"  uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "body" TEXT NOT NULL,
    "user_id" uuid REFERENCES "users" ("id") ON DELETE CASCADE NOT NULL,
    "disabled" BOOLEAN NOT NULL,
     "comments" uuid
                                                  );


CREATE TABLE IF NOT EXISTS "comments" (
                                          "id" uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    "body" TEXT NOT NULL,
    "user_id" uuid REFERENCES "users" ( "id") ON DELETE CASCADE NOT NULL,
    "parent" uuid REFERENCES "comments" ("id") ON DELETE CASCADE,
                                          "children" uuid REFERENCES "comments" ("id") ON DELETE CASCADE,
                                          "post" uuid REFERENCES "posts" ("id") ON DELETE CASCADE
                                                  );



