// NFT管理类
class NFTManager {
    constructor() {
        this.userNFTs = [];
        this.currentNFT = null;
        this.init();
    }

    // 初始化NFT管理器
    init() {
        this.bindEvents();
    }

    // 绑定事件
    bindEvents() {
        const viewTxBtn = document.getElementById('viewTransaction');
        
        if (viewTxBtn) {
            viewTxBtn.addEventListener('click', () => this.viewTransaction());
        }

        // 监听钱包连接事件
        window.addEventListener('walletConnected', () => {
            this.loadUserNFTs();
        });

        window.addEventListener('walletDisconnected', () => {
            this.clearUserNFTs();
        });
    }

    // 显示NFT模态框
    showNFTModal(nftData) {
        this.currentNFT = nftData;
        
        const modal = document.getElementById('nftDisplay');
        const nftImage = document.getElementById('nftImage');
        const nftName = document.getElementById('nftName');
        const nftRarity = document.getElementById('nftRarity');
        const nftTokenId = document.getElementById('nftTokenId');

        if (modal) {
            modal.classList.remove('hidden');
        }

        // 设置NFT信息
        if (nftImage && nftData.metadata && nftData.metadata.image) {
            nftImage.src = nftData.metadata.image;
            nftImage.alt = nftData.metadata.name || 'NFT';
        } else if (nftImage) {
            // 使用默认图片或生成SVG
            nftImage.src = this.generateDefaultNFTImage(nftData.token_id);
        }

        if (nftName) {
            nftName.textContent = nftData.metadata?.name || `NFT #${nftData.token_id}`;
        }

        if (nftRarity) {
            nftRarity.textContent = this.getRarityText(nftData.metadata?.rarity || 'common');
        }

        if (nftTokenId) {
            nftTokenId.textContent = `Token ID: ${nftData.token_id}`;
        }

        // 播放显示动画
        this.playShowAnimation();
    }

    // 隐藏NFT模态框
    hideNFTModal() {
        const modal = document.getElementById('nftDisplay');
        if (modal) {
            modal.classList.add('hidden');
        }
        this.currentNFT = null;
    }

    // 查看交易
    viewTransaction() {
        if (!this.currentNFT || !this.currentNFT.transaction_hash) {
            this.showNotification('交易哈希不可用', 'error');
            return;
        }

        const txURL = window.apiManager.getTransactionURL(this.currentNFT.transaction_hash);
        window.open(txURL, '_blank');
    }

    // 加载用户NFT列表
    async loadUserNFTs() {
        try {
            const walletAddress = window.walletManager.getAddress();
            if (!walletAddress) return;

            const result = await window.apiManager.getUserNFTs(walletAddress);
            if (result.success) {
                this.userNFTs = result.data || [];
                this.renderNFTGrid();
            } else {
                console.error('加载NFT列表失败:', result.error);
            }
        } catch (error) {
            console.error('加载NFT列表失败:', error);
        }
    }

    // 清空用户NFT
    clearUserNFTs() {
        this.userNFTs = [];
        this.renderNFTGrid();
    }

    // 渲染NFT网格
    renderNFTGrid() {
        const grid = document.getElementById('nftGrid');
        if (!grid) return;

        grid.innerHTML = '';

        if (this.userNFTs.length === 0) {
            grid.innerHTML = `
                <div style="grid-column: 1 / -1; text-align: center; padding: 40px; color: var(--text-secondary);">
                    <i class="fas fa-images" style="font-size: 48px; margin-bottom: 16px; opacity: 0.5;"></i>
                    <div>还没有NFT收藏</div>
                    <div style="font-size: 14px; margin-top: 8px;">开始游戏来获得你的第一个NFT！</div>
                </div>
            `;
            return;
        }

        this.userNFTs.forEach(nft => {
            const nftElement = this.createNFTElement(nft);
            grid.appendChild(nftElement);
        });
    }

