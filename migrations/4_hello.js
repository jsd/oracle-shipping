const HelloWorld = artifacts.require('HelloWorld')


module.exports = async (deployer, network, [defaultAccount]) => {
    // contract automatically retrieve the correct address for you
    deployer.deploy(HelloWorld)
  
}
