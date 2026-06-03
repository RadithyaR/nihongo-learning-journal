Database Specification Final

Tabel Users
Column Type Constraint
id UUID PK
name VARCHAR(100) NOT NULL
email VARCHAR(255) UNIQUE NOT NULL
password_hash TEXT NOT NULL
avatar_url TEXT NULL
created_at TIMESTAMP NOT NULL
updated_at TIMESTAMP NOT NULL
deleted_at TIMESTAMP NULL

Tabel User Sessions
Column Type Constraint
id UUID PK
user_id UUID FK
refresh_token_hash TEXT NOT NULL
device_name VARCHAR(255) NULL
ip_address VARCHAR(255) NULL
user_agent TEXT NULL
expires_at TIMESTAMP NOT NULL
last_used_at TIMESTAMP NULL
created_at TIMESTAMP NOT NULL
updated_at TIMESTAMP NOT NULL

Tabel Vocabularies
Column Type
id UUID
user_id UUID
word VARCHAR(255)
reading VARCHAR(255)
meaning TEXT
source VARCHAR(255)
note TEXT
status VARCHAR(20)
favorite BOOLEAN
created_at TIMESTAMP
updated_at TIMESTAMP
deleted_at TIMESTAMP
Constraint:
UNIQUE(user_id, word)

Tabel Grammars
Column Type
id UUID
user_id UUID
title VARCHAR(255)
pattern TEXT
meaning TEXT
note TEXT
example TEXT
source VARCHAR(255)
status VARCHAR(20)
favorite BOOLEAN
review_count INTEGER
last_reviewed_at TIMESTAMP
created_at TIMESTAMP
updated_at TIMESTAMP
deleted_at TIMESTAMP
Constraint:
UNIQUE(user_id, title)

Tabel Grammar Attachments
Column Type
id UUID
grammar_id UUID
file_name VARCHAR(255)
file_url TEXT
created_at TIMESTAMP
Tabel Kanjis
Column Type
id UUID
user_id UUID
kanji VARCHAR(20)
reading TEXT
meaning TEXT
note TEXT
status VARCHAR(20)
favorite BOOLEAN
created_at TIMESTAMP
updated_at TIMESTAMP
deleted_at TIMESTAMP
Constraint:
UNIQUE(user_id, kanji)

Tabel Study Sessions
Column Type
id UUID
user_id UUID
study_date DATE
duration_minutes INTEGER
note TEXT
title VARCHAR(255)
created_at TIMESTAMP
updated_at TIMESTAMP
deleted_at TIMESTAMP

Tabel Goals
Column Type
id UUID
user_id UUID
title VARCHAR(255)
description TEXT
target_level VARCHAR(10)
target_date DATE
status VARCHAR(20)
created_at TIMESTAMP
updated_at TIMESTAMP
deleted_at TIMESTAMP

Tabel Tags
Column Type
id UUID
user_id UUID
name VARCHAR(100)
created_at TIMESTAMP
updated_at TIMESTAMP
deleted_at TIMESTAMP

Constraint:
UNIQUE(user_id, name)

Tabel Review Logs
Column Type
id UUID
user_id UUID
entity_type VARCHAR(20)
entity_id UUID
previous_status VARCHAR(20)
new_status VARCHAR(20)
reviewed_at TIMESTAMP
created_at TIMESTAMP

vocabulary_tags
Column Type
vocabulary_id UUID
tag_id UUID
PK:
PRIMARY KEY(vocabulary_id, tag_id)

grammar_tags
Column Type
grammar_id UUID
tag_id UUID
PK:
PRIMARY KEY(grammar_id, tag_id)

kanji_tags
Column Type
kanji_id UUID
tag_id UUID
PK:
PRIMARY KEY(kanji_id, tag_id)

study_session_vocabularies
Column Type
study_session_id UUID
vocabulary_id UUID
PK:
PRIMARY KEY(study_session_id, vocabulary_id)

study_session_grammars
Column Type
study_session_id UUID
grammar_id UUID
PK:
PRIMARY KEY(study_session_id, grammar_id)

study_session_kanjis
Column Type
study_session_id UUID
kanji_id UUID
PK:
PRIMARY KEY(study_session_id, kanji_id)

Status Enum

Learning Status
NEW
LEARNING
REVIEWING
MASTERED

Goal Status
ACTIVE
COMPLETED
CANCELLED

Entity Type
VOCABULARY
GRAMMAR
KANJI

Relasi Final
User
│
├── Vocabularies
├── Grammars
├── Kanjis
├── Study Sessions
├── Goals
├── Tags
├── Review Logs
└── User Sessions
