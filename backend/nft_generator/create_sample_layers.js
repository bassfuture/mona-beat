const fs = require("fs");
const { createCanvas } = require("canvas");

console.log("🎨 开始创建BAYC风格的图层文件...");

// 创建完全程序化的BAYC风格图层，不依赖外部图像

// 1. 创建背景图层 (Background)
const backgrounds = [
  { name: "blue", color: "#87CEEB", weight: 30 },
  { name: "green", color: "#90EE90", weight: 25 },
  { name: "purple", color: "#DDA0DD", weight: 20 },
  { name: "orange", color: "#FFA500", weight: 15 },
  { name: "red", color: "#FF6B6B", weight: 10 }
];

for (const bg of backgrounds) {
  // 为每个背景创建新的画布
  const bgCanvas = createCanvas(512, 512);
  const bgCtx = bgCanvas.getContext("2d");
  
  bgCtx.fillStyle = bg.color;
  bgCtx.fillRect(0, 0, 512, 512);
  
  const buffer = bgCanvas.toBuffer("image/png");
  fs.writeFileSync(`layers/Background/${bg.name}#${bg.weight}.png`, buffer);
  console.log(`✅ 创建背景: ${bg.name}`);
}

// 2. 创建毛发图层 (Fur) - 程序化生成猿猴基础形状和毛发纹理
const furs = [
  { name: "brown", color: "#8B4513", weight: 40 },
  { name: "black", color: "#2F2F2F", weight: 30 },
  { name: "golden", color: "#DAA520", weight: 20 },
  { name: "silver", color: "#C0C0C0", weight: 10 }
];

for (const fur of furs) {
  // 创建新的透明画布
  const furCanvas = createCanvas(512, 512);
  const furCtx = furCanvas.getContext("2d");
  
  // 绘制猿猴基础形状
  furCtx.fillStyle = fur.color;
  
  // 头部轮廓
  furCtx.beginPath();
  furCtx.ellipse(256, 200, 120, 100, 0, 0, 2 * Math.PI);
  furCtx.fill();
  
  // 身体轮廓
  furCtx.beginPath();
  furCtx.ellipse(256, 380, 100, 120, 0, 0, 2 * Math.PI);
  furCtx.fill();
  
  // 手臂
  furCtx.beginPath();
  furCtx.ellipse(180, 350, 40, 80, -0.3, 0, 2 * Math.PI);
  furCtx.fill();
  
  furCtx.beginPath();
  furCtx.ellipse(332, 350, 40, 80, 0.3, 0, 2 * Math.PI);
  furCtx.fill();
  
  const buffer = furCanvas.toBuffer("image/png");
  fs.writeFileSync(`layers/Fur/${fur.name}#${fur.weight}.png`, buffer);
  console.log(`✅ 创建毛发: ${fur.name}`);
}

// 3. 创建眼睛图层 (Eyes)
const eyes = [
  { name: "normal", weight: 40 },
  { name: "big", weight: 30 },
  { name: "sleepy", weight: 20 },
  { name: "cyborg", weight: 5 }, // 稀有特征
  { name: "laser", weight: 5 }   // 稀有特征
];

for (const eye of eyes) {
  // 创建新的透明画布
  const eyeCanvas = createCanvas(512, 512);
  const eyeCtx = eyeCanvas.getContext("2d");
  
  switch (eye.name) {
    case "normal":
      // 普通眼睛 - 白色眼球，黑色瞳孔
      eyeCtx.fillStyle = "#FFFFFF";
      eyeCtx.fillRect(180, 200, 40, 40);
      eyeCtx.fillRect(290, 200, 40, 40);
      eyeCtx.fillStyle = "#000000";
      eyeCtx.fillRect(190, 210, 20, 20);
      eyeCtx.fillRect(300, 210, 20, 20);
      break;
    case "big":
      // 大眼睛
      eyeCtx.fillStyle = "#FFFFFF";
      eyeCtx.fillRect(170, 190, 60, 60);
      eyeCtx.fillRect(280, 190, 60, 60);
      eyeCtx.fillStyle = "#000000";
      eyeCtx.fillRect(185, 205, 30, 30);
      eyeCtx.fillRect(295, 205, 30, 30);
      break;
    case "sleepy":
      // 困倦眼睛
      eyeCtx.fillStyle = "#FFFFFF";
      eyeCtx.fillRect(180, 210, 40, 20);
      eyeCtx.fillRect(290, 210, 40, 20);
      eyeCtx.fillStyle = "#000000";
      eyeCtx.fillRect(190, 215, 20, 10);
      eyeCtx.fillRect(300, 215, 20, 10);
      break;
    case "cyborg":
      // 机械眼 - 稀有特征
      eyeCtx.fillStyle = "#FF0000";
      eyeCtx.fillRect(180, 200, 40, 40);
      eyeCtx.fillStyle = "#00FF00";
      eyeCtx.fillRect(290, 200, 40, 40);
      break;
    case "laser":
      // 激光眼 - 稀有特征
      eyeCtx.fillStyle = "#FF0000";
      eyeCtx.fillRect(180, 200, 40, 40);
      eyeCtx.fillRect(290, 200, 40, 40);
      // 激光效果
      eyeCtx.fillStyle = "#FFFF00";
      eyeCtx.fillRect(200, 220, 100, 5);
      eyeCtx.fillRect(310, 220, 100, 5);
      break;
  }
  
  const buffer = eyeCanvas.toBuffer("image/png");
  fs.writeFileSync(`layers/Eyes/${eye.name}#${eye.weight}.png`, buffer);
  console.log(`✅ 创建眼睛: ${eye.name}`);
}

