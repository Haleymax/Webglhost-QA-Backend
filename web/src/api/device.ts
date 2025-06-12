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

export interface GetPhoneInfo {
  "host": string,
  "serial": string,
}

export const addNode = (node: Node) => {
    return request({
        url: '/api/v1/nodes/add',
        method: 'post',
        data: node
    })
}

export const getNodes = () => {
    return request({
        url: '/api/v1/nodes/get',
        method: 'get'
    })
}

export const updataNode = (node: updata) => {
    return request({
        url: '/api/v1/nodes/updata',
        method: 'put',
        data: node
    })
}

export const deleteNode = (host: string) => {
    return request({
        url: '/api/v1/nodes/remove',
        method: 'delete',
        data: host
    })
}

export const uploadFile = (data: FormData) => {
    return request({
        url: '/api/v1/nodes/upload',
        method: 'post',
        data: data,
    })
}

export const getPhoneList = (host: string) => {
    return request({
        url: '/api/v1/nodes/get_phone',
        method: 'get',
        params: { host }
    })
}


export const getPhoneInfo = (data: GetPhoneInfo) => {
    return request({
        url: '/api/v1/nodes/phone_info',
        method: 'post',
        data: data
    })
}
