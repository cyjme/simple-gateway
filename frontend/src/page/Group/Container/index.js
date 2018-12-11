import React, {Component} from 'react';
import {Collapse, Table, Button} from 'antd';
import {inject, observer} from 'mobx-react';

const Panel = Collapse.Panel;

function callback(key) {
    console.log(key)
}

@inject('stores')
@observer
class Container extends Component {
    constructor(props) {
        super(props);
        this.groupStore = this.props.stores.groupStore;
        this.groupStore.getGroups();
    }

    editGroup(e,group) {
        this.groupStore.showGroupModal();
        this.groupStore.setGroupModalType('edit');
        this.groupStore.setCurrentGroupModalData(group);
        e.stopPropagation();
    }

    deleteGroup(e,groupId) {
        this.groupStore.deleteGroupById(groupId);
        e.stopPropagation();
    }

    editPath(record) {
        this.groupStore.showPathModal();
        this.groupStore.setPathModalType('edit');
        console.log("record",record)
        this.groupStore.setCurrentPathModalData(record)
    }

    createPath(e,groupId) {
        this.groupStore.showPathModal();
        this.groupStore.setPathModalType('create');
        this.groupStore.setCurrentPathModalGroupId(groupId);
        e.stopPropagation();
    }

    deletePath(pathId) {
        this.groupStore.deletePath(pathId);
    }

    render() {
        const groupStore = this.props.stores.groupStore;
        const columns = [
            {
                title: ' 名称',
                dataIndex: 'name',
                key: 'name'
            },
            {
                title: 'path',
                dataIndex: 'path',
                key: 'path'
            },
            {
                title: 'serverMethod',
                dataIndex: 'serverMethod',
                key: 'serverMethod'
            }, 
            {
                title: 'serverScheme',
                dataIndex: 'serverScheme',
                key: 'serverScheme'
            }, 
            {
                title: 'serviceName',
                dataIndex: 'serviceName',
                key: 'serviceName'
            }, 
            {
                title: 'serverHost',
                dataIndex: 'serverHost',
                key: 'serverHost'
            }, 
            {
                title: 'serverPort',
                dataIndex: 'serverPort',
                key: 'serverPort'
            }, 
            {
                title: 'serverPath',
                dataIndex: 'serverPath',
                key: 'serverPath'
            }, 
            {
                title: '调用次数',
                dataIndex: 'time',
                key: 'time'
            }, 
            {
                title: '认证',
                dataIndex: 'auth',
                key: 'auth'
            },
            {
                title: ' 操作',
                dataIndex: 'operate',
                key: 'operate',
                render: (text, record) => (
                    <span className='operate-record'>
                        <a onClick={() => this.editPath(record)}>编辑</a>
                        <a onClick={() => this.deletePath(record.id)}>删除</a>
                    </span>
                )
            },
        ];

        return (
            <Collapse defaultActiveKey={['0']} onchange={callback}>
                {groupStore.groupList.map((group, index) => {
                    return (
                        <Panel className='panel-box'
                               key={index}
                               header={
                                   <div className='panel-header'>
                                       <span>{group.name}</span>
                                       <div className="btn-group">
                                           <Button onClick={(e) => this.editGroup(e,group)}>编辑</Button>
                                           <Button onClick={(e) => this.createPath(e,group.id)}>新建</Button>
                                           <Button onClick={(e) => this.deleteGroup(e,group.id)}>删除</Button>
                                       </div>
                                   </div>}>
                            <Table columns={columns} dataSource={group.path} pagination={false}/>
                        </Panel>
                    );
                })}
            </Collapse>
        )
    }
}

export default Container;
