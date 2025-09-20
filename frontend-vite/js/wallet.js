// 钱包管理类
class WalletManager {
    constructor() {
        this.isConnected = false;
        this.walletAddress = null;
        this.web3 = null;
        this.init();
    }

    // 初始化钱包管理器
    init() {
        this.checkConnection();
        this.bindEvents();
    }

    // 绑定事件
    bindEvents() {
        const connectBtn = document.getElementById('connectWallet');
        const disconnectBtn = document.getElementById('disconnectWallet');

        if (connectBtn) {
            connectBtn.addEventListener('click', () => this.connectWallet());
        }

        if (disconnectBtn) {
            disconnectBtn.addEventListener('click', () => this.disconnectWallet());
        }

        // 监听账户变化
        if (window.ethereum) {
            window.ethereum.on('accountsChanged', (accounts) => {
                if (accounts.length === 0) {
                    this.disconnectWallet();
                } else {
                    this.handleAccountChange(accounts[0]);
                }
            });

            window.ethereum.on('chainChanged', (chainId) => {
                window.location.reload();
            });
        }
    }

    // 检查是否已连接
    async checkConnection() {
        if (typeof window.ethereum !== 'undefined') {
            try {
                const accounts = await window.ethereum.request({ method: 'eth_accounts' });
                if (accounts.length > 0) {
                    this.walletAddress = accounts[0];
                    this.isConnected = true;
                    this.updateUI();
                    
                    // 检查网络
                    await this.checkNetwork();
                }
            } catch (error) {
                console.error('检查连接失败:', error);
            }
        }
    }

    // 连接钱包
    async connectWallet() {
        if (typeof window.ethereum === 'undefined') {
            this.showNotification('请安装MetaMask钱包', 'error');
            return false;
        }

        try {
            this.showLoading('连接钱包中...');
            
            const accounts = await window.ethereum.request({ 
                method: 'eth_requestAccounts' 
            });

            if (accounts.length > 0) {
                this.walletAddress = accounts[0];
                this.isConnected = true;
                
                // 检查网络
                await this.checkNetwork();
                
                this.updateUI();
                this.showNotification('钱包连接成功！', 'success');
                
                // 触发连接成功事件
                this.onWalletConnected();
                
                return true;
            }
        } catch (error) {
            console.error('连接钱包失败:', error);
            this.showNotification('连接钱包失败: ' + error.message, 'error');
            return false;
        } finally {
            this.hideLoading();
        }
    }

    // 断开钱包连接
    disconnectWallet() {
        this.isConnected = false;
        this.walletAddress = null;
        this.updateUI();
        this.showNotification('钱包已断开连接', 'warning');
        
        // 触发断开连接事件
        this.onWalletDisconnected();
    }

    // 处理账户变化
    handleAccountChange(newAccount) {
        this.walletAddress = newAccount;
        this.updateUI();
        this.showNotification('账户已切换', 'success');
        
        // 触发账户变化事件
        this.onAccountChanged(newAccount);
    }

    // 检查网络
    async checkNetwork() {
        try {
            const chainId = await window.ethereum.request({ method: 'eth_chainId' });
            const currentChain = parseInt(chainId, 16);
            
            // 这里可以根据需要检查特定网络
            // 例如：以太坊主网 (1)、Polygon (137)、BSC (56) 等
            console.log('当前网络 Chain ID:', currentChain);
            
            return currentChain;
        } catch (error) {
            console.error('检查网络失败:', error);
            return null;
        }
    }

    // 切换网络
    async switchNetwork(chainId) {
        try {
            await window.ethereum.request({
                method: 'wallet_switchEthereumChain',
                params: [{ chainId: `0x${chainId.toString(16)}` }],
            });
            return true;
        } catch (error) {
            console.error('切换网络失败:', error);
            this.showNotification('切换网络失败: ' + error.message, 'error');
            return false;
        }
    }

