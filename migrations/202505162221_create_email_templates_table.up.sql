CREATE TABLE email_templates (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL UNIQUE,
	subject TEXT NOT NULL,
	message TEXT NOT NULL, 
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
