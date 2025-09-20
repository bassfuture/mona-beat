const fs = require("fs");
const { createCanvas, loadImage } = require("canvas");
const { layerConfigurations } = require("./config.js");

// 画布设置
const format = {
  width: 512,
  height: 512,
  smoothing: false,
};

// 构建目录
const buildDir = `${__dirname}/build`;
const layersDir = `${__dirname}/layers`;

// 确保目录存在
if (!fs.existsSync(buildDir)) {
  fs.mkdirSync(buildDir, { recursive: true });
}

if (!fs.existsSync(`${buildDir}/json`)) {
  fs.mkdirSync(`${buildDir}/json`);
}

if (!fs.existsSync(`${buildDir}/images`)) {
  fs.mkdirSync(`${buildDir}/images`);
}

// 获取图层元素
const getElements = (path) => {
  if (!fs.existsSync(path)) {
    return [];
  }
  
  const elements = fs
    .readdirSync(path)
    .filter((item) => !/(^|\/)\.[^\/.]/g.test(item))
    .filter((item) => item.endsWith('.png'))
    .map((i, index) => {
      const sublayer = i.split("#");
      const weight = sublayer[1] ? parseInt(sublayer[1].replace('.png', '')) : 1;
      return {
        id: index,
        name: sublayer[0],
        filename: i,
        path: `${path}${i}`,
        weight: weight,
      };
    });
  
  // 计算总权重用于稀有度计算
  const totalWeight = elements.reduce((sum, element) => sum + element.weight, 0);
  
  // 为每个元素添加稀有度百分比
  return elements.map(element => ({
    ...element,
    rarity: ((element.weight / totalWeight) * 100).toFixed(1) + '%'
  }));
};

// 获取图层配置
const layersSetup = (layersOrder) => {
  const layers = layersOrder.map((layerObj, index) => ({
    id: index,
    elements: getElements(`${layersDir}/${layerObj.name}/`),
    name: layerObj.name,
    blend: layerObj.blend || "source-over",
    opacity: layerObj.opacity || 1,
  }));
  return layers;
};

// 随机数生成
const random = () => {
  return Math.random();
};

// 加权随机选择
const getRandomWeightedElement = (elements) => {
  const totalWeight = elements.reduce((acc, element) => acc + element.weight, 0);
  let randomWeight = random() * totalWeight;
  
  for (let element of elements) {
    randomWeight -= element.weight;
    if (randomWeight <= 0) {
      return element;
    }
  }
  
  return elements[0];
};

// 生成DNA
const createDna = (_layers) => {
  let randNum = [];
  _layers.forEach((layer) => {
    if (layer.elements.length > 0) {
      const selectedElement = getRandomWeightedElement(layer.elements);
      randNum.push(selectedElement.id);
    } else {
      randNum.push(0);
    }
  });
  return randNum.join("");
};

// 构建图层到DNA
const constructLayerToDna = (_dna = "", _layers = []) => {
  let mappedDnaToLayers = _layers.map((layer, index) => {
    let selectedElement = layer.elements[_dna.split("")[index]] || layer.elements[0];
    return {
      name: layer.name,
      blend: layer.blend,
      opacity: layer.opacity,
      selectedElement: selectedElement,
    };
  });
  return mappedDnaToLayers;
};

// 绘制图层
const drawElement = (ctx, _renderObject, _index, _layersLen) => {
  ctx.globalAlpha = _renderObject.layer.opacity;
  ctx.globalCompositeOperation = _renderObject.layer.blend;
  
  if (_renderObject.loadedImage) {
    ctx.drawImage(
      _renderObject.loadedImage,
      0,
      0,
      format.width,
      format.height
    );
  }
};

// 保存图片
const saveImage = (canvas, _editionCount) => {
  const imageName = `${_editionCount}.png`;
  const imagePath = `${buildDir}/images/${imageName}`;
  
  fs.writeFileSync(imagePath, canvas.toBuffer("image/png"));
  console.log(`图片已保存: ${imagePath}`);
  return imageName;
};

// 生成单个NFT
const generateSingleNFT = async (editionNumber = 1) => {
  try {
    // 创建新的画布和上下文
    const canvas = createCanvas(format.width, format.height);
    const ctx = canvas.getContext("2d");
    ctx.imageSmoothingEnabled = format.smoothing;
    
    const layers = layersSetup(layerConfigurations[0].layersOrder);
    const dna = createDna(layers);
    const results = constructLayerToDna(dna, layers);
    
    // 清空画布
    ctx.clearRect(0, 0, format.width, format.height);
    
    // 加载并绘制图层
    const renderObjectArray = [];
    
    for (let i = 0; i < results.length; i++) {
      const layer = results[i];
      if (layer.selectedElement && layer.selectedElement.path) {
        try {
          const loadedImage = await loadImage(layer.selectedElement.path);
          renderObjectArray.push({
            layer: layer,
            loadedImage: loadedImage,
          });
        } catch (error) {
          // Skip layer if image cannot be loaded
        }
      }
    }
    
    // 绘制所有图层
    renderObjectArray.forEach((renderObject, index) => {
      drawElement(ctx, renderObject, index, renderObjectArray.length);
    });
    
    // 保存图片
    const imageName = `${editionNumber}.png`;
    const imagePath = `${buildDir}/images/${imageName}`;
    
    fs.writeFileSync(imagePath, canvas.toBuffer("image/png"));
    
    // 生成属性
    const attributes = results
      .filter(layer => layer.selectedElement)
      .map(layer => ({
        trait_type: layer.name,
        value: layer.selectedElement.name,
        rarity: layer.selectedElement.rarity || "0%",
      }));
    
    // 生成元数据
    const metadata = {
      name: `NFT #${editionNumber}`,
      description: "Generated NFT",
      image: `${imageName}`,
      dna: dna,
      edition: editionNumber,
      date: Date.now(),
      attributes: attributes,
      compiler: "NFT Generator v1.0.0",
    };
    
    // 保存元数据
    const metadataPath = `${buildDir}/json/${editionNumber}.json`;
    fs.writeFileSync(metadataPath, JSON.stringify(metadata, null, 2));
    
    return {
      metadata: metadata,
      dna: dna,
      success: true,
      imagePath: imagePath,
      metadataPath: metadataPath
    };
    
  } catch (error) {
    return {
      success: false,
      error: error.message
    };
  }
};

// 如果直接运行此文件
if (require.main === module) {
  (async () => {
    try {
      const result = await generateSingleNFT(1);
      console.log(JSON.stringify(result));
    } catch (error) {
      console.log(JSON.stringify({
        success: false,
        error: error.message
      }));
    }
  })();
}

module.exports = { generateSingleNFT };