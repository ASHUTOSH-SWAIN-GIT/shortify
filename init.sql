-- Create the urls table
CREATE TABLE IF NOT EXISTS urls (
    short_code VARCHAR(10) PRIMARY KEY,
    long_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert some sample data (optional)
-- INSERT INTO urls (short_code, long_url) VALUES 
-- ('test', 'https://example.com'),
-- ('demo', 'https://google.com');
