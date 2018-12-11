import React, {Component} from 'react';
import {Button,} from 'antd';
import './index.less';
import GroupModal from "./GroupModal";
import PathModal from './PathModal';
import {inject, observer} from 'mobx-react';
import Container from './Container';

@inject('stores')
@observer
class Group extends Component {
    constructor(props) {
        super(props);
        this.groupStore = this.props.stores.groupStore;
    }

    handleClickGroup() {
        this.groupStore.showGroupModal();
        this.groupStore.setGroupModalType('create');
    }

    render() {

        return (
            <div className='group-page'>
                <div className="header">
                    <span className='title'>分组</span>
                    <Button onClick={() => this.handleClickGroup()}>新建分组</Button>
                </div>
                <div className="main-container">
                    <Container/>
                </div>
                <GroupModal/>
                <PathModal/>
            </div>
        )
    }
}

export default Group;
