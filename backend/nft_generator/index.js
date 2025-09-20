const { startCreating } = require("./generator.js");

console.log("🎨 启动像素艺术NFT生成器...");
console.log("📁 检查图层文件夹结构...");

// 检查必要的文件夹是否存在
const fs = require("fs");
const path = require("path");

const requiredDirs = [
  "./layers",
  "./layers/Background",
  "./layers/Base", 
  "./layers/Eyes",
  "./layers/Hat",
  "./layers/Accessories"
];

// 创建必要的目录
requiredDirs.forEach(dir => {
  if (!fs.existsSync(dir)) {
    fs.mkdirSync(dir, { recursive: true });
    console.log(`✅ 创建目录: ${dir}`);
  }
});

// 检查是否有基础图像文件
const baseImagePath = "./layers/Base/base#100.png";
if (!fs.existsSync(baseImagePath)) {
  console.log("⚠️  请将b.png复制到 ./layers/Base/base#100.png");
  console.log("💡 提示: 文件名格式为 [名称]#[权重].png，权重越高出现概率越大");
}

console.log("🚀 开始生成NFT集合...");

// 开始生成
startCreating().then(() => {
  console.log("✨ NFT生成完成!");
  console.log("📂 生成的文件位于 ./build/ 目录");
  console.log("🖼️  图片文件: ./build/images/");
  console.log("📄 元数据文件: ./build/json/");
}).catch((error) => {
  console.error("❌ 生成过程中出现错误:", error);
});