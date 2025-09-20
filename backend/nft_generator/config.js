// NFT生成器配置文件 - 基于HashLips Art Engine原理
const fs = require("fs");
const path = require("path");

const layersDir = `${__dirname}/layers`;

// 获取图层元素的函数
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

// 基于真实BAYC的7个特征类别：Background, Fur, Eyes, Mouth, Clothes, Hat, Earrings
const layerConfigurations = [
  {
    growEditionSizeTo: 100, // 生成100个NFT
    layersOrder: [
      {
        name: "Background",
        elements: getElements(`${layersDir}/Background/`),
        blend: "source-over",
        opacity: 1,
      },
      {
        name: "Fur", // BAYC使用Fur而不是Base
        elements: getElements(`${layersDir}/Fur/`),
        blend: "source-over",
        opacity: 1,
      },
      {
        name: "Eyes",
        elements: getElements(`${layersDir}/Eyes/`),
        blend: "source-over",
        opacity: 1,
      },
      {
        name: "Mouth",
        elements: getElements(`${layersDir}/Mouth/`),
        blend: "source-over",
        opacity: 1,
      },
      {
        name: "Clothes",
        elements: getElements(`${layersDir}/Clothes/`),
        blend: "source-over",
        opacity: 1,
      },
      {
        name: "Hat",
        elements: getElements(`${layersDir}/Hat/`),
        blend: "source-over",
        opacity: 1,
      },
      {
        name: "Earrings",
        elements: getElements(`${layersDir}/Earrings/`),
        blend: "source-over",
        opacity: 1,
      },
    ],
  },
];

const shuffleLayerConfigurations = false;

const debugLogs = false;

const extraMetadata = {
  creator: "BAYC Style NFT Generator",
};

const rarityDelimiter = "#";

const uniqueDnaTorrance = 10000;

const preview = {
  thumbPerRow: 5,
  thumbWidth: 50,
  imageRatio: 1,
  imageName: "preview.png",
};

const solanaMetadata = {
  symbol: "BAYC",
  seller_fee_basis_points: 1000, // 10%
  external_url: "",
  creators: [
    {
      address: "",
      share: 100,
    },
  ],
};

module.exports = {
  layerConfigurations,
  getElements,
  layersDir,
  shuffleLayerConfigurations,
  debugLogs,
  extraMetadata,
  rarityDelimiter,
  uniqueDnaTorrance,
  preview,
  solanaMetadata,
};