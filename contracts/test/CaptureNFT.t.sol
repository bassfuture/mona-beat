// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/CaptureNFT.sol";

contract CaptureNFTTest is Test {
    CaptureNFT public captureNFT;
    address public player1 = address(0x1);
    address public player2 = address(0x2);

    function setUp() public {
        captureNFT = new CaptureNFT("CaptureNFT", "CNFT");
        // 授权测试合约为铸造者
        captureNFT.authorizeMinter(address(this));
    }

    function testAttemptCapture() public {
        uint256 tokenId = captureNFT.attemptCapture(
            player1,
            "capture1",
            1,
            CaptureNFT.Rarity.COMMON,
            true,
            "https://example.com/token1.json"
        );
        
        assertEq(captureNFT.ownerOf(tokenId), player1);
        assertEq(captureNFT.balanceOf(player1), 1);
    }

    function testMultipleCaptures() public {
        uint256 tokenId1 = captureNFT.attemptCapture(
            player1,
            "capture1",
            1,
            CaptureNFT.Rarity.COMMON,
            true,
            "https://example.com/token1.json"
        );
        
        uint256 tokenId2 = captureNFT.attemptCapture(
            player2,
            "capture2",
            2,
            CaptureNFT.Rarity.RARE,
            true,
            "https://example.com/token2.json"
        );
        
        assertEq(captureNFT.ownerOf(tokenId1), player1);
        assertEq(captureNFT.ownerOf(tokenId2), player2);
        assertEq(captureNFT.balanceOf(player1), 1);
        assertEq(captureNFT.balanceOf(player2), 1);
    }

    function testGetNFTDetails() public {
        uint256 tokenId = captureNFT.attemptCapture(
            player1,
            "capture1",
            1,
            CaptureNFT.Rarity.COMMON,
            true,
            "https://example.com/token1.json"
        );
        
        (
            address owner,
            CaptureNFT.Rarity rarity,
            uint256 difficulty,
            string memory captureId,
            uint256 timestamp,
            string memory tokenURI
        ) = captureNFT.getNFTDetails(tokenId);
        
        assertEq(owner, player1);
        assertTrue(uint8(rarity) >= 0 && uint8(rarity) <= 4);
        assertEq(difficulty, 1);
        assertEq(captureId, "capture1");
        assertTrue(timestamp > 0);
        assertTrue(bytes(tokenURI).length > 0);
    }

    function testGetPlayerStats() public {
        captureNFT.attemptCapture(
            player1,
            "capture1",
            1,
            CaptureNFT.Rarity.COMMON,
            true,
            "https://example.com/token1.json"
        );
        
        captureNFT.attemptCapture(
            player1,
            "capture2",
            2,
            CaptureNFT.Rarity.RARE,
            true,
            "https://example.com/token2.json"
        );
        
        (uint256 totalCaptures, uint256 successfulCaptures, uint256 nftCount, uint256 successRate) = captureNFT.getPlayerStats(player1);
        
        assertEq(totalCaptures, 2);
        assertEq(successfulCaptures, 2);
        assertEq(nftCount, 2);
        assertEq(successRate, 100);
    }

    function testTokenURI() public {
        uint256 tokenId = captureNFT.attemptCapture(
            player1,
            "capture1",
            1,
            CaptureNFT.Rarity.COMMON,
            true,
            "https://example.com/token1.json"
        );
        
        string memory tokenURI = captureNFT.tokenURI(tokenId);
        assertTrue(bytes(tokenURI).length > 0);
    }

    function testSupportsInterface() public {
        assertTrue(captureNFT.supportsInterface(0x80ac58cd)); // ERC721
        assertTrue(captureNFT.supportsInterface(0x5b5e139f)); // ERC721Metadata
    }
}