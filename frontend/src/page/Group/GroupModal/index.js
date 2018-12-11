import React, {Component} from 'react';
import {Modal, Form, Input} from 'antd';
import {inject, observer} from 'mobx-react';

const FormItem = Form.Item;

@inject('stores')
@observer
class GroupModal extends Component {
    constructor(props) {
        super(props);
        this.groupStore = this.props.stores.groupStore;
    }

    handleOk = () => {
        this.props.form.validateFields((err, values) => {
            if (!err) {
                const groupData = {
                    name: values.name,
                    direction: values.direction,
                    type: 'http',
                    status: '0',
                    id: this.groupStore.currentGroupData.id
                };
                this.groupStore.hiddenGroupModal();
                if (this.groupStore.groupModalType === 'create') {
                    this.groupStore.creatrGroup(groupData);
                    return;
                }
                this.groupStore.updateGroup(groupData);
            }
        })
    };

    render() {
        const {getFieldDecorator} = this.props.form;
        const {groupStore} = this.props.stores;
        const {currentGroupData} = groupStore;
        const {groupModalType} = groupStore;
        const formItemLayout = {
            labelCol: {
                span: 6
            },
            wrapperCol: {
                span: 16
            }
        };
        return (
            <Modal visible={groupStore.groupModalVisible}
                   title={groupStore.groupModalType === 'create' ? '新建分组' : '编辑' + currentGroupData.name}
                   onOk={() => this.handleOk()}
                   onCancel={() => this.groupStore.hiddenGroupModal()}>
                <Form>
                    <FormItem {...formItemLayout} label="分组名称">
                        {getFieldDecorator('name', {
                            initialValue: groupModalType === 'create' ? '' : currentGroupData.name,
                            rules: [{
                                required: true, message: '昵称不能为空'
                            }]
                        })(
                            <Input/>
                        )}
                    </FormItem>
                    <FormItem {...formItemLayout} label="分组描述">
                        {getFieldDecorator('direction', {
                            initialValue: groupModalType === 'create' ? '' : currentGroupData.direction,
                            rules: [{
                                required: true, message: '分组不能为空'
                            }]
                        })(
                            <Input/>
                        )}
                    </FormItem>
                </Form>
            </Modal>
        )
    }
}

export default Form.create()(GroupModal);
