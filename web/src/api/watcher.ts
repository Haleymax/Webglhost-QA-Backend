import request from '@/utils/request'

export interface Watcher {
  id: string | null
  name: string
  resource: string
  click: string
  brand: string[]
  tag: string[]
}


export const getWatchers = () => {
  return request({
    url: '/api/v1/watchers/find',
    method: 'get'
  })
}

export const addWatcher = (data: Watcher) => {
  return request({
    url: '/api/v1/watchers/add',
    method: 'post',
    data
  })
}


export const deleteWatcher = (data: Watcher) => {
  return request({
    url: '/api/v1/watchers/delete',
    method: 'delete',
    data
  })
}



export const updateWatcher = (data: Watcher) => {
  return request({
    url: '/api/v1/watchers/update',
    method: 'put',
    data
  })
}


export const updateRedis = (data: any) => {
  return request({
    url: '/api/v1/watchers/refresh',
    method: 'post',
    data
  })
}
