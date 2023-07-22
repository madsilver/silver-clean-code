
DROP TABLE IF EXISTS "Account";
CREATE TABLE "Account" (
    "AccountID" SERIAL,
    "DocumentNumber" varchar(20),
    CONSTRAINT "PK_Account" PRIMARY KEY ("AccountID")
);

DROP TABLE IF EXISTS "Transaction";
CREATE TABLE "Transaction" (
    "TransactionID" SERIAL,
    "AccountID" int,
    "OperationTypeID" int,
    "Amount" decimal(15,2),
    "EventDate" timestamp,
    CONSTRAINT "PK_Transaction" PRIMARY KEY ("TransactionID")
);