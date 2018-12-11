import React, {Component} from 'react';
import {Modal, Form, Input, Select,InputNumber,Switch} from 'antd';
import {inject, observer} from 'mobx-react';

const FormItem = Form.Item;
const Option = Select.Option;

@inject('stores')
@observer
class PathModal extends Component {
    constructor(props) {
        super(props);
        this.groupStore = this.props.stores.groupStore;
    }

    handleOk = () => {
        this.props.form.validateFields((err, values) => {
            if (!err) {
                let pathData ={
                    ...values,
                    serverPort: parseInt(values.serverPort),
                    id: this.groupStore.currentPathData.id,
                    groupId: this.groupStore.currentPathData.groupId,
                }
                this.groupStore.hiddenPathModal();
                if (this.groupStore.pathModalType === 'create') {
                    this.groupStore.createPath(pathData);
                    return;
                }
                this.groupStore.updatePath(pathData)
            }
        })
    };

    render() {
        const {getFieldDecorator} = this.props.form;
        const {groupStore} = this.props.stores;
        const {currentPathData} = groupStore;
        const {pathModalType} = groupStore;
        const formItemLayout = {
            labelCol: {
                span: 6
            },
            wrapperCol: {
                span: 16
            }
        };
        return (
            <Modal visible={groupStore.pathModalVisible}
                   title={groupStore.pathModalType === 'create' ? '新建路由' : '编辑'}
                   onOk={() => this.handleOk()}
                   onCancel={() => this.props.stores.groupStore.hiddenPathModal()}>
                <Form>
                    <FormItem {...formItemLayout} label="name">
                        {getFieldDecorator('name', {
                            initialValue: pathModalType === 'create' ? '' : currentPathData.name,
                        })(
                            <Input placeholder="非必填" />
                        )}
                    </FormItem>
                    <FormItem {...formItemLayout} label="path">
                        {getFieldDecorator('path', {
                            initialValue: pathModalType === 'create' ? '' : currentPathData.path,
                            rules: [{
                                required: true, message: '路径不能为空'
                            }, {
                                pattern: '\^/S*', message: '路径格式不正确'
                            }]
                        })(
                            <Input addonBefore="http://gateway.letsgo.tech/api"/>
                        )}
                    </FormItem>
                    <FormItem {...formItemLayout} label="serverMethod">
                        {getFieldDecorator('serverMethod', {
                            initialValue: pathModalType === 'create' ? 'GET' : currentPathData.serverMethod,
                            rules: [{
                                required: true, message: '必填'
                            }]
                        })(
                            <Select>
                                <Option value="GET">GET</Option>
                                <Option value="POST">POST</Option>
                                <Option value="PUT">PUT</Option>
                                <Option value="PATCH">PATCH</Option>
                                <Option value="DELETE">DELETE</Option>
                                <Option value="OPTION">OPTION</Option>
                                <Option value="HEADER">HEADER</Option>
                            </Select>
                        )}
                    </FormItem>
                    <FormItem {...formItemLayout} label="serverScheme">
                        {getFieldDecorator('serverScheme', {
                            initialValue: pathModalType === 'create' ? 'http' : currentPathData.serverScheme,
                            rules: [{
                                required: true, message: '必填'
                            }]
                        })(
                            <Select>
                                <Option value="http">http</Option>
                                <Option value="https">https</Option>
                            </Select>
                        )}
                    </FormItem>
                    <FormItem {...formItemLayout} label="serviceName">
                        {getFieldDecorator('serviceName', {
                            initialValue: pathModalType === 'create' ? '' : currentPathData.serviceName
                        })(
                            <Input/>
                        )}
                    </FormItem>
                    <FormItem {...formItemLayout} label="serverHost">
                        {getFieldDecorator('serverHost', {
                            initialValue: pathModalType === 'create' ? '' : currentPathData.serverHost,
                        })(
                            <Input/>
                        )}
                    </FormItem>
                    <FormItem {...formItemLayout} label="serverPort">
                        {getFieldDecorator('serverPort', {
                            initialValue: pathModalType === 'create' ? '' : currentPathData.serverPort,
                        })(
                            <InputNumber />
                        )}
                    </FormItem>
                    <FormItem {...formItemLayout} label="serverPath">
                        {getFieldDecorator('serverPath', {
                            initialValue: pathModalType === 'create' ? '' : currentPathData.serverPath,
                            rules: [{
                                required: true, message: '必填'
                            }, {
                                pattern: '\^/S*',
                                message: '格式不正确'
                            }]
                        })(
                            <Input/>
                        )}
                    </FormItem>
                    <FormItem {...formItemLayout} label="认证">
                        {getFieldDecorator('auth', {
                            initialValue: pathModalType === 'create' ? true : currentPathData.status,
                            rules: [{
                                required: true, message: '必填'
                            }]
                        })(
                            <Switch checkedChildren="开启" unCheckedChildren="关闭" defaultChecked />
                        )}
                    </FormItem>
                    <FormItem {...formItemLayout} label="状态">
                        {getFieldDecorator('status', {
                            initialValue: pathModalType === 'create' ? true : currentPathData.status,
                            rules: [{
                                required: true, message: '必填'
                            }]
                        })(
                            <Switch checkedChildren="启用" unCheckedChildren="停用" defaultChecked />
                        )}
                    </FormItem>
                </Form>
            </Modal>
        )
    }
}

export default Form.create()(PathModal);
