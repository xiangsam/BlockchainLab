/*
 * @Author: Samrito
 * @Date: 2021-11-20 20:49:35
 * @LastEditors: Samrito
 * @LastEditTime: 2021-11-20 20:54:50
 */
const ConvertLib = artifacts.require("ConvertLib");
const Enrollment = artifacts.require("Enrollment");

module.exports = function(deployer) {
    deployer.deploy(ConvertLib);
    deployer.link(ConvertLib, Enrollment);
    deployer.deploy(Enrollment);
}