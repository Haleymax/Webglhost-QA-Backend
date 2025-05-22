import request from '@/utils/request'

export interface Node {
    host: string
    port: number
    user: string
    password: string
}

export interface updata {
    host: string
    name: string
    user: string
    password: string
}

export const addNode = (node: Node) => {
    return request({
        url: '/api/nodes/add',
        method: 'post',
        data: node
    })
}

export const getNodes = () => {
    return request({
        url: '/api/nodes/get',
        method: 'get'
    })
}

export const updataNode = (node: updata) => {
    return request({
        url: '/api/nodes/updata',
        method: 'put',
        data: node
    })
}

export const deleteNode = (host: string) => {
    return request({
        url: '/api/nodes/remove',
        method: 'delete',
        data: host
    })
}

export const uploadFile = (data: FormData) => {
    return request({
        url: '/api/nodes/upload',
        method: 'post',
        data: data,
    })
}