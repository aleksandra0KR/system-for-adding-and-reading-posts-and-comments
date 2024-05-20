CREATE TABLE IF NOT EXISTS "users" (
                                       "userId" uuid DEFAULT gen_random_uuid() PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS "posts" (
     "postId"  uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "body" TEXT NOT NULL,
    "userId" uuid REFERENCES "users" ("userId") ON DELETE CASCADE NOT NULL,
    "disabledComments" BOOLEAN NOT NULL,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone
                                                  );


CREATE TABLE IF NOT EXISTS "comments" (
                                          "commentId" uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    "body" TEXT NOT NULL,
    "userId" uuid REFERENCES "users" ( "userId") ON DELETE CASCADE NOT NULL,
    "parent" uuid REFERENCES "comments" ("commentId") ON DELETE CASCADE,
    "updatedAt" timestamp with time zone
                                                  );



