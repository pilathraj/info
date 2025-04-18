// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleStorage{
    // this will initialized to zero
    uint256 public faviouriteNumber;

    struct People {
        string name;
        uint256 faviouriteNumber;
    }
    People[] public people;
    mapping(string=>uint256) public nameToFaviouriateNumber;

    People public person = People({name:"hello", faviouriteNumber:0});

    function storeFaviouriteNumber(uint256 _faviouriteNumber) public  {
        faviouriteNumber = _faviouriteNumber;
    }
    // read only function no state change - know as view returns
    function retrive() public view returns(uint256) {
        return faviouriteNumber;
    }
    // do some mathematical without change the state and return anything - pure function
    function retrive2(uint256 _faviouriteNumber) public pure {
        _faviouriteNumber+_faviouriteNumber;
    }

    function addPersion(string memory _name, uint256 _faviouriteNumber) public {
        people.push(People({name:_name, faviouriteNumber:_faviouriteNumber}));
        nameToFaviouriateNumber[_name]=_faviouriteNumber;
    }

}
