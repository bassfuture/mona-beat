const { startCreating } = require('./generator.js');

console.log('🎨 开始生成NFT图片...');

startCreating()
  .then(() => {
    console.log('✅ NFT图片生成完成！');
  })
  .catch((error) => {
    console.error('❌ 生成失败:', error);
  });