const { createCanvas } = require('canvas');
const fs = require('fs');

console.log('开始Canvas测试...');

// 创建一个简单的测试图像
const canvas = createCanvas(512, 512);
const ctx = canvas.getContext('2d');

console.log('Canvas创建成功，尺寸:', canvas.width, 'x', canvas.height);

// 设置背景为白色
ctx.fillStyle = '#FFFFFF';
ctx.fillRect(0, 0, 512, 512);
console.log('白色背景绘制完成');

// 绘制一个蓝色矩形
ctx.fillStyle = '#0000FF';
ctx.fillRect(100, 100, 200, 200);
console.log('蓝色矩形绘制完成');

// 绘制一个红色圆圈
ctx.fillStyle = '#FF0000';
ctx.beginPath();
ctx.arc(300, 300, 50, 0, 2 * Math.PI);
ctx.fill();
console.log('红色圆圈绘制完成');

// 保存图像
try {
  const buffer = canvas.toBuffer('image/png');
  fs.writeFileSync('test_image.png', buffer);
  console.log('测试图像已创建: test_image.png');
  console.log('文件大小:', buffer.length, '字节');
  
  // 检查buffer的前几个字节
  console.log('Buffer前16字节:', buffer.slice(0, 16));
} catch (error) {
  console.error('保存图像时出错:', error);
}