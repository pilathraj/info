// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./SimpleStorage.sol";

contract StorageFactory is SimpleStorage {
    SimpleStorage[] public simpleStorageArray;
    function createSimpleStorage() public {
        SimpleStorage simpleStorage = new SimpleStorage();
        simpleStorageArray.push(simpleStorage);
    }

    function sfStorage(uint256 index, uint256 number) public {
        simpleStorageArray[index].storeFaviouriteNumber(number);
    }

    function sfRetrive(uint256 index) public view returns(uint256){
        return simpleStorageArray[index].retrive();
    }
}
