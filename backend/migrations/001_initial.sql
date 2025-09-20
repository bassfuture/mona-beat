-- 创建用户表
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    wallet_address VARCHAR(42) UNIQUE NOT NULL,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100),
    total_captures INTEGER DEFAULT 0,
    successful_captures INTEGER DEFAULT 0,
    total_nfts INTEGER DEFAULT 0,
    last_capture_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 创建捕捉表
CREATE TABLE IF NOT EXISTS captures (
    id VARCHAR(64) PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    wallet_address VARCHAR(42) NOT NULL,
    success BOOLEAN NOT NULL DEFAULT FALSE,
    rarity VARCHAR(20),
    nft_token_id BIGINT,
    transaction_hash VARCHAR(66),
    metadata JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 创建NFT表
CREATE TABLE IF NOT EXISTS nfts (
    token_id BIGINT PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    wallet_address VARCHAR(42) NOT NULL,
    capture_id VARCHAR(64) NOT NULL REFERENCES captures(id),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    image_url VARCHAR(500) NOT NULL,
    metadata_url VARCHAR(500) NOT NULL,
    rarity VARCHAR(20) NOT NULL,
    attributes JSONB,
    transaction_hash VARCHAR(66),
    minted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 创建索引
CREATE INDEX IF NOT EXISTS idx_users_wallet_address ON users(wallet_address);
CREATE INDEX IF NOT EXISTS idx_captures_user_id ON captures(user_id);
CREATE INDEX IF NOT EXISTS idx_captures_wallet_address ON captures(wallet_address);
CREATE INDEX IF NOT EXISTS idx_nfts_user_id ON nfts(user_id);
CREATE INDEX IF NOT EXISTS idx_nfts_wallet_address ON nfts(wallet_address);
CREATE INDEX IF NOT EXISTS idx_nfts_capture_id ON nfts(capture_id);
CREATE INDEX IF NOT EXISTS idx_nfts_rarity ON nfts(rarity);