// 4. 创建嘴巴图层 (Mouth)
const mouths = [
  { name: "bored", weight: 40 },
  { name: "smile", weight: 30 },
  { name: "grin", weight: 20 },
  { name: "dagger", weight: 5 }, // 稀有特征
  { name: "phoneme", weight: 5 } // 稀有特征
];

for (const mouth of mouths) {
  // 创建新的透明画布
  const mouthCanvas = createCanvas(512, 512);
  const mouthCtx = mouthCanvas.getContext("2d");
  
  switch (mouth.name) {
    case "bored":
      // 无聊的嘴
      mouthCtx.fillStyle = "#000000";
      mouthCtx.fillRect(230, 300, 50, 10);
      break;
    case "smile":
      // 微笑
      mouthCtx.fillStyle = "#000000";
      mouthCtx.beginPath();
      mouthCtx.arc(256, 300, 25, 0, Math.PI);
      mouthCtx.fill();
      break;
    case "grin":
      // 咧嘴笑
      mouthCtx.fillStyle = "#FFFFFF";
      mouthCtx.fillRect(220, 290, 70, 20);
      mouthCtx.fillStyle = "#000000";
      mouthCtx.fillRect(225, 295, 60, 10);
      break;
    case "dagger":
      // 嘴里叼着匕首 - 稀有特征
      mouthCtx.fillStyle = "#C0C0C0";
      mouthCtx.fillRect(240, 280, 30, 5);
      mouthCtx.fillStyle = "#8B4513";
      mouthCtx.fillRect(270, 275, 15, 15);
      break;
    case "phoneme":
      // 特殊嘴型 - 稀有特征
      mouthCtx.fillStyle = "#FF69B4";
      mouthCtx.beginPath();
      mouthCtx.ellipse(256, 300, 20, 15, 0, 0, 2 * Math.PI);
      mouthCtx.fill();
      break;
  }
  
  const buffer = mouthCanvas.toBuffer("image/png");
  fs.writeFileSync(`layers/Mouth/${mouth.name}#${mouth.weight}.png`, buffer);
  console.log(`✅ 创建嘴巴: ${mouth.name}`);
}

// 5. 创建服装图层 (Clothes)
const clothes = [
  { name: "none", weight: 50 },
  { name: "tshirt", weight: 25 },
  { name: "hoodie", weight: 15 },
  { name: "suit", weight: 8 },
  { name: "service", weight: 2 } // 稀有特征
];

for (const cloth of clothes) {
  // 创建新的透明画布
  const clothCanvas = createCanvas(512, 512);
  const clothCtx = clothCanvas.getContext("2d");
  
  switch (cloth.name) {
    case "none":
      // 无服装 - 不绘制任何内容
      break;
    case "tshirt":
      // T恤
      clothCtx.fillStyle = "#4169E1";
      clothCtx.fillRect(180, 350, 150, 100);
      break;
    case "hoodie":
      // 连帽衫
      clothCtx.fillStyle = "#696969";
      clothCtx.fillRect(170, 340, 170, 120);
      // 帽子部分
      clothCtx.fillRect(200, 150, 110, 80);
      break;
    case "suit":
      // 西装
      clothCtx.fillStyle = "#000000";
      clothCtx.fillRect(180, 350, 150, 100);
      clothCtx.fillStyle = "#FFFFFF";
      clothCtx.fillRect(240, 360, 30, 80);
      break;
    case "service":
      // 军装 - 稀有特征
      clothCtx.fillStyle = "#228B22";
      clothCtx.fillRect(180, 350, 150, 100);
      clothCtx.fillStyle = "#FFD700";
      clothCtx.fillRect(200, 370, 10, 10);
      clothCtx.fillRect(280, 370, 10, 10);
      break;
  }
  
  const buffer = clothCanvas.toBuffer("image/png");
  fs.writeFileSync(`layers/Clothes/${cloth.name}#${cloth.weight}.png`, buffer);
  console.log(`✅ 创建服装: ${cloth.name}`);
}

