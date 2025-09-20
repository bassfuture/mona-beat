// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/**
 * @title CaptureNFT
 * @dev NFT捕捉游戏的主要合约，实现NFT的铸造、转移和游戏逻辑
 */
contract CaptureNFT is ERC721, ERC721URIStorage, ERC721Enumerable, Ownable, ReentrancyGuard {
    // 代币ID计数器
    uint256 private _nextTokenId = 1;

    // 稀有度枚举
    enum Rarity {
        COMMON,     // 普通 - 0
        UNCOMMON,   // 不常见 - 1
        RARE,       // 稀有 - 2
        EPIC,       // 史诗 - 3
        LEGENDARY   // 传说 - 4
    }

    // NFT元数据结构
    struct NFTMetadata {
        Rarity rarity;
        uint256 difficulty;
        string captureId;
        uint256 timestamp;
        string attributes;
    }

    // 捕捉事件结构
    struct CaptureEvent {
        address player;
        string captureId;
        bool success;
        Rarity rarity;
        uint256 difficulty;
        uint256 timestamp;
        uint256 tokenId; // 如果成功则为NFT的tokenId，否则为0
    }

    // 存储NFT元数据
    mapping(uint256 => NFTMetadata) public nftMetadata;
    
    // 存储捕捉事件
    mapping(string => CaptureEvent) public captureEvents;
    
    // 玩家统计
    mapping(address => uint256) public playerTotalCaptures;
    mapping(address => uint256) public playerSuccessfulCaptures;
    
    // 稀有度统计
    mapping(Rarity => uint256) public rarityCount;
    
    // 游戏配置
    mapping(uint256 => uint256) public difficultySuccessRate; // 难度 => 成功率(百分比)
    
    // 授权的铸造者（后端服务）
    mapping(address => bool) public authorizedMinters;

    // 事件定义
    event CaptureAttempted(
        address indexed player,
        string indexed captureId,
        bool success,
        Rarity rarity,
        uint256 difficulty,
        uint256 tokenId
    );
    
    event NFTMinted(
        address indexed to,
        uint256 indexed tokenId,
        Rarity rarity,
        string captureId
    );
    
    event MinterAuthorized(address indexed minter);
    event MinterRevoked(address indexed minter);

    constructor(
        string memory name,
        string memory symbol
    ) ERC721(name, symbol) Ownable(msg.sender) {
        // 初始化难度成功率配置
        difficultySuccessRate[1] = 80; // 简单 - 80%
        difficultySuccessRate[2] = 60; // 中等 - 60%
        difficultySuccessRate[3] = 40; // 困难 - 40%
        difficultySuccessRate[4] = 20; // 极难 - 20%
        difficultySuccessRate[5] = 10; // 传说 - 10%
    }

    /**
     * @dev 授权铸造者
     */
    function authorizeMinter(address minter) external onlyOwner {
        authorizedMinters[minter] = true;
        emit MinterAuthorized(minter);
    }

    /**
     * @dev 撤销铸造者授权
     */
    function revokeMinter(address minter) external onlyOwner {
        authorizedMinters[minter] = false;
        emit MinterRevoked(minter);
    }

    /**
     * @dev 修饰符：只有授权的铸造者可以调用
     */
    modifier onlyAuthorizedMinter() {
        require(authorizedMinters[msg.sender] || msg.sender == owner(), "Not authorized minter");
        _;
    }

    /**
     * @dev 尝试捕捉（由后端调用）
     */
    function attemptCapture(
        address player,
        string memory captureId,
        uint256 difficulty,
        Rarity rarity,
        bool success,
        string memory tokenURI
    ) external onlyAuthorizedMinter nonReentrant returns (uint256) {
        require(player != address(0), "Invalid player address");
        require(bytes(captureId).length > 0, "Invalid capture ID");
        require(difficulty >= 1 && difficulty <= 5, "Invalid difficulty");
        require(captureEvents[captureId].timestamp == 0, "Capture ID already exists");

        // 更新玩家统计
        playerTotalCaptures[player]++;
        
        uint256 tokenId = 0;
        
        if (success) {
            // 铸造NFT
            tokenId = _mintNFT(player, rarity, difficulty, captureId, tokenURI);
            playerSuccessfulCaptures[player]++;
            rarityCount[rarity]++;
        }

        // 记录捕捉事件
        captureEvents[captureId] = CaptureEvent({
            player: player,
            captureId: captureId,
            success: success,
            rarity: rarity,
            difficulty: difficulty,
            timestamp: block.timestamp,
            tokenId: tokenId
        });

        emit CaptureAttempted(player, captureId, success, rarity, difficulty, tokenId);
        
        return tokenId;
    }

    /**
     * @dev 内部铸造NFT函数
     */
    function _mintNFT(
        address to,
        Rarity rarity,
        uint256 difficulty,
        string memory captureId,
        string memory tokenURI
    ) internal returns (uint256) {
        uint256 tokenId = _nextTokenId;
        _nextTokenId++;

        _safeMint(to, tokenId);
        _setTokenURI(tokenId, tokenURI);

        // 存储NFT元数据
        nftMetadata[tokenId] = NFTMetadata({
            rarity: rarity,
            difficulty: difficulty,
            captureId: captureId,
            timestamp: block.timestamp,
            attributes: ""
        });

        emit NFTMinted(to, tokenId, rarity, captureId);
        
        return tokenId;
    }

    /**
     * @dev 获取玩家拥有的所有NFT
     */
    function getPlayerNFTs(address player) external view returns (uint256[] memory) {
        uint256 balance = balanceOf(player);
        uint256[] memory tokenIds = new uint256[](balance);
        
        for (uint256 i = 0; i < balance; i++) {
            tokenIds[i] = tokenOfOwnerByIndex(player, i);
        }
        
        return tokenIds;
    }

    /**
     * @dev 获取NFT的详细信息
     */
    function getNFTDetails(uint256 tokenId) external view returns (
        address owner,
        Rarity rarity,
        uint256 difficulty,
        string memory captureId,
        uint256 timestamp,
        string memory tokenURI
    ) {
        require(_ownerOf(tokenId) != address(0), "Token does not exist");
        
        NFTMetadata memory metadata = nftMetadata[tokenId];
        
        return (
            ownerOf(tokenId),
            metadata.rarity,
            metadata.difficulty,
            metadata.captureId,
            metadata.timestamp,
            super.tokenURI(tokenId)
        );
    }

    /**
     * @dev 获取玩家统计信息
     */
    function getPlayerStats(address player) external view returns (
        uint256 totalCaptures,
        uint256 successfulCaptures,
        uint256 nftCount,
        uint256 successRate
    ) {
        totalCaptures = playerTotalCaptures[player];
        successfulCaptures = playerSuccessfulCaptures[player];
        nftCount = balanceOf(player);
        
        if (totalCaptures > 0) {
            successRate = (successfulCaptures * 100) / totalCaptures;
        } else {
            successRate = 0;
        }
    }

    /**
     * @dev 获取稀有度统计
     */
    function getRarityStats() external view returns (uint256[5] memory) {
        return [
            rarityCount[Rarity.COMMON],
            rarityCount[Rarity.UNCOMMON],
            rarityCount[Rarity.RARE],
            rarityCount[Rarity.EPIC],
            rarityCount[Rarity.LEGENDARY]
        ];
    }

    /**
     * @dev 设置难度成功率
     */
    function setDifficultySuccessRate(uint256 difficulty, uint256 successRate) external onlyOwner {
        require(difficulty >= 1 && difficulty <= 5, "Invalid difficulty");
        require(successRate <= 100, "Success rate cannot exceed 100%");
        difficultySuccessRate[difficulty] = successRate;
    }

    /**
     * @dev 获取当前总供应量
     */
    function getCurrentTokenId() external view returns (uint256) {
        return _nextTokenId - 1;
    }

    // 重写必要的函数以解决继承冲突
    function _update(
        address to,
        uint256 tokenId,
        address auth
    ) internal override(ERC721, ERC721Enumerable) returns (address) {
        return super._update(to, tokenId, auth);
    }

    function _increaseBalance(address account, uint128 value) internal override(ERC721, ERC721Enumerable) {
        super._increaseBalance(account, value);
    }



    function tokenURI(uint256 tokenId) public view override(ERC721, ERC721URIStorage) returns (string memory) {
        return super.tokenURI(tokenId);
    }

    function supportsInterface(bytes4 interfaceId) public view override(ERC721, ERC721Enumerable, ERC721URIStorage) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}