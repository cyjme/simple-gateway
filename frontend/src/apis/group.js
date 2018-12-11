import request from '../utils/request';

export async function getGroups() {
    return await request.get('/routeGroup');
}

export async function createGroup(groupData) {
    return await request.post('/routeGroup', groupData);
}

export async function updateGroup(groupData) {
    return await request.put('/routeGroup/' + groupData.id, groupData);
}

export async function deleteGroupById(groupId) {
    return await request.delete('/routeGroup/' + groupId);
}

export async function getPathByGroupId(groupId) {
    return await request.get('/routeGroup/' + groupId + '/routes');

}

export async function createPath(pathData) {
    return await request.post('/routes', pathData);
}

export async function updatePath(pathData) {
    return await request.put('/routes/' + pathData.id, pathData);
}

export async function deletePath(pathId) {
    return await request.delete('/routes/' + pathId);
}
