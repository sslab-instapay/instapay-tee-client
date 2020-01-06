pragma solidity ^0.4.23;

contract InstaPay {
    enum ChannelStatus {CLOSED, OPEN}
    enum Stage {PRE_UPDATE, POST_UPDATE}

    struct Channel {
        address owner;
        address receiver;
        uint256 deposit;
        ChannelStatus status;
    }

    struct Ejection {
        bool registered;
        Stage stage;
    }

    uint256 acc_id = 0;
    uint256 public readme = 4321;
    mapping (uint256 => Channel) public channels;
    mapping (uint256 => Ejection) public ejections;

    event EventCreateChannel(uint256 id, address owner, address receiver, uint256 deposit);
    event EventCloseChannel(uint256 id, uint256 ownerbal, uint256 receiverbal);
    event EventEject(uint256 pn, Stage registeredstage);

    modifier onlyOwnerOrReceiver(uint256 id, address sender) {
        require(channels[id].owner == sender || channels[id].receiver == sender);
        _;
    }

    function create_channel(address receiver) public payable {
        require(channels[acc_id + 1].status == ChannelStatus.CLOSED);
        acc_id++;

        channels[acc_id].owner = msg.sender;
        channels[acc_id].receiver = receiver;
        channels[acc_id].deposit = msg.value;
        channels[acc_id].status = ChannelStatus.OPEN;

        emit EventCreateChannel(acc_id, msg.sender, receiver, msg.value);
    }

    function close_channel(uint256 id, uint256 owner_bal, uint256 receiver_bal) public onlyOwnerOrReceiver(id, msg.sender) {
        require(channels[id].status == ChannelStatus.OPEN);

        channels[id].owner.transfer(owner_bal * 1 ether);
        channels[id].receiver.transfer(receiver_bal * 1 ether);
        channels[id].status = ChannelStatus.CLOSED;

        emit EventCloseChannel(id, owner_bal, receiver_bal);
    }

    function settle_channels(Stage registered_stage, Stage stage, uint256[] ids, uint256[] bals, uint256 v) internal {
        if(registered_stage == Stage.PRE_UPDATE) {
            if(stage == Stage.PRE_UPDATE) {
                for(uint256 i = 0; i < ids.length; i++) {
                    if(channels[ids[i]].status == ChannelStatus.CLOSED) continue;
                    close_channel(ids[i], bals[i+i], bals[i+i+1]);
                }
            }
            else {
                for(uint256 j = 0; j < ids.length; j++) {
                    if(channels[ids[j]].status == ChannelStatus.CLOSED) continue;
                    close_channel(ids[j], bals[j+j] + v, bals[j+j+1] - v);
                }
            }
        }
        else {  // stage == Stage.POST_UPDATE
            if(stage == Stage.PRE_UPDATE) {
                for(uint256 k = 0; k < ids.length; k++) {
                    if(channels[ids[k]].status == ChannelStatus.CLOSED) continue;
                    close_channel(ids[k], bals[k+k] - v, bals[k+k+1] + v);
                }
            }
            else {
                for(uint256 l = 0; l < ids.length; l++) {
                    if(channels[ids[l]].status == ChannelStatus.CLOSED) continue;
                    close_channel(ids[l], bals[l+l], bals[l+l+1]);
                }
            }
        }
    }

    function eject(uint256 pn, Stage stage, uint256[] ids, uint256[] bals, uint256 v) public {
        if(ejections[pn].registered == true) {
            settle_channels(ejections[pn].stage, stage, ids, bals, v);
        }
        else {
            ejections[pn].registered = true;
            ejections[pn].stage = stage;

            for(uint256 i = 0; i < ids.length; i++) {
                close_channel(ids[i], bals[i+i], bals[i+i+1]);
            }

            emit EventEject(pn, stage);
        }
    }
}