// 6. 创建帽子图层 (Hat)
const hats = [
  { name: "none", weight: 50 },
  { name: "cap", weight: 25 },
  { name: "crown", weight: 15 },
  { name: "wizard", weight: 8 },
  { name: "trippy", weight: 2 } // 稀有特征
];

for (const hat of hats) {
  // 创建新的透明画布
  const hatCanvas = createCanvas(512, 512);
  const hatCtx = hatCanvas.getContext("2d");
  
  switch (hat.name) {
    case "none":
      // 无帽子 - 不绘制任何内容
      break;
    case "cap":
      // 棒球帽
      hatCtx.fillStyle = "#FF0000";
      hatCtx.fillRect(200, 100, 110, 60);
      // 帽檐
      hatCtx.fillRect(180, 150, 150, 20);
      break;
    case "crown":
      // 皇冠
      hatCtx.fillStyle = "#FFD700";
      hatCtx.fillRect(200, 80, 110, 40);
      // 皇冠尖
      hatCtx.fillRect(220, 60, 20, 20);
      hatCtx.fillRect(250, 50, 20, 30);
      hatCtx.fillRect(280, 60, 20, 20);
      break;
    case "wizard":
      // 巫师帽
      hatCtx.fillStyle = "#4B0082";
      hatCtx.fillRect(230, 50, 50, 100);
      hatCtx.fillRect(200, 140, 110, 20);
      break;
    case "trippy":
      // 迷幻帽 - 稀有特征
      hatCtx.fillStyle = "#FF1493";
      hatCtx.fillRect(200, 100, 110, 60);
      hatCtx.fillStyle = "#00FF00";
      hatCtx.fillRect(210, 110, 20, 20);
      hatCtx.fillRect(250, 110, 20, 20);
      hatCtx.fillRect(290, 110, 20, 20);
      break;
  }
  
  const buffer = hatCanvas.toBuffer("image/png");
  fs.writeFileSync(`layers/Hat/${hat.name}#${hat.weight}.png`, buffer);
  console.log(`✅ 创建帽子: ${hat.name}`);
}

// 7. 创建耳环图层 (Earrings)
const earrings = [
  { name: "none", weight: 60 },
  { name: "silver", weight: 25 },
  { name: "gold", weight: 10 },
  { name: "cross", weight: 4 }, // 稀有特征
  { name: "diamond", weight: 1 } // 极稀有特征
];

for (const earring of earrings) {
  // 创建新的透明画布
  const earringCanvas = createCanvas(512, 512);
  const earringCtx = earringCanvas.getContext("2d");
  
  switch (earring.name) {
    case "none":
      // 无耳环 - 不绘制任何内容
      break;
    case "silver":
      // 银耳环
      earringCtx.fillStyle = "#C0C0C0";
      earringCtx.beginPath();
      earringCtx.arc(150, 250, 8, 0, 2 * Math.PI);
      earringCtx.fill();
      break;
    case "gold":
      // 金耳环
      earringCtx.fillStyle = "#FFD700";
      earringCtx.beginPath();
      earringCtx.arc(150, 250, 10, 0, 2 * Math.PI);
      earringCtx.fill();
      break;
    case "cross":
      // 十字耳环 - 稀有特征
      earringCtx.fillStyle = "#C0C0C0";
      earringCtx.fillRect(148, 240, 4, 20);
      earringCtx.fillRect(140, 248, 20, 4);
      break;
    case "diamond":
      // 钻石耳环 - 极稀有特征
      earringCtx.fillStyle = "#B9F2FF";
      earringCtx.beginPath();
      earringCtx.moveTo(150, 240);
      earringCtx.lineTo(160, 250);
      earringCtx.lineTo(150, 260);
      earringCtx.lineTo(140, 250);
      earringCtx.closePath();
      earringCtx.fill();
      break;
  }
  
  const buffer = earringCanvas.toBuffer("image/png");
  fs.writeFileSync(`layers/Earrings/${earring.name}#${earring.weight}.png`, buffer);
  console.log(`✅ 创建耳环: ${earring.name}`);
}

console.log("🎨 所有BAYC风格图层创建完成!");
console.log("📊 特征统计:");
console.log(`- 背景: ${backgrounds.length} 种`);
console.log(`- 毛发: ${furs.length} 种`);
console.log(`- 眼睛: ${eyes.length} 种`);
console.log(`- 嘴巴: ${mouths.length} 种`);
console.log(`- 服装: ${clothes.length} 种`);
console.log(`- 帽子: ${hats.length} 种`);
console.log(`- 耳环: ${earrings.length} 种`);
console.log(`🔢 总计可能组合: ${backgrounds.length * furs.length * eyes.length * mouths.length * clothes.length * hats.length * earrings.length} 种`);