    // 获取余额
    async getBalance() {
        if (!this.isConnected || !this.walletAddress) {
            return null;
        }

        try {
            const balance = await window.ethereum.request({
                method: 'eth_getBalance',
                params: [this.walletAddress, 'latest']
            });
            
            // 转换为 ETH
            const ethBalance = parseInt(balance, 16) / Math.pow(10, 18);
            return ethBalance;
        } catch (error) {
            console.error('获取余额失败:', error);
            return null;
        }
    }

    // 更新UI
    updateUI() {
        const connectBtn = document.getElementById('connectWallet');
        const walletInfo = document.getElementById('walletInfo');
        const walletAddress = document.getElementById('walletAddress');
        const startBtn = document.getElementById('startGame');

        if (this.isConnected && this.walletAddress) {
            // 显示钱包信息
            if (connectBtn) connectBtn.classList.add('hidden');
            if (walletInfo) walletInfo.classList.remove('hidden');
            if (walletAddress) {
                walletAddress.textContent = this.formatAddress(this.walletAddress);
            }
            
            // 启用开始游戏按钮
            if (startBtn) {
                startBtn.disabled = false;
                startBtn.innerHTML = '<i class="fas fa-play"></i> 开始游戏';
            }
        } else {
            // 隐藏钱包信息
            if (connectBtn) connectBtn.classList.remove('hidden');
            if (walletInfo) walletInfo.classList.add('hidden');
            
            // 禁用开始游戏按钮
            if (startBtn) {
                startBtn.disabled = true;
                startBtn.innerHTML = '<i class="fas fa-wallet"></i> 请先连接钱包';
            }
        }
    }

    // 格式化地址显示
    formatAddress(address) {
        if (!address) return '';
        return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`;
    }

    // 显示加载状态
    showLoading(message = '处理中...') {
        const overlay = document.getElementById('loadingOverlay');
        const text = overlay?.querySelector('.loading-text');
        
        if (overlay) {
            overlay.classList.remove('hidden');
            if (text) text.textContent = message;
        }
    }

    // 隐藏加载状态
    hideLoading() {
        const overlay = document.getElementById('loadingOverlay');
        if (overlay) {
            overlay.classList.add('hidden');
        }
    }

    // 显示通知
    showNotification(message, type = 'info') {
        const container = document.getElementById('notifications');
        if (!container) return;

        const notification = document.createElement('div');
        notification.className = `notification ${type}`;
        notification.innerHTML = `
            <div style="display: flex; align-items: center; justify-content: space-between;">
                <span>${message}</span>
                <button onclick="this.parentElement.parentElement.remove()" style="background: none; border: none; color: inherit; cursor: pointer; padding: 0; margin-left: 10px;">
                    <i class="fas fa-times"></i>
                </button>
            </div>
        `;

        container.appendChild(notification);

        // 自动移除通知
        setTimeout(() => {
            if (notification.parentNode) {
                notification.remove();
            }
        }, 5000);
    }

    // 钱包连接成功回调
    onWalletConnected() {
        // 可以在这里添加连接成功后的逻辑
        console.log('钱包连接成功:', this.walletAddress);
        
        // 触发自定义事件
        window.dispatchEvent(new CustomEvent('walletConnected', {
            detail: { address: this.walletAddress }
        }));
    }

    // 钱包断开连接回调
    onWalletDisconnected() {
        // 可以在这里添加断开连接后的逻辑
        console.log('钱包已断开连接');
        
        // 触发自定义事件
        window.dispatchEvent(new CustomEvent('walletDisconnected'));
    }

    // 账户变化回调
    onAccountChanged(newAccount) {
        console.log('账户已切换:', newAccount);
        
        // 触发自定义事件
        window.dispatchEvent(new CustomEvent('accountChanged', {
            detail: { address: newAccount }
        }));
    }

    // 获取当前钱包地址
    getAddress() {
        return this.walletAddress;
    }

    // 检查是否已连接
    getIsConnected() {
        return this.isConnected;
    }
}

// 创建全局钱包管理器实例
window.walletManager = new WalletManager();