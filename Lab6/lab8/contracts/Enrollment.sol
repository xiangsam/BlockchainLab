pragma solidity >=0.4.22 <0.9.0;
pragma experimental ABIEncoderV2;

contract Enrollment {
    address public administrator;
    struct Conference {
        string title;
        string detail;
        uint limitation;
        uint current;
        bool isFull;
    }
    struct Participant {
        string username;
        string[] myconferences;
    }
    // 参与者address 对应 参与者
    mapping(address => Participant) participants;
    // 委托人address 对应 参与者
    mapping(address => Participant[]) trustees;
    Conference[] public conferences;

    constructor() public {
        administrator = msg.sender;
    }

    modifier onlyOwner(){
        require(msg.sender == administrator);
        _;
    }

    event NewConference(string,string);
    event ConferenceExpire(string);
    event MyNewConference(string);

    function signUp(string _username) public returns(string){
        require(bytes(participants[msg.sender].username).length<=0,"User exists.");
        participants[msg.sender].username = _username;
        return participants[msg.sender].username;
    }

    function delegate(address trustee) public {
        require(trustee != msg.sender);
        trustees[trustee].push(participants[msg.sender]);
    }


    function enroll(string _title) public returns(string){
        bool finded = false;
        for(uint i = 0; i < conferences.length; ++i){
            if(keccak256(abi.encodePacked(_title)) == keccak256(abi.encodePacked(conferences[i].title))){
                if(!conferences[i].isFull){
                    participants[msg.sender].myconferences.push(_title);
                    conferences[i].current++;
                    if(conferences[i].current == conferences[i].limitation){
                        conferences[i].isFull = true;
                    }
                    emit MyNewConference(_title);
                    finded = true;
                    break;
                }
                emit ConferenceExpire(_title);
            }
        }
        if(!finded){
            revert('fail to enroll');
        }
        return _title;
    }

    function enrollFor(string _title, string _username) public returns(string){
        bool finded = false;
        uint userId = 0;
        for(uint i = 0; i < trustees[msg.sender].length; ++i){
            if(keccak256(abi.encodePacked(trustees[msg.sender][i].username))==keccak256(abi.encodePacked(_username))){
                finded = true;
                userId = i;
            }
        }
        if(!finded){
            revert('fail to enroll');
        }
        finded = false;
        for(i = 0; i < conferences.length; ++i){
            if(keccak256(abi.encodePacked(_title)) == keccak256(abi.encodePacked(conferences[i].title))){
                if(!conferences[i].isFull){
                    trustees[msg.sender][userId].myconferences.push(_title);
                    conferences[i].current++;
                    if(conferences[i].current == conferences[i].limitation){
                        conferences[i].isFull = true;
                    }
                    finded = true;
                    break;
                }
            }
        }
        if(!finded){
            revert('fail to enroll');
        }
        return _title;
    }

    function newConference(string _title, string _detail,uint _limitation) public onlyOwner returns(string){
        conferences.push(Conference(_title, _detail,_limitation, 0,false));
        emit NewConference(_title,_detail);
        return _title;
    }

    function queryConfList()  public view returns(string[], string[]) {
        string[] memory title = new string[](conferences.length);
        string[] memory detail = new string[](conferences.length);
        uint count = 0;
        for(uint i = 0; i < conferences.length; ++i){
            title[count] = conferences[i].title;
            detail[count] = conferences[i].detail;
            count++;
        }
        return (title, detail);
    }

    function queryMyConf() public view returns(string[]) {
        return participants[msg.sender].myconferences;
    }

    function destruct() private onlyOwner{
        selfdestruct(administrator);
    }

    
}