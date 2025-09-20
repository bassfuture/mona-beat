// 主应用程序入口
class App {
    constructor() {
        this.isInitialized = false;
        this.init();
    }

    // 初始化应用程序
    async init() {
        try {
            console.log('🚀 NFT Capture Game 正在启动...');
            
            // 等待DOM加载完成
            if (document.readyState === 'loading') {
                document.addEventListener('DOMContentLoaded', () => this.startApp());
            } else {
                this.startApp();
            }
        } catch (error) {
            console.error('应用程序初始化失败:', error);
            this.showError('应用程序启动失败，请刷新页面重试');
        }
    }

    // 启动应用程序
    async startApp() {
        try {
            console.log('📱 初始化应用组件...');
            
            // 检查浏览器兼容性
            this.checkBrowserCompatibility();
            
            // 初始化粒子背景
            this.initParticles();
            
            // 检查服务器连接
            await this.checkServerConnection();
            
            // 设置全局错误处理
            this.setupErrorHandling();
            
            // 设置键盘快捷键
            this.setupKeyboardShortcuts();
            
            // 设置性能监控
            this.setupPerformanceMonitoring();
            
            // 标记为已初始化
            this.isInitialized = true;
            
            console.log('✅ NFT Capture Game 启动成功！');
            this.showWelcomeMessage();
            
        } catch (error) {
            console.error('应用程序启动失败:', error);
            this.showError('应用程序启动失败: ' + error.message);
        }
    }

    // 检查浏览器兼容性
    checkBrowserCompatibility() {
        const requiredFeatures = [
            'fetch',
            'Promise',
            'localStorage',
            'addEventListener'
        ];

        const missingFeatures = requiredFeatures.filter(feature => 
            typeof window[feature] === 'undefined'
        );

        if (missingFeatures.length > 0) {
            throw new Error(`浏览器不支持以下功能: ${missingFeatures.join(', ')}`);
        }

        // 检查Web3支持
        if (typeof window.ethereum === 'undefined') {
            this.showNotification('未检测到Web3钱包，请安装MetaMask等钱包扩展', 'warning');
        }

        console.log('✅ 浏览器兼容性检查通过');
    }

    // 初始化粒子背景
    initParticles() {
        try {
            if (typeof particlesJS !== 'undefined') {
                console.log('🌟 初始化粒子背景...');
                // 粒子配置在 particles-config.js 中
            } else {
                console.warn('⚠️ 粒子库未加载，跳过粒子背景');
            }
        } catch (error) {
            console.warn('粒子背景初始化失败:', error);
        }
    }

    // 检查服务器连接
    async checkServerConnection() {
        try {
            console.log('🔗 检查服务器连接...');
            
            const isConnected = await window.apiManager.checkServerStatus();
            if (isConnected) {
                console.log('✅ 服务器连接正常');
                this.showNotification('服务器连接正常', 'success');
            } else {
                console.warn('⚠️ 服务器连接异常');
                this.showNotification('服务器连接异常，部分功能可能不可用', 'warning');
            }
        } catch (error) {
            console.error('服务器连接检查失败:', error);
            this.showNotification('无法连接到服务器，请检查网络连接', 'error');
        }
    }

    // 设置全局错误处理
    setupErrorHandling() {
        // 捕获未处理的Promise错误
        window.addEventListener('unhandledrejection', (event) => {
            console.error('未处理的Promise错误:', event.reason);
            this.showNotification('发生了一个错误，请重试', 'error');
            event.preventDefault();
        });

        // 捕获JavaScript错误
        window.addEventListener('error', (event) => {
            console.error('JavaScript错误:', event.error);
            this.showNotification('页面发生错误，请刷新重试', 'error');
        });

        console.log('✅ 全局错误处理已设置');
    }

    // 设置键盘快捷键
    setupKeyboardShortcuts() {
        document.addEventListener('keydown', (event) => {
            // 防止在输入框中触发快捷键
            if (event.target.tagName === 'INPUT' || event.target.tagName === 'TEXTAREA') {
                return;
            }

            switch (event.key) {
                case ' ': // 空格键 - 敲击
                    event.preventDefault();
                    if (window.gameManager && window.gameManager.isGameActive) {
                        const tapTarget = document.getElementById('tapTarget');
                        if (tapTarget) {
                            tapTarget.click();
                        }
                    }
                    break;
                    
                case 'Enter': // 回车键 - 开始游戏
                    event.preventDefault();
                    const startBtn = document.getElementById('startGame');
                    if (startBtn && !startBtn.disabled) {
                        startBtn.click();
                    }
                    break;
                    
                case 'Escape': // ESC键 - 关闭模态框
                    event.preventDefault();
                    if (window.nftManager) {
                        window.nftManager.hideNFTModal();
                    }
                    break;
                    
                case 'c': // C键 - 连接钱包
                    if (event.ctrlKey || event.metaKey) return; // 避免与复制冲突
                    event.preventDefault();
                    const connectBtn = document.getElementById('connectWallet');
                    if (connectBtn && !connectBtn.classList.contains('hidden')) {
                        connectBtn.click();
                    }
                    break;
            }
        });

        console.log('⌨️ 键盘快捷键已设置');
        this.showKeyboardHelp();
    }

    // 显示键盘快捷键帮助
    showKeyboardHelp() {
        const helpText = `
            键盘快捷键：
            • 空格键：敲击目标
            • 回车键：开始游戏
            • ESC键：关闭弹窗
            • C键：连接钱包
        `;
        
        console.log(helpText);
    }

    // 设置性能监控
    setupPerformanceMonitoring() {
        // 监控页面加载性能
        window.addEventListener('load', () => {
            setTimeout(() => {
                const perfData = performance.getEntriesByType('navigation')[0];
                if (perfData) {
                    const loadTime = perfData.loadEventEnd - perfData.loadEventStart;
                    console.log(`📊 页面加载时间: ${loadTime.toFixed(2)}ms`);
                    
                    if (loadTime > 3000) {
                        console.warn('⚠️ 页面加载较慢，可能影响用户体验');
                    }
                }
            }, 0);
        });

        // 监控内存使用
        if ('memory' in performance) {
            setInterval(() => {
                const memory = performance.memory;
                const usedMB = (memory.usedJSHeapSize / 1024 / 1024).toFixed(2);
                const totalMB = (memory.totalJSHeapSize / 1024 / 1024).toFixed(2);
                
                console.log(`💾 内存使用: ${usedMB}MB / ${totalMB}MB`);
                
                // 内存使用过高警告
                if (memory.usedJSHeapSize / memory.jsHeapSizeLimit > 0.9) {
                    console.warn('⚠️ 内存使用过高，建议刷新页面');
                }
            }, 30000); // 每30秒检查一次
        }

        console.log('📊 性能监控已启用');
    }

    // 显示欢迎消息
    showWelcomeMessage() {
        const messages = [
            '🎮 欢迎来到NFT Capture Game！',
            '💎 连接钱包开始你的NFT收集之旅',
            '🚀 每次敲击都可能获得珍贵的NFT',
            '🏆 收集更多NFT，成为顶级收藏家！'
        ];

        let index = 0;
        const showMessage = () => {
            if (index < messages.length) {
                this.showNotification(messages[index], 'info');
                index++;
                setTimeout(showMessage, 2000);
            }
        };

        setTimeout(showMessage, 1000);
    }

    // 显示错误信息
    showError(message) {
        const errorContainer = document.createElement('div');
        errorContainer.style.cssText = `
            position: fixed;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
            background: #ef4444;
            color: white;
            padding: 20px;
            border-radius: 10px;
            z-index: 9999;
            text-align: center;
            box-shadow: 0 10px 25px rgba(0,0,0,0.5);
        `;
        errorContainer.innerHTML = `
            <h3>❌ 错误</h3>
            <p>${message}</p>
            <button onclick="location.reload()" style="
                background: white;
                color: #ef4444;
                border: none;
                padding: 10px 20px;
                border-radius: 5px;
                margin-top: 10px;
                cursor: pointer;
            ">刷新页面</button>
        `;
        
        document.body.appendChild(errorContainer);
    }

    // 显示通知
    showNotification(message, type = 'info') {
        if (window.walletManager) {
            window.walletManager.showNotification(message, type);
        } else {
            console.log(`[${type.toUpperCase()}] ${message}`);
        }
    }

    // 获取应用状态
    getAppStatus() {
        return {
            initialized: this.isInitialized,
            walletConnected: window.walletManager?.getIsConnected() || false,
            gameActive: window.gameManager?.isGameActive || false,
            nftCount: window.nftManager?.getUserNFTCount() || 0
        };
    }

    // 重启应用
    restart() {
        console.log('🔄 重启应用程序...');
        location.reload();
    }

    // 清理资源
    cleanup() {
        console.log('🧹 清理应用资源...');
        
        // 清理事件监听器
        window.removeEventListener('unhandledrejection', this.handleUnhandledRejection);
        window.removeEventListener('error', this.handleError);
        
        // 清理定时器
        // 这里可以添加清理定时器的代码
        
        console.log('✅ 资源清理完成');
    }
}

// 应用程序启动
console.log('🎯 NFT Capture Game - 区块链NFT捕获游戏');
console.log('🔧 版本: 1.0.0');
console.log('👨‍💻 开发者: AI Assistant');

// 创建应用实例
window.app = new App();

// 导出到全局作用域以便调试
window.debug = {
    app: window.app,
    walletManager: window.walletManager,
    gameManager: window.gameManager,
    nftManager: window.nftManager,
    apiManager: window.apiManager,
    
    // 调试方法
    getStatus: () => window.app.getAppStatus(),
    restart: () => window.app.restart(),
    cleanup: () => window.app.cleanup(),
    
    // 测试方法
    testNotification: (message, type) => window.app.showNotification(message, type),
    testAPI: () => window.apiManager.checkServerStatus(),
    testWallet: () => window.walletManager.connectWallet()
};

console.log('🔍 调试工具已加载，使用 window.debug 访问');
console.log('📖 使用说明：');
console.log('  - window.debug.getStatus() - 获取应用状态');
console.log('  - window.debug.testNotification("消息", "类型") - 测试通知');
console.log('  - window.debug.testAPI() - 测试API连接');
console.log('  - window.debug.testWallet() - 测试钱包连接');