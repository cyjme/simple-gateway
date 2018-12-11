import {observable} from 'mobx';
import {
    createGroup,
    updateGroup,
    createPath,
    updatePath,
    deletePath,
    getGroups,
    deleteGroupById,
    getPathByGroupId
} from "../apis/group";
import {message} from 'antd';

class GroupStore {
    @observable groupModalVisible = false;
    @observable pathModalVisible = false;
    @observable groupModalType = "";
    @observable pathModalType = "";
    @observable currentGroupData = {};
    @observable currentPathData = {};
    @observable groupList = [];

    showGroupModal() {
        this.groupModalVisible = true;
    }

    hiddenGroupModal() {
        this.groupModalVisible = false;
    }

    setGroupModalType(type) {
        this.groupModalType = type
    }

    setCurrentGroupModalData(group) {
        this.currentGroupData = {
            id: group.id,
            name: group.name,
            direction: group.direction,
        }
    }

    showPathModal() {
        this.pathModalVisible = true;
    }

    hiddenPathModal() {
        this.pathModalVisible = false;
    }

    setCurrentPathModalGroupId(groupId) {
        this.currentPathData.groupId = groupId;
    }

    setPathModalType(type) {
        this.pathModalType = type
    }

    setCurrentPathModalData(path) {
        this.currentPathData = path
    }

    //获取 group 数组
    //map group 数组
    //map 过程中为数组中的每一项增加一个属性 path
    //path 为一个数组
    //如果没有复数，也可以用 listGroup 表示获取数组，用 fetchGroup 表示获取单个
    async getGroups() {
        //result 后面不需要修改，所以用  const result
        const result = await getGroups();
        if (result.status === 200) {
            this.groupList = result.data;
            this.groupList.map(async group => {
                group.path = await this.getPathByGroupId(group.id);
            });
            return;
        }
        message.error('获取失败');
    }

    async creatrGroup(groupData) {
        const result = await createGroup(groupData);
        if (result.status === 200) {
            this.getGroups();
            return;
        }
        message.error('新建分组失败');
    }

    async updateGroup(groupData) {
        const result = await updateGroup(groupData);
        if (result.status === 200) {
            this.getGroups();
            return;
        }
        message.error('分组修改失败');
    }

    async deleteGroupById(groupId) {
        const result = await deleteGroupById(groupId);
        if (result.status === 200) {
            this.getGroups();
            return;
        }
        message.error('删除失败');
    }

    async getPathByGroupId(groupId) {
        let result = await getPathByGroupId(groupId);
        if (result.status === 200) {
            result.data.map(item => {
                item.key = item.id
            });
            return result.data;
        }
        message.error('path 加载失败');
    }

    async createPath(pathData) {
        const result = await createPath(pathData);
        if (result.status === 200) {
            this.getGroups();
            return;
        }
        message.error('path添加失败');
    };

    async updatePath(pathData) {
        const result = await updatePath(pathData);
        if (result.status === 200) {
            this.getGroups();
            return;
        }
        message.error(' path修改失败');
    }

    async deletePath(pathId) {
        const result = await deletePath(pathId);
        if (result.status === 200) {
            this.getGroups();
            return;
        }
        message.error('删除失败');
    }
}

export default GroupStore;
