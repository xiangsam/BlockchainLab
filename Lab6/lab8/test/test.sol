pragma solidity >=0.4.22 <0.9.0 ;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/Enrollment.sol";

contract TestEnrollment{
    function testAdministratorIdentification() {
        Enrollment enroll = Enrollment(DeployedAddresses.Enrollment());
        Assert.equal(address(enroll.administrator), DeployedAddresses.Enrollment(), 'administior error');
    }
    function testSignUp() public {
        Enrollment en = new Enrollment();
        string memory name = en.signUp("alice");
        string memory expected_name = "alice";
        Assert.equal(name,expected_name,"signup failed");
    }
    function testNewConf() public {
        Enrollment enroll = new Enrollment();
        string memory expected = "conf1";
        Assert.equal(enroll.newConference("conf1","beijing",30),expected,"new conference failed");
    }
    function  testEnroll() public {
        Enrollment enroll = new Enrollment();
        enroll.newConference("conf1", "beijing", 30);
        string memory expected = "conf1";
        Assert.equal(enroll.enroll("conf1"),expected,"enroll failed");
    }

}