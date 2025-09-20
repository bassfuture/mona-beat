#!/usr/bin/env node

const { generateSingleNFT, writeMetaData } = require('./simple-generator.js');

// 从命令行参数获取edition number
const args = process.argv.slice(2);
const editionNumber = args[0] ? parseInt(args[0]) : Math.floor(Math.random() * 10000);

// 可选的自定义名称和描述
const customName = args[1] || null;
const customDescription = args[2] || null;

try {
  console.log(`🎨 生成NFT #${editionNumber}...`);
  
  const result = generateSingleNFT(editionNumber);
  
  // 如果提供了自定义名称和描述，则使用它们
  if (customName) {
    result.metadata.name = customName;
  }
  if (customDescription) {
    result.metadata.description = customDescription;
  }
  
  // 保存元数据到文件
  const filename = `nft_${editionNumber}.json`;
  writeMetaData(JSON.stringify(result.metadata, null, 2), filename);
  
  // 输出结果
  const output = {
    success: true,
    metadata: result.metadata,
    dna: result.dna,
    filename: filename
  };
  
  console.log(JSON.stringify(output));
  
} catch (error) {
  const errorOutput = {
    success: false,
    error: error.message
  };
  
  console.error('❌ 生成失败:', error.message);
  console.log(JSON.stringify(errorOutput));
  process.exit(1);
}