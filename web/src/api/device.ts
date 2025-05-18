import request from '@/utils/request'

export interface Node {
    host: string
    port: number
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