    // 创建NFT元素
    createNFTElement(nft) {
        const element = document.createElement('div');
        element.className = 'nft-item';
        element.onclick = () => this.showNFTDetails(nft);

        const imageUrl = nft.metadata?.image || this.generateDefaultNFTImage(nft.token_id);
        const name = nft.metadata?.name || `NFT #${nft.token_id}`;

        element.innerHTML = `
            <div class="nft-item-image">
                <img src="${imageUrl}" alt="${name}" onerror="this.src='${this.generateDefaultNFTImage(nft.token_id)}'">
            </div>
            <div class="nft-item-info">
                <div class="nft-item-name">${name}</div>
                <div class="nft-item-id">#${nft.token_id}</div>
            </div>
        `;

        return element;
    }

    // 显示NFT详情
    async showNFTDetails(nft) {
        try {
            // 如果需要更详细的信息，可以调用API获取
            const result = await window.apiManager.getNFTDetails(nft.token_id);
            if (result.success) {
                this.showNFTModal(result.data);
            } else {
                // 使用现有数据显示
                this.showNFTModal(nft);
            }
        } catch (error) {
            console.error('获取NFT详情失败:', error);
            this.showNFTModal(nft);
        }
    }

    // 生成默认NFT图片
    generateDefaultNFTImage(tokenId) {
        const colors = [
            '#6366f1', '#8b5cf6', '#f59e0b', '#10b981', '#ef4444',
            '#06b6d4', '#84cc16', '#f97316', '#ec4899', '#8b5a2b'
        ];
        
        const color = colors[parseInt(tokenId) % colors.length];
        const shapes = ['circle', 'square', 'triangle', 'diamond', 'star'];
        const shape = shapes[parseInt(tokenId) % shapes.length];
        
        return `data:image/svg+xml,${encodeURIComponent(`
            <svg width="150" height="150" xmlns="http://www.w3.org/2000/svg">
                <defs>
                    <linearGradient id="grad${tokenId}" x1="0%" y1="0%" x2="100%" y2="100%">
                        <stop offset="0%" style="stop-color:${color};stop-opacity:1" />
                        <stop offset="100%" style="stop-color:${this.lightenColor(color, 20)};stop-opacity:1" />
                    </linearGradient>
                </defs>
                <rect width="150" height="150" fill="url(#grad${tokenId})"/>
                ${this.generateShape(shape, 75, 75, 30, '#ffffff')}
                <text x="75" y="120" text-anchor="middle" fill="#ffffff" font-family="Arial" font-size="12" font-weight="bold">
                    NFT #${tokenId}
                </text>
            </svg>
        `)}`;
    }

    // 生成形状
    generateShape(shape, x, y, size, color) {
        switch (shape) {
            case 'circle':
                return `<circle cx="${x}" cy="${y}" r="${size}" fill="${color}" opacity="0.8"/>`;
            case 'square':
                return `<rect x="${x - size}" y="${y - size}" width="${size * 2}" height="${size * 2}" fill="${color}" opacity="0.8"/>`;
            case 'triangle':
                return `<polygon points="${x},${y - size} ${x - size},${y + size} ${x + size},${y + size}" fill="${color}" opacity="0.8"/>`;
            case 'diamond':
                return `<polygon points="${x},${y - size} ${x + size},${y} ${x},${y + size} ${x - size},${y}" fill="${color}" opacity="0.8"/>`;
            case 'star':
                const points = [];
                for (let i = 0; i < 10; i++) {
                    const angle = (i * Math.PI) / 5;
                    const radius = i % 2 === 0 ? size : size / 2;
                    const px = x + radius * Math.cos(angle - Math.PI / 2);
                    const py = y + radius * Math.sin(angle - Math.PI / 2);
                    points.push(`${px},${py}`);
                }
                return `<polygon points="${points.join(' ')}" fill="${color}" opacity="0.8"/>`;
            default:
                return `<circle cx="${x}" cy="${y}" r="${size}" fill="${color}" opacity="0.8"/>`;
        }
    }

