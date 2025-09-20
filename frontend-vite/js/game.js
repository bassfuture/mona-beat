// 游戏管理类
class GameManager {
    constructor() {
        this.isGameActive = false;
        this.gameData = null;
        this.tapCount = 0;
        this.maxTaps = 0;
        this.successRate = 0;
        this.totalNFTs = 0;
        this.isProcessing = false;
        
        // 游戏配置
        this.gameConfig = {
            targetType: 'common', // 默认目标类型
            difficulty: 1         // 默认难度
        };
        
        this.init();
    }

    // 初始化游戏
    init() {
        this.bindEvents();
        this.setupAnimations();
    }

    // 绑定事件
    bindEvents() {
        const startBtn = document.getElementById('startGame');
        const tapTarget = document.getElementById('tapTarget');
        const continueBtn = document.getElementById('continueGame');

        if (startBtn) {
            startBtn.addEventListener('click', () => this.startGame());
        }

        if (tapTarget) {
            tapTarget.addEventListener('click', (e) => this.handleTap(e));
        }

        if (continueBtn) {
            continueBtn.addEventListener('click', () => this.continueGame());
        }

        // 监听钱包连接事件
        window.addEventListener('walletConnected', () => {
            this.onWalletConnected();
        });

        window.addEventListener('walletDisconnected', () => {
            this.onWalletDisconnected();
        });
    }

    // 设置动画
    setupAnimations() {
        // 可以在这里添加初始动画设置
    }

    // 开始游戏
    async startGame() {
        if (!window.walletManager.getIsConnected()) {
            this.showNotification('请先连接钱包', 'error');
            return;
        }

        if (this.isProcessing) {
            return;
        }

        this.isProcessing = true;
        this.showLoading('初始化游戏...');

        try {
            const walletAddress = window.walletManager.getAddress();
            const result = await window.apiManager.startGame(
                walletAddress, 
                this.gameConfig.targetType, 
                this.gameConfig.difficulty
            );

            if (result.success) {
                this.gameData = result.data;
                this.maxTaps = this.gameData.remaining_taps || 10;
                this.tapCount = 0;
                this.isGameActive = true;

                this.showGameScreen();
                this.updateGameStats();
                this.showNotification('游戏开始！点击目标来捕获NFT', 'success');
            } else {
                this.showNotification('游戏启动失败: ' + result.error, 'error');
            }
        } catch (error) {
            console.error('启动游戏失败:', error);
            this.showNotification('游戏启动失败: ' + error.message, 'error');
        } finally {
            this.isProcessing = false;
            this.hideLoading();
        }
    }

    // 处理敲击
    async handleTap(event) {
        if (!this.isGameActive || this.isProcessing) {
            return;
        }

        if (this.tapCount >= this.maxTaps) {
            this.showNotification('敲击次数已用完，请重新开始游戏', 'warning');
            return;
        }

        this.isProcessing = true;
        
        // 播放敲击动画
        this.playTapAnimation(event);
        
        // 增加敲击计数
        this.tapCount++;
        this.updateProgress();

        try {
            const walletAddress = window.walletManager.getAddress();
            const result = await window.apiManager.tap(
                walletAddress, 
                this.gameConfig.targetType, 
                this.gameConfig.difficulty
            );

            if (result.success) {
                const tapResult = result.data;
                
                if (tapResult.nft_generated) {
                    // 生成了NFT
                    this.handleNFTGenerated(tapResult);
                } else {
                    // 没有生成NFT
                    this.handleNormalTap(tapResult);
                }

                // 更新游戏统计
                this.updateGameStats();
                
                // 检查游戏是否结束
                if (this.tapCount >= this.maxTaps) {
                    this.endGame();
                }
            } else {
                this.showNotification('敲击失败: ' + result.error, 'error');
                this.tapCount--; // 回滚敲击计数
            }
        } catch (error) {
            console.error('敲击处理失败:', error);
            this.showNotification('敲击处理失败: ' + error.message, 'error');
            this.tapCount--; // 回滚敲击计数
        } finally {
            this.isProcessing = false;
            this.updateProgress();
        }
    }

    // 播放敲击动画
    playTapAnimation(event) {
        const target = event.currentTarget;
        const ripple = target.querySelector('.tap-ripple');
        const effectsContainer = document.getElementById('tapEffects');

        // 波纹效果
        if (ripple) {
            ripple.classList.remove('animate');
            setTimeout(() => {
                ripple.classList.add('animate');
            }, 10);
        }

        // 点击效果文字
        if (effectsContainer) {
            const effect = document.createElement('div');
            effect.className = 'tap-effect';
            effect.style.left = Math.random() * 100 + '%';
            effect.style.top = Math.random() * 100 + '%';
            effect.style.color = this.getRandomColor();
            effect.textContent = this.getRandomTapText();
            
            effectsContainer.appendChild(effect);
            
            // 移除效果元素
            setTimeout(() => {
                if (effect.parentNode) {
                    effect.remove();
                }
            }, 1000);
        }

        // 目标缩放效果
        target.style.transform = 'scale(0.95)';
        setTimeout(() => {
            target.style.transform = 'scale(1)';
        }, 150);
    }

    // 获取随机颜色
    getRandomColor() {
        const colors = ['#6366f1', '#8b5cf6', '#f59e0b', '#10b981', '#ef4444'];
        return colors[Math.floor(Math.random() * colors.length)];
    }

    // 获取随机敲击文字
    getRandomTapText() {
        const texts = ['💎', '⭐', '🔥', '💫', '✨', '+1', 'Nice!', 'Good!'];
        return texts[Math.floor(Math.random() * texts.length)];
    }

