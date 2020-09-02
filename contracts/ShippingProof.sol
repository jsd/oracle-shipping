pragma solidity ^0.6.0;

import "@chainlink/contracts/src/v0.6/ChainlinkClient.sol";

contract ShippingProof is ChainlinkClient {

    bytes32 public data;

    address private oracle;
    bytes32 private jobId;
    uint256 private fee;

    /**
     * Network: Ropsten
     * Oracle: Chainlink - 0xc99B3D447826532722E41bc36e644ba3479E4365
     * Job ID: Chainlink - 3cff0a3524694ff8834bda9cf9c779a1
     * Fee: 0.1 LINK
     */
    constructor(address _link) public {
        if(_link == address(0)) {
          setPublicChainlinkToken();
        } else {
          setChainlinkToken(_link);
        }
        oracle = 0x0bE828811DB2bf0Eb0d28F0a697B5514A7cc70B7;
        jobId = "f08e3ca2aba94c6aafcd610843d3d4d9";
        fee =  1 * LINK;  // 0.1 LINK
    }

    function requestShippingStatus(address _oracle, string memory _jobId) public returns (bytes32 requestId) {
        Chainlink.Request memory req = buildChainlinkRequest(stringToBytes32(_jobId), address(this), this.fulfill.selector);
        req.add("id", "300403513353");
        req.add("copyPath", "status");
        return sendChainlinkRequestTo(_oracle, req, fee);
    }

    function stringToBytes32(string memory source) private pure returns (bytes32 result) {
        bytes memory tempEmptyStringTest = bytes(source);
            if (tempEmptyStringTest.length == 0) {
              return 0x0;
            }

        assembly { // solhint-disable-line no-inline-assembly
          result := mload(add(source, 32))
        }
    }

    /**
     * Receive the response in the form of bytes32
     */
     function fulfill(bytes32 _requestId, bytes32 _data) public recordChainlinkFulfillment(_requestId)
    {
        data = _data;
    }
}
