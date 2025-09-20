// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/CaptureNFT.sol";

contract DeployScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        // Deploy the CaptureNFT contract
        CaptureNFT captureNFT = new CaptureNFT("CaptureNFT", "CNFT");
        
        console.log("CaptureNFT deployed to:", address(captureNFT));

        vm.stopBroadcast();
    }
}