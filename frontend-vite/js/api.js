// API管理类
class APIManager {
    constructor() {
        this.baseURL = '/api'; // 通过 Vite 代理到后端，并在代理中重写为 /api/v1
        this.endpoints = {
            startGame: '/captures/attempt',
            tap: '/captures/attempt',
            getUserInfo: '/users/{address}',
            getUserNFTs: '/nfts/{address}',
            getNFTDetails: '/nfts/token/{tokenId}',
            health: '/health',
            nftStats: '/nfts/stats',
            recentNFTs: '/nfts/recent',
            userStats: '/users/{address}/stats',
            leaderboard: '/users/leaderboard'
        };
    }

    // 通用请求方法
    async request(endpoint, options = {}) {
        const url = `${this.baseURL}${endpoint}`;
        
        const defaultOptions = {
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            }
        };

        const finalOptions = {
            ...defaultOptions,
            ...options,
            headers: {
                ...defaultOptions.headers,
                ...options.headers
            }
        };

        try {
            console.log(`API请求: ${options.method || 'GET'} ${url}`);
            
            const response = await fetch(url, finalOptions);
            
            if (!response.ok) {
                const errorText = await response.text();
                throw new Error(`HTTP ${response.status}: ${errorText}`);
            }

            const data = await response.json();
            console.log('API响应:', data);
            
            return {
                success: true,
                data: data
            };
        } catch (error) {
            console.error('API请求失败:', error);
            return {
                success: false,
                error: error.message
            };
        }
    }

    // 开始游戏
    async startGame(walletAddress, targetType = 'common', difficulty = 1) {
        if (!walletAddress) {
            return {
                success: false,
                error: '钱包地址不能为空'
            };
        }

        return await this.request(this.endpoints.startGame, {
            method: 'POST',
            body: JSON.stringify({
                wallet_address: walletAddress,
                target_type: targetType,
                difficulty: difficulty
            })
        });
    }

    // 敲击操作
    async tap(walletAddress, targetType = 'common', difficulty = 1) {
        if (!walletAddress) {
            return {
                success: false,
                error: '钱包地址不能为空'
            };
        }

        return await this.request(this.endpoints.tap, {
            method: 'POST',
            body: JSON.stringify({
                wallet_address: walletAddress,
                target_type: targetType,
                difficulty: difficulty
            })
        });
    }

    // 获取用户信息
    async getUserInfo(walletAddress) {
        if (!walletAddress) {
            return {
                success: false,
                error: '钱包地址不能为空'
            };
        }

        const endpoint = this.endpoints.getUserInfo.replace('{address}', walletAddress);
        return await this.request(endpoint, {
            method: 'GET'
        });
    }

    // 获取用户NFT列表
    async getUserNFTs(walletAddress) {
        if (!walletAddress) {
            return {
                success: false,
                error: '钱包地址不能为空'
            };
        }

        const endpoint = this.endpoints.getUserNFTs.replace('{address}', walletAddress);
        return await this.request(endpoint, {
            method: 'GET'
        });
    }

    // 获取NFT详情
    async getNFTDetails(tokenId) {
        if (!tokenId) {
            return {
                success: false,
                error: 'Token ID不能为空'
            };
        }

        const endpoint = this.endpoints.getNFTDetails.replace('{tokenId}', tokenId);
        return await this.request(endpoint, {
            method: 'GET'
        });
    }

    // 批量获取NFT详情
    async getBatchNFTDetails(tokenIds) {
        if (!tokenIds || tokenIds.length === 0) {
            return {
                success: true,
                data: []
            };
        }

        const promises = tokenIds.map(tokenId => this.getNFTDetails(tokenId));
        const results = await Promise.allSettled(promises);
        
        const nfts = [];
        results.forEach((result, index) => {
            if (result.status === 'fulfilled' && result.value.success) {
                nfts.push(result.value.data);
            } else {
                console.error(`获取NFT ${tokenIds[index]} 详情失败:`, result.reason);
            }
        });

        return {
            success: true,
            data: nfts
        };
    }

    // 检查服务器状态
    async checkServerStatus() {
        try {
            const response = await fetch(`${this.baseURL}/health`, {
                method: 'GET'
            });
            
            return response.ok;
        } catch (error) {
            console.error('服务器状态检查失败:', error);
            return false;
        }
    }

    // 获取区块链交易详情
    getTransactionURL(txHash, network = 'ethereum') {
        const explorers = {
            ethereum: 'https://etherscan.io/tx/',
            polygon: 'https://polygonscan.com/tx/',
            bsc: 'https://bscscan.com/tx/',
            sepolia: 'https://sepolia.etherscan.io/tx/'
        };

        return explorers[network] + txHash;
    }

    // 获取NFT在OpenSea的链接
    getOpenSeaURL(contractAddress, tokenId, network = 'ethereum') {
        const baseURLs = {
            ethereum: 'https://opensea.io/assets/ethereum/',
            polygon: 'https://opensea.io/assets/matic/',
            sepolia: 'https://testnets.opensea.io/assets/sepolia/'
        };

        return `${baseURLs[network]}${contractAddress}/${tokenId}`;
    }

    // 格式化错误消息
    formatError(error) {
        if (typeof error === 'string') {
            return error;
        }

        if (error.message) {
            return error.message;
        }

        return '未知错误';
    }

    // 重试机制
    async requestWithRetry(endpoint, options = {}, maxRetries = 3) {
        let lastError;
        
        for (let i = 0; i < maxRetries; i++) {
            try {
                const result = await this.request(endpoint, options);
                if (result.success) {
                    return result;
                }
                lastError = result.error;
            } catch (error) {
                lastError = error.message;
                
                // 如果不是最后一次重试，等待一段时间
                if (i < maxRetries - 1) {
                    await new Promise(resolve => setTimeout(resolve, 1000 * (i + 1)));
                }
            }
        }

        return {
            success: false,
            error: `请求失败，已重试${maxRetries}次: ${lastError}`
        };
    }
}

// 创建全局API管理器实例
window.apiManager = new APIManager();