    // 颜色变亮
    lightenColor(color, percent) {
        const num = parseInt(color.replace("#", ""), 16);
        const amt = Math.round(2.55 * percent);
        const R = (num >> 16) + amt;
        const G = (num >> 8 & 0x00FF) + amt;
        const B = (num & 0x0000FF) + amt;
        return "#" + (0x1000000 + (R < 255 ? R < 1 ? 0 : R : 255) * 0x10000 +
            (G < 255 ? G < 1 ? 0 : G : 255) * 0x100 +
            (B < 255 ? B < 1 ? 0 : B : 255)).toString(16).slice(1);
    }

    // 获取稀有度文本
    getRarityText(rarity) {
        const rarityMap = {
            'common': '🔵 普通',
            'uncommon': '🟢 不常见',
            'rare': '🟡 稀有',
            'epic': '🟣 史诗',
            'legendary': '🟠 传说',
            'mythic': '🔴 神话'
        };
        
        return rarityMap[rarity.toLowerCase()] || '🔵 普通';
    }

    // 播放显示动画
    playShowAnimation() {
        const modal = document.getElementById('nftDisplay');
        const card = modal?.querySelector('.nft-card');
        
        if (card) {
            card.style.animation = 'none';
            setTimeout(() => {
                card.style.animation = 'slideIn 0.5s ease-out';
            }, 10);
        }
    }

    // 获取NFT在OpenSea的链接
    getOpenSeaLink(nft) {
        if (!nft.contract_address || !nft.token_id) {
            return null;
        }
        
        return window.apiManager.getOpenSeaURL(nft.contract_address, nft.token_id);
    }

    // 分享NFT
    shareNFT(nft) {
        if (navigator.share) {
            navigator.share({
                title: nft.metadata?.name || `NFT #${nft.token_id}`,
                text: '看看我在NFT Capture Game中获得的NFT！',
                url: this.getOpenSeaLink(nft) || window.location.href
            });
        } else {
            // 复制链接到剪贴板
            const link = this.getOpenSeaLink(nft) || window.location.href;
            navigator.clipboard.writeText(link).then(() => {
                this.showNotification('链接已复制到剪贴板', 'success');
            });
        }
    }

    // 下载NFT图片
    async downloadNFTImage(nft) {
        try {
            const imageUrl = nft.metadata?.image;
            if (!imageUrl) {
                this.showNotification('NFT图片不可用', 'error');
                return;
            }

            const response = await fetch(imageUrl);
            const blob = await response.blob();
            const url = window.URL.createObjectURL(blob);
            
            const a = document.createElement('a');
            a.href = url;
            a.download = `${nft.metadata?.name || 'NFT'}_${nft.token_id}.png`;
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
            
            window.URL.revokeObjectURL(url);
            this.showNotification('图片下载成功', 'success');
        } catch (error) {
            console.error('下载图片失败:', error);
            this.showNotification('下载图片失败', 'error');
        }
    }

    // 显示通知
    showNotification(message, type = 'info') {
        if (window.walletManager) {
            window.walletManager.showNotification(message, type);
        }
    }

    // 获取用户NFT数量
    getUserNFTCount() {
        return this.userNFTs.length;
    }

    // 获取最新NFT
    getLatestNFT() {
        if (this.userNFTs.length === 0) return null;
        
        return this.userNFTs.reduce((latest, current) => {
            const latestTime = new Date(latest.created_at || 0);
            const currentTime = new Date(current.created_at || 0);
            return currentTime > latestTime ? current : latest;
        });
    }

    // 按稀有度过滤NFT
    filterByRarity(rarity) {
        return this.userNFTs.filter(nft => 
            (nft.metadata?.rarity || 'common').toLowerCase() === rarity.toLowerCase()
        );
    }

    // 搜索NFT
    searchNFTs(query) {
        const lowerQuery = query.toLowerCase();
        return this.userNFTs.filter(nft => {
            const name = (nft.metadata?.name || '').toLowerCase();
            const tokenId = nft.token_id.toString();
            const description = (nft.metadata?.description || '').toLowerCase();
            
            return name.includes(lowerQuery) || 
                   tokenId.includes(lowerQuery) || 
                   description.includes(lowerQuery);
        });
    }
}

// 创建全局NFT管理器实例
window.nftManager = new NFTManager();