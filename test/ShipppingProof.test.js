
const ShippingProof = artifacts.require('ShippingProof');
const { oracle } = require('@chainlink/test-helpers');
const { LinkToken } = require('@chainlink/contracts/truffle/v0.4/LinkToken')

contract("ShippingProof", accounts => {
  let cc, link, request;
  const defaultAccount = accounts[0];
  const oracleAddress ='0x6d95ad369f8fecf4bb3010bdc99490f252c22701'
  const jobId ='7b8f231558d4412a8bf0157af69a52da'
  const payment = '1000000000000000000'

  beforeEach(async () => {
    cc = await ShippingProof.new("0xa36085F69e2889c224210F603D836748e7dC0088", { from: defaultAccount });

    //cc = await ShippingProof.deployed()

    link = await LinkToken.new({ from: defaultAccount })
    await link.transfer(cc.address, web3.utils.toWei('1', 'ether'), {
      from: defaultAccount,
    })

    //const tx = await cc.requestShippingStatus(oracleAddress,jobId);
    //request = oracle.decodeRunRequest(tx.receipt.rawLogs[3])
  })

  it("Should return a tx id",  async callback => {
    //const tx = await cc.requestShippingStatus(oracleAddress,jobId);
  console.log('Creating request on contract:', cc.address)
  //const tx = await cc.requestShippingStatus(oracleAddress,jobId);
    //request = cc.decodeRunRequest(tx.receipt.rawLogs[3])
  //  const result  =  await cc.data.call();
  //  console.log("data ", result )
   callback(tx.tx);

    assert(1 == 1);

  } )



})
