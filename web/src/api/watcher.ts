import request from '@/utils/request'

export interface Watcher {
  id: string
  name: string
  resource: string
  click: string
  brand: string[]
  tag: string[]
}


export const getWatchers = () => {
  return request({
    url: '/api/watchers/find',
    method: 'get'
  })
}
