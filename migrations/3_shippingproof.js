const ShippingProof = artifacts.require('ShippingProof')
const { LinkToken } = require('@chainlink/contracts/truffle/v0.4/LinkToken')
const { Oracle } = require('@chainlink/contracts/truffle/v0.6/Oracle')

module.exports = async (deployer, network, [defaultAccount]) => {

    // contract automatically retrieve the correct address for you
    deployer.deploy(ShippingProof, "0xa36085F69e2889c224210F603D836748e7dC0088")

}
