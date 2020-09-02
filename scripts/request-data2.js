const ShippingProof = artifacts.require('ShippingProof')

/*
  This script allows for a Chainlink request to be created from
  the requesting contract. Defaults to the Chainlink oracle address
  on this page: https://docs.chain.link/docs/testnet-oracles
*/

const oracleAddress ='0x6d95ad369f8fecf4bb3010bdc99490f252c22701'
const jobId ='7b8f231558d4412a8bf0157af69a52da'
const payment = '1000000000000000000'

module.exports = async callback => {
  const mc = await ShippingProof.deployed()
  console.log('Creating request on contract:', mc.address)
  const tx = await mc.requestShippingStatus(oracleAddress,jobId)
  callback(tx.tx)
}
