const fs = require("fs");
const { createCanvas, loadImage } = require("canvas");
const console = require("console");
const { layerConfigurations } = require("./config.js");

// 画布设置
const format = {
  width: 512,
  height: 512,
  smoothing: false,
};

// 全局变量
const canvas = createCanvas(format.width, format.height);
const ctx = canvas.getContext("2d");
ctx.imageSmoothingEnabled = format.smoothing;

let metadataList = [];
let attributesList = [];
let dnaList = new Set();
let editionCount = 1;

// 构建目录
const buildDir = `${__dirname}/build`;
const layersDir = `${__dirname}/layers`;

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
  return fs
    .readdirSync(path)
    .filter((item) => !/(^|\/)\.[^\/.]/g.test(item))
    .map((i, index) => {
      const sublayer = i.split("#");
      return {
        id: index,
        name: sublayer[0],
        filename: i,
        path: `${path}${i}`,
        weight: sublayer[1] ? parseInt(sublayer[1]) : 1,
      };
    });
};

// 获取图层
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

// 保存图像
const saveImage = (_editionCount) => {
  fs.writeFileSync(
    `${buildDir}/images/${_editionCount}.png`,
    canvas.toBuffer("image/png")
  );
};

// 生成随机数
const random = () => {
  return Math.random();
};

// 根据权重选择元素
const getRandomWeightedElement = (elements) => {
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

// 检查DNA是否存在
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
    var totalWeight = 0;
    layer.elements.forEach((element) => {
      totalWeight += element.weight;
    });
    let random = Math.floor(Math.random() * totalWeight);
    for (var i = 0; i < layer.elements.length; i++) {
      random -= layer.elements[i].weight;
      if (random < 0) {
        return randNum.push(
          `${layer.elements[i].id}:${layer.elements[i].filename}${
            layer.bypassDNA ? "?1" : ""
          }`
        );
      }
    }
  });
  return randNum.join("-");
};

// 构建设置
const constructLayerToDna = (_dna = "", _layers = []) => {
  let mappedDnaToLayers = _layers.map((layer, index) => {
    let selectedElement = layer.elements.find(
      (e) => e.id == cleanDna(_dna.split("-")[index])
    );
    return {
      name: layer.name,
      blend: layer.blend,
      opacity: layer.opacity,
      selectedElement: selectedElement,
    };
  });
  return mappedDnaToLayers;
};

// 加载图层图像
const loadLayerImg = async (_layer) => {
  try {
    return new Promise(async (resolve, reject) => {
      const image = await loadImage(_layer.selectedElement.path);
      resolve({ layer: _layer, loadedImage: image });
    });
  } catch (error) {
    console.error("Error loading image:", _layer.selectedElement.path, error);
    throw error;
  }
};

// 绘制元素
const drawElement = (_renderObject, _index) => {
  ctx.globalAlpha = _renderObject.layer.opacity;
  ctx.globalCompositeOperation = _renderObject.layer.blend;
  ctx.drawImage(
    _renderObject.loadedImage,
    0,
    0,
    format.width,
    format.height
  );
};

// 添加元数据
const addMetadata = (_dna, _edition) => {
  let dateTime = Date.now();
  let tempMetadata = {
    name: `#${_edition}`,
    description: "Pixel Art NFT Collection",
    image: `${_edition}.png`,
    dna: _dna,
    edition: _edition,
    date: dateTime,
    attributes: attributesList,
    compiler: "Pixel Art NFT Generator",
  };
  metadataList.push(tempMetadata);
  attributesList = [];
};

// 添加属性
const addAttributes = (_element) => {
  let selectedElement = _element.layer.selectedElement;
  attributesList.push({
    trait_type: _element.layer.name,
    value: cleanName(selectedElement.name),
    rarity: getRarityWeight(selectedElement.filename),
  });
};

// 开始创建
const startCreating = async () => {
  let layerConfigIndex = 0;
  let editionCount = 1;
  let failedCount = 0;
  let abstractedIndexes = [];
  
  for (
    let i = 1;
    i <= layerConfigurations[layerConfigIndex].growEditionSizeTo;
    i++
  ) {
    abstractedIndexes.push(i);
  }
  
  while (layerConfigIndex < layerConfigurations.length) {
    const layers = layersSetup(
      layerConfigurations[layerConfigIndex].layersOrder
    );
    while (
      editionCount <= layerConfigurations[layerConfigIndex].growEditionSizeTo
    ) {
      let newDna = createDna(layers);
      if (isDnaUnique(dnaList, newDna)) {
        let results = constructLayerToDna(newDna, layers);
        let loadedElements = [];

        results.forEach((layer) => {
          loadedElements.push(loadLayerImg(layer));
        });

        await Promise.all(loadedElements).then((renderObjectArray) => {
          ctx.clearRect(0, 0, format.width, format.height);
          renderObjectArray.forEach((renderObject, index) => {
            drawElement(renderObject, index);
            addAttributes(renderObject);
          });
          saveImage(editionCount);
          addMetadata(newDna, editionCount);
          console.log(
            `Created edition: ${editionCount}, with DNA: ${newDna}`
          );
        });
        dnaList.add(filterDNAOptions(newDna));
        editionCount++;
        failedCount = 0;
      } else {
        console.log("DNA exists!");
        failedCount++;
        if (failedCount >= 10000) {
          console.log(
            `You need more layers or elements to grow your edition to ${layerConfigurations[layerConfigIndex].growEditionSizeTo} artworks!`
          );
          break;
        }
      }
    }
    layerConfigIndex++;
  }
  writeMetaData(JSON.stringify(metadataList, null, 2));
};

// 写入元数据
const writeMetaData = (_data) => {
  fs.writeFileSync(`${buildDir}/json/_metadata.json`, _data);
};

// 主函数
module.exports = { startCreating };