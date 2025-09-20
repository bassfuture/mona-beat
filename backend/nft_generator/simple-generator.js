const fs = require("fs");
const path = require("path");
const sha1 = require("sha1");
const { layerConfigurations } = require("./config.js");

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

// 全局变量
let metadataList = [];
let dnaList = new Set();

// 获取图层元素
const getElements = (path) => {
  if (!fs.existsSync(path)) {
    console.log(`⚠️  图层目录不存在: ${path}`);
    return [];
  }
  
  return fs
    .readdirSync(path)
    .filter((item) => !/(^|\/)\.[^\/.]/g.test(item))
    .filter((item) => item.endsWith('.png'))
    .map((i, index) => {
      const sublayer = i.split("#");
      return {
        id: index,
        name: sublayer[0],
        filename: i,
        path: `${path}${i}`,
        weight: sublayer[1] ? parseInt(sublayer[1].replace('.png', '')) : 1,
      };
    });
};

// 获取图层配置
const layersSetup = (layersOrder) => {
  const layers = layersOrder.map((layerObj, index) => ({
    id: index,
    elements: getElements(`${layersDir}/${layerObj.name}/`),
    name: layerObj.name,
    blend: layerObj.blend || "source-over",
    opacity: layerObj.opacity || 1,
    bypassDNA: layerObj.options?.bypassDNA || false,
  }));
  return layers;
};

// 生成随机数
const random = () => {
  return Math.random();
};

// 根据权重选择元素
const getRandomWeightedElement = (elements) => {
  if (elements.length === 0) return null;
  
  const weights = elements.map((element) => element.weight);
  const totalWeight = weights.reduce((acc, weight) => acc + weight, 0);
  const randomNum = random() * totalWeight;
  let weightSum = 0;
  
  for (let i = 0; i < elements.length; i++) {
    weightSum += weights[i];
    if (randomNum <= weightSum) {
      return elements[i];
    }
  }
  return elements[0]; // 备用返回
};

// 清理DNA
const cleanDna = (_str) => {
  const withoutOptions = _str.split(":").shift();
  return Number(withoutOptions);
};

// 清理名称
const cleanName = (_str) => {
  let nameWithoutExtension = _str.slice(0, -4);
  let nameWithoutWeight = nameWithoutExtension.split("#").shift();
  return nameWithoutWeight;
};

// 获取稀有度权重
const getRarityWeight = (_str) => {
  let nameWithoutExtension = _str.slice(0, -4);
  let nameWithoutWeight = nameWithoutExtension.split("#").pop();
  return nameWithoutWeight;
};

// 检查DNA是否唯一
const isDnaUnique = (_DnaList = new Set(), _dna = "") => {
  const _filteredDNA = filterDNAOptions(_dna);
  return !_DnaList.has(_filteredDNA);
};

// 过滤DNA选项
const filterDNAOptions = (_dna) => {
  const dnaItems = _dna.split("-");
  const filteredDNA = dnaItems.filter((element) => {
    const query = element.split(":")[0];
    const queryBool = !query.includes("?");
    return queryBool;
  });
  return filteredDNA.join("-");
};

// 创建DNA
const createDna = (_layers) => {
  let randNum = [];
  _layers.forEach((layer) => {
    if (layer.elements.length === 0) {
      console.log(`⚠️  图层 ${layer.name} 没有元素`);
      return;
    }
    
    const selectedElement = getRandomWeightedElement(layer.elements);
    if (selectedElement) {
      randNum.push(
        `${selectedElement.id}:${selectedElement.filename}${
          layer.bypassDNA ? "?1" : ""
        }`
      );
    }
  });
  return randNum.join("-");
};

// 构建DNA到图层的映射
const constructLayerToDna = (_dna = "", _layers = []) => {
  let mappedDnaToLayers = _layers.map((layer, index) => {
    const dnaSegments = _dna.split("-");
    if (index >= dnaSegments.length) return null;
    
    let selectedElement = layer.elements.find(
      (e) => e.id == cleanDna(dnaSegments[index])
    );
    
    if (!selectedElement && layer.elements.length > 0) {
      selectedElement = layer.elements[0]; // 备用选择
    }
    
    return {
      name: layer.name,
      blend: layer.blend,
      opacity: layer.opacity,
      selectedElement: selectedElement,
    };
  }).filter(layer => layer !== null && layer.selectedElement !== null);
  
  return mappedDnaToLayers;
};

// 生成单个NFT元数据
const generateSingleNFT = (editionCount = 1) => {
  const layers = layersSetup(layerConfigurations[0].layersOrder);
  const newDna = createDna(layers);
  
  if (!isDnaUnique(dnaList, newDna)) {
    console.log("DNA已存在，重新生成...");
    return generateSingleNFT(editionCount);
  }
  
  const results = constructLayerToDna(newDna, layers);
  const attributes = [];
  
  // 生成属性
  results.forEach((layer) => {
    if (layer.selectedElement) {
      attributes.push({
        trait_type: layer.name,
        value: cleanName(layer.selectedElement.name),
        rarity: getRarityWeight(layer.selectedElement.filename),
      });
    }
  });
  
  // 生成元数据
  const metadata = {
    name: `Capture NFT #${editionCount}`,
    description: "A unique NFT generated from successful creature capture",
    image: `${editionCount}.png`,
    dna: newDna,
    edition: editionCount,
    date: Date.now(),
    attributes: attributes,
    compiler: "Capture Game NFT Generator",
  };
  
  dnaList.add(filterDNAOptions(newDna));
  
  return {
    metadata,
    layers: results,
    dna: newDna
  };
};

// 写入元数据
const writeMetaData = (_data, filename = "_metadata.json") => {
  fs.writeFileSync(`${buildDir}/json/${filename}`, _data);
};

// 导出函数
module.exports = { 
  generateSingleNFT,
  writeMetaData,
  layersSetup,
  getElements
};

// 如果直接运行此文件
if (require.main === module) {
  console.log("🎨 启动简化版NFT生成器...");
  
  try {
    const result = generateSingleNFT(1);
    console.log("✅ 生成成功:");
    console.log("DNA:", result.dna);
    console.log("属性:", result.metadata.attributes);
    
    writeMetaData(JSON.stringify(result.metadata, null, 2), "sample.json");
    console.log("📄 元数据已保存到 build/json/sample.json");
  } catch (error) {
    console.error("❌ 生成失败:", error);
  }
}