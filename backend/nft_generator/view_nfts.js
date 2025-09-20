const http = require('http');
const fs = require('fs');
const path = require('path');

const PORT = 1001;
const buildDir = path.join(__dirname, 'build');

// 获取所有图片文件
function getImageFiles() {
  const imagesDir = path.join(buildDir, 'images');
  if (!fs.existsSync(imagesDir)) {
    return [];
  }
  
  return fs.readdirSync(imagesDir)
    .filter(file => file.endsWith('.png'))
    .sort((a, b) => {
      const numA = parseInt(a.replace('.png', ''));
      const numB = parseInt(b.replace('.png', ''));
      return numA - numB;
    });
}

// 获取元数据
function getMetadata() {
  const metadataPath = path.join(buildDir, 'json', '_metadata.json');
  if (!fs.existsSync(metadataPath)) {
    return [];
  }
  
  try {
    const data = fs.readFileSync(metadataPath, 'utf8');
    return JSON.parse(data);
  } catch (error) {
    console.error('读取元数据失败:', error);
    return [];
  }
}

// 生成HTML页面
function generateHTML() {
  const images = getImageFiles();
  const metadata = getMetadata();
  
  let html = `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>NFT图层混合展示</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 20px;
            background-color: #f0f0f0;
        }
        .header {
            text-align: center;
            margin-bottom: 30px;
        }
        .gallery {
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
            gap: 20px;
            max-width: 1200px;
            margin: 0 auto;
        }
        .nft-card {
            background: white;
            border-radius: 10px;
            padding: 15px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
            transition: transform 0.2s;
        }
        .nft-card:hover {
            transform: translateY(-5px);
        }
        .nft-image {
            width: 100%;
            height: 200px;
            object-fit: contain;
            border-radius: 5px;
            background: #f8f8f8;
        }
        .nft-info {
            margin-top: 10px;
        }
        .nft-title {
            font-size: 18px;
            font-weight: bold;
            margin-bottom: 5px;
        }
        .nft-dna {
            font-size: 12px;
            color: #666;
            margin-bottom: 10px;
            word-break: break-all;
        }
        .attributes {
            display: flex;
            flex-wrap: wrap;
            gap: 5px;
        }
        .attribute {
            background: #e3f2fd;
            padding: 3px 8px;
            border-radius: 15px;
            font-size: 11px;
            color: #1976d2;
        }
        .stats {
            text-align: center;
            margin-bottom: 20px;
            padding: 15px;
            background: white;
            border-radius: 10px;
            box-shadow: 0 2px 5px rgba(0,0,0,0.1);
        }
    </style>
</head>
<body>
    <div class="header">
        <h1>🎨 NFT图层混合展示</h1>
        <p>展示通过Canvas图层混合技术生成的NFT作品</p>
    </div>
    
    <div class="stats">
        <h3>生成统计</h3>
        <p>总共生成了 <strong>${images.length}</strong> 张NFT图片</p>
        <p>每张图片都是通过多个图层混合而成，包含背景、身体、眼睛、嘴巴、帽子等元素</p>
    </div>
    
    <div class="gallery">
  `;
  
  images.slice(0, 20).forEach((image, index) => {
    const imageNumber = parseInt(image.replace('.png', ''));
    const nftMetadata = metadata.find(m => m.edition === imageNumber) || {};
    
    html += `
        <div class="nft-card">
            <img src="/image/${image}" alt="NFT #${imageNumber}" class="nft-image">
            <div class="nft-info">
                <div class="nft-title">NFT #${imageNumber}</div>
                <div class="nft-dna">DNA: ${nftMetadata.dna || 'N/A'}</div>
                <div class="attributes">
    `;
    
    if (nftMetadata.attributes) {
      nftMetadata.attributes.forEach(attr => {
        if (attr.value && attr.value.trim()) {
          html += `<span class="attribute">${attr.trait_type}: ${attr.value}</span>`;
        }
      });
    }
    
    html += `
                </div>
            </div>
        </div>
    `;
  });
  
  html += `
    </div>
    
    <div style="text-align: center; margin-top: 30px; color: #666;">
        <p>显示前20张图片 | 每张图片都是独特的图层组合</p>
    </div>
</body>
</html>
  `;
  
  return html;
}

// 创建服务器
const server = http.createServer((req, res) => {
  const url = req.url;
  
  if (url === '/' || url === '/index.html') {
    // 返回主页
    res.writeHead(200, { 'Content-Type': 'text/html; charset=utf-8' });
    res.end(generateHTML());
    
  } else if (url.startsWith('/image/')) {
    // 返回图片
    const imageName = url.replace('/image/', '');
    const imagePath = path.join(buildDir, 'images', imageName);
    
    if (fs.existsSync(imagePath)) {
      const imageData = fs.readFileSync(imagePath);
      res.writeHead(200, { 'Content-Type': 'image/png' });
      res.end(imageData);
    } else {
      res.writeHead(404, { 'Content-Type': 'text/plain' });
      res.end('Image not found');
    }
    
  } else {
    // 404
    res.writeHead(404, { 'Content-Type': 'text/plain' });
    res.end('Not found');
  }
});

server.listen(PORT, '0.0.0.0', () => {
  console.log(`🌐 NFT展示服务器启动成功！`);
  console.log(`📱 本地访问: http://localhost:${PORT}`);
  console.log(`🌍 外部访问: http://59.110.161.193:${PORT}`);
  console.log(`🎨 展示 ${getImageFiles().length} 张混合生成的NFT图片`);
});