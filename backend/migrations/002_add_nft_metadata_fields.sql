-- MySQL migration for adding NFT metadata fields
-- Add new columns to nfts table
ALTER TABLE nfts 
ADD COLUMN full_metadata JSON COMMENT 'Complete NFT metadata including all attributes',
ADD COLUMN dna VARCHAR(255) COMMENT 'Unique DNA string for the NFT',
ADD COLUMN image_path VARCHAR(500) COMMENT 'File system path to the generated image',
ADD COLUMN metadata_path VARCHAR(500) COMMENT 'File system path to the metadata JSON file';

-- Create indexes for better query performance
CREATE INDEX idx_nfts_dna ON nfts(dna);
CREATE INDEX idx_nfts_image_path ON nfts(image_path);
CREATE INDEX idx_nfts_metadata_path ON nfts(metadata_path);

-- Add comments to existing columns for clarity
ALTER TABLE nfts MODIFY COLUMN attributes JSON COMMENT 'NFT attributes in JSON format';