    // 处理NFT生成
    handleNFTGenerated(tapResult) {
        this.totalNFTs++;
        this.showNotification('🎉 恭喜！你获得了一个NFT！', 'success');
        
        // 显示NFT详情
        window.nftManager.showNFTModal(tapResult);
        
        // 播放庆祝动画
        this.playCelebrationAnimation();
        
        // 更新NFT收藏
        this.refreshNFTCollection();
    }

    // 处理普通敲击
    handleNormalTap(tapResult) {
        const messages = [
            '继续努力！',
            '再试一次！',
            '差一点就成功了！',
            '不要放弃！',
            '下次一定行！'
        ];
        
        const message = messages[Math.floor(Math.random() * messages.length)];
        this.showNotification(message, 'info');
    }

    // 播放庆祝动画
    playCelebrationAnimation() {
        // 创建庆祝粒子效果
        const container = document.querySelector('.game-area');
        if (!container) return;

        for (let i = 0; i < 20; i++) {
            const particle = document.createElement('div');
            particle.style.position = 'absolute';
            particle.style.left = Math.random() * 100 + '%';
            particle.style.top = Math.random() * 100 + '%';
            particle.style.fontSize = '20px';
            particle.style.pointerEvents = 'none';
            particle.style.zIndex = '1000';
            particle.textContent = ['🎉', '🎊', '✨', '💎', '🏆'][Math.floor(Math.random() * 5)];
            
            container.appendChild(particle);
            
            // 动画
            particle.animate([
                { transform: 'translateY(0) scale(1)', opacity: 1 },
                { transform: 'translateY(-100px) scale(1.5)', opacity: 0 }
            ], {
                duration: 2000,
                easing: 'ease-out'
            }).onfinish = () => {
                if (particle.parentNode) {
                    particle.remove();
                }
            };
        }
    }

    // 显示游戏界面
    showGameScreen() {
        const startScreen = document.getElementById('startScreen');
        const gameScreen = document.getElementById('gameScreen');

        if (startScreen) startScreen.classList.add('hidden');
        if (gameScreen) gameScreen.classList.remove('hidden');
    }

    // 显示开始界面
    showStartScreen() {
        const startScreen = document.getElementById('startScreen');
        const gameScreen = document.getElementById('gameScreen');

        if (startScreen) startScreen.classList.remove('hidden');
        if (gameScreen) gameScreen.classList.add('hidden');
    }

    // 更新游戏统计
    updateGameStats() {
        const remainingTaps = document.getElementById('remainingTaps');
        const totalNFTsElement = document.getElementById('totalNFTs');
        const successRateElement = document.getElementById('successRate');

        if (remainingTaps) {
            remainingTaps.textContent = Math.max(0, this.maxTaps - this.tapCount);
        }

        if (totalNFTsElement) {
            totalNFTsElement.textContent = this.totalNFTs;
        }

        if (successRateElement) {
            const rate = this.tapCount > 0 ? ((this.totalNFTs / this.tapCount) * 100).toFixed(1) : 0;
            successRateElement.textContent = rate + '%';
        }
    }

    // 更新进度条
    updateProgress() {
        const progressFill = document.getElementById('progressFill');
        if (progressFill) {
            const progress = (this.tapCount / this.maxTaps) * 100;
            progressFill.style.width = progress + '%';
        }
    }

    // 继续游戏
    continueGame() {
        window.nftManager.hideNFTModal();
        
        if (this.tapCount >= this.maxTaps) {
            this.endGame();
        }
    }

    // 结束游戏
    endGame() {
        this.isGameActive = false;
        this.showStartScreen();
        
        const finalMessage = this.totalNFTs > 0 
            ? `游戏结束！你总共获得了 ${this.totalNFTs} 个NFT！`
            : '游戏结束！这次没有获得NFT，再试一次吧！';
            
        this.showNotification(finalMessage, this.totalNFTs > 0 ? 'success' : 'info');
        
        // 重置游戏状态
        this.resetGame();
    }

    // 重置游戏
    resetGame() {
        this.tapCount = 0;
        this.maxTaps = 0;
        this.gameData = null;
        this.updateProgress();
    }

    // 刷新NFT收藏
    async refreshNFTCollection() {
        if (window.nftManager) {
            await window.nftManager.loadUserNFTs();
        }
    }

    // 钱包连接成功回调
    onWalletConnected() {
        this.loadUserData();
    }

    // 钱包断开连接回调
    onWalletDisconnected() {
        this.isGameActive = false;
        this.showStartScreen();
        this.resetGame();
        this.totalNFTs = 0;
        this.updateGameStats();
    }

    // 加载用户数据
    async loadUserData() {
        try {
            const walletAddress = window.walletManager.getAddress();
            if (!walletAddress) return;

            const result = await window.apiManager.getUserInfo(walletAddress);
            if (result.success) {
                const userInfo = result.data;
                this.totalNFTs = userInfo.total_nfts || 0;
                this.updateGameStats();
            }
        } catch (error) {
            console.error('加载用户数据失败:', error);
        }
    }

    // 显示通知
    showNotification(message, type = 'info') {
        if (window.walletManager) {
            window.walletManager.showNotification(message, type);
        }
    }

    // 显示加载状态
    showLoading(message = '处理中...') {
        if (window.walletManager) {
            window.walletManager.showLoading(message);
        }
    }

    // 隐藏加载状态
    hideLoading() {
        if (window.walletManager) {
            window.walletManager.hideLoading();
        }
    }
}

// 创建全局游戏管理器实例
window.gameManager = new